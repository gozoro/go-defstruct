package defstruct

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

const tag_default = "default"

// Sets default value for scalar types from tags.
// Returns the error.
func SetDefaultFromTags(structure interface{}) error {

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

		if defaultVal, tagIsset := tfield.Tag.Lookup(tag_default); tagIsset && vfield.IsZero() {

			if err := setField(vfield, defaultVal); err != nil {
				return err
			}
		}

		switch vfield.Kind() {
		case reflect.Struct:

			SetDefaultFromTags(vfield.Addr().Interface())

		case reflect.Ptr:

			if vfield.Addr().Elem().Type().Elem().Kind() == reflect.Struct {
				SetDefaultFromTags(vfield.Interface())
			}
		}
	}

	return nil
}

func setField(field reflect.Value, value string) error {

	if !field.CanSet() {
		return nil
	}

	switch field.Kind() {

	case reflect.String:
		field.SetString(value)

	case reflect.Bool:
		if val, err := strconv.ParseBool(value); err == nil {
			field.SetBool(val)
		}

	case reflect.Int:
		if val, err := strconv.ParseInt(value, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}

	case reflect.Int8:
		if val, err := strconv.ParseInt(value, 0, 8); err == nil {
			field.Set(reflect.ValueOf(int8(val)).Convert(field.Type()))
		}

	case reflect.Int16:
		if val, err := strconv.ParseInt(value, 0, 16); err == nil {
			field.Set(reflect.ValueOf(int16(val)).Convert(field.Type()))
		}

	case reflect.Int32:
		if val, err := strconv.ParseInt(value, 0, 32); err == nil {
			field.Set(reflect.ValueOf(int32(val)).Convert(field.Type()))
		}

	case reflect.Int64:

		fieldType := fmt.Sprintf("%v", field.Type())

		if fieldType == "time.Duration" {

			if val, err := time.ParseDuration(value); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}

		} else {

			if val, err := strconv.ParseInt(value, 0, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		}

	case reflect.Uint:
		if val, err := strconv.ParseUint(value, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(uint(val)).Convert(field.Type()))
		}

	case reflect.Uint8:
		if val, err := strconv.ParseUint(value, 0, 8); err == nil {
			field.Set(reflect.ValueOf(uint8(val)).Convert(field.Type()))
		}

	case reflect.Uint16:
		if val, err := strconv.ParseUint(value, 0, 16); err == nil {
			field.Set(reflect.ValueOf(uint16(val)).Convert(field.Type()))
		}

	case reflect.Uint32:
		if val, err := strconv.ParseUint(value, 0, 32); err == nil {
			field.Set(reflect.ValueOf(uint32(val)).Convert(field.Type()))
		}
	case reflect.Uint64:
		if val, err := strconv.ParseUint(value, 0, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	case reflect.Uintptr:
		if val, err := strconv.ParseUint(value, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(uintptr(val)).Convert(field.Type()))
		}

	case reflect.Float32:
		if val, err := strconv.ParseFloat(value, 32); err == nil {
			field.Set(reflect.ValueOf(float32(val)).Convert(field.Type()))
		}

	case reflect.Float64:
		if val, err := strconv.ParseFloat(value, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	}

	return nil
}
