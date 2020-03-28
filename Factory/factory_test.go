package Factory

import (
	"fmt"
	"testing"
)

func TestDatabaseFactory(t *testing.T) {
	env1:="production"
	env2:="development"

	db1:= DatabaseFactory(env1)
	db2:= DatabaseFactory(env2)

	db1.PutData("test","this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test","this is sqlite")
	fmt.Println(db2.GetData("test"))
}

func TestAbstractFactory(t *testing.T) {
	env:="production"
	fs := AbstractFactory("filesystem")
	db:= AbstractFactory("database")

	d:= db(env).(Database)
	f:= fs(env).(FileSystem)

	fmt.Println(d)
	fmt.Println(f)
}