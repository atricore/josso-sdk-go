package cli

import (
	api "github.com/atricore/josso-api-go"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type FiledTestStruct struct {
	name     string
	cmp      func() bool
	expected *string
	received *string
}

func ValidateField(f FiledTestStruct) error {
	var err error
	if !f.cmp() {
		err = errors.Errorf("invalid %s, expected [%s],  received[%s]",
			f.name, StrDeref(f.expected), StrDeref(f.received))
	}
	return err
}

func ValidateFields(fts []FiledTestStruct) error {
	var result error
	for _, ft := range fts {
		if !ft.cmp() {
			err := errors.Errorf("invalid %s, expected [%s],  received[%s]",
				ft.name, StrDeref(ft.expected), StrDeref(ft.received))
			multierror.Append(result, err)
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

// Returns true if a and b have the same elements
func Oauth2ClientsEquals(a *[]api.OAuth2ClientDTO, b *[]api.OAuth2ClientDTO) bool {
	// Check both a and b are nil -> true

	// TODO  code

	// If len(a) != len(b) -> false

	// TODO code

	// We know both lists have values and the same lenght
	// Compare elements, list may have different sort

	// TODO code use Oauth2ClientEquals

	return false // TODO

}

func Oauth2ClientEquals(a *api.OAuth2ClientDTO, b api.OAuth2ClientDTO) bool {
	// Compare each field of a wiht b:
	return StrPtrEquals(a.Secret, b.Secret) // TODO : add more fields: BaseUrl and Id

}
