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
	crudName := "VirtP-a"
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
		t.Errorf("creating vp : %v", err)
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
		t.Errorf("creating vp : %v", err)
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
	element1, err := createTestVirtualSaml2ServiceProviderDTO("VirtP-1")
	if err != nil {
		t.Error(err)
		return
	}
	listOfCreated[0], _ = s.client.CreateVirtSaml2Sp(*appliance.Name, *element1)

	element2, err := createTestVirtualSaml2ServiceProviderDTO("VirtP-2")
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
	tData.SetDisplayName(fmt.Sprintf("Virtual Provider %s", name))

	var snip api.SubjectNameIdentifierPolicyDTO

	snip.SetDescriptionKey("")
	snip.SetName("Principal")
	snip.SetSubjectAttribute("")
	snip.SetType("PRINCIPAL")
	snip.AdditionalProperties = make(map[string]interface{})
	snip.AdditionalProperties["@c"] = "com.atricore.idbus.console.services.dto.SubjectNameIdentifierPolicyDTO"
	tData.SetSubjectNameIDPolicy(snip)

	rs := api.NewResourceDTOInit("test-ks", "test ks", keystore)
	rs.SetUri(fmt.Sprintf("ks-%s.jks", name))

	// Use a p12 file
	ks := api.NewKeystoreDTOInit(fmt.Sprintf("%s-ks", name), fmt.Sprintf("%s keystore", name), rs)
	ks.SetCertificateAlias("jetty")
	ks.SetPassword("@WSX3edc")
	ks.SetPrivateKeyName("jetty")
	ks.SetPrivateKeyPassword("@WSX3edc")
	ks.SetType("PKCS#12")

	// Inject in IdP using IDP Config
	idpCfg := api.NewSamlR2IDPConfigDTO()
	idpCfg.SetName(fmt.Sprintf("%s-cfg", name))
	idpCfg.SetUseSampleStore(false)
	idpCfg.SetUseSystemStore(false)
	idpCfg.SetSigner(*ks)
	idpCfg.SetEncrypter(*ks)
	cfg, _ := idpCfg.ToProviderConfig()
	tData.SetConfig(*cfg)

	// SAML2 IdP config as serialized by CXF (Additional properties)
	var conf api.SamlR2IDPConfigDTO
	conf.SetUseSampleStore(false)
	conf.SetUseSystemStore(false)
	err := tData.SetSamlR2IDPConfig(&conf)
	if err != nil {
		return nil, err
	}

	var idMapping api.IdentityMappingPolicyDTO
	idMapping.SetMappingType("LOCAL")
	idMapping.SetUseLocalId(true)
	tData.SetIdentityMappingPolicy(idMapping)

	var linkage api.AccountLinkagePolicyDTO
	linkage.SetLinkEmitterType("EMAIL")
	tData.SetAccountLinkagePolicy(linkage)

	tData.SetDashboardUrl("http://my-dashbaord.mycompany.com/ui")
	tData.SetEnableMetadataEndpoint(true)
	tData.SetEnableProxyExtension(true)

	tData.SetEncryptAssertion(true)
	tData.SetEncryptAssertionAlgorithm("http://www.w3.org/200|/04/xmlenc#aes128-cbc")
	tData.SetErrorBinding("JSON")
	tData.SetIdpSignatureHash("SHA256")

	tData.SetIgnoreRequestedNameIDPolicy(true)
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
		// {
		// 	name:     "account_linkage_policy",
		// 	cmp:      func() bool { return StrPtrEquals(e.AccountLinkagePolicy, r.AccountLinkagePolicy) },
		// 	expected: StrDeref(e.AccountLinkagePolicy),
		// 	received: StrDeref(r.AccountLinkagePolicy),
		// },
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "dashboard_url",
			cmp:      func() bool { return StrPtrEquals(e.DashboardUrl, r.DashboardUrl) },
			expected: StrDeref(e.DashboardUrl),
			received: StrDeref(r.DashboardUrl),
		},
		{
			name:     "display_Name",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: StrDeref(e.DisplayName),
			received: StrDeref(r.DisplayName),
		},
		{
			name:     "enable_metadata_endpoint",
			cmp:      func() bool { return BoolPtrEquals(e.EnableMetadataEndpoint, r.EnableMetadataEndpoint) },
			expected: strconv.FormatBool(BoolDeref(e.EnableMetadataEndpoint)),
			received: strconv.FormatBool(BoolDeref(r.EnableMetadataEndpoint)),
		},
		{
			name:     "enable_proxy_extension",
			cmp:      func() bool { return BoolPtrEquals(e.EnableProxyExtension, r.EnableProxyExtension) },
			expected: strconv.FormatBool(BoolDeref(e.EnableProxyExtension)),
			received: strconv.FormatBool(BoolDeref(r.EnableProxyExtension)),
		},
		{
			name:     "encrypt_assertion",
			cmp:      func() bool { return BoolPtrEquals(e.EncryptAssertion, r.EncryptAssertion) },
			expected: strconv.FormatBool(BoolDeref(e.EncryptAssertion)),
			received: strconv.FormatBool(BoolDeref(r.EncryptAssertion)),
		},
		{
			name:     "encrypt_assertion_algorithm",
			cmp:      func() bool { return StrPtrEquals(e.EncryptAssertionAlgorithm, r.EncryptAssertionAlgorithm) },
			expected: StrDeref(e.EncryptAssertionAlgorithm),
			received: StrDeref(r.EncryptAssertionAlgorithm),
		},
		{
			name:     "error_binding",
			cmp:      func() bool { return StrPtrEquals(e.ErrorBinding, r.ErrorBinding) },
			expected: StrDeref(e.ErrorBinding),
			received: StrDeref(r.ErrorBinding),
		},
		{
			name:     "idp_signature_hash",
			cmp:      func() bool { return StrPtrEquals(e.IdpSignatureHash, r.IdpSignatureHash) },
			expected: StrDeref(e.IdpSignatureHash),
			received: StrDeref(r.IdpSignatureHash),
		},
		{
			name:     "ignore_requested_name_id_policy",
			cmp:      func() bool { return BoolPtrEquals(e.IgnoreRequestedNameIDPolicy, r.IgnoreRequestedNameIDPolicy) },
			expected: strconv.FormatBool(BoolDeref(e.IgnoreRequestedNameIDPolicy)),
			received: strconv.FormatBool(BoolDeref(r.IgnoreRequestedNameIDPolicy)),
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
			name:     "oauth_2_enabled",
			cmp:      func() bool { return BoolPtrEquals(e.Oauth2Enabled, r.Oauth2Enabled) },
			expected: strconv.FormatBool(BoolDeref(e.Oauth2Enabled)),
			received: strconv.FormatBool(BoolDeref(r.Oauth2Enabled)),
		},
		{
			name:     "oauth_2_key",
			cmp:      func() bool { return StrPtrEquals(e.Oauth2Key, r.Oauth2Key) },
			expected: StrDeref(e.Oauth2Key),
			received: StrDeref(r.Oauth2Key),
		},
		{
			name:     "oauth_2_remember_me_token_validity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2RememberMeTokenValidity, r.Oauth2RememberMeTokenValidity) },
			expected: strconv.FormatInt(Int64Deref(e.Oauth2RememberMeTokenValidity), 10),
			received: strconv.FormatInt(Int64Deref(r.Oauth2RememberMeTokenValidity), 10),
		},
		{
			name:     "oauth_2_token_validity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2TokenValidity, r.Oauth2TokenValidity) },
			expected: strconv.FormatInt(Int64Deref(e.Oauth2TokenValidity), 10),
			received: strconv.FormatInt(Int64Deref(r.Oauth2TokenValidity), 10),
		},
		{
			name:     "oidc_access_token_time_to_live",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAccessTokenTimeToLive, r.OidcAccessTokenTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcAccessTokenTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcAccessTokenTimeToLive))),
		},
		{
			name:     "oidc_authz_code_time_to_live",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAuthzCodeTimeToLive, r.OidcAuthzCodeTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcAuthzCodeTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcAuthzCodeTimeToLive))),
		},
		{
			name:     "oidc_id_token_time_to_live",
			cmp:      func() bool { return Int32PtrEquals(e.OidcIdTokenTimeToLive, r.OidcIdTokenTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcIdTokenTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcIdTokenTimeToLive))),
		},
		{
			name:     "open_id_enabled",
			cmp:      func() bool { return BoolPtrEquals(e.OpenIdEnabled, r.OpenIdEnabled) },
			expected: strconv.FormatBool(BoolDeref(e.OpenIdEnabled)),
			received: strconv.FormatBool(BoolDeref(r.OpenIdEnabled)),
		},
		{
			name:     "sign_authentication_requests",
			cmp:      func() bool { return BoolPtrEquals(e.SignAuthenticationRequests, r.SignAuthenticationRequests) },
			expected: strconv.FormatBool(BoolDeref(e.SignAuthenticationRequests)),
			received: strconv.FormatBool(BoolDeref(r.SignAuthenticationRequests)),
		},
		{
			name:     "sign_requests",
			cmp:      func() bool { return BoolPtrEquals(e.SignRequests, r.SignRequests) },
			expected: strconv.FormatBool(BoolDeref(e.SignRequests)),
			received: strconv.FormatBool(BoolDeref(r.SignRequests)),
		},
		{
			name:     "sp_signature_hash",
			cmp:      func() bool { return StrPtrEquals(e.SpSignatureHash, r.SpSignatureHash) },
			expected: StrDeref(e.SpSignatureHash),
			received: StrDeref(r.SpSignatureHash),
		},
		{
			name:     "sso_session_timeout",
			cmp:      func() bool { return Int32PtrEquals(e.SsoSessionTimeout, r.SsoSessionTimeout) },
			expected: strconv.Itoa(int(Int32Deref(e.SsoSessionTimeout))),
			received: strconv.Itoa(int(Int32Deref(r.SsoSessionTimeout))),
		},
		{
			name:     "want_assertion_signed",
			cmp:      func() bool { return BoolPtrEquals(e.WantAssertionSigned, r.WantAssertionSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantAssertionSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantAssertionSigned)),
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
			name:     "want_signed_requests",
			cmp:      func() bool { return BoolPtrEquals(e.WantSignedRequests, r.WantSignedRequests) },
			expected: strconv.FormatBool(BoolDeref(e.WantSignedRequests)),
			received: strconv.FormatBool(BoolDeref(r.WantSignedRequests)),
		},
	}
}

//Fields to validate after VirtualSaml2Sp update
func VirtualSaml2SpFieldTestUpdate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{}
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
