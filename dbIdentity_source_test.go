package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliDbIdentitySourceDTO_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "DdIdentityVauld-a"
	var orig *api.DbIdentitySourceDTO
	var created api.DbIdentitySourceDTO
	orig = createTestDbIdentitySourceDTO(crudName)

	// Test CREATE
	created, err = s.client.createDbIdentitySourceDTO(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := DbIdentitySourceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating db identity source : %v", err)
		return
	}

	// Test READ
	var read api.DbIdentitySourceDTO
	read, err = s.client.GetDbIdentitySourceDTO(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = DbIdentitySourceDTOValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating db identity source : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	updated, err := s.client.UpdateDbIdentitySourceDTO(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = DbIdentitySourceDTOValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteDbIdentitySourceDTO(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetDbIdentitySourceDTOs(*appliance.Name)
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
	var listOfCreated [2]api.DbIdentitySourceDTO

	// Test list of #2 elements
	element1 := createTestDbIdentitySourceDTO("DbIdentitySource-2")
	listOfCreated[0], _ = s.client.createDbIdentitySourceDTO(*appliance.Name, *element1)

	element2 := createTestDbIdentitySourceDTO("DbIdentitySource-2")
	listOfCreated[1], _ = s.client.createDbIdentitySourceDTO(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetDbIdentitySourceDTOs(*appliance.Name)
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
		if err = DbIdentitySourceDTOValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestDbIdentitySourceDTO(name string) *api.DbIdentitySourceDTO {
	orig := api.NewDbIdentitySourceDTO()
	var cust api.CustomClassDTO
	var custprop []api.CustomClassPropertyDTO
	custprop1 := api.NewCustomClassPropertyDTO()
	custprop1.SetName("")
	custprop1.SetValue("")
	custprop = append(custprop, *custprop1)
	custprop2 := api.NewCustomClassPropertyDTO()
	custprop2.SetName("")
	custprop2.SetValue("")
	custprop = append(custprop, *custprop1)

	cust.SetCardinality("")
	cust.SetFqcn("")
	cust.SetOsgiFilter("")
	cust.SetOsgiService(true)
	cust.SetProperties(custprop)
	cust.SetTimeoutSecs(1)
	cust.SetType("")

	orig.SetAcquireIncrement(1)
	orig.SetAdmin("")
	orig.SetConnectionUrl("")
	orig.SetCredentialsQueryString("")
	orig.SetCustomClass(cust)
	orig.SetDescription("")
	orig.SetDriverName("")
	orig.SetElementId("")
	orig.SetIdleConnectionTestPeriod(1)
	orig.SetInitialPoolSize(1)
	orig.SetMaxIdleTime(1)
	orig.SetMaxPoolSize(1)
	orig.SetMinPoolSize(1)
	orig.SetPassword("")
	orig.SetPooledDatasource(true)
	orig.SetRelayCredentialQueryString("")
	orig.SetResetCredentialDml("")
	orig.SetRolesQueryString("")
	orig.SetUseColumnNamesAsPropertyNames(true)
	orig.SetUserPropertiesQueryString("")
	orig.SetUserQueryString("")
	orig.SetName(name)
	orig.SetId(-1)
	return orig
}

func (s *AccTestSuite) TestAccCliDbIdentitySourceDTO_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliDbIdentitySourceDTO_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func DbIdentitySourceFieldTestCreate(
	e *api.DbIdentitySourceDTO,
	r *api.DbIdentitySourceDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "acquireincrement",
			cmp:      func() bool { return Int32PtrEquals(e.AcquireIncrement, r.AcquireIncrement) },
			expected: strconv.Itoa(int(Int32Deref(e.AcquireIncrement))),
			received: strconv.Itoa(int(Int32Deref(r.AcquireIncrement))),
		},
		{
			name:     "admin",
			cmp:      func() bool { return StrPtrEquals(e.Admin, r.Admin) },
			expected: StrDeref(e.Admin),
			received: StrDeref(r.Admin),
		},
		{
			name:     "connecturl",
			cmp:      func() bool { return StrPtrEquals(e.ConnectionUrl, r.ConnectionUrl) },
			expected: StrDeref(e.ConnectionUrl),
			received: StrDeref(r.ConnectionUrl),
		},
		{
			name:     "credentialsquerystring",
			cmp:      func() bool { return StrPtrEquals(e.CredentialsQueryString, r.CredentialsQueryString) },
			expected: StrDeref(e.CredentialsQueryString),
			received: StrDeref(r.CredentialsQueryString),
		},
		//{
		//	name:     "customclass",
		//	cmp:      func() bool { return StrPtrEquals(e.customclass, r.customclass) },
		//	expected: StrDeref(e.customclass),
		//	received: StrDeref(r.customclass),
		//},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "drivername",
			cmp:      func() bool { return StrPtrEquals(e.DriverName, r.DriverName) },
			expected: StrDeref(e.DriverName),
			received: StrDeref(r.DriverName),
		},
		{
			name:     "elementid",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
		{
			name:     "idleconnectiontestperiod",
			cmp:      func() bool { return Int32PtrEquals(e.IdleConnectionTestPeriod, r.IdleConnectionTestPeriod) },
			expected: strconv.Itoa(int(Int32Deref(e.IdleConnectionTestPeriod))),
			received: strconv.Itoa(int(Int32Deref(r.IdleConnectionTestPeriod))),
		},
		{
			name:     "initialpollsize",
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
			name:     "password",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "pooleddatasource",
			cmp:      func() bool { return BoolPtrEquals(e.PooledDatasource, r.PooledDatasource) },
			expected: strconv.FormatBool(BoolDeref(e.PooledDatasource)),
			received: strconv.FormatBool(BoolDeref(r.PooledDatasource)),
		},
		{
			name:     "relaycredentialquerystring",
			cmp:      func() bool { return StrPtrEquals(e.RelayCredentialQueryString, r.RelayCredentialQueryString) },
			expected: StrDeref(e.RelayCredentialQueryString),
			received: StrDeref(r.RelayCredentialQueryString),
		},
		{
			name:     "resetcretentialdml",
			cmp:      func() bool { return StrPtrEquals(e.ResetCredentialDml, r.ResetCredentialDml) },
			expected: StrDeref(e.ResetCredentialDml),
			received: StrDeref(r.ResetCredentialDml),
		},
		{
			name:     "rolesquerystring",
			cmp:      func() bool { return StrPtrEquals(e.RolesQueryString, r.RolesQueryString) },
			expected: StrDeref(e.RolesQueryString),
			received: StrDeref(r.RolesQueryString),
		},
		{
			name:     "usecolumnnamesaspropertynames",
			cmp:      func() bool { return BoolPtrEquals(e.UseColumnNamesAsPropertyNames, r.UseColumnNamesAsPropertyNames) },
			expected: strconv.FormatBool(BoolDeref(e.UseColumnNamesAsPropertyNames)),
			received: strconv.FormatBool(BoolDeref(r.UseColumnNamesAsPropertyNames)),
		},
		{
			name:     "userpropertiesquerystring",
			cmp:      func() bool { return StrPtrEquals(e.UserPropertiesQueryString, r.UserPropertiesQueryString) },
			expected: StrDeref(e.UserPropertiesQueryString),
			received: StrDeref(r.UserPropertiesQueryString),
		}, {
			name:     "userquerystring",
			cmp:      func() bool { return StrPtrEquals(e.UserQueryString, r.UserQueryString) },
			expected: StrDeref(e.UserQueryString),
			received: StrDeref(r.UserQueryString),
		},
	}
}

//Fields to validate after DbIdentitySource update
func DbIdentitySourceFieldTestUpdate(
	e *api.DbIdentitySourceDTO,
	r *api.DbIdentitySourceDTO) []FiledTestStruct {

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
	return append(t, DbIdentitySourceFieldTestCreate(e, r)...)
}

// Compares the expected DbIdentitySource with the received one.
func DbIdentitySourceValidateCreate(
	e *api.DbIdentitySourceDTO,
	r *api.DbIdentitySourceDTO) error {

	return ValidateFields(DbIdentitySourceFieldTestCreate(e, r))
}

// Compares the expected DbIdentitySource with the received one.
func DbIdentitySourceDTOValidateUpdate(
	e *api.DbIdentitySourceDTO,
	r *api.DbIdentitySourceDTO) error {

	return ValidateFields(DbIdentitySourceFieldTestUpdate(e, r))
}
