module github.com/atricore/josso-sdk-go

go 1.13

replace github.com/atricore/josso-api-go => ../josso-api-go

require (
	github.com/atricore/josso-api-go v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
)
