package gomysql

import (
	"log"
	"os"
	"testing"
)

var (
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
	DBPort     string
)

var (
	db *GoMysql
)

/**
 * Get Test Config If env not set take default one
 */
func getConfig(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

/**
 * Setup Tests
 */
func init() {
	log.Println("Starting GoMysql Testing")
	DBHost = getConfig("GOMYSQL_TEST_HOST", "localhost")
	DBUsername = getConfig("GOMYSQL_TEST_USERNAME", "root")
	DBPassword = getConfig("GOMYSQL_TEST_PASSWORD", "")
	DBName = getConfig("GOMYSQL_TEST_DBNAME", "gomysql_test")
	DBPort = getConfig("GOMYSQL_TEST_DBPORT", "3306")
}

/**
 * Test Mysql Connect With
 */
func TestConnectionWithGoodCredentials(t *testing.T) {
	gomysql, err := Connect(DBHost, DBUsername, DBPassword, DBName, DBPort)
	if err != nil {
		t.Error("Failed to Connect Database Using Good Credentials")
	} else {
		db = gomysql
		t.Log("Success Mysql Connect Passed Using Right Credentials")
	}
}

/**
 * Test Mysql Connect With Wrong Credentials
 */
func TestConnectionWithBadCredentials(t *testing.T) {
	_, err := Connect(DBHost+"wrong", DBUsername+"wrong", DBPassword+"wrong", DBName+"wrong", DBPort+"wrong")
	if err == nil {
		t.Error("Not Thow Error with Wrong DB Credentials")
	} else {
		t.Log("Unable to Connect with bad credentials")
	}
}

/**
 * Test Set Table
 */
func TestTable(t *testing.T) {
	tbl := "users_table"
	db.Table(tbl)
	if db.tableName == tbl {
		t.Log("Table Set Working")
	} else {
		t.Error("Unable to SetTable Name")
	}
}

/**
 * Test From Table Name Set
 */
func TestFrom(t *testing.T) {
	tbl := "users_table"
	db.From(tbl)
	if db.tableName == tbl {
		t.Log("Table Set Working using From ")
	} else {
		t.Error("Unable to SetTable Name")
	}
}

/**
 * Select Test
 */
func TestSelect(t *testing.T) {
	db.Select("id")
	if len(db.fields) == 1 && db.fields[0] == "id" {
		t.Log("Single column select working")
	} else {
		t.Error("Problem in Select Field Set Single Select")

	}

	db.Select("username")
	if len(db.fields) == 2 && db.fields[0] == "id" && db.fields[1] == "username" {
		t.Log("Single column select working")
	} else {

		t.Error("Problem in Select Field Set Single Select")
	}

	db.Select("age,role,sex")
	if len(db.fields) == 5 && db.fields[0] == "id" && db.fields[1] == "username" && db.fields[2] == "age" && db.fields[3] == "role" && db.fields[4] == "sex" {
		t.Log("Single column select working")
	} else {

		t.Error("Problem in Select Field Set Single Select")
	}

}

/**
 * Test Clear Select
 */
func TestClearSelect(t *testing.T) {
	db.Select("id,username,password")
	totalBeforeClear := len(db.fields)
	db.ClearSelect()
	totalAfterClear := len(db.fields)
	if totalBeforeClear > 0 && totalAfterClear == 0 {
		t.Log("Clear Select Working")
	} else {
		t.Error("Clear Select Not Working")
	}
}

/**
 * Test Chain Select
 */
func TestChainSelect(t *testing.T) {
	db.ClearSelect()
	q := db.Select("id")
	q.Select("username,password")
	if len(db.fields) == 3 && db.fields[0] == "id" && db.fields[1] == "username" && db.fields[2] == "password" {
		t.Log("Chain Select Working")
	} else {
		t.Error("Problem in Chain Select")
	}
}

/**
 * Test Where Value Set
 */

func TestWhere(t *testing.T) {
	db.Where("username", "=", "biswarupadhikari")
	if len(db.conditions) == 1 && db.conditions[0] == "username= ?" && db.conditionValues[0] == "biswarupadhikari" {
		t.Log("Single Where Working")
	} else {
		t.Error("Problem in Single Where")
	}
}

/**
 * Test Clear Where Condition
 */
func TestClearWhere(t *testing.T) {
	db.Where("username", "=", "biswarupadhikari")
	db.Where("password", "=", "secret")

	totalCondBeforeClear := len(db.conditions)
	totalCondVBeforeClear := len(db.conditionValues)

	db.ClearWhere()

	totalCondAfterClear := len(db.conditions)
	totalCondVAfterClear := len(db.conditionValues)

	if totalCondBeforeClear > 0 && totalCondVBeforeClear > 0 && totalCondAfterClear == 0 && totalCondVAfterClear == 0 {
		t.Log("Clear Where Working")
	} else {
		t.Error("Clear Where Not Working")
	}
}

/**
 * Test Query Using Create Table
 */
func TestQueryCreateTable(t *testing.T) {
	sqlQuery := `
    CREATE TABLE IF NOT EXISTS test_users (
    id int(11) NOT NULL AUTO_INCREMENT,
    username varchar(250) NOT NULL,
    password varchar(250) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY username (username)
  ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;
  `
	_, err := db.Query(sqlQuery)
	if err != nil {
		t.Error("Table Create Failed")
	} else {
		t.Log("test users Table created or exists")
	}
	//Check Table is Exist or not
	_, err = db.Query("DESC test_users")
	if err != nil {
		t.Error("Table not Exist After Create")
	} else {
		t.Log("Table Exist")
	}
}

/**
 * Test Query Using Drop Table
 */
func TestQueryDropTable(t *testing.T) {
	sqlQuery := `
    CREATE TABLE IF NOT EXISTS test_users (
    id int(11) NOT NULL AUTO_INCREMENT,
    username varchar(250) NOT NULL,
    password varchar(250) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY username (username)
  ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;
  `
	db.Query(sqlQuery)
	//Drop Table
	_, err := db.Query("DROP TABLE test_users")
	if err != nil {
		t.Error("Error droping Table")
	} else {
		t.Log("Table Dropped")
	}
	//Confirm Table Dropped or Not
	_, err = db.Query("DESC test_users")
	if err != nil {
		t.Log("Confirmed Table Droped")
	} else {
		t.Error("Table Not Dropped Properly")
	}
}

/**
 * Test Query Using Insert Table
 */
func TestQueryInsertIntoTable(t *testing.T) {
	sqlQuery := `
    CREATE TABLE IF NOT EXISTS test_users (
    id int(11) NOT NULL AUTO_INCREMENT,
    username varchar(250) NOT NULL,
    password varchar(250) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY username (username)
  ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;
  `
	db.Query(sqlQuery)
	//Run Insert Query
	result, err := db.Query("INSERT INTO test_users(id,username,password) VALUES(14599,'biswarupadhikari','secret')")
	if err != nil {
		t.Error("Insert Data into test_users table failed")
	} else {
		t.Log("Data Inserted into test_users")
	}
	newId, _ := result.LastInsertId()
	if newId != 14599 {
		t.Error("Insert Data Error insert id not matching")
	} else {
		t.Log("Insert id Matching working Insert")
	}

}
