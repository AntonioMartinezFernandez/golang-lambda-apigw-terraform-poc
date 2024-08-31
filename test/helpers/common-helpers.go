package helpers

import (
	"errors"
	"fmt"
	"path"
	"time"

	json_schema "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/json-schema"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"

	"github.com/bxcodec/faker/v3"
)

var errNotImplemented = errors.New("not implemented")

func Validator() json_schema.JsonSchemaValidator {
	validatorBasePath := fmt.Sprintf("%s/", path.Join("./schemas"))
	return json_schema.NewJsonSchemaValidator(validatorBasePath)
}

func RandomBool() bool {
	a, _ := faker.RandomInt(0, 1, 1)
	return a[0] == 1
}

func EmptyHeaders() map[string]string {
	return map[string]string{}
}

func BearerHeader(token string) map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
}

type FakeDto struct{}

func (f *FakeDto) Id() string { return "fake" }

func RandomInt() int {
	a, _ := faker.RandomInt(1, 1000000, 1)
	return a[0]
}

func RandomName() string {
	return faker.Name()
}

func RandomTimeRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func RandomIntOrNil() *int {
	var result *int
	if RandomBool() {
		return result
	}

	result = utils.Ptr(RandomInt())

	return result
}

func RandomStringOrNil() *string {
	var result *string
	if RandomBool() {
		return result
	}

	result = utils.Ptr(faker.Word())

	return result
}

func RandomStringIntOrNull() string {
	number := RandomIntOrNil()
	if number == nil {
		return "null"
	}
	return fmt.Sprintf("%d", *number)
}

func RandomStringOrNull() string {
	word := RandomStringOrNil()
	if word == nil {
		return "null"
	}
	return *word
}
