package cli

import (
	"context"
	b64 "encoding/base64"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) ImportAppliance(applianceJson string) (api.IdentityApplianceDefinitionDTO, error) {

	var result api.IdentityApplianceDefinitionDTO

	c.logger.Debugf("Importing appliance from JSON")

	encoded := b64.StdEncoding.EncodeToString([]byte(applianceJson))

	sc, err := c.IdbusServerForOperation("DefaultApiService.ImportAppliance") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.ImportAppliance(ctx)
	req = req.ImportApplianceReq(api.ImportApplianceReq{
		Base64Json: &encoded,
		Modify:     PtrBool(false),
	})
	res, _, err := c.apiClient.DefaultApi.ImportApplianceExecute(req)
	if err != nil {
		c.logger.Errorf("importAppliance. Error %v", err)
		return result, err

	}
	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("importAppliance. Error %s", msg)
		return result, errors.New(msg)
	}
	c.logger.Debugf("importAppliance. ID: %d [%s]", *res.Appliance.Id, *res.Appliance.Name)

	result = *res.Appliance
	return result, err
}

// Creates a new identity appliance. Name must not exist, even in other namespaces.  Namespaces must also be unique.
// It returs the created appliance object.
func (c *IdbusApiClient) CreateAppliance(appliance api.IdentityApplianceDefinitionDTO) (api.IdentityApplianceDefinitionDTO, error) {
	var result api.IdentityApplianceDefinitionDTO

	if appliance.Name == nil || appliance.Namespace == nil {
		return result, errors.New("appliance name and namespace are required")
	}
	c.logger.Debugf("createAppliance : %s %s", *appliance.Name, *appliance.Namespace)

	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateAppliance") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateAppliance(ctx)
	req = req.StoreApplianceReq(api.StoreApplianceReq{Appliance: &appliance})
	res, _, err := c.apiClient.DefaultApi.CreateApplianceExecute(req)
	if err != nil {
		c.logger.Errorf("createAppliance. Error %v", err)
		return result, err

	}
	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
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

	c.logger.Debugf("updateAppliance : %s %s", *appliance.Name, *appliance.Namespace)

	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateAppliance") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateAppliance(ctx)
	req = req.StoreApplianceReq(api.StoreApplianceReq{Appliance: &appliance})
	res, _, err := c.apiClient.DefaultApi.UpdateApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("updateAppliance. Error %v", err)
	}
	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateAppliance. Error %s", msg)
		return result, errors.New(msg)
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

	c.logger.Debugf("getAppliances. found appliances %d", len(res.Appliances))
	return res.Appliances, nil

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
		c.logger.Debugf("getAppliance. %d found for ID/name %s", *result.Id, idOrName)
	} else {
		c.logger.Debugf("getAppliance. not found for ID/name %s", idOrName)
	}

	return result, err
}

func (c *IdbusApiClient) DeleteAppliance(id string) (bool, error) {

	c.logger.Debugf("deleteAppliance id: %s", id)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteAppliance") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteAppliance. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteAppliance(ctx)
	req = req.DeleteReq(api.DeleteReq{Name: &id})
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

// Received an identity appliance name and starts it.
func (c *IdbusApiClient) StartAppliance(name string) error {

	c.logger.Debugf("startAppliance id: %s", name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.StartAppliance") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("startAppliance. Error %v", err)
		return err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.StartAppliance(ctx)
	req = req.SetApplianceStateReq(api.SetApplianceStateReq{IdaName: &name})
	res, _, err := c.apiClient.DefaultApi.StartApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("startAppliance. Error %v", err)
		return err
	}

	if res.Error != nil {
		c.logger.Errorf("startAppliance. Error %v", *res.Error)
		return errors.New(*res.Error)
	}

	c.logger.Debugf("startAppliance. Deleted %s : %t", name)

	return err
}

func (c *IdbusApiClient) StopAppliance(name string) error {

	c.logger.Debugf("stopAppliance id: %s", name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.stopAppliance") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("stopAppliance. Error %v", err)
		return err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.StopAppliance(ctx)
	req = req.SetApplianceStateReq(api.SetApplianceStateReq{IdaName: &name})
	res, _, err := c.apiClient.DefaultApi.StopApplianceExecute(req)

	if err != nil {
		c.logger.Errorf("stopAppliance. Error %v", err)
		return err
	}

	if res.Error != nil {
		c.logger.Errorf("stopAppliance. Error %v", *res.Error)
		return errors.New(*res.Error)
	}

	c.logger.Debugf("stopAppliance. Deleted %s : %t", name)

	return err
}

func (c *IdbusApiClient) GetApplianceContainers() ([]api.IdentityApplianceContainerDTO, error) {
	sc, err := c.IdbusServerForOperation("DefaultApiService.GetApplianceContainers") // Also hard-coded in generated client
	if err != nil {
		return nil, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetApplianceContainers(ctx)
	res, _, err := c.apiClient.DefaultApi.GetApplianceContainersExecute(req)

	if err != nil {
		c.logger.Errorf("getApplianceContainers. Error %v", err)
		return nil, err
	}

	if res.Error != nil {
		c.logger.Errorf("getApplianceContainers. Error %s", *res.Error)
		return nil, errors.New(*res.Error)
	}

	c.logger.Debugf("getApplianceContainers. found appliances %d", len(res.Appliances))
	return res.Appliances, nil

}

func (c *IdbusApiClient) GetApplianceContainer(idOrName string) (api.IdentityApplianceContainerDTO, error) {

	var result api.IdentityApplianceContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetApplianceContainer") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("getApplianceContainer. Error %v", err)
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)

	req := c.apiClient.DefaultApi.GetApplianceContainer(ctx)
	req = req.GetApplianceReq(api.GetApplianceReq{IdOrName: &idOrName})
	res, _, err := c.apiClient.DefaultApi.GetApplianceContainerExecute(req)

	if err != nil {
		c.logger.Errorf("getApplianceContainer. Error %v", err)
	}

	if res.Error != nil {
		c.logger.Errorf("getApplianceContainer. Error %v", *res.Error)
		return result, errors.New(*res.Error)
	}

	if res.Appliance != nil {
		result = *res.Appliance
		c.logger.Debugf("getAppliance. %d found for ID/name %s", *result.GetAppliance().Id, idOrName)
	} else {
		c.logger.Debugf("getAppliance. not found for ID/name %s", idOrName)
	}

	return result, err
}
