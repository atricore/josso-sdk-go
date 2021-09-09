package cli

import (
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliOidcRp_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	crudName := "rp-a"
	var orig *api.ExternalOpenIDConnectRelayingPartyDTO
	var created api.ExternalOpenIDConnectRelayingPartyDTO
	orig = createTestExternalOpenIDConnectRelayingPartyDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateOidcRp(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := OidcRpValidateCreate(orig, &created); err != nil {
		t.Errorf("creating rp : %v", err)
		return
	}

	// Test READ
	var read api.ExternalOpenIDConnectRelayingPartyDTO
	read, err = s.client.GetOidcRp(*appliance.Name, "rp-2")
	if err != nil {
		t.Error(err)
		return
	}
	if err = OidcRpValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating rp : %v", err)
		return
	}

	// Test UPDATE
	// 1. Modify an existing OIDCRP store in read, set description to a new value
	read.Description = api.PtrString("My updated description")
	read.ClientId = api.PtrString("1234")
	read.ClientType = api.PtrString("type1")

	// 2. Send update request to server
	updated, err := s.client.UpdateOidcRp(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}

	// 3. Validate updated vs original name, descriptions must be OK
	if err = OidcRpValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}

	// Test Delete
	toDelete := "rp-2" // Name of the RP to be deleted
	// 1. Send delete request to server using ODIC RP name
	deleted, err := s.client.DeleteOidcRp(*appliance.Name, toDelete)
	if err != nil {
		t.Error(err)
		return
	}
	// 2. Validate that the RP was deleted
	if !deleted {
		t.Errorf("Not deleted! %s", toDelete)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetOidcRps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name
	var listOfCreated [2]api.ExternalOpenIDConnectRelayingPartyDTO
	// Test list of #2 elements
	element1 := createTestExternalOpenIDConnectRelayingPartyDTO("rp-1")
	listOfCreated[0], _ = s.client.CreateOidcRp(*appliance.Name, *element1)

	element2 := createTestExternalOpenIDConnectRelayingPartyDTO("rp-2")
	listOfCreated[1], _ = s.client.CreateOidcRp(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetOidcRps(*appliance.Name)
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
		if err = OidcRpValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}
}

func createTestExternalOpenIDConnectRelayingPartyDTO(name string) *api.ExternalOpenIDConnectRelayingPartyDTO {
	orig := api.NewExternalOpenIDConnectRelayingPartyDTO()

	var Grants []string
	Grants = append(Grants, "AUTHORIZATION_CODE")

	var RType []string
	RType = append(RType, "CODE")

	var LogUris []string
	LogUris = append(LogUris, "http://localhost:8080/app/logout")

	var AuthUrl []string
	AuthUrl = append(AuthUrl, " http://localhost:8080/app", "http://localhost:8080/app/secure")

	// orig.SetActiveBindings()
	// orig.SetActiveProfiles()
	orig.SetAuthorizedURIs(AuthUrl)
	orig.SetClientAuthnMethod("CLIENT_SECRET_BASIC")
	orig.SetClientCert("")
	orig.SetClientId("")
	orig.SetClientSecret("")
	orig.SetClientType("")
	//orig.SetConfig(conf)
	orig.SetDescription("")
	orig.SetDisplayName("")
	orig.SetElementId("")
	orig.SetEncryptionAlg("")
	orig.SetEncryptionMethod("")
	orig.SetGrants(Grants)
	orig.SetId(-1)
	orig.SetIdTokenEncryptionAlg("")
	orig.SetIdTokenEncryptionMethod("")
	//orig.SetIdentityAppliance(identityAppliance)
	//orig.SetIdentityLookups(IdentityLookupDTO)
	orig.SetIsRemote(true)
	//orig.SetLocation(locat)
	//orig.SetMetadata(ResourceDTO)
	orig.SetName("rp-2")
	orig.SetPostLogoutRedirectionURIs(LogUris)
	orig.SetRemote(true)
	orig.SetResponseTypes(RType)
	//orig.SetRole("")
	orig.SetSigningAlg("RS256")

	return orig
}

func (s *AccTestSuite) TestAccCliOidcRp_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliOidcRp_updateFailOnDupName() {

	// TODO ! implement me!

}

//Fields to validate after appliance creation
func OidcRpFieldTestCreate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) []FiledTestStruct {

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
			name:     "clientId",
			cmp:      func() bool { return StrPtrEquals(e.ClientId, r.ClientId) },
			expected: StrDeref(e.ClientId),
			received: StrDeref(r.ClientId),
		},
		{
			name:     "clientType",
			cmp:      func() bool { return StrPtrEquals(e.ClientType, r.ClientType) },
			expected: StrDeref(e.ClientType),
			received: StrDeref(r.ClientType),
		},
	}
}

//Fields to validate after OidcRp update
func OidcRpFieldTestUpdate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
	}

	return append(t, OidcRpFieldTestCreate(e, r)...)
}

// Compares the expected OidcRp with the received one.
func OidcRpValidateCreate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) error {

	return ValidateFields(OidcRpFieldTestCreate(e, r))
}

// Compares the expected OidcRp with the received one.
func OidcRpValidateUpdate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) error {

	return ValidateFields(OidcRpFieldTestUpdate(e, r))
}
