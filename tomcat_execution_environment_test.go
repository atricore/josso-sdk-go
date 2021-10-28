package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliTomcat_Exec_Env() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Tomcat-Exect-Env"
	var orig *api.TomcatExecutionEnvironmentDTO
	var created api.TomcatExecutionEnvironmentDTO
	orig = createTestTomcatExecutionEnvironmentDTO(crudName)
	if err != nil {
		t.Error(err)
		return
	}

	// Test CREATE
	created, err = s.client.CreateTomcatExeEnv(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := TomcatExeEnvValidateCreate(orig, &created); err != nil {
		t.Errorf("creating vp : %v", err)
		return
	}

	// Test READ
	var read api.TomcatExecutionEnvironmentDTO
	read, err = s.client.GetTomcatExeEnv(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = TomcatExeEnvValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating vp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateTomcatExeEnv(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = TomcatExeEnvValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteTomcatExeEnv(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetTomcatExeEnvs(*appliance.Name)
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
	var listOfCreated [2]api.TomcatExecutionEnvironmentDTO
	// Test list of #2 elements
	element1 := createTestTomcatExecutionEnvironmentDTO("tomcat-1")
	listOfCreated[0], _ = s.client.CreateTomcatExeEnv(*appliance.Name, *element1)

	element2 := createTestTomcatExecutionEnvironmentDTO("tomcat-2")
	listOfCreated[1], _ = s.client.CreateTomcatExeEnv(*appliance.Name, *element2)
	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetTomcatExeEnvs(*appliance.Name)
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
		if err = TomcatExeEnvValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestTomcatExecutionEnvironmentDTO(name string) *api.TomcatExecutionEnvironmentDTO {
	tData := api.NewTomcatExecutionEnvironmentDTO()

	tData.SetActive(true)
	// tData.SetDescription()
	// tData.SetDisplayName()
	// tData.SetElementId()
	// tData.SetId()
	// tData.SetInstallDemoApps()
	// tData.SetinstallUri()
	// tData.SetName(name)
	// tData.SetOverwriteOriginalSetup()
	// tData.SetPlatformId()
	// tData.SetTargetJDK()
	// tData.SetType()

	return tData
}

func (s *AccTestSuite) TestAccCliTomcatExeEnvField_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccTomcatExeEnvField_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func TomcatExeEnvFieldTestCreate(
	e *api.TomcatExecutionEnvironmentDTO,
	r *api.TomcatExecutionEnvironmentDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "active",
			cmp:      func() bool { return BoolPtrEquals(e.Active, r.Active) },
			expected: strconv.FormatBool(BoolDeref(e.Active)),
			received: strconv.FormatBool(BoolDeref(r.Active)),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "display_name",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: StrDeref(e.DisplayName),
			received: StrDeref(r.DisplayName),
		},
		{
			name:     "element_id",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
		{
			name:     "install_demo_apps",
			cmp:      func() bool { return BoolPtrEquals(e.InstallDemoApps, r.InstallDemoApps) },
			expected: strconv.FormatBool(BoolDeref(e.InstallDemoApps)),
			received: strconv.FormatBool(BoolDeref(r.InstallDemoApps)),
		},
		{
			name:     "install_uri",
			cmp:      func() bool { return StrPtrEquals(e.InstallUri, r.InstallUri) },
			expected: StrDeref(e.InstallUri),
			received: StrDeref(r.InstallUri),
		},
		{
			name:     "location",
			cmp:      func() bool { return StrPtrEquals(e.Location, r.Location) },
			expected: StrDeref(e.Location),
			received: StrDeref(r.Location),
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "overwrite_original_setup",
			cmp:      func() bool { return BoolPtrEquals(e.OverwriteOriginalSetup, r.OverwriteOriginalSetup) },
			expected: strconv.FormatBool(BoolDeref(e.OverwriteOriginalSetup)),
			received: strconv.FormatBool(BoolDeref(r.OverwriteOriginalSetup)),
		},
		{
			name:     "plataform_id",
			cmp:      func() bool { return StrPtrEquals(e.PlatformId, r.PlatformId) },
			expected: StrDeref(e.PlatformId),
			received: StrDeref(r.PlatformId),
		},
		{
			name:     "target_jdk",
			cmp:      func() bool { return StrPtrEquals(e.TargetJDK, r.TargetJDK) },
			expected: StrDeref(e.TargetJDK),
			received: StrDeref(r.TargetJDK),
		},
		{
			name:     "type",
			cmp:      func() bool { return StrPtrEquals(e.Type, r.Type) },
			expected: StrDeref(e.Type),
			received: StrDeref(r.Type),
		},
	}
}

//Fields to validate after TomcatExeEnv update
func TomcatExeEnvFieldTestUpdate(
	e *api.TomcatExecutionEnvironmentDTO,
	r *api.TomcatExecutionEnvironmentDTO) []FiledTestStruct {

	t := []FiledTestStruct{}
	return append(t, TomcatExeEnvFieldTestCreate(e, r)...)
}

// Compares the expected TomcatExeEnv with the received one.
func TomcatExeEnvValidateCreate(
	e *api.TomcatExecutionEnvironmentDTO,
	r *api.TomcatExecutionEnvironmentDTO) error {

	return ValidateFields(TomcatExeEnvFieldTestCreate(e, r))
}

// Compares the expected ExternalSaml2Sp with the received one.
func TomcatExeEnvValidateUpdate(
	e *api.TomcatExecutionEnvironmentDTO,
	r *api.TomcatExecutionEnvironmentDTO) error {

	return ValidateFields(TomcatExeEnvFieldTestUpdate(e, r))
}
