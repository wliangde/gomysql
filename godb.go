package godb
import(
"log"
//"strings"
"database/sql"
_"github.com/go-sql-driver/mysql"
)
/**
 * GoDB Structure
 */
 type GoDB struct{	
 	db *sql.DB
 	fields []string
 	tableName string
 	conditions []string
 	dataValues []interface{}
 }
/**
 * Connect To DataBase
 */
 func (godb *GoDB)Connect(dbUsername string,dbPassword string,dbName string)(*GoDB){
 	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
 	if err != nil {
 		log.Fatal(err)
 	}
 	err = db.Ping()
 	if err != nil {
 		log.Fatal(err)
 	}
 	godb.db=db
 	return godb
 }
