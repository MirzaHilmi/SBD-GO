package fakers

import (
	"reflect"

	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
	"pgregory.net/rand"
)

func LoadCustomFaker() {
	_ = faker.AddProvider("ulid", func(v reflect.Value) (interface{}, error) {
		return ulid.Make().String(), nil
	})

	_ = faker.AddProvider("gender", func(v reflect.Value) (interface{}, error) {
		gender := [2]string{"MALE", "FEMALE"}
		return gender[0], nil
	})

	_ = faker.AddProvider("age", func(v reflect.Value) (interface{}, error) {
		return uint(rand.Intn(40)), nil
	})

	_ = faker.AddProvider("street", func(v reflect.Value) (interface{}, error) {
		address := faker.GetRealAddress()

		return address.Address, nil
	})

	_ = faker.AddProvider("city", func(v reflect.Value) (interface{}, error) {
		address := faker.GetRealAddress()

		return address.City, nil
	})

	_ = faker.AddProvider("postalcode", func(v reflect.Value) (interface{}, error) {
		address := faker.GetRealAddress()

		return address.PostalCode, nil
	})
}
