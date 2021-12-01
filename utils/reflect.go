package utils

import (
	"reflect"
	"strings"
)

/*
	q must be a pointer to the structure
	For each string type field in the structure, execute TrimSpace function
 */
func TrimChildStr(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Ptr {
		v := reflect.Indirect(reflect.ValueOf(q))
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				if v.Field(i).Kind().String() == "string" {
					s := strings.TrimSpace(v.Field(i).String())
					v.Field(i).SetString(s)
				}
			}
		}
	}
}

