package cli

import (
	"errors"
	"os"
	"strconv"
	"testing"

	api "github.com/atricore/josso-api-go"
	"github.com/stretchr/testify/suite"
)

type AccTestSuite struct {
	suite.Suite
	client *IdbusApiClient
}

func (s *AccTestSuite) SetupSuite() {
	var t = s.T()
	var err error
	t.Log("creating client")

	s.client, err = createClient()
	if err != nil {
		t.Errorf("failed to create client: %v", err)
		return
	}

	t.Logf("created test client: %v", s.client.config.Servers)

	if err = s.accClearData(); err != nil {
		t.Errorf("error while clearing data %v", err)
	}
}

func (s *AccTestSuite) TearDownSuite() {
	if err := s.accClearData(); err != nil {
		s.T().Errorf("error clearing data %v", err)
	}
}

// All methods that begin with "Test" are run as tests within a
// suite.
func TestAccCliSuite(t *testing.T) {
	s := new(AccTestSuite)
	suite.Run(t, s)
}

func (s *AccTestSuite) TestAccCliIdentityAppliance_crud() {
	var t = s.T()

	// Test CRUD

	// Create
	l, _ := StrToLocation("htpt://localhost")
	orig := api.IdentityApplianceDefinitionDTO{
		Name:        api.PtrString("ida-a"),
		Namespace:   api.PtrString("com.atricore.idbus.ida.a"),
		Location:    l,
		Description: api.PtrString("IDA-A TEST !"),
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
	retrieved, err := s.client.GetAppliance("ida-a")
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

func TestMarshalJSON(t *testing.T) {

	a := api.NewIdentityApplianceDefinitionDTO()
	a.Name = Cstring("ida-a")
	a.Namespace = Cstring("com.atricore.idaa")
	a.Description = Cstring("Test A")

	a.Location = api.NewLocationDTO()
	a.Location.Host = Cstring("localhost")
	a.Location.Protocol = Cstring("http")
	a.Location.Port = Cint32(8081)

	json, err := a.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshalling to JSON %v", err)
	}

	b := api.NewIdentityApplianceDefinitionDTO()
	err = b.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("Error unmarshalling to JSON %v", err)
	}

	if b.Name == nil {
		t.Errorf("Name not found")
	} else if *a.Name != *b.Name {
		t.Errorf("Invalid name, %s, expeded %s", *b.Name, *a.Name)
	}

	if b.Namespace == nil {
		t.Errorf("Name not found")
	} else if *a.Namespace != *b.Namespace {
		t.Errorf("Invalid name, %s, expeded %s", *b.Namespace, *a.Namespace)
	}

}

// ----------------------------------

// Creates a new TEST client.  You can enable/disable debugging and message exchage tracing
func createClient() (*IdbusApiClient, error) {

	s, err := accPreCheck()
	if err != nil {
		return nil, err
	}

	c := NewIdbusApiClient(&DefaultLogger{debug: true}, false)
	err = c.RegisterServer(s, "")
	if err != nil {
		return nil, err
	}

	err = c.Authn()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Check environment and build server configuration
func accPreCheck() (*IdbusServer, error) {

	clientSecret := os.Getenv("JOSSO_API_SECRET")
	clientId := os.Getenv("JOSSO_API_CLIENT_ID")
	endpoint := os.Getenv("JOSSO_API_ENDPOINT")
	username := os.Getenv("JOSSO_API_USERNAME")
	password := os.Getenv("JOSSO_API_PASSWORD")
	if clientSecret == "" || clientId == "" || endpoint == "" || username == "" || password == "" {
		return nil, errors.New("JOSSO variables must be set for acceptance tests")
	}

	s := IdbusServer{
		Config: &api.ServerConfiguration{
			URL:         endpoint,
			Description: "JOSSO Test server",
		},
		Credentials: &ServerCredentials{
			ClientId: clientId,
			Secret:   clientSecret,
			Username: username,
			Password: password,
		},
	}
	return &s, nil
}

func (s *AccTestSuite) accClearData() error {
	s.T().Logf("clearing test data")
	as, err := s.client.GetAppliances()
	if err != nil {
		return err
	}

	for _, a := range as {
		if *a.Name == "ida-a" {
			s.T().Logf("deleting appliance %d", a.GetId())
			r, err1 := s.client.DeleteAppliance(strconv.FormatInt(a.GetId(), 10))
			if err1 != nil {
				return err
			}
			if !r {
				s.T().Logf("appliance not found %s", *a.Name)
				return nil
			}

		}
	}
	return nil

}

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
