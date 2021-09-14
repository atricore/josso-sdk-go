package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	pprof "runtime/pprof"

	api "github.com/atricore/josso-api-go"
	"github.com/stretchr/testify/suite"
)

type AccTestSuite struct {
	suite.Suite
	client *IdbusApiClient
	idx    int
}

func (s *AccTestSuite) nextIdx() int {
	s.idx++
	return s.idx
}

func (s *AccTestSuite) DumpHeap(prefix string) error {

	f, err := os.Create(buildFileName(prefix, s.nextIdx()))
	if err != nil {
		return err
	}

	pprof.WriteHeapProfile(f)

	return nil
}

func buildFileName(prefix string, suffix int) string {
	return fmt.Sprintf("%s-heap-%s-%d.prof", prefix, time.Now().Format("20060102150405"), suffix)
}

const (
	keystore          = "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"
	testApplianceName = "testacc-a"
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
		return
	}

	t.Log("SetupSuite complete")

}

func (s *AccTestSuite) TearDownSuite() {
	var t = s.T()
	t.Log("Teardown suite")

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
		Description: api.PtrString("TESTACC-T TEST !"),
		DisplayName: api.PtrString("TESTACC-T TEST !"),
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

	v := os.Getenv("ACCTEST_CLEAR_DATA")
	s.T().Logf("ACCTEST_CLEAR_DATA: %s", v)
	if v == "" {
		v = "true"
	}
	clearData, _ := strconv.ParseBool(v)

	if clearData {
		s.T().Logf("clearing test data")
		as, err := s.client.GetAppliances()
		if err != nil {
			return err
		}

		s.T().Logf("found %d appliances", len(as))

		for _, a := range as {

			if strings.HasPrefix(*a.Name, "testacc") {

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

	} else {
		s.T().Logf("ENABLE DELETE")
	}

	return nil

}
