package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdVault_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	_, err = s.client.GetIdVaults(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}

	crudName := "idVault-A"
	var orig *api.EmbeddedIdentityVaultDTO
	var created api.EmbeddedIdentityVaultDTO
	orig = createTestEmbeddedIdentityVaultDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdVault(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdVaultValidateCreate(orig, &created); err != nil {
		t.Errorf("creating idVault : %v", err)
		return
	}
	// Test READ
	var read api.EmbeddedIdentityVaultDTO
	read, err = s.client.GetIdVault(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdVaultValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idVault : %v", err)
		return

	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.ElementId = api.PtrString("12345")

	updated, err := s.client.UpdateIdVault(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdVaultValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIdVault(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdVaults(*appliance.Name)
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
	var listOfCreated [2]api.EmbeddedIdentityVaultDTO
	// Test list of #2 elements
	element1 := createTestEmbeddedIdentityVaultDTO("IdVault-1")
	listOfCreated[0], _ = s.client.CreateIdVault(*appliance.Name, *element1)

	element2 := createTestEmbeddedIdentityVaultDTO("IdVault-2")
	listOfCreated[1], _ = s.client.CreateIdVault(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetIdVaults(*appliance.Name)
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
		if err = IdVaultValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestEmbeddedIdentityVaultDTO(name string) *api.EmbeddedIdentityVaultDTO {
	tData := api.NewEmbeddedIdentityVaultDTO()
	tData.SetName(name)
	tData.SetId(-1)
	tData.SetIdentityConnectorName("connector-default")
	return tData
}

func (s *AccTestSuite) TestAccCliIdVault_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdVault_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func IdVaultFieldTestCreate(
	e *api.EmbeddedIdentityVaultDTO,
	r *api.EmbeddedIdentityVaultDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "IdentityConnectorName",
			cmp:      func() bool { return StrPtrEquals(e.IdentityConnectorName, r.IdentityConnectorName) },
			expected: StrDeref(e.IdentityConnectorName),
			received: StrDeref(r.IdentityConnectorName),
		},
	}
}

//Fields to validate after IdVault update
func IdVaultFieldTestUpdate(
	e *api.EmbeddedIdentityVaultDTO,
	r *api.EmbeddedIdentityVaultDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}

	return append(t, IdVaultFieldTestCreate(e, r)...)
}

// Compares the expected IdVault with the received one.
func IdVaultValidateCreate(
	e *api.EmbeddedIdentityVaultDTO,
	r *api.EmbeddedIdentityVaultDTO) error {

	return ValidateFields(IdVaultFieldTestCreate(e, r))
}

// Compares the expected IdVault with the received one.
func IdVaultValidateUpdate(
	e *api.EmbeddedIdentityVaultDTO,
	r *api.EmbeddedIdentityVaultDTO) error {

	return ValidateFields(IdVaultFieldTestUpdate(e, r))
}
