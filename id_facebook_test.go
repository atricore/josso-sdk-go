package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdFacebook_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	_, err = s.client.GetIdpFacebooks(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}

	crudName := "idFacebook-A"
	var orig *api.FacebookOpenIDConnectIdentityProviderDTO
	var created api.FacebookOpenIDConnectIdentityProviderDTO
	orig = createTestFacebookOpenIDConnectIdentityDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdFacebook(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdFacebookValidateCreate(orig, &created); err != nil {
		t.Errorf("creating IdFacebook : %v", err)
		return
	}
	// Test READ
	var read api.FacebookOpenIDConnectIdentityProviderDTO
	read, err = s.client.GetIdpFacebook(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdFacebookValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating IdFacebook : %v", err)
		return

	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateIdpFacebook(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdFacebookValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIdpFacebook(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdpFacebooks(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name
	var listOfCreated [2]api.FacebookOpenIDConnectIdentityProviderDTO
	// Test list of #2 elements
	element1 := createTestFacebookOpenIDConnectIdentityDTO("idFacebook-1")
	listOfCreated[0], _ = s.client.CreateIdFacebook(*appliance.Name, *element1)

	element2 := createTestFacebookOpenIDConnectIdentityDTO("idFacebook-2")
	listOfCreated[1], _ = s.client.CreateIdFacebook(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetIdpFacebooks(*appliance.Name)
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
		if err = IdFacebookValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestFacebookOpenIDConnectIdentityDTO(name string) *api.FacebookOpenIDConnectIdentityProviderDTO {
	tData := api.NewFacebookOpenIDConnectIdentityProviderDTO()

	var locationToAuthz api.LocationDTO
	locationToAuthz.SetProtocol("https")
	locationToAuthz.SetHost("www.facebook.com")
	locationToAuthz.SetPort(443)
	locationToAuthz.SetContext("dialog")
	locationToAuthz.SetUri("oauth")
	//https://www.facebook.com:443/dialog/oauth

	var locationToTokenServ api.LocationDTO
	locationToTokenServ.SetProtocol("https")
	locationToTokenServ.SetHost("www.facebook.com")
	locationToTokenServ.SetPort(443)
	locationToTokenServ.SetContext("oauth")
	locationToTokenServ.SetUri("access_token")
	//https://graph.facebook.com:443/oauth/access_token

	tData.SetName(name)
	tData.SetId(-1)
	tData.SetElementId("")
	tData.SetDescription(fmt.Sprintf("Description for %s", name))
	tData.SetClientId("")
	tData.SetClientSecret("")
	tData.SetServerKey("")
	tData.SetAuthzTokenService(locationToAuthz)
	tData.SetAccessTokenService(locationToTokenServ)
	tData.SetScopes("email")
	tData.SetUserFields("gender")
	return tData
}

func (s *AccTestSuite) TestAccCliIdFacebook_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdFacebook_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func IdFacebookFieldTestCreate(
	e *api.FacebookOpenIDConnectIdentityProviderDTO,
	r *api.FacebookOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "client_id",
			cmp:      func() bool { return StrPtrEquals(e.ClientId, r.ClientId) },
			expected: StrDeref(e.ClientId),
			received: StrDeref(r.ClientId),
		},
		{
			name:     "client_secret",
			cmp:      func() bool { return StrPtrEquals(e.ClientSecret, r.ClientSecret) },
			expected: StrDeref(e.ClientSecret),
			received: StrDeref(r.ClientSecret),
		},
		{
			name:     "server_key",
			cmp:      func() bool { return StrPtrEquals(e.ServerKey, r.ServerKey) },
			expected: StrDeref(e.ServerKey),
			received: StrDeref(r.ServerKey),
		},
		{
			name:     "authz_token_service",
			cmp:      func() bool { return LocationPtrEquals(e.AuthzTokenService, r.AuthzTokenService) },
			expected: LocationToStr(e.AuthzTokenService),
			received: LocationToStr(r.AuthzTokenService),
		},
		{
			name:     "access_token_service",
			cmp:      func() bool { return LocationPtrEquals(e.AccessTokenService, r.AccessTokenService) },
			expected: LocationToStr(e.AccessTokenService),
			received: LocationToStr(r.AccessTokenService),
		},
		{
			name:     "scopes",
			cmp:      func() bool { return StrPtrEquals(e.Scopes, r.Scopes) },
			expected: StrDeref(e.Scopes),
			received: StrDeref(r.Scopes),
		},
		{
			name:     "user_fields",
			cmp:      func() bool { return StrPtrEquals(e.UserFields, r.UserFields) },
			expected: StrDeref(e.UserFields),
			received: StrDeref(r.UserFields),
		},
	}
}

//Fields to validate after IdVault update
func IdFacebookFieldTestUpdate(
	e *api.FacebookOpenIDConnectIdentityProviderDTO,
	r *api.FacebookOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

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
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
	}

	return append(t, IdFacebookFieldTestCreate(e, r)...)
}

// Compares the expected IdVault with the received one.
func IdFacebookValidateCreate(
	e *api.FacebookOpenIDConnectIdentityProviderDTO,
	r *api.FacebookOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdFacebookFieldTestCreate(e, r))
}

// Compares the expected IdVault with the received one.
func IdFacebookValidateUpdate(
	e *api.FacebookOpenIDConnectIdentityProviderDTO,
	r *api.FacebookOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdFacebookFieldTestUpdate(e, r))
}
