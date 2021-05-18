package cli

import (
	"context"
	"errors"
	"fmt"

	api "github.com/atricore/josso-api-go"
)

const (
	DEFAULT_SVR = "__default__"
)

type (
	IdbusApiClient struct {
		apiClient *api.APIClient
		config    *api.Configuration
		servers   map[string]*ServerConnection
		logger    Logger
	}

	IdbusServer struct {
		Config      *api.ServerConfiguration
		Credentials *ServerCredentials
	}

	ServerCredentials struct {
		ClientId string
		Secret   string
		Username string
		Password string
	}

	ServerConnection struct {
		Authn  ServerAuthn
		Server *IdbusServer
	}

	ServerAuthn struct {
		AccessToken  string
		RefreshToken string
	}
)

func NewIdbusApiClientWithDefaults() *IdbusApiClient {
	cfg := config(false)
	cli := &IdbusApiClient{
		config:    cfg,
		apiClient: api.NewAPIClient(cfg),
		servers:   make(map[string]*ServerConnection),
		logger:    &DefaultLogger{debug: true},
	}
	return cli
}

func NewIdbusApiClient(l Logger, trace bool) *IdbusApiClient {
	cfg := config(trace)
	cli := &IdbusApiClient{
		config:    cfg,
		apiClient: api.NewAPIClient(cfg),
		servers:   make(map[string]*ServerConnection),
		logger:    l,
	}
	return cli
}

func (c *IdbusApiClient) Logger() Logger {
	return c.logger
}

/*
* Register a new server
 */
func (c *IdbusApiClient) RegisterServer(svr *IdbusServer, operation string) error {

	key := operation
	if key == "" {
		key = DEFAULT_SVR
	}
	c.logger.Tracef("registering server %s", svr.Config.URL)

	// We replace configuration if the server is already registerd for the URL
	sc := ServerConnection{
		Server: svr,
	}
	if ok := c.servers[key]; ok != nil {
		c.logger.Tracef("replacing server registration")
		found := false
		for _, sc := range c.apiClient.GetConfig().Servers {
			if sc.URL == svr.Config.URL {
				c.logger.Tracef("replacing server configuration for %s", sc.URL)
				sc.Description = svr.Config.Description
				sc.Variables = svr.Config.Variables
				found = true
				break
			}
		}
		if !found {
			c.logger.Errorf("server registered, but config not found for %s", key)
			return fmt.Errorf("server registered, but config not found for %s", key)
		}
	} else {
		c.logger.Tracef("adding server configuration for %s", svr.Config.URL)
		c.apiClient.GetConfig().Servers = append(c.apiClient.GetConfig().Servers, *svr.Config)

	}
	c.servers[key] = &sc

	if operation != "" {
		scs := c.apiClient.GetConfig().OperationServers[operation]
		scs = append(scs, *svr.Config)
		c.apiClient.GetConfig().OperationServers[operation] = scs
	}

	return nil
}

func (c *IdbusApiClient) Authn() error {

	sc, err := c.IdbusServerForOperation("DefaultApiService.SignOn") // Also hard-coded in generated openapi
	if err != nil {
		return err
	}

	c.logger.Tracef("authn: %s %t/%s %t",
		sc.Server.Credentials.ClientId,
		sc.Server.Credentials.Secret != "",
		sc.Server.Credentials.Username,
		sc.Server.Credentials.Password != "")

	req := c.apiClient.DefaultApi.SignOn(context.Background())
	req = req.OIDCSignOnRequest(api.OIDCSignOnRequest{
		ClientId: &sc.Server.Credentials.ClientId,
		Secret:   &sc.Server.Credentials.Secret,
		Username: &sc.Server.Credentials.Username,
		Password: &sc.Server.Credentials.Password})

	res, _, err := c.apiClient.DefaultApi.SignOnExecute(req)
	if err != nil {
		return err
	}

	sc.Authn.AccessToken = *res.AccessToken
	sc.Authn.RefreshToken = *res.RefreshToken

	return nil

}

// Creates a new identity appliance. Name must not exist, even in other namespaces.  Namespaces must also be unique.
// It returs the created appliance object.
func (c *IdbusApiClient) CreateAppliance(appliance api.IdentityApplianceDefinitionDTO) (api.IdentityApplianceDefinitionDTO, error) {
	var result api.IdentityApplianceDefinitionDTO

	if appliance.Name == nil || appliance.Namespace == nil {
		return result, errors.New("appliance name and namespace are required")
	}
	c.logger.Debugf("creating identity appliance : %s %s", *appliance.Name, *appliance.Namespace)

	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateAppliance") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateAppliance(ctx)
	req = req.CreateApplianceReq(api.CreateApplianceReq{Appliance: &appliance})
	res, _, err := c.apiClient.DefaultApi.CreateApplianceExecute(req)
	if err != nil {
		c.logger.Errorf("createAppliance. Error %v", err)
		return result, err

	}
	if res.Error != nil {

		var msg string
		if res.ValidationErrors != nil && len(*res.ValidationErrors) > 0 {
			msg = fmt.Sprintf("%s : %#v", *res.Error, *res.ValidationErrors)
		} else {
			msg = *res.Error
		}

		c.logger.Errorf("createAppliance. Error %s", msg)
		return result, errors.New(msg)
	}
	c.logger.Debugf("createAppliance. ID: %d", *res.Appliance.Id)

	result = *res.Appliance
	return result, err

}

// Updates a new identity appliance. Name must not exist, even in other namespaces.  Namespaces must also be unique.
// It returs the created appliance object.
func (c *IdbusApiClient) UpdateAppliance(appliance api.IdentityApplianceDefinitionDTO) (api.IdentityApplianceDefinitionDTO, error) {

	var result api.IdentityApplianceDefinitionDTO

	c.logger.Debugf("updating identity appliance : %s %s", *appliance.Name, *appliance.Namespace)

	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateAppliance") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateAppliance(ctx)
	req = req.UpdateApplianceReq(api.UpdateApplianceReq{Appliance: &appliance})
	res, _, err := c.apiClient.DefaultApi.UpdateApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("updateAppliance. Error %v", err)
	}
	if res.Error != nil {
		c.logger.Errorf("updateAppliance. Error %v", err)
		return result, errors.New(*res.Error)
	}

	result = *res.Appliance
	c.logger.Debugf("updateAppliance. Updated: %d", *res.Appliance.Id)

	return result, err

}

func (c *IdbusApiClient) GetAppliances() ([]api.IdentityApplianceDefinitionDTO, error) {
	sc, err := c.IdbusServerForOperation("DefaultApiService.GetAppliances") // Also hard-coded in generated client
	if err != nil {
		return nil, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetAppliances(ctx)
	res, _, err := c.apiClient.DefaultApi.GetAppliancesExecute(req)

	if err != nil {
		c.logger.Errorf("getAppliances. Error %v", err)
		return nil, err
	}

	if res.Error != nil {
		c.logger.Errorf("getAppliances. Error %s", *res.Error)
		return nil, errors.New(*res.Error)
	}

	c.logger.Debugf("found appliances %d", len(*res.Appliances))
	return *res.Appliances, nil

}

func (c *IdbusApiClient) GetAppliance(idOrName string) (api.IdentityApplianceDefinitionDTO, error) {

	var result api.IdentityApplianceDefinitionDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetAppliance") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("getAppliance. Error %v", err)
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)

	req := c.apiClient.DefaultApi.GetAppliance(ctx)
	req = req.GetApplianceReq(api.GetApplianceReq{IdOrName: &idOrName})
	res, _, err := c.apiClient.DefaultApi.GetApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("getAppliance. Error %v", err)
	}

	if res.Error != nil {
		c.logger.Errorf("getAppliance. Error %v", *res.Error)
		return result, errors.New(*res.Error)
	}

	if res.Appliance != nil {
		result = *res.Appliance
		c.logger.Debugf("appliance ID %d found for ID/name %s", *result.Id, idOrName)
	} else {
		c.logger.Debugf("appliance ID not found for ID/name %s", idOrName)
	}

	return result, err
}

func (c *IdbusApiClient) DeleteAppliance(id string) (bool, error) {

	c.logger.Debugf("delete appliance id: %s", id)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteAppliance") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteAppliance. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteAppliance(ctx)
	req = req.DeleteApplianceReq(api.DeleteApplianceReq{Id: &id})
	res, _, err := c.apiClient.DefaultApi.DeleteApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("deleteAppliance. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteAppliance. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteAppliance. Deleted %s : %t", id, *res.Removed)

	return *res.Removed, err
}

// Create default configuration
func config(debug bool) *api.Configuration {
	return &api.Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            debug,
		Servers:          api.ServerConfigurations{},
		OperationServers: make(map[string]api.ServerConfigurations), // Servers for specific operations
	}
}

func (c *IdbusApiClient) IdbusServerForOperation(operation string) (*ServerConnection, error) {
	sc, ok := c.servers[operation]
	if ok {
		return sc, nil
	} else {
		return c.servers[DEFAULT_SVR], nil
	}
}

func Cstring(v string) *string {
	return &v
}

func Cint32(v int32) *int32 {
	return &v
}
