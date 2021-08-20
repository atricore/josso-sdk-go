package cli

import (
	"encoding/base64"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

const (
	metadata = `<?xml version="1.0" encoding="UTF-8"?>
	<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="http://www.webex.com"><md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><md:KeyDescriptor use="signing"><ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:X509Data><ds:X509Certificate>MIIC5jCCAc6gAwIBAgIGATtZlyKOMA0GCSqGSIb3DQEBBQUAMDQxCzAJBgNVBAYTAlVTMSUwIwYDVQQDExxXZWJFeCBDb21tdW5pY2F0aW9ucyBJbmMuIENBMB4XDTEyMTIwMjAzMDkzNVoXDTIyMTEzMDAzMDkzNVowNDELMAkGA1UEBhMCVVMxJTAjBgNVBAMTHFdlYkV4IENvbW11bmljYXRpb25zIEluYy4gQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCM925ZnQNyrYX2jyI+x+AEswI5hT8Ueyc4ignaYdFdwfOuUVBOk/fLUW9XbVKCqKxALyXmf1tdzvPM/PHJMIG8YK+EsZ7OCeVmDpO7obMHpcZk+sSF8CAUkwVmWi5bhypec+7tdtCvdnxJEOcNP2URXHMgxOFETXXo+WQ6hAqG92TpaL1V3w515OAixhDXfOlDGqxgrDgM59ChVherdCuwOilTNWFWdSoOpZi1niyuG1ukDsefl2YfAT6d3fj7UFkvr7n68EyqlFmmksDglEN0PSybw3ZFOBaSYwhF0S6orPBZAupIG6UZvJoSeOzsRSTbvNpBtwmIlCIGqA3srRnlAgMBAAEwDQYJKoZIhvcNAQEFBQADggEBABVY4TH/jajsE+dOFNL8sq3PIN/ZpRww7bskTNQKy9mpqCu4Rp5UpCbgz6VDAdBRyQo1jlyWMUltyVJugOeXyw7ebgB73iVM8F7Dl5hym6kJwB6TGa5IZnabHQQHWXH1zWdeflbcgv06yDKOszCmVLJskcnRmDisA4xs4FGBYZ+Dn4nI8DltQtDTt4yW6Clc3ZU+3VJUiKOxaM887adlBR6nUAkvEvXgr918AEmGHWIX6/yLnN+CFQoa3Gg/InKRCXJmgCAD6g/4yaDatB78M7BrXtm6cSFQ8ub+xGqvBIJ/H17dXiDC+x6TJmlwPghGNNkkP8jxz0QG5781sDbY7nY=</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</md:NameIDFormat><md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://shrm.webex.com/dispatcher/SAML2AuthService?siteurl=shrm" index="0" isDefault="true"/></md:SPSSODescriptor><md:Organization><md:OrganizationName xml:lang="en">Cisco WebEx</md:OrganizationName><md:OrganizationDisplayName xml:lang="en">Cisco WebEx</md:OrganizationDisplayName><md:OrganizationURL xml:lang="en"/></md:Organization><md:ContactPerson contactType="technical"><md:Company>Cisco WebEx</md:Company><md:GivenName/><md:SurName/><md:EmailAddress/><md:TelephoneNumber/></md:ContactPerson></md:EntityDescriptor>`
)

func (s *AccTestSuite) TestAccCliExtSaml2_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Extsmal2-a"
	var orig *api.ExternalSaml2ServiceProviderDTO
	var created api.ExternalSaml2ServiceProviderDTO
	orig = createTestExternalSaml2ServiceProviderDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateExtSaml2Sp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := ExternalSaml2SpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.ExternalSaml2ServiceProviderDTO
	read, err = s.client.GetExtSaml2Sp(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ExternalSaml2SpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating Sp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.DisplayName = api.PtrString("Atricore")
	updated, err := s.client.UpdateExtSaml2Sp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ExternalSaml2SpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	//Test Delete
	deleted, err := s.client.DeleteExtSaml2Sp(*appliance.Name, crudName)
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

	listOfAll, err := s.client.GetExtSaml2Sps(*appliance.Name)
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
	var listOfCreated [2]api.ExternalSaml2ServiceProviderDTO

	// Test list of #2 elements
	element1 := createTestExternalSaml2ServiceProviderDTO("Extsmal2-1")
	listOfCreated[0], _ = s.client.CreateExtSaml2Sp(*appliance.Name, *element1)

	element2 := createTestExternalSaml2ServiceProviderDTO("Extsmal2-2")
	listOfCreated[1], _ = s.client.CreateExtSaml2Sp(*appliance.Name, *element2)

	// ------------------------
	// Get list from server
	listOfRead, err := s.client.GetExtSaml2Sps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	// The list should have listOfCreated lenght elemetns
	if len(listOfRead) != len(listOfCreated) {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expected %d", len(listOfAll), len(listOfCreated))
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
		if err = ExternalSaml2SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestExternalSaml2ServiceProviderDTO(name string) *api.ExternalSaml2ServiceProviderDTO {
	encMetadata := base64.StdEncoding.EncodeToString([]byte(metadata))
	//encMetadata := metadata
	orig := api.NewExternalSaml2ServiceProviderDTO()
	var fedconn []api.FederatedConnectionDTO
	fedconn1 := api.NewFederatedConnectionDTO()
	fedconn2 := api.NewFederatedConnectionDTO()
	var AuthenticationMechanismDTO []api.AuthenticationMechanismDTO
	AuthenticationMechanismDTO1 := api.NewAuthenticationMechanismDTO()
	AuthenticationMechanismDTO2 := api.NewAuthenticationMechanismDTO()
	var FederatedChannelDTO api.FederatedChannelDTO
	var FederatedProviderDTO api.FederatedProviderDTO
	var ResourceDTO api.ResourceDTO
	var identityAppliance api.IdentityApplianceDefinitionDTO
	var AuthenticationServiceDTO api.AuthenticationServiceDTO
	var AuthenticationService []api.AuthenticationServiceDTO
	AuthenticationService1 := api.NewAuthenticationServiceDTO()
	AuthenticationService2 := api.NewAuthenticationServiceDTO()
	var locat api.LocationDTO
	var poi []api.PointDTO
	var conf api.ProviderConfigDTO
	var OAuth2ClientDTO []api.OAuth2ClientDTO
	OAuth2ClientDTO1 := api.NewOAuth2ClientDTO()
	OAuth2ClientDTO2 := api.NewOAuth2ClientDTO()
	var DelegatedAuthenticationDTO []api.DelegatedAuthenticationDTO
	DelegatedAuthenticationDTO1 := api.NewDelegatedAuthenticationDTO()
	DelegatedAuthenticationDTO2 := api.NewDelegatedAuthenticationDTO()
	var SubjectNameIdentifierPolicyDTO []api.SubjectNameIdentifierPolicyDTO
	SubjectNameIdentifierPolicyDTO1 := api.NewSubjectNameIdentifierPolicyDTO()
	SubjectNameIdentifierPolicyDTO2 := api.NewSubjectNameIdentifierPolicyDTO()
	var IdentityProviderDTO api.IdentityProviderDTO
	var AttributeProfileDTO api.AttributeProfileDTO
	var AuthenticationContractDTO api.AuthenticationContractDTO
	var AuthenticationAssertionEmissionPolicyDTO api.AuthenticationAssertionEmissionPolicyDTO
	var ExtensionDTO api.ExtensionDTO
	var SessionManagerFactoryDTO api.SessionManagerFactoryDTO
	var SubjectNameIdentifierPolicy api.SubjectNameIdentifierPolicyDTO
	var DelegatedAuthentication api.DelegatedAuthenticationDTO
	var ExecutionEnvironmentDTO []api.ExecutionEnvironmentDTO
	ExecutionEnvironmentDTO1 := api.NewExecutionEnvironmentDTO()
	ExecutionEnvironmentDTO2 := api.NewExecutionEnvironmentDTO()
	var ActivationDTO []api.ActivationDTO
	ActivationDTO1 := api.NewActivationDTO()
	ActivationDTO2 := api.NewActivationDTO()
	var ExecutionEnvironment api.ExecutionEnvironmentDTO
	var ServiceResourceDTO api.ServiceResourceDTO
	var ServiceConnectionDTO api.ServiceConnectionDTO
	var Activation api.ActivationDTO
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
	var FederatedConnection api.FederatedConnectionDTO
	var IdentityLookupDTO []api.IdentityLookupDTO
	IdentityLookupDTO1 := api.NewIdentityLookupDTO()
	IdentityLookupDTO2 := api.NewIdentityLookupDTO()
	var Provider api.ProviderDTO
	var IdentitySource api.IdentitySourceDTO

	IdentitySource.SetDescription("")
	IdentitySource.SetElementId("")
	IdentitySource.SetId(1)
	IdentitySource.SetName("")

	// Provider.SetActiveBindings() //
	// Provider.SetActiveProfiles() //
	Provider.SetConfig(conf)
	Provider.SetDescription("")
	Provider.SetDisplayName("")
	Provider.SetElementId("")
	Provider.SetId(1)
	Provider.SetIdentityAppliance(identityAppliance)
	Provider.SetIsRemote(true)
	Provider.SetLocation(locat)
	Provider.SetMetadata(ResourceDTO)
	Provider.SetName("")
	Provider.SetRemote(true)
	Provider.SetRole("")

	IdentityLookupDTO1.SetDescription("")
	IdentityLookupDTO1.SetElementId("")
	IdentityLookupDTO1.SetId(1)
	IdentityLookupDTO1.SetIdentitySource(IdentitySource)
	IdentityLookupDTO1.SetName("")
	IdentityLookupDTO1.SetProvider(Provider)
	IdentityLookupDTO1.SetWaypoints(poi)
	IdentityLookupDTO = append(IdentityLookupDTO, *IdentityLookupDTO1)
	IdentityLookupDTO2.SetDescription("")
	IdentityLookupDTO2.SetElementId("")
	IdentityLookupDTO2.SetId(2)
	IdentityLookupDTO2.SetIdentitySource(IdentitySource)
	IdentityLookupDTO2.SetName("")
	IdentityLookupDTO2.SetProvider(Provider)
	IdentityLookupDTO2.SetWaypoints(poi)
	IdentityLookupDTO = append(IdentityLookupDTO, *IdentityLookupDTO2)

	FederatedConnection.SetChannelA(FederatedChannelDTO)
	FederatedConnection.SetChannelB(FederatedChannelDTO)
	FederatedConnection.SetDescription("")
	FederatedConnection.SetElementId("")
	FederatedConnection.SetId(1)
	FederatedConnection.SetName("")
	FederatedConnection.SetRoleA(FederatedProviderDTO)
	FederatedConnection.SetRoleB(FederatedProviderDTO)
	FederatedConnection.SetWaypoints(poi)

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

	ServiceResourceDTO.SetActivation(Activation)
	ServiceResourceDTO.SetDescription("")
	ServiceResourceDTO.SetElementId("")
	ServiceResourceDTO.SetId(1)
	ServiceResourceDTO.SetName("")
	ServiceResourceDTO.SetServiceConnection(ServiceConnectionDTO)

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

	DelegatedAuthentication.SetAuthnService(AuthenticationServiceDTO)
	DelegatedAuthentication.SetDescription("")
	DelegatedAuthentication.SetElementId("")
	DelegatedAuthentication.SetId(1)
	DelegatedAuthentication.SetIdp(IdentityProviderDTO)
	DelegatedAuthentication.SetName("")
	DelegatedAuthentication.SetWaypoints(poi)

	SubjectNameIdentifierPolicy.SetDescriptionKey("")
	SubjectNameIdentifierPolicy.SetId("")
	SubjectNameIdentifierPolicy.SetName("")
	SubjectNameIdentifierPolicy.SetSubjectAttribute("")
	SubjectNameIdentifierPolicy.SetType("")

	SubjectNameIdentifierPolicyDTO1.SetDescriptionKey("")
	SubjectNameIdentifierPolicyDTO1.SetId("")
	SubjectNameIdentifierPolicyDTO1.SetName("")
	SubjectNameIdentifierPolicyDTO1.SetSubjectAttribute("")
	SubjectNameIdentifierPolicyDTO1.SetType("")
	SubjectNameIdentifierPolicyDTO = append(SubjectNameIdentifierPolicyDTO, *SubjectNameIdentifierPolicyDTO1)
	SubjectNameIdentifierPolicyDTO2.SetDescriptionKey("")
	SubjectNameIdentifierPolicyDTO2.SetId("")
	SubjectNameIdentifierPolicyDTO2.SetName("")
	SubjectNameIdentifierPolicyDTO2.SetSubjectAttribute("")
	SubjectNameIdentifierPolicyDTO2.SetType("")
	SubjectNameIdentifierPolicyDTO = append(SubjectNameIdentifierPolicyDTO, *SubjectNameIdentifierPolicyDTO2)

	OAuth2ClientDTO1.SetBaseURL("")
	OAuth2ClientDTO1.SetId("")
	OAuth2ClientDTO1.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO1)
	OAuth2ClientDTO2.SetBaseURL("")
	OAuth2ClientDTO2.SetId("")
	OAuth2ClientDTO2.SetSecret("")
	OAuth2ClientDTO = append(OAuth2ClientDTO, *OAuth2ClientDTO2)

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

	SessionManagerFactoryDTO.SetDescription("")
	SessionManagerFactoryDTO.SetName("")

	ExtensionDTO.SetClassifier("")
	ExtensionDTO.SetId("")
	ExtensionDTO.SetName("")
	ExtensionDTO.SetNamespace("")
	ExtensionDTO.SetProvider("")
	ExtensionDTO.SetVersion("")

	AuthenticationAssertionEmissionPolicyDTO.SetElementId("")
	AuthenticationAssertionEmissionPolicyDTO.SetId(1)
	AuthenticationAssertionEmissionPolicyDTO.SetName("")

	AuthenticationContractDTO.SetElementId("")
	AuthenticationContractDTO.SetId(1)
	AuthenticationContractDTO.SetName("")

	AttributeProfileDTO.SetElementId("")
	AttributeProfileDTO.SetId(1)
	AttributeProfileDTO.SetName("")
	AttributeProfileDTO.SetProfileType("")

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

	AuthenticationServiceDTO.SetDelegatedAuthentications(DelegatedAuthenticationDTO)
	AuthenticationServiceDTO.SetDescription("")
	AuthenticationServiceDTO.SetDisplayName("")
	AuthenticationServiceDTO.SetElementId("")
	AuthenticationServiceDTO.SetId(1)
	AuthenticationServiceDTO.SetName("")

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

	conf.SetDescription("")
	conf.SetDisplayName("")
	conf.SetElementId("")
	conf.SetId(1)
	conf.SetName("")

	poi1 := api.NewPointDTO()
	poi1.SetX(1)
	poi1.SetY(1)
	poi = append(poi, *poi1)
	poi2 := api.NewPointDTO()
	poi2.SetX(1)
	poi2.SetY(1)
	poi = append(poi, *poi1)

	locat.SetContext("")
	locat.SetElementId("")
	locat.SetHost("")
	locat.SetId(1)
	locat.SetLocationAsString("")
	locat.SetPort(1)
	locat.SetProtocol("")
	locat.SetUri("")

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

	// orig.SetActiveBindings("") //////////// preguntar
	// orig.SetActiveProfiles("") //////////// preguntar
	orig.SetConfig(conf)
	orig.SetDescription("")
	orig.SetDisplayName("")
	orig.SetElementId("")
	orig.SetFederatedConnectionsA(fedconn)
	orig.SetFederatedConnectionsB(fedconn)
	orig.SetId(1)
	orig.SetIdentityAppliance(identityAppliance)
	orig.SetIdentityLookups(IdentityLookupDTO)
	orig.SetIsRemote(true)
	orig.SetLocation(locat)
	metadata := api.NewResourceDTO()
	metadata.SetValue(encMetadata)
	orig.SetMetadata(*metadata)
	orig.SetName("")
	orig.SetRemote(true)
	orig.SetRole("")
	return orig
}

func (s *AccTestSuite) TestAccCliExtSaml2_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliExtSaml2_crud_updateFailOnDupName() {

	// TODO ! implement me!

}

// --------------------------------------------------------

//Fields to validate after appliance creation
func ExternalSaml2SpFieldTestCreate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		/* TODO : Change for MD test
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: LocationToStr(e.Location),
			received: LocationToStr(r.Location),
		}, */
	}
}

//Fields to validate after ExternalSaml2Sp update
func ExternalSaml2SpFieldTestUpdate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}
	return append(t, ExternalSaml2SpFieldTestCreate(e, r)...)
}

// Compares the expected ExternalSaml2Sp with the received one.
func ExternalSaml2SpValidateCreate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) error {

	return ValidateFields(ExternalSaml2SpFieldTestCreate(e, r))
}

// Compares the expected ExternalSaml2Sp with the received one.
func ExternalSaml2SpValidateUpdate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) error {

	return ValidateFields(ExternalSaml2SpFieldTestUpdate(e, r))
}
