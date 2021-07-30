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
	if err := SpValidateCreate(orig, &created); err != nil {
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
	if err = SpValidateUpdate(&read, &created); err != nil {
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
	if err = SpValidateUpdate(&read, &updated); err != nil {
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
		if err = SpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestExternalSaml2ServiceProviderDTO(name string) *api.ExternalSaml2ServiceProviderDTO {
	encMetadata := base64.StdEncoding.EncodeToString([]byte(metadata))
	//encMetadata := metadata
	orig := api.NewExternalSaml2ServiceProviderDTO()
	orig.SetName(name)
	orig.SetId(-1)
	orig.SetDescription("My Sp 2")
	metadata := api.NewResourceDTO()
	metadata.SetValue(encMetadata)
	orig.SetMetadata(*metadata)
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
func SpFieldTestCreate(
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

//Fields to validate after Sp update
func SpFieldTestUpdate(
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
	return append(t, SpFieldTestCreate(e, r)...)
}

// Compares the expected Sp with the received one.
func SpValidateCreate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) error {

	return ValidateFields(SpFieldTestCreate(e, r))
}

// Compares the expected Sp with the received one.
func SpValidateUpdate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) error {

	return ValidateFields(SpFieldTestUpdate(e, r))
}
