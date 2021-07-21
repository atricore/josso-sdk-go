go mod tidy
go mod vendor
==> Checking that code complies with gofmt requirements...
go clean -testcache
TF_ACC=1 go test $(go list ./... |grep -v 'vendor') -v   -timeout 120m
=== RUN   TestAccCliSuite
    client_test.go:25: creating client
2021/07/21 14:21:50 newIdbusApiClient TRACE: true
2021/07/21 14:21:50 Using client TRACE ON
2021/07/21 14:21:50 registering server http://localhost:8081/atricore-rest/services
2021/07/21 14:21:50 adding server configuration for http://localhost:8081/atricore-rest/services
2021/07/21 14:21:50 authn: idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7 true/admin true
2021/07/21 14:21:50 
POST /atricore-rest/services/iam-authn/sign-on HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 127
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"clientId":"idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7","password":"atricore","secret":"7oUHlv(HLT%vxK4L","username":"admin"}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 2206
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"server":null,"validationErrors":null,"idToken":"eyJhbGciOiJIUzI1NiJ9.eyJhdF9oYXNoIjoiYXZHVi1PUkxoNjk2bkZtUW5kWVBhdyIsInN1YiI6ImFkbWluIiwiYXVkIjoiaWRidXMtZjJmNzI0NGUtYmJjZS00NGNhLThiMzMtZjVjMGJkZTMzOWY3IiwiYXV0aF90aW1lIjoxNjI2ODg4MTEwLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsIm5hbWUiOiJBZG1pbmlzdHJhdG9yIiwiYWNjb3VudERpc2FibGVkIjoiZmFsc2UiLCJncm91cHMiOlsiQWRtaW5pc3RyYXRvcnMiXSwiZXhwIjoxNjI2ODg4NDEwLCJnaXZlbl9uYW1lIjoiQWRtaW5pc3RyYXRvciIsImlhdCI6MTYyNjg4ODExMCwiZmFtaWx5X25hbWUiOiJBZG1pbmlzdHJhdG9yIn0.Z4MFJ4l2xrRnAiq_VhGPr9pyOAvEr4PxA1QOgt1GxUE","accessToken":"eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM","refreshToken":"W--5YASE7WpYPEzFbaFhEvXM_rOxvE0nbLOQizEF8PY","authenticatedUser":{"@id":2,"id":"1","userName":"admin","firstName":"Administrator","surename":"Administrator","commonName":"Administrator","givenName":"Administrator","initials":null,"generationQualifier":null,"distinguishedName":null,"email":null,"telephoneNumber":null,"facsimilTelephoneNumber":null,"countryName":"Administrator","localityName":null,"stateOrProvinceName":null,"streetAddress":null,"organizationName":null,"organizationUnitName":null,"personalTitle":null,"businessCategory":null,"postalAddress":null,"postalCode":null,"postOfficeBox":null,"language":null,"groups":[{"@id":3,"id":"1","name":"Administrators","description":"Administrators","extraAttributes":[]}],"accountDisabled":false,"accountExpires":null,"accountExpirationDate":null,"limitSimultaneousLogin":null,"maximunLogins":null,"terminatePreviousSession":null,"preventNewSession":null,"allowUserToChangePassword":true,"forcePeriodicPasswordChanges":null,"daysBetweenChanges":null,"passwordExpirationDate":null,"notifyPasswordExpiration":null,"daysBeforeExpiration":null,"userPassword":"YHRsG4tlSmYpl6//r8OoIafxuqwne0plJ7HuzRZbtYY=","userCertificate":null,"automaticallyGeneratePassword":null,"emailNewPasword":null,"extraAttributes":[]}}
    client_test.go:33: created test client: [{http://localhost:8081/atricore-rest/services JOSSO Test server map[]}]
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/07/21 14:21:50 
GET /atricore-rest/services/iam-deploy/appliances HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 5
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

null

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 60
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"appliances":[]}
2021/07/21 14:21:50 getAppliances. found appliances 0
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud
2021/07/21 14:21:50 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 61
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":null}
2021/07/21 14:21:50 getAppliance. not found for ID/name ida-a
2021/07/21 14:21:50 createAppliance : ida-a com.atricore.idbus.ida.t
2021/07/21 14:21:50 
POST /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 191
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"appliance":{"description":"IDA-T TEST !","location":{"context":"IDBUS","host":"localhost","port":80,"protocol":"http","uri":"IDA-T"},"name":"ida-a","namespace":"com.atricore.idbus.ida.t"}}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":1,"lastModification":"2021-07-21T14:21:50-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:50 createAppliance. ID: 15
2021/07/21 14:21:50 createExtSaml2Sp : Extsmal2-2 [ida-a]
2021/07/21 14:21:50 
POST /atricore-rest/services/iam-deploy/extsaml2sp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 7512
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","sp":{"@c":".ExternalSaml2ServiceProviderDTO","description":"My SP 2","id":-1,"metadata":{"uri":"metadata-a.xml","value":["\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n\t\t\t\t\u003cns6:EntityDescriptor \n\t\t\t\t\txmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" \n\t\t\t\t\txmlns:ns7=\"urn:org:atricore:idbus:common:sso:1.0:protocol\" \n\t\t\t\t\txmlns:ns6=\"urn:oasis:names:tc:SAML:2.0:metadata\" \n\t\t\t\t\txmlns:ns5=\"urn:oasis:names:tc:SAML:2.0:idbus\" \n\t\t\t\t\txmlns:samlp=\"urn:oasis:names:tc:SAML:2.0:protocol\" \n\t\t\t\t\txmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" \n\t\t\t\t\txmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\" \n\t\t\t\t\txmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" \n\t\t\t\t\tID=\"_D21F6A50-6C94-4E1B-B416-08A26E996882\" \n\t\t\t\t\tentityID=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MD\"\u003e\n\t\t\t\t\t\u003cns6:SPSSODescriptor \n\t\t\t\t\t\tWantAssertionsSigned=\"false\" \n\t\t\t\t\t\tAuthnRequestsSigned=\"false\" \n\t\t\t\t\t\tprotocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\" \n\t\t\t\t\t\tID=\"_D21F6A50-6C94-4E1B-B416-08A26E996882sp\"\u003e\n\t\t\t\t\t\t\u003cns6:KeyDescriptor use=\"signing\"\u003e\n\t\t\t\t\t\t\t\u003cds:KeyInfo\u003e\n\t\t\t\t\t\t\t\t\u003cds:X509Data\u003e\n\t\t\t\t\t\t\t\t\t\u003cds:X509Certificate\u003eMIIDojCCAooCCQCVTd3p5WnWmjANBgkqhkiG9w0BAQsFADCBkjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2NvMREwDwYDVQQKDAhhdHJpY29yZTENMAsGA1UECwwEZGVtbzEXMBUGA1UEAwwOam9zc28tcHJvdmlkZXIxIzAhBgkqhkiG9w0BCQEWFHN1cHBvcnRAYXRyaWNvcmUuY29tMB4XDTE2MDIwMjE3MDIwM1oXDTI2MDEzMDE3MDIwM1owgZIxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzERMA8GA1UECgwIYXRyaWNvcmUxDTALBgNVBAsMBGRlbW8xFzAVBgNVBAMMDmpvc3NvLXByb3ZpZGVyMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGF0cmljb3JlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKCBJiMEjYh2Id50qMGGuZzivqFy7t3IwsJgjbS+xV3Jf5MmPyXh1AsYpk8eKSYDb+H8+hROeqxbSneXjAi5msrD+oCJnMwz0/uMUPsmntjlrbWSe2P2vGfLWLp708YLh2RyAA3Iz2Vx5fdbN+14zPfdMF/uNuD4e8XTU7PJcX4cIPna58P1ko3mCMVoPFI2KLess/EafBvc5OBBmTo3KeQ59hGRdNtCe5oeuLHapfLWnl36MHHkV/sdV+xVV/NsO5lVJ4al/n7snOsqBvUm++Zbey1OI3CWp9+q1CnnqFxzRiJySahYF5FoSiWJKpw7tXHkyU93FCVeBV5c5zxqVykCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAU27Ag+jrg+xVbRZc3Dqk40PitlvLiT619U8eyt0LHAhX+ZGy/Ao+pJAxSWHLP6YofG+EO3Fl4sgJ5S9py+PZwDgRQR1xfUsZ5a8tk6c0NPHpcHBU2pMuYQA+OoE7g5EIeAhPsmMeM2IH4Yz6qmzhvYBAvbDvGJYHi+Udxp8JHlKYjkieVw+9kI580YKeUIKXng4XXSuFHspYRLS2iDRfmeJsveOUYr9y7L4XrbLJIG/kVcpFiLkzsWJp1j6hwqPe748wekASae/+96l3NjT1AyNnD7rzyskUiNI6wb28OZeJoPczgzIedKXYdmFqLRuLeSLDJK2EiUATRUqE3ys7Fw==\u003c/ds:X509Certificate\u003e\n\t\t\t\t\t\t\t\t\u003c/ds:X509Data\u003e\n\t\t\t\t\t\t\t\u003c/ds:KeyInfo\u003e\n\t\t\t\t\t\t\u003c/ns6:KeyDescriptor\u003e\n\t\t\t\t\t\t\u003cns6:KeyDescriptor \n\t\t\t\t\t\t\tuse=\"encryption\"\u003e\n\t\t\t\t\t\t\u003cds:KeyInfo\u003e\n\t\t\t\t\t\t\t\u003cds:X509Data\u003e\n\t\t\t\t\t\t\t\t\u003cds:X509Certificate\u003eMIIDojCCAooCCQCVTd3p5WnWmjANBgkqhkiG9w0BAQsFADCBkjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2NvMREwDwYDVQQKDAhhdHJpY29yZTENMAsGA1UECwwEZGVtbzEXMBUGA1UEAwwOam9zc28tcHJvdmlkZXIxIzAhBgkqhkiG9w0BCQEWFHN1cHBvcnRAYXRyaWNvcmUuY29tMB4XDTE2MDIwMjE3MDIwM1oXDTI2MDEzMDE3MDIwM1owgZIxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzERMA8GA1UECgwIYXRyaWNvcmUxDTALBgNVBAsMBGRlbW8xFzAVBgNVBAMMDmpvc3NvLXByb3ZpZGVyMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGF0cmljb3JlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKCBJiMEjYh2Id50qMGGuZzivqFy7t3IwsJgjbS+xV3Jf5MmPyXh1AsYpk8eKSYDb+H8+hROeqxbSneXjAi5msrD+oCJnMwz0/uMUPsmntjlrbWSe2P2vGfLWLp708YLh2RyAA3Iz2Vx5fdbN+14zPfdMF/uNuD4e8XTU7PJcX4cIPna58P1ko3mCMVoPFI2KLess/EafBvc5OBBmTo3KeQ59hGRdNtCe5oeuLHapfLWnl36MHHkV/sdV+xVV/NsO5lVJ4al/n7snOsqBvUm++Zbey1OI3CWp9+q1CnnqFxzRiJySahYF5FoSiWJKpw7tXHkyU93FCVeBV5c5zxqVykCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAU27Ag+jrg+xVbRZc3Dqk40PitlvLiT619U8eyt0LHAhX+ZGy/Ao+pJAxSWHLP6YofG+EO3Fl4sgJ5S9py+PZwDgRQR1xfUsZ5a8tk6c0NPHpcHBU2pMuYQA+OoE7g5EIeAhPsmMeM2IH4Yz6qmzhvYBAvbDvGJYHi+Udxp8JHlKYjkieVw+9kI580YKeUIKXng4XXSuFHspYRLS2iDRfmeJsveOUYr9y7L4XrbLJIG/kVcpFiLkzsWJp1j6hwqPe748wekASae/+96l3NjT1AyNnD7rzyskUiNI6wb28OZeJoPczgzIedKXYdmFqLRuLeSLDJK2EiUATRUqE3ys7Fw==\u003c/ds:X509Certificate\u003e\n\t\t\t\t\t\t\t\u003c/ds:X509Data\u003e\n\t\t\t\t\t\t\u003c/ds:KeyInfo\u003e\n\t\t\t\t\t\u003cns6:EncryptionMethod \n\t\t\t\t\t\tAlgorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"\u003e\n\t\t\t\t\t\t\u003cenc:KeySize\u003e128\u003c/enc:KeySize\u003e\n\t\t\t\t\u003c/ns6:EncryptionMethod\u003e\n\t\t\t\t\u003cns6:EncryptionMethod \n\t\t\t\t\tAlgorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"\u003e\n\t\t\t\t\t\u003cenc:KeySize\u003e256\u003c/enc:KeySize\u003e\n\t\t\t\t\u003c/ns6:EncryptionMethod\u003e\n\t\t\t\t\u003cns6:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"\u003e\n\t\t\t\t\t\u003cenc:KeySize\u003e192\u003c/enc:KeySize\u003e\n\t\t\t\t\u003c/ns6:EncryptionMethod\u003e\n\t\t\t\u003c/ns6:KeyDescriptor\u003e\n\t\t\t\u003cns6:ArtifactResolutionService \n\t\t\t\tisDefault=\"true\" \n\t\t\t\tindex=\"0\" \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/ARTIFACT/SOAP\"\n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:SOAP\"/\u003e\n\t\t\t\u003cns6:SingleLogoutService \n\t\t\t\tResponseLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO_RESPONSE/POST\"\n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO/POST\"\n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"/\u003e\n\t\t\t\u003cns6:SingleLogoutService \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/SLO/SOAP\" \n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:SOAP\"/\u003e\n\t\t\t\u003cns6:ManageNameIDService \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/SOAP\" \n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:SOAP\"/\u003e\n\t\t\t\u003cns6:ManageNameIDService \n\t\t\t\tResponseLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI_RESPONSE/POST\" \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/POST\" \n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"/\u003e\n\t\t\t\u003cns6:ManageNameIDService \n\t\t\t\tResponseLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI_RESPONSE/REDIR\" \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/MNI/REDIR\" \n\t\t\t\tBinding=\"urn:oasis:n\t\t\t\tprotocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\" ames:tc:SAML:2.0:bindings:HTTP-Redirect\"/\u003e\n\t\t\t\u003cns6:NameIDFormat\u003eurn:oasis:names:tc:SAML:1.1:nameid-format:unspecified\u003c/ns6:NameIDFormat\u003e\n\t\t\t\u003cns6:AssertionConsumerService \n\t\t\t\tisDefault=\"true\" \n\t\t\t\tindex=\"0\" \n\t\t\t\tLocation=\"http://cloudsso:8082/IDBUS/ACCT-03C/SP-1/SAML2/ACS/POST\" \n\t\t\t\tBinding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"/\u003e\n\t\t\t\u003c/ns6:SPSSODescriptor\u003e\n\t\t\t\u003cns6:Organization\u003e\n\t\t\t\t\u003cns6:OrganizationName xml:lang=\"en\"\u003eAtricore IDBUs SAMLR2 JOSSO SP Sample\u003c/ns6:OrganizationName\u003e\n\t\t\t\t\u003cns6:OrganizationDisplayName xml:lang=\"en\"\u003eAtricore, Inc.\u003c/ns6:OrganizationDisplayName\u003e\n\t\t\t\t\u003cns6:OrganizationURL xml:lang=\"en\"\u003ehttp://www.atricore.org\u003c/ns6:OrganizationURL\u003e\n\t\t\t\u003c/ns6:Organization\u003e\n\t\t\t\u003cns6:ContactPerson contactType=\"other\"/\u003e\n\t\t\u003c/ns6:EntityDescriptor\u003e"]},"name":"Extsmal2-2"}}

2021/07/21 14:21:50 
HTTP/1.1 500 Internal Server Error
Content-Length: 66
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":"server","message":"See server logs for details"}
2021/07/21 14:21:50 CreateExtSaml2Sp. Error 500 Internal Server Error
    extsaml2sp_test.go:121: 500 Internal Server Error
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdP_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdP_crud
2021/07/21 14:21:50 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":1,"lastModification":"2021-07-21T14:21:50-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:50 getAppliance. 15 found for ID/name ida-a
2021/07/21 14:21:50 createIdP : idp-2 [ida-a]
2021/07/21 14:21:50 
POST /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 122
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","idp":{"@c":".IdentityProviderDTO","id":-1,"name":"idp-2","userDashboardBranding":"josso25-branding"}}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 2970
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"idp":{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_1D8C791D-F2AD-4CC3-95B9-5E3E470A4912","name":"idp-2","location":{"@id":3,"id":0,"elementId":"id5ACBF3F79CA56A","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"id1A34003EBFCFDF","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":null,"x":3500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"id":0,"elementId":"id2DF7F522DA992B","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":8,"name":"default","description":"default"},"subjectAuthnPolicies":null}}
2021/07/21 14:21:50 getIdp. idp-2 [ida-a]
2021/07/21 14:21:50 
GET /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 35
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"idp-2"}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 2984
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"idp":{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_1D8C791D-F2AD-4CC3-95B9-5E3E470A4912","name":"idp-2","location":{"@id":3,"id":0,"elementId":"id5ACBF3F79CA56A","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"id1A34003EBFCFDF","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":null,"x":3500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"id":0,"elementId":"id2DF7F522DA992B","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":8,"name":"default","description":"default"},"subjectAuthnPolicies":null},"config":null}
2021/07/21 14:21:50 getIdP. %!d(string=idp-2) found for ID/name idp-2
2021/07/21 14:21:50 updateIdP. : idp-2 [ida-a]
2021/07/21 14:21:50 
PUT /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 2272
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","idp":{"@c":".IdentityProviderDTO","activeBindings":[],"activeProfiles":[],"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"displayName":"idp-2-basic-authn","elementId":"id2DF7F522DA992B","enabled":true,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","id":0,"ignorePasswordCase":false,"ignoreUsernameCase":false,"impersonateUserPolicy":null,"name":"idp-2-basic-authn","priority":0,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0,"y":0}],"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"elementId":"id1A34003EBFCFDF","encrypter":null,"id":0,"signer":null,"useSampleStore":true,"useSystemStore":false},"dashboardUrl":"12345","description":"Updated description","destroyPreviousSession":false,"displayName":"Atricore","elementId":"_1D8C791D-F2AD-4CC3-95B9-5E3E470A4912","enableMetadataEndpoint":false,"encryptAssertion":false,"errorBinding":"ARTIFACT","externallyHostedIdentityConfirmationTokenService":false,"federatedConnectionsA":[],"federatedConnectionsB":[],"id":0,"identityConfirmationEnabled":false,"ignoreRequestedNameIDPolicy":true,"isRemote":false,"location":{"@id":3,"context":"IDBUS","elementId":"id5ACBF3F79CA56A","host":"localhost","id":0,"port":80,"protocol":"http","uri":"IDA-T/IDP-2"},"maxSessionsPerUser":-1,"messageTtl":300,"messageTtlTolerance":300,"name":"idp-2","oauth2Clients":[],"oauth2Enabled":false,"oauth2RememberMeTokenValidity":0,"oauth2TokenValidity":0,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"oidcIdTokenTimeToLive":3600,"openIdEnabled":false,"pwdlessAuthnEnabled":false,"role":"SSOIdentityProvider","sessionManagerFactory":{"@id":8,"description":"default","name":"default"},"signRequests":false,"signatureHash":"SHA256","ssoSessionTimeout":30,"subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"descriptionKey":"samlr2.principal","id":"samlr2-unspecified-nameidpolicy","name":"Principal","type":"PRINCIPAL"},"userDashboardBranding":"josso25-branding","wantAuthnRequestsSigned":false,"wantSignedRequests":false,"x":3500,"y":350}}

2021/07/21 14:21:50 
HTTP/1.1 200 OK
Content-Length: 2990
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:50 GMT

{"@id":1,"error":null,"validationErrors":[],"idp":{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_1D8C791D-F2AD-4CC3-95B9-5E3E470A4912","name":"idp-2","location":{"@id":3,"id":0,"elementId":"id5ACBF3F79CA56A","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":"Updated description","displayName":"Atricore","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"id1A34003EBFCFDF","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":null,"x":3500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"id":0,"elementId":"id2DF7F522DA992B","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":8,"name":"default","description":"default"},"subjectAuthnPolicies":null}}
2021/07/21 14:21:50 deleteIdp. idp-2 [ida-a]
2021/07/21 14:21:50 
DELETE /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 35
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"idp-2"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:51 deleteIdp. Deleted idp-2 : true
2021/07/21 14:21:51 get idps: all [ida-a]
2021/07/21 14:21:51 
GET /atricore-rest/services/iam-deploy/idps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 54
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idps":[]}
2021/07/21 14:21:51 createIdP : idp-1 [ida-a]
2021/07/21 14:21:51 
POST /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 79
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","idp":{"@c":".IdentityProviderDTO","id":-1,"name":"idp-1"}}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 2970
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idp":{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_673D52D3-FDFA-4548-9614-296CBE0FB3A8","name":"idp-1","location":{"@id":3,"id":0,"elementId":"id58023516018FFB","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-1","locationAsString":null},"description":null,"displayName":"idp-1","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"id3DD81462328458","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":null,"x":3750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"id":0,"elementId":"id9EE8F12AA3325E","name":"idp-1-basic-authn","displayName":"idp-1-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":8,"name":"default","description":"default"},"subjectAuthnPolicies":null}}
2021/07/21 14:21:51 createIdP : idp-2 [ida-a]
2021/07/21 14:21:51 
POST /atricore-rest/services/iam-deploy/idp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 79
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","idp":{"@c":".IdentityProviderDTO","id":-1,"name":"idp-2"}}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 2970
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idp":{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_80ECFB62-E102-4306-A6E0-2EAA6AFBC04F","name":"idp-2","location":{"@id":3,"id":0,"elementId":"id54F45F61BDE928","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"idFC19FC5AD6432B","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":null,"x":4000.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":5,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":6,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":7,"id":0,"elementId":"id61E051D7E21558","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":8,"name":"default","description":"default"},"subjectAuthnPolicies":null}}
2021/07/21 14:21:51 get idps: all [ida-a]
2021/07/21 14:21:51 
GET /atricore-rest/services/iam-deploy/idps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 13379
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idps":[{"@c":".IdentityProviderDTO","@id":2,"id":0,"elementId":"_673D52D3-FDFA-4548-9614-296CBE0FB3A8","name":"idp-1","location":{"@id":3,"id":0,"elementId":"id58023516018FFB","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-1","locationAsString":null},"description":null,"displayName":"idp-1","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":4,"id":0,"elementId":"id3DD81462328458","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":{"@id":5,"id":0,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":6,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":6,"lastModification":"2021-07-21T14:21:51-03:00","activeFeatures":[],"supportedRoles":[],"providers":[2,{"@c":".IdentityProviderDTO","@id":7,"id":0,"elementId":"_80ECFB62-E102-4306-A6E0-2EAA6AFBC04F","name":"idp-2","location":{"@id":8,"id":0,"elementId":"id54F45F61BDE928","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":9,"id":0,"elementId":"idFC19FC5AD6432B","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":5,"x":4000.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":10,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":11,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":12,"id":0,"elementId":"id61E051D7E21558","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":13,"name":"default","description":"default"},"subjectAuthnPolicies":null}],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":14,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":15,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null},"x":3750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":16,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":17,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":18,"id":0,"elementId":"id9EE8F12AA3325E","name":"idp-1-basic-authn","displayName":"idp-1-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":19,"name":"default","description":"default"},"subjectAuthnPolicies":null},{"@c":".IdentityProviderDTO","@id":20,"id":0,"elementId":"_80ECFB62-E102-4306-A6E0-2EAA6AFBC04F","name":"idp-2","location":{"@id":21,"id":0,"elementId":"id54F45F61BDE928","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":22,"id":0,"elementId":"idFC19FC5AD6432B","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":{"@id":23,"id":0,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":24,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":6,"lastModification":"2021-07-21T14:21:51-03:00","activeFeatures":[],"supportedRoles":[],"providers":[{"@c":".IdentityProviderDTO","@id":25,"id":0,"elementId":"_673D52D3-FDFA-4548-9614-296CBE0FB3A8","name":"idp-1","location":{"@id":26,"id":0,"elementId":"id58023516018FFB","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-1","locationAsString":null},"description":null,"displayName":"idp-1","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":27,"id":0,"elementId":"id3DD81462328458","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":23,"x":3750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":28,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":29,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":30,"id":0,"elementId":"id9EE8F12AA3325E","name":"idp-1-basic-authn","displayName":"idp-1-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":31,"name":"default","description":"default"},"subjectAuthnPolicies":null},20],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":32,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":33,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null},"x":4000.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":34,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":35,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":36,"id":0,"elementId":"id61E051D7E21558","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":37,"name":"default","description":"default"},"subjectAuthnPolicies":null}]}
=== RUN   TestAccCliSuite/TestAccCliIdP_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdS_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdS_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdSourceLdap_crud
2021/07/21 14:21:51 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":6,"lastModification":"2021-07-21T14:21:51-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:51 getAppliance. 15 found for ID/name ida-a
2021/07/21 14:21:51 createIdSourceLdap : ids-2 [ida-a]
2021/07/21 14:21:51 
POST /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 626
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idSourceLdap":{"@c":".LdapIdentitySourceDTO","elementId":"air","id":-1,"initialContextFactory":"true","ldapSearchScope":"subtree","name":"ids-2","principalUidAttributeID":"sAMAccountName","providerUrl":"ldap://192.168.0.97:389","referrals":"follow","roleAttributeID":"sAMAccountName","roleMatchingMode":"manager","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","securityAuthentication":"authenticated","securityCredential":"@WSX3edc","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","uidAttributeID":"member","userPropertiesQueryString":"space","usersCtxDN":"CN=Users,DC=mycompany,DC=com"},"idaName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 862
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdap":{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"idDA61329408C619","name":"ids-2","description":null,"x":800.0,"y":450.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}}
2021/07/21 14:21:51 getIdSourceLdap. ids-2 [ida-a]
2021/07/21 14:21:51 
GET /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 35
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"ids-2"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 862
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdap":{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"idDA61329408C619","name":"ids-2","description":null,"x":800.0,"y":450.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}}
2021/07/21 14:21:51 getIdSourceLdap. %!d(string=ids-2) found for ID/name ids-2
2021/07/21 14:21:51 updateIdSourceLdap. : ids-2 [ida-a]
2021/07/21 14:21:51 
PUT /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 775
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idSourceLdap":{"@c":".LdapIdentitySourceDTO","credentialQueryString":"N/A","description":"Updated description","elementId":"dirt","id":0,"includeOperationalAttributes":false,"initialContextFactory":"true","ldapSearchScope":"subtree","name":"ids-2","principalUidAttributeID":"sAMAccountName","providerUrl":"ldap://192.168.0.97:389","referrals":"follow","roleAttributeID":"sAMAccountName","roleMatchingMode":"manager","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","securityAuthentication":"authenticated","securityCredential":"@WSX3edc","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","uidAttributeID":"member","updatePasswordEnabled":false,"userPropertiesQueryString":"space","usersCtxDN":"CN=Users,DC=mycompany,DC=com","x":800,"y":450},"idaName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 862
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdap":{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"idDA61329408C619","name":"ids-2","description":null,"x":800.0,"y":450.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}}
2021/07/21 14:21:51 deleteIdSourceLdap. ids-2 [ida-a]
2021/07/21 14:21:51 
DELETE /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 35
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"ids-2"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:51 deleteIdSourceLdap. Deleted ids-2 : true
2021/07/21 14:21:51 get idSourceLdaps: all [ida-a]
2021/07/21 14:21:51 
GET /atricore-rest/services/iam-deploy/idsourceldaps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:51 
HTTP/1.1 200 OK
Content-Length: 63
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:51 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdaps":[]}
2021/07/21 14:21:51 createIdSourceLdap : ids-1 [ida-a]
2021/07/21 14:21:51 
POST /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 626
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idSourceLdap":{"@c":".LdapIdentitySourceDTO","elementId":"air","id":-1,"initialContextFactory":"true","ldapSearchScope":"subtree","name":"ids-1","principalUidAttributeID":"sAMAccountName","providerUrl":"ldap://192.168.0.97:389","referrals":"follow","roleAttributeID":"sAMAccountName","roleMatchingMode":"manager","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","securityAuthentication":"authenticated","securityCredential":"@WSX3edc","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","uidAttributeID":"member","userPropertiesQueryString":"space","usersCtxDN":"CN=Users,DC=mycompany,DC=com"},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 862
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdap":{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"id381FA5B7D585ED","name":"ids-1","description":null,"x":850.0,"y":500.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}}
2021/07/21 14:21:52 createIdSourceLdap : ids-2 [ida-a]
2021/07/21 14:21:52 
POST /atricore-rest/services/iam-deploy/idsourceldap HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 626
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idSourceLdap":{"@c":".LdapIdentitySourceDTO","elementId":"air","id":-1,"initialContextFactory":"true","ldapSearchScope":"subtree","name":"ids-2","principalUidAttributeID":"sAMAccountName","providerUrl":"ldap://192.168.0.97:389","referrals":"follow","roleAttributeID":"sAMAccountName","roleMatchingMode":"manager","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","securityAuthentication":"authenticated","securityCredential":"@WSX3edc","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","uidAttributeID":"member","userPropertiesQueryString":"space","usersCtxDN":"CN=Users,DC=mycompany,DC=com"},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 862
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdap":{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"id10FCA91DA18B3B","name":"ids-2","description":null,"x":900.0,"y":550.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}}
2021/07/21 14:21:52 get idSourceLdaps: all [ida-a]
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/idsourceldaps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 1668
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idSourceLdaps":[{"@c":".LdapIdentitySourceDTO","@id":2,"id":0,"elementId":"id381FA5B7D585ED","name":"ids-1","description":null,"x":850.0,"y":500.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false},{"@c":".LdapIdentitySourceDTO","@id":3,"id":0,"elementId":"id10FCA91DA18B3B","name":"ids-2","description":null,"x":900.0,"y":550.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false}]}
=== RUN   TestAccCliSuite/TestAccCliIdVault_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdVault_crud
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 873
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":11,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:52 getAppliance. 15 found for ID/name ida-a
2021/07/21 14:21:52 get idVaults: all [ida-a]
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/idvaults HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 58
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVaults":[]}
2021/07/21 14:21:52 createIdVault : idVault-2 [ida-a]
2021/07/21 14:21:52 
POST /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 136
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idVault":{"@c":".EmbeddedIdentityVaultDTO","id":-1,"identityConnectorName":"connector-default","name":"idVault-2"},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 254
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVault":{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"_FD699C34-33D1-49E3-BA09-FFB1E05144E2","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}}
2021/07/21 14:21:52 getIdVault. idVault-2 [ida-a]
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 39
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"idVault-2"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 254
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVault":{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"_FD699C34-33D1-49E3-BA09-FFB1E05144E2","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}}
2021/07/21 14:21:52 getIdVault. %!d(string=idVault-2) found for ID/name idVault-2
2021/07/21 14:21:52 updateIdVault. : idVault-2 [ida-a]
2021/07/21 14:21:52 
PUT /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 203
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idVault":{"@c":".EmbeddedIdentityVaultDTO","description":"Updated description","elementId":"12345","id":0,"identityConnectorName":"connector-default","name":"idVault-2","x":0,"y":0},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 222
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVault":{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"12345","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}}
2021/07/21 14:21:52 deleteIdVault. idVault-2 [ida-a]
2021/07/21 14:21:52 
DELETE /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 39
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"idVault-2"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:52 deleteIdVault. Deleted idVault-2 : true
2021/07/21 14:21:52 get idVaults: all [ida-a]
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/idvaults HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 58
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVaults":[]}
2021/07/21 14:21:52 createIdVault : idVault-1 [ida-a]
2021/07/21 14:21:52 
POST /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 92
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idVault":{"@c":".EmbeddedIdentityVaultDTO","id":-1,"name":"idVault-1"},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 254
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVault":{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"_728A5BFF-EE7B-4706-A00B-62A7FEBFF9AE","name":"idVault-1","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}}
2021/07/21 14:21:52 createIdVault : idVault-2 [ida-a]
2021/07/21 14:21:52 
POST /atricore-rest/services/iam-deploy/idvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 92
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idVault":{"@c":".EmbeddedIdentityVaultDTO","id":-1,"name":"idVault-2"},"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 254
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVault":{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"_2D70E5D9-7EBB-4E4A-BC11-F305E31F8B64","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}}
2021/07/21 14:21:52 get idVaults: all [ida-a]
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/idvaults HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 457
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"idVaults":[{"@c":".EmbeddedIdentityVaultDTO","@id":2,"id":0,"elementId":"_728A5BFF-EE7B-4706-A00B-62A7FEBFF9AE","name":"idVault-1","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"},{"@c":".EmbeddedIdentityVaultDTO","@id":3,"id":0,"elementId":"_2D70E5D9-7EBB-4E4A-BC11-F305E31F8B64","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}]}
=== RUN   TestAccCliSuite/TestAccCliIdVault_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdentityAppliance_basic
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 873
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":16,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:52 getAppliance. 15 found for ID/name ida-a
    identity_appliance_test.go:17: Appliance: jossoappi.IdentityApplianceDefinitionDTO{ActiveFeatures:(*[]string)(0xc0002efaa0), AuthenticationServices:(*[]jossoappi.AuthenticationServiceDTO)(0xc0002efbd8), Description:(*string)(0xc0004802e0), DisplayName:(*string)(0xc000480180), ElementId:(*string)(0xc000480160), ExecutionEnvironments:(*[]jossoappi.ExecutionEnvironmentDTO)(0xc0002efba8), Id:(*int64)(0xc0002df6a0), IdentitySources:(*[]jossoappi.IdentitySourceDTO)(0xc0002efb48), IdpSelector:(*jossoappi.EntitySelectionStrategyDTO)(0xc0002efc80), Keystore:(*jossoappi.KeystoreDTO)(nil), LastModification:(*time.Time)(0xc0002efa88), Location:(*jossoappi.LocationDTO)(0xc0001261e0), ModelVersion:(*string)(nil), Name:(*string)(0xc000480170), Namespace:(*string)(0xc0004802f0), Providers:(*[]jossoappi.ProviderDTO)(0xc0002efb18), RequiredBundles:(*[]string)(0xc0002efd10), Revision:(*int32)(0xc0002df778), SecurityConfig:(*jossoappi.IdentityApplianceSecurityConfigDTO)(nil), ServiceResources:(*[]jossoappi.ServiceResourceDTO)(0xc0002efb78), SupportedRoles:(*[]string)(0xc0002efae8), UserDashboardBranding:(*jossoappi.UserDashboardBrandingDTO)(0xc0002efc08), AdditionalProperties:map[string]interface {}{"@id":2}}
=== RUN   TestAccCliSuite/TestAccCliIdentityAppliance_crud
2021/07/21 14:21:52 createAppliance : ida-b com.atricore.idbus.ida.b
2021/07/21 14:21:52 
POST /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 191
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"appliance":{"description":"IDA-B TEST !","location":{"context":"IDBUS","host":"localhost","port":80,"protocol":"http","uri":"IDA-B"},"name":"ida-b","namespace":"com.atricore.idbus.ida.b"}}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":16,"elementId":"_2BDBC000-C4A8-4C1D-8DD0-DC71F6D2624E","name":"ida-b","displayName":"ida-b","location":{"@id":3,"id":0,"elementId":"id28DE725F7A4162","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-B","locationAsString":null},"description":"IDA-B TEST !","namespace":"com.atricore.idbus.ida.b","revision":1,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:52 createAppliance. ID: 16
2021/07/21 14:21:52 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-b"}

2021/07/21 14:21:52 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:52 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":16,"elementId":"_2BDBC000-C4A8-4C1D-8DD0-DC71F6D2624E","name":"ida-b","displayName":"ida-b","location":{"@id":3,"id":0,"elementId":"id28DE725F7A4162","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-B","locationAsString":null},"description":"IDA-B TEST !","namespace":"com.atricore.idbus.ida.b","revision":1,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:52 getAppliance. 16 found for ID/name ida-b
2021/07/21 14:21:52 updateAppliance : ida-b com.atricore.ida.a.mod
2021/07/21 14:21:52 
PUT /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 746
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"appliance":{"@id":2,"activeFeatures":[],"authenticationServices":[],"description":"IDA-B TEST !","displayName":"ida-b","elementId":"_2BDBC000-C4A8-4C1D-8DD0-DC71F6D2624E","executionEnvironments":[],"id":16,"identitySources":[],"idpSelector":{"@id":5,"description":"Requested, then Preferred","name":"requested-preferred-idp-selection"},"lastModification":"2021-07-21T14:21:52-03:00","location":{"@id":3,"context":"IDBUS","elementId":"id28DE725F7A4162","host":"localhost","id":0,"port":80,"protocol":"http","uri":"IDA-B"},"name":"ida-b","namespace":"com.atricore.ida.a.mod","providers":[],"requiredBundles":[],"revision":1,"serviceResources":[],"supportedRoles":[],"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"}}}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 870
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":16,"elementId":"_2BDBC000-C4A8-4C1D-8DD0-DC71F6D2624E","name":"ida-b","displayName":"ida-b","location":{"@id":3,"id":0,"elementId":"id28DE725F7A4162","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-B","locationAsString":null},"description":"IDA-B TEST !","namespace":"com.atricore.ida.a.mod","revision":2,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:53 updateAppliance. Updated: 16
2021/07/21 14:21:53 deleteAppliance id: 16
2021/07/21 14:21:53 
DELETE /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 14
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"name":"16"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:53 deleteAppliance. Deleted 16 : true
=== RUN   TestAccCliSuite/TestAccCliOidcRp_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliOidcRp_crud
2021/07/21 14:21:53 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 873
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":16,"lastModification":"2021-07-21T14:21:52-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/07/21 14:21:53 getAppliance. 15 found for ID/name ida-a
2021/07/21 14:21:53 createOidcRp : rp-2 [ida-a]
2021/07/21 14:21:53 
POST /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 99
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","id":-1,"name":"rp-2"}}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 933
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"id7247D3EF15504C","name":"rp-2","location":{"@id":3,"id":0,"elementId":"id70BFF9CA5BAD97","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":null,"displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":null,"x":4250.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}}
2021/07/21 14:21:53 getOidcRp. rp-2 [ida-a]
2021/07/21 14:21:53 
GET /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 34
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"rp-2"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 933
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"id7247D3EF15504C","name":"rp-2","location":{"@id":3,"id":0,"elementId":"id70BFF9CA5BAD97","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":null,"displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":null,"x":4250.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}}
2021/07/21 14:21:53 getOidcRp. %!d(string=rp-2) found for ID/name rp-2
2021/07/21 14:21:53 updateOidcRp. : rp-2 [ida-a]
2021/07/21 14:21:53 
PUT /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 602
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","activeBindings":[],"activeProfiles":[],"authorizedURIs":[],"clientId":"1234","clientType":"type1","description":"My updated description","displayName":"rp-2","elementId":"id7247D3EF15504C","federatedConnectionsA":[],"federatedConnectionsB":[],"grants":[],"id":0,"identityLookups":[],"isRemote":false,"location":{"@id":3,"context":"IDBUS","elementId":"id70BFF9CA5BAD97","host":"localhost","id":0,"port":80,"protocol":"http","uri":"IDA-T/RP-2"},"name":"rp-2","postLogoutRedirectionURIs":[],"responseTypes":[],"x":4250,"y":350}}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 955
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"id7247D3EF15504C","name":"rp-2","location":{"@id":3,"id":0,"elementId":"id70BFF9CA5BAD97","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":"My updated description","displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":null,"x":4250.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":"1234","clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}}
2021/07/21 14:21:53 deleteOidcRp. rp-2 [ida-a]
2021/07/21 14:21:53 
DELETE /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 34
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"rp-2"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:53 deleteOidcRp. Deleted rp-2 : true
2021/07/21 14:21:53 get oidcRps: all [ida-a]
2021/07/21 14:21:53 
GET /atricore-rest/services/iam-deploy/oidcrps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 57
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRps":[]}
2021/07/21 14:21:53 createOidcRp : rp-1 [ida-a]
2021/07/21 14:21:53 
POST /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 99
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","id":-1,"name":"rp-1"}}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 933
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"idB355DB07EE986F","name":"rp-1","location":{"@id":3,"id":0,"elementId":"id7999460521E716","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-1","locationAsString":null},"description":null,"displayName":"rp-1","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":null,"x":4500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}}
2021/07/21 14:21:53 createOidcRp : rp-2 [ida-a]
2021/07/21 14:21:53 
POST /atricore-rest/services/iam-deploy/oidcrp HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 99
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","id":-1,"name":"rp-2"}}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 933
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRp":{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"id22B55BDD8D93D0","name":"rp-2","location":{"@id":3,"id":0,"elementId":"id108CAAA33716D0","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":null,"displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":null,"x":4750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}}
2021/07/21 14:21:53 get oidcRps: all [ida-a]
2021/07/21 14:21:53 
GET /atricore-rest/services/iam-deploy/oidcrps HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 20920
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"oidcRps":[{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":2,"id":0,"elementId":"idB355DB07EE986F","name":"rp-1","location":{"@id":3,"id":0,"elementId":"id7999460521E716","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-1","locationAsString":null},"description":null,"displayName":"rp-1","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":{"@id":4,"id":0,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":5,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":21,"lastModification":"2021-07-21T14:21:53-03:00","activeFeatures":[],"supportedRoles":[],"providers":[{"@c":".IdentityProviderDTO","@id":6,"id":0,"elementId":"_673D52D3-FDFA-4548-9614-296CBE0FB3A8","name":"idp-1","location":{"@id":7,"id":0,"elementId":"id58023516018FFB","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-1","locationAsString":null},"description":null,"displayName":"idp-1","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":8,"id":0,"elementId":"id3DD81462328458","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":4,"x":3750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":9,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":10,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":11,"id":0,"elementId":"id9EE8F12AA3325E","name":"idp-1-basic-authn","displayName":"idp-1-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":12,"name":"default","description":"default"},"subjectAuthnPolicies":null},{"@c":".IdentityProviderDTO","@id":13,"id":0,"elementId":"_80ECFB62-E102-4306-A6E0-2EAA6AFBC04F","name":"idp-2","location":{"@id":14,"id":0,"elementId":"id54F45F61BDE928","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":15,"id":0,"elementId":"idFC19FC5AD6432B","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":4,"x":4000.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":16,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":17,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":18,"id":0,"elementId":"id61E051D7E21558","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":19,"name":"default","description":"default"},"subjectAuthnPolicies":null},2,{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":20,"id":0,"elementId":"id22B55BDD8D93D0","name":"rp-2","location":{"@id":21,"id":0,"elementId":"id108CAAA33716D0","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":null,"displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":4,"x":4750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}],"identitySources":[{"@c":".LdapIdentitySourceDTO","@id":22,"id":0,"elementId":"id381FA5B7D585ED","name":"ids-1","description":null,"x":850.0,"y":500.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false},{"@c":".LdapIdentitySourceDTO","@id":23,"id":0,"elementId":"id10FCA91DA18B3B","name":"ids-2","description":null,"x":900.0,"y":550.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false},{"@c":".EmbeddedIdentityVaultDTO","@id":24,"id":0,"elementId":"_728A5BFF-EE7B-4706-A00B-62A7FEBFF9AE","name":"idVault-1","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"},{"@c":".EmbeddedIdentityVaultDTO","@id":25,"id":0,"elementId":"_2D70E5D9-7EBB-4E4A-BC11-F305E31F8B64","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":26,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":27,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null},"x":4500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]},{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":28,"id":0,"elementId":"id22B55BDD8D93D0","name":"rp-2","location":{"@id":29,"id":0,"elementId":"id108CAAA33716D0","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-2","locationAsString":null},"description":null,"displayName":"rp-2","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":{"@id":30,"id":0,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":31,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":21,"lastModification":"2021-07-21T14:21:53-03:00","activeFeatures":[],"supportedRoles":[],"providers":[{"@c":".IdentityProviderDTO","@id":32,"id":0,"elementId":"_673D52D3-FDFA-4548-9614-296CBE0FB3A8","name":"idp-1","location":{"@id":33,"id":0,"elementId":"id58023516018FFB","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-1","locationAsString":null},"description":null,"displayName":"idp-1","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":34,"id":0,"elementId":"id3DD81462328458","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":30,"x":3750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":35,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":36,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":37,"id":0,"elementId":"id9EE8F12AA3325E","name":"idp-1-basic-authn","displayName":"idp-1-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":38,"name":"default","description":"default"},"subjectAuthnPolicies":null},{"@c":".IdentityProviderDTO","@id":39,"id":0,"elementId":"_80ECFB62-E102-4306-A6E0-2EAA6AFBC04F","name":"idp-2","location":{"@id":40,"id":0,"elementId":"id54F45F61BDE928","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/IDP-2","locationAsString":null},"description":null,"displayName":"idp-2","isRemote":false,"config":{"@c":".SamlR2IDPConfigDTO","@id":41,"id":0,"elementId":"idFC19FC5AD6432B","name":null,"displayName":null,"description":null,"signer":null,"encrypter":null,"useSampleStore":true,"useSystemStore":false},"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":"SSOIdentityProvider","identityAppliance":30,"x":4000.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"wantAuthnRequestsSigned":false,"signRequests":false,"wantSignedRequests":false,"encryptAssertion":false,"encryptAssertionAlgorithm":null,"signatureHash":"SHA256","ignoreRequestedNameIDPolicy":true,"ssoSessionTimeout":30,"maxSessionsPerUser":-1,"destroyPreviousSession":false,"oauth2ClientsConfig":null,"oauth2Clients":[],"oauth2Key":null,"oauth2TokenValidity":0,"oauth2RememberMeTokenValidity":0,"pwdlessAuthnEnabled":false,"pwdlessAuthnFrom":null,"pwdlessAuthnSubject":null,"pwdlessAuthnTemplate":null,"pwdlessAuthnTo":null,"oauth2Enabled":false,"openIdEnabled":false,"oidcIdTokenTimeToLive":3600,"oidcAccessTokenTimeToLive":3600,"oidcAuthzCodeTimeToLive":300,"dashboardUrl":null,"errorBinding":"ARTIFACT","userDashboardBranding":"josso25-branding","subjectNameIDPolicy":{"@c":".SubjectNameIdentifierPolicyDTO","@id":42,"id":"samlr2-unspecified-nameidpolicy","name":"Principal","descriptionKey":"samlr2.principal","type":"PRINCIPAL","subjectAttribute":null},"attributeProfile":{"@c":".BuiltInAttributeProfileDTO","@id":43,"id":940388578,"elementId":null,"name":"basic-built-in","profileType":"BASIC"},"authenticationMechanisms":[{"@c":".BasicAuthenticationDTO","@id":44,"id":0,"elementId":"id61E051D7E21558","name":"idp-2-basic-authn","displayName":"idp-2-basic-authn","priority":0,"delegatedAuthentication":null,"hashAlgorithm":"SHA-256","hashEncoding":"BASE64","ignoreUsernameCase":false,"ignorePasswordCase":false,"saltLength":0,"saltPrefix":null,"saltSuffix":null,"impersonateUserPolicy":null,"enabled":true,"simpleAuthnSaml2AuthnCtxClass":"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport","x":0.0,"y":0.0}],"authenticationContract":null,"emissionPolicy":null,"delegatedAuthentications":null,"messageTtl":300,"messageTtlTolerance":300,"identityConfirmationEnabled":false,"identityConfirmationPolicy":null,"identityConfirmationOAuth2ClientId":null,"identityConfirmationOAuth2ClientSecret":null,"externallyHostedIdentityConfirmationTokenService":false,"identityConfirmationOAuth2AuthorizationServerEndpoint":null,"enableMetadataEndpoint":false,"sessionManagerFactory":{"@id":45,"name":"default","description":"default"},"subjectAuthnPolicies":null},{"@c":".ExternalOpenIDConnectRelayingPartyDTO","@id":46,"id":0,"elementId":"idB355DB07EE986F","name":"rp-1","location":{"@id":47,"id":0,"elementId":"id7999460521E716","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T/RP-1","locationAsString":null},"description":null,"displayName":"rp-1","isRemote":false,"config":null,"activeBindings":[],"activeProfiles":[],"identityLookups":[],"metadata":null,"role":null,"identityAppliance":30,"x":4500.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]},28],"identitySources":[{"@c":".LdapIdentitySourceDTO","@id":48,"id":0,"elementId":"id381FA5B7D585ED","name":"ids-1","description":null,"x":850.0,"y":500.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false},{"@c":".LdapIdentitySourceDTO","@id":49,"id":0,"elementId":"id10FCA91DA18B3B","name":"ids-2","description":null,"x":900.0,"y":550.0,"initialContextFactory":"true","providerUrl":"ldap://192.168.0.97:389","securityPrincipal":"CN=Administrator,CN=Users,DC=mycompany,DC=com","securityCredential":"@WSX3edc","securityAuthentication":"authenticated","ldapSearchScope":"subtree","usersCtxDN":"CN=Users,DC=mycompany,DC=com","principalUidAttributeID":"sAMAccountName","roleMatchingMode":"manager","uidAttributeID":"member","rolesCtxDN":"CN=Users,DC=mycompany,DC=com","roleAttributeID":"sAMAccountName","credentialQueryString":"N/A","updateableCredentialAttribute":null,"userPropertiesQueryString":"space","referrals":"follow","includeOperationalAttributes":false,"customClass":null,"updatePasswordEnabled":false},{"@c":".EmbeddedIdentityVaultDTO","@id":50,"id":0,"elementId":"_728A5BFF-EE7B-4706-A00B-62A7FEBFF9AE","name":"idVault-1","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"},{"@c":".EmbeddedIdentityVaultDTO","@id":51,"id":0,"elementId":"_2D70E5D9-7EBB-4E4A-BC11-F305E31F8B64","name":"idVault-2","description":null,"x":0.0,"y":0.0,"identityConnectorName":"connector-default"}],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":52,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":53,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null},"x":4750.0,"y":350.0,"federatedConnectionsA":[],"federatedConnectionsB":[],"clientId":null,"clientSecret":null,"clientCert":null,"clientType":null,"signingAlg":null,"encryptionAlg":null,"encryptionMethod":null,"idTokenSigningAlg":null,"idTokenEncryptionAlg":null,"idTokenEncryptionMethod":null,"postLogoutRedirectionURIs":[],"grants":[],"responseModes":[],"responseTypes":[],"clientAuthnMethod":null,"authorizedURIs":[]}]}
=== RUN   TestAccCliSuite/TestAccCliOidcRp_updateFailOnDupName
=== CONT  TestAccCliSuite
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/07/21 14:21:53 
GET /atricore-rest/services/iam-deploy/appliances HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 5
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

null

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 876
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"appliances":[{"@id":2,"id":15,"elementId":"_96810EC9-FFA9-4141-B038-3EABE83A3493","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id4F63283A60F44D","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":21,"lastModification":"2021-07-21T14:21:53-03:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}]}
2021/07/21 14:21:53 getAppliances. found appliances 1
    client_test.go:205: deleting appliance 15
2021/07/21 14:21:53 deleteAppliance id: 15
2021/07/21 14:21:53 
DELETE /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 14
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYyNjg5MTcxMCwianRpIjoiaWQtNjRiZTgyOWQtZjU0Zi00MTY4LWJlNDQtNWQ4ZGI3NDk4NzUxIn0.deyKzBZg-QL3485UUf-Ttd71V7CIWbv341dpby70tAM
Content-Type: application/json
Accept-Encoding: gzip

{"name":"15"}

2021/07/21 14:21:53 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Wed, 21 Jul 2021 17:21:53 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/07/21 14:21:53 deleteAppliance. Deleted 15 : true
--- FAIL: TestAccCliSuite (3.57s)
    --- FAIL: TestAccCliSuite/TestAccCliExtSaml2_crud (0.10s)
    --- PASS: TestAccCliSuite/TestAccCliExtSaml2_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliExtSaml2_crud_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdP_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdP_crud (1.07s)
    --- PASS: TestAccCliSuite/TestAccCliIdP_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdS_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdS_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdSourceLdap_crud (0.54s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_crud (0.68s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdentityAppliance_basic (0.01s)
    --- PASS: TestAccCliSuite/TestAccCliIdentityAppliance_crud (0.28s)
    --- PASS: TestAccCliSuite/TestAccCliOidcRp_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliOidcRp_crud (0.66s)
    --- PASS: TestAccCliSuite/TestAccCliOidcRp_updateFailOnDupName (0.00s)
=== RUN   TestMarshalJSON
--- PASS: TestMarshalJSON (0.00s)
=== RUN   TestStringToLocation
=== RUN   TestStringToLocation/StrToLoc-1
=== RUN   TestStringToLocation/StrToLoc-2
=== RUN   TestStringToLocation/StrToLoc-3
=== RUN   TestStringToLocation/StrToLoc-4
=== RUN   TestStringToLocation/StrToLoc-5
--- PASS: TestStringToLocation (0.00s)
    --- PASS: TestStringToLocation/StrToLoc-1 (0.00s)
    --- PASS: TestStringToLocation/StrToLoc-2 (0.00s)
    --- PASS: TestStringToLocation/StrToLoc-3 (0.00s)
    --- PASS: TestStringToLocation/StrToLoc-4 (0.00s)
    --- PASS: TestStringToLocation/StrToLoc-5 (0.00s)
=== RUN   TestLocationToString
=== RUN   TestLocationToString/LocToStr-1
=== RUN   TestLocationToString/LocToStr-2
=== RUN   TestLocationToString/LocToStr-3
=== RUN   TestLocationToString/LocToStr-4
=== RUN   TestLocationToString/LocToStr-5
=== RUN   TestLocationToString/LocToStr-6
--- PASS: TestLocationToString (0.00s)
    --- PASS: TestLocationToString/LocToStr-1 (0.00s)
    --- PASS: TestLocationToString/LocToStr-2 (0.00s)
    --- PASS: TestLocationToString/LocToStr-3 (0.00s)
    --- PASS: TestLocationToString/LocToStr-4 (0.00s)
    --- PASS: TestLocationToString/LocToStr-5 (0.00s)
    --- PASS: TestLocationToString/LocToStr-6 (0.00s)
FAIL
FAIL	github.com/atricore/josso-sdk-go	3.576s
FAIL
