package cli

import (
	"encoding/base64"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIntSaml2_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Intsmal2-a"
	var orig *api.InternalSaml2ServiceProviderDTO
	var created api.InternalSaml2ServiceProviderDTO
	orig = createTestInternalSaml2ServiceProviderDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIntSaml2Sp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := InteralSaml2SpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.InternalSaml2ServiceProviderDTO
	read, err = s.client.GetIntSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = InternalSaml2SpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateIntSaml2Sp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = InternalSaml2SpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteIntSaml2Sp(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetIntSaml2Sps(*appliance.Name)
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
	var listOfCreated [2]api.InternalSaml2ServiceProviderDTO

	// Test list of #2 elements
	element1 := createTestInternalSaml2ServiceProviderDTO("Intsmal2-1")
	listOfCreated[0], _ = s.client.CreateIntSaml2Sp(*appliance.Name, *element1)

	element2 := createTestInternalSaml2ServiceProviderDTO("Intsmal2-2")
	listOfCreated[1], _ = s.client.CreateIntSaml2Sp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetIntSaml2Sps(*appliance.Name)
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
		if err = InternalSaml2SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestInternalSaml2ServiceProviderDTO(name string) *api.InternalSaml2ServiceProviderDTO {
	encMetadata := base64.StdEncoding.EncodeToString([]byte(metadata))
	//encMetadata := metadata
	orig := api.NewInternalSaml2ServiceProviderDTO()
	acts1 := api.NewActivationDTO()
	var alp api.AccountLinkagePolicyDTO
	var conf api.ProviderConfigDTO
	var imp api.IdentityMappingPolicyDTO
	var scdto api.ServiceConnectionDTO
	var serc api.ServiceConnectionDTO
	var res api.ServiceResourceDTO
	var act api.ActivationDTO
	var exe api.ExecutionEnvironmentDTO
	var loca api.LocationDTO
	var poi []api.PointDTO
	var acts []api.ActivationDTO
	var AuthenticationContractDTO api.AuthenticationContractDTO
	var AuthenticationMechanism api.AuthenticationMechanismDTO
	var fedconn []api.FederatedConnectionDTO
	fedconn1 := api.NewFederatedConnectionDTO()
	fedconn2 := api.NewFederatedConnectionDTO()
	var identityAppliance api.IdentityApplianceDefinitionDTO
	var locat api.LocationDTO
	var DelegatedAuthentication api.DelegatedAuthenticationDTO
	var FederatedProviderDTO api.FederatedProviderDTO
	var FederatedChannelDTO api.FederatedChannelDTO
	var UserDashboardBrandingDTO api.UserDashboardBrandingDTO
	var ServiceResource []api.ServiceResourceDTO
	ServiceResource1 := api.NewServiceResourceDTO()
	ServiceResource2 := api.NewServiceResourceDTO()
	var IdentityApplianceSecurityConfigDTO api.IdentityApplianceSecurityConfigDTO
	var ProviderDTO []api.ProviderDTO
	ProviderDTO1 := api.NewProviderDTO()
	ProviderDTO2 := api.NewProviderDTO()
	var keystore api.KeystoreDTO
	var EntitySelectionStrategyDTO api.EntitySelectionStrategyDTO
	var IdentitySourceDTO []api.IdentitySourceDTO
	IdentitySourceDTO1 := api.NewIdentitySourceDTO()
	IdentitySourceDTO2 := api.NewIdentitySourceDTO()
	var ExecutionEnvironmentDTO []api.ExecutionEnvironmentDTO
	ExecutionEnvironmentDTO1 := api.NewExecutionEnvironmentDTO()
	ExecutionEnvironmentDTO2 := api.NewExecutionEnvironmentDTO()
	var AuthenticationService []api.AuthenticationServiceDTO
	AuthenticationService1 := api.NewAuthenticationServiceDTO()
	AuthenticationService2 := api.NewAuthenticationServiceDTO()
	var IdentityProviderDTO api.IdentityProviderDTO
	var AuthenticationServiceDTO api.AuthenticationServiceDTO
	var FederatedConnection api.FederatedConnectionDTO
	var ResourceDTO api.ResourceDTO
	var DelegatedAuthenticationDTO []api.DelegatedAuthenticationDTO
	DelegatedAuthenticationDTO1 := api.NewDelegatedAuthenticationDTO()
	DelegatedAuthenticationDTO2 := api.NewDelegatedAuthenticationDTO()
	var SubjectNameIdentifierPolicy api.SubjectNameIdentifierPolicyDTO
	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO
	var OAuth2ClientDTO []api.OAuth2ClientDTO
	OAuth2ClientDTO1 := api.NewOAuth2ClientDTO()
	OAuth2ClientDTO2 := api.NewOAuth2ClientDTO()
	var ExtensionDTO api.ExtensionDTO
	var AuthenticationAssertionEmissionPolicyDTO api.AuthenticationAssertionEmissionPolicyDTO
	var AttributeProfileDTO api.AttributeProfileDTO
	var AuthenticationMechanismDTO []api.AuthenticationMechanismDTO
	AuthenticationMechanismDTO1 := api.NewAuthenticationMechanismDTO()
	AuthenticationMechanismDTO2 := api.NewAuthenticationMechanismDTO()

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

	AttributeProfileDTO.SetElementId("")
	AttributeProfileDTO.SetId(1)
	AttributeProfileDTO.SetName("")
	AttributeProfileDTO.SetProfileType("")
	AuthenticationAssertionEmissionPolicyDTO.SetElementId("")
	AuthenticationAssertionEmissionPolicyDTO.SetId(1)
	AuthenticationAssertionEmissionPolicyDTO.SetName("")

	OAuth2ClientDTO1.SetBaseURL("")
	OAuth2ClientDTO1.SetId("")
	OAuth2ClientDTO1.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO1)
	OAuth2ClientDTO2.SetBaseURL("")
	OAuth2ClientDTO2.SetId("")
	OAuth2ClientDTO2.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO2)

	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")

	SubjectNameIdentifierPolicy.SetDescriptionKey("")
	SubjectNameIdentifierPolicy.SetId("")
	SubjectNameIdentifierPolicy.SetName("")
	SubjectNameIdentifierPolicy.SetSubjectAttribute("")
	SubjectNameIdentifierPolicy.SetType("")

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

	ResourceDTO.SetDisplayName("")
	ResourceDTO.SetElementId("")
	ResourceDTO.SetId(1)
	ResourceDTO.SetName("")
	ResourceDTO.SetUri("")
	ResourceDTO.SetValue("")

	FederatedConnection.SetChannelA(FederatedChannelDTO)
	FederatedConnection.SetChannelB(FederatedChannelDTO)
	FederatedConnection.SetDescription("")
	FederatedConnection.SetElementId("")
	FederatedConnection.SetId(1)
	FederatedConnection.SetName("")
	FederatedConnection.SetRoleA(FederatedProviderDTO)
	FederatedConnection.SetRoleB(FederatedProviderDTO)
	FederatedConnection.SetWaypoints(poi)

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

	ExecutionEnvironmentDTO1.SetActivations(acts)
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
	ExecutionEnvironmentDTO2.SetActivations(acts)
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

	EntitySelectionStrategyDTO.SetDescription("")
	EntitySelectionStrategyDTO.SetName("")

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

	ServiceResource1.SetActivation(act)
	ServiceResource1.SetDescription("")
	ServiceResource1.SetElementId("")
	ServiceResource1.SetId(1)
	ServiceResource1.SetName("")
	ServiceResource1.SetServiceConnection(scdto)
	ServiceResource = append(ServiceResource, *ServiceResource1)
	ServiceResource2.SetActivation(act)
	ServiceResource2.SetDescription("")
	ServiceResource2.SetElementId("")
	ServiceResource2.SetId(1)
	ServiceResource2.SetName("")
	ServiceResource2.SetServiceConnection(scdto)
	ServiceResource = append(ServiceResource, *ServiceResource2)

	UserDashboardBrandingDTO.SetId("")
	UserDashboardBrandingDTO.SetName("")

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

	DelegatedAuthentication.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthentication.SetDescription("")
	DelegatedAuthentication.SetElementId("")
	DelegatedAuthentication.SetId(1)
	DelegatedAuthentication.SetIdp(IdentityProviderDTO)
	DelegatedAuthentication.SetName("")
	DelegatedAuthentication.SetWaypoints(poi)

	locat.SetContext("")
	locat.SetElementId("")
	locat.SetHost("")
	locat.SetId(1)
	locat.SetLocationAsString("")
	locat.SetPort(1)
	locat.SetProtocol("")
	locat.SetUri("")

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

	AuthenticationMechanism.SetDelegatedAuthentication(DelegatedAuthentication)
	AuthenticationMechanism.SetDisplayName("")
	AuthenticationMechanism.SetElementId("")
	AuthenticationMechanism.SetId(1)
	AuthenticationMechanism.SetName("")
	AuthenticationMechanism.SetPriority(1)

	AuthenticationContractDTO.SetElementId("")
	AuthenticationContractDTO.SetId(1)
	AuthenticationContractDTO.SetName("")

	alp.SetCustomLinkEmitter("http://host1:80/")
	alp.SetElementId("my-secret1")
	alp.SetId(-1)
	alp.SetLinkEmitterType("")
	alp.SetName("Account")

	conf.SetDescription("") //JSON SAY ITS NULL
	conf.SetDescription("") //JSON SAY ITS NULL
	conf.SetElementId("")   //JSON SAY ITS NULL
	conf.SetId(49)
	conf.SetName("") // JSON SAY ITS NULL

	imp.SetCustomMapper("")
	imp.SetElementId("")
	imp.SetId(1) // NO LO ENCONTRE EN JSON
	imp.SetMappingType("")
	imp.SetName("")
	imp.SetUseLocalId(true)

	scdto.SetDescription("") //JSON SAY ITS NULL
	scdto.SetElementId("_CDD050B9-E247-4C28-BFC7-E5CDD6AE969B")
	scdto.SetId(50)
	scdto.SetName("sp-2-to-app-2-svc")
	scdto.SetResource(res)
	scdto.SetSp(*orig)

	serc.SetDescription("")
	serc.SetElementId("")
	serc.SetId(1)
	serc.SetName("")
	serc.SetResource(res)
	serc.SetSp(*orig)

	res.SetActivation(act)
	res.SetDescription("") //JSON SAY ITS ""
	res.SetElementId("_FA0CB14C-7937-4144-82C1-3EEDE46642AF")
	res.SetId(51)
	res.SetName("app-2")
	res.SetServiceConnection(serc)
	res.SetX(603.0)
	res.SetY(122.0)

	act.SetDescription("") //JSON SAY ITS ""
	act.SetElementId("_71C07040-57A8-422B-8FB6-621423A7C769")
	act.SetExecutionEnv(exe)
	act.SetId(52)
	act.SetName("app-2-to-tc-2-activation")
	act.SetResource(res)
	act.SetSp(*orig)

	exe.SetActive(false)
	exe.SetBindingLocation(loca)
	exe.SetDescription("") //JSON SAY ITS ""
	exe.SetDisplayName("") //JSON SAY ITS NULL
	exe.SetElementId("_D4464EBE-CFEC-4A78-A645-36DC8DF4567F")
	exe.SetId(53)
	exe.SetInstallDemoApps(false)
	exe.SetInstallUri("") //JSON SAY ITS ""
	exe.SetLocation("")   //JSON SAY ITS ""
	exe.SetName("tc-2")
	exe.SetOverwriteOriginalSetup(false)
	exe.SetPlatformId("tc85")
	exe.SetTargetJDK("") //JSON SAY ITS NULL
	exe.SetType("LOCAL")

	loca.SetContext("")
	loca.SetElementId("IDBUS")
	loca.SetHost("josso")
	loca.SetId(0)
	loca.SetLocationAsString("")
	loca.SetPort(8081)
	loca.SetProtocol("http")
	loca.SetUri("ACC-03/SP-2")

	poi1 := api.NewPointDTO()
	poi1.SetX(1)
	poi1.SetY(1)
	poi = append(poi, *poi1)
	poi2 := api.NewPointDTO()
	poi2.SetX(1)
	poi2.SetY(1)
	poi = append(poi, *poi1)
	scdto.SetWaypoints(poi)
	serc.SetWaypoints(poi)
	act.SetWaypoints(poi)

	acts1.SetDescription("")
	acts1.SetElementId("")
	acts1.SetExecutionEnv(exe)
	acts1.SetId(1)
	acts1.SetName("")
	acts1.SetResource(res)
	acts1.SetSp(*orig)
	acts1.SetWaypoints(poi)
	acts2 := api.NewActivationDTO()
	acts2.SetDescription("")
	acts2.SetElementId("")
	acts2.SetExecutionEnv(exe)
	acts2.SetId(1)
	acts2.SetName("")
	acts2.SetResource(res)
	acts2.SetSp(*orig)
	acts2.SetWaypoints(poi)
	acts = append(acts, *acts1, *acts2)
	exe.SetActivations(acts)

	orig.SetAccountLinkagePolicy(alp)
	// orig.SetActiveBindings() //
	// orig.SetActiveProfiles() //
	orig.SetAuthenticationContract(AuthenticationContractDTO)
	orig.SetAuthenticationMechanism(AuthenticationMechanism)
	orig.SetConfig(conf)
	orig.SetDashboardUrl("") // JSON SAY ITS ""
	orig.SetDescription("IntSaml2Sp One")
	orig.SetDisplayName("")
	orig.SetElementId("_B9FCB070-1856-4375-8429-C04BF79E457A")
	orig.SetEnableMetadataEndpoint(true)
	orig.SetErrorBinding("JSON")
	orig.SetFederatedConnectionsA(fedconn)
	orig.SetFederatedConnectionsB(fedconn)
	orig.SetId(47)
	orig.SetIdentityAppliance(identityAppliance)
	orig.SetIdentityMappingPolicy(imp)
	orig.SetIsRemote(true)
	orig.SetLocation(loca)
	orig.SetMessageTtl(300)
	orig.SetMessageTtlTolerance(300)
	orig.SetName(name)
	orig.SetRemote(true)
	orig.SetRole("")
	orig.SetServiceConnection(scdto)
	orig.SetSignAuthenticationRequests(true)
	orig.SetSignRequests(true)
	orig.SetSignatureHash("")
	orig.SetWantAssertionSigned(false)
	orig.SetWantSLOResponseSigned(false)
	orig.SetWantAssertionSigned(false)
	metadata := api.NewResourceDTO()
	metadata.SetValue(encMetadata)
	orig.SetMetadata(*metadata)
	return orig
}

func (s *AccTestSuite) TestAccCliIntSaml2_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIntSaml2_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func InternalSaml2SpFieldTestCreate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		// {
		// 	name:     "AccountLinkagePolicy",
		// 	cmp:      func() bool { return StrPtrEquals(e.AccountLinkagePolicy, r.AccountLinkagePolicy) },
		// 	expected: StrDeref(e.AccountLinkagePolicy),
		// 	received: StrDeref(r.AccountLinkagePolicy),
		// },
		// {
		// 	name:     "Config",
		// 	cmp:      func() bool { return StrPtrEquals(e.Config, r.Config) },
		// 	expected: StrDeref(e.Config),
		// 	received: StrDeref(r.Config),
		// },
		{
			name:     "dashboardUrl",
			cmp:      func() bool { return StrPtrEquals(e.DashboardUrl, r.DashboardUrl) },
			expected: StrDeref(e.DashboardUrl),
			received: StrDeref(r.DashboardUrl),
		},
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "enablemetadataendpoint",
			cmp:      func() bool { return BoolPtrEquals(e.EnableMetadataEndpoint, r.EnableMetadataEndpoint) },
			expected: strconv.FormatBool(BoolDeref(e.EnableMetadataEndpoint)),
			received: strconv.FormatBool(BoolDeref(r.EnableMetadataEndpoint)),
		},
		{
			name:     "errorBinding",
			cmp:      func() bool { return StrPtrEquals(e.ErrorBinding, r.ErrorBinding) },
			expected: StrDeref(e.ErrorBinding),
			received: StrDeref(r.ErrorBinding),
		},
		// {
		// 	name:     "IdentityMappingPolicy",
		// 	cmp:      func() bool { return Int32PtrEquals(e.IdentityMappingPolicy, r.IdentityMappingPolicy) },
		// 	expected: strconv.Itoa(int(Int32Deref(e.IdentityMappingPolicy))),
		// 	received: strconv.Itoa(int(Int32Deref(r.IdentityMappingPolicy))),
		// },
		{
			name:     "messageTtl",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtl, r.MessageTtl) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtl))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtl))),
		},
		{
			name:     "messageTtlTolerance",
			cmp:      func() bool { return Int32PtrEquals(e.MessageTtlTolerance, r.MessageTtlTolerance) },
			expected: strconv.Itoa(int(Int32Deref(e.MessageTtlTolerance))),
			received: strconv.Itoa(int(Int32Deref(r.MessageTtlTolerance))),
		},
		// {
		// 	name:     "ServiceConnection",
		// 	cmp:      func() bool { return StrPtrEquals(e.ServiceConnection, r.ServiceConnection) },
		// 	expected: StrDeref(e.ServiceConnection),
		// 	received: StrDeref(r.ServiceConnection),
		// },
		{
			name:     "SignRequests",
			cmp:      func() bool { return BoolPtrEquals(e.SignRequests, r.SignRequests) },
			expected: strconv.FormatBool(BoolDeref(e.SignRequests)),
			received: strconv.FormatBool(BoolDeref(r.SignRequests)),
		},
		{
			name:     "SignatureHash",
			cmp:      func() bool { return StrPtrEquals(e.SignatureHash, r.SignatureHash) },
			expected: StrDeref(e.SignatureHash),
			received: StrDeref(r.SignatureHash),
		},
		{
			name:     "WantAssertionSigned",
			cmp:      func() bool { return BoolPtrEquals(e.WantAssertionSigned, r.WantAssertionSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantAssertionSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantAssertionSigned)),
		},
		{
			name:     "WantSLOResponseSigned",
			cmp:      func() bool { return BoolPtrEquals(e.WantSLOResponseSigned, r.WantSLOResponseSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantSLOResponseSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantSLOResponseSigned)),
		},
		{
			name:     "WantAssertionSigned",
			cmp:      func() bool { return BoolPtrEquals(e.WantAssertionSigned, r.WantAssertionSigned) },
			expected: strconv.FormatBool(BoolDeref(e.WantAssertionSigned)),
			received: strconv.FormatBool(BoolDeref(r.WantAssertionSigned)),
		},
	}
}

//Fields to validate after Sp update
func InteralSaml2SpFieldTestUpdate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		},
		{
			name:     "elementId",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
	}
	return append(t, InternalSaml2SpFieldTestCreate(e, r)...)
}

// Compares the expected Sp with the received one.
func InteralSaml2SpValidateCreate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) error {

	return ValidateFields(InternalSaml2SpFieldTestCreate(e, r))
}

// Compares the expected Sp with the received one.
func InternalSaml2SpValidateUpdate(
	e *api.InternalSaml2ServiceProviderDTO,
	r *api.InternalSaml2ServiceProviderDTO) error {

	return ValidateFields(InteralSaml2SpFieldTestUpdate(e, r))
}
