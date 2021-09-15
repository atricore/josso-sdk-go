package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliVirtSaml2_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Virtsaml2-a"
	var orig *api.VirtualSaml2ServiceProviderDTO
	var created api.VirtualSaml2ServiceProviderDTO
	orig, err = createTestVirtualSaml2ServiceProviderDTO(crudName)
	if err != nil {
		t.Error(err)
		return
	}

	// Test CREATE
	created, err = s.client.CreateVirtSaml2Sp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := VirtualSaml2SpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.VirtualSaml2ServiceProviderDTO
	read, err = s.client.GetVirtSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = VirtualSaml2SpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateVirtSaml2Sp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = VirtualSaml2SpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteVirtSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// ------------------------------------------------------------------------------------------------------------------
	// Test empty list

	listOfAll, err := s.client.GetVirtSaml2Sps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// ------------------------
	// List of created elements, order by Name
	var listOfCreated [2]api.VirtualSaml2ServiceProviderDTO

	// Test list of #2 elements
	element1, err := createTestVirtualSaml2ServiceProviderDTO("Extsmal2-1")
	if err != nil {
		t.Error(err)
		return
	}
	listOfCreated[0], _ = s.client.CreateVirtSaml2Sp(*appliance.Name, *element1)

	element2, err := createTestVirtualSaml2ServiceProviderDTO("Extsmal2-2")
	if err != nil {
		t.Error(err)
		return
	}
	listOfCreated[1], _ = s.client.CreateVirtSaml2Sp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetVirtSaml2Sps(*appliance.Name)
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

	// Order list of read by Name
	sort.SliceStable(listOfRead,
		func(i, j int) bool {
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) < 0
		},
	)

	// Validate each element from the list of created with the list of read
	for idx, r := range listOfCreated {
		if err = VirtualSaml2SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestVirtualSaml2ServiceProviderDTO(name string) (*api.VirtualSaml2ServiceProviderDTO, error) {
	//encMetadata := metadata
	tData := api.NewVirtualSaml2ServiceProviderDTO()

	tData.SetName(name)

	var snip api.SubjectNameIdentifierPolicyDTO

	snip.SetDescriptionKey("")
	snip.SetName("Principal")
	snip.SetSubjectAttribute("")
	snip.SetType("PRINCIPAL")
	snip.AdditionalProperties = make(map[string]interface{})
	snip.AdditionalProperties["@c"] = "com.atricore.idbus.console.services.dto.SubjectNameIdentifierPolicyDTO"
	tData.SetSubjectNameIDPolicy(snip)

	var rs api.ResourceDTO
	rs.SetValue(keystore)
	rs.SetUri(fmt.Sprintf("ks-%s.jks", name))

	var ks api.KeystoreDTO
	ks.SetCertificateAlias("jetty")
	ks.SetPassword("@WSX3edc")
	ks.SetPrivateKeyName("jetty")
	ks.SetPrivateKeyPassword("@WSX3edc")
	ks.SetStore(rs)
	ks.SetType("JKS")
	ks.SetName(fmt.Sprintf("%s-ks", name))
	ks.SetStore(rs)
	// TODO : Inject in VP

	var IdentityMappingPolicyDTO api.IdentityMappingPolicyDTO
	IdentityMappingPolicyDTO.SetMappingType("LOCAL")
	IdentityMappingPolicyDTO.SetUseLocalId(true)
	tData.SetIdentityMappingPolicy(IdentityMappingPolicyDTO)

	/*
		var AttributeProfileDTO api.AttributeProfileDTO
		AttributeProfileDTO.SetProfileType("")
		tData.SetAttributeProfile(AttributeProfileDTO)
	*/

	var conf api.SamlR2IDPConfigDTO
	conf.SetDescription("")
	conf.SetDisplayName("")
	conf.SetElementId("")
	conf.SetId(19)
	conf.SetName("")
	conf.SetUseSampleStore(true)
	conf.SetUseSystemStore(false)
	idpConf, err := FromIdPConfig(&conf)
	if err != nil {
		return nil, err
	}
	tData.SetConfig(idpConf)

	var AccountLinkagePolicyDTO api.AccountLinkagePolicyDTO
	AccountLinkagePolicyDTO.SetLinkEmitterType("EMAIL")
	tData.SetAccountLinkagePolicy(AccountLinkagePolicyDTO)

	tData.SetDashboardUrl("http://my-dashbaord.mycompany.com/ui")
	tData.SetDisplayName("")
	tData.SetEnableMetadataEndpoint(true)
	tData.SetEnableProxyExtension(true)

	tData.SetEncryptAssertion(true)
	tData.SetEncryptAssertionAlgorithm("http://www.w3.org/200|/04/xmlenc#aes128-cbc")
	tData.SetErrorBinding("JSON")
	tData.SetIdpSignatureHash("SHA256")

	tData.SetIgnoreRequestedNameIDPolicy(true)
	tData.SetIsRemote(true)
	tData.SetMessageTtl(300)
	tData.SetMessageTtlTolerance(300)
	tData.SetOauth2Enabled(true)
	tData.SetOauth2Key(fmt.Sprintf("%s-oauth-key", name))
	tData.SetOauth2RememberMeTokenValidity(43201)
	tData.SetOauth2TokenValidity(303)
	tData.SetOidcAccessTokenTimeToLive(3610)
	tData.SetOidcAuthzCodeTimeToLive(305)
	tData.SetOidcIdTokenTimeToLive(3620)
	tData.SetOpenIdEnabled(true)
	tData.SetSignAuthenticationRequests(true)
	tData.SetSignRequests(true)
	tData.SetSpSignatureHash("SHA256")
	tData.SetSsoSessionTimeout(1)
	tData.SetWantAssertionSigned(true)
	tData.SetWantAuthnRequestsSigned(true)
	tData.SetWantSLOResponseSigned(true)
	tData.SetWantSignedRequests(true)

	tData.SetDescription(fmt.Sprintf("%s descr", name))

	return tData, nil
}

func (s *AccTestSuite) TestAccCliVirSaml2_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliVirSaml2_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func VirtualSaml2SpFieldTestCreate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
	}
}

//Fields to validate after VirtualSaml2Sp update
func VirtualSaml2SpFieldTestUpdate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}
	return append(t, VirtualSaml2SpFieldTestCreate(e, r)...)
}

// Compares the expected VirtualSaml2Sp with the received one.
func VirtualSaml2SpValidateCreate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) error {

	return ValidateFields(VirtualSaml2SpFieldTestCreate(e, r))
}

// Compares the expected ExternalSaml2Sp with the received one.
func VirtualSaml2SpValidateUpdate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) error {

	return ValidateFields(VirtualSaml2SpFieldTestUpdate(e, r))
}
