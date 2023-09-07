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
	tData := api.NewExternalOpenIDConnectRelayingPartyDTO()

	var Grants []string
	Grants = append(Grants, "AUTHORIZATION_CODE")
	tData.SetGrants(Grants)

	var RType []string
	RType = append(RType, "CODE")
	tData.SetResponseTypes(RType)

	var LogUris []string
	LogUris = append(LogUris, "http://localhost:8080/app/logout")
	tData.SetPostLogoutRedirectionURIs(LogUris)

	var AuthUrl []string
	AuthUrl = append(AuthUrl, " http://localhost:8080/app", "http://localhost:8080/app/secure")
	tData.SetAuthorizedURIs(AuthUrl)

	tData.SetClientAuthnMethod("CLIENT_SECRET_BASIC")
	tData.SetClientCert("")
	tData.SetClientId("")
	tData.SetClientSecret("")
	tData.SetClientType("")
	tData.SetDescription("")
	tData.SetDisplayName("")
	tData.SetEncryptionAlg("")
	tData.SetEncryptionMethod("")
	tData.SetId(-1)
	tData.SetIdTokenEncryptionAlg("")
	tData.SetIdTokenEncryptionMethod("")
	//orig.SetIdentityLookups(IdentityLookupDTO)
	//tData.SetIsRemote(true)
	tData.SetName("rp-2")
	//tData.SetRemote(true)
	tData.SetSigningAlg("RS256")

	return tData
}

func (s *AccTestSuite) TestAccCliOidcRp_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliOidcRp_updateFailOnDupName() {

	// TODO ! implement me!

}

// Fields to validate after appliance creation
func OidcRpFieldTestCreate(
	e *api.ExternalOpenIDConnectRelayingPartyDTO,
	r *api.ExternalOpenIDConnectRelayingPartyDTO) []FiledTestStruct {

	return []FiledTestStruct{
		// {
		// 	name:     "grants",
		// 	cmp:      func() bool { return StrPtrEquals(e.Grants, r.Grants) },
		// 	expected: StrDeref(e.Grants),
		// 	received: StrDeref(r.Grants),
		// },
		// {
		// 	name:     "responseTypes",
		// 	cmp:      func() bool { return StrPtrEquals(e.ResponseTypes, r.ResponseTypes) },
		// 	expected: StrDeref(e.ResponseTypes),
		// 	received: StrDeref(r.ResponseTypes),
		// },
		// {
		// 	name:     "postlogoutredirectionuris",
		// 	cmp:      func() bool { return StrPtrEquals(e.PostLogoutRedirectionURIs, r.PostLogoutRedirectionURIs) },
		// 	expected: StrDeref(e.PostLogoutRedirectionURIs),
		// 	received: StrDeref(r.PostLogoutRedirectionURIs),
		// },
		// {
		// 	name:     "authorizeduris",
		// 	cmp:      func() bool { return StrPtrEquals(e.AuthorizedURIs, r.AuthorizedURIs) },
		// 	expected: StrDeref(e.AuthorizedURIs),
		// 	received: StrDeref(r.AuthorizedURIs),
		// },
		{
			name:     "client_authn_method",
			cmp:      func() bool { return StrPtrEquals(e.ClientAuthnMethod, r.ClientAuthnMethod) },
			expected: StrDeref(e.ClientAuthnMethod),
			received: StrDeref(r.ClientAuthnMethod),
		},
		{
			name:     "client_cert",
			cmp:      func() bool { return StrPtrEquals(e.ClientCert, r.ClientCert) },
			expected: StrDeref(e.ClientCert),
			received: StrDeref(r.ClientCert),
		},
		{
			name:     "client_secret",
			cmp:      func() bool { return StrPtrEquals(e.ClientSecret, r.ClientSecret) },
			expected: StrDeref(e.ClientSecret),
			received: StrDeref(r.ClientSecret),
		},
		{
			name:     "client_type",
			cmp:      func() bool { return StrPtrEquals(e.ClientType, r.ClientType) },
			expected: StrDeref(e.ClientType),
			received: StrDeref(r.ClientType),
		},
		{
			name:     "description(",
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
		{
			name:     "encryption_alg",
			cmp:      func() bool { return StrPtrEquals(e.EncryptionAlg, r.EncryptionAlg) },
			expected: StrDeref(e.EncryptionAlg),
			received: StrDeref(r.EncryptionAlg),
		},
		{
			name:     "encryption_method",
			cmp:      func() bool { return StrPtrEquals(e.EncryptionAlg, r.EncryptionAlg) },
			expected: StrDeref(e.EncryptionAlg),
			received: StrDeref(r.EncryptionAlg),
		},
		{
			name:     "id_token_encryption_alg(",
			cmp:      func() bool { return StrPtrEquals(e.IdTokenEncryptionAlg, r.IdTokenEncryptionAlg) },
			expected: StrDeref(e.IdTokenEncryptionAlg),
			received: StrDeref(r.IdTokenEncryptionAlg),
		},
		{
			name:     "id_token_encryption_method",
			cmp:      func() bool { return StrPtrEquals(e.IdTokenEncryptionAlg, r.IdTokenEncryptionAlg) },
			expected: StrDeref(e.IdTokenEncryptionAlg),
			received: StrDeref(r.IdTokenEncryptionAlg),
		},
		// {
		// 	name:     "identitylookups",
		// 	cmp:      func() bool { return StrPtrEquals(e.IdentityLookups, r.IdentityLookups) },
		// 	expected: StrDeref(e.IdentityLookups),
		// 	received: StrDeref(r.IdentityLookups),
		// },

		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "signing_alg",
			cmp:      func() bool { return StrPtrEquals(e.SigningAlg, r.SigningAlg) },
			expected: StrDeref(e.SigningAlg),
			received: StrDeref(r.SigningAlg),
		},
	}
}

// Fields to validate after OidcRp update
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
