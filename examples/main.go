package main

import (
	"github.com/biswarupadhikari/gomysql"
	"log"
)

func main() {
	log.Println("GoMysql Testing App")
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
	// db.Table("users").Where("id", "=", 155).Delete()
	/**
	 * Select All Records
	 */
	// rows := db.Select("id,username,password").From("users").Where("id", ">", 7).Get()
	// log.Println(rows)
	// log.Println(db.Select("id").Get())
	// //log.Println(db.Get())
	// //log.Println(db.Get())
	/**
	 * Left Join
	//  */
	// query := db.GetQuery()
	// query.Select("*")
	// query.From("users as u")
	// query.Join("LEFT", "roles as r", "u.role_id=r.id")
	// rows := query.Get()
	// log.Println(rows)
	/**
	 * Left Join With Condition
	 */

	// query := db.GetQuery()
	// query.Select("*")
	// query.From("users as u")
	// query.Join("LEFT", "roles as r", "u.role_id=r.id AND u.id>?",7)
	// query.Where("u.username","=","biswarupadhikari")
	// sql,params:=query.GetSQL()
	// log.Println(sql,params)
	//
	/**
	 * Get Insert SQL
	 */
	// query:=db.GetQuery()
	// query.Table("users")
	// data:=make(map[string]interface{})
	// data["username"]="asdas"
	// data["password"]="xxxxx"
	// sql,params:=query.InsertSQL(data)
	// log.Println(sql,params)
	/**
	  * Get Update SQL
	//   */
	// query:=db.GetQuery()
	// query.Table("users")
	// query.Where("id","=",9)
	// query.Where("username","=","Biswu")
	// data:=make(map[string]interface{})
	// data["username"]="asdas"
	// data["password"]="xxxxx"
	// sql,params:=query.UpdateSQL(data)
	// log.Println(sql,params)
	/**
	 * Get Delete SQL
	 */
	query := db.GetQuery()
	query.Table("users")
	query.Where("id", "=", 9)
	query.Where("username", "=", "Biswu")
	sql, params := query.DeleteSQL()
	log.Println(sql, params)
}
