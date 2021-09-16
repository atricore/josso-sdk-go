package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

/*
func (s *AccTestSuite) TestAccCliIdP_a001() {
	var t = s.T()
	idp, err := s.client.GetIdp("ida-1", "idp-1")
	if err != nil {
		s.client.Logger().Errorf("cannot get idp %v", err)
		t.Error(err)
		return
	}

	fcs := idp.GetFederatedConnectionsA()
	for _, fc := range fcs {

		spChannel := fc.GetChannelA()  // This MUST be the SP channel
		idpChannel := fc.GetChannelB() // This MUST be the IDP channel

		if spChannel.Id == nil {
			//
		}

		if idpChannel.Id == nil {
			//
		}

		// TODO : Convert to specific channel

	}

	fmt.Println(len(fcs))
}
*/

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
	tData := api.NewIdentityProviderDTO()
	var ResourceDTO api.ResourceDTO
	var identityAppliance api.IdentityApplianceDefinitionDTO

	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO

	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")

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
	// TODO : Inject in IdP

	// SAML2 IdP config as serialized by CXF (Additional properties)
	var conf api.SamlR2IDPConfigDTO
	conf.SetDescription("")
	conf.SetDisplayName("")
	conf.SetUseSampleStore(true)
	conf.SetUseSystemStore(false)
	err := tData.SetSamlR2IDPConfig(&conf)
	if err != nil {
		return nil, err
	}

	// Attribute profile (TODO : Make configurable)
	// TODO : Use custom profile
	var atp api.AttributeProfileDTO
	atp.SetElementId("")
	atp.SetId(97)
	atp.SetName("basic-built-in")
	atp.SetProfileType("BASIC")

	atp.AdditionalProperties = make(map[string]interface{})
	atp.AdditionalProperties["@c"] = "com.atricore.idbus.console.services.dto.BuiltInAttributeProfileDTO"
	tData.SetAttributeProfile(atp)

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
	tData.SetAuthenticationContract(auc)

	tData.SetDashboardUrl("http://localhost:8080/myapp")

	tData.SetAuthenticationMechanisms(authn)

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

	tData.SetDescription("IdP One")
	tData.SetDestroyPreviousSession(true)
	tData.SetDisplayName("")
	tData.SetElementId("")
	tData.SetEmissionPolicy(AuthenticationAssertionEmissionPolicyDTO)
	tData.SetEnableMetadataEndpoint(true)
	tData.SetEncryptAssertion(true)
	tData.SetEncryptAssertionAlgorithm("http://www.w3.org/200|/04/xmlenc#aes128-cbc")
	tData.SetErrorBinding("JSON")
	tData.SetExternallyHostedIdentityConfirmationTokenService(true)
	tData.SetIdentityAppliance(identityAppliance)
	tData.SetIdentityConfirmationEnabled(true)
	tData.SetIdentityConfirmationOAuth2AuthorizationServerEndpoint("")
	tData.SetIdentityConfirmationOAuth2ClientId("")
	tData.SetIdentityConfirmationOAuth2ClientSecret("my-secret")

	tData.SetIgnoreRequestedNameIDPolicy(true)
	tData.SetIsRemote(true)
	tData.SetMaxSessionsPerUser(5)
	tData.SetMessageTtl(301)
	tData.SetMessageTtlTolerance(302)
	tData.SetMetadata(ResourceDTO)
	tData.SetName(name)

	// OAuth2 authentication
	var oac []api.OAuth2ClientDTO
	oac1 := api.NewOAuth2ClientDTO()
	oac1.SetBaseURL("http://host1:80/")
	oac1.SetSecret("my-secret")
	oac = append(oac, *oac1)
	oac2 := api.NewOAuth2ClientDTO()
	oac2.SetBaseURL("http://host2:80/")
	oac2.SetSecret("my-secret")
	oac = append(oac, *oac2)
	tData.SetOauth2Clients(oac)

	tData.SetOauth2ClientsConfig("")
	tData.SetOauth2Enabled(true)
	tData.SetOauth2Key("secret")
	tData.SetOauth2RememberMeTokenValidity(43201)
	tData.SetOauth2TokenValidity(303)
	tData.SetPwdlessAuthnEnabled(false)
	tData.SetPwdlessAuthnFrom("")
	tData.SetPwdlessAuthnSubject("")
	tData.SetPwdlessAuthnTemplate("")
	tData.SetPwdlessAuthnTo("")

	/// OpenID Connect
	tData.SetOidcAccessTokenTimeToLive(3610)
	tData.SetOidcAuthzCodeTimeToLive(305)
	tData.SetOidcIdTokenTimeToLive(3620)
	tData.SetOpenIdEnabled(true)

	tData.SetSessionManagerFactory(SessionManagerFactoryDTO)
	tData.SetSignRequests(true)
	tData.SetSignatureHash("")
	tData.SetSsoSessionTimeout(1)

	tData.SetUserDashboardBranding("josso2-branding")
	tData.SetWantAuthnRequestsSigned(true)
	tData.SetSignRequests(true)

	return tData, nil
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
