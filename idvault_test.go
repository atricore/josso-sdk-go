package cli

import (
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

	orig := api.EmbeddedIdentityVaultDTO{
		Name:                  api.PtrString("idVault-2"),
		Id:                    api.PtrInt64(-1),
		IdentityConnectorName: api.PtrString("connector-default"),
	}

	created, err := s.client.CreateIdVault(*appliance.Name, orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdVaultValidateCreate(&orig, &created); err != nil {
		t.Errorf("creating idVault : %v", err)
		return
	}

	var read api.EmbeddedIdentityVaultDTO
	read, err = s.client.GetIdVault(*appliance.Name, "idVault-2")
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdVaultValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idVault : %v", err)
		return
	}
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
			expected: e.Name,
			received: r.Name,
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
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: e.Name,
			received: r.Name,
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
