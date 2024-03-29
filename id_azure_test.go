package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdpAzure_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	_, err = s.client.GetIdpAzures(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}

	crudName := "idpAzure-A"
	var orig *api.AzureOpenIDConnectIdentityProviderDTO
	var created api.AzureOpenIDConnectIdentityProviderDTO
	orig = createTestAzureOpenIDConnectIdentityDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdpAzure(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdAzureValidateCreate(orig, &created); err != nil {
		t.Errorf("creating idpAzure : %v", err)
		return
	}
	// Test READ
	var read api.AzureOpenIDConnectIdentityProviderDTO
	read, err = s.client.GetIdpAzure(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdAzureValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idpAzure : %v", err)
		return

	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateIdpAzure(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdAzureValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIdpAzure(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdpAzures(*appliance.Name)
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
	var listOfCreated [2]api.AzureOpenIDConnectIdentityProviderDTO
	// Test list of #2 elements
	element1 := createTestAzureOpenIDConnectIdentityDTO("idpAzure-1")
	listOfCreated[0], _ = s.client.CreateIdpAzure(*appliance.Name, *element1)

	element2 := createTestAzureOpenIDConnectIdentityDTO("idpAzure-2")
	listOfCreated[1], _ = s.client.CreateIdpAzure(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetIdpAzures(*appliance.Name)
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
		if err = IdAzureValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestAzureOpenIDConnectIdentityDTO(name string) *api.AzureOpenIDConnectIdentityProviderDTO {
	tData := api.NewAzureOpenIDConnectIdentityProviderDTO()
	var location api.LocationDTO
	location.SetProtocol("http")
	location.SetHost("www.localhost.com")
	location.SetPort(8081)
	location.SetContext("IDBUS")
	location.SetUri("TESTACC-2146919007984762839/IDP-2")
	tData.SetLocation(location)

	var locationToAuthz api.LocationDTO
	locationToAuthz.SetProtocol("https")
	locationToAuthz.SetHost("login.microsoft.com")
	locationToAuthz.SetPort(443)
	locationToAuthz.SetContext("<change-me>")
	locationToAuthz.SetUri("/oauth2/v2.0/token")

	var locationToTokenServ api.LocationDTO
	locationToTokenServ.SetProtocol("https")
	locationToTokenServ.SetHost("login.microsoft.com")
	locationToTokenServ.SetPort(443)
	locationToTokenServ.SetContext("<change-me>")
	locationToTokenServ.SetUri("/oauth2/v2.0/authorize")

	tData.SetName(name)
	tData.SetId(-1)
	tData.SetElementId("")
	tData.SetDescription(fmt.Sprintf("Description for %s", name))
	tData.SetClientId("")
	tData.SetClientSecret("")
	tData.SetServerKey("")
	tData.SetAuthzTokenService(locationToAuthz)
	tData.SetAccessTokenService(locationToTokenServ)
	return tData
}

func (s *AccTestSuite) TestAccCliIdAzure_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdAzure_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

// Fields to validate after appliance creation
func IdAzureFieldTestCreate(
	e *api.AzureOpenIDConnectIdentityProviderDTO,
	r *api.AzureOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

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
	}
}

// Fields to validate after IdVault update
func IdAzureFieldTestUpdate(
	e *api.AzureOpenIDConnectIdentityProviderDTO,
	r *api.AzureOpenIDConnectIdentityProviderDTO) []FiledTestStruct {

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

	return append(t, IdAzureFieldTestCreate(e, r)...)
}

// Compares the expected IdVault with the received one.
func IdAzureValidateCreate(
	e *api.AzureOpenIDConnectIdentityProviderDTO,
	r *api.AzureOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdAzureFieldTestCreate(e, r))
}

// Compares the expected IdVault with the received one.
func IdAzureValidateUpdate(
	e *api.AzureOpenIDConnectIdentityProviderDTO,
	r *api.AzureOpenIDConnectIdentityProviderDTO) error {

	return ValidateFields(IdAzureFieldTestUpdate(e, r))
}
