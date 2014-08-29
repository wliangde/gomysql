package main 
import(
//"log"
"github.com/biswarupadhikari/gomysql"
)
func main() {
	db:=gomysql.Connect("root","rootwdp","go")
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
	// data:=make(map[string]interface{})
	
	// data["username"]="joyantak"
	// data["password"]="ttttt"
	// db.Table("users").Insert(data)
	
	/**
	 * Update Table
	 */
	data:=make(map[string]interface{})
	
	data["username"]="new Username"
	data["password"]="new Password"
	db.Table("users").Where("id","=",25).Update(data)
}