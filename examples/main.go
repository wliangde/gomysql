package main 
import(
"log"
"github.com/biswarupadhikari/gomysql"
)
func main() {
	db:=gomysql.Connect("root","rootwdp","go")
	db.Select("id,username,email")
	log.Println("HI",db)
}