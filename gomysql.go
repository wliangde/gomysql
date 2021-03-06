package gomysql

import (
"database/sql"
"errors"
_ "github.com/go-sql-driver/mysql"
_ "log"
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
 	joins           []string
 	dataValues      []interface{}
 	conditionValues []interface{}
 	joinDataValues  []interface{}
 	lastResult      sql.Result
 	lastRow         *sql.Row
 	lastRows        *sql.Rows
 }
/**
 * Connect To DataBase using username pass and dbname Host Name and Port
 * host localhost
 * port 3306
 */
 func Connect(dbHost string, dbUsername string, dbPassword string, dbName string, params ...string) (*GoMysql, error) {
 	var dbPort string
 	gomysql := new(GoMysql)
 	if len(params) > 0 {
 		dbPort = params[0]
 	} else {
 		dbPort = "3306"
 	}
 	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
 	if err != nil {
 		return gomysql, err
 	}
 	err = db.Ping()
 	if err != nil {
 		return gomysql, err
 	}
 	gomysql.db = db
 	return gomysql, nil
 }
/**
 * Ping Database
 */
 func (gomysql *GoMysql) Ping() bool {
 	err := gomysql.db.Ping()
 	if err != nil {
 		return false
 	}
 	return true;
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
 * Clear Select
 */
 func (gomysql *GoMysql) ClearSelect() *GoMysql {
 	gomysql.fields = gomysql.fields[:0]
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
 	gomysql.applyCondition(""+key+""+operator+" ?", " AND ")
 	gomysql.conditionValues = append(gomysql.conditionValues, dataValue)
 	return gomysql
 }

/**
 * Clear Where
 */
 func (gomysql *GoMysql) ClearWhere() *GoMysql {
 	gomysql.conditions = gomysql.conditions[:0]
 	gomysql.conditionValues = gomysql.conditionValues[:0]
 	return gomysql
 }

/**
 * Add OR Where Conditions
 */
 func (gomysql *GoMysql) ORWhere(key string, operator string, dataValue interface{}) *GoMysql {
 	gomysql.applyCondition(""+key+""+operator+" ?", " OR ")
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
 * Join Query
 */
 func (gomysql *GoMysql) Join(joinType string, joinTable string, joinCond string, joinDataValues ...interface{}) *GoMysql {
 	joinSQL := " " + joinType + " JOIN " + joinTable + " ON " + joinCond
 	for _, joinDataValue := range joinDataValues {
 		gomysql.joinDataValues = append(gomysql.joinDataValues, joinDataValue)
 	}
 	gomysql.joins = append(gomysql.joins, joinSQL)
 	return gomysql
 }

/**
 * Generate Select SQL using Select Fields From Table and Condition
 */
 func (gomysql *GoMysql) generateSelectSQL() string {
 	var sqlQuery string
 	sqlQuery += "Select " + strings.Join(gomysql.fields, ",") + " FROM " + gomysql.tableName
 	if len(gomysql.joins) > 0 {
 		sqlQuery += " " + strings.Join(gomysql.joins, " ")
 	}
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

 	for _, value := range gomysql.joinDataValues {
 		values = append(values, value)
 	}

 	for _, value := range gomysql.conditionValues {
 		values = append(values, value)
 	}

 	return values
 }

/**
 * Reset Query Data
 */
 func (gomysql *GoMysql) RessetQuery() {
 	gomysql.tableName = ""
 	gomysql.fields = make([]string, 0)
 	gomysql.conditions = make([]string, 0)
 	gomysql.joins = make([]string, 0)
 	gomysql.dataValues = make([]interface{}, 0)
 	gomysql.conditionValues = make([]interface{}, 0)
 	gomysql.joinDataValues = make([]interface{}, 0)
 }

/**
 * Get New Query Start New Query
 */
 func (gomysql *GoMysql) GetQuery() *GoMysql {
 	gomysql.RessetQuery()
 	return gomysql
 }

/**
 * Get Records
 */
 func (gomysql *GoMysql) Get() ([]map[string]interface{}, error) {
 	var rows *sql.Rows
 	var err error
 	var columns []string
 	items := make([]map[string]interface{}, 0)

 	if gomysql.tableName == "" {
 		return items, errors.New("Please Select Table")
 	}

 	rows, err = gomysql.db.Query(gomysql.generateSelectSQL(), gomysql.GetMappedValues()...)
 	if err != nil {
 		return items, err
 	}
 	columns, err = rows.Columns()
 	if err != nil {
 		return items, err
 	}
 	values := make([]sql.RawBytes, len(columns))
 	scanArgs := make([]interface{}, len(values))
 	for i := range values {
 		scanArgs[i] = &values[i]
 	}
 	for rows.Next() {
 		item := make(map[string]interface{})
 		err = rows.Scan(scanArgs...)
 		if err != nil {
 			return items, err
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
 	gomysql.RessetQuery()
 	return items, nil
 }

/**
 * Get Sql Query
 */
 func (gomysql *GoMysql) GetSQL() (string, []interface{}) {
 	return gomysql.generateSelectSQL(), gomysql.GetMappedValues()
 }

/**
 * Insert Data Into Table Using Data
 */
 func (gomysql *GoMysql) Insert(data map[string]interface{}) (sql.Result, error) {
 	var result sql.Result
 	var err error

 	var fieldName string
 	var fieldValue interface{}

 	if gomysql.tableName == "" {
 		return result, errors.New("Please Select Table")
 	}

 	for fieldName, fieldValue = range data {
 		gomysql.fields = append(gomysql.fields, fieldName)
 		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
 	}

 	result, err = gomysql.Query(gomysql.generateInsertSQL(), gomysql.GetMappedValues()...)
 	if err != nil {
 		return result, err
 	}
 	gomysql.lastResult = result
 	gomysql.RessetQuery()
 	return result, nil
 }

/**
 * Get Insert Sql Query
 */
 func (gomysql *GoMysql) InsertSQL(data map[string]interface{}) (string, []interface{}) {
 	var fieldName string
 	var fieldValue interface{}
 	for fieldName, fieldValue = range data {
 		gomysql.fields = append(gomysql.fields, fieldName)
 		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
 	}
 	return gomysql.generateInsertSQL(), gomysql.GetMappedValues()
 }

/**
 * Update Data
 */
 func (gomysql *GoMysql) Update(data map[string]interface{}) (sql.Result, error) {
 	var result sql.Result
 	var err error

 	var fieldName string
 	var fieldValue interface{}

 	if gomysql.tableName == "" {
 		return result, errors.New("Please Select Table")
 	}

 	for fieldName, fieldValue = range data {
 		gomysql.fields = append(gomysql.fields, fieldName)
 		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
 	}
 	result, err = gomysql.Query(gomysql.generateUpdateSQL(), gomysql.GetMappedValues()...)
 	if err != nil {
 		return result, err
 	}
 	gomysql.lastResult = result
 	gomysql.RessetQuery()
 	return result, nil
 }

/**
 * Get Insert Sql Query
 */
 func (gomysql *GoMysql) UpdateSQL(data map[string]interface{}) (string, []interface{}) {
 	var fieldName string
 	var fieldValue interface{}
 	for fieldName, fieldValue = range data {
 		gomysql.fields = append(gomysql.fields, fieldName)
 		gomysql.dataValues = append(gomysql.dataValues, fieldValue)
 	}
 	return gomysql.generateUpdateSQL(), gomysql.GetMappedValues()
 }

/**
 * Delete Records From Mysql DataBase Table
 */
 func (gomysql *GoMysql) Delete() (sql.Result, error) {
 	var result sql.Result
 	var err error

 	if gomysql.tableName == "" {
 		return result, errors.New("Please Select Table")
 	}
 	result, err = gomysql.Query(gomysql.generateDeleteSQL(), gomysql.GetMappedValues()...)
 	if err != nil {
 		return result, err
 	}
 	gomysql.lastResult = result
 	gomysql.RessetQuery()
 	return result, nil
 }

/**
 * Get Delete Sql Query
 */
 func (gomysql *GoMysql) DeleteSQL() (string, []interface{}) {
 	return gomysql.generateDeleteSQL(), gomysql.GetMappedValues()
 }

/**
 * Run Custom Query
 */
 func (gomysql *GoMysql) Query(sqlQuery string, args ...interface{}) (sql.Result, error) {
 	var result sql.Result
 	var err error
 	result, err = gomysql.db.Exec(sqlQuery, args...)
 	if err != nil {
 		return result, err
 	}
 	gomysql.lastResult = result
 	return result, nil
 }

/**
 * Query and Get Rows
 */
 func (gomysql *GoMysql) QueryRows(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
 	var rows *sql.Rows
 	var err error
 	rows, err = gomysql.db.Query(sqlQuery, args...)
 	if err != nil {
 		return rows, err
 	}
 	gomysql.lastRows = rows
 	return rows, nil
 }

/**
 * Query and Get Row
 */
 func (gomysql *GoMysql) QueryRow(sqlQuery string, args ...interface{}) *sql.Row {
 	var row *sql.Row
 	row = gomysql.db.QueryRow(sqlQuery, args...)
 	gomysql.lastRow = row
 	return row
 }
