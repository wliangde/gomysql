package gomysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

/**
 * GoMysql Structure
 */
type GoMysql struct {
	db              *sql.DB
	tableName       string
	fields          []string
	conditions      []string
	dataValues      []interface{}
	conditionValues []interface{}
}

/**
 * Connect To DataBase
 */
func Connect(dbUsername string, dbPassword string, dbName string) *GoMysql {
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	gomysql := new(GoMysql)
	gomysql.db = db
	return gomysql
}

/**
 * Select Table
 */
func (gomysql *GoMysql) Table(tableName string) *GoMysql {
	gomysql.tableName = tableName
	return gomysql
}

/**
 * Select Fields
 */
func (gomysql *GoMysql) Select(fields string) *GoMysql {
	cols := strings.Split(fields, ",")
	for _, field := range cols {
		gomysql.fields = append(gomysql.fields, field)
	}
	return gomysql
}

/**
 * Select From a Table
 */
func (gomysql *GoMysql) From(tableName string) *GoMysql {
	gomysql.tableName = tableName
	return gomysql
}

/**
 * Add AND Where Conditions
 */
func (gomysql *GoMysql) Where(key string, operator string, dataValue interface{}) *GoMysql {
	gomysql.applyCondition("`"+key+"`"+operator+" ?", " AND ")
	gomysql.conditionValues = append(gomysql.conditionValues, dataValue)
	return gomysql
}

/**
 * Add OR Where Conditions
 */
func (gomysql *GoMysql) ORWhere(key string, operator string, dataValue interface{}) *GoMysql {
	gomysql.applyCondition("`"+key+"`"+operator+" ?", " OR ")
	gomysql.conditionValues = append(gomysql.conditionValues, dataValue)
	return gomysql
}

/**
 * Add OR Where Conditions
 */
func (gomysql *GoMysql) RawWhere(condition string, dataValues ...interface{}) *GoMysql {
	gomysql.conditions = append(gomysql.conditions, condition)
	for _, dataValue := range dataValues {
		gomysql.conditionValues = append(gomysql.conditionValues, dataValue)
	}

	return gomysql
}

/**
 * Apply Condition And Add Logical Operator(AND,OR)
 */
func (gomysql *GoMysql) applyCondition(condition string, operatorCondition string) {
	if len(gomysql.conditions) > 0 {
		gomysql.conditions = append(gomysql.conditions, operatorCondition+condition)
	} else {
		gomysql.conditions = append(gomysql.conditions, condition)
	}
}

/**
 * Generate Select SQL using Select Fields From Table and Condition
 */
func (gomysql *GoMysql) generateSelectSQL() string {
	var sqlQuery string
	sqlQuery += "Select " + strings.Join(gomysql.fields, ",") + " FROM " + gomysql.tableName
	if len(gomysql.conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(gomysql.conditions, " ")
	}
	return sqlQuery
}

/**
 * Generate Insert SQL Query
 */
func (gomysql *GoMysql) generateInsertSQL() string {
	var sqlQuery string
	var placeholders string
	sqlQuery += "INSERT INTO " + gomysql.tableName + "(" + strings.Join(gomysql.fields, ",") + ")"

	for i := 0; i < len(gomysql.fields); i++ {
		placeholders += "?,"
	}
	placeholders = strings.TrimRight(placeholders, ",")
	sqlQuery += " VALUES(" + placeholders + ")"
	return sqlQuery
}

/**
 * Generate Update SQL Query
 */
func (gomysql *GoMysql) generateUpdateSQL() string {
	var sqlQuery string
	var fieldMaping string
	sqlQuery += "UPDATE " + gomysql.tableName + " SET "
	for _, field := range gomysql.fields {
		fieldMaping += field + "=?,"
	}
	fieldMaping = strings.TrimRight(fieldMaping, ",")
	sqlQuery += fieldMaping
	if len(gomysql.conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(gomysql.conditions, " ")
	}
	return sqlQuery
}

/**
 * Generate Delete SQL Query
 */
func (gomysql *GoMysql) generateDeleteSQL() string {
	var sqlQuery string
	sqlQuery += "DELETE FROM " + gomysql.tableName
	if len(gomysql.conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(gomysql.conditions, " ")
	}

	return sqlQuery
}

/**
 *Get Combined DataValues and Condtion Mapped Values
 */
func (gomysql *GoMysql) GetMappedValues() []interface{} {
	values := make([]interface{}, 0)
	for _, value := range gomysql.dataValues {
		values = append(values, value)
	}
	for _, value := range gomysql.conditionValues {
		values = append(values, value)
	}

	return values
}

/**
 * Get Records
 */
func (gomysql *GoMysql) Get() []map[string]interface{} {
	item := make(map[string]interface{})
	items := make([]map[string]interface{}, 0)
	sqlQuery := gomysql.generateSelectSQL()
	log.Println(sqlQuery)
	//return
	rows, err := gomysql.db.Query(sqlQuery, gomysql.GetMappedValues()...)
	if err != nil {
		log.Fatal(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			item[columns[i]] = value
		}
		items = append(items, item)
	}
	return items
}

/**
 * Insert Data Into Table Using Data
 */
func (gomysql *GoMysql) Insert(data map[string]interface{}) {
	var fieldName string
	var fieldValue interface{}
	for fieldName, fieldValue = range data {
		gomysql.fields = append(gomysql.fields, fieldName)
		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
	}
	sqlQuery := gomysql.generateInsertSQL()
	stmtIns, err := gomysql.db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(gomysql.GetMappedValues()...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sqlQuery, gomysql.dataValues)
}

/**
 * Update Data
 */
func (gomysql *GoMysql) Update(data map[string]interface{}) {
	var fieldName string
	var fieldValue interface{}
	for fieldName, fieldValue = range data {
		gomysql.fields = append(gomysql.fields, fieldName)
		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
	}
	sqlQuery := gomysql.generateUpdateSQL()
	stmtIns, err := gomysql.db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(gomysql.GetMappedValues()...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sqlQuery, gomysql.GetMappedValues())
}

/**
 * Delete Records From Mysql DataBase Table
 */
func (gomysql *GoMysql) Delete() {
	sqlQuery := gomysql.generateDeleteSQL()
	stmtIns, err := gomysql.db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(gomysql.GetMappedValues()...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sqlQuery, gomysql.GetMappedValues())
}
