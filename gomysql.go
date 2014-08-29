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
   db         *sql.DB
   fields     []string
   tableName  string
   conditions []string
   dataValues []interface{}
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
   gomysql.dataValues = append(gomysql.dataValues, dataValue)
   return gomysql
 }

/**
 * Add OR Where Conditions
 */
 func (gomysql *GoMysql) ORWhere(key string, operator string, dataValue interface{}) *GoMysql {
   gomysql.applyCondition("`"+key+"`"+operator+" ?", " OR ")
   gomysql.dataValues = append(gomysql.dataValues, dataValue)
   return gomysql
 }

/**
 * Add OR Where Conditions
 */
 func (gomysql *GoMysql) RawWhere(condition string, dataValues ...interface{}) *GoMysql {
   gomysql.conditions = append(gomysql.conditions, condition)
   for _, dataValue := range dataValues {
    gomysql.dataValues = append(gomysql.dataValues, dataValue)
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
  var sql string
  sql += "Select " + strings.Join(gomysql.fields, ",") + " FROM " + gomysql.tableName
  if len(gomysql.conditions) > 0 {
    sql += " WHERE " + strings.Join(gomysql.conditions, " ")
  }
  return sql
}


/**
 * Generate Insert SQL using Table Name
 */
 func (gomysql *GoMysql) generateInsertSQL() string {
  var sql string
  var placeholders string
  sql+="INSERT INTO "+gomysql.tableName+"("+strings.Join(gomysql.fields, ",")+")"

  for i:=0;i<len(gomysql.fields);i++{
    placeholders+="?,"
  }
  placeholders=strings.TrimRight(placeholders,",")
  sql+=" VALUES("+placeholders+")"
  return sql
}

/**
 * Get Records
 */
 func (gomysql *GoMysql) Get() {
   sql := gomysql.generateSelectSQL()
   log.Println(sql, gomysql.dataValues)
 }

/**
 * Insert Data Into Table Using Data
 */
 func (gomysql *GoMysql) Insert(data map[string]interface{}) {
   var fieldName string
   var fieldValue interface{}
   for fieldName,fieldValue=range data{
    gomysql.fields = append(gomysql.fields, fieldName)
    gomysql.dataValues = append(gomysql.dataValues, fieldValue)
  }
  sql:=gomysql.generateInsertSQL()
  log.Println(sql, gomysql.dataValues)
}