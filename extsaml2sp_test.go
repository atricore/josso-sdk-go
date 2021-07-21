package cli

import (
	"sort"
	"strings"

	api "github.com/atricore/josso-api-go"
)

const (
	metadata = `<?xml version="1.0" encoding="UTF-8"?>
				<ns6:EntityDescriptor 
					xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
					xmlns:ns7="urn:org:atricore:idbus:common:sso:1.0:protocol" 
					xmlns:ns6="urn:oasis:names:tc:SAML:2.0:metadata" 
					xmlns:ns5="urn:oasis:names:tc:SAML:2.0:idbus" 
					xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol" 
					xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion" 
					xmlns:ds="http://www.w3.org/2000/09/xmldsig#" 
					xmlns:enc="http://www.w3.org/2001/04/xmlenc#" 
					ID="_D21F6A50-6C94-4E1B-B416-08A26E996882" 
					entityID="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MD">
					<ns6:SPSSODescriptor 
						WantAssertionsSigned="false" 
						AuthnRequestsSigned="false" 
						protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol" 
						ID="_D21F6A50-6C94-4E1B-B416-08A26E996882sp">
						<ns6:KeyDescriptor use="signing">
							<ds:KeyInfo>
								<ds:X509Data>
									<ds:X509Certificate>MIIDojCCAooCCQCVTd3p5WnWmjANBgkqhkiG9w0BAQsFADCBkjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2NvMREwDwYDVQQKDAhhdHJpY29yZTENMAsGA1UECwwEZGVtbzEXMBUGA1UEAwwOam9zc28tcHJvdmlkZXIxIzAhBgkqhkiG9w0BCQEWFHN1cHBvcnRAYXRyaWNvcmUuY29tMB4XDTE2MDIwMjE3MDIwM1oXDTI2MDEzMDE3MDIwM1owgZIxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzERMA8GA1UECgwIYXRyaWNvcmUxDTALBgNVBAsMBGRlbW8xFzAVBgNVBAMMDmpvc3NvLXByb3ZpZGVyMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGF0cmljb3JlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKCBJiMEjYh2Id50qMGGuZzivqFy7t3IwsJgjbS+xV3Jf5MmPyXh1AsYpk8eKSYDb+H8+hROeqxbSneXjAi5msrD+oCJnMwz0/uMUPsmntjlrbWSe2P2vGfLWLp708YLh2RyAA3Iz2Vx5fdbN+14zPfdMF/uNuD4e8XTU7PJcX4cIPna58P1ko3mCMVoPFI2KLess/EafBvc5OBBmTo3KeQ59hGRdNtCe5oeuLHapfLWnl36MHHkV/sdV+xVV/NsO5lVJ4al/n7snOsqBvUm++Zbey1OI3CWp9+q1CnnqFxzRiJySahYF5FoSiWJKpw7tXHkyU93FCVeBV5c5zxqVykCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAU27Ag+jrg+xVbRZc3Dqk40PitlvLiT619U8eyt0LHAhX+ZGy/Ao+pJAxSWHLP6YofG+EO3Fl4sgJ5S9py+PZwDgRQR1xfUsZ5a8tk6c0NPHpcHBU2pMuYQA+OoE7g5EIeAhPsmMeM2IH4Yz6qmzhvYBAvbDvGJYHi+Udxp8JHlKYjkieVw+9kI580YKeUIKXng4XXSuFHspYRLS2iDRfmeJsveOUYr9y7L4XrbLJIG/kVcpFiLkzsWJp1j6hwqPe748wekASae/+96l3NjT1AyNnD7rzyskUiNI6wb28OZeJoPczgzIedKXYdmFqLRuLeSLDJK2EiUATRUqE3ys7Fw==</ds:X509Certificate>
								</ds:X509Data>
							</ds:KeyInfo>
						</ns6:KeyDescriptor>
						<ns6:KeyDescriptor 
							use="encryption">
						<ds:KeyInfo>
							<ds:X509Data>
								<ds:X509Certificate>MIIDojCCAooCCQCVTd3p5WnWmjANBgkqhkiG9w0BAQsFADCBkjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2NvMREwDwYDVQQKDAhhdHJpY29yZTENMAsGA1UECwwEZGVtbzEXMBUGA1UEAwwOam9zc28tcHJvdmlkZXIxIzAhBgkqhkiG9w0BCQEWFHN1cHBvcnRAYXRyaWNvcmUuY29tMB4XDTE2MDIwMjE3MDIwM1oXDTI2MDEzMDE3MDIwM1owgZIxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzERMA8GA1UECgwIYXRyaWNvcmUxDTALBgNVBAsMBGRlbW8xFzAVBgNVBAMMDmpvc3NvLXByb3ZpZGVyMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGF0cmljb3JlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKCBJiMEjYh2Id50qMGGuZzivqFy7t3IwsJgjbS+xV3Jf5MmPyXh1AsYpk8eKSYDb+H8+hROeqxbSneXjAi5msrD+oCJnMwz0/uMUPsmntjlrbWSe2P2vGfLWLp708YLh2RyAA3Iz2Vx5fdbN+14zPfdMF/uNuD4e8XTU7PJcX4cIPna58P1ko3mCMVoPFI2KLess/EafBvc5OBBmTo3KeQ59hGRdNtCe5oeuLHapfLWnl36MHHkV/sdV+xVV/NsO5lVJ4al/n7snOsqBvUm++Zbey1OI3CWp9+q1CnnqFxzRiJySahYF5FoSiWJKpw7tXHkyU93FCVeBV5c5zxqVykCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAU27Ag+jrg+xVbRZc3Dqk40PitlvLiT619U8eyt0LHAhX+ZGy/Ao+pJAxSWHLP6YofG+EO3Fl4sgJ5S9py+PZwDgRQR1xfUsZ5a8tk6c0NPHpcHBU2pMuYQA+OoE7g5EIeAhPsmMeM2IH4Yz6qmzhvYBAvbDvGJYHi+Udxp8JHlKYjkieVw+9kI580YKeUIKXng4XXSuFHspYRLS2iDRfmeJsveOUYr9y7L4XrbLJIG/kVcpFiLkzsWJp1j6hwqPe748wekASae/+96l3NjT1AyNnD7rzyskUiNI6wb28OZeJoPczgzIedKXYdmFqLRuLeSLDJK2EiUATRUqE3ys7Fw==</ds:X509Certificate>
							</ds:X509Data>
						</ds:KeyInfo>
					<ns6:EncryptionMethod 
						Algorithm="http://www.w3.org/2001/04/xmlenc#aes128-cbc">
						<enc:KeySize>128</enc:KeySize>
				</ns6:EncryptionMethod>
				<ns6:EncryptionMethod 
					Algorithm="http://www.w3.org/2001/04/xmlenc#aes256-cbc">
					<enc:KeySize>256</enc:KeySize>
				</ns6:EncryptionMethod>
				<ns6:EncryptionMethod Algorithm="http://www.w3.org/2001/04/xmlenc#tripledes-cbc">
					<enc:KeySize>192</enc:KeySize>
				</ns6:EncryptionMethod>
			</ns6:KeyDescriptor>
			<ns6:ArtifactResolutionService 
				isDefault="true" 
				index="0" 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/ARTIFACT/SOAP"
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"/>
			<ns6:SingleLogoutService 
				ResponseLocation="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO_RESPONSE/POST"
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO/POST"
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"/>
			<ns6:SingleLogoutService 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO/SOAP" 
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"/>
			<ns6:ManageNameIDService 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/SOAP" 
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"/>
			<ns6:ManageNameIDService 
				ResponseLocation="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI_RESPONSE/POST" 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/POST" 
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"/>
			<ns6:ManageNameIDService 
				ResponseLocation="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI_RESPONSE/REDIR" 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/REDIR" 
				Binding="urn:oasis:n				protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol" ames:tc:SAML:2.0:bindings:HTTP-Redirect"/>
			<ns6:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</ns6:NameIDFormat>
			<ns6:AssertionConsumerService 
				isDefault="true" 
				index="0" 
				Location="http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/ACS/POST" 
				Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"/>
			</ns6:SPSSODescriptor>
			<ns6:Organization>
				<ns6:OrganizationName xml:lang="en">Atricore IDBUs SAMLR2 JOSSO SP Sample</ns6:OrganizationName>
				<ns6:OrganizationDisplayName xml:lang="en">Atricore, Inc.</ns6:OrganizationDisplayName>
				<ns6:OrganizationURL xml:lang="en">http://www.atricore.org</ns6:OrganizationURL>
			</ns6:Organization>
			<ns6:ContactPerson contactType="other"/>
		</ns6:EntityDescriptor>`
)

func toStringArr(s string) *[]string {
	v := []string{s} // Creates array with string in first position
	return &v        // retunrs pointer to array
}

func (s *AccTestSuite) TestAccCliExtSaml2_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	var created api.ExternalSaml2ServiceProviderDTO
	orig := api.ExternalSaml2ServiceProviderDTO{
		Name:        api.PtrString("Extsmal2-2"),
		Id:          api.PtrInt64(-1),
		Description: api.PtrString("My SP 2"),
		Metadata: &api.ResourceDTO{
			Value: toStringArr(metadata),
			Uri:   api.PtrString("metadata-a.xml"),
		},
	}
	// Test CREATE
	created, err = s.client.CreateExtSaml2Sp(*appliance.Name, orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := SpValidateCreate(&orig, &created); err != nil {
		t.Errorf("creating sp : %v", err)
		return
	}

	// Test READ
	var read api.ExternalSaml2ServiceProviderDTO
	read, err = s.client.GetExtSaml2Sp(*appliance.Name, "Extsmal2-2")
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
	//read.DashboardUrl = api.PtrString("12345")
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
	toDelete := "Extsmal2-2"
	deleted, err := s.client.DeleteExtSaml2Sp(*appliance.Name, toDelete)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", toDelete)
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
	element1 := api.ExternalSaml2ServiceProviderDTO{
		Name: api.PtrString("Extsmal2-1"),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[0], _ = s.client.CreateExtSaml2Sp(*appliance.Name, element1)

	element2 := api.ExternalSaml2ServiceProviderDTO{
		Name: api.PtrString("Extsmal2-2"),
		Id:   api.PtrInt64(-1),
	}
	listOfCreated[1], _ = s.client.CreateExtSaml2Sp(*appliance.Name, element2)

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
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) > 0
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
			expected: e.Name,
			received: r.Name,
		},
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
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "location",
			cmp:      func() bool { return LocationPtrEquals(e.Location, r.Location) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "DisplayName",
			cmp:      func() bool { return StrPtrEquals(e.DisplayName, r.DisplayName) },
			expected: e.DisplayName,
			received: r.DisplayName,
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
