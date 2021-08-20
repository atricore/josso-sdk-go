package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliDbIdentityvaultDto() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "DdIdentityVauld-a"
	var orig *api.DbIdentityVaultDTO
	var created api.DbIdentityVaultDTO
	orig = createTestDbIdentityVaultDto(crudName)

	// Test CREATE
	created, err = s.client.CreateDbIdentitySource(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := DbIdentityVaultValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.DbIdentityVaultDTO
	read, err = s.client.GetDbIdentityVaultDto(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = DbIdentityVaultValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateDbIdentityVaultDto(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = DbIdentityVaultValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteDbIdentityVaultDto(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetDbIdentityVaultDtos(*appliance.Name)
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
	var listOfCreated [2]api.DbIdentityVaultDTO

	// Test list of #2 elements
	element1 := createTestDbIdentityVaultDto("DbIdentityVault-2")
	listOfCreated[0], _ = s.client.CreateDbIdentitySource(*appliance.Name, *element1)

	element2 := createTestDbIdentityVaultDto("DbIdentityVault-2")
	listOfCreated[1], _ = s.client.CreateDbIdentitySource(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetDbIdentityVaultDtos(*appliance.Name)
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
		if err = DbIdentityVaultValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestDbIdentityVaultDto(name string) *api.DbIdentityVaultDTO {
	orig := api.NewDbIdentityVaultDTO()
	orig.SetAcquireIncrement(1)
	orig.SetConnectionUrl("")
	orig.SetDescription("")
	orig.SetDriverName("")
	orig.SetElementId("")
	orig.SetExternalDB(true)
	orig.SetHashAlgorithm("")
	orig.SetHashEncoding("")
	orig.SetId(1)
	orig.SetIdleConnectionTestPeriod(1)
	orig.SetInitialPoolSize(1)
	orig.SetMaxIdleTime(1)
	orig.SetMaxPoolSize(1)
	orig.SetMinPoolSize(1)
	orig.SetName("")
	orig.SetPassword("")
	orig.SetPooledDatasource(true)
	orig.SetSaltLength(1)
	orig.SetSaltValue("")
	orig.SetUsername("")
	return orig
}

func (s *AccTestSuite) TestAccCliDbIdentitySourceDto_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func DbIdentityVaultFieldTestCreate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "acquireincrement",
			cmp:      func() bool { return Int32PtrEquals(e.AcquireIncrement, r.AcquireIncrement) },
			expected: strconv.Itoa(int(Int32Deref(e.AcquireIncrement))),
			received: strconv.Itoa(int(Int32Deref(r.AcquireIncrement))),
		},
		{
			name:     "connectionurl",
			cmp:      func() bool { return StrPtrEquals(e.ConnectionUrl, r.ConnectionUrl) },
			expected: StrDeref(e.ConnectionUrl),
			received: StrDeref(r.ConnectionUrl),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "DriverName",
			cmp:      func() bool { return StrPtrEquals(e.DriverName, r.DriverName) },
			expected: StrDeref(e.DriverName),
			received: StrDeref(r.DriverName),
		},
		{
			name:     "externaldb",
			cmp:      func() bool { return BoolPtrEquals(e.ExternalDB, r.ExternalDB) },
			expected: strconv.FormatBool(BoolDeref(e.ExternalDB)),
			received: strconv.FormatBool(BoolDeref(r.ExternalDB)),
		},
		{
			name:     "hashalgorith",
			cmp:      func() bool { return StrPtrEquals(e.HashAlgorithm, r.HashAlgorithm) },
			expected: StrDeref(e.HashAlgorithm),
			received: StrDeref(r.HashAlgorithm),
		},
		{
			name:     "hashencoding",
			cmp:      func() bool { return StrPtrEquals(e.HashEncoding, r.HashEncoding) },
			expected: StrDeref(e.HashEncoding),
			received: StrDeref(r.HashEncoding),
		},
		{
			name:     "idleconnectiontestperiod",
			cmp:      func() bool { return Int32PtrEquals(e.IdleConnectionTestPeriod, r.IdleConnectionTestPeriod) },
			expected: strconv.Itoa(int(Int32Deref(e.IdleConnectionTestPeriod))),
			received: strconv.Itoa(int(Int32Deref(r.IdleConnectionTestPeriod))),
		},
		{
			name:     "initialpoolsize",
			cmp:      func() bool { return Int32PtrEquals(e.InitialPoolSize, r.InitialPoolSize) },
			expected: strconv.Itoa(int(Int32Deref(e.InitialPoolSize))),
			received: strconv.Itoa(int(Int32Deref(r.InitialPoolSize))),
		},
		{
			name:     "maxidletime",
			cmp:      func() bool { return Int32PtrEquals(e.MaxIdleTime, r.MaxIdleTime) },
			expected: strconv.Itoa(int(Int32Deref(e.MaxIdleTime))),
			received: strconv.Itoa(int(Int32Deref(r.MaxIdleTime))),
		},
		{
			name:     "maxpoolsize",
			cmp:      func() bool { return Int32PtrEquals(e.MaxPoolSize, r.MaxPoolSize) },
			expected: strconv.Itoa(int(Int32Deref(e.MaxPoolSize))),
			received: strconv.Itoa(int(Int32Deref(r.MaxPoolSize))),
		},
		{
			name:     "minpoolsize",
			cmp:      func() bool { return Int32PtrEquals(e.MinPoolSize, r.MinPoolSize) },
			expected: strconv.Itoa(int(Int32Deref(e.MinPoolSize))),
			received: strconv.Itoa(int(Int32Deref(r.MinPoolSize))),
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "password",
			cmp:      func() bool { return StrPtrEquals(e.Password, r.Password) },
			expected: StrDeref(e.Password),
			received: StrDeref(r.Password),
		},
		{
			name:     "pooleddatasource",
			cmp:      func() bool { return BoolPtrEquals(e.PooledDatasource, r.PooledDatasource) },
			expected: strconv.FormatBool(BoolDeref(e.PooledDatasource)),
			received: strconv.FormatBool(BoolDeref(r.PooledDatasource)),
		},
		{
			name:     "saltlength",
			cmp:      func() bool { return Int32PtrEquals(e.SaltLength, r.SaltLength) },
			expected: strconv.Itoa(int(Int32Deref(e.SaltLength))),
			received: strconv.Itoa(int(Int32Deref(r.SaltLength))),
		},
		{
			name:     "saltvalue",
			cmp:      func() bool { return StrPtrEquals(e.SaltValue, r.SaltValue) },
			expected: StrDeref(e.SaltValue),
			received: StrDeref(r.SaltValue),
		},
		{
			name:     "username",
			cmp:      func() bool { return StrPtrEquals(e.Username, r.Username) },
			expected: StrDeref(e.Username),
			received: StrDeref(r.Username),
		},
	}
}

//Fields to validate after DbIdentityVault update
func DbIdentityVaultFieldTestUpdate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
		{
			name:     "elementid",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
	}
	return append(t, DbIdentityVaultFieldTestCreate(e, r)...)
}

// Compares the expected DbIdentityVault with the received one.
func DbIdentityVaultValidateCreate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) error {

	return ValidateFields(DbIdentityVaultFieldTestCreate(e, r))
}

// Compares the expected DbIdentityVault with the received one.
func DbIdentityVaultValidateUpdate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) error {

	return ValidateFields(DbIdentityVaultFieldTestUpdate(e, r))
}
