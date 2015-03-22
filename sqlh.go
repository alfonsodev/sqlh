package sqlh

import (
	"fmt"
	"reflect"
	"strings"
)

func StructListKeys(u interface{}) (string, string, []interface{}) {
	val := reflect.ValueOf(u).Elem()
	keys := ""
	values := ""
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		//TODO: get sql tag if exists the fieldName  tag.Get("tag_name")
		//		tag := val.Type().Field(i).Tag
		fieldName := strings.ToLower(val.Type().Field(i).Name)
		if i > 0 {
			keys += fmt.Sprintf(", %s", fieldName)
		}
		fmt.Printf(":::::%v ", valueField.Type().Name())

		v[i] = valueField.Field(0).Addr().Interface()

		// keys += fmt.Sprintf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}

	if string(keys[0]) == "," {
		keys = keys[1:]
	}

	for j := 1; j < len(v); j++ {
		values += fmt.Sprintf(", $%v", j)
	}

	values = values[1:]

	return keys, values, v[1:]
}

func StructToKeyValue(u interface{}) (string, []interface{}) {
	val := reflect.ValueOf(u).Elem()
	keys := ""
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		fmt.Printf("\n[VALUE]:%+v\n", valueField.String())
		// TODO: get sql tag if exists the fieldName  tag.Get("tag_name"
		// tag := val.Type().Field(i).Tag
		// if string(val.Type().Field(i).Type.Name()) == "int" && valueField > 0 {
		// 	fieldName := strings.ToLower(val.Type().Field(i).Name)
		// 	keys += fmt.Sprintf(", %s = $%v", fieldName, i+1)
		// 	v[i] = valueField
		// } else if val.Type().Field(i).Type == "string" && valueField != "" {
		// 	fieldName := strings.ToLower(val.Type().Field(i).Name)
		// 	keys += fmt.Sprintf(", %s = $%v", fieldName, i+1)
		// 	v[i] = valueField
		// }

		fieldName := strings.ToLower(val.Type().Field(i).Name)
		keys += fmt.Sprintf(", %s = $%v", fieldName, i+1)
		v[i] = valueField

		// if valueField.IsNil() {
		// 	fmt.Printf("\n[:o]:%s is nil \n", fieldName)
		// }
		// keys += fmt.Sprintf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))

	}

	if string(keys[0]) == "," {
		keys = keys[1:]
	}

	return keys, v
}

// func main() {
// 	var b User
// 	b.Name = "a"
// 	keys, _ := Reflect(&b)
// 	q := "UPDATE " + reflect.TypeOf(b).Name() + " SET " + keys
// 	fmt.Printf(q)
// }
