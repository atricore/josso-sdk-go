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
	orig = createTestIdentityProviderDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("TTL %d", created.GetMessageTtl())
	t.Logf("TTL %d", orig.GetMessageTtl())
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
	if err = IdPValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}
	if read.Name == nil {
		t.Errorf("IdP not found for name %s", crudName)
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
	element1 := createTestIdentityProviderDTO("ids-1")
	listOfCreated[0], _ = s.client.CreateIdp(*appliance.Name, *element1)

	element2 := createTestIdentityProviderDTO("ids-2")
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
func createTestIdentityProviderDTO(name string) *api.IdentityProviderDTO {
	var oac []api.OAuth2ClientDTO
	oac1 := api.NewOAuth2ClientDTO()
	oac1.SetBaseURL("http://host1:80/")
	oac1.SetSecret("my-secret1")
	oac = append(oac, *oac1)
	oac2 := api.NewOAuth2ClientDTO()
	oac2.SetBaseURL("http://host2:80/")
	oac2.SetSecret("my-secret2")
	oac = append(oac, *oac1)
	orig := api.NewIdentityProviderDTO()
	orig.SetName(name)
	orig.SetDescription("IdP One")
	orig.SetDashboardUrl("http://localhost:8080/myapp")
	orig.SetDestroyPreviousSession(true)
	orig.SetEncryptAssertion(true)
	orig.SetEncryptAssertionAlgorithm("http://www.w3.org/2001/04/xmlenc#aes128-cbc")
	orig.SetErrorBinding("JSON")
	orig.SetMaxSessionsPerUser(5)
	orig.SetMessageTtl(301)
	orig.SetMessageTtlTolerance(302)
	orig.SetOauth2Clients(oac)
	orig.SetOauth2ClientsConfig("")
	orig.SetOauth2Enabled(false)
	orig.SetOauth2Key("secret")
	orig.SetOauth2RememberMeTokenValidity(43201)
	orig.SetOauth2TokenValidity(303)
	orig.SetOidcAccessTokenTimeToLive(3610)
	orig.SetOidcAuthzCodeTimeToLive(305)
	orig.SetOidcIdTokenTimeToLive(3620)
	orig.SetOpenIdEnabled(false)
	orig.SetPwdlessAuthnEnabled(false)
	orig.SetPwdlessAuthnFrom("")
	orig.SetPwdlessAuthnSubject("")
	orig.SetPwdlessAuthnTemplate("")
	orig.SetPwdlessAuthnTo("")
	orig.SetId(-1)
	orig.SetUserDashboardBranding("josso2-branding")
	return orig
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
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
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
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: StrDeref(e.DisplayName),
			received: StrDeref(r.DisplayName),
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
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "MessageTtl",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtl, r.MessageTtl) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "MessageTtlTolerance",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtlTolerance, r.MessageTtlTolerance) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		// {
		// 	name:     "Oauth2Clients",
		// 	cmp:      func() bool { return StrPtrEquals(e.Oauth2Clients, r.Oauth2Clients) }, //TODO: NewIdentityProviderDTO
		// 	expected: StrDeref(e.Name),
		// 	received: StrDeref(r.Name),
		// },
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
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "OidcAuthzCodeTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAuthzCodeTimeToLive, r.OidcAuthzCodeTimeToLive) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "OidcIdTokenTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcIdTokenTimeToLive, r.OidcIdTokenTimeToLive) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
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
			name:     "ElementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
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
