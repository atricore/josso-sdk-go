package cli

import (
	"errors"
	"os"
	"testing"

	api "github.com/atricore/josso-api-go"
	"github.com/stretchr/testify/suite"
)

type AccTestSuite struct {
	suite.Suite
	client *IdbusApiClient
}

const (
	testApplianceName = "ida-a"
)

func (s *AccTestSuite) SetupSuite() {
	var t = s.T()
	var err error
	t.Log("creating client")

	s.client, err = createClient()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
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

func TestMarshalJSON(t *testing.T) {

	a := api.NewIdentityApplianceDefinitionDTO()
	a.Name = api.PtrString(testApplianceName)
	a.Namespace = api.PtrString("com.atricore.idaa")
	a.Description = api.PtrString("Test A")

	a.Location = api.NewLocationDTO()
	a.Location.Host = api.PtrString("localhost")
	a.Location.Protocol = api.PtrString("http")
	a.Location.Port = api.PtrInt32(8081)

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

	trace, err := GetenvBool("JOSSO_API_TRACE")
	if err != nil {
		trace = false
	}

	l := DefaultLogger{debug: true}

	c := NewIdbusApiClient(&l, trace)
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

func getTestAppliance(t *testing.T, client *IdbusApiClient) (api.IdentityApplianceDefinitionDTO, error) {
	var read api.IdentityApplianceDefinitionDTO

	read, err := client.GetAppliance(testApplianceName)
	if err != nil {
		return read, err
	}

	if read.Name == nil || *read.Name == "" {
		read, err = createTestAppliance(t, client)
	}
	return read, err

}

func createTestAppliance(t *testing.T, client *IdbusApiClient) (api.IdentityApplianceDefinitionDTO, error) {

	var created api.IdentityApplianceDefinitionDTO

	l, _ := StrToLocation("http://localhost/IDBUS/IDA-T")
	orig := api.IdentityApplianceDefinitionDTO{
		Name:        api.PtrString(testApplianceName),
		Namespace:   api.PtrString("com.atricore.idbus.ida.t"),
		Location:    l,
		Description: api.PtrString("IDA-T TEST !"),
	}
	created, err := client.CreateAppliance(orig)
	if err != nil {
		return created, err
	}
	if err := IdApplianceValidateCreate(&orig, &created); err != nil {
		return created, err
	}

	return created, nil

}

func (s *AccTestSuite) accClearData() error {
	/*
		s.T().Logf("clearing test data")
		as, err := s.client.GetAppliances()
		if err != nil {
			return err
		}

		for _, a := range as {
			if *a.Name == testApplianceName {

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
	*/
	s.T().Logf("ENABLE DELETE")
	return nil

}
