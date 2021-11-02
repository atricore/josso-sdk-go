package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliJossoResourcejosso_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	var execenv api.TomcatExecutionEnvironmentDTO
	execenv = *createTestTomcatExecutionEnvironmentDTO("execenv-a")
	execenv, err = s.client.CreateTomcatExeEnv(*appliance.Name, execenv)
	if err != nil {
		t.Error(err)
		return
	}

	crudName := "Josso1-a"
	var orig *api.JOSSO1ResourceDTO
	var created api.JOSSO1ResourceDTO
	orig = createTestJOSSO1ResourceDTO(crudName, execenv.GetName())

	// Test CREATE
	created, err = s.client.CreateJossoresource(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := JOSSO1ResourceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating JossoRs: %v", err)
		return
	}

	// Test READ
	var read api.JOSSO1ResourceDTO
	read, err = s.client.GetJosso1Resource(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = JOSSO1ResourceValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	updated, err := s.client.UpdateJosso1Resource(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = JOSSO1ResourceValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteJosso1Resource(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// ------------------------------------------------------------------------------------------------------------------
	// Test empty list

	listOfAll, err := s.client.GetJosso1Resources(*appliance.Name)
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
	var listOfCreated [2]api.JOSSO1ResourceDTO

	var execenv1 api.TomcatExecutionEnvironmentDTO
	execenv1 = *createTestTomcatExecutionEnvironmentDTO("execenv-1")
	execenv1, err = s.client.CreateTomcatExeEnv(*appliance.Name, execenv)
	if err != nil {
		t.Error(err)
		return
	}
	// Test list of #2 elements
	element1 := createTestJOSSO1ResourceDTO("Josso1-1", execenv1.GetName())
	listOfCreated[0], _ = s.client.CreateJossoresource(*appliance.Name, *element1)

	var execenv2 api.TomcatExecutionEnvironmentDTO
	execenv2 = *createTestTomcatExecutionEnvironmentDTO("execenv-2")
	execenv2, err = s.client.CreateTomcatExeEnv(*appliance.Name, execenv)
	if err != nil {
		t.Error(err)
		return
	}

	element2 := createTestJOSSO1ResourceDTO("Josso1-2", execenv2.GetName())
	listOfCreated[1], _ = s.client.CreateJossoresource(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetJosso1Resources(*appliance.Name)
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
		if err = JOSSO1ResourceValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestJOSSO1ResourceDTO(name string, execEnv string) *api.JOSSO1ResourceDTO {

	tData := api.NewJOSSO1ResourceDTO()
	var loca api.LocationDTO
	loca.SetContext("IDBUS")
	loca.SetHost("localhost")
	loca.SetPort(8081)
	loca.SetProtocol("http")
	loca.SetUri(strings.ToUpper(name))
	tData.SetSloLocation(loca)
	tData.SetPartnerAppLocation(loca)

	var sc api.ServiceConnectionDTO
	sc.SetName(execEnv)
	tData.SetServiceConnection(sc)

	var IgnoreR []string
	IgnoreR = append(IgnoreR, "")
	tData.SetIgnoredWebResources(IgnoreR)

	tData.SetDefaultResource("")
	tData.SetDescription(fmt.Sprintf("Desc %s", name))
	tData.SetElementId("")
	tData.SetId(1)
	tData.SetName(name)
	tData.SetSloLocationEnabled(true)

	return tData
}

func (s *AccTestSuite) TestAccCliJossoResourcejosso_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliJossoResourcejosso_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func JOSSO1ResourceFieldTestCreate(
	e *api.JOSSO1ResourceDTO,
	r *api.JOSSO1ResourceDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "slo_location",
			cmp:      func() bool { return LocationPtrEquals(e.SloLocation, r.SloLocation) },
			expected: LocationToStr(e.SloLocation),
			received: LocationToStr(r.SloLocation),
		},
		{
			name:     "partner_app_location(",
			cmp:      func() bool { return LocationPtrEquals(e.PartnerAppLocation, r.PartnerAppLocation) },
			expected: LocationToStr(e.PartnerAppLocation),
			received: LocationToStr(r.PartnerAppLocation),
		},
		// {
		// 	name:     "ignored_web_resources",
		// 	cmp:      func() bool { return StrPtrEquals(e.IgnoredWebResources, r.IgnoredWebResources) },
		// 	expected: StrDeref(e.Name),
		// 	received: StrDeref(r.Name),
		// },
		{
			name:     "default_resource",
			cmp:      func() bool { return StrPtrEquals(e.DefaultResource, r.DefaultResource) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "slo_location_enabled",
			cmp:      func() bool { return BoolPtrEquals(e.SloLocationEnabled, r.SloLocationEnabled) },
			expected: strconv.FormatBool(BoolDeref(e.SloLocationEnabled)),
			received: strconv.FormatBool(BoolDeref(r.SloLocationEnabled)),
		},
	}
}

//Fields to validate after Sp update
func JOSSO1ResourceFieldTestUpdate(
	e *api.JOSSO1ResourceDTO,
	r *api.JOSSO1ResourceDTO) []FiledTestStruct {

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
	}
	return append(t, JOSSO1ResourceFieldTestCreate(e, r)...)
}

// Compares the expected Sp with the received one.
func JOSSO1ResourceValidateCreate(
	e *api.JOSSO1ResourceDTO,
	r *api.JOSSO1ResourceDTO) error {

	return ValidateFields(JOSSO1ResourceFieldTestCreate(e, r))
}

// Compares the expected Sp with the received one.
func JOSSO1ResourceValidateUpdate(
	e *api.JOSSO1ResourceDTO,
	r *api.JOSSO1ResourceDTO) error {

	return ValidateFields(JOSSO1ResourceFieldTestUpdate(e, r))
}
