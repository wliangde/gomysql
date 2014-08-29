package gomysql
import(
"log"
"strings"
"database/sql"
_"github.com/go-sql-driver/mysql"
)
/**
 * GoMysql Structure
 */
 type GoMysql struct{	
 	db *sql.DB
 	fields []string
 	tableName string
 	conditions []string
 	dataValues []interface{}
 }
/**
 * Connect To DataBase
 */
 func Connect(dbUsername string,dbPassword string,dbName string)*GoMysql{
 	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
 	if err != nil {
 		log.Fatal(err)
 	}
 	err = db.Ping()
 	if err != nil {
 		log.Fatal(err)
 	}
 	gomysql:=new(GoMysql)
 	gomysql.db=db
 	return gomysql
 }
/**
 * Select Fields
 */
 func (gomysql *GoMysql)Select(fields string)*GoMysql{
 	cols:=strings.Split(fields,",")
 	for _,field:=range cols{
 		gomysql.fields=append(gomysql.fields,field)
 	}
 	return gomysql
 }
 /**
  * Select From a Table 
  */
  func (gomysql *GoMysql)From(tableName string)*GoMysql{
  	gomysql.tableName=tableName
  	return gomysql
  }