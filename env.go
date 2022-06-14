package defstruct

import (
	"errors"
	"os"
	"reflect"
)

const tag_env = "env"

// Sets ENV value for scalar types from tags.
// Returns the error.
func SetEnvFromTags(structure interface{}) error {

	v := reflect.ValueOf(structure)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("Invalid pointer. The parameter must be a pointer to structure.")
	}

	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return errors.New("The pointer value is not a structure.")
	}

	for i := 0; i < t.NumField(); i++ {

		tfield := t.Field(i)
		vfield := v.Field(i)

		if envName, tagIsset := tfield.Tag.Lookup(tag_env); tagIsset {

			if envValue, envIsset := os.LookupEnv(envName); envIsset {

				if err := setField(vfield, envValue); err != nil {
					return err
				}
			}
		}

		switch vfield.Kind() {
		case reflect.Struct:

			SetEnvFromTags(vfield.Addr().Interface())

		case reflect.Ptr:

			if vfield.Addr().Elem().Type().Elem().Kind() == reflect.Struct {
				SetEnvFromTags(vfield.Interface())
			}
		}
	}

	return nil
}
