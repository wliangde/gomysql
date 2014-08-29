package main 
import(
//"log"
"github.com/biswarupadhikari/gomysql"
)
func main() {
	db:=gomysql.Connect("root","rootwdp","go")
	db.Select("id,username as p,email").From("users").Where("id",">",15)
	db.Where("username","=","biswarupadhikari")
	db.ORWhere("id","<",150)
	db.RawWhere(" NOR password=? ","taste")
	db.RawWhere(" XOR password=? ",15)
	db.Get()
}