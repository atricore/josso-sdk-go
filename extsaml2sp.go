package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new SP in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateExtSaml2Sp(ida string, sp api.ExternalSaml2ServiceProviderDTO) (api.ExternalSaml2ServiceProviderDTO, error) {
	var result api.ExternalSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("createSP : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateSp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	// initSP(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateExtSaml2Sp(ctx)
	req = req.StoreExtSaml2SpReq(api.StoreExtSaml2SpReq{IdaName: &ida, Sp: &sp})
	res, _, err := c.apiClient.DefaultApi.CreateExtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("createSP. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, *res.ValidationErrors)
		c.logger.Errorf("createSP. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no sp received after creation")
	}

	result = *res.Sp

	return result, nil
}

// TODO : UPDATE, DELETE, GET , GETS
