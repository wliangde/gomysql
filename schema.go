package gomysql

import (
	"database/sql"
	"strconv"
	"strings"
)

/**
 * Table Cloumn Field Structure
 */
type SchemaField struct {
	FieldName         string
	FieldType         string
	FieldSize         string
	FieldIsNull       bool
	FieldIsPrimaryKey bool
	FieldIsUniqueKey  bool
	FieldIsIndexKey   bool
	FieldDefaultValue string
	Extra             string
	schema            *Schema
}

/**
 * Tbale Schema Structure
 */
type Schema struct {
	name               string
	autoIncrement      string
	autoIncrementValue int
	engine             string
	charset            string
	db                 *GoMysql
	fields             []*SchemaField
}

/**
 * Initialize Schema
 */
func (gomysql *GoMysql) Schema(name string) *Schema {
	schema := new(Schema)
	schema.name = name
	schema.engine = "InnoDB"
	schema.charset = "latin1"
	schema.autoIncrementValue = 1
	schema.db = gomysql
	return schema
}

/**
 * Make Field as Primary Key
 */
func (schemaField *SchemaField) Primary() *SchemaField {
	schemaField.FieldIsPrimaryKey = true
	return schemaField
}

/**
 * Make Field as Unique Key
 */
func (schemaField *SchemaField) Unique() *SchemaField {
	schemaField.FieldIsUniqueKey = true
	return schemaField
}

/**
 * Make Field as Index
 */
func (schemaField *SchemaField) Index() *SchemaField {
	schemaField.FieldIsIndexKey = true
	return schemaField
}

/**
 * Set Field Default Value
 */
func (schemaField *SchemaField) Default(defaultValue string) *SchemaField {
	schemaField.FieldDefaultValue = defaultValue
	return schemaField
}

/**
 * Make Field Nullable by default NOT NULL
 */
func (schemaField *SchemaField) Nullable() *SchemaField {
	schemaField.FieldIsNull = true
	return schemaField
}

/**
 * Change Field Length or Size
 */
func (schemaField *SchemaField) Size(sizeValues string) *SchemaField {
	schemaField.FieldSize = sizeValues
	return schemaField
}

/**
 * Create Number Field
 */
func (schema *Schema) Number(name string) *SchemaField {
	schemaField := new(SchemaField)
	schemaField.schema = schema
	schemaField.FieldName = name
	schemaField.FieldType = "Int"
	schemaField.FieldSize = "11"
	schemaField.FieldIsNull = false
	schema.fields = append(schema.fields, schemaField)
	return schemaField
}

/**
 * Create Autoincrement Field
 */
func (schema *Schema) Increment(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "Int"
	schemaField.FieldSize = "11"
	schemaField.FieldIsNull = false
	schemaField.FieldIsPrimaryKey = true
	schema.autoIncrement = schemaField.FieldName
	return schemaField
}

/**
 * Create Integer Field
 */
func (schema *Schema) Int(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "INT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create SmallInt Field
 */
func (schema *Schema) SmallInt(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "SMALLINT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create TinyInt Field
 */
func (schema *Schema) TinyInt(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "TINYINT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create MEDIUMINT Field
 */
func (schema *Schema) MediumInt(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "MEDIUMINT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create BIGINT Field
 */
func (schema *Schema) BigInt(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "BIGINT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create DECIMAL Field
 */
func (schema *Schema) Decimal(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "DECIMAL"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create FLOAT Field
 */
func (schema *Schema) Float(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "FLOAT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create DOUBLE Field
 */
func (schema *Schema) Double(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "DOUBLE"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create REAL Field
 */
func (schema *Schema) Real(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "REAL"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create BIT Field
 */
func (schema *Schema) Bit(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "BIT"
	return schemaField
}

/**
 * Create BOOLEAN Field
 */
func (schema *Schema) Boolean(name string) *SchemaField {
	schemaField := schema.Number(name)
	schemaField.FieldType = "BOOLEAN"
	return schemaField
}

/**
 * Create String Field
 */
func (schema *Schema) String(name string) *SchemaField {
	schemaField := new(SchemaField)
	schemaField.schema = schema
	schemaField.FieldName = name
	schemaField.FieldType = "Varchar"
	schemaField.FieldSize = "255"
	schemaField.FieldIsNull = false
	schema.fields = append(schema.fields, schemaField)
	return schemaField
}

/**
 * Create CHAR Field
 */
func (schema *Schema) Char(name string) *SchemaField {
	schemaField := schema.String(name)
	schemaField.FieldType = "CHAR"
	return schemaField
}

/**
 * Create Varchar Field
 */
func (schema *Schema) Varchar(name string) *SchemaField {
	schemaField := schema.String(name)
	return schemaField
}

/**
 * Create TINYTEXT Field
 */
func (schema *Schema) TinyText(name string) *SchemaField {
	schemaField := schema.String(name)
	schemaField.FieldType = "TINYTEXT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create Text Field
 */
func (schema *Schema) Text(name string) *SchemaField {
	schemaField := schema.String(name)
	schemaField.FieldType = "TEXT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create MEDIUMTEXT Field
 */
func (schema *Schema) MediumText(name string) *SchemaField {
	schemaField := schema.String(name)
	schemaField.FieldType = "MEDIUMTEXT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create LONGTEXT Field
 */
func (schema *Schema) LongText(name string) *SchemaField {
	schemaField := schema.String(name)
	schemaField.FieldType = "LONGTEXT"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create Date Field
 */
func (schema *Schema) Date(name string) *SchemaField {
	schemaField := new(SchemaField)
	schemaField.schema = schema
	schemaField.FieldName = name
	schemaField.FieldType = "DATE"
	schemaField.FieldSize = ""
	schemaField.FieldIsNull = false
	schema.fields = append(schema.fields, schemaField)
	return schemaField
}

/**
 * Create DATETIME Field
 */
func (schema *Schema) DateTime(name string) *SchemaField {
	schemaField := schema.Date(name)
	schemaField.FieldType = "DateTime"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create TIMESTAMP Field
 */
func (schema *Schema) TimeStamp(name string) *SchemaField {
	schemaField := schema.Date(name)
	schemaField.FieldType = "TIMESTAMP"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create TIME Field
 */
func (schema *Schema) Time(name string) *SchemaField {
	schemaField := schema.Date(name)
	schemaField.FieldType = "TIME"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create YEAR Field
 */
func (schema *Schema) Year(name string) *SchemaField {
	schemaField := schema.Date(name)
	schemaField.FieldType = "YEAR"
	schemaField.FieldSize = ""
	return schemaField
}

/**
 * Create Enum Field
 */
func (schema *Schema) Enum(name string) *SchemaField {
	schemaField := new(SchemaField)
	schemaField.schema = schema
	schemaField.FieldName = name
	schemaField.FieldType = "ENUM"
	schemaField.FieldSize = ""
	schemaField.FieldIsNull = false
	schema.fields = append(schema.fields, schemaField)
	return schemaField
}

/**
 * Generate Schema SQL
 */
func (schema *Schema) CreateSQL() string {
	var sqlQuery string
	sqlQuery += "\nCREATE TABLE IF NOT EXISTS " + schema.name + "(\n"
	var primaryKeys []string
	var UniqueKeys []string
	var IndexKeys []string
	for _, field := range schema.fields {
		var fieldSQL string
		if field.FieldIsPrimaryKey {
			primaryKeys = append(primaryKeys, field.FieldName)
		}
		if field.FieldIsUniqueKey {
			UniqueKeys = append(UniqueKeys, field.FieldName)
		}
		if field.FieldIsIndexKey {
			IndexKeys = append(IndexKeys, field.FieldName)
		}
		fieldSQL += field.FieldName + " " + field.FieldType
		if field.FieldSize != "" {
			fieldSQL += "(" + field.FieldSize + ")"
		}
		if field.FieldIsNull {
			fieldSQL += " NULL"
		} else {
			fieldSQL += " NOT NULL"
		}
		if field.FieldDefaultValue != "" {
			fieldSQL += " DEFAULT " + field.FieldDefaultValue
		}
		if schema.autoIncrement == field.FieldName {
			fieldSQL += "  AUTO_INCREMENT"
		}
		sqlQuery += "   " + fieldSQL + ",\n"
	}
	if len(primaryKeys) > 0 {
		sqlQuery += "   PRIMARY KEY (" + strings.Join(primaryKeys, ",") + "),\n"
	}
	for _, uniqueKey := range UniqueKeys {
		sqlQuery += "   UNIQUE KEY " + uniqueKey + " (" + uniqueKey + "),\n"
	}
	for _, IndexKey := range IndexKeys {
		sqlQuery += "   KEY " + IndexKey + " (" + IndexKey + "),\n"
	}
	sqlQuery = strings.TrimRight(sqlQuery, ",\n") + "\n"
	sqlQuery += ")"
	sqlQuery += "  ENGINE=" + schema.engine + " DEFAULT CHARSET=" + schema.charset
	if schema.autoIncrement == "" {
		sqlQuery += " ;"
	} else {
		sqlQuery += " AUTO_INCREMENT=" + strconv.Itoa(schema.autoIncrementValue) + " ;"
	}
	return sqlQuery
}

/**
 * Create Schema into Database
 */
func (schema *Schema) Create() (sql.Result, error) {
	schemaSql := schema.CreateSQL()
	return schema.db.Query(schemaSql)
}

/**
 * Drop Schema
 */
func (schema *Schema) Drop() (sql.Result, error) {
	return schema.db.Query("DROP TABLE IF EXISTS " + schema.name)
}

/**
 * Rename Schema
 */
func (schema *Schema) Rename(newTableName string) (sql.Result, error) {
	return schema.db.Query("RENAME TABLE " + schema.name + " TO " + newTableName)
}

/**
 * Generate AddColumn SQL
 */
func (schemaField *SchemaField) AddColumnSQL(params ...string) string {
	sqlQuery := "ALTER TABLE " + schemaField.schema.name + " ADD "
	sqlQuery += schemaField.FieldName + " " + schemaField.FieldType + "(" + schemaField.FieldSize + ")"
	if schemaField.FieldIsNull {
		sqlQuery += " NULL"
	} else {
		sqlQuery += " NOT NULL"
	}
	if schemaField.FieldDefaultValue != "" {
		sqlQuery += " DEFAULT " + schemaField.FieldDefaultValue + ""
	}
	if len(params) == 1 {
		sqlQuery += " FIRST "
	}
	if len(params) == 2 {
		sqlQuery += " AFTER " + params[1]
	}
	sqlQuery += ",\n"
	if schemaField.FieldIsPrimaryKey {
		sqlQuery += "ADD PRIMARY KEY(" + schemaField.FieldName + "),\n"
	}
	if schemaField.FieldIsUniqueKey {
		sqlQuery += "ADD UNIQUE KEY(" + schemaField.FieldName + "),\n"
	}
	if schemaField.FieldIsIndexKey {
		sqlQuery += "ADD INDEX(" + schemaField.FieldName + "),\n"
	}
	sqlQuery = strings.TrimRight(sqlQuery, ",\n") + ";\n"
	return sqlQuery
}

/**
 * Add Column To Table
 */
func (schemaField *SchemaField) AddColumn(params ...string) (sql.Result, error) {
	return schemaField.schema.db.Query(schemaField.AddColumnSQL(params...))
}

/**
 * Add Column To Table After Some COlumn
 */
func (schemaField *SchemaField) AddColumnFirst() (sql.Result, error) {
	return schemaField.schema.db.Query(schemaField.AddColumnSQL("FIRST"))
}

/**
 * Add Column To Table After Some COlumn
 */
func (schemaField *SchemaField) AddColumnAfter(afterColumnName string) (sql.Result, error) {
	return schemaField.schema.db.Query(schemaField.AddColumnSQL("AFTER", afterColumnName))
}
