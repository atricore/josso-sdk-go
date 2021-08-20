package cli

import (
	"strconv"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdentityAppliance_crud() {
	var t = s.T()

	// Test CRUD
	crudName := "ida-z"
	var orig *api.IdentityApplianceDefinitionDTO
	var created api.IdentityApplianceDefinitionDTO
	orig = createTestIdentityApplianceDefinitionDTO(crudName)
	// Create

	created, err := s.client.CreateAppliance(*orig)
	if err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}
	if err := IdApplianceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating identity appliance : %v", err)
		return
	}

	// Retrieve
	retrieved, err := s.client.GetAppliance(crudName)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&created, &retrieved); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Update
	retrieved.Namespace = api.PtrString("com.atricore.ida.a.mod")
	updated, err := s.client.UpdateAppliance(retrieved)
	if err != nil {
		t.Errorf("retrieving identity appilance : %v", err)
		return
	}
	if err := IdApplianceValidateUpdate(&retrieved, &updated); err != nil {
		t.Errorf("retrieving identity appliance : %v", err)
		return
	}

	// Delete
	removed, err := s.client.DeleteAppliance(strconv.FormatInt(*updated.Id, 10))
	if err != nil {
		t.Errorf("deleting identity appilance : %v", err)
		return
	}
	if !removed {
		t.Errorf("deleting identity appliance : not found %d", *updated.Id)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetAppliances()
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name, (these elements must have all the variables of the structure)
	var listOfCreated [2]api.IdentityApplianceDefinitionDTO

	element1 := createTestIdentityApplianceDefinitionDTO("ida-1")
	listOfCreated[0], _ = s.client.CreateAppliance(*element1)

	element2 := createTestIdentityApplianceDefinitionDTO("ida-2")
	listOfCreated[1], _ = s.client.CreateAppliance(*element2)

	// Get list from server
	listOfRead, err := s.client.GetAppliances()
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
}
func createTestIdentityApplianceDefinitionDTO(name string) *api.IdentityApplianceDefinitionDTO {
	l, _ := StrToLocation("http://localhost/IDBUS/IDA-B")
	orig := api.NewIdentityApplianceDefinitionDTO()
	var AuthenticationService []api.AuthenticationServiceDTO
	AuthenticationService1 := api.NewAuthenticationServiceDTO()
	AuthenticationService2 := api.NewAuthenticationServiceDTO()
	var ExecutionEnvironmentDTO []api.ExecutionEnvironmentDTO
	ExecutionEnvironmentDTO1 := api.NewExecutionEnvironmentDTO()
	ExecutionEnvironmentDTO2 := api.NewExecutionEnvironmentDTO()
	var IdentitySourceDTO []api.IdentitySourceDTO
	IdentitySourceDTO1 := api.NewIdentitySourceDTO()
	IdentitySourceDTO2 := api.NewIdentitySourceDTO()
	var EntitySelectionStrategyDTO api.EntitySelectionStrategyDTO
	var keystore api.KeystoreDTO
	var ProviderDTO []api.ProviderDTO
	ProviderDTO1 := api.NewProviderDTO()
	ProviderDTO2 := api.NewProviderDTO()
	var IdentityApplianceSecurityConfigDTO api.IdentityApplianceSecurityConfigDTO
	var ServiceResource []api.ServiceResourceDTO
	ServiceResource1 := api.NewServiceResourceDTO()
	ServiceResource2 := api.NewServiceResourceDTO()
	var UserDashboardBrandingDTO api.UserDashboardBrandingDTO
	var DelegatedAuthenticationDTO []api.DelegatedAuthenticationDTO
	DelegatedAuthenticationDTO1 := api.NewDelegatedAuthenticationDTO()
	DelegatedAuthenticationDTO2 := api.NewDelegatedAuthenticationDTO()
	var locat api.LocationDTO
	var ActivationDTO []api.ActivationDTO
	ActivationDTO1 := api.NewActivationDTO()
	ActivationDTO2 := api.NewActivationDTO()
	var ResourceDTO api.ResourceDTO
	var identityAppliance api.IdentityApplianceDefinitionDTO
	var conf api.ProviderConfigDTO
	var ServiceConnectionDTO api.ServiceConnectionDTO
	var Activation api.ActivationDTO
	var poi []api.PointDTO
	poi1 := api.NewPointDTO()
	poi2 := api.NewPointDTO()
	var IdentityProviderDTO api.IdentityProviderDTO
	var AuthenticationServiceDTO api.AuthenticationServiceDTO
	var ServiceResourceDTO api.ServiceResourceDTO
	var SubjectNameIdentifierPolicy api.SubjectNameIdentifierPolicyDTO
	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO
	var fedconn []api.FederatedConnectionDTO
	fedconn1 := api.NewFederatedConnectionDTO()
	fedconn2 := api.NewFederatedConnectionDTO()
	var OAuth2ClientDTO []api.OAuth2ClientDTO
	OAuth2ClientDTO1 := api.NewOAuth2ClientDTO()
	OAuth2ClientDTO2 := api.NewOAuth2ClientDTO()
	var ExtensionDTO api.ExtensionDTO
	var AuthenticationAssertionEmissionPolicyDTO api.AuthenticationAssertionEmissionPolicyDTO
	var AuthenticationMechanismDTO []api.AuthenticationMechanismDTO
	AuthenticationMechanismDTO1 := api.NewAuthenticationMechanismDTO()
	AuthenticationMechanismDTO2 := api.NewAuthenticationMechanismDTO()
	var AuthenticationContractDTO api.AuthenticationContractDTO
	var AttributeProfileDTO api.AttributeProfileDTO
	var FederatedProviderDTO api.FederatedProviderDTO
	var FederatedChannelDTO api.FederatedChannelDTO
	var DelegatedAuthentication api.DelegatedAuthenticationDTO
	var FederatedConnection api.FederatedConnectionDTO

	FederatedConnection.SetChannelA(FederatedChannelDTO)
	FederatedConnection.SetChannelB(FederatedChannelDTO)
	FederatedConnection.SetDescription("")
	FederatedConnection.SetElementId("")
	FederatedConnection.SetId(1)
	FederatedConnection.SetName("")
	FederatedConnection.SetRoleA(FederatedProviderDTO)
	FederatedConnection.SetRoleB(FederatedProviderDTO)
	FederatedConnection.SetWaypoints(poi)

	DelegatedAuthentication.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthentication.SetDescription("")
	DelegatedAuthentication.SetElementId("")
	DelegatedAuthentication.SetId(1)
	DelegatedAuthentication.SetIdp(IdentityProviderDTO)
	DelegatedAuthentication.SetName("")
	DelegatedAuthentication.SetWaypoints(poi)

	// FederatedChannelDTO.SetActiveBindings("") //////////// preguntar
	// FederatedChannelDTO.SetActiveProfiles("") //////////// preguntar
	FederatedChannelDTO.SetConnectionA(FederatedConnection)
	FederatedChannelDTO.SetConnectionB(FederatedConnection)
	FederatedChannelDTO.SetDescription("")
	FederatedChannelDTO.SetDisplayName("")
	FederatedChannelDTO.SetElementId("")
	FederatedChannelDTO.SetId(1)
	FederatedChannelDTO.SetLocation(locat)
	FederatedChannelDTO.SetName("")
	FederatedChannelDTO.SetOverrideProviderSetup(true)

	// FederatedProviderDTO.SetActiveBindings("") //
	// FederatedProviderDTO.SetActiveProfiles("") //
	FederatedProviderDTO.SetConfig(conf)
	FederatedProviderDTO.SetDescription("")
	FederatedProviderDTO.SetDisplayName("")
	FederatedProviderDTO.SetElementId("")
	FederatedProviderDTO.SetId(1)
	FederatedProviderDTO.SetIdentityAppliance(identityAppliance)
	FederatedProviderDTO.SetIsRemote(true)
	FederatedProviderDTO.SetLocation(locat)
	FederatedProviderDTO.SetMetadata(ResourceDTO)
	FederatedProviderDTO.SetName("")
	FederatedProviderDTO.SetRemote(true)
	FederatedProviderDTO.SetRole("")

	AttributeProfileDTO.SetElementId("")
	AttributeProfileDTO.SetId(1)
	AttributeProfileDTO.SetName("")
	AttributeProfileDTO.SetProfileType("")

	AuthenticationContractDTO.SetElementId("")
	AuthenticationContractDTO.SetId(1)
	AuthenticationContractDTO.SetName("")

	AuthenticationMechanismDTO1.SetDelegatedAuthentication(DelegatedAuthentication)
	AuthenticationMechanismDTO1.SetDisplayName("")
	AuthenticationMechanismDTO1.SetElementId("")
	AuthenticationMechanismDTO1.SetId(1)
	AuthenticationMechanismDTO1.SetName("")
	AuthenticationMechanismDTO1.SetPriority(1)
	AuthenticationMechanismDTO = append(AuthenticationMechanismDTO, *AuthenticationMechanismDTO1)
	AuthenticationMechanismDTO2.SetDelegatedAuthentication(DelegatedAuthentication)
	AuthenticationMechanismDTO2.SetDisplayName("")
	AuthenticationMechanismDTO2.SetElementId("")
	AuthenticationMechanismDTO2.SetId(1)
	AuthenticationMechanismDTO2.SetName("")
	AuthenticationMechanismDTO2.SetPriority(1)
	AuthenticationMechanismDTO = append(AuthenticationMechanismDTO, *AuthenticationMechanismDTO2)

	AuthenticationAssertionEmissionPolicyDTO.SetElementId("")
	AuthenticationAssertionEmissionPolicyDTO.SetId(1)
	AuthenticationAssertionEmissionPolicyDTO.SetName("")

	ExtensionDTO.SetClassifier("")
	ExtensionDTO.SetId("")
	ExtensionDTO.SetName("")
	ExtensionDTO.SetNamespace("")
	ExtensionDTO.SetProvider("")
	ExtensionDTO.SetVersion("")

	OAuth2ClientDTO1.SetBaseURL("")
	OAuth2ClientDTO1.SetId("")
	OAuth2ClientDTO1.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO1)
	OAuth2ClientDTO2.SetBaseURL("")
	OAuth2ClientDTO2.SetId("")
	OAuth2ClientDTO2.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO2)

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

	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")

	SubjectNameIdentifierPolicy.SetDescriptionKey("")
	SubjectNameIdentifierPolicy.SetId("")
	SubjectNameIdentifierPolicy.SetName("")
	SubjectNameIdentifierPolicy.SetSubjectAttribute("")
	SubjectNameIdentifierPolicy.SetType("")

	ServiceResourceDTO.SetActivation(Activation)
	ServiceResourceDTO.SetDescription("")
	ServiceResourceDTO.SetElementId("")
	ServiceResourceDTO.SetId(1)
	ServiceResourceDTO.SetName("")
	ServiceResourceDTO.SetServiceConnection(ServiceConnectionDTO)

	AuthenticationServiceDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationServiceDTO.SetDescription("")
	AuthenticationServiceDTO.SetDisplayName("")
	AuthenticationServiceDTO.SetElementId("")
	AuthenticationServiceDTO.SetId(1)
	AuthenticationServiceDTO.SetName("")

	// IdentityProviderDTO.SetActiveBindings("") //
	// IdentityProviderDTO.SetActiveProfiles("") //
	IdentityProviderDTO.SetAttributeProfile(AttributeProfileDTO)
	IdentityProviderDTO.SetAuthenticationContract(AuthenticationContractDTO)
	IdentityProviderDTO.SetAuthenticationMechanisms(AuthenticationMechanismDTO)
	IdentityProviderDTO.SetConfig(conf)
	IdentityProviderDTO.SetDashboardUrl("")
	IdentityProviderDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
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
	IdentityProviderDTO.SetFederatedConnectionsA(fedconn)
	IdentityProviderDTO.SetFederatedConnectionsB(fedconn)
	IdentityProviderDTO.SetId(1)
	IdentityProviderDTO.SetIdentityAppliance(identityAppliance)
	IdentityProviderDTO.SetIdentityConfirmationEnabled(true)
	IdentityProviderDTO.SetIdentityConfirmationOAuth2AuthorizationServerEndpoint("")
	IdentityProviderDTO.SetIdentityConfirmationOAuth2ClientId("")
	IdentityProviderDTO.SetIdentityConfirmationOAuth2ClientSecret("")
	IdentityProviderDTO.SetIdentityConfirmationPolicy(ExtensionDTO)
	IdentityProviderDTO.SetIgnoreRequestedNameIDPolicy(true)
	IdentityProviderDTO.SetIsRemote(true)
	IdentityProviderDTO.SetLocation(locat)
	IdentityProviderDTO.SetMaxSessionsPerUser(1)
	IdentityProviderDTO.SetMessageTtl(300)
	IdentityProviderDTO.SetMessageTtlTolerance(300)
	IdentityProviderDTO.SetMetadata(ResourceDTO)
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
	IdentityProviderDTO.SetSessionManagerFactory(SessionManagerFactoryDTO)
	IdentityProviderDTO.SetSignRequests(true)
	IdentityProviderDTO.SetSignatureHash("")
	IdentityProviderDTO.SetSsoSessionTimeout(1)
	IdentityProviderDTO.SetSubjectAuthnPolicies(*IdentityProviderDTO.SubjectAuthnPolicies)
	IdentityProviderDTO.SetSubjectNameIDPolicy(SubjectNameIdentifierPolicy)
	IdentityProviderDTO.SetUserDashboardBranding("")
	IdentityProviderDTO.SetWantAuthnRequestsSigned(true)
	IdentityProviderDTO.SetWantSignedRequests(true)
	var ExecutionEnvironment api.ExecutionEnvironmentDTO

	ExecutionEnvironment.SetActivations(ActivationDTO)
	ExecutionEnvironment.SetActive(true)
	ExecutionEnvironment.SetBindingLocation(locat)
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

	poi1.SetX(1)
	poi1.SetY(1)
	poi = append(poi, *poi1)
	poi2.SetX(1)
	poi2.SetY(1)
	poi = append(poi, *poi2)

	Activation.SetDescription("")
	Activation.SetElementId("")
	Activation.SetExecutionEnv(ExecutionEnvironment)
	Activation.SetId(1)
	Activation.SetName("")
	Activation.SetResource(ServiceResourceDTO)
	//Activation.SetSp()
	Activation.SetWaypoints(poi)

	ServiceConnectionDTO.SetDescription("")
	ServiceConnectionDTO.SetElementId("")
	ServiceConnectionDTO.SetId(1)
	ServiceConnectionDTO.SetName("")
	ServiceConnectionDTO.SetResource(ServiceResourceDTO)
	//ServiceConnectionDTO.SetSp()
	ServiceConnectionDTO.SetWaypoints(poi)

	conf.SetDescription("")
	conf.SetDisplayName("")
	conf.SetElementId("")
	conf.SetId(1)
	conf.SetName("")

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

	ResourceDTO.SetDisplayName("")
	ResourceDTO.SetElementId("")
	ResourceDTO.SetId(1)
	ResourceDTO.SetName("")
	ResourceDTO.SetUri("")
	ResourceDTO.SetValue("")

	ActivationDTO1.SetDescription("")
	ActivationDTO1.SetElementId("")
	ActivationDTO1.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO1.SetId(1)
	ActivationDTO1.SetName("")
	ActivationDTO1.SetResource(ServiceResourceDTO)
	//ActivationDTO1.SetSp()
	ActivationDTO1.SetWaypoints(poi)
	ActivationDTO = append(ActivationDTO, *ActivationDTO1)
	ActivationDTO2.SetDescription("")
	ActivationDTO2.SetElementId("")
	ActivationDTO2.SetExecutionEnv(ExecutionEnvironment)
	ActivationDTO2.SetId(1)
	ActivationDTO2.SetName("")
	ActivationDTO2.SetResource(ServiceResourceDTO)
	//ActivationDTO2.SetSp()
	ActivationDTO2.SetWaypoints(poi)
	ActivationDTO = append(ActivationDTO, *ActivationDTO2)

	locat.SetContext("")
	locat.SetElementId("")
	locat.SetHost("")
	locat.SetId(1)
	locat.SetLocationAsString("")
	locat.SetPort(1)
	locat.SetProtocol("")
	locat.SetUri("")

	DelegatedAuthenticationDTO1.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO1.SetDescription("")
	DelegatedAuthenticationDTO1.SetElementId("")
	DelegatedAuthenticationDTO1.SetId(1)
	DelegatedAuthenticationDTO1.SetIdp(IdentityProviderDTO)
	DelegatedAuthenticationDTO1.SetName("")
	DelegatedAuthenticationDTO1.SetWaypoints(poi)
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO1)
	DelegatedAuthenticationDTO2.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthenticationDTO2.SetDescription("")
	DelegatedAuthenticationDTO2.SetElementId("")
	DelegatedAuthenticationDTO2.SetId(1)
	DelegatedAuthenticationDTO2.SetIdp(IdentityProviderDTO)
	DelegatedAuthenticationDTO2.SetName("")
	DelegatedAuthenticationDTO2.SetWaypoints(poi)
	DelegatedAuthenticationDTO = append(DelegatedAuthenticationDTO, *DelegatedAuthenticationDTO2)

	UserDashboardBrandingDTO.SetId("")
	UserDashboardBrandingDTO.SetName("")

	ServiceResource1.SetActivation(Activation)
	ServiceResource1.SetDescription("")
	ServiceResource1.SetElementId("")
	ServiceResource1.SetId(1)
	ServiceResource1.SetName("")
	ServiceResource1.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource1)
	ServiceResource2.SetActivation(Activation)
	ServiceResource2.SetDescription("")
	ServiceResource2.SetElementId("")
	ServiceResource2.SetId(1)
	ServiceResource2.SetName("")
	ServiceResource2.SetServiceConnection(ServiceConnectionDTO)
	ServiceResource = append(ServiceResource, *ServiceResource2)

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

	// ProviderDTO1.SetActiveBindings() //
	// ProviderDTO1.SetActiveProfiles() //
	ProviderDTO1.SetConfig(conf)
	ProviderDTO1.SetDescription("")
	ProviderDTO1.SetDisplayName("")
	ProviderDTO1.SetElementId("")
	ProviderDTO1.SetId(1)
	ProviderDTO1.SetIdentityAppliance(identityAppliance)
	ProviderDTO1.SetIsRemote(true)
	ProviderDTO1.SetLocation(locat)
	ProviderDTO1.SetMetadata(ResourceDTO)
	ProviderDTO1.SetName("")
	ProviderDTO1.SetRemote(true)
	ProviderDTO1.SetRole("")
	ProviderDTO = append(ProviderDTO, *ProviderDTO1)
	// ProviderDTO2.SetActiveBindings() //
	// ProviderDTO2.SetActiveProfiles() //
	ProviderDTO2.SetConfig(conf)
	ProviderDTO2.SetDescription("")
	ProviderDTO2.SetDisplayName("")
	ProviderDTO2.SetElementId("")
	ProviderDTO2.SetId(1)
	ProviderDTO2.SetIdentityAppliance(identityAppliance)
	ProviderDTO2.SetIsRemote(true)
	ProviderDTO2.SetLocation(locat)
	ProviderDTO2.SetMetadata(ResourceDTO)
	ProviderDTO2.SetName("")
	ProviderDTO2.SetRemote(true)
	ProviderDTO2.SetRole("")
	ProviderDTO = append(ProviderDTO, *ProviderDTO2)

	keystore.SetCertificateAlias("")
	keystore.SetDisplayName("")
	keystore.SetElementId("")
	keystore.SetId(1)
	keystore.SetKeystorePassOnly(true)
	keystore.SetName("")
	keystore.SetPassword("")
	keystore.SetPrivateKeyName("")
	keystore.SetPrivateKeyPassword("")
	keystore.SetStore(ResourceDTO)
	keystore.SetType("")

	EntitySelectionStrategyDTO.SetDescription("")
	EntitySelectionStrategyDTO.SetName("")

	IdentitySourceDTO1.SetDescription("")
	IdentitySourceDTO1.SetElementId("")
	IdentitySourceDTO1.SetId(1)
	IdentitySourceDTO1.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO1)
	IdentitySourceDTO2.SetDescription("")
	IdentitySourceDTO2.SetElementId("")
	IdentitySourceDTO2.SetId(1)
	IdentitySourceDTO2.SetName("")
	IdentitySourceDTO = append(IdentitySourceDTO, *IdentitySourceDTO2)

	ExecutionEnvironmentDTO1.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO1.SetActive(true)
	ExecutionEnvironmentDTO1.SetBindingLocation(locat)
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
	ExecutionEnvironmentDTO2.SetActivations(ActivationDTO)
	ExecutionEnvironmentDTO2.SetActive(true)
	ExecutionEnvironmentDTO2.SetBindingLocation(locat)
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

	AuthenticationService1.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService1.SetDescription("")
	AuthenticationService1.SetDisplayName("")
	AuthenticationService1.SetElementId("")
	AuthenticationService1.SetId(1)
	AuthenticationService1.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService1)
	AuthenticationService2.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationService2.SetDescription("")
	AuthenticationService2.SetDisplayName("")
	AuthenticationService2.SetElementId("")
	AuthenticationService2.SetId(1)
	AuthenticationService2.SetName("")
	AuthenticationService = append(AuthenticationService, *AuthenticationService2)

	// orig.SetActiveFeatures() //
	orig.SetAuthenticationServices(AuthenticationService)
	orig.SetDescription("IDA-B TEST !")
	orig.SetDisplayName("")
	orig.SetElementId("")
	orig.SetExecutionEnvironments(ExecutionEnvironmentDTO)
	orig.SetId(1)
	orig.SetIdentitySources(IdentitySourceDTO)
	orig.SetIdpSelector(EntitySelectionStrategyDTO)
	orig.SetKeystore(keystore)
	// orig.SetLastModification() // TIME.TIME ??
	orig.SetLocation(*l)
	orig.SetModelVersion("")
	orig.SetName(name)
	orig.SetNamespace("com.atricore.idbus.ida.b")
	orig.SetProviders(ProviderDTO)
	// orig.SetRequiredBundles() //
	orig.SetRevision(1)
	orig.SetSecurityConfig(IdentityApplianceSecurityConfigDTO)
	orig.SetServiceResources(ServiceResource)
	// orig.SetSupportedRoles() //
	orig.SetUserDashboardBranding(UserDashboardBrandingDTO)

	return orig
}

// -------------------------------------------------

//Fields to validate after appliance creation
func ApplianceFieldTestCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "namespace",
			cmp:      func() bool { return StrPtrEquals(e.Namespace, r.Namespace) },
			expected: StrDeref(e.Namespace),
			received: StrDeref(r.Namespace),
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
	}
}

//Fields to validate after appliance update
func ApplianceFieldTestUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}
	return append(t, ApplianceFieldTestCreate(e, r)...)
}

// Compares the expected appliance with the received one.
func IdApplianceValidateCreate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestCreate(e, r))
}

// Compares the expected appliance with the received one.
func IdApplianceValidateUpdate(
	e *api.IdentityApplianceDefinitionDTO,
	r *api.IdentityApplianceDefinitionDTO) error {

	return ValidateFields(ApplianceFieldTestUpdate(e, r))
}
