package cli

import (
	"context"
	"fmt"
	"log"

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
	return NewIdbusApiClient(&DefaultLogger{debug: true}, false)
}

func NewIdbusApiClient(l Logger, trace bool) *IdbusApiClient {
	l.Debugf("newIdbusApiClient TRACE: %t", trace)

	if trace {
		log.Print("Using client TRACE ON")
	}
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
