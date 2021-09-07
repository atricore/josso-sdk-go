package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdP_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "idp-1"
	var orig *api.IdentityProviderDTO
	var created api.IdentityProviderDTO

	var authn []api.AuthenticationMechanismDTO
	authn = append(authn, createTestBasicAuthn())
	orig, err = createTestIdentityProviderDTO(crudName, authn)
	if err != nil {
		t.Error(err)
		return
	}

	// Test CREATE
	created, err = s.client.CreateIdp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdPValidateCreate(orig, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}

	// Test READ
	var read api.IdentityProviderDTO
	read, err = s.client.GetIdp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if read.Name == nil {
		t.Errorf("IdP not found for name %s", crudName)
		return
	}
	if err = IdPValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DashboardUrl = api.PtrString("12345")
	read.DisplayName = api.PtrString("null")
	updated, err := s.client.UpdateIdp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdPValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIdp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdps(*appliance.Name)
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
	var listOfCreated [2]api.IdentityProviderDTO
	// Test list of #2 elements

	// Idp - 1
	var authns1 []api.AuthenticationMechanismDTO
	authns1 = append(authns1, createTestBasicAuthn())
	element1, err := createTestIdentityProviderDTO("ids-1", authns1)
	if err != nil {
		t.Error(err)
		return
	}
	listOfCreated[0], _ = s.client.CreateIdp(*appliance.Name, *element1)

	// Idp - 2
	var authns2 []api.AuthenticationMechanismDTO
	authns2 = append(authns2, createTestBasicAuthn())
	element2, err := createTestIdentityProviderDTO("ids-2", authns2)
	if err != nil {
		t.Error(err)
		return
	}
	listOfCreated[1], _ = s.client.CreateIdp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetIdps(*appliance.Name)
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
		if err = IdPValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

// Creates an new AuthenticationMechanisDTO for basic authentication.
func createTestBasicAuthn() api.AuthenticationMechanismDTO {
	var authn api.AuthenticationMechanismDTO

	authn.SetName("idp-basic-authn")
	authn.SetDisplayName("idp-basic-authn")
	authn.SetPriority(1)

	authn.AdditionalProperties = make(map[string]interface{})

	authn.AdditionalProperties["@c"] = ".BasicAuthenticationDTO"
	authn.AdditionalProperties["hashAlgorithm"] = "SHA-512"
	authn.AdditionalProperties["hashEncoding"] = "BASE64"
	authn.AdditionalProperties["ignoreUsernamecase"] = false
	authn.AdditionalProperties["ignorePassowordCase"] = false
	authn.AdditionalProperties["SaltLength"] = 0
	authn.AdditionalProperties["saltPrefix"] = ""
	authn.AdditionalProperties["saltSuffix"] = ""
	//authn.AdditionalProperties["impersonateUserPolicy"]
	authn.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"] = "urn:oasis:names:tc:SAML:2.0:ac:classes:Password"

	return authn
}

// Creates an new AuthenticationMechanisDTO for two-factor authentication.
func createTest2FactorAuthn() api.AuthenticationMechanismDTO {
	var authn api.AuthenticationMechanismDTO

	authn.SetDelegatedAuthentication(*authn.DelegatedAuthentication)
	authn.SetDisplayName("")
	authn.SetName("idp-2factor-authn")
	authn.SetPriority(1)

	authn.AdditionalProperties[""] = 0

	authn.AdditionalProperties["@c"] = ".TwoFactorAuthenticationDTO"

	return authn
}

// Receives an array of authentication mechanisms.  An authneticaiton mehcanism may or may not have a Delegated authentication.
// All delegated authentications must be used : idp.SetDelegatedAuthentications.  Mechanisms go into idp.AuthenticationMechanisms
func createTestIdentityProviderDTO(name string, authn []api.AuthenticationMechanismDTO) (*api.IdentityProviderDTO, error) {

	var AuthenticationAssertionEmissionPolicyDTO api.AuthenticationAssertionEmissionPolicyDTO
	var fedconn []api.FederatedConnectionDTO
	orig := api.NewIdentityProviderDTO()
	var ResourceDTO api.ResourceDTO
	var identityAppliance api.IdentityApplianceDefinitionDTO
	var EntitySelectionStrategyDTO api.EntitySelectionStrategyDTO
	var UserDashboardBrandingDTO api.UserDashboardBrandingDTO
	var ExtensionDTO api.ExtensionDTO
	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO

	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")

	// TODO : Set valid name/classifier
	// ExtensionDTO.SetClassifier("")
	// ExtensionDTO.SetName("")

	UserDashboardBrandingDTO.SetId("")
	UserDashboardBrandingDTO.SetName("")

	// TODO : Use valid name
	EntitySelectionStrategyDTO.SetDescription("")
	EntitySelectionStrategyDTO.SetName("")

	// AuthenticationAssertionEmissionPolicyDTO.SetElementId("")
	// AuthenticationAssertionEmissionPolicyDTO.SetId(1)
	// AuthenticationAssertionEmissionPolicyDTO.SetName("")
	var snip api.SubjectNameIdentifierPolicyDTO

	snip.SetDescriptionKey("")
	snip.SetName("Principal")
	snip.SetSubjectAttribute("")
	snip.SetType("PRINCIPAL")
	snip.AdditionalProperties = make(map[string]interface{})
	snip.AdditionalProperties["@c"] = "com.atricore.idbus.console.services.dto.SubjectNameIdentifierPolicyDTO"
	orig.SetSubjectNameIDPolicy(snip)

	// TODO : Use valid names(Share)
	var key api.KeystoreDTO
	key.SetCertificateAlias("")
	key.SetDisplayName("")
	key.SetElementId("")
	key.SetId(1)
	key.SetKeystorePassOnly(true)
	key.SetName("")
	key.SetPassword("")
	key.SetPrivateKeyName("")
	key.SetPrivateKeyPassword("")
	key.SetStore(ResourceDTO)
	key.SetType("")

	// SAML2 IdP config as serialized by CXF (Additional properties)
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
	orig.SetConfig(idpConf)

	// Attribute profile (TODO : Make configurable)
	// TODO : Use custom profile
	var atp api.AttributeProfileDTO
	atp.SetElementId("")
	atp.SetId(97)
	atp.SetName("basic-built-in")
	atp.SetProfileType("BASIC")

	atp.AdditionalProperties = make(map[string]interface{})
	atp.AdditionalProperties["@c"] = "com.atricore.idbus.console.services.dto.BuiltInAttributeProfileDTO"
	orig.SetAttributeProfile(atp)

	var AuthenticationServiceDTO api.AuthenticationServiceDTO

	AuthenticationServiceDTO.SetDescription("")
	AuthenticationServiceDTO.SetDisplayName("")
	AuthenticationServiceDTO.SetElementId("")
	AuthenticationServiceDTO.SetId(1)
	AuthenticationServiceDTO.SetName("")
	AuthenticationServiceDTO.SetX(1)
	AuthenticationServiceDTO.SetY(1)

	// Authentication contract
	var auc api.AuthenticationContractDTO
	auc.SetElementId("")
	auc.SetId(72)
	orig.SetAuthenticationContract(auc)

	orig.SetDashboardUrl("http://localhost:8080/myapp")

	orig.SetAuthenticationMechanisms(authn)
	/*
		var delegatedauthns []api.DelegatedAuthenticationDTO
		delegatedauthns1 := api.NewDelegatedAuthenticationDTO()
		delegatedauthns2 := api.NewDelegatedAuthenticationDTO()
		delegatedauthns1.SetAuthnService(AuthenticationServiceDTO)
		delegatedauthns1.SetDescription("")
		delegatedauthns1.SetElementId("")
		delegatedauthns1.SetId(1)
		delegatedauthns1.SetIdp(*orig)
		delegatedauthns1.SetName("")
		delegatedauthns = append(delegatedauthns, *delegatedauthns1)
		delegatedauthns2.SetAuthnService(AuthenticationServiceDTO)
		delegatedauthns2.SetDescription("")
		delegatedauthns2.SetElementId("")
		delegatedauthns2.SetId(1)
		delegatedauthns2.SetIdp(*orig)
		delegatedauthns2.SetName("")
		delegatedauthns = append(delegatedauthns, *delegatedauthns2)
		orig.SetDelegatedAuthentications(delegatedauthns)
	*/
	orig.SetDescription("IdP One")
	orig.SetDestroyPreviousSession(true)
	orig.SetDisplayName("")
	orig.SetElementId("")
	orig.SetEmissionPolicy(AuthenticationAssertionEmissionPolicyDTO)
	orig.SetEnableMetadataEndpoint(true)
	orig.SetEncryptAssertion(true)
	orig.SetEncryptAssertionAlgorithm("http://www.w3.org/200|/04/xmlenc#aes128-cbc")
	orig.SetErrorBinding("JSON")
	orig.SetExternallyHostedIdentityConfirmationTokenService(true)
	orig.SetFederatedConnectionsA(fedconn)
	orig.SetFederatedConnectionsB(fedconn) // preguntar
	orig.SetId(-1)
	orig.SetIdentityAppliance(identityAppliance)
	orig.SetIdentityConfirmationEnabled(true)
	orig.SetIdentityConfirmationOAuth2AuthorizationServerEndpoint("")
	orig.SetIdentityConfirmationOAuth2ClientId("")
	orig.SetIdentityConfirmationOAuth2ClientSecret("")
	orig.SetIdentityConfirmationPolicy(ExtensionDTO)
	orig.SetIgnoreRequestedNameIDPolicy(true)
	orig.SetIsRemote(true)
	orig.SetMaxSessionsPerUser(5)
	orig.SetMessageTtl(301)
	orig.SetMessageTtlTolerance(302)
	orig.SetMetadata(ResourceDTO)
	orig.SetName(name)

	// OAuth2 authentication
	var oac []api.OAuth2ClientDTO
	oac1 := api.NewOAuth2ClientDTO()
	oac1.SetBaseURL("http://host1:80/")
	oac1.SetSecret("my-secret1")
	oac = append(oac, *oac1)
	oac2 := api.NewOAuth2ClientDTO()
	oac2.SetBaseURL("http://host2:80/")
	oac2.SetSecret("my-secret2")
	oac = append(oac, *oac2)

	orig.SetOauth2Clients(oac)
	orig.SetOauth2ClientsConfig("")
	orig.SetOauth2Enabled(true)
	orig.SetOauth2Key("secret")
	orig.SetOauth2RememberMeTokenValidity(43201)
	orig.SetOauth2TokenValidity(303)
	orig.SetPwdlessAuthnEnabled(false)
	orig.SetPwdlessAuthnFrom("")
	orig.SetPwdlessAuthnSubject("")
	orig.SetPwdlessAuthnTemplate("")
	orig.SetPwdlessAuthnTo("")

	/// OpenID Connect
	orig.SetOidcAccessTokenTimeToLive(3610)
	orig.SetOidcAuthzCodeTimeToLive(305)
	orig.SetOidcIdTokenTimeToLive(3620)
	orig.SetOpenIdEnabled(true)

	orig.SetSessionManagerFactory(SessionManagerFactoryDTO)
	orig.SetSignRequests(true)
	orig.SetSignatureHash("")
	orig.SetSsoSessionTimeout(1)

	orig.SetUserDashboardBranding("josso2-branding")
	orig.SetWantAuthnRequestsSigned(true)
	orig.SetSignRequests(true)

	return orig, nil
}

func AddBasicAuthentication(idp *api.IdentityProviderDTO) {

}

func (s *AccTestSuite) TestAccCliIdP_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdP_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func IdPFieldTestCreate(
	e *api.IdentityProviderDTO,
	r *api.IdentityProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "brandign",
			cmp:      func() bool { return StrPtrEquals(e.UserDashboardBranding, r.UserDashboardBranding) },
			expected: StrDeref(e.UserDashboardBranding),
			received: StrDeref(r.UserDashboardBranding),
		},
		{
			name:     "DashboardUrl",
			cmp:      func() bool { return StrPtrEquals(e.DashboardUrl, r.DashboardUrl) },
			expected: StrDeref(e.DashboardUrl),
			received: StrDeref(r.DashboardUrl),
		},
		{
			name:     "Description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "DestroyPreviousSession",
			cmp:      func() bool { return BoolPtrEquals(e.DestroyPreviousSession, r.DestroyPreviousSession) },
			expected: strconv.FormatBool(BoolDeref(e.DestroyPreviousSession)),
			received: strconv.FormatBool(BoolDeref(r.DestroyPreviousSession)),
		},
		{
			name:     "EncryptAssertion",
			cmp:      func() bool { return BoolPtrEquals(e.EncryptAssertion, r.EncryptAssertion) },
			expected: strconv.FormatBool(BoolDeref(e.EncryptAssertion)),
			received: strconv.FormatBool(BoolDeref(r.EncryptAssertion)),
		},
		{
			name:     "EncryptAssertionAlgorithm",
			cmp:      func() bool { return StrPtrEquals(e.EncryptAssertionAlgorithm, r.EncryptAssertionAlgorithm) },
			expected: StrDeref(e.EncryptAssertionAlgorithm),
			received: StrDeref(r.EncryptAssertionAlgorithm),
		},
		{
			name:     "ErrorBinding",
			cmp:      func() bool { return StrPtrEquals(e.ErrorBinding, r.ErrorBinding) },
			expected: StrDeref(e.ErrorBinding),
			received: StrDeref(r.ErrorBinding),
		},
		{
			name:     "MaxSessionsPerUser",
			cmp:      func() bool { return Int32PtrEquals(e.MaxSessionsPerUser, r.MaxSessionsPerUser) },
			expected: strconv.Itoa(int(Int32Deref(e.MaxSessionsPerUser))),
			received: strconv.Itoa(int(Int32Deref(r.MaxSessionsPerUser))),
		},
		{
			name:     "MessageTtl",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtl, r.MessageTtl) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtl))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtl))),
		},
		{
			name:     "MessageTtlTolerance",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtlTolerance, r.MessageTtlTolerance) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtlTolerance))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtlTolerance))),
		},
		{
			name:     "Oauth2Clients",
			cmp:      func() bool { return Oauth2ClientsEquals(e.Oauth2Clients, r.Oauth2Clients) }, //TODO: NewIdentityProviderDTO
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "Oauth2ClientsConfig",
			cmp:      func() bool { return StrPtrEquals(e.Oauth2ClientsConfig, r.Oauth2ClientsConfig) },
			expected: StrDeref(e.Oauth2ClientsConfig),
			received: StrDeref(r.Oauth2ClientsConfig),
		},
		{
			name:     "Oauth2Enabled",
			cmp:      func() bool { return BoolPtrEquals(e.Oauth2Enabled, r.Oauth2Enabled) },
			expected: strconv.FormatBool(BoolDeref(e.Oauth2Enabled)),
			received: strconv.FormatBool(BoolDeref(r.Oauth2Enabled)),
		},
		{
			name:     "Oauth2Key",
			cmp:      func() bool { return StrPtrEquals(e.Oauth2Key, r.Oauth2Key) },
			expected: StrDeref(e.Oauth2Key),
			received: StrDeref(r.Oauth2Key),
		},
		{
			name:     "Oauth2RememberMeTokenValidity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2RememberMeTokenValidity, r.Oauth2RememberMeTokenValidity) },
			expected: strconv.FormatInt(Int64Deref(e.Oauth2RememberMeTokenValidity), 10),
			received: strconv.FormatInt(Int64Deref(r.Oauth2RememberMeTokenValidity), 10),
		},
		{
			name:     "Oauth2TokenValidity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2TokenValidity, r.Oauth2TokenValidity) },
			expected: strconv.FormatInt(Int64Deref(e.Oauth2TokenValidity), 10),
			received: strconv.FormatInt(Int64Deref(r.Oauth2TokenValidity), 10),
		},
		{
			name:     "OidcAccessTokenTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAccessTokenTimeToLive, r.OidcAccessTokenTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcAccessTokenTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcAccessTokenTimeToLive))),
		},
		{
			name:     "OidcAuthzCodeTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAuthzCodeTimeToLive, r.OidcAuthzCodeTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcAuthzCodeTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcAuthzCodeTimeToLive))),
		},
		{
			name:     "OidcIdTokenTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcIdTokenTimeToLive, r.OidcIdTokenTimeToLive) },
			expected: strconv.Itoa(int(Int32Deref(e.OidcIdTokenTimeToLive))),
			received: strconv.Itoa(int(Int32Deref(r.OidcIdTokenTimeToLive))),
		},
		{
			name:     "OpenIdEnabled",
			cmp:      func() bool { return BoolPtrEquals(e.OpenIdEnabled, r.OpenIdEnabled) },
			expected: strconv.FormatBool(BoolDeref(e.OpenIdEnabled)),
			received: strconv.FormatBool(BoolDeref(r.OpenIdEnabled)),
		},
		{
			name:     "PwdlessAuthnEnabled",
			cmp:      func() bool { return BoolPtrEquals(e.PwdlessAuthnEnabled, r.PwdlessAuthnEnabled) },
			expected: strconv.FormatBool(BoolDeref(e.PwdlessAuthnEnabled)),
			received: strconv.FormatBool(BoolDeref(r.PwdlessAuthnEnabled)),
		},
		{
			name:     "PwdlessAuthnSubject",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnSubject, r.PwdlessAuthnSubject) },
			expected: StrDeref(e.PwdlessAuthnSubject),
			received: StrDeref(r.PwdlessAuthnSubject),
		},
		{
			name:     "PwdlessAuthnTemplate",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTemplate, r.PwdlessAuthnTemplate) },
			expected: StrDeref(e.PwdlessAuthnTemplate),
			received: StrDeref(r.PwdlessAuthnTemplate),
		},
		{
			name:     "PwdlessAuthnTemplate",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTemplate, r.PwdlessAuthnTemplate) },
			expected: StrDeref(e.PwdlessAuthnTemplate),
			received: StrDeref(r.PwdlessAuthnTemplate),
		},
		{
			name:     "PwdlessAuthnTo",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTo, r.PwdlessAuthnTo) },
			expected: StrDeref(e.PwdlessAuthnTo),
			received: StrDeref(r.PwdlessAuthnTo),
		},
		// {
		// 	name:     "AuthenticationMechanisms",
		// 	cmp:      func() bool { return StrPtrEquals(e.AuthenticationMechanisms, r.AuthenticationMechanisms) },
		// 	expected: StrDeref(PtrString(e.AuthenticationMechanisms)),
		// 	received: StrDeref(PtrString((r.AuthenticationMechanisms))),
		// },

		// TODO : Add validator for Authn. Mechanisms property : AuthenticationMechanisms
		// TODO : Add validator for DelegatedAuthentications property
	}
}

func IdPFieldTestUpdate(
	e *api.IdentityProviderDTO,
	r *api.IdentityProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
		{
			name:     "elementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
	}

	return append(t, IdPFieldTestCreate(e, r)...)
}

// Compares the expected IdP with the received one.
func IdPValidateCreate(
	e *api.IdentityProviderDTO,
	r *api.IdentityProviderDTO) error {

	return ValidateFields(IdPFieldTestCreate(e, r))
}

// Compares the expected IdP with the received one.
func IdPValidateUpdate(
	e *api.IdentityProviderDTO,
	r *api.IdentityProviderDTO) error {

	return ValidateFields(IdPFieldTestUpdate(e, r))
}

/*
	var AuthenticationService []api.AuthenticationServiceDTO
	AuthenticationService1 := api.NewAuthenticationServiceDTO()
	AuthenticationService2 := api.NewAuthenticationServiceDTO()

	AuthenticationService1.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService1.SetDescription("")
	AuthenticationService1.SetDisplayName("")
	AuthenticationService1.SetElementId("")
	AuthenticationService1.SetId(1)
	AuthenticationService1.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService1)
	AuthenticationService2.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService2.SetDescription("")
	AuthenticationService2.SetDisplayName("")
	AuthenticationService2.SetElementId("")
	AuthenticationService2.SetId(1)
	AuthenticationService2.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService2)
*/
/*
	DelegatedAuthenticationDTO1.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO1.SetDescription("")
	DelegatedAuthenticationDTO1.SetElementId("")
	DelegatedAuthenticationDTO1.SetId(1)
	DelegatedAuthenticationDTO1.SetIdp(*orig)
	DelegatedAuthenticationDTO1.SetName("")
	DelegatedAuthenticationDTO1.SetWaypoints(poi)
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO1)
	DelegatedAuthenticationDTO2.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO2.SetDescription("")
	DelegatedAuthenticationDTO2.SetElementId("")
	DelegatedAuthenticationDTO2.SetId(1)
	DelegatedAuthenticationDTO2.SetIdp(*orig)
	DelegatedAuthenticationDTO2.SetName("")
	DelegatedAuthenticationDTO2.SetWaypoints(poi)
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO2)

	AuthenticationServiceDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationServiceDTO.SetDescription("")
	AuthenticationServiceDTO.SetDisplayName("")
	AuthenticationServiceDTO.SetElementId("")
	AuthenticationServiceDTO.SetId(1)
	AuthenticationServiceDTO.SetName("")
	AuthenticationServiceDTO.SetX(1)
	AuthenticationServiceDTO.SetY(1)

	dauh.SetAuthnService(AuthenticationServiceDTO)
	dauh.SetDescription("")
	dauh.SetElementId("")
	dauh.SetId(1)
	dauh.SetIdp(*orig)
	dauh.SetName("")
	dauh.SetWaypoints(poi)

	auc2 := api.NewAuthenticationMechanismDTO()
	auc2.SetDelegatedAuthentication(dauh)
	auc2.SetDisplayName("")
	auc2.SetElementId("")
	auc2.SetId(1)
	auc2.SetName("")
	auc2.SetPriority(1)
	auc1 = append(auc1, *auc2)
	auc3 := api.NewAuthenticationMechanismDTO()
	auc3.SetDelegatedAuthentication(dauh)
	auc3.SetDisplayName("")
	auc3.SetElementId("")
	auc3.SetId(1)
	auc3.SetName("")
	auc3.SetPriority(1)
	auc1 = append(auc1, *auc2)


*/

/*
	IdentitySourceDTO1.SetDescription("")
	IdentitySourceDTO1.SetElementId("")
	IdentitySourceDTO1.SetId(1)
	IdentitySourceDTO1.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO1)
	IdentitySourceDTO2.SetDescription("")
	IdentitySourceDTO2.SetElementId("")
	IdentitySourceDTO2.SetId(1)
	IdentitySourceDTO2.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO2)

	ExecutionEnvironmentDTO1.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO1.SetActive(true)
	ExecutionEnvironmentDTO1.SetBindingLocation(locat)
	ExecutionEnvironmentDTO1.SetDescription("")
	ExecutionEnvironmentDTO1.SetDisplayName("")
	ExecutionEnvironmentDTO1.SetElementId("")
	ExecutionEnvironmentDTO1.SetId(1)
	ExecutionEnvironmentDTO1.SetInstallDemoApps(true)
	ExecutionEnvironmentDTO1.SetInstallUri("")
	ExecutionEnvironmentDTO1.SetLocation("")
	ExecutionEnvironmentDTO1.SetName("")
	ExecutionEnvironmentDTO1.SetOverwriteOriginalSetup(true)
	ExecutionEnvironmentDTO1.SetPlatformId("")
	ExecutionEnvironmentDTO1.SetTargetJDK("")
	ExecutionEnvironmentDTO1.SetType("")
	ExecutionEnvironmentDTO = append(ExecutionEnvironmentDTO, *ExecutionEnvironmentDTO1)
	ExecutionEnvironmentDTO2.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO2.SetActive(true)
	ExecutionEnvironmentDTO2.SetBindingLocation(locat)
	ExecutionEnvironmentDTO2.SetDescription("")
	ExecutionEnvironmentDTO2.SetDisplayName("")
	ExecutionEnvironmentDTO2.SetElementId("")
	ExecutionEnvironmentDTO2.SetId(1)
	ExecutionEnvironmentDTO2.SetInstallDemoApps(true)
	ExecutionEnvironmentDTO2.SetInstallUri("")
	ExecutionEnvironmentDTO2.SetLocation("")
	ExecutionEnvironmentDTO2.SetName("")
	ExecutionEnvironmentDTO2.SetOverwriteOriginalSetup(true)
	ExecutionEnvironmentDTO2.SetPlatformId("")
	ExecutionEnvironmentDTO2.SetTargetJDK("")
	ExecutionEnvironmentDTO2.SetType("")
	ExecutionEnvironmentDTO = append(ExecutionEnvironmentDTO, *ExecutionEnvironmentDTO2)

	// identityAppliance.SetActiveFeatures("") //
	identityAppliance.SetAuthenticationServices(AuthenticationService)
	identityAppliance.SetDescription("")
	identityAppliance.SetDisplayName("")
	identityAppliance.SetElementId("")
	identityAppliance.SetExecutionEnvironments(ExecutionEnvironmentDTO)
	identityAppliance.SetId(1)
	identityAppliance.SetIdentitySources(IdentitySourceDTO)
	identityAppliance.SetIdpSelector(EntitySelectionStrategyDTO)
	identityAppliance.SetKeystore(keystore)
	// identityAppliance.SetLastModification() //////////*TIME.TIME?
	identityAppliance.SetLocation(locat)
	identityAppliance.SetModelVersion("")
	identityAppliance.SetName("")
	identityAppliance.SetNamespace("")
	identityAppliance.SetProviders(ProviderDTO)
	// identityAppliance.SetRequiredBundles() //
	identityAppliance.SetRevision(1)
	identityAppliance.SetSecurityConfig(IdentityApplianceSecurityConfigDTO)
	identityAppliance.SetServiceResources(ServiceResource)
	// identityAppliance.SetSupportedRoles() //
	identityAppliance.SetUserDashboardBranding(UserDashboardBrandingDTO)

	ResourceDTO.SetDisplayName("")
	ResourceDTO.SetElementId("")
	ResourceDTO.SetId(1)
	ResourceDTO.SetName("")
	ResourceDTO.SetUri("")
	ResourceDTO.SetValue("")


	// FederatedProviderDTO.SetActiveBindings("") //
	// FederatedProviderDTO.SetActiveProfiles("") //
	FederatedProviderDTO.SetConfig(conf)
	FederatedProviderDTO.SetDescription("")
	FederatedProviderDTO.SetDisplayName("")
	FederatedProviderDTO.SetElementId("")
	FederatedProviderDTO.SetId(1)
	FederatedProviderDTO.SetIdentityAppliance(identityAppliance)
	FederatedProviderDTO.SetIsRemote(true)
	FederatedProviderDTO.SetLocation(locat)
	FederatedProviderDTO.SetMetadata(ResourceDTO)
	FederatedProviderDTO.SetName("")
	FederatedProviderDTO.SetRemote(true)
	FederatedProviderDTO.SetRole("")

	FederatedConnection.SetChannelA(FederatedChannelDTO)
	FederatedConnection.SetChannelB(FederatedChannelDTO)
	FederatedConnection.SetDescription("")
	FederatedConnection.SetElementId("")
	FederatedConnection.SetId(1)
	FederatedConnection.SetName("")
	FederatedConnection.SetRoleA(FederatedProviderDTO)
	FederatedConnection.SetRoleB(FederatedProviderDTO)
	FederatedConnection.SetWaypoints(poi)

	// FederatedChannelDTO.SetActiveBindings("") //////////// preguntar
	// FederatedChannelDTO.SetActiveProfiles("") //////////// preguntar
	FederatedChannelDTO.SetConnectionA(FederatedConnection)
	FederatedChannelDTO.SetConnectionB(FederatedConnection)
	FederatedChannelDTO.SetDescription("")
	FederatedChannelDTO.SetDisplayName("")
	FederatedChannelDTO.SetElementId("")
	FederatedChannelDTO.SetId(1)
	FederatedChannelDTO.SetLocation(locat)
	FederatedChannelDTO.SetName("")
	FederatedChannelDTO.SetOverrideProviderSetup(true)

	fedconn1.SetChannelA(FederatedChannelDTO)
	fedconn1.SetChannelB(FederatedChannelDTO)
	fedconn1.SetDescription("")
	fedconn1.SetElementId("")
	fedconn1.SetId(1)
	fedconn1.SetName("")
	fedconn1.SetRoleA(FederatedProviderDTO)
	fedconn1.SetRoleB(FederatedProviderDTO)
	fedconn1.SetWaypoints(poi)
	fedconn = append(fedconn, *fedconn1)
	fedconn2.SetChannelA(FederatedChannelDTO)
	fedconn2.SetChannelB(FederatedChannelDTO)
	fedconn2.SetDescription("")
	fedconn2.SetElementId("")
	fedconn2.SetId(1)
	fedconn2.SetName("")
	fedconn2.SetRoleA(FederatedProviderDTO)
	fedconn2.SetRoleB(FederatedProviderDTO)
	fedconn2.SetWaypoints(poi)
	fedconn = append(fedconn, *fedconn2)
*/

/*
	ServiceResource1.SetActivation(Activation)
	ServiceResource1.SetDescription("")
	ServiceResource1.SetElementId("")
	ServiceResource1.SetId(1)
	ServiceResource1.SetName("")
	ServiceResource1.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource1)
	ServiceResource2.SetActivation(Activation)
	ServiceResource2.SetDescription("")
	ServiceResource2.SetElementId("")
	ServiceResource2.SetId(1)
	ServiceResource2.SetName("")
	ServiceResource2.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource2)

	IdentityApplianceSecurityConfigDTO.SetEncryptSensitiveData(true)
	IdentityApplianceSecurityConfigDTO.SetEncryption("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionConfig("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionConfigFile("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionPassword("")
	IdentityApplianceSecurityConfigDTO.SetExternalConfig(true)
	IdentityApplianceSecurityConfigDTO.SetExternalConfigFile("")
	IdentityApplianceSecurityConfigDTO.SetPasswordProperty("")
	IdentityApplianceSecurityConfigDTO.SetSalt("")
	IdentityApplianceSecurityConfigDTO.SetSaltProperty("")
	IdentityApplianceSecurityConfigDTO.SetSaltValue("")



	keystore.SetCertificateAlias("")
	keystore.SetDisplayName("")
	keystore.SetElementId("")
	keystore.SetId(1)
	keystore.SetKeystorePassOnly(true)
	keystore.SetName("")
	keystore.SetPassword("")
	keystore.SetPrivateKeyName("")
	keystore.SetPrivateKeyPassword("")
	keystore.SetStore(ResourceDTO)
	keystore.SetType("")
*/

/*
	ExecutionEnvironment.SetActivations(ActivationDTO)
	ExecutionEnvironment.SetActive(true)
	ExecutionEnvironment.SetBindingLocation(locat)
	ExecutionEnvironment.SetDescription("")
	ExecutionEnvironment.SetDisplayName("")
	ExecutionEnvironment.SetElementId("")
	ExecutionEnvironment.SetId(1)
	ExecutionEnvironment.SetInstallDemoApps(true)
	ExecutionEnvironment.SetInstallUri("")
	ExecutionEnvironment.SetLocation("")
	ExecutionEnvironment.SetName("")
	ExecutionEnvironment.SetOverwriteOriginalSetup(true)
	ExecutionEnvironment.SetPlatformId("")
	ExecutionEnvironment.SetTargetJDK("")
	ExecutionEnvironment.SetType("")

	ServiceResourceDTO.SetActivation(Activation)
	ServiceResourceDTO.SetDescription("")
	ServiceResourceDTO.SetElementId("")
	ServiceResourceDTO.SetId(1)
	ServiceResourceDTO.SetName("")
	ServiceResourceDTO.SetServiceConnection(ServiceConnectionDTO)

	Activation.SetDescription("")
	Activation.SetElementId("")
	Activation.SetExecutionEnv(ExecutionEnvironment)
	Activation.SetId(1)
	Activation.SetName("")
	Activation.SetResource(ServiceResourceDTO)
	//Activation.SetSp()
	Activation.SetWaypoints(poi)

	ServiceConnectionDTO.SetDescription("")
	ServiceConnectionDTO.SetElementId("")
	ServiceConnectionDTO.SetId(1)
	ServiceConnectionDTO.SetName("")
	ServiceConnectionDTO.SetResource(ServiceResourceDTO)
	//ServiceConnectionDTO.SetSp()
	ServiceConnectionDTO.SetWaypoints(poi)

	ActivationDTO1.SetDescription("")
	ActivationDTO1.SetElementId("")
	ActivationDTO1.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO1.SetId(1)
	ActivationDTO1.SetName("")
	ActivationDTO1.SetResource(ServiceResourceDTO)
	//ActivationDTO1.SetSp()
	ActivationDTO1.SetWaypoints(poi)
	ActivationDTO = append(ActivationDTO, *ActivationDTO1)
	ActivationDTO2.SetDescription("")
	ActivationDTO2.SetElementId("")
	ActivationDTO2.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO2.SetId(1)
	ActivationDTO2.SetName("")
	ActivationDTO2.SetResource(ServiceResourceDTO)
	//ActivationDTO2.SetSp()
	ActivationDTO2.SetWaypoints(poi)
	ActivationDTO = append(ActivationDTO, *ActivationDTO2)
*/
