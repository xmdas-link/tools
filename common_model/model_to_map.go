package common_model

import (
	"reflect"
)

func ModelToMap(data interface{}) map[string]interface{} {

	var (
		m      = map[string]interface{}{}
		types  = reflect.TypeOf(data)
		values = reflect.ValueOf(data)
	)

	for i := 0; i < types.NumField(); i++ {
		m[types.Field(i).Name] = values.Field(i).Interface()
	}

	return m
}

func ModelToMapAndIgnore(data interface{}, ignores []string) map[string]interface{} {
	var (
		m      = map[string]interface{}{}
		types  = reflect.TypeOf(data)
		values = reflect.ValueOf(data)
	)

	for i := 0; i < types.NumField(); i++ {
		ignore := false
		for _, n := range ignores {
			if types.Field(i).Name == n {
				ignore = true
			}
		}
		if ignore {
			continue
		}
		m[types.Field(i).Name] = values.Field(i).Interface()
	}

	return m
}
