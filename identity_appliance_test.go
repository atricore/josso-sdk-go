package cli

import (
	"fmt"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdentityAppliance_import() {

	var t = s.T()

	a, err := s.ImportAppliance("./samples/testacc-01.idmn")
	if err != nil {
		t.Errorf("importing test appliance : %v", err)
		return
	}
	t.Logf("Imported appliance %s [%d]", *a.Name, *a.Id)

}

func (s *AccTestSuite) TestAccCliIdentityAppliance_crud() {
	var t = s.T()

	s.accClearData()

	// Test CRUD
	crudName := "testacc-z"
	var orig *api.IdentityApplianceDefinitionDTO
	var created api.IdentityApplianceDefinitionDTO
	orig = createTestIdentityApplianceDefinitionDTO(crudName)
	// Create

	created, err := s.client.CreateAppliance(*orig)
	if err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}
	if err := IdApplianceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}

	// Retrieve
	retrieved, err := s.client.GetAppliance(crudName)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&created, &retrieved); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Update
	retrieved.Namespace = api.PtrString("com.atricore.ida.a.mod")
	updated, err := s.client.UpdateAppliance(retrieved)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&retrieved, &updated); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Delete
	removed, err := s.client.DeleteAppliance(strconv.FormatInt(*updated.Id, 10))
	if err != nil {
		t.Errorf("deleting identity appilance : %v", err)
		return
	}
	if !removed {
		t.Errorf("deleting identity appliance : not found %d", *updated.Id)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetAppliances()
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name, (these elements must have all the variables of the structure)
	var listOfCreated [2]api.IdentityApplianceDefinitionDTO

	element1 := createTestIdentityApplianceDefinitionDTO("testacc-1")
	listOfCreated[0], _ = s.client.CreateAppliance(*element1)

	element2 := createTestIdentityApplianceDefinitionDTO("testacc-2")
	listOfCreated[1], _ = s.client.CreateAppliance(*element2)

	// Get list from server
	listOfRead, err := s.client.GetAppliances()
	if err != nil {
		t.Error(err)
		return
	}
	// The list should have 2 elemetns
	if len(listOfRead) != 2 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expected 2", len(listOfAll))
		return
	}
}

func (s *AccTestSuite) TestAccCliIdentityAppliance_z010() {
	var t = s.T()
	t.Log("Acceptance test z010 : basic appliance")

	a, err := s.ImportAppliance("./samples/testacc-01.idmn")
	if err != nil {
		t.Errorf("z010, importing appliance : %v ", err)
		return
	}

	sp, err := s.client.GetExtSaml2Sp(*a.Name, "sp-1")
	if err != nil {
		t.Errorf("z010, getting sp : %v ", err)
		return
	}

	if sp.GetName() != "sp-1" {
		t.Errorf("z010, unexpected sp name %s", sp.GetName())
		return
	}

	fcsBLen := len(sp.GetFederatedConnectionsB())
	if fcsBLen != 1 {
		t.Errorf("z010, unexpedted number of federated connections B, got %d", fcsBLen)
		return
	}

	for _, fcB := range sp.GetFederatedConnectionsB() {
		t.Logf("Federated Connection %s", fcB.GetName())

		idpC, err := fcB.GetIDPChannel()
		if err != nil {
			t.Errorf("%v", err)
		}
		spC, err := fcB.GetSPChannel()
		if err != nil {
			t.Errorf("%v", err)
		}

		t.Logf("%#v", fcB)
		t.Logf("%#v", idpC)
		t.Logf("%#v", spC)
	}

}

// Simple appliance test:
//  - one idp
//     - basic authn
//  - db identity source
//  - external saml2 sp
//  - partnerapp tomcat
//
func (s *AccTestSuite) TestAccCliIdentityAppliance_z020() {
	var t = s.T()
	t.Log("Acceptance test #020 : basic idp")

}

// Simple appliance test:
//  - one idp
//     - basic authn
//  - identity vault
//  - external saml2 sp
//
func (s *AccTestSuite) TestAccCliIdentityAppliance_z030() {
	var t = s.T()
	t.Log("Acceptance test #030 : basic idp")

	// 0. Create test appliance
	var template api.IdentityApplianceDefinitionDTO

	appliance, err := s.client.CreateAppliance(template)
	// 1. Create identity vault
	var origIdVauld *api.EmbeddedIdentityVaultDTO
	var createdIdVaul api.EmbeddedIdentityVaultDTO

	createdIdVaul, err = s.client.CreateIdVault("idVault", *origIdVauld)
	if err != nil {
		t.Errorf("create identity appliance: %v", err)
		return
	}
	if err := IdVaultValidateCreate(origIdVauld, &createdIdVaul); err != nil {
		t.Errorf("creating idVault : %v", err)
		return
	}

	// 2. Create IdP using DB identity vault
	var authns []api.AuthenticationMechanismDTO
	authns = append(authns, createTestBasicAuthn())
	idp, err := createTestIdentityProviderDTO("idp-1", authns)
	if err != nil {
		t.Errorf("Create identity appliance : %v", err)
		return
	}

	*idp, err = s.client.CreateIdp(appliance.GetName(), *idp)

	idp.AddIdentityLookup("id-lookup-1")

	// 3. Create external SAML 2 sp, using test metadata and connect it to the IdP
	var origsaml2 *api.EmbeddedIdentityVaultDTO
	var createdsaml2 api.EmbeddedIdentityVaultDTO

	createdsaml2, err = s.client.CreateIdVault("Saml2Sp", *origsaml2)
	if err != nil {
		t.Errorf("create Saml2Sp: %v", err)
		return
	}
	if err := IdVaultValidateCreate(origsaml2, &createdsaml2); err != nil {
		t.Errorf("creating Saml2Sp : %v", err)
		return
	}
	// 4. Build/Start identity appliance

	// s, err := s.client.SetApplianceState("STARTED")

}

// ---------------------------------------------------------

func createTestIdentityApplianceDefinitionDTO(name string) *api.IdentityApplianceDefinitionDTO {
	tData := api.NewIdentityApplianceDefinitionDTO()

	var locat api.LocationDTO
	locat.SetContext("IDBUS")
	locat.SetHost("localhost")
	locat.SetPort(8081)
	locat.SetProtocol("http")
	locat.SetUri(strings.ToUpper(name))
	tData.SetLocation(locat)

	var branding api.UserDashboardBrandingDTO
	branding.SetName("josso25-branding")
	tData.SetUserDashboardBranding(branding)

	var scfg api.IdentityApplianceSecurityConfigDTO
	scfg.SetEncryptSensitiveData(false)
	//scfg.SetEncryption("AES128-CBC/SHA256")
	//scfg.SetEncryptionConfig("")
	scfg.SetEncryptionConfigFile(fmt.Sprintf("cfg-%s.properties", name))
	scfg.SetEncryptionPassword(fmt.Sprintf("pwd%s", name))
	scfg.SetExternalConfig(false)
	// scfg.SetExternalConfigFile(fmt.Sprintf("sec-%s.properties", name))
	// scfg.SetPasswordProperty("")
	// scfg.SetSalt("")
	// scfg.SetSaltProperty("")
	// scfg.SetSaltValue("")
	tData.SetSecurityConfig(scfg)

	var es api.EntitySelectionStrategyDTO
	es.SetName("requested-preferred-idp-selection")
	tData.SetIdpSelector(es)

	tData.SetDescription(fmt.Sprintf("Desc %s", name))
	tData.SetDisplayName(fmt.Sprintf("DispName %s", name))
	tData.SetName(name)
	tData.SetNamespace(fmt.Sprintf("com.atricore.idbus.ida.%s", sanitizeName(name)))
	// orig.SetRequiredBundles() //

	return tData
}

// -------------------------------------------------

//Fields to validate after appliance creation
func ApplianceFieldTestCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

	return []FiledTestStruct{
		// {
		// 	name:     "userdashboardbranding",
		// 	cmp:      func() bool { return UserBrandingPtrEquals(e.UserDashboardBranding, r.UserDashboardBranding) },
		// 	expected: StrDeref(e.UserDashboardBranding),
		// 	received: StrDeref(r.UserDashboardBranding),
		// },
		// {
		// 	name:     "securityconfig",
		// 	cmp:      func() bool { return IDASCPtrEquals(e.SecurityConfig, r.SecurityConfig) },
		// 	expected: StrDeref(e.SecurityConfig),
		// 	received: StrDeref(r.SecurityConfig),
		// },
		// {
		// 	name:     "keystore",
		// 	cmp:      func() bool { return KeystorePtrEquals(e.Keystore, r.Keystore) },
		// 	expected: StrDeref(e.Keystore),
		// 	received: StrDeref(r.Keystore),
		// },
		// {
		// 	name:     "idpselector",
		// 	cmp:      func() bool { return EntitySelectionPtrEquals(e.IdpSelector, r.IdpSelector) },
		// 	expected: StrDeref(e.IdpSelector),
		// 	received: StrDeref(r.IdpSelector),
		// },
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "displayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: StrDeref(e.DisplayName),
			received: StrDeref(r.DisplayName),
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "namespace",
			cmp:      func() bool { return StrPtrEquals(e.Namespace, r.Namespace) },
			expected: StrDeref(e.Namespace),
			received: StrDeref(r.Namespace),
		},
	}
}

//Fields to validate after appliance update
func ApplianceFieldTestUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
	}
	return append(t, ApplianceFieldTestCreate(e, r)...)
}

// Compares the expected appliance with the received one.
func IdApplianceValidateCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestCreate(e, r))
}

// Compares the expected appliance with the received one.
func IdApplianceValidateUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestUpdate(e, r))
}
