go mod tidy
go mod vendor
==> Checking that code complies with gofmt requirements...
go clean -testcache
TF_ACC=1 go test $(go list ./... |grep -v 'vendor') -v  -run TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud -timeout 120m
=== RUN   TestAccCliSuite
    client_test.go:25: creating client
2021/09/06 14:58:27 newIdbusApiClient TRACE: true
2021/09/06 14:58:27 Using client TRACE ON
2021/09/06 14:58:27 registering server http://localhost:8081/atricore-rest/services
2021/09/06 14:58:27 adding server configuration for http://localhost:8081/atricore-rest/services
2021/09/06 14:58:27 authn: idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7 true/admin true
2021/09/06 14:58:27 
POST /atricore-rest/services/iam-authn/sign-on HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 127
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"clientId":"idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7","password":"atricore","secret":"7oUHlv(HLT%vxK4L","username":"admin"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 2206
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"server":null,"validationErrors":null,"idToken":"eyJhbGciOiJIUzI1NiJ9.eyJhdF9oYXNoIjoiSWp0cjFlN1RTeEJaMzJGdVplUWMtUSIsInN1YiI6ImFkbWluIiwiYXVkIjoiaWRidXMtZjJmNzI0NGUtYmJjZS00NGNhLThiMzMtZjVjMGJkZTMzOWY3IiwiYXV0aF90aW1lIjoxNjMwOTU0NzA3LCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsIm5hbWUiOiJBZG1pbmlzdHJhdG9yIiwiYWNjb3VudERpc2FibGVkIjoiZmFsc2UiLCJncm91cHMiOlsiQWRtaW5pc3RyYXRvcnMiXSwiZXhwIjoxNjMwOTU1MDA3LCJnaXZlbl9uYW1lIjoiQWRtaW5pc3RyYXRvciIsImlhdCI6MTYzMDk1NDcwNywiZmFtaWx5X25hbWUiOiJBZG1pbmlzdHJhdG9yIn0.AkSIO_4WyIphuHdcDzoSHWTEvBz_xyio-GSyiR_uBjI","accessToken":"eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA","refreshToken":"YUCWmWwcUWbDU6q_P0ltS9Ja_lB48yMBoLO2CI4Vgmk","authenticatedUser":{"@id":2,"id":"1","userName":"admin","firstName":"Administrator","surename":"Administrator","commonName":"Administrator","givenName":"Administrator","initials":null,"generationQualifier":null,"distinguishedName":null,"email":null,"telephoneNumber":null,"facsimilTelephoneNumber":null,"countryName":"Administrator","localityName":null,"stateOrProvinceName":null,"streetAddress":null,"organizationName":null,"organizationUnitName":null,"personalTitle":null,"businessCategory":null,"postalAddress":null,"postalCode":null,"postOfficeBox":null,"language":null,"groups":[{"@id":3,"id":"1","name":"Administrators","description":"Administrators","extraAttributes":[]}],"accountDisabled":false,"accountExpires":null,"accountExpirationDate":null,"limitSimultaneousLogin":null,"maximunLogins":null,"terminatePreviousSession":null,"preventNewSession":null,"allowUserToChangePassword":true,"forcePeriodicPasswordChanges":null,"daysBetweenChanges":null,"passwordExpirationDate":null,"notifyPasswordExpiration":null,"daysBeforeExpiration":null,"userPassword":"YHRsG4tlSmYpl6//r8OoIafxuqwne0plJ7HuzRZbtYY=","userCertificate":null,"automaticallyGeneratePassword":null,"emailNewPasword":null,"extraAttributes":[]}}
    client_test.go:33: created test client: [{http://localhost:8081/atricore-rest/services JOSSO Test server map[]}]
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/appliances HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 5
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

null

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 1054
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"appliances":[{"@id":2,"id":6,"elementId":"_B5D518B5-C7C7-471B-B847-863BE4C58C44","name":"ida-1","displayName":null,"location":{"@id":3,"id":0,"elementId":"id682ED029831CA9","protocol":"http","host":"localhost","port":8081,"context":"IDBUS","uri":"IDA-1","locationAsString":null},"description":null,"namespace":"com.atricore.idmn.authoring.ida1","revision":2,"lastModification":"2021-09-02T17:13:26-04:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":null,"securityConfig":{"@id":5,"externalConfig":false,"externalConfigFile":null,"encryptSensitiveData":false,"encryption":null,"encryptionConfig":null,"encryptionPassword":null,"salt":null,"saltValue":null,"encryptionConfigFile":null,"passwordProperty":null,"saltProperty":null},"requiredBundles":[],"modelVersion":"3.0.3-SNAPSHOT-unstable"}]}
2021/09/06 14:58:27 getAppliances. found appliances 1
=== RUN   TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 21
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"idOrName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 61
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":null}
2021/09/06 14:58:27 getAppliance. not found for ID/name ida-a
2021/09/06 14:58:27 createAppliance : ida-a com.atricore.idbus.ida.t
2021/09/06 14:58:27 
POST /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 191
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"appliance":{"description":"IDA-T TEST !","location":{"context":"IDBUS","host":"localhost","port":80,"protocol":"http","uri":"IDA-T"},"name":"ida-a","namespace":"com.atricore.idbus.ida.t"}}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 872
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"appliance":{"@id":2,"id":55,"elementId":"_C0DC8687-4237-44EA-B050-0E6763BA82F2","name":"ida-a","displayName":"ida-a","location":{"@id":3,"id":0,"elementId":"id9C28D7173C28B0","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":1,"lastModification":"2021-09-06T14:58:27-04:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":5,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}}
2021/09/06 14:58:27 createAppliance. ID: 55
2021/09/06 14:58:27 CreateDbIdentityVault : DdIdentityVault-a [ida-a]
2021/09/06 14:58:27 
POST /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 551
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"dbIdVault":{"@c":".DbIdentityVaultDTO","acquireIncrement":1,"connectionUrl":"jdbc:mysql:localhost/DdIdentityVault-a?create=true","description":"DescriptionDdIdentityVault-a","driverName":"org.mysql.driver\n","externalDB":true,"hashAlgorithm":"SHA267","hashEncoding":"BASE64","idleConnectionTestPeriod":1,"initialPoolSize":10,"maxIdleTime":15,"maxPoolSize":20,"minPoolSize":1,"name":"DdIdentityVault-a","password":"pdwDdIdentityVault-a","pooledDatasource":true,"saltLength":55,"saltValue":"salt#","username":"dbDdIdentityVault-a"},"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 658
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVault":{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_81A4A4E3-EBF3-4FBB-B5B8-71EDF7CE4108","name":"DdIdentityVault-a","description":"DescriptionDdIdentityVault-a","x":0.0,"y":0.0,"username":"dbDdIdentityVault-a","password":"pdwDdIdentityVault-a","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DdIdentityVault-a?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15}}
2021/09/06 14:58:27 GetDbIdentityVaultDto. DdIdentityVault-a [ida-a]
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 47
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"DdIdentityVault-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 672
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVault":{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_81A4A4E3-EBF3-4FBB-B5B8-71EDF7CE4108","name":"DdIdentityVault-a","description":"DescriptionDdIdentityVault-a","x":0.0,"y":0.0,"username":"dbDdIdentityVault-a","password":"pdwDdIdentityVault-a","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DdIdentityVault-a?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15},"config":null}
2021/09/06 14:58:27 GetDbIdentityVaultDto. DdIdentityVault-a found for ID/name DdIdentityVault-a
2021/09/06 14:58:27 UpdateDbIdentityVaultDto. : DdIdentityVault-a [ida-a]
2021/09/06 14:58:27 
PUT /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 613
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"dbIdVault":{"@c":".DbIdentityVaultDTO","acquireIncrement":1,"connectionUrl":"jdbc:mysql:localhost/DdIdentityVault-a?create=true","description":"Updated description","driverName":"org.mysql.driver\n","elementId":"_81A4A4E3-EBF3-4FBB-B5B8-71EDF7CE4108","externalDB":true,"hashAlgorithm":"SHA267","hashEncoding":"BASE64","id":0,"idleConnectionTestPeriod":1,"initialPoolSize":10,"maxIdleTime":15,"maxPoolSize":20,"minPoolSize":1,"name":"DdIdentityVault-a","password":"pdwDdIdentityVault-a","pooledDatasource":true,"saltLength":55,"saltValue":"salt#","username":"dbDdIdentityVault-a","x":0,"y":0},"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 649
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVault":{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_81A4A4E3-EBF3-4FBB-B5B8-71EDF7CE4108","name":"DdIdentityVault-a","description":"Updated description","x":0.0,"y":0.0,"username":"dbDdIdentityVault-a","password":"pdwDdIdentityVault-a","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DdIdentityVault-a?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15}}
2021/09/06 14:58:27 deleteDbIdentityVaultDto. DdIdentityVault-a [ida-a]
2021/09/06 14:58:27 
DELETE /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 47
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a","name":"DdIdentityVault-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/09/06 14:58:27 deleteIntSaml2Ss. Deleted DdIdentityVault-a : true
2021/09/06 14:58:27 get DbIdentityVaultDtos: all [ida-a]
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/dbidvaults HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 60
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVaults":[]}
2021/09/06 14:58:27 CreateDbIdentityVault : DbIdentityVault-2 [ida-a]
2021/09/06 14:58:27 
POST /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 551
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"dbIdVault":{"@c":".DbIdentityVaultDTO","acquireIncrement":1,"connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","description":"DescriptionDbIdentityVault-2","driverName":"org.mysql.driver\n","externalDB":true,"hashAlgorithm":"SHA267","hashEncoding":"BASE64","idleConnectionTestPeriod":1,"initialPoolSize":10,"maxIdleTime":15,"maxPoolSize":20,"minPoolSize":1,"name":"DbIdentityVault-2","password":"pdwDbIdentityVault-2","pooledDatasource":true,"saltLength":55,"saltValue":"salt#","username":"dbDbIdentityVault-2"},"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 658
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVault":{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_35E39FCA-59B9-48F0-90EE-BD60425476E0","name":"DbIdentityVault-2","description":"DescriptionDbIdentityVault-2","x":0.0,"y":0.0,"username":"dbDbIdentityVault-2","password":"pdwDbIdentityVault-2","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15}}
2021/09/06 14:58:27 CreateDbIdentityVault : DbIdentityVault-2 [ida-a]
2021/09/06 14:58:27 
POST /atricore-rest/services/iam-deploy/dbidvault HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 551
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"dbIdVault":{"@c":".DbIdentityVaultDTO","acquireIncrement":1,"connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","description":"DescriptionDbIdentityVault-2","driverName":"org.mysql.driver\n","externalDB":true,"hashAlgorithm":"SHA267","hashEncoding":"BASE64","idleConnectionTestPeriod":1,"initialPoolSize":10,"maxIdleTime":15,"maxPoolSize":20,"minPoolSize":1,"name":"DbIdentityVault-2","password":"pdwDbIdentityVault-2","pooledDatasource":true,"saltLength":55,"saltValue":"salt#","username":"dbDbIdentityVault-2"},"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 658
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVault":{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_35E39FCA-59B9-48F0-90EE-BD60425476E0","name":"DbIdentityVault-2","description":"DescriptionDbIdentityVault-2","x":0.0,"y":0.0,"username":"dbDbIdentityVault-2","password":"pdwDbIdentityVault-2","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15}}
2021/09/06 14:58:27 get DbIdentityVaultDtos: all [ida-a]
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/dbidvaults HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 20
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"idaName":"ida-a"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 1263
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"dbIdVaults":[{"@c":".DbIdentityVaultDTO","@id":2,"id":0,"elementId":"_35E39FCA-59B9-48F0-90EE-BD60425476E0","name":"DbIdentityVault-2","description":"DescriptionDbIdentityVault-2","x":0.0,"y":0.0,"username":"dbDbIdentityVault-2","password":"pdwDbIdentityVault-2","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15},{"@c":".DbIdentityVaultDTO","@id":3,"id":0,"elementId":"_07F38811-5E6C-4E0C-9148-0A98C633C60B","name":"DbIdentityVault-2","description":"DescriptionDbIdentityVault-2","x":0.0,"y":0.0,"username":"dbDbIdentityVault-2","password":"pdwDbIdentityVault-2","externalDB":true,"driverName":"org.mysql.driver\n","connectionUrl":"jdbc:mysql:localhost/DbIdentityVault-2?create=true","hashAlgorithm":"SHA267","hashEncoding":"BASE64","saltLength":55,"saltValue":"salt#","pooledDatasource":true,"acquireIncrement":1,"initialPoolSize":10,"minPoolSize":1,"maxPoolSize":20,"idleConnectionTestPeriod":1,"maxIdleTime":15}]}
    dbIdentity_vault_test.go:126: 1 error occurred:
        	* invalid elementid, expected [_35E39FCA-59B9-48F0-90EE-BD60425476E0],  received[_07F38811-5E6C-4E0C-9148-0A98C633C60B]
        
        
=== RUN   TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName
=== CONT  TestAccCliSuite
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/09/06 14:58:27 
GET /atricore-rest/services/iam-deploy/appliances HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 5
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

null

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 1870
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"appliances":[{"@id":2,"id":6,"elementId":"_B5D518B5-C7C7-471B-B847-863BE4C58C44","name":"ida-1","displayName":null,"location":{"@id":3,"id":0,"elementId":"id682ED029831CA9","protocol":"http","host":"localhost","port":8081,"context":"IDBUS","uri":"IDA-1","locationAsString":null},"description":null,"namespace":"com.atricore.idmn.authoring.ida1","revision":2,"lastModification":"2021-09-02T17:13:26-04:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":4,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":null,"securityConfig":{"@id":5,"externalConfig":false,"externalConfigFile":null,"encryptSensitiveData":false,"encryption":null,"encryptionConfig":null,"encryptionPassword":null,"salt":null,"saltValue":null,"encryptionConfigFile":null,"passwordProperty":null,"saltProperty":null},"requiredBundles":[],"modelVersion":"3.0.3-SNAPSHOT-unstable"},{"@id":6,"id":55,"elementId":"_C0DC8687-4237-44EA-B050-0E6763BA82F2","name":"ida-a","displayName":"ida-a","location":{"@id":7,"id":0,"elementId":"id9C28D7173C28B0","protocol":"http","host":"localhost","port":80,"context":"IDBUS","uri":"IDA-T","locationAsString":null},"description":"IDA-T TEST !","namespace":"com.atricore.idbus.ida.t","revision":6,"lastModification":"2021-09-06T14:58:27-04:00","activeFeatures":[],"supportedRoles":[],"providers":[],"identitySources":[],"serviceResources":[],"executionEnvironments":[],"authenticationServices":[],"keystore":null,"userDashboardBranding":{"@id":8,"id":"josso25-branding","name":"JOSSO 2.5+"},"idpSelector":{"@id":9,"name":"requested-preferred-idp-selection","description":"Requested, then Preferred"},"securityConfig":null,"requiredBundles":[],"modelVersion":null}]}
2021/09/06 14:58:27 getAppliances. found appliances 2
    client_test.go:205: deleting appliance 55
2021/09/06 14:58:27 deleteAppliance id: 55
2021/09/06 14:58:27 
DELETE /atricore-rest/services/iam-deploy/appliance HTTP/1.1
Host: localhost:8081
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 14
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpZGJ1cy1mMmY3MjQ0ZS1iYmNlLTQ0Y2EtOGIzMy1mNWMwYmRlMzM5ZjciLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MVwvSURCVVNcL0RFRkFVTFRcL1dCLU9QXC9PSURDXC9NRCIsImV4cCI6MTYzMDk1ODMwNywianRpIjoiaWQtYTIzMTdiMTEtMTViOC00MWQxLTg1YzAtMDI2ZjMxZDVhMzg0In0.MCLrvJenruS8jup4O2Ta_VRHPJdgoSV1o_n4Ar2earA
Content-Type: application/json
Accept-Encoding: gzip

{"name":"55"}

2021/09/06 14:58:27 
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json
Date: Mon, 06 Sep 2021 18:58:27 GMT

{"@id":1,"error":null,"validationErrors":[],"removed":true}
2021/09/06 14:58:27 deleteAppliance. Deleted 55 : true
--- FAIL: TestAccCliSuite (0.55s)
    --- FAIL: TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud (0.39s)
    --- PASS: TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName (0.00s)
FAIL
FAIL	github.com/atricore/josso-sdk-go	0.556s
FAIL
