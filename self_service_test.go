package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliSelfServiceResource_crud() {
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

	crudName := "selfService-a"
	var orig *api.SelfServicesResourceDTO
	var created api.SelfServicesResourceDTO
	orig = createTestSelfServiceDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateSelfServiceresource(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := selfServiceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating self Service : %v", err)
		return
	}
	// Test READ
	var read api.SelfServicesResourceDTO
	read, err = s.client.GetSelfServiceResource(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = selfServiceValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating self Service : %v", err)
		return

	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateSelfServiceResource(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = selfServiceValidateUpdate(&read, &updated); err != nil {
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
	listOfAll, err := s.client.GetSelfServiceResources(*appliance.Name)
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
	var listOfCreated [2]api.SelfServicesResourceDTO
	// Test list of #2 elements
	element1 := createTestSelfServiceDTO("selfService-1")
	listOfCreated[0], _ = s.client.CreateSelfServiceresource(*appliance.Name, *element1)

	element2 := createTestSelfServiceDTO("selfService-2")
	listOfCreated[1], _ = s.client.CreateSelfServiceresource(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetSelfServiceResources(*appliance.Name)
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
		if err = selfServiceValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestSelfServiceDTO(name string) *api.SelfServicesResourceDTO {
	tData := api.NewSelfServicesResourceDTO()
	var location api.LocationDTO
	location.SetProtocol("http")
	location.SetHost("www.localhost.com")
	location.SetPort(8081)
	location.SetContext("IDBUS")
	location.SetUri("TESTACC-2146919007984762839")

	var serviceConnection api.ServiceConnectionDTO
	serviceConnection.SetDescription("")
	serviceConnection.SetElementId("")
	serviceConnection.SetId(-1)
	serviceConnection.SetName("")
	//serviceConnection.SetSp()
	//serviceConnection.SetWaypoints()

	tData.SetDescription(fmt.Sprintf("Description for %s", name))
	tData.SetElementId("")
	tData.SetId(-1)
	tData.SetLocation(location)
	tData.SetName(name)
	tData.SetSecret("")
	tData.SetServiceConnection(serviceConnection)

	return tData
}

func (s *AccTestSuite) TestAccCliSelfService_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliSelfService_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func selfServiceFieldTestCreate(
	e *api.SelfServicesResourceDTO,
	r *api.SelfServicesResourceDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "secret",
			cmp:      func() bool { return StrPtrEquals(e.Secret, r.Secret) },
			expected: StrDeref(e.Secret),
			received: StrDeref(r.Secret),
		},
	}
}

//Fields to validate after IdVault update
func SelfServiceFieldTestUpdate(
	e *api.SelfServicesResourceDTO,
	r *api.SelfServicesResourceDTO) []FiledTestStruct {

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

	return append(t, selfServiceFieldTestCreate(e, r)...)
}

// Compares the expected IdVault with the received one.
func selfServiceValidateCreate(
	e *api.SelfServicesResourceDTO,
	r *api.SelfServicesResourceDTO) error {

	return ValidateFields(selfServiceFieldTestCreate(e, r))
}

// Compares the expected IdVault with the received one.
func selfServiceValidateUpdate(
	e *api.SelfServicesResourceDTO,
	r *api.SelfServicesResourceDTO) error {

	return ValidateFields(SelfServiceFieldTestUpdate(e, r))
}
