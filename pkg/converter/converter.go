package converter

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"time"
)

type TagMap map[string]reflect.Value

func ConvertSlice[A any, B any](from []A) ([]B, error) {
	if len(from) == 0 {
		return nil, nil
	}

	newB := make([]B, 0, len(from))

	fromTagMap := tagMap(from[0])

	for _, valueA := range from {
		valueB, err := convertStructWithTagMap[A, B](valueA, fromTagMap)
		if err != nil {
			return nil, err
		}
		newB = append(newB, valueB)
	}

	return newB, nil
}

func ConvertPointerSlice[A any, B any](from []*A) ([]B, error) {
	if len(from) == 0 {
		return nil, nil
	}

	newB := make([]B, 0, len(from))

	fromTagMap := tagMap(from[0])

	for _, valueA := range from {
		valueB, err := convertStructWithTagMap[A, B](*valueA, fromTagMap)
		if err != nil {
			return nil, err
		}
		newB = append(newB, valueB)
	}

	return newB, nil
}

func ConvertToPointerSlice[A any, B any](from []A) ([]*B, error) {
	if len(from) == 0 {
		return nil, nil
	}

	newB := make([]*B, 0, len(from))

	fromTagMap := tagMap(from[0])

	for _, valueA := range from {
		valueB, err := convertStructWithTagMap[A, B](valueA, fromTagMap)
		if err != nil {
			return nil, err
		}
		newB = append(newB, &valueB)
	}

	return newB, nil
}

func ConvertPointerStruct[A any, B any](from *A) (B, error) {
	return ConvertStruct[A, B](*from)
}

// ConvertStruct convert from struct A to struct B
func ConvertStruct[A any, B any](from A) (B, error) {
	var newValue B

	aValue := reflect.ValueOf(&from)

	bValue := reflect.ValueOf(&newValue).Elem()

	if err := convertStruct(aValue, bValue, tagMap(from)); err != nil {
		return newValue, err
	}

	return newValue, nil
}

func convertStructWithTagMap[A any, B any](from A, fromTagMap TagMap) (B, error) {
	var newValue B

	aValue := reflect.ValueOf(&from)

	bValue := reflect.ValueOf(&newValue).Elem()

	if err := convertStruct(aValue, bValue, fromTagMap); err != nil {
		return newValue, err
	}

	return newValue, nil
}

func convertStruct(from reflect.Value, to reflect.Value, fromTagMap TagMap) error {

	bType := to.Type()

	for idx := 0; idx < bType.NumField(); idx++ {
		bField := bType.Field(idx)
		tagName, ok := bField.Tag.Lookup("convert")

		if value, has := fromTagMap[tagName]; ok && has {
			err := setValue(to.Field(idx), value)
			if err != nil {
				return err
			}

			continue
		}

		aFieldValue := reflect.Indirect(from).FieldByName(bField.Name)
		if aFieldValue.IsValid() {
			err := setValue(to.Field(idx), aFieldValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setIfTime(dst, src reflect.Value) (bool, error) {
	if dst.Type().String() == "string" && src.Type().String() == "time.Time" {
		srcTime := src.Interface().(time.Time)
		dst.SetString(srcTime.Format(time.RFC3339))
		return true, nil
	}

	if dst.Type().String() == "time.Time" && src.Type().String() == "string" {
		dstTime, err := time.Parse(time.RFC3339, src.String())
		if err != nil {
			return false, err
		}

		dst.Set(reflect.ValueOf(dstTime))
		return true, nil
	}

	if dst.Type().String() == "string" && src.Type().String() == "*time.Time" {
		if src.IsNil() {
			return true, nil
		}
		srcTime := reflect.Indirect(src).Interface().(time.Time)
		dst.SetString(srcTime.Format(time.RFC3339))
		return true, nil
	}

	if dst.Type().String() == "*time.Time" && src.Type().String() == "string" {
		dstTime, err := time.Parse(time.RFC3339, src.String())
		if err != nil {
			return false, err
		}

		dst.Set(reflect.ValueOf(&dstTime))
		return true, nil
	}

	return false, nil
}

func nullable(dst, src reflect.Value) bool {
	if strings.TrimPrefix(dst.Type().String(), "*") != strings.TrimPrefix(src.Type().String(), "*") {
		return false
	}

	if dst.Kind() == reflect.Pointer {
		if src.IsZero() {
			return true
		}

		dst.Set(src.Addr())
		return true
	}

	// src is a pointer
	if src.IsNil() {
		return true
	}

	dst.Set(reflect.Indirect(src))
	return true
}

func similarType(dst, src reflect.Value) bool {
	if !src.CanConvert(dst.Type()) {
		return false
	}

	dst.Set(src.Convert(dst.Type()))
	return true
}

func setValue(dst, src reflect.Value) error {
	if dst.Type() == src.Type() {
		dst.Set(src)
		return nil
	}

	if ok, err := setIfTime(dst, src); ok || err != nil {
		return err
	}

	if nullable(dst, src) {
		return nil
	}

	if similarType(dst, src) {
		return nil
	}

	return errors.Errorf("unkhown case: %s %s", src.Type(), dst.Type())
}

func tagMap(src any) TagMap {
	tags := TagMap{}
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)

	for idx := 0; idx < srcType.NumField(); idx++ {
		srcField := srcType.Field(idx)
		tagName, ok := srcField.Tag.Lookup("convert")
		if !ok {
			continue
		}
		tags[tagName] = srcValue.Field(idx)
	}

	return tags
}
