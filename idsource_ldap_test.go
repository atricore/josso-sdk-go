package cli

import (
	"sort"
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

	var created api.LdapIdentitySourceDTO
	orig := api.NewLdapIdentitySourceDTO()
	orig.SetName("ids-2")
	orig.SetId(-1)
	orig.SetProviderUrl("ldap://192.168.0.97:389")
	orig.SetSecurityPrincipal("CN=Administrator,CN=Users,DC=mycompany,DC=com")
	orig.SetSecurityCredential("@WSX3edc")
	orig.SetUsersCtxDN("CN=Users,DC=mycompany,DC=com")
	orig.SetRolesCtxDN("CN=Users,DC=mycompany,DC=com")
	orig.SetUidAttributeID("member")
	orig.SetPrincipalUidAttributeID("sAMAccountName")
	orig.SetRoleAttributeID("sAMAccountName")
	orig.SetReferrals("follow")
	orig.SetLdapSearchScope("subtree")
	orig.SetInitialContextFactory("true")
	orig.SetRoleMatchingMode("manager")
	orig.SetUserPropertiesQueryString("space")
	orig.SetSecurityAuthentication("authenticated")

	// Test CREATE
	created, err = s.client.CreateIdSourceLdap(*appliance.Name, *orig)
	if err != nil {
		t.Error(err)
		return
	}
	if err := IdSValidateUpdate(orig, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}

	// Test READ
	var read api.LdapIdentitySourceDTO
	read, err = s.client.GetIdSourceLdap(*appliance.Name, "ids-2")
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdSValidateUpdate(&read, &created); err != nil {
		t.Errorf("creating idp : %v", err)
		return
	}

	// Test Update
	read.Description = api.PtrString("Updated description")
	read.ElementId = api.PtrString("dirt")
	updated, err := s.client.UpdateIdSourceLdap(*appliance.Name, read)
	if err != nil {
		t.Error(err)
		return
	}
	if err = IdSValidateUpdate(&read, &updated); err != nil {
		t.Error(err)
		return
	}
	//.
	//Test Delete
	toDelete := "ids-2"
	deleted, err := s.client.DeleteIdSourceLdap(*appliance.Name, toDelete)
	if err != nil {
		t.Error(err)
		return
	}
	if !deleted {
		t.Errorf("Not deleted! %s", toDelete)
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

	element1 := api.LdapIdentitySourceDTO{
		Name:                      api.PtrString("ids-1"),
		Id:                        api.PtrInt64(-1),
		ElementId:                 api.PtrString("air"),
		ProviderUrl:               api.PtrString("ldap://192.168.0.97:389"),
		SecurityPrincipal:         api.PtrString("CN=Administrator,CN=Users,DC=mycompany,DC=com"),
		SecurityCredential:        api.PtrString("@WSX3edc"),
		UsersCtxDN:                api.PtrString("CN=Users,DC=mycompany,DC=com"),
		RolesCtxDN:                api.PtrString("CN=Users,DC=mycompany,DC=com"),
		UidAttributeID:            api.PtrString("member"),
		PrincipalUidAttributeID:   api.PtrString("sAMAccountName"),
		RoleAttributeID:           api.PtrString("sAMAccountName"),
		Referrals:                 api.PtrString("follow"),
		LdapSearchScope:           api.PtrString("subtree"),
		InitialContextFactory:     api.PtrString("true"),
		RoleMatchingMode:          api.PtrString("manager"),
		UserPropertiesQueryString: api.PtrString("space"),
		SecurityAuthentication:    api.PtrString("authenticated"),
	} // Modifi "CreateIdSourceLdap" Because not accept (orig := api.NewLdapIdentitySourceDTO())
	listOfCreated[0], _ = s.client.CreateIdSourceLdap(*appliance.Name, element1)

	element2 := api.LdapIdentitySourceDTO{
		Name:                      api.PtrString("ids-2"),
		Id:                        api.PtrInt64(-1),
		ElementId:                 api.PtrString("air"),
		ProviderUrl:               api.PtrString("ldap://192.168.0.97:389"),
		SecurityPrincipal:         api.PtrString("CN=Administrator,CN=Users,DC=mycompany,DC=com"),
		SecurityCredential:        api.PtrString("@WSX3edc"),
		UsersCtxDN:                api.PtrString("CN=Users,DC=mycompany,DC=com"),
		RolesCtxDN:                api.PtrString("CN=Users,DC=mycompany,DC=com"),
		UidAttributeID:            api.PtrString("member"),
		PrincipalUidAttributeID:   api.PtrString("sAMAccountName"),
		RoleAttributeID:           api.PtrString("sAMAccountName"),
		Referrals:                 api.PtrString("follow"),
		LdapSearchScope:           api.PtrString("subtree"),
		InitialContextFactory:     api.PtrString("true"),
		RoleMatchingMode:          api.PtrString("manager"),
		UserPropertiesQueryString: api.PtrString("space"),
		SecurityAuthentication:    api.PtrString("authenticated"),
	}
	listOfCreated[1], _ = s.client.CreateIdSourceLdap(*appliance.Name, element2)

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
			return strings.Compare(*listOfRead[i].Name, *listOfRead[j].Name) > 0
		},
	)

	// Validate each element from the list of created with the list of read
	for idx, r := range listOfCreated {
		if err = IdSValidateUpdate(&r, &listOfRead[idx]); err != nil {
			t.Error(err)
			return
		}
	}

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
			expected: e.Name,
			received: r.Name,
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
			name:     "elementid",
			cmp:      func() bool { return StrPtrEquals(e.ElementId, r.ElementId) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "providerurl",
			cmp:      func() bool { return StrPtrEquals(e.ProviderUrl, r.ProviderUrl) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "securityprincipal",
			cmp:      func() bool { return StrPtrEquals(e.SecurityPrincipal, r.SecurityPrincipal) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "usersctxdn",
			cmp:      func() bool { return StrPtrEquals(e.UsersCtxDN, r.UsersCtxDN) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "rolesctxdn",
			cmp:      func() bool { return StrPtrEquals(e.RolesCtxDN, r.RolesCtxDN) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "uidattributeid",
			cmp:      func() bool { return StrPtrEquals(e.UidAttributeID, r.UidAttributeID) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "principaluiattributeid",
			cmp:      func() bool { return StrPtrEquals(e.PrincipalUidAttributeID, r.PrincipalUidAttributeID) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "roleattributeid",
			cmp:      func() bool { return StrPtrEquals(e.RoleAttributeID, r.RoleAttributeID) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "referrals",
			cmp:      func() bool { return StrPtrEquals(e.Referrals, r.Referrals) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "ldapsearchscope",
			cmp:      func() bool { return StrPtrEquals(e.LdapSearchScope, r.LdapSearchScope) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "initialcontextfactory",
			cmp:      func() bool { return StrPtrEquals(e.InitialContextFactory, r.InitialContextFactory) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "rolematchingmode",
			cmp:      func() bool { return StrPtrEquals(e.RoleMatchingMode, r.RoleMatchingMode) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "userpropertiesquerystring",
			cmp:      func() bool { return StrPtrEquals(e.UserPropertiesQueryString, r.UserPropertiesQueryString) },
			expected: e.Name,
			received: r.Name,
		},
		{
			name:     "securityauthentication",
			cmp:      func() bool { return StrPtrEquals(e.SecurityAuthentication, r.SecurityAuthentication) },
			expected: e.Name,
			received: r.Name,
		},
	}

	return append(t, IdSFieldTestCreate(e, r)...)
}

// Compares the expected IdP with the received one.
func CreateIdSourceLdap(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) error {

	return ValidateFields(IdSFieldTestCreate(e, r))
}

// Compares the expected IdP with the received one.
func IdSValidateUpdate(
	e *api.LdapIdentitySourceDTO,
	r *api.LdapIdentitySourceDTO) error {

	return ValidateFields(IdSFieldTestUpdate(e, r))
}
