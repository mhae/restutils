package restutils

import (
	"reflect"
	"fmt"
)

// Simple mapper function to convert between domain objects such as REST and Service layer
// TODO: Improve performance
func Mapper(src interface{}, dest interface{})  {

	st := reflect.TypeOf(src).Elem()
	dv := reflect.ValueOf(dest).Elem()
	sv := reflect.ValueOf(src).Elem()

	if st.Kind() != reflect.Struct || dv.Kind() != reflect.Struct {
		return
	}

	for i:=0; i<st.NumField(); i++ {

		df := dv.FieldByName(st.Field(i).Name)

		if !(df.IsValid() && df.CanSet()) {
			continue
		}

		if df.Kind() != st.Field(i).Type.Kind() {
			continue
		}

		if df.Kind() == reflect.Struct {
			mapEmbeddedField(sv.Field(i), df)
		} else {

			df.Set(sv.Field(i))
		}
	}
}

func mapEmbeddedField(sv reflect.Value, dv reflect.Value) {

	st := sv.Type()
	for i:=0; i<st.NumField(); i++ {

		df := dv.FieldByName(st.Field(i).Name)

		if !(df.IsValid() && df.CanSet()) {
			continue
		}

		if df.Kind() != st.Field(i).Type.Kind() {
			continue
		}

		if df.Kind() == reflect.Struct {
			fmt.Println("Embedded struct: sv=", sv.Field(i), " df=", df)
			mapEmbeddedField(sv.Field(i), df)
		} else {

			df.Set(sv.Field(i))
		}
	}

}
