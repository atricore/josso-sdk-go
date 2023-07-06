package cli

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"

	api "github.com/atricore/josso-api-go"
)

const (
	sp_metadata = `<?xml version="1.0" encoding="UTF-8"?>
	<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="http://www.webex.com"><md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><md:KeyDescriptor use="signing"><ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:X509Data><ds:X509Certificate>MIIC5jCCAc6gAwIBAgIGATtZlyKOMA0GCSqGSIb3DQEBBQUAMDQxCzAJBgNVBAYTAlVTMSUwIwYDVQQDExxXZWJFeCBDb21tdW5pY2F0aW9ucyBJbmMuIENBMB4XDTEyMTIwMjAzMDkzNVoXDTIyMTEzMDAzMDkzNVowNDELMAkGA1UEBhMCVVMxJTAjBgNVBAMTHFdlYkV4IENvbW11bmljYXRpb25zIEluYy4gQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCM925ZnQNyrYX2jyI+x+AEswI5hT8Ueyc4ignaYdFdwfOuUVBOk/fLUW9XbVKCqKxALyXmf1tdzvPM/PHJMIG8YK+EsZ7OCeVmDpO7obMHpcZk+sSF8CAUkwVmWi5bhypec+7tdtCvdnxJEOcNP2URXHMgxOFETXXo+WQ6hAqG92TpaL1V3w515OAixhDXfOlDGqxgrDgM59ChVherdCuwOilTNWFWdSoOpZi1niyuG1ukDsefl2YfAT6d3fj7UFkvr7n68EyqlFmmksDglEN0PSybw3ZFOBaSYwhF0S6orPBZAupIG6UZvJoSeOzsRSTbvNpBtwmIlCIGqA3srRnlAgMBAAEwDQYJKoZIhvcNAQEFBQADggEBABVY4TH/jajsE+dOFNL8sq3PIN/ZpRww7bskTNQKy9mpqCu4Rp5UpCbgz6VDAdBRyQo1jlyWMUltyVJugOeXyw7ebgB73iVM8F7Dl5hym6kJwB6TGa5IZnabHQQHWXH1zWdeflbcgv06yDKOszCmVLJskcnRmDisA4xs4FGBYZ+Dn4nI8DltQtDTt4yW6Clc3ZU+3VJUiKOxaM887adlBR6nUAkvEvXgr918AEmGHWIX6/yLnN+CFQoa3Gg/InKRCXJmgCAD6g/4yaDatB78M7BrXtm6cSFQ8ub+xGqvBIJ/H17dXiDC+x6TJmlwPghGNNkkP8jxz0QG5781sDbY7nY=</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</md:NameIDFormat><md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</md:NameIDFormat><md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://shrm.webex.com/dispatcher/SAML2AuthService?siteurl=shrm" index="0" isDefault="true"/></md:SPSSODescriptor><md:Organization><md:OrganizationName xml:lang="en">Cisco WebEx</md:OrganizationName><md:OrganizationDisplayName xml:lang="en">Cisco WebEx</md:OrganizationDisplayName><md:OrganizationURL xml:lang="en"/></md:Organization><md:ContactPerson contactType="technical"><md:Company>Cisco WebEx</md:Company><md:GivenName/><md:SurName/><md:EmailAddress/><md:TelephoneNumber/></md:ContactPerson></md:EntityDescriptor>`
)

func (s *AccTestSuite) TestAccCliExtSaml2SP_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}
	crudName := "Extsaml2-a"
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
	element1 := createTestExternalSaml2ServiceProviderDTO("Extsaml2-1")
	listOfCreated[0], _ = s.client.CreateExtSaml2Sp(*appliance.Name, *element1)

	element2 := createTestExternalSaml2ServiceProviderDTO("Extsaml2-2")
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
	encMetadata := base64.StdEncoding.EncodeToString([]byte(sp_metadata))
	//encMetadata := metadata
	tData := api.NewExternalSaml2ServiceProviderDTO()
	tData.SetDescription(fmt.Sprintf("Desc %s", name))
	tData.SetDisplayName(fmt.Sprintf("DispName %s", name))
	md := api.NewResourceDTO()
	md.SetValue(encMetadata)
	tData.SetMetadata(*md)
	tData.SetName(name)
	return tData
}

func (s *AccTestSuite) TestAccCliExtSaml2SP_crud_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliExtSaml2SP_crud_updateFailOnDupName() {

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
		{
			name:     "description",
			cmp:      func() bool { return StrPtrEquals(e.Description, r.Description) },
			expected: StrDeref(e.Description),
			received: StrDeref(r.Description),
		},
		{
			name:     "display_name",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: StrDeref(e.DisplayName),
			received: StrDeref(r.DisplayName),
		},
		// {
		// 	name:     "metadata",
		// 	cmp:      func() bool { return StrPtrEquals(e.Metadata, r.Metadata) },
		// 	expected: StrDeref(e.Metadata),
		// 	received: StrDeref(r.Metadata),
		// },
	}
}

//Fields to validate after ExternalSaml2Sp update
func ExternalSaml2SpFieldTestUpdate(
	e *api.ExternalSaml2ServiceProviderDTO,
	r *api.ExternalSaml2ServiceProviderDTO) []FiledTestStruct {

	t := []FiledTestStruct{}
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
