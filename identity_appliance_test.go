package cli

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

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

// Removes unsupported chars from name
func sanitizeName(name string) string {
	// Replace unsupported chars

	chars := []string{"]", "^", "\\\\", "[", "(", ")", "-"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	name = re.ReplaceAllString(name, "")

	return strings.ToLower(name)
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
