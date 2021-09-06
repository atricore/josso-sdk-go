go mod tidy
go mod vendor
==> Checking that code complies with gofmt requirements...
go clean -testcache
TF_ACC=1 go test $(go list ./... |grep -v 'vendor') -v   -timeout 120m
=== RUN   TestAccCliSuite
    client_test.go:25: creating client
2021/09/06 13:42:20 newIdbusApiClient TRACE: false
2021/09/06 13:42:20 registering server http://localhost:8081/atricore-rest/services
2021/09/06 13:42:20 adding server configuration for http://localhost:8081/atricore-rest/services
2021/09/06 13:42:20 authn: idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7 true/admin true
    client_test.go:33: created test client: [{http://localhost:8081/atricore-rest/services JOSSO Test server map[]}]
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/09/06 13:42:20 getAppliances. found appliances 1
=== RUN   TestAccCliSuite/TestAccCliDbIdentitySourceDTO_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliDbIdentitySourceDTO_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliDbIdentitySourceDto
2021/09/06 13:42:20 getAppliance. not found for ID/name ida-a
2021/09/06 13:42:20 createAppliance : ida-a com.atricore.idbus.ida.t
2021/09/06 13:42:21 createAppliance. ID: 12
2021/09/06 13:42:21 CreateDbIdentitySource : DdIdentityVauld-a [ida-a]
2021/09/06 13:42:21 CreateDbIdentitySource. Error 500 Internal Server Error
    dbIdentity_source_test.go:28: 500 Internal Server Error
=== RUN   TestAccCliSuite/TestAccCliDbIdentitySourceDto_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliDbIdentityvaultDto
2021/09/06 13:42:21 getAppliance. 12 found for ID/name ida-a
2021/09/06 13:42:21 CreateDbIdentitySource : DdIdentityVauld-a [ida-a]
2021/09/06 13:42:21 CreateDbIdentitySource. Error 500 Internal Server Error
    dbIdentity_vauld_test.go:29: 500 Internal Server Error
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud
2021/09/06 13:42:21 getAppliance. 12 found for ID/name ida-a
    suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
        goroutine 8 [running]:
        runtime/debug.Stack(0xc00034aa78, 0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc000001c80)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestExternalSaml2ServiceProviderDTO(0x9c4e72, 0xa, 0xc0000bf5c0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/extsaml2sp_test.go:557 +0x481f
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliExtSaml2_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/extsaml2sp_test.go:29 +0x2ac
        reflect.Value.call(0xc00016c480, 0xc000010270, 0x13, 0x9c2d53, 0x4, 0xc000055e30, 0x1, 0x1, 0xc0000444f8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc000010270, 0x13, 0xc000044630, 0x1, 0x1, 0xb14281, 0x2e, 0x439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc000001c80)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc000001c80, 0xc000078990)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliExtSaml2_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdP_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdP_crud
2021/09/06 13:42:21 getAppliance. 12 found for ID/name ida-a
    suite.go:63: test panicked: assignment to entry in nil map
        goroutine 31 [running]:
        runtime/debug.Stack(0xc00035ff00, 0x93bb80, 0xa592f0)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc000083c80)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93bb80, 0xa592f0)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestBasicAuthn(0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/idp_test.go:159 +0x105
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliIdP_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/idp_test.go:25 +0x2b8
        reflect.Value.call(0xc00016c480, 0xc0000102d0, 0x13, 0x9c2d53, 0x4, 0xc000051e30, 0x1, 0x1, 0xc00003ecf8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc0000102d0, 0x13, 0xc00003ee30, 0x1, 0x1, 0xb14281, 0x2e, 0xc000000439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc000083c80)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc000083c80, 0xc000078bd0)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== RUN   TestAccCliSuite/TestAccCliIdP_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdS_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdS_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdSourceLdap_crud
2021/09/06 13:42:21 getAppliance. 12 found for ID/name ida-a
2021/09/06 13:42:21 createIdSourceLdap : ids-a [ida-a]
2021/09/06 13:42:21 getIdSourceLdap. ids-a [ida-a]
2021/09/06 13:42:21 getIdSourceLdap. %!d(string=ids-a) found for ID/name ids-a
2021/09/06 13:42:21 updateIdSourceLdap. : ids-a [ida-a]
2021/09/06 13:42:21 deleteIdSourceLdap. ids-a [ida-a]
2021/09/06 13:42:21 deleteIdSourceLdap. Deleted ids-a : true
2021/09/06 13:42:21 get idSourceLdaps: all [ida-a]
2021/09/06 13:42:21 createIdSourceLdap : ids-1 [ida-a]
2021/09/06 13:42:21 createIdSourceLdap : ids-2 [ida-a]
2021/09/06 13:42:21 get idSourceLdaps: all [ida-a]
=== RUN   TestAccCliSuite/TestAccCliIdVault_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdVault_crud
2021/09/06 13:42:21 getAppliance. 12 found for ID/name ida-a
2021/09/06 13:42:21 get idVaults: all [ida-a]
2021/09/06 13:42:21 createIdVault : idVault-A [ida-a]
2021/09/06 13:42:21 getIdVault. idVault-A [ida-a]
2021/09/06 13:42:21 getIdVault. %!d(string=idVault-A) found for ID/name idVault-A
2021/09/06 13:42:21 updateIdVault. : idVault-A [ida-a]
2021/09/06 13:42:21 deleteIdVault. idVault-A [ida-a]
2021/09/06 13:42:21 deleteIdVault. Deleted idVault-A : true
2021/09/06 13:42:21 get idVaults: all [ida-a]
2021/09/06 13:42:21 createIdVault : IdVault-1 [ida-a]
2021/09/06 13:42:22 createIdVault : IdVault-2 [ida-a]
2021/09/06 13:42:22 get idVaults: all [ida-a]
=== RUN   TestAccCliSuite/TestAccCliIdVault_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIdentityAppliance_crud
    suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
        goroutine 37 [running]:
        runtime/debug.Stack(0xc00034ba80, 0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc0002b8180)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestIdentityApplianceDefinitionDTO(0x9c337f, 0x5, 0x0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/identity_appliance_test.go:346 +0x27bf
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliIdentityAppliance_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/identity_appliance_test.go:16 +0x7c
        reflect.Value.call(0xc00016c480, 0xc000010390, 0x13, 0x9c2d53, 0x4, 0xc00021fe30, 0x1, 0x1, 0xc00003e4f8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc000010390, 0x13, 0xc00003e630, 0x1, 0x1, 0xb14281, 0x2e, 0xc000000439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc0002b8180)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc0002b8180, 0xc000079050)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== RUN   TestAccCliSuite/TestAccCliIntSaml2_crud
2021/09/06 13:42:22 getAppliance. 12 found for ID/name ida-a
    suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
        goroutine 38 [running]:
        runtime/debug.Stack(0xc0003a6628, 0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc0002b8480)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestInternalSaml2ServiceProviderDTO(0x9c4ee0, 0xa, 0xc00038aa50)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/Intsaml2sp_test.go:331 +0x203f
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliIntSaml2_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/Intsaml2sp_test.go:24 +0x2ac
        reflect.Value.call(0xc00016c480, 0xc0000103a8, 0x13, 0x9c2d53, 0x4, 0xc00021fe30, 0x1, 0x1, 0xc00003e4f8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc0000103a8, 0x13, 0xc00003e630, 0x1, 0x1, 0xb14281, 0x2e, 0xc000000439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc0002b8480)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc0002b8480, 0xc0000790e0)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== RUN   TestAccCliSuite/TestAccCliIntSaml2_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliIntSaml2_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliOidcRp_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliOidcRp_crud
2021/09/06 13:42:22 getAppliance. 12 found for ID/name ida-a
    suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
        goroutine 53 [running]:
        runtime/debug.Stack(0xc00034af60, 0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc000382780)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestExternalOpenIDConnectRelayingPartyDTO(0x9c2f9f, 0x4, 0xc00038b3b0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/oidcrp_test.go:344 +0x1e5f
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliOidcRp_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/oidcrp_test.go:24 +0x2ac
        reflect.Value.call(0xc00016c480, 0xc000010408, 0x13, 0x9c2d53, 0x4, 0xc000051e30, 0x1, 0x1, 0xc00003ecf8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc000010408, 0x13, 0xc00003ee30, 0x1, 0x1, 0xb14281, 0x2e, 0xc000000439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc000382780)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc000382780, 0xc000079320)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== RUN   TestAccCliSuite/TestAccCliOidcRp_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliVirSaml2_crud_createFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliVirSaml2_crud_updateFailOnDupName
=== RUN   TestAccCliSuite/TestAccCliVirtSaml2_crud
2021/09/06 13:42:22 getAppliance. 12 found for ID/name ida-a
    suite.go:63: test panicked: runtime error: invalid memory address or nil pointer dereference
        goroutine 57 [running]:
        runtime/debug.Stack(0xc0003a6850, 0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/debug/stack.go:24 +0x9f
        github.com/stretchr/testify/suite.failOnPanic(0xc000382f00)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:63 +0x5b
        panic(0x93ba00, 0xd38e90)
        	/data/atricore/tools/go/src/runtime/panic.go:965 +0x1b9
        github.com/atricore/josso-sdk-go.createTestVirtualSaml2ServiceProviderDTO(0x9c57a5, 0xb, 0xc00037b758)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/virsaml2sp_test.go:314 +0x1a9f
        github.com/atricore/josso-sdk-go.(*AccTestSuite).TestAccCliVirtSaml2_crud(0xc0000d63a0)
        	/home/sgonzalez/wa/git/go/src/josso-sdk-go/virsaml2sp_test.go:23 +0x2cc
        reflect.Value.call(0xc00016c480, 0xc000010468, 0x13, 0x9c2d53, 0x4, 0xc000051e30, 0x1, 0x1, 0xc00003ecf8, 0x40d6ea, ...)
        	/data/atricore/tools/go/src/reflect/value.go:476 +0x8e7
        reflect.Value.Call(0xc00016c480, 0xc000010468, 0x13, 0xc00003ee30, 0x1, 0x1, 0xb14281, 0x2e, 0xc000000439)
        	/data/atricore/tools/go/src/reflect/value.go:337 +0xb9
        github.com/stretchr/testify/suite.Run.func1(0xc000382f00)
        	/home/sgonzalez/wa/git/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:158 +0x379
        testing.tRunner(0xc000382f00, 0xc000079560)
        	/data/atricore/tools/go/src/testing/testing.go:1194 +0xef
        created by testing.(*T).Run
        	/data/atricore/tools/go/src/testing/testing.go:1239 +0x2b3
=== CONT  TestAccCliSuite
    client_test.go:189: ACCTEST_CLEAR_DATA: 
    client_test.go:196: clearing test data
2021/09/06 13:42:22 getAppliances. found appliances 2
    client_test.go:205: deleting appliance 12
2021/09/06 13:42:22 deleteAppliance id: 12
2021/09/06 13:42:22 deleteAppliance. Deleted 12 : true
--- FAIL: TestAccCliSuite (1.41s)
    --- PASS: TestAccCliSuite/TestAccCliDbIdentitySourceDTO_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliDbIdentitySourceDTO_crud_updateFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliDbIdentitySourceDto (0.07s)
    --- PASS: TestAccCliSuite/TestAccCliDbIdentitySourceDto_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliDbIdentityVaultDto_crud_updateFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliDbIdentityvaultDto (0.01s)
    --- FAIL: TestAccCliSuite/TestAccCliExtSaml2_crud (0.02s)
    --- PASS: TestAccCliSuite/TestAccCliExtSaml2_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliExtSaml2_crud_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdP_createFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliIdP_crud (0.01s)
    --- PASS: TestAccCliSuite/TestAccCliIdP_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdS_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdS_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdSourceLdap_crud (0.28s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_crud (0.81s)
    --- PASS: TestAccCliSuite/TestAccCliIdVault_updateFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliIdentityAppliance_crud (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliIntSaml2_crud (0.02s)
    --- PASS: TestAccCliSuite/TestAccCliIntSaml2_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliIntSaml2_crud_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliOidcRp_createFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliOidcRp_crud (0.02s)
    --- PASS: TestAccCliSuite/TestAccCliOidcRp_updateFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliVirSaml2_crud_createFailOnDupName (0.00s)
    --- PASS: TestAccCliSuite/TestAccCliVirSaml2_crud_updateFailOnDupName (0.00s)
    --- FAIL: TestAccCliSuite/TestAccCliVirtSaml2_crud (0.02s)
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
FAIL	github.com/atricore/josso-sdk-go	1.416s
FAIL
