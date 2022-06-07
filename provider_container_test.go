package cli

import (
	api "github.com/atricore/josso-api-go"
	"github.com/stretchr/testify/assert"
)

func (s *AccTestSuite) TestAccProviderContainer_read() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	var authn []api.AuthenticationMechanismDTO
	authn = append(authn, createTestBasicAuthn())
	orig, err := createTestIdentityProviderDTO("idp-1", authn)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.client.CreateIdp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}

	// Test READ
	var providers []api.ProviderContainerDTO
	providers, err = s.client.GetProviders(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, len(providers), 1, "Total providers")

	provider, err := s.client.GetProvider(*appliance.Name, "idp-1")
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "IdentityProvider", provider.GetType())
	assert.Equal(t, "idp-1", provider.GetName())
	assert.NotNil(t, provider.FederatedProvider, "provider is nil")

	//fmt.Printf("PROVIDER:\n%#v\n", provider.FederatedProvider.AdditionalProperties)
}
