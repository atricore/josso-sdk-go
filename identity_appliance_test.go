package cli

import (
	"strconv"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdentityAppliance_crud() {
	var t = s.T()

	// Test CRUD
	crudName := "ida-z"
	var orig *api.IdentityApplianceDefinitionDTO
	var created api.IdentityApplianceDefinitionDTO
	orig = createTestIdentityApplianceDefinitionDTO(crudName)
	// Create

	created, err := s.client.CreateAppliance(*orig)
	if err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}
	if err := IdApplianceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}

	// Retrieve
	retrieved, err := s.client.GetAppliance(crudName)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&created, &retrieved); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Update
	retrieved.Namespace = api.PtrString("com.atricore.ida.a.mod")
	updated, err := s.client.UpdateAppliance(retrieved)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&retrieved, &updated); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Delete
	removed, err := s.client.DeleteAppliance(strconv.FormatInt(*updated.Id, 10))
	if err != nil {
		t.Errorf("deleting identity appilance : %v", err)
		return
	}
	if !removed {
		t.Errorf("deleting identity appliance : not found %d", *updated.Id)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetAppliances()
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name, (these elements must have all the variables of the structure)
	var listOfCreated [2]api.IdentityApplianceDefinitionDTO

	element1 := createTestIdentityApplianceDefinitionDTO("ida-1")
	listOfCreated[0], _ = s.client.CreateAppliance(*element1)

	element2 := createTestIdentityApplianceDefinitionDTO("ida-2")
	listOfCreated[1], _ = s.client.CreateAppliance(*element2)

	// Get list from server
	listOfRead, err := s.client.GetAppliances()
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
}
func createTestIdentityApplianceDefinitionDTO(name string) *api.IdentityApplianceDefinitionDTO {
	l, _ := StrToLocation("http://localhost/IDBUS/IDA-B")
	orig := api.NewIdentityApplianceDefinitionDTO()
	orig.SetName(name)
	orig.SetNamespace("com.atricore.idbus.ida.b")
	orig.SetLocation(*l)
	orig.SetDescription("IDA-B TEST !")
	return orig
}

// -------------------------------------------------

//Fields to validate after appliance creation
func ApplianceFieldTestCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

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
			name:     "namespace",
			cmp:      func() bool { return StrPtrEquals(e.Namespace, r.Namespace) },
			expected: StrDeref(e.Namespace),
			received: StrDeref(r.Namespace),
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
	}
}

//Fields to validate after appliance update
func ApplianceFieldTestUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}
	return append(t, ApplianceFieldTestCreate(e, r)...)
}

// Compares the expected appliance with the received one.
func IdApplianceValidateCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestCreate(e, r))
}

// Compares the expected appliance with the received one.
func IdApplianceValidateUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestUpdate(e, r))
}
