package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func (s *AccTestSuite) TestAccCliIdSourceLdap_crud() {
	var t = s.T()

	appliance, err := getTestAppliance(s.T(), s.client)
	if err != nil {
		s.client.Logger().Errorf("cannot get test appliance %v", err)
		t.Error(err)
		return
	}

	crudName := "ids-a"
	var orig *api.LdapIdentitySourceDTO
	var created api.LdapIdentitySourceDTO
	orig = createTestLdapIdentitySourceDTO(crudName)

	// Test CREATE
	created, err = s.client.CreateIdSourceLdap(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := LdapIdentitySourceValidateCreate(orig, &created); err != nil {
		t.Errorf("creating IdSourceLDap : %v", err)
		return
	}

	// Test READ
	var read api.LdapIdentitySourceDTO
	read, err = s.client.GetIdSourceLdap(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = LdapIdentitySourceValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating IdSourceLDap : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	updated, err := s.client.UpdateIdSourceLdap(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = LdapIdentitySourceValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}
	//.
	//Test Delete
	deleted, err := s.client.DeleteIdSourceLdap(*appliance.Name, crudName)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", crudName)
		return
	}

	// Test empty list
	listOfAll, err := s.client.GetIdSourceLdaps(*appliance.Name)
	if err != nil {
		t.Error(err)
		return
	}
	if len(listOfAll) != 0 {
		// The list should be emtpy
		t.Errorf("Invalid number of elements found %d, expeted 0", len(listOfAll))
		return
	}

	// List of created elements, order by Name, (these elements must have all the variables of the structure)
	var listOfCreated [2]api.LdapIdentitySourceDTO

	element1 := createTestLdapIdentitySourceDTO("ids-1")
	listOfCreated[0], _ = s.client.CreateIdSourceLdap(*appliance.Name, *element1)

	element2 := createTestLdapIdentitySourceDTO("ids-2")
	listOfCreated[1], _ = s.client.CreateIdSourceLdap(*appliance.Name, *element2)

	// Get list from server
	listOfRead, err := s.client.GetIdSourceLdaps(*appliance.Name)
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
		if err = LdapIdentitySourceValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

}

func createTestLdapIdentitySourceDTO(name string) *api.LdapIdentitySourceDTO {
	orig := api.NewLdapIdentitySourceDTO()
	orig.SetName(name)
	orig.SetId(-1)
	orig.SetProviderUrl("ldap://192.168.0.97:389")
	orig.SetSecurityPrincipal(fmt.Sprintf("CN=%s,CN=Users,DC=mycompany,DC=com", name))
	orig.SetSecurityCredential("@WSX3edc%s")
	orig.SetUsersCtxDN(fmt.Sprintf("CN=%s,CN=Users,DC=mycompany,DC=com", name))
	orig.SetRolesCtxDN(fmt.Sprintf("CN=%s,CN=Users,DC=mycompany,DC=com", name))
	orig.SetUidAttributeID("member")
	orig.SetPrincipalUidAttributeID("sAMAccountName")
	orig.SetRoleAttributeID("sAMAccountName")
	orig.SetReferrals("follow")
	orig.SetLdapSearchScope("subtree")
	orig.SetInitialContextFactory("true")
	orig.SetRoleMatchingMode("manager")
	orig.SetUserPropertiesQueryString("space")
	orig.SetSecurityAuthentication("authenticated")

	return orig
}

func (s *AccTestSuite) TestAccCliIdS_createFailOnDupName() {

	// TODO ! implement me!

}

func (s *AccTestSuite) TestAccCliIdS_updateFailOnDupName() {

	// TODO ! implement me!

}

//Fields to validate after appliance creation
func IdSFieldTestCreate(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) []FiledTestStruct {

	return []FiledTestStruct{
		{
			name:     "name",
			cmp:      func() bool { return StrPtrEquals(e.Name, r.Name) },
			expected: StrDeref(e.Name),
			received: StrDeref(r.Name),
		},
		{
			name:     "provider_url",
			cmp:      func() bool { return StrPtrEquals(e.ProviderUrl, r.ProviderUrl) },
			expected: StrDeref(e.ProviderUrl),
			received: StrDeref(r.ProviderUrl),
		},
		{
			name:     "security_principal",
			cmp:      func() bool { return StrPtrEquals(e.SecurityPrincipal, r.SecurityPrincipal) },
			expected: StrDeref(e.SecurityPrincipal),
			received: StrDeref(r.SecurityPrincipal),
		},
		{
			name:     "users_ctx_dn",
			cmp:      func() bool { return StrPtrEquals(e.UsersCtxDN, r.UsersCtxDN) },
			expected: StrDeref(e.UsersCtxDN),
			received: StrDeref(r.UsersCtxDN),
		},
		{
			name:     "roles_ctx_dn",
			cmp:      func() bool { return StrPtrEquals(e.RolesCtxDN, r.RolesCtxDN) },
			expected: StrDeref(e.RolesCtxDN),
			received: StrDeref(r.RolesCtxDN),
		},
		{
			name:     "uid_attribute_id",
			cmp:      func() bool { return StrPtrEquals(e.UidAttributeID, r.UidAttributeID) },
			expected: StrDeref(e.UidAttributeID),
			received: StrDeref(r.UidAttributeID),
		},
		{
			name:     "principal_uid_attribute_id",
			cmp:      func() bool { return StrPtrEquals(e.PrincipalUidAttributeID, r.PrincipalUidAttributeID) },
			expected: StrDeref(e.PrincipalUidAttributeID),
			received: StrDeref(r.PrincipalUidAttributeID),
		},
		{
			name:     "role_attribute_id",
			cmp:      func() bool { return StrPtrEquals(e.RoleAttributeID, r.RoleAttributeID) },
			expected: StrDeref(e.RoleAttributeID),
			received: StrDeref(r.RoleAttributeID),
		},
		{
			name:     "referrals",
			cmp:      func() bool { return StrPtrEquals(e.Referrals, r.Referrals) },
			expected: StrDeref(e.Referrals),
			received: StrDeref(r.Referrals),
		},
		{
			name:     "ldap_search_scope",
			cmp:      func() bool { return StrPtrEquals(e.LdapSearchScope, r.LdapSearchScope) },
			expected: StrDeref(e.LdapSearchScope),
			received: StrDeref(r.LdapSearchScope),
		},
		{
			name:     "initial_context_factory",
			cmp:      func() bool { return StrPtrEquals(e.InitialContextFactory, r.InitialContextFactory) },
			expected: StrDeref(e.InitialContextFactory),
			received: StrDeref(r.InitialContextFactory),
		},
		{
			name:     "role_matching_mode",
			cmp:      func() bool { return StrPtrEquals(e.RoleMatchingMode, r.RoleMatchingMode) },
			expected: StrDeref(e.RoleMatchingMode),
			received: StrDeref(r.RoleMatchingMode),
		},
		{
			name:     "user_properties_query_string",
			cmp:      func() bool { return StrPtrEquals(e.UserPropertiesQueryString, r.UserPropertiesQueryString) },
			expected: StrDeref(e.UserPropertiesQueryString),
			received: StrDeref(r.UserPropertiesQueryString),
		},
		{
			name:     "security_authentication",
			cmp:      func() bool { return StrPtrEquals(e.SecurityAuthentication, r.SecurityAuthentication) },
			expected: StrDeref(e.SecurityAuthentication),
			received: StrDeref(r.SecurityAuthentication),
		},
	}
}

//Fields to validate after IdS update
func IdSFieldTestUpdate(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) []FiledTestStruct {

	t := []FiledTestStruct{
		{
			name:     "id",
			cmp:      func() bool { return Int64PtrEquals(e.Id, r.Id) },
			expected: strconv.FormatInt(Int64Deref(e.Id), 10),
			received: strconv.FormatInt(Int64Deref(r.Id), 10),
		},
		{
			name:     "element_id",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: StrDeref(e.ElementId),
			received: StrDeref(r.ElementId),
		},
	}
	return append(t, IdSFieldTestCreate(e, r)...)
}

// Compares the expected IdSourceLDap with the received one.
func LdapIdentitySourceValidateCreate(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) error {

	return ValidateFields(IdSFieldTestCreate(e, r))
}

// Compares the expected IdSourceLDap with the received one.
func LdapIdentitySourceValidateUpdate(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) error {

	return ValidateFields(IdSFieldTestUpdate(e, r))
}
