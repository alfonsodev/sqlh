package sqlh

import (
	"fmt"
	"testing"
)

type User struct {
	Id   int // `db:"user_id"`
	Name string
	Age  int
}

func (u *User) Update() {
	names, _ := StructToKeyValue(u)
	sql := "UPDATE myschema.Users SET "
	sql += names
	sql += " WHERE myschema.Users = $1"
	fmt.Println(sql)
}

func (u *User) Insert() {
	keys, values, _ := StructListKeys(u)
	sql := "INSERT INTO " + keys
	sql += " VALUES(" + values + ") "
	fmt.Println(sql)
}

func TestReflectSql(t *testing.T) {
	u := User{Name: "Mikel", Age: 32}
	u.Update()
}

// func TestStructListKeys(t *testing.T) {
// 	u := User{Name: "jhon", Age: 32}
// 	u.Insert()
// }
