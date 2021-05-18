GENERATOR=/opt/atricore/tools/openapitools/openapi-generator-cli.sh
SWAGGER_FILE=~/.m2/repository/com/atricore/idbus/console/console-api/1.4.3-SNAPSHOT/console-api-1.4.3-SNAPSHOT-swagger.yaml
PGK_NAME=jossoappi

default: build

build:
	go install

dep: # Download required dependencies
	go mod tidy
	go mod vendor

generate:
	$(GENERATOR) generate -i $(SWAGGER_FILE) -g go -o . --additional-properties=packageName=$(PGK_NAME) --additional-properties=disallowAdditionalPropertiesIfNotPresent=false --git-repo-id=josso-api-go --git-user-id=atricore
