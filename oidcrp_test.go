package cli

import (
	"sort"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliOidcRp_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	var created api.ExternalOpenIDConnectRelayingPartyDTO
	orig := api.ExternalOpenIDConnectRelayingPartyDTO{
		Name: api.PtrString("rp-2"),
		Id:   api.PtrInt64(-1),
	}

	// Test CREATE
	created, err = s.client.CreateOidcRp(*appliance.Name, orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := OidcRpValidateCreate(&orig, &created); err != nil {
		t.Errorf("creating rp : %v", err)
		return
	}

	// Test READ
	var read api.ExternalOpenIDConnectRelayingPartyDTO
	read, err = s.client.GetOidcRp(*appliance.Name, "rp-2")
	if err != nil {
		t.Error(err)
		return
	}
	if err = OidcRpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating rp : %v", err)
		return
	}

	// Test UPDATE
	// 1. Modify an existing OIDCRP store in read, set description to a new value
	read.Description = api.PtrString("My updated description")
	read.ClientId = api.PtrString("1234")
	read.ClientType = api.PtrString("type1")

	// 2. Send update request to server
	updated, err := s.client.UpdateOidcRp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}

	// 3. Validate updated vs original name, descriptions must be OK
	if err = OidcRpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	// Test Delete
	toDelete := "rp-2" // Name of the RP to be deleted
	// 1. Send delete request to server using ODIC RP name
	deleted, err := s.client.DeleteOidcRp(*appliance.Name, toDelete)
	if err != nil {
		t.Error(err)
		return
	}
	// 2. Validate that the RP was deleted
	if !deleted {
		t.Errorf("Not deleted! %s", toDelete)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetOidcRps(*appliance.Name)
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// ------------------------
	// List of created elements, order by Name
	var listOfCreated [2]api.ExternalOpenIDConnectRelayingPartyDTO
	// Test list of #2 elements
	element1 := api.ExternalOpenIDConnectRelayingPartyDTO{
		Name: api.PtrString("rp-1"),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[0], _ = s.client.CreateOidcRp(*appliance.Name, element1)

	element2 := api.ExternalOpenIDConnectRelayingPartyDTO{
		Name: api.PtrString("rp-2"),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[1], _ = s.client.CreateOidcRp(*appliance.Name, element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetOidcRps(*appliance.Name)

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
		if err = OidcRpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func (s *AccTestSuite) TestAccCliOidcRp_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliOidcRp_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------update------------------

//Fields to validate after appliance creation
func OidcRpFieldTestCreate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: e.Name,
			received: r.Name,
		},
	}
}

//Fields to validate after OidcRp update
func OidcRpFieldTestUpdate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) []FiledTestStruct {

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
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: e.Description,
			received: r.Description,
		},
		{
			name:     "clientId",
			cmp:      func() bool { return StrPtrEquals(e.ClientId, r.ClientId) },
			expected: e.ClientId,
			received: r.ClientId,
		},
		{
			name:     "clientType",
			cmp:      func() bool { return StrPtrEquals(e.ClientType, r.ClientType) },
			expected: e.ClientType,
			received: r.ClientType,
		},
	}

	return append(t, OidcRpFieldTestCreate(e, r)...)
}

// Compares the expected OidcRp with the received one.
func OidcRpValidateCreate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) error {

	return ValidateFields(OidcRpFieldTestCreate(e, r))
}

// Compares the expected OidcRp with the received one.
func OidcRpValidateUpdate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) error {

	return ValidateFields(OidcRpFieldTestUpdate(e, r))
}
