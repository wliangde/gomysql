package gomysql
import(
"testing"
)
var(
	DBHost="localhost"
	DBUsername="root"
	DBPassword="rootwdp"
	DBName="go"
	DBPort="3306"
)
func init(){
	//log.Println("Starting GoMysql Testing")
}
/**
 * Test Mysql Connect With 
 */
 func TestConnectionWithGoodCredentials(t *testing.T){
 	_,err:=Connect(DBHost,DBUsername,DBPassword,DBName,DBPort)
 	if err!=nil{
 		t.Error("Failed to Connect Database Using Good Credentials")
 	}else{
 		t.Log("Success Mysql Connect Passed Using Right Credentials")
 	}
 }
 /**
  * Test Mysql Connect With Wrong Credentials
  */
  func TestConnectionWithBadCredentials(t *testing.T){
 	_,err:=Connect(DBHost+"wrong",DBUsername+"wrong",DBPassword+"wrong",DBName+"wrong",DBPort+"wrong")
 	if err==nil{
 		t.Error("Not Thow Error with Wrong DB Credentials")
 	}else{
 		t.Log("Unable to Connect with bad credentials")
 	}
 }