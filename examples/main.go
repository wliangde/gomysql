package main

import (
	//"log"
	"github.com/biswarupadhikari/gomysql"
)

func main() {
	db := gomysql.Connect("root", "rootwdp", "go")
	/**
	 * Select Records
	 */
	// db.Select("id,username as p,email").From("users").Where("id",">",15)
	// db.Where("username","=","biswarupadhikari")
	// db.ORWhere("id","<",150)
	// db.RawWhere(" NOR password=? ","taste")
	// db.RawWhere(" XOR password=? ",15)
	// db.Get()

	/**
	 * Insert Record
	 */
	// data := make(map[string]interface{})

	// data["username"] = "samirgoswami"
	// data["password"] = "9514753"
	// db.Table("users").Insert(data)

	/**
	 * Update Table
	 */
	// data := make(map[string]interface{})

	// data["username"] = "samirgoswami Modified"
	// data["password"] = "samirgoswami pass Modified"
	// db.Table("users").Where("username", "=", "samirgoswami").Update(data)

	/**
	 * Delete Records
	 */
	db.Table("users").Where("id", "=", 155).Delete()
}
