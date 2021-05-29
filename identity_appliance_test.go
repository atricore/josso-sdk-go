package cli

import (
	"strconv"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdentityAppliance_basic() {
	var t = s.T()
	// Retrieve
	retrieved, err := s.client.GetAppliance("ida-a")
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	t.Logf("Appliance: %#v", retrieved)
}

func (s *AccTestSuite) TestAccCliIdentityAppliance_crud() {
	var t = s.T()

	// Test CRUD

	// Create
	l, _ := StrToLocation("http://localhost/IDBUS/IDA-B")
	orig := api.IdentityApplianceDefinitionDTO{
		Name:        api.PtrString("ida-b"),
		Namespace:   api.PtrString("com.atricore.idbus.ida.b"),
		Location:    l,
		Description: api.PtrString("IDA-B TEST !"),
	}
	created, err := s.client.CreateAppliance(orig)
	if err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}
	if err := IdApplianceValidateCreate(&orig, &created); err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}

	// Retrieve
	retrieved, err := s.client.GetAppliance("ida-b")
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
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "namespace",
			cmp:      func() bool { return StrPtrEquals(e.Namespace, r.Namespace) },
			expected: e.Namespace,
			received: r.Namespace,
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: e.Name,
			received: r.Name,
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
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "elementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: e.Name,
			received: r.Name,
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
