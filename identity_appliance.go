package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

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
		msg := buildErrorMsg(*res.Error, *res.ValidationErrors)
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
		msg := buildErrorMsg(*res.Error, *res.ValidationErrors)
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
