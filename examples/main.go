package main 
import(
"log"
"github.com/biswarupadhikari/godb"
)
func main() {
	db:=new(godb.GoDB)
	db.Connect("root","rootwdp","go")
	log.Println("HI",db)
}