package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliVirtSaml2_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Virtsaml2-a"
	var orig *api.VirtualSaml2ServiceProviderDTO
	var created api.VirtualSaml2ServiceProviderDTO
	orig = createTestVirtualSaml2ServiceProviderDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateVirtSaml2Sp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := VirtualSaml2SpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.VirtualSaml2ServiceProviderDTO
	read, err = s.client.GetVirtSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = VirtualSaml2SpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateVirtSaml2Sp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = VirtualSaml2SpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteVirtSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// ------------------------------------------------------------------------------------------------------------------
	// Test empty list

	listOfAll, err := s.client.GetVirtSaml2Sps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// ------------------------
	// List of created elements, order by Name
	var listOfCreated [2]api.VirtualSaml2ServiceProviderDTO

	// Test list of #2 elements
	element1 := createTestVirtualSaml2ServiceProviderDTO("Extsmal2-1")
	listOfCreated[0], _ = s.client.CreateVirtSaml2Sp(*appliance.Name, *element1)

	element2 := createTestVirtualSaml2ServiceProviderDTO("Extsmal2-2")
	listOfCreated[1], _ = s.client.CreateVirtSaml2Sp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetVirtSaml2Sps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	// The list should have 2 elemetns
	if len(listOfRead) != 2 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expected 2", len(listOfAll))
		return
	}

	// Order list of read by Name
	sort.SliceStable(listOfRead,
		func(i, j int) bool {
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) < 0
		},
	)

	// Validate each element from the list of created with the list of read
	for idx, r := range listOfCreated {
		if err = VirtualSaml2SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestVirtualSaml2ServiceProviderDTO(name string) *api.VirtualSaml2ServiceProviderDTO {
	//encMetadata := metadata
	tData := api.NewVirtualSaml2ServiceProviderDTO()

	var DelegatedAuthentication api.DelegatedAuthenticationDTO
	DelegatedAuthentication.SetDescription("")
	DelegatedAuthentication.SetElementId("")
	DelegatedAuthentication.SetId(1)
	DelegatedAuthentication.SetName("")

	var AuthenticationContractDTO api.AuthenticationContractDTO
	AuthenticationContractDTO.SetElementId("")
	AuthenticationContractDTO.SetId(1)
	AuthenticationContractDTO.SetName("")

	var AuthenticationMechanismDTO []api.AuthenticationMechanismDTO
	AuthenticationMechanismDTO1 := api.NewAuthenticationMechanismDTO()
	AuthenticationMechanismDTO1.SetDelegatedAuthentication(DelegatedAuthentication)
	AuthenticationMechanismDTO1.SetDisplayName("")
	AuthenticationMechanismDTO1.SetElementId("")
	AuthenticationMechanismDTO1.SetId(1)
	AuthenticationMechanismDTO1.SetName("")
	AuthenticationMechanismDTO1.SetPriority(1)
	AuthenticationMechanismDTO = append(AuthenticationMechanismDTO, *AuthenticationMechanismDTO1)
	AuthenticationMechanismDTO2 := api.NewAuthenticationMechanismDTO()
	AuthenticationMechanismDTO2.SetDelegatedAuthentication(DelegatedAuthentication)
	AuthenticationMechanismDTO2.SetDisplayName("")
	AuthenticationMechanismDTO2.SetElementId("")
	AuthenticationMechanismDTO2.SetId(1)
	AuthenticationMechanismDTO2.SetName("")
	AuthenticationMechanismDTO2.SetPriority(1)
	AuthenticationMechanismDTO = append(AuthenticationMechanismDTO, *AuthenticationMechanismDTO2)

	var AuthenticationAssertionEmissionPolicyDTO api.AuthenticationAssertionEmissionPolicyDTO
	AuthenticationAssertionEmissionPolicyDTO.SetElementId("")
	AuthenticationAssertionEmissionPolicyDTO.SetId(1)
	AuthenticationAssertionEmissionPolicyDTO.SetName("")

	var ExtensionDTO api.ExtensionDTO
	ExtensionDTO.SetClassifier("")
	ExtensionDTO.SetId("")
	ExtensionDTO.SetName("")
	ExtensionDTO.SetNamespace("")
	ExtensionDTO.SetProvider("")
	ExtensionDTO.SetVersion("")

	var OAuth2ClientDTO []api.OAuth2ClientDTO
	OAuth2ClientDTO1 := api.NewOAuth2ClientDTO()
	OAuth2ClientDTO1.SetBaseURL("")
	OAuth2ClientDTO1.SetId("")
	OAuth2ClientDTO1.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO1)
	OAuth2ClientDTO2 := api.NewOAuth2ClientDTO()
	OAuth2ClientDTO2.SetBaseURL("")
	OAuth2ClientDTO2.SetId("")
	OAuth2ClientDTO2.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO2)

	var SubjectNameIdentifierPolicy api.SubjectNameIdentifierPolicyDTO
	SubjectNameIdentifierPolicy.SetDescriptionKey("")
	SubjectNameIdentifierPolicy.SetId("")
	SubjectNameIdentifierPolicy.SetName("")
	SubjectNameIdentifierPolicy.SetSubjectAttribute("")
	SubjectNameIdentifierPolicy.SetType("")

	var AuthenticationServiceDTO api.AuthenticationServiceDTO
	AuthenticationServiceDTO.SetDescription("")
	AuthenticationServiceDTO.SetDisplayName("")
	AuthenticationServiceDTO.SetElementId("")
	AuthenticationServiceDTO.SetId(1)
	AuthenticationServiceDTO.SetName("")

	var IdentityProviderDTO api.IdentityProviderDTO
	// IdentityProviderDTO.SetActiveBindings("") //
	// IdentityProviderDTO.SetActiveProfiles("") //
	IdentityProviderDTO.SetAuthenticationContract(AuthenticationContractDTO)
	IdentityProviderDTO.SetAuthenticationMechanisms(AuthenticationMechanismDTO)
	IdentityProviderDTO.SetDashboardUrl("")
	IdentityProviderDTO.SetDescription("")
	IdentityProviderDTO.SetDestroyPreviousSession(true)
	IdentityProviderDTO.SetDisplayName("")
	IdentityProviderDTO.SetElementId("")
	IdentityProviderDTO.SetEmissionPolicy(AuthenticationAssertionEmissionPolicyDTO)
	IdentityProviderDTO.SetEnableMetadataEndpoint(true)
	IdentityProviderDTO.SetEncryptAssertion(true)
	IdentityProviderDTO.SetEncryptAssertionAlgorithm("")
	IdentityProviderDTO.SetErrorBinding("")
	IdentityProviderDTO.SetExternallyHostedIdentityConfirmationTokenService(true)
	IdentityProviderDTO.SetId(1)
	IdentityProviderDTO.SetIdentityConfirmationEnabled(true)
	IdentityProviderDTO.SetIdentityConfirmationOAuth2AuthorizationServerEndpoint("")
	IdentityProviderDTO.SetIdentityConfirmationOAuth2ClientId("")
	IdentityProviderDTO.SetIdentityConfirmationOAuth2ClientSecret("")
	IdentityProviderDTO.SetIdentityConfirmationPolicy(ExtensionDTO)
	IdentityProviderDTO.SetIgnoreRequestedNameIDPolicy(true)
	IdentityProviderDTO.SetIsRemote(true)
	IdentityProviderDTO.SetMaxSessionsPerUser(1)
	IdentityProviderDTO.SetMessageTtl(300)
	IdentityProviderDTO.SetMessageTtlTolerance(300)
	IdentityProviderDTO.SetName("")
	IdentityProviderDTO.SetOauth2Clients(OAuth2ClientDTO)
	IdentityProviderDTO.SetOauth2ClientsConfig("")
	IdentityProviderDTO.SetOauth2Enabled(true)
	IdentityProviderDTO.SetOauth2Key("")
	IdentityProviderDTO.SetOauth2RememberMeTokenValidity(1)
	IdentityProviderDTO.SetOauth2TokenValidity(1)
	IdentityProviderDTO.SetOidcAccessTokenTimeToLive(1)
	IdentityProviderDTO.SetOidcAuthzCodeTimeToLive(1)
	IdentityProviderDTO.SetOidcIdTokenTimeToLive(1)
	IdentityProviderDTO.SetOpenIdEnabled(true)
	IdentityProviderDTO.SetPwdlessAuthnEnabled(true)
	IdentityProviderDTO.SetPwdlessAuthnFrom("")
	IdentityProviderDTO.SetPwdlessAuthnSubject("")
	IdentityProviderDTO.SetPwdlessAuthnTemplate("")
	IdentityProviderDTO.SetPwdlessAuthnTo("")
	IdentityProviderDTO.SetRemote(true)
	IdentityProviderDTO.SetRole("")
	IdentityProviderDTO.SetSignRequests(true)
	IdentityProviderDTO.SetSignatureHash("")
	IdentityProviderDTO.SetSsoSessionTimeout(1)
	IdentityProviderDTO.SetSubjectAuthnPolicies(*IdentityProviderDTO.SubjectAuthnPolicies)
	IdentityProviderDTO.SetSubjectNameIDPolicy(SubjectNameIdentifierPolicy)
	IdentityProviderDTO.SetUserDashboardBranding("")
	IdentityProviderDTO.SetWantAuthnRequestsSigned(true)
	IdentityProviderDTO.SetWantSignedRequests(true)
	DelegatedAuthentication.SetIdp(IdentityProviderDTO)

	var ExecutionEnvironment api.ExecutionEnvironmentDTO
	ExecutionEnvironment.SetActive(true)
	ExecutionEnvironment.SetDescription("")
	ExecutionEnvironment.SetDisplayName("")
	ExecutionEnvironment.SetElementId("")
	ExecutionEnvironment.SetId(1)
	ExecutionEnvironment.SetInstallDemoApps(true)
	ExecutionEnvironment.SetInstallUri("")
	ExecutionEnvironment.SetLocation("")
	ExecutionEnvironment.SetName("")
	ExecutionEnvironment.SetOverwriteOriginalSetup(true)
	ExecutionEnvironment.SetPlatformId("")
	ExecutionEnvironment.SetTargetJDK("")
	ExecutionEnvironment.SetType("")

	var ServiceResourceDTO api.ServiceResourceDTO
	ServiceResourceDTO.SetDescription("")
	ServiceResourceDTO.SetElementId("")
	ServiceResourceDTO.SetId(1)
	ServiceResourceDTO.SetName("")

	var DelegatedAuthenticationDTO []api.DelegatedAuthenticationDTO
	DelegatedAuthenticationDTO1 := api.NewDelegatedAuthenticationDTO()
	DelegatedAuthenticationDTO1.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO1.SetDescription("")
	DelegatedAuthenticationDTO1.SetElementId("")
	DelegatedAuthenticationDTO1.SetId(1)
	DelegatedAuthenticationDTO1.SetIdp(IdentityProviderDTO)
	DelegatedAuthenticationDTO1.SetName("")
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO1)
	DelegatedAuthenticationDTO2 := api.NewDelegatedAuthenticationDTO()
	DelegatedAuthenticationDTO2.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO2.SetDescription("")
	DelegatedAuthenticationDTO2.SetElementId("")
	DelegatedAuthenticationDTO2.SetId(1)
	DelegatedAuthenticationDTO2.SetIdp(IdentityProviderDTO)
	DelegatedAuthenticationDTO2.SetName("")
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO2)
	AuthenticationServiceDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	IdentityProviderDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)

	var Activation api.ActivationDTO
	Activation.SetDescription("")
	Activation.SetElementId("")
	Activation.SetExecutionEnv(ExecutionEnvironment)
	Activation.SetId(1)
	Activation.SetName("")
	Activation.SetResource(ServiceResourceDTO)
	//Activation.SetSp()
	ServiceResourceDTO.SetActivation(Activation)

	var ActivationDTO []api.ActivationDTO
	ActivationDTO1 := api.NewActivationDTO()
	ActivationDTO1.SetDescription("")
	ActivationDTO1.SetElementId("")
	ActivationDTO1.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO1.SetId(1)
	ActivationDTO1.SetName("")
	ActivationDTO1.SetResource(ServiceResourceDTO)
	//ActivationDTO1.SetSp()
	ActivationDTO = append(ActivationDTO, *ActivationDTO1)
	ActivationDTO2 := api.NewActivationDTO()
	ActivationDTO2.SetDescription("")
	ActivationDTO2.SetElementId("")
	ActivationDTO2.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO2.SetId(1)
	ActivationDTO2.SetName("")
	ActivationDTO2.SetResource(ServiceResourceDTO)
	//ActivationDTO2.SetSp()
	ActivationDTO = append(ActivationDTO, *ActivationDTO2)
	ExecutionEnvironment.SetActivations(ActivationDTO)

	var ServiceConnectionDTO api.ServiceConnectionDTO
	ServiceConnectionDTO.SetDescription("")
	ServiceConnectionDTO.SetElementId("")
	ServiceConnectionDTO.SetId(1)
	ServiceConnectionDTO.SetName("")
	ServiceConnectionDTO.SetResource(ServiceResourceDTO)
	//ServiceConnectionDTO.SetSp()
	ServiceResourceDTO.SetServiceConnection(ServiceConnectionDTO)

	var FederatedConnection api.FederatedConnectionDTO
	FederatedConnection.SetDescription("")
	FederatedConnection.SetElementId("")
	FederatedConnection.SetId(1)
	FederatedConnection.SetName("")

	var IdentitySource api.IdentitySourceDTO
	IdentitySource.SetDescription("")
	IdentitySource.SetElementId("")
	IdentitySource.SetId(1)
	IdentitySource.SetName("")

	var Provider api.ProviderDTO
	// Provider.SetActiveBindings() //
	// Provider.SetActiveProfiles() //
	Provider.SetDescription("")
	Provider.SetDisplayName("")
	Provider.SetElementId("")
	Provider.SetId(1)
	Provider.SetIsRemote(true)
	Provider.SetName("")
	Provider.SetRemote(true)
	Provider.SetRole("")

	var AuthenticationService []api.AuthenticationServiceDTO
	AuthenticationService1 := api.NewAuthenticationServiceDTO()
	AuthenticationService1.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService1.SetDescription("")
	AuthenticationService1.SetDisplayName("")
	AuthenticationService1.SetElementId("")
	AuthenticationService1.SetId(1)
	AuthenticationService1.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService1)
	AuthenticationService2 := api.NewAuthenticationServiceDTO()
	AuthenticationService2.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService2.SetDescription("")
	AuthenticationService2.SetDisplayName("")
	AuthenticationService2.SetElementId("")
	AuthenticationService2.SetId(1)
	AuthenticationService2.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService2)
	DelegatedAuthentication.SetAuthnService(AuthenticationServiceDTO)

	var ExecutionEnvironmentDTO []api.ExecutionEnvironmentDTO
	ExecutionEnvironmentDTO1 := api.NewExecutionEnvironmentDTO()
	ExecutionEnvironmentDTO1.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO1.SetActive(true)
	ExecutionEnvironmentDTO1.SetDescription("")
	ExecutionEnvironmentDTO1.SetDisplayName("")
	ExecutionEnvironmentDTO1.SetElementId("")
	ExecutionEnvironmentDTO1.SetId(1)
	ExecutionEnvironmentDTO1.SetInstallDemoApps(true)
	ExecutionEnvironmentDTO1.SetInstallUri("")
	ExecutionEnvironmentDTO1.SetLocation("")
	ExecutionEnvironmentDTO1.SetName("")
	ExecutionEnvironmentDTO1.SetOverwriteOriginalSetup(true)
	ExecutionEnvironmentDTO1.SetPlatformId("")
	ExecutionEnvironmentDTO1.SetTargetJDK("")
	ExecutionEnvironmentDTO1.SetType("")
	ExecutionEnvironmentDTO = append(ExecutionEnvironmentDTO, *ExecutionEnvironmentDTO1)
	ExecutionEnvironmentDTO2 := api.NewExecutionEnvironmentDTO()
	ExecutionEnvironmentDTO2.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO2.SetActive(true)
	ExecutionEnvironmentDTO2.SetDescription("")
	ExecutionEnvironmentDTO2.SetDisplayName("")
	ExecutionEnvironmentDTO2.SetElementId("")
	ExecutionEnvironmentDTO2.SetId(1)
	ExecutionEnvironmentDTO2.SetInstallDemoApps(true)
	ExecutionEnvironmentDTO2.SetInstallUri("")
	ExecutionEnvironmentDTO2.SetLocation("")
	ExecutionEnvironmentDTO2.SetName("")
	ExecutionEnvironmentDTO2.SetOverwriteOriginalSetup(true)
	ExecutionEnvironmentDTO2.SetPlatformId("")
	ExecutionEnvironmentDTO2.SetTargetJDK("")
	ExecutionEnvironmentDTO2.SetType("")
	ExecutionEnvironmentDTO = append(ExecutionEnvironmentDTO, *ExecutionEnvironmentDTO2)

	var IdentitySourceDTO []api.IdentitySourceDTO
	IdentitySourceDTO1 := api.NewIdentitySourceDTO()
	IdentitySourceDTO1.SetDescription("")
	IdentitySourceDTO1.SetElementId("")
	IdentitySourceDTO1.SetId(1)
	IdentitySourceDTO1.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO1)
	IdentitySourceDTO2 := api.NewIdentitySourceDTO()
	IdentitySourceDTO2.SetDescription("")
	IdentitySourceDTO2.SetElementId("")
	IdentitySourceDTO2.SetId(1)
	IdentitySourceDTO2.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO2)

	var ProviderDTO []api.ProviderDTO
	ProviderDTO1 := api.NewProviderDTO()
	// ProviderDTO1.SetActiveBindings() //
	// ProviderDTO1.SetActiveProfiles() //
	ProviderDTO1.SetDescription("")
	ProviderDTO1.SetDisplayName("")
	ProviderDTO1.SetElementId("")
	ProviderDTO1.SetId(1)
	ProviderDTO1.SetIsRemote(true)
	ProviderDTO1.SetName("")
	ProviderDTO1.SetRemote(true)
	ProviderDTO1.SetRole("")
	ProviderDTO = append(ProviderDTO, *ProviderDTO1)
	ProviderDTO2 := api.NewProviderDTO()
	// ProviderDTO2.SetActiveBindings() //
	// ProviderDTO2.SetActiveProfiles() //
	ProviderDTO2.SetDescription("")
	ProviderDTO2.SetDisplayName("")
	ProviderDTO2.SetElementId("")
	ProviderDTO2.SetId(1)
	ProviderDTO2.SetIsRemote(true)
	ProviderDTO2.SetName("")
	ProviderDTO2.SetRemote(true)
	ProviderDTO2.SetRole("")
	ProviderDTO = append(ProviderDTO, *ProviderDTO2)

	var EntitySelectionStrategyDTO api.EntitySelectionStrategyDTO
	EntitySelectionStrategyDTO.SetDescription("")
	EntitySelectionStrategyDTO.SetName("")

	var keystore api.KeystoreDTO
	keystore.SetCertificateAlias("")
	keystore.SetDisplayName("")
	keystore.SetElementId("")
	keystore.SetId(1)
	keystore.SetKeystorePassOnly(true)
	keystore.SetName("")
	keystore.SetPassword("")
	keystore.SetPrivateKeyName("")
	keystore.SetPrivateKeyPassword("")
	keystore.SetType("")

	var IdentityApplianceSecurityConfigDTO api.IdentityApplianceSecurityConfigDTO
	IdentityApplianceSecurityConfigDTO.SetEncryptSensitiveData(true)
	IdentityApplianceSecurityConfigDTO.SetEncryption("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionConfig("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionConfigFile("")
	IdentityApplianceSecurityConfigDTO.SetEncryptionPassword("")
	IdentityApplianceSecurityConfigDTO.SetExternalConfig(true)
	IdentityApplianceSecurityConfigDTO.SetExternalConfigFile("")
	IdentityApplianceSecurityConfigDTO.SetPasswordProperty("")
	IdentityApplianceSecurityConfigDTO.SetSalt("")
	IdentityApplianceSecurityConfigDTO.SetSaltProperty("")
	IdentityApplianceSecurityConfigDTO.SetSaltValue("")

	var ServiceResource []api.ServiceResourceDTO
	ServiceResource1 := api.NewServiceResourceDTO()
	ServiceResource1.SetActivation(Activation)
	ServiceResource1.SetDescription("")
	ServiceResource1.SetElementId("")
	ServiceResource1.SetId(1)
	ServiceResource1.SetName("")
	ServiceResource1.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource1)
	ServiceResource2 := api.NewServiceResourceDTO()
	ServiceResource2.SetActivation(Activation)
	ServiceResource2.SetDescription("")
	ServiceResource2.SetElementId("")
	ServiceResource2.SetId(1)
	ServiceResource2.SetName("")
	ServiceResource2.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource2)

	var UserDashboardBrandingDTO api.UserDashboardBrandingDTO
	UserDashboardBrandingDTO.SetId("")
	UserDashboardBrandingDTO.SetName("")

	var FederatedChannelDTO api.FederatedChannelDTO
	// FederatedChannelDTO.SetActiveBindings("") //////////// preguntar
	// FederatedChannelDTO.SetActiveProfiles("") //////////// preguntar
	FederatedChannelDTO.SetConnectionA(FederatedConnection)
	FederatedChannelDTO.SetConnectionB(FederatedConnection)
	FederatedChannelDTO.SetDescription("")
	FederatedChannelDTO.SetDisplayName("")
	FederatedChannelDTO.SetElementId("")
	FederatedChannelDTO.SetId(1)
	FederatedChannelDTO.SetName("")
	FederatedChannelDTO.SetOverrideProviderSetup(true)
	FederatedConnection.SetChannelA(FederatedChannelDTO)
	FederatedConnection.SetChannelB(FederatedChannelDTO)

	var FederatedProviderDTO api.FederatedProviderDTO
	// FederatedProviderDTO.SetActiveBindings("") //
	// FederatedProviderDTO.SetActiveProfiles("") //
	FederatedProviderDTO.SetDescription("")
	FederatedProviderDTO.SetDisplayName("")
	FederatedProviderDTO.SetElementId("")
	FederatedProviderDTO.SetId(1)
	FederatedProviderDTO.SetIsRemote(true)
	FederatedProviderDTO.SetName("")
	FederatedProviderDTO.SetRemote(true)
	FederatedProviderDTO.SetRole("")
	FederatedConnection.SetRoleA(FederatedProviderDTO)
	FederatedConnection.SetRoleB(FederatedProviderDTO)

	var poi []api.PointDTO
	poi1 := api.NewPointDTO()
	poi1.SetX(1)
	poi1.SetY(1)
	poi = append(poi, *poi1)
	poi2 := api.NewPointDTO()
	poi2.SetX(1)
	poi2.SetY(1)
	poi = append(poi, *poi1)
	Activation.SetWaypoints(poi)
	DelegatedAuthentication.SetWaypoints(poi)
	DelegatedAuthenticationDTO1.SetWaypoints(poi)
	ActivationDTO1.SetWaypoints(poi)
	ActivationDTO2.SetWaypoints(poi)
	FederatedConnection.SetWaypoints(poi)
	ServiceConnectionDTO.SetWaypoints(poi)
	DelegatedAuthenticationDTO2.SetWaypoints(poi)

	var SubjectAuthenticationPolicyDTO []api.SubjectAuthenticationPolicyDTO
	SubjectAuthenticationPolicyDTO1 := api.NewSubjectAuthenticationPolicyDTO()
	SubjectAuthenticationPolicyDTO1.SetDescription("")
	SubjectAuthenticationPolicyDTO1.SetName("")
	SubjectAuthenticationPolicyDTO = append(SubjectAuthenticationPolicyDTO, *SubjectAuthenticationPolicyDTO1)
	SubjectAuthenticationPolicyDTO2 := api.NewSubjectAuthenticationPolicyDTO()
	SubjectAuthenticationPolicyDTO2.SetDescription("")
	SubjectAuthenticationPolicyDTO2.SetName("")
	SubjectAuthenticationPolicyDTO = append(SubjectAuthenticationPolicyDTO, *SubjectAuthenticationPolicyDTO2)
	tData.SetSubjectAuthnPolicies(SubjectAuthenticationPolicyDTO)

	var SubjectNameIdentifierPolicyDTO api.SubjectNameIdentifierPolicyDTO
	SubjectNameIdentifierPolicyDTO.SetDescriptionKey("")
	SubjectNameIdentifierPolicyDTO.SetId("")
	SubjectNameIdentifierPolicyDTO.SetName("")
	SubjectNameIdentifierPolicyDTO.SetSubjectAttribute("")
	SubjectNameIdentifierPolicyDTO.SetType("")
	tData.SetSubjectNameIDPolicy(SubjectNameIdentifierPolicyDTO)

	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO
	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")
	tData.SetSessionManagerFactory(SessionManagerFactoryDTO)
	IdentityProviderDTO.SetSessionManagerFactory(SessionManagerFactoryDTO)

	var ResourceDTO api.ResourceDTO
	ResourceDTO.SetDisplayName("")
	ResourceDTO.SetElementId("")
	ResourceDTO.SetId(1)
	ResourceDTO.SetName("")
	ResourceDTO.SetUri("")
	ResourceDTO.SetValue("")
	IdentityProviderDTO.SetMetadata(ResourceDTO)
	Provider.SetMetadata(ResourceDTO)
	ProviderDTO1.SetMetadata(ResourceDTO)
	ProviderDTO2.SetMetadata(ResourceDTO)
	keystore.SetStore(ResourceDTO)
	FederatedProviderDTO.SetMetadata(ResourceDTO)
	tData.SetMetadata(ResourceDTO)

	var locat api.LocationDTO
	locat.SetContext("")
	locat.SetElementId("")
	locat.SetHost("")
	locat.SetId(1)
	locat.SetLocationAsString("")
	locat.SetPort(1)
	locat.SetProtocol("")
	locat.SetUri("")
	IdentityProviderDTO.SetLocation(locat)
	ExecutionEnvironment.SetBindingLocation(locat)
	Provider.SetLocation(locat)
	ExecutionEnvironmentDTO1.SetBindingLocation(locat)
	ExecutionEnvironmentDTO2.SetBindingLocation(locat)
	ProviderDTO1.SetLocation(locat)
	ProviderDTO2.SetLocation(locat)
	FederatedChannelDTO.SetLocation(locat)
	FederatedProviderDTO.SetLocation(locat)
	tData.SetLocation(locat)

	var IdentityMappingPolicyDTO api.IdentityMappingPolicyDTO
	IdentityMappingPolicyDTO.SetCustomMapper("")
	IdentityMappingPolicyDTO.SetElementId("")
	IdentityMappingPolicyDTO.SetId(1)
	IdentityMappingPolicyDTO.SetMappingType("")
	IdentityMappingPolicyDTO.SetName("")
	IdentityMappingPolicyDTO.SetUseLocalId(true)
	tData.SetIdentityMappingPolicy(IdentityMappingPolicyDTO)

	var IdentityLookupDTO []api.IdentityLookupDTO
	IdentityLookupDTO1 := api.NewIdentityLookupDTO()
	IdentityLookupDTO1.SetDescription("")
	IdentityLookupDTO1.SetElementId("")
	IdentityLookupDTO1.SetId(1)
	IdentityLookupDTO1.SetIdentitySource(IdentitySource)
	IdentityLookupDTO1.SetName("")
	IdentityLookupDTO1.SetProvider(Provider)
	IdentityLookupDTO1.SetWaypoints(poi)
	IdentityLookupDTO = append(IdentityLookupDTO, *IdentityLookupDTO1)
	IdentityLookupDTO2 := api.NewIdentityLookupDTO()
	IdentityLookupDTO2.SetDescription("")
	IdentityLookupDTO2.SetElementId("")
	IdentityLookupDTO2.SetId(2)
	IdentityLookupDTO2.SetIdentitySource(IdentitySource)
	IdentityLookupDTO2.SetName("")
	IdentityLookupDTO2.SetProvider(Provider)
	IdentityLookupDTO2.SetWaypoints(poi)
	IdentityLookupDTO = append(IdentityLookupDTO, *IdentityLookupDTO2)
	tData.SetIdentityLookups(IdentityLookupDTO)

	var identityAppliance api.IdentityApplianceDefinitionDTO
	// identityAppliance.SetActiveFeatures("") //
	identityAppliance.SetAuthenticationServices(AuthenticationService)
	identityAppliance.SetDescription("")
	identityAppliance.SetDisplayName("")
	identityAppliance.SetElementId("")
	identityAppliance.SetExecutionEnvironments(ExecutionEnvironmentDTO)
	identityAppliance.SetId(1)
	identityAppliance.SetIdentitySources(IdentitySourceDTO)
	identityAppliance.SetIdpSelector(EntitySelectionStrategyDTO)
	identityAppliance.SetKeystore(keystore)
	// identityAppliance.SetLastModification() //////////*TIME.TIME?
	identityAppliance.SetLocation(locat)
	identityAppliance.SetModelVersion("")
	identityAppliance.SetName("")
	identityAppliance.SetNamespace("")
	identityAppliance.SetProviders(ProviderDTO)
	// identityAppliance.SetRequiredBundles() //
	identityAppliance.SetRevision(1)
	identityAppliance.SetSecurityConfig(IdentityApplianceSecurityConfigDTO)
	identityAppliance.SetServiceResources(ServiceResource)
	// identityAppliance.SetSupportedRoles() //
	identityAppliance.SetUserDashboardBranding(UserDashboardBrandingDTO)
	IdentityProviderDTO.SetIdentityAppliance(identityAppliance)
	Provider.SetIdentityAppliance(identityAppliance)
	ProviderDTO1.SetIdentityAppliance(identityAppliance)
	ProviderDTO2.SetIdentityAppliance(identityAppliance)
	FederatedProviderDTO.SetIdentityAppliance(identityAppliance)
	tData.SetIdentityAppliance(identityAppliance)

	var fedconn []api.FederatedConnectionDTO
	fedconn1 := api.NewFederatedConnectionDTO()
	fedconn1.SetChannelA(FederatedChannelDTO)
	fedconn1.SetChannelB(FederatedChannelDTO)
	fedconn1.SetDescription("")
	fedconn1.SetElementId("")
	fedconn1.SetId(1)
	fedconn1.SetName("")
	fedconn1.SetRoleA(FederatedProviderDTO)
	fedconn1.SetRoleB(FederatedProviderDTO)
	fedconn1.SetWaypoints(poi)
	fedconn = append(fedconn, *fedconn1)
	fedconn2 := api.NewFederatedConnectionDTO()
	fedconn2.SetChannelA(FederatedChannelDTO)
	fedconn2.SetChannelB(FederatedChannelDTO)
	fedconn2.SetDescription("")
	fedconn2.SetElementId("")
	fedconn2.SetId(1)
	fedconn2.SetName("")
	fedconn2.SetRoleA(FederatedProviderDTO)
	fedconn2.SetRoleB(FederatedProviderDTO)
	fedconn2.SetWaypoints(poi)
	fedconn = append(fedconn, *fedconn2)
	tData.SetFederatedConnectionsA(fedconn)
	tData.SetFederatedConnectionsB(fedconn)
	IdentityProviderDTO.SetFederatedConnectionsA(fedconn)
	IdentityProviderDTO.SetFederatedConnectionsB(fedconn)

	var AttributeProfileDTO api.AttributeProfileDTO
	AttributeProfileDTO.SetElementId("")
	AttributeProfileDTO.SetId(1)
	AttributeProfileDTO.SetName("")
	AttributeProfileDTO.SetProfileType("")
	tData.SetAttributeProfile(AttributeProfileDTO)
	IdentityProviderDTO.SetAttributeProfile(AttributeProfileDTO)

	var conf api.ProviderConfigDTO
	conf.SetDescription("")
	conf.SetDisplayName("")
	conf.SetElementId("")
	conf.SetId(1)
	conf.SetName("")
	IdentityProviderDTO.SetConfig(conf)
	Provider.SetConfig(conf)
	ProviderDTO2.SetConfig(conf)
	ProviderDTO1.SetConfig(conf)
	FederatedProviderDTO.SetConfig(conf)
	tData.SetConfig(conf)

	var AccountLinkagePolicyDTO api.AccountLinkagePolicyDTO
	AccountLinkagePolicyDTO.SetCustomLinkEmitter("")
	AccountLinkagePolicyDTO.SetElementId("")
	AccountLinkagePolicyDTO.SetId(1)
	AccountLinkagePolicyDTO.SetLinkEmitterType("")
	AccountLinkagePolicyDTO.SetName("")
	tData.SetAccountLinkagePolicy(AccountLinkagePolicyDTO)

	tData.SetDashboardUrl("")
	tData.SetDescription("")
	tData.SetDisplayName("")
	tData.SetElementId("")
	tData.SetEnableMetadataEndpoint(true)
	tData.SetEnableProxyExtension(true)
	tData.SetEncryptAssertion(true)
	tData.SetEncryptAssertionAlgorithm("")
	tData.SetErrorBinding("")
	tData.SetId(1)
	tData.SetIdpSignatureHash("")
	tData.SetIgnoreRequestedNameIDPolicy(true)
	tData.SetIsRemote(true)
	tData.SetMessageTtl(300)
	tData.SetMessageTtlTolerance(300)
	tData.SetName("")
	tData.SetOauth2Enabled(true)
	tData.SetOauth2Key("")
	tData.SetOauth2RememberMeTokenValidity(1)
	tData.SetOauth2TokenValidity(1)
	tData.SetOidcAccessTokenTimeToLive(1)
	tData.SetOidcAuthzCodeTimeToLive(1)
	tData.SetOidcIdTokenTimeToLive(1)
	tData.SetOpenIdEnabled(true)
	tData.SetRemote(true)
	tData.SetRole("")
	tData.SetSignAuthenticationRequests(true)
	tData.SetSignRequests(true)
	tData.SetSpSignatureHash("")
	tData.SetSsoSessionTimeout(1)
	tData.SetWantAssertionSigned(true)
	tData.SetWantAuthnRequestsSigned(true)
	tData.SetWantSLOResponseSigned(true)
	tData.SetWantSignedRequests(true)

	tData.SetName(name)
	tData.SetId(-1)
	tData.SetDescription("My Sp 2")

	return tData
}

func (s *AccTestSuite) TestAccCliVirSaml2_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliVirSaml2_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func VirtualSaml2SpFieldTestCreate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
	}
}

//Fields to validate after VirtualSaml2Sp update
func VirtualSaml2SpFieldTestUpdate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}
	return append(t, VirtualSaml2SpFieldTestCreate(e, r)...)
}

// Compares the expected VirtualSaml2Sp with the received one.
func VirtualSaml2SpValidateCreate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) error {

	return ValidateFields(VirtualSaml2SpFieldTestCreate(e, r))
}

// Compares the expected ExternalSaml2Sp with the received one.
func VirtualSaml2SpValidateUpdate(
	e *api.VirtualSaml2ServiceProviderDTO,
	r *api.VirtualSaml2ServiceProviderDTO) error {

	return ValidateFields(VirtualSaml2SpFieldTestUpdate(e, r))
}
