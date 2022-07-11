package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliSharePointResource_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "app-0"
	var orig *api.SharepointResourceDTO
	var created api.SharepointResourceDTO
	orig = createTestSharePointExecutionEnvironmentDTO(crudName)
	if err != nil {
		t.Error(err)
		return
	}

	// Test CREATE
	created, err = s.client.CreateSharePointresource(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := SharePointExeEnvValidateCreate(orig, &created); err != nil {
		t.Errorf("creating appexecenv : %v", err)
		return
	}

	// Test READ
	var read api.SharepointResourceDTO
	read, err = s.client.GetSharePointResource(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ShPointExeEnvFieldTestUpdate(&read, &created); err != nil {
		t.Errorf("creating appexecenv : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	updated, err := s.client.UpdateSharePointResource(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ShPointExeEnvFieldTestUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteSharePointResource(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetIssExeEnvs(*appliance.Name)
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
	var listOfCreated [2]api.SharepointResourceDTO
	// Test list of #2 elements
	element1 := createTestSharePointExecutionEnvironmentDTO("app-1")
	listOfCreated[0], _ = s.client.CreateSharePointresource(*appliance.Name, *element1)

	element2 := createTestSharePointExecutionEnvironmentDTO("app-2")
	listOfCreated[1], _ = s.client.CreateSharePointresource(*appliance.Name, *element2)
	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetSharePointResources(*appliance.Name)
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
	sort.Slice(listOfRead,
		func(i, j int) bool {
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) < 0
		},
	)

	// Validate each element from the list of created with the list of read
	for idx, r := range listOfCreated {
		if err = ShPointExeEnvFieldTestUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestSharePointExecutionEnvironmentDTO(name string) *api.SharepointResourceDTO {
	tData := api.NewSharepointResourceDTO()

	var loca api.LocationDTO
	loca.SetContext("myapp")
	loca.SetHost("mycompany")
	loca.SetPort(8080)
	loca.SetProtocol("http")
	loca.SetUri(strings.ToUpper(name))
	tData.SetSloLocation(loca)

	tData.SetDescription(fmt.Sprintf("sharePoint description %s", name))
	tData.SetName(name)
	tData.SetStsSigningCertSubject("")
	tData.SetSloLocationEnabled(false)
	tData.SetStsEncryptingCertSubject("")
	return tData
}

func (s *AccTestSuite) TestAccCliSharePointExeEnvField_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccSharePointExeEnvField_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func SharePointExeEnvValidateUpdate(
	e *api.SharepointResourceDTO,
	r *api.SharepointResourceDTO) []FiledTestStruct {

	return []FiledTestStruct{

		{
			name:     "slo_location",
			cmp:      func() bool { return LocationPtrEquals(e.SloLocation, r.SloLocation) },
			expected: LocationToStr(e.SloLocation),
			received: LocationToStr(r.SloLocation),
		},
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
			name:     "sts_signing_cert_subject",
			cmp:      func() bool { return StrPtrEquals(e.StsSigningCertSubject, r.StsSigningCertSubject) },
			expected: StrDeref(e.StsSigningCertSubject),
			received: StrDeref(r.StsSigningCertSubject),
		},
		{
			name:     "slo_location_enabled",
			cmp:      func() bool { return BoolPtrEquals(e.SloLocationEnabled, r.SloLocationEnabled) },
			expected: strconv.FormatBool(BoolDeref(e.SloLocationEnabled)),
			received: strconv.FormatBool(BoolDeref(r.SloLocationEnabled)),
		},
		{
			name:     "sts_encrypting_cert_subject",
			cmp:      func() bool { return StrPtrEquals(e.StsEncryptingCertSubject, r.StsEncryptingCertSubject) },
			expected: StrDeref(e.StsEncryptingCertSubject),
			received: StrDeref(r.StsEncryptingCertSubject),
		},
	}
}

//Fields to validate after IssExeEnv update
func SharePointExeEnvFieldTestUpdate(
	e *api.SharepointResourceDTO,
	r *api.SharepointResourceDTO) []FiledTestStruct {

	t := []FiledTestStruct{}
	return append(t, SharePointExeEnvValidateUpdate(e, r)...)
}

// Compares the expected IssExeEnv with the received one.
func SharePointExeEnvValidateCreate(
	e *api.SharepointResourceDTO,
	r *api.SharepointResourceDTO) error {

	return ValidateFields(SharePointExeEnvFieldTestUpdate(e, r))
}

// Compares the expected ExternalSaml2Sp with the received one.
func ShPointExeEnvFieldTestUpdate(
	e *api.SharepointResourceDTO,
	r *api.SharepointResourceDTO) error {

	return ValidateFields(SharePointExeEnvFieldTestUpdate(e, r))
}
