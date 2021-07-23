package cli

import (
	"sort"
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

	var created api.IdentityProviderDTO
	orig := api.IdentityProviderDTO{
		Name:                          api.PtrString("idp-1"),
		Description:                   api.PtrString(""),
		DashboardUrl:                  api.PtrString(""),
		ElementId:                     api.PtrString("_BFD218B4-0F7A-4C7A-AAF9-41883AAE3598"),
		DestroyPreviousSession:        api.PtrBool(true),
		EncryptAssertion:              api.PtrBool(true),
		EncryptAssertionAlgorithm:     api.PtrString(""),
		ErrorBinding:                  api.PtrString(""),
		MaxSessionsPerUser:            api.PtrInt32(1),
		MessageTtl:                    api.PtrInt32(1),
		MessageTtlTolerance:           api.PtrInt32(1),
		Oauth2Clients:                 api.NewIdentityProviderDTO().Oauth2Clients,
		Oauth2ClientsConfig:           api.PtrString(""),
		Oauth2Enabled:                 api.PtrBool(true),
		Oauth2Key:                     api.PtrString(""),
		Oauth2RememberMeTokenValidity: api.PtrInt64(-1),
		Oauth2TokenValidity:           appliance.Id,
		OidcAccessTokenTimeToLive:     api.PtrInt32(-1),
		OidcAuthzCodeTimeToLive:       api.PtrInt32(-1),
		OidcIdTokenTimeToLive:         api.PtrInt32(-1),
		OpenIdEnabled:                 api.PtrBool(true),
		PwdlessAuthnEnabled:           api.PtrBool(true),
		PwdlessAuthnFrom:              api.PtrString(""),
		PwdlessAuthnSubject:           api.PtrString(""),
		PwdlessAuthnTemplate:          api.PtrString(""),
		PwdlessAuthnTo:                api.PtrString(""),
		Id:                            api.PtrInt64(-1),
		UserDashboardBranding:         api.PtrString("josso25-branding"),
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
	read, err = s.client.GetIdp(*appliance.Name, "idp-2")
	if err != nil {
		t.Error(err)
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
	toDelete := "idp-2"
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
		Name: api.PtrString("idp-1"),
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
			expected: e.DashboardUrl,
			received: r.DashboardUrl,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "Description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: e.Description,
			received: r.Description,
		},
		{
			name:     "ElementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: e.ElementId,
			received: r.ElementId,
		},
		{
			name:     "DestroyPreviousSession",
			cmp:      func() bool { return BoolPtrEquals(e.DestroyPreviousSession, r.DestroyPreviousSession) },
			expected: PtrString("e.DestroyPreviousSession"), // TODO : Convert type to bool
			received: PtrString("r.DestroyPreviousSession"), // TODO : Convert type to bool
		},
		{
			name:     "EncryptAssertion",
			cmp:      func() bool { return BoolPtrEquals(e.EncryptAssertion, r.EncryptAssertion) },
			expected: PtrString("e.EncryptAssertion"), // TODO : Convert type to bool
			received: PtrString("r.EncryptAssertion"), // TODO : Convert type to bool
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
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
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
