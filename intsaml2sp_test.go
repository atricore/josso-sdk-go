package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIntSaml2_crud() { // MODIFIED
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Intsmal2-a"
	var orig *api.InternalSaml2ServiceProviderDTO
	var created api.InternalSaml2ServiceProviderDTO
	orig = createTestInternalSaml2ServiceProviderDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIntSaml2Sp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := InteralSaml2SpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.InternalSaml2ServiceProviderDTO
	read, err = s.client.GetIntSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = InternalSaml2SpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateIntSaml2Sp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = InternalSaml2SpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIntSaml2Sp(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetIntSaml2Sps(*appliance.Name)
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
	var listOfCreated [2]api.InternalSaml2ServiceProviderDTO

	// Test list of #2 elements
	element1 := createTestInternalSaml2ServiceProviderDTO("Intsmal2-1")
	listOfCreated[0], _ = s.client.CreateIntSaml2Sp(*appliance.Name, *element1)

	element2 := createTestInternalSaml2ServiceProviderDTO("Intsmal2-2")
	listOfCreated[1], _ = s.client.CreateIntSaml2Sp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetIntSaml2Sps(*appliance.Name)
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
		if err = InternalSaml2SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestInternalSaml2ServiceProviderDTO(name string) *api.InternalSaml2ServiceProviderDTO {

	tData := api.NewInternalSaml2ServiceProviderDTO()

	tData.SetDashboardUrl(fmt.Sprintf("https://my-page-%s.com", name))
	tData.SetDescription("IntSaml2Sp One")
	tData.SetDisplayName("")
	tData.SetEnableMetadataEndpoint(true)
	tData.SetErrorBinding("JSON")
	tData.SetId(47)
	tData.SetIsRemote(true)
	tData.SetMessageTtl(300)
	tData.SetMessageTtlTolerance(300)
	tData.SetName(name)
	tData.SetRemote(true)
	tData.SetSignAuthenticationRequests(true)
	tData.SetSignRequests(true)
	tData.SetSignatureHash("SHA256")
	tData.SetWantAssertionSigned(false)
	tData.SetWantSLOResponseSigned(false)
	tData.SetWantAssertionSigned(false)

	var accountLink api.AccountLinkagePolicyDTO
	accountLink.SetName("One To One")
	accountLink.SetLinkEmitterType("ONE_TO_ONE") // "ONE_TO_ONE", EMAIL, CUSTOM, UID
	tData.SetAccountLinkagePolicy(accountLink)

	var idMapping api.IdentityMappingPolicyDTO
	idMapping.SetName("Aggregate")
	idMapping.SetMappingType("MERGED") //  LOCAL("Use Ours"), REMOTE("Use Theirs"), default: MERGED("Aggregate"), CUSTOM("Custom");
	idMapping.SetUseLocalId(false)

	var rs api.ResourceDTO
	rs.SetValue(keystore)
	rs.SetUri(fmt.Sprintf("ks-%s.jks", name))

	ks := api.NewKeystoreDTOWithOK()
	ks.SetCertificateAlias("jetty")
	ks.SetPassword("@WSX3edc")
	ks.SetPrivateKeyName("jetty")
	ks.SetPrivateKeyPassword("@WSX3edc")
	ks.SetStore(rs)
	ks.SetType("JKS")
	ks.SetName(fmt.Sprintf("%s-ks", name))
	ks.SetStore(rs)

	var spCfg api.SamlR2SPConfigDTO
	spCfg.SetUseSampleStore(false)
	spCfg.SetUseSystemStore(false)
	spCfg.SetSigner(*ks)
	cfg, _ := spCfg.ToProviderConfig()
	tData.SetConfig(*cfg)

	tData.SetIdentityMappingPolicy(idMapping)

	return tData
}

func (s *AccTestSuite) TestAccCliIntSaml2_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIntSaml2_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func InternalSaml2SpFieldTestCreate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "dashboard_Url",
			cmp:      func() bool { return StrPtrEquals(e.DashboardUrl, r.DashboardUrl) },
			expected: StrDeref(e.DashboardUrl),
			received: StrDeref(r.DashboardUrl),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "enable_metadata_endpoint",
			cmp:      func() bool { return BoolPtrEquals(e.EnableMetadataEndpoint, r.EnableMetadataEndpoint) },
			expected: strconv.FormatBool(BoolDeref(e.EnableMetadataEndpoint)),
			received: strconv.FormatBool(BoolDeref(r.EnableMetadataEndpoint)),
		},
		{
			name:     "error_binding",
			cmp:      func() bool { return StrPtrEquals(e.ErrorBinding, r.ErrorBinding) },
			expected: StrDeref(e.ErrorBinding),
			received: StrDeref(r.ErrorBinding),
		},
		{
			name:     "message_ttl",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtl, r.MessageTtl) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtl))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtl))),
		},
		{
			name:     "message_ttl_tolerance",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtlTolerance, r.MessageTtlTolerance) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtlTolerance))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtlTolerance))),
		},
		{
			name:     "sign_requests",
			cmp:      func() bool { return BoolPtrEquals(e.SignRequests, r.SignRequests) },
			expected: strconv.FormatBool(BoolDeref(e.SignRequests)),
			received: strconv.FormatBool(BoolDeref(r.SignRequests)),
		},
		{
			name:     "signature_hash",
			cmp:      func() bool { return StrPtrEquals(e.SignatureHash, r.SignatureHash) },
			expected: StrDeref(e.SignatureHash),
			received: StrDeref(r.SignatureHash),
		},
		{
			name:     "want_assertion_signed",
			cmp:      func() bool { return BoolPtrEquals(e.WantAssertionSigned, r.WantAssertionSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantAssertionSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantAssertionSigned)),
		},
		{
			name:     "want_slo_response_signed",
			cmp:      func() bool { return BoolPtrEquals(e.WantSLOResponseSigned, r.WantSLOResponseSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantSLOResponseSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantSLOResponseSigned)),
		},
		{
			name:     "want_assertion_signed",
			cmp:      func() bool { return BoolPtrEquals(e.WantAssertionSigned, r.WantAssertionSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantAssertionSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantAssertionSigned)),
		},
	}
}

//Fields to validate after Sp update
func InteralSaml2SpFieldTestUpdate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
		{
			name:     "element_id",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
	}
	return append(t, InternalSaml2SpFieldTestCreate(e, r)...)
}

// Compares the expected Sp with the received one.
func InteralSaml2SpValidateCreate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) error {

	return ValidateFields(InternalSaml2SpFieldTestCreate(e, r))
}

// Compares the expected Sp with the received one.
func InternalSaml2SpValidateUpdate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) error {

	return ValidateFields(InteralSaml2SpFieldTestUpdate(e, r))
}
