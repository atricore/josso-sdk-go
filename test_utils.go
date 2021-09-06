package cli

import (
	api "github.com/atricore/josso-api-go"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type FiledTestStruct struct {
	name     string
	cmp      func() bool
	expected string
	received string
}

func ValidateField(f FiledTestStruct) error {
	var err error
	if !f.cmp() {
		err = errors.Errorf("invalid %s, expected [%s],  received[%s]",
			f.name, f.expected, f.received)
	}
	return err
}

func ValidateFields(fts []FiledTestStruct) error {
	var result error
	for _, ft := range fts {
		//fmt.Printf("ValidateField: %s=%t\n", ft.name, ft.cmp())
		if !ft.cmp() {
			err := errors.Errorf("invalid %s, expected [%s],  received[%s]",
				ft.name, ft.expected, ft.received)
			result = multierror.Append(result, err)
		}
	}

	return result
}

func LocationPtrEquals(a *api.LocationDTO, b *api.LocationDTO) bool {
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return LocationEquals(*a, *b)
}

func LocationEquals(a api.LocationDTO, b api.LocationDTO) bool {
	return StrPtrEquals(a.Protocol, b.Protocol) &&
		StrPtrEquals(a.Host, b.Host) &&
		Int32PtrEquals(a.Port, b.Port) &&
		StrPtrEquals(a.Context, b.Context) &&
		StrPtrEquals(a.Uri, b.Uri)
}

func Int32PtrEquals(a *int32, b *int32) bool {
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return *a == *b
}

func Int64PtrEquals(a *int64, b *int64) bool {
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return *a == *b
}

// Compares if ptrs are nil, then compares values
func StrPtrEquals(a *string, b *string) bool {

	// a == nil means be must b nil
	if a == nil {
		return b == nil
	}

	// a != nil
	if b == nil {
		return false
	}

	return *a == *b
}

// Compares if ptrs are nil, then compares values
func BoolPtrEquals(a *bool, b *bool) bool {

	// a == nil means be must b nil
	if a == nil {
		return b == nil
	}

	// a != nil
	if b == nil {
		return false
	}

	return *a == *b
}

func SubjectNamePolicyContains(needle api.SubjectNameIdentifierPolicyDTO, haystack *[]api.SubjectNameIdentifierPolicyDTO) bool {
	for _, matchValue := range *haystack {
		if SubNameEquals(&matchValue, &needle) {
			return true
		}
	}
	return false
}

func SubjectNamePolicyEquals(a api.SubjectNameIdentifierPolicyDTO, b api.SubjectNameIdentifierPolicyDTO) bool {
	return StrPtrEquals(a.DescriptionKey, b.DescriptionKey) &&
		StrPtrEquals(a.Id, b.Id) &&
		StrPtrEquals(a.Name, b.Name) &&
		StrPtrEquals(a.SubjectAttribute, b.SubjectAttribute) &&
		StrPtrEquals(a.Type, b.Type)

}

// Returns true if a and b have the same elements
func SubNamePolEquals(a *[]api.SubjectNameIdentifierPolicyDTO, b *[]api.SubjectNameIdentifierPolicyDTO) bool {
	// Check both a and b are nil -> true

	if *a == nil {
		return *b == nil
	}

	if *b == nil {
		return false
	}

	if len(*a) != len(*b) {
		return false
	}
	for _, value := range *a {
		if !SubjectNamePolicyContains(value, b) {
			return false
		}
	}
	return true

}

func SubNameEquals(a *api.SubjectNameIdentifierPolicyDTO, b *api.SubjectNameIdentifierPolicyDTO) bool {
	// Compare each field of a wiht b:
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return SubNameEquals(a, b)

}

//=====================================================================

func ProviderConfigContains(needle api.ProviderConfigDTO, haystack *[]api.ProviderConfigDTO) bool {
	for _, matchValue := range *haystack {
		if ProviderconfEquals(&matchValue, &needle) {
			return true
		}
	}
	return false
}

func ProvConfEquals(a api.ProviderConfigDTO, b api.ProviderConfigDTO) bool {
	return StrPtrEquals(a.Description, b.Description) &&
		StrPtrEquals(a.DisplayName, b.DisplayName) &&
		StrPtrEquals(a.ElementId, b.ElementId) &&
		Int64PtrEquals(a.Id, b.Id) &&
		StrPtrEquals(a.Name, b.Name)

}

// Returns true if a and b have the same elements
func ProvConfigEquals(a *[]api.ProviderConfigDTO, b *[]api.ProviderConfigDTO) bool {
	// Check both a and b are nil -> true

	if *a == nil {
		return *b == nil
	}

	if *b == nil {
		return false
	}

	if len(*a) != len(*b) {
		return false
	}
	for _, value := range *a {
		if !ProviderConfigContains(value, b) {
			return false
		}
	}
	return true

}

func ProviderconfEquals(a *api.ProviderConfigDTO, b *api.ProviderConfigDTO) bool {
	// Compare each field of a wiht b:
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return ProviderconfEquals(a, b)

}

//==================================

func IdentitylokContains(needle api.IdentityLookupDTO, haystack *[]api.IdentityLookupDTO) bool {
	for _, matchValue := range *haystack {
		if IdentitylokEquals(&matchValue, &needle) {
			return true
		}
	}
	return false
}

func IdentitylookupEquals(a api.IdentityLookupDTO, b api.IdentityLookupDTO) bool {
	return StrPtrEquals(a.Description, b.Description) &&
		StrPtrEquals(a.ElementId, b.ElementId) &&
		Int64PtrEquals(a.Id, b.Id) &&
		//StrPtrEquals(a.IdentitySource, b.IdentitySource) &&
		StrPtrEquals(a.Name, b.Name)
	//StrPtrEquals(a.Provider, b.Provider) &&
	//StrPtrEquals(a.Waypoints, b.Waypoints)

}

// Returns true if a and b have the same elements
func IdentitylookEquals(a *[]api.IdentityLookupDTO, b *[]api.IdentityLookupDTO) bool {
	// Check both a and b are nil -> true

	if *a == nil {
		return *b == nil
	}

	if *b == nil {
		return false
	}

	if len(*a) != len(*b) {
		return false
	}
	for _, value := range *a {
		if !IdentitylokContains(value, b) {
			return false
		}
	}
	return true

}

func IdentitylokEquals(a *api.IdentityLookupDTO, b *api.IdentityLookupDTO) bool {
	// Compare each field of a wiht b:
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return IdentitylokEquals(a, b)

}

// =======================================================================================

func DelegatedAuthenticationDTOEquals(a api.DelegatedAuthenticationDTO, b api.DelegatedAuthenticationDTO) bool {
	return StrPtrEquals(a.Description, b.Description) &&
		//DelegatedAuthenticationEquals(a.AuthenticationService, b.AuthenticationService) &&
		StrPtrEquals(a.ElementId, b.ElementId) &&
		Int64PtrEquals(a.Id, b.Id) &&
		//StrPtrEquals(a.Idp, b.Idp)
		StrPtrEquals(a.Name, b.Name)
	//StrPtrEquals(a.Waypoints, b.Waypoints)

}

func AuthenticationMechanismDTOContains(needle api.AuthenticationMechanismDTO, haystack *[]api.AuthenticationMechanismDTO) bool {
	for _, matchValue := range *haystack {
		if AuthenticationMechanismEquals(&matchValue, &needle) {
			return true
		}
	}
	return false
}

func AuthMechanismDTOEquals(a api.AuthenticationMechanismDTO, b api.AuthenticationMechanismDTO) bool {
	return StrPtrEquals(a.DisplayName, b.DisplayName) &&
		//DelegatedAuthenticationDTOEquals(a.DelegatedAuthentication, b.DelegatedAuthentication) &&
		StrPtrEquals(a.ElementId, b.ElementId) &&
		Int64PtrEquals(a.Id, b.Id) &&
		StrPtrEquals(a.Name, b.Name) &&
		Int32PtrEquals(a.Priority, b.Priority)

}

// Returns true if a and b have the same elements
func AuthenticationMechanismsDTOEquals(a *[]api.AuthenticationMechanismDTO, b *[]api.AuthenticationMechanismDTO) bool {
	// Check both a and b are nil -> true

	if *a == nil {
		return *b == nil
	}

	if *b == nil {
		return false
	}

	if len(*a) != len(*b) {
		return false
	}
	for _, value := range *a {
		if !AuthenticationMechanismDTOContains(value, b) {
			return false
		}
	}
	return true

}

func AuthenticationMechanismEquals(a *api.AuthenticationMechanismDTO, b *api.AuthenticationMechanismDTO) bool {
	// Compare each field of a wiht b:
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return AuthenticationMechanismEquals(a, b)

}

func OAuth2ClientContains(needle api.OAuth2ClientDTO, haystack *[]api.OAuth2ClientDTO) bool {
	for _, matchValue := range *haystack {
		if OAuth2ClientEquals(matchValue, needle) {
			return true
		}
	}
	return false
}

func OAuth2ClientEquals(a api.OAuth2ClientDTO, b api.OAuth2ClientDTO) bool {
	return StrPtrEquals(a.BaseURL, b.BaseURL) &&
		StrPtrEquals(a.Id, b.Id) &&
		StrPtrEquals(a.Secret, b.Secret)

}

// Returns true if a and b have the same elements
func Oauth2ClientsEquals(a *[]api.OAuth2ClientDTO, b *[]api.OAuth2ClientDTO) bool {
	// Check both a and b are nil -> true

	if *a == nil {
		return *b == nil
	}

	if *b == nil {
		return false
	}

	if len(*a) != len(*b) {
		return false
	}
	for _, value := range *a {
		if !OAuth2ClientContains(value, b) {
			return false
		}
	}
	return true
}

func Oauth2ClientEquals(a *api.OAuth2ClientDTO, b *api.OAuth2ClientDTO) bool {
	// Compare each field of a wiht b:
	if a == nil {
		return b == nil
	}

	if b == nil {
		return false
	}

	return StrPtrEquals(a.Secret, b.Secret)

}
