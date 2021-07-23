package cli

import (
	"sort"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdP_crud() {
	var t = s.T()

	idpName := "idp-1"

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	var created api.IdentityProviderDTO
	orig := api.IdentityProviderDTO{
		Name:                      api.PtrString(idpName),
		Description:               api.PtrString("IdP one"),
		DashboardUrl:              api.PtrString("http://localhost:8080/myapp"),
		ElementId:                 api.PtrString("_BFD218B4-0F7A-4C7A-AAF9-41883AAE3598"),
		DestroyPreviousSession:    api.PtrBool(true),
		EncryptAssertion:          api.PtrBool(false),
		EncryptAssertionAlgorithm: api.PtrString("http://www.w3.org/2001/04/xmlenc#aes128-cbc"),
		ErrorBinding:              api.PtrString("ARTIFACT"),
		MaxSessionsPerUser:        api.PtrInt32(-1),
		MessageTtl:                api.PtrInt32(300),
		MessageTtlTolerance:       api.PtrInt32(300),
		Oauth2Clients:             api.NewIdentityProviderDTO().Oauth2Clients,
		//		Oauth2ClientsConfig:           api.PtrString("null"),
		Oauth2Enabled:                 api.PtrBool(false),
		Oauth2Key:                     api.PtrString("secret"),
		Oauth2RememberMeTokenValidity: api.PtrInt64(43200),
		Oauth2TokenValidity:           api.PtrInt64(300),
		OidcAccessTokenTimeToLive:     api.PtrInt32(3600),
		OidcAuthzCodeTimeToLive:       api.PtrInt32(300),
		OidcIdTokenTimeToLive:         api.PtrInt32(3600),
		OpenIdEnabled:                 api.PtrBool(false),
		PwdlessAuthnEnabled:           api.PtrBool(false),
		//		PwdlessAuthnFrom:              api.PtrString("null"),
		//		PwdlessAuthnSubject:           api.PtrString("null"),
		//		PwdlessAuthnTemplate:          api.PtrString("null"),
		//		PwdlessAuthnTo:                api.PtrString("null"),
		Id:                    api.PtrInt64(-1),
		UserDashboardBranding: api.PtrString("josso25-branding"),
	}
	// Test CREATE
	created, err = s.client.CreateIdp(*appliance.Name, orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdPValidateCreate(&orig, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}

	// Test READ
	var read api.IdentityProviderDTO
	read, err = s.client.GetIdp(*appliance.Name, idpName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdPValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}
	if read.Name == nil {
		t.Errorf("IdP not found for name %s", idpName)
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
	toDelete := idpName
	deleted, err := s.client.DeleteIdp(*appliance.Name, toDelete)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", toDelete)
		return
	}

	// ------------------------------------------------------------------------------------------------------------------
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
	element1 := api.IdentityProviderDTO{
		Name: api.PtrString(idpName),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[0], _ = s.client.CreateIdp(*appliance.Name, element1)

	element2 := api.IdentityProviderDTO{
		Name: api.PtrString("idp-2"),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[1], _ = s.client.CreateIdp(*appliance.Name, element2)

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
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) > 0
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
			expected: e.Name,
			received: r.Name,
		},
	}
}

// TODO
//Fields to validate after IdP update
func IdPFieldTestUpdate(
	e *api.IdentityProviderDTO,
	r *api.IdentityProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "brandign",
			cmp:      func() bool { return StrPtrEquals(e.UserDashboardBranding, r.UserDashboardBranding) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "DashboardUrl",
			cmp:      func() bool { return StrPtrEquals(e.DashboardUrl, r.DashboardUrl) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "Description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "ElementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "DestroyPreviousSession",
			cmp:      func() bool { return BoolPtrEquals(e.DestroyPreviousSession, r.DestroyPreviousSession) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "EncryptAssertion",
			cmp:      func() bool { return BoolPtrEquals(e.EncryptAssertion, r.EncryptAssertion) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "EncryptAssertionAlgorithm",
			cmp:      func() bool { return StrPtrEquals(e.EncryptAssertionAlgorithm, r.EncryptAssertionAlgorithm) },
			expected: e.EncryptAssertionAlgorithm,
			received: r.EncryptAssertionAlgorithm,
		},
		{
			name:     "ErrorBinding",
			cmp:      func() bool { return StrPtrEquals(e.ErrorBinding, r.ErrorBinding) },
			expected: e.ErrorBinding,
			received: r.ErrorBinding,
		},
		{
			name:     "MaxSessionsPerUser",
			cmp:      func() bool { return Int32PtrEquals(e.MaxSessionsPerUser, r.MaxSessionsPerUser) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "MessageTtl",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtl, r.MessageTtl) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "MessageTtlTolerance",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtlTolerance, r.MessageTtlTolerance) },
			expected: e.Name,
			received: r.Name,
		},
		// {
		// 	name:     "Oauth2Clients",
		// 	cmp:      func() bool { return StrPtrEquals(e.Oauth2Clients, r.Oauth2Clients) }, //TODO: NewIdentityProviderDTO
		// 	expected: e.Name,
		// 	received: r.Name,
		// },
		{
			name:     "Oauth2ClientsConfig",
			cmp:      func() bool { return StrPtrEquals(e.Oauth2ClientsConfig, r.Oauth2ClientsConfig) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "Oauth2Enabled",
			cmp:      func() bool { return BoolPtrEquals(e.Oauth2Enabled, r.Oauth2Enabled) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "Oauth2Key",
			cmp:      func() bool { return StrPtrEquals(e.Oauth2Key, r.Oauth2Key) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "Oauth2RememberMeTokenValidity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2RememberMeTokenValidity, r.Oauth2RememberMeTokenValidity) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "Oauth2TokenValidity",
			cmp:      func() bool { return Int64PtrEquals(e.Oauth2TokenValidity, r.Oauth2TokenValidity) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "OidcAccessTokenTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAccessTokenTimeToLive, r.OidcAccessTokenTimeToLive) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "OidcAuthzCodeTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcAuthzCodeTimeToLive, r.OidcAuthzCodeTimeToLive) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "OidcIdTokenTimeToLive",
			cmp:      func() bool { return Int32PtrEquals(e.OidcIdTokenTimeToLive, r.OidcIdTokenTimeToLive) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "OpenIdEnabled",
			cmp:      func() bool { return BoolPtrEquals(e.OpenIdEnabled, r.OpenIdEnabled) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "PwdlessAuthnEnabled",
			cmp:      func() bool { return BoolPtrEquals(e.PwdlessAuthnEnabled, r.PwdlessAuthnEnabled) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "PwdlessAuthnSubject",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnSubject, r.PwdlessAuthnSubject) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "PwdlessAuthnTemplate",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTemplate, r.PwdlessAuthnTemplate) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "PwdlessAuthnTemplate",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTemplate, r.PwdlessAuthnTemplate) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "PwdlessAuthnTo",
			cmp:      func() bool { return StrPtrEquals(e.PwdlessAuthnTo, r.PwdlessAuthnTo) },
			expected: e.Name,
			received: r.Name,
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
