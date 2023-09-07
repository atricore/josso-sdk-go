package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdpGoogle_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	_, err = s.client.GetIdpGoogles(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}

	crudName := "idpAzure-A"
	var orig *api.GoogleOpenIDConnectIdentityProviderDTO
	var created api.GoogleOpenIDConnectIdentityProviderDTO
	orig = createTestGoogleOpenIDConnectIdentityDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdpGoogle(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdGoogleValidateCreate(orig, &created); err != nil {
		t.Errorf("creating idpAzure : %v", err)
		return
	}
	// Test READ
	var read api.GoogleOpenIDConnectIdentityProviderDTO
	read, err = s.client.GetIdpGoogle(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdGoogleValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idpAzure : %v", err)
		return

	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateIdpGoogle(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdGoogleValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIdpGoogle(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdpGoogles(*appliance.Name)
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
	var listOfCreated [2]api.GoogleOpenIDConnectIdentityProviderDTO
	// Test list of #2 elements
	element1 := createTestGoogleOpenIDConnectIdentityDTO("idpAzure-1")
	listOfCreated[0], _ = s.client.CreateIdpGoogle(*appliance.Name, *element1)

	element2 := createTestGoogleOpenIDConnectIdentityDTO("idpAzure-2")
	listOfCreated[1], _ = s.client.CreateIdpGoogle(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetIdpGoogles(*appliance.Name)
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
		if err = IdGoogleValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestGoogleOpenIDConnectIdentityDTO(name string) *api.GoogleOpenIDConnectIdentityProviderDTO {
	tData := api.NewGoogleOpenIDConnectIdentityProviderDTO()

	var location api.LocationDTO
	location.SetProtocol("http")
	location.SetHost("www.localhost.com")
	location.SetPort(8081)
	location.SetContext("IDBUS")
	location.SetUri("TESTACC-2146919007984762839/IDP-1")
	tData.SetLocation(location)

	var locationToAuthz api.LocationDTO
	locationToAuthz.SetProtocol("https")
	locationToAuthz.SetHost("accounts.google.com")
	locationToAuthz.SetPort(443)
	locationToAuthz.SetContext("o")
	locationToAuthz.SetUri("oauth2/auth")
	tData.SetAuthzTokenService(locationToAuthz)

	var locationToTokenServ api.LocationDTO
	locationToTokenServ.SetProtocol("https")
	locationToTokenServ.SetHost("accounts.google.com")
	locationToTokenServ.SetPort(443)
	locationToTokenServ.SetContext("o")
	locationToTokenServ.SetUri("oauth2/token")
	tData.SetAccessTokenService(locationToTokenServ)

	tData.SetName(name)
	tData.SetId(-1)
	tData.SetElementId("")
	tData.SetDescription(fmt.Sprintf("Description for %s", name))
	tData.SetClientId("")
	tData.SetClientSecret("")
	tData.SetServerKey("")

	tData.SetGoogleAppsDomain("Google Suite")
	return tData
}

func (s *AccTestSuite) TestAccCliIdGoogle_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdGoogle_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

// Fields to validate after appliance creation
func IdGoogleFieldTestCreate(
	e *api.GoogleOpenIDConnectIdentityProviderDTO,
	r *api.GoogleOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

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
			name:     "google_apps_domain",
			cmp:      func() bool { return StrPtrEquals(e.GoogleAppsDomain, r.GoogleAppsDomain) },
			expected: StrDeref(e.GoogleAppsDomain),
			received: StrDeref(r.GoogleAppsDomain),
		},
	}
}

// Fields to validate after IdVault update
func IdGoogleFieldTestUpdate(
	e *api.GoogleOpenIDConnectIdentityProviderDTO,
	r *api.GoogleOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

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

	return append(t, IdGoogleFieldTestCreate(e, r)...)
}

// Compares the expected IdVault with the received one.
func IdGoogleValidateCreate(
	e *api.GoogleOpenIDConnectIdentityProviderDTO,
	r *api.GoogleOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdGoogleFieldTestCreate(e, r))
}

// Compares the expected IdVault with the received one.
func IdGoogleValidateUpdate(
	e *api.GoogleOpenIDConnectIdentityProviderDTO,
	r *api.GoogleOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdGoogleFieldTestUpdate(e, r))
}
