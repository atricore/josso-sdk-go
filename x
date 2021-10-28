go mod tidy
go mod vendor
==> Checking that code complies with gofmt requirements...
go test $(go list ./... |grep -v 'vendor') || exit 1
2021/10/04 16:05:56 newIdbusApiClient TRACE: false
2021/10/04 16:05:56 registering server http://localhost:8081/atricore-rest/services
2021/10/04 16:05:56 adding server configuration for http://localhost:8081/atricore-rest/services
2021/10/04 16:05:56 authn: idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7 true/admin true
2021/10/04 16:05:56 getAppliances. found appliances 0
2021/10/04 16:05:56 getAppliance. not found for ID/name testacc-a
2021/10/04 16:05:56 createAppliance : testacc-a com.atricore.idbus.ida.t
2021/10/04 16:05:56 createAppliance. ID: 174
2021/10/04 16:05:56 CreateDbIdentitySource : DdIdentityVauld-a [testacc-a]
2021/10/04 16:05:56 GetDbIdentitySourceDTO. DdIdentityVauld-a [testacc-a]
2021/10/04 16:05:56 GetDbIdentitySourceDTO. DdIdentityVauld-a found for ID/name DdIdentityVauld-a
2021/10/04 16:05:56 UpdateDbIdentitySourceDTO. : DdIdentityVauld-a [testacc-a]
2021/10/04 16:05:56 deleteDbIdentitySourceDTO. DdIdentityVauld-a [testacc-a]
2021/10/04 16:05:56 deleteIntSaml2Ss. Deleted DdIdentityVauld-a : true
2021/10/04 16:05:56 get DbIdentitySourceDTOs: all [testacc-a]
2021/10/04 16:05:56 CreateDbIdentitySource : DbIdentitySource-2 [testacc-a]
2021/10/04 16:05:56 CreateDbIdentitySource : DbIdentitySource-2 [testacc-a]
2021/10/04 16:05:56 get DbIdentitySourceDTOs: all [testacc-a]
2021/10/04 16:05:56 getAppliance. 174 found for ID/name testacc-a
2021/10/04 16:05:56 CreateDbIdentityVault : DdIdentityVault-a [testacc-a]
2021/10/04 16:05:56 GetDbIdentityVaultDto. DdIdentityVault-a [testacc-a]
2021/10/04 16:05:56 GetDbIdentityVaultDto. DdIdentityVault-a found for ID/name DdIdentityVault-a
2021/10/04 16:05:56 UpdateDbIdentityVaultDto. : DdIdentityVault-a [testacc-a]
2021/10/04 16:05:56 deleteDbIdentityVaultDto. DdIdentityVault-a [testacc-a]
2021/10/04 16:05:56 deleteIntSaml2Ss. Deleted DdIdentityVault-a : true
2021/10/04 16:05:56 get DbIdentityVaultDtos: all [testacc-a]
2021/10/04 16:05:56 CreateDbIdentityVault : DbIdentityVault-2 [testacc-a]
2021/10/04 16:05:56 CreateDbIdentityVault : DbIdentityVault-3 [testacc-a]
2021/10/04 16:05:56 get DbIdentityVaultDtos: all [testacc-a]
2021/10/04 16:05:56 getAppliance. 174 found for ID/name testacc-a
2021/10/04 16:05:56 createExtSaml2Sp : Extsaml2-a [testacc-a]
2021/10/04 16:05:56 GetExtSaml2Sp. Extsaml2-a [testacc-a]
2021/10/04 16:05:56 GetExtSaml2Sp. Extsaml2-a found for ID/name Extsaml2-a
2021/10/04 16:05:56 UpdateExtSaml2Sp. : Extsaml2-a [testacc-a]
2021/10/04 16:05:57 deleteExtSaml2Sp. Extsaml2-a [testacc-a]
2021/10/04 16:05:57 deletesp. Deleted Extsaml2-a : true
2021/10/04 16:05:57 get ExtSaml2Sps: all [testacc-a]
2021/10/04 16:05:57 createExtSaml2Sp : Extsaml2-1 [testacc-a]
2021/10/04 16:05:57 createExtSaml2Sp : Extsaml2-2 [testacc-a]
2021/10/04 16:05:58 get ExtSaml2Sps: all [testacc-a]
2021/10/04 16:05:58 getAppliance. 174 found for ID/name testacc-a
2021/10/04 16:05:58 createIdP : idp-1 [testacc-a]
2021/10/04 16:05:58 getIdp. idp-1 [testacc-a]
2021/10/04 16:05:58 getIdP. idp-1 found for ID/name idp-1
2021/10/04 16:05:58 updateIdP. : idp-1 [testacc-a]
2021/10/04 16:05:58 deleteIdp. idp-1 [testacc-a]
2021/10/04 16:05:59 deleteIdp. Deleted idp-1 : true
2021/10/04 16:05:59 get idps: all [testacc-a]
2021/10/04 16:05:59 createIdP : ids-1 [testacc-a]
2021/10/04 16:05:59 createIdP : ids-2 [testacc-a]
2021/10/04 16:06:00 get idps: all [testacc-a]
2021/10/04 16:06:00 getAppliance. 174 found for ID/name testacc-a
2021/10/04 16:06:00 createIdSourceLdap : ids-a [testacc-a]
2021/10/04 16:06:00 getIdSourceLdap. ids-a [testacc-a]
2021/10/04 16:06:00 getIdSourceLdap. ids-a found for ID/name ids-a
2021/10/04 16:06:00 updateIdSourceLdap. : ids-a [testacc-a]
2021/10/04 16:06:01 deleteIdSourceLdap. ids-a [testacc-a]
2021/10/04 16:06:01 deleteIdSourceLdap. Deleted ids-a : true
2021/10/04 16:06:01 get idSourceLdaps: all [testacc-a]
2021/10/04 16:06:01 createIdSourceLdap : ids-1 [testacc-a]
2021/10/04 16:06:02 createIdSourceLdap : ids-2 [testacc-a]
2021/10/04 16:06:03 get idSourceLdaps: all [testacc-a]
2021/10/04 16:06:03 getAppliance. 174 found for ID/name testacc-a
2021/10/04 16:06:03 get idVaults: all [testacc-a]
2021/10/04 16:06:03 createIdVault : idVault-A [testacc-a]
2021/10/04 16:06:03 getIdVault. idVault-A [testacc-a]
2021/10/04 16:06:03 getIdVault. idVault-A found for ID/name idVault-A
2021/10/04 16:06:03 updateIdVault. : idVault-A [testacc-a]
2021/10/04 16:06:03 deleteIdVault. idVault-A [testacc-a]
2021/10/04 16:06:04 deleteIdVault. Deleted idVault-A : true
2021/10/04 16:06:04 get idVaults: all [testacc-a]
2021/10/04 16:06:04 createIdVault : IdVault-1 [testacc-a]
2021/10/04 16:06:05 createIdVault : IdVault-2 [testacc-a]
2021/10/04 16:06:05 get idVaults: all [testacc-a]
2021/10/04 16:06:05 getAppliances. found appliances 1
2021/10/04 16:06:05 deleteAppliance id: 174
2021/10/04 16:06:05 deleteAppliance. Deleted 174 : true
2021/10/04 16:06:05 createAppliance : testacc-z com.atricore.idbus.ida.testaccz
2021/10/04 16:06:05 createAppliance. ID: 175
2021/10/04 16:06:05 getAppliance. 175 found for ID/name testacc-z
2021/10/04 16:06:05 updateAppliance : testacc-z com.atricore.ida.a.mod
2021/10/04 16:06:05 updateAppliance. Updated: 175
2021/10/04 16:06:05 deleteAppliance id: 175
2021/10/04 16:06:05 deleteAppliance. Deleted 175 : true
2021/10/04 16:06:05 getAppliances. found appliances 0
2021/10/04 16:06:05 createAppliance : testacc-1 com.atricore.idbus.ida.testacc1
2021/10/04 16:06:05 createAppliance. ID: 176
2021/10/04 16:06:05 createAppliance : testacc-2 com.atricore.idbus.ida.testacc2
2021/10/04 16:06:05 createAppliance. ID: 177
2021/10/04 16:06:05 getAppliances. found appliances 2
2021/10/04 16:06:05 Importing appliance from JSON
2021/10/04 16:06:05 Importing appliance from JSON
2021/10/04 16:06:06 importAppliance. Error com.atricore.idbus.console.lifecycle.main.exception.ApplianceValidationException: There are 2 validation  errors for appliance testacc-01 : []string{"Appliance namespace is already used in some other appliance", "Appliance name already in use 'testacc-01' by 178"}
2021/10/04 16:06:06 getAppliance. not found for ID/name testacc-a
2021/10/04 16:06:06 createAppliance : testacc-a com.atricore.idbus.ida.t
2021/10/04 16:06:06 createAppliance. ID: 179
2021/10/04 16:06:06 createOidcRp : rp-2 [testacc-a]
2021/10/04 16:06:06 getOidcRp. rp-2 [testacc-a]
2021/10/04 16:06:06 getOidcRp. rp-2 found for ID/name rp-2
2021/10/04 16:06:06 updateOidcRp. : rp-2 [testacc-a]
2021/10/04 16:06:06 deleteOidcRp. rp-2 [testacc-a]
2021/10/04 16:06:06 deleteOidcRp. Deleted rp-2 : true
2021/10/04 16:06:06 get oidcRps: all [testacc-a]
2021/10/04 16:06:06 createOidcRp : rp-2 [testacc-a]
2021/10/04 16:06:06 createOidcRp : rp-2 [testacc-a]
2021/10/04 16:06:06 get oidcRps: all [testacc-a]
2021/10/04 16:06:06 getAppliance. 179 found for ID/name testacc-a
2021/10/04 16:06:06 createVirtSaml2Sp : VirtP-a [testacc-a]
2021/10/04 16:06:06 CreateVirtSaml2Sp. Error com.atricore.idbus.console.lifecycle.main.exception.ApplianceValidationException: There are 2 validation  errors for appliance testacc-a : []string{"No preferred Identity Provider Channel defined for SP VirtP-a", "No external/internal SAML 2 Identity Provider connected to VirtP-a"}
2021/10/04 16:06:06 getAppliances. found appliances 4
2021/10/04 16:06:06 deleteAppliance id: 176
2021/10/04 16:06:06 deleteAppliance. Deleted 176 : true
2021/10/04 16:06:06 deleteAppliance id: 177
2021/10/04 16:06:06 deleteAppliance. Deleted 177 : true
2021/10/04 16:06:06 deleteAppliance id: 178
2021/10/04 16:06:06 deleteAppliance. Deleted 178 : true
2021/10/04 16:06:06 deleteAppliance id: 179
2021/10/04 16:06:06 deleteAppliance. Deleted 179 : true
--- FAIL: TestAccCliSuite (10.16s)
    client_test.go:71: creating client
    client_test.go:79: created test client: [{http://localhost:8081/atricore-rest/services JOSSO Test server map[]}]
    client_test.go:200: ACCTEST_CLEAR_DATA: 
    client_test.go:207: clearing test data
    client_test.go:213: found 0 appliances
    client_test.go:86: SetupSuite complete
    --- FAIL: TestAccCliSuite/TestAccCliIdentityAppliance_import (0.21s)
        suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
            goroutine 54 [running]:
            runtime/debug.Stack(0xc00050b148, 0x93b5e0, 0xd36eb0)
            	/opt/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
            github.com/stretchr/testify/suite.failOnPanic(0xc000082c00)
            	/home/fbosch/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
            panic(0x93b5e0, 0xd36eb0)
            	/opt/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
            github.com/atricore/josso-sdk-go.(*IdbusApiClient).ImportAppliance(0xc0001af080, 0xc0004f4000, 0x1f02, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
            	/home/fbosch/wa/git/go/src/josso-sdk-go/identity_appliance.go:40 +0x621
            github.com/atricore/josso-sdk-go.(*AccTestSuite).ImportAppliance(0xc0001013e0, 0x9cf0e4, 0x19, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
            	/home/fbosch/wa/git/go/src/josso-sdk-go/client_test.go:37 +0x16a
            github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliIdentityAppliance_import(0xc0001013e0)
            	/home/fbosch/wa/git/go/src/josso-sdk-go/identity_appliance_test.go:15 +0x92
            reflect.Value.call(0xc0001b4660, 0xc000011120, 0x13, 0x9c53b7, 0x4, 0xc0005b3e30, 0x1, 0x1, 0xc0004f04f8, 0x40d6ea, ...)
            	/opt/atricore/tools/go/src/reflect/value.go:476 +0x8e7
            reflect.Value.Call(0xc0001b4660, 0xc000011120, 0x13, 0xc0004f0630, 0x1, 0x1, 0xb0fc1f, 0x2d, 0x438)
            	/opt/atricore/tools/go/src/reflect/value.go:337 +0xb9
            github.com/stretchr/testify/suite.Run.func1(0xc000082c00)
            	/home/fbosch/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
            testing.tRunner(0xc000082c00, 0xc0001c14d0)
            	/opt/atricore/tools/go/src/testing/testing.go:1193 +0xef
            created by testing.(*T).Run
            	/opt/atricore/tools/go/src/testing/testing.go:1238 +0x2b3
    --- FAIL: TestAccCliSuite/TestAccCliIdentityAppliance_z010 (0.28s)
        identity_appliance_test.go:123: Acceptance test z010 : basic appliance
        identity_appliance_test.go:127: z010, importing appliance : com.atricore.idbus.console.lifecycle.main.exception.ApplianceValidationException: There are 2 validation  errors for appliance testacc-01 : []string{"Appliance namespace is already used in some other appliance", "Appliance name already in use 'testacc-01' by 178"} 
    --- FAIL: TestAccCliSuite/TestAccCliIdentityAppliance_z030 (0.00s)
        identity_appliance_test.go:193: Acceptance test #030 : basic idp
        suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
            goroutine 56 [running]:
            runtime/debug.Stack(0xc0005bac80, 0x93b5e0, 0xd36eb0)
            	/opt/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
            github.com/stretchr/testify/suite.failOnPanic(0xc000082f00)
            	/home/fbosch/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
            panic(0x93b5e0, 0xd36eb0)
            	/opt/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
            github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliIdentityAppliance_z030(0xc0001013e0)
            	/home/fbosch/wa/git/go/src/josso-sdk-go/identity_appliance_test.go:203 +0x1f7
            reflect.Value.call(0xc0001b4660, 0xc000011168, 0x13, 0x9c53b7, 0x4, 0xc0005b3e30, 0x1, 0x1, 0xc0004f04f8, 0x40d6ea, ...)
            	/opt/atricore/tools/go/src/reflect/value.go:476 +0x8e7
            reflect.Value.Call(0xc0001b4660, 0xc000011168, 0x13, 0xc0004f0630, 0x1, 0x1, 0xb0fc1f, 0x2d, 0x438)
            	/opt/atricore/tools/go/src/reflect/value.go:337 +0xb9
            github.com/stretchr/testify/suite.Run.func1(0xc000082f00)
            	/home/fbosch/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
            testing.tRunner(0xc000082f00, 0xc0001c1680)
            	/opt/atricore/tools/go/src/testing/testing.go:1193 +0xef
            created by testing.(*T).Run
            	/opt/atricore/tools/go/src/testing/testing.go:1238 +0x2b3
    --- FAIL: TestAccCliSuite/TestAccCliVirtSaml2_crud (0.04s)
        virsaml2sp_test.go:33: com.atricore.idbus.console.lifecycle.main.exception.ApplianceValidationException: There are 2 validation  errors for appliance testacc-a : []string{"No preferred Identity Provider Channel defined for SP VirtP-a", "No external/internal SAML 2 Identity Provider connected to VirtP-a"}
    client_test.go:92: Teardown suite
    client_test.go:200: ACCTEST_CLEAR_DATA: 
    client_test.go:207: clearing test data
    client_test.go:213: found 4 appliances
    client_test.go:219: deleting appliance 176
    client_test.go:219: deleting appliance 177
    client_test.go:219: deleting appliance 178
    client_test.go:219: deleting appliance 179
FAIL
FAIL	github.com/atricore/josso-sdk-go	10.167s
?   	github.com/atricore/josso-sdk-go/cmd/josso-cli	[no test files]
FAIL
