package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliDbIdentityVaultDto_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "DdIdentityVault-a"
	var orig *api.DbIdentityVaultDTO
	var created api.DbIdentityVaultDTO
	orig = createTestDbIdentityVaultDto(crudName)

	// Test CREATE
	created, err = s.client.CreateDbIdentityVault(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := DbIdentityVaultValidateCreate(orig, &created); err != nil {
		t.Errorf("creating dbv identity vault : %v", err)
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
		t.Errorf("reading dbv identity vault : %v", err)
		return
	}
	if read.Name == nil {
		t.Errorf("dbv identity vault not found for name %s", crudName)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")

	updated, err := s.client.UpdateDbIdentityVaultDTO(*appliance.Name, read)
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
	listOfCreated[0], _ = s.client.CreateDbIdentityVault(*appliance.Name, *element1)

	element2 := createTestDbIdentityVaultDto("DbIdentityVault-3")
	listOfCreated[1], _ = s.client.CreateDbIdentityVault(*appliance.Name, *element2)

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
	tData := api.NewDbIdentityVaultDTO()
	tData.SetAcquireIncrement(1)
	tData.SetConnectionUrl(fmt.Sprintf("jdbc:mysql:localhost/%s?create=true", name))
	tData.SetDescription(fmt.Sprint("Description", name))
	tData.SetDriverName(fmt.Sprintln("org.mysql.driver"))
	tData.SetExternalDB(true)
	tData.SetHashAlgorithm("SHA267")
	tData.SetHashEncoding("BASE64")
	tData.SetIdleConnectionTestPeriod(1)
	tData.SetInitialPoolSize(10)
	tData.SetMaxIdleTime(15)
	tData.SetMaxPoolSize(20)
	tData.SetMinPoolSize(1)
	tData.SetName(name)
	tData.SetPassword(fmt.Sprint("pdw", name))
	tData.SetPooledDatasource(true)
	tData.SetSaltLength(55)
	tData.SetSaltValue("salt#")
	tData.SetUsername(fmt.Sprint("db", name))
	return tData
}

func (s *AccTestSuite) TestAccCliDbIdentitySourceDto_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

// Fields to validate after appliance creation
func DbIdentityVaultFieldTestCreate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "acquire_increment",
			cmp:      func() bool { return Int32PtrEquals(e.AcquireIncrement, r.AcquireIncrement) },
			expected: strconv.Itoa(int(Int32Deref(e.AcquireIncrement))),
			received: strconv.Itoa(int(Int32Deref(r.AcquireIncrement))),
		},
		{
			name:     "connection_url",
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
			name:     "driver_name",
			cmp:      func() bool { return StrPtrEquals(e.DriverName, r.DriverName) },
			expected: StrDeref(e.DriverName),
			received: StrDeref(r.DriverName),
		},
		{
			name:     "external_db",
			cmp:      func() bool { return BoolPtrEquals(e.ExternalDB, r.ExternalDB) },
			expected: strconv.FormatBool(BoolDeref(e.ExternalDB)),
			received: strconv.FormatBool(BoolDeref(r.ExternalDB)),
		},
		{
			name:     "hash_algorith",
			cmp:      func() bool { return StrPtrEquals(e.HashAlgorithm, r.HashAlgorithm) },
			expected: StrDeref(e.HashAlgorithm),
			received: StrDeref(r.HashAlgorithm),
		},
		{
			name:     "hash_encoding",
			cmp:      func() bool { return StrPtrEquals(e.HashEncoding, r.HashEncoding) },
			expected: StrDeref(e.HashEncoding),
			received: StrDeref(r.HashEncoding),
		},
		{
			name:     "idle_connection_test_period",
			cmp:      func() bool { return Int32PtrEquals(e.IdleConnectionTestPeriod, r.IdleConnectionTestPeriod) },
			expected: strconv.Itoa(int(Int32Deref(e.IdleConnectionTestPeriod))),
			received: strconv.Itoa(int(Int32Deref(r.IdleConnectionTestPeriod))),
		},
		{
			name:     "initial_pool_size",
			cmp:      func() bool { return Int32PtrEquals(e.InitialPoolSize, r.InitialPoolSize) },
			expected: strconv.Itoa(int(Int32Deref(e.InitialPoolSize))),
			received: strconv.Itoa(int(Int32Deref(r.InitialPoolSize))),
		},
		{
			name:     "max_idle_time",
			cmp:      func() bool { return Int32PtrEquals(e.MaxIdleTime, r.MaxIdleTime) },
			expected: strconv.Itoa(int(Int32Deref(e.MaxIdleTime))),
			received: strconv.Itoa(int(Int32Deref(r.MaxIdleTime))),
		},
		{
			name:     "max_pool_size",
			cmp:      func() bool { return Int32PtrEquals(e.MaxPoolSize, r.MaxPoolSize) },
			expected: strconv.Itoa(int(Int32Deref(e.MaxPoolSize))),
			received: strconv.Itoa(int(Int32Deref(r.MaxPoolSize))),
		},
		{
			name:     "min_pool_size",
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
			name:     "pooled_data_source",
			cmp:      func() bool { return BoolPtrEquals(e.PooledDatasource, r.PooledDatasource) },
			expected: strconv.FormatBool(BoolDeref(e.PooledDatasource)),
			received: strconv.FormatBool(BoolDeref(r.PooledDatasource)),
		},
		{
			name:     "salt_length",
			cmp:      func() bool { return Int32PtrEquals(e.SaltLength, r.SaltLength) },
			expected: strconv.Itoa(int(Int32Deref(e.SaltLength))),
			received: strconv.Itoa(int(Int32Deref(r.SaltLength))),
		},
		{
			name:     "salt_value",
			cmp:      func() bool { return StrPtrEquals(e.SaltValue, r.SaltValue) },
			expected: StrDeref(e.SaltValue),
			received: StrDeref(r.SaltValue),
		},
		{
			name:     "user_name",
			cmp:      func() bool { return StrPtrEquals(e.Username, r.Username) },
			expected: StrDeref(e.Username),
			received: StrDeref(r.Username),
		},
	}
}

// Fields to validate after DbIdentityVault update
func DbIdentityVaultFieldTestUpdate(
	e *api.DbIdentityVaultDTO,
	r *api.DbIdentityVaultDTO) []FiledTestStruct {

	t := []FiledTestStruct{}
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
