package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Received an identity appliance name and starts it.
func (c *IdbusApiClient) ActivateExecEnv(
	appliance_id string,
	exec_env_name string,
	exec_env_folder string,
	force bool,
	replace_config bool,
	install_samples bool,
) error {

	c.logger.Debugf("activateExecEnv [%s] id: %s", appliance_id, exec_env_name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.activateExecEnv") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("activateExecEnv. Error %v", err)
		return err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.ActivateExecEnv(ctx)
	req = req.ActivateExecEnvReq(api.ActivateExecEnvReq{
		ApplianceId:     &appliance_id,
		IdaName:         &appliance_id,
		ExecEnvName:     &exec_env_name,
		ExecEnvFolder:   &exec_env_folder,
		Reactivate:      &force,
		Replace:         &replace_config,
		ActivateSamples: &install_samples,
	})
	res, _, err := c.apiClient.DefaultApi.ActivateExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("activateExecEnv. Error %v", err)
		return err
	}

	if res.Error != nil {
		c.logger.Errorf("activateExecEnv. Error %v", *res.Error)
		return errors.New(*res.Error)
	}

	c.logger.Debugf("activateExecEnv. [%s] Activated %s", appliance_id, exec_env_name)

	return err
}
