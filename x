go mod tidy
go mod vendor
==> Checking that code complies with gofmt requirements...
go clean -testcache
TF_ACC=1 go test $(go list ./... |grep -v 'vendor') -v   -timeout 120m
