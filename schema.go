package gomysql
import(
"database/sql"
"strconv"
"strings"
)
/**
 * Table Cloumn Field Structure 
 */
type SchemaField struct{
	FieldName string
	FieldType string
	FieldSize string
	FieldIsNull bool
	FieldIsPrimaryKey bool
	FieldIsUniqueKey bool
	FieldIsIndexKey bool
	FieldDefaultValue string
	Extra string
}
/**
 * Tbale Schema Structure
 */
type Schema struct{
	name string
	autoIncrement string
	autoIncrementValue int
	engine string
	charset string
	db *GoMysql
	fields[]*SchemaField

}
/**
 * Initialize Schema
 */
func (gomysql *GoMysql) Schema(name string) *Schema {
	schema:=new(Schema)
	schema.name=name
	schema.engine="InnoDB"
	schema.charset="latin1"
	schema.autoIncrementValue=1
	schema.db=gomysql
	return schema
}
/**
 * Make Field as Primary Key
 */
func(schemaField *SchemaField) Primary()(*SchemaField){
	schemaField.FieldIsPrimaryKey=true
	return schemaField
}
/**
 * Make Field as Unique Key
 */
func(schemaField *SchemaField) Unique()(*SchemaField){
	schemaField.FieldIsUniqueKey=true
	return schemaField
}
/**
 * Make Field as Index
 */
func(schemaField *SchemaField) Index()(*SchemaField){
	schemaField.FieldIsIndexKey=true
	return schemaField
}
/**
 * Set Field Default Value
 */
func(schemaField *SchemaField) Default(defaultValue string)(*SchemaField){
	schemaField.FieldDefaultValue=defaultValue
	return schemaField
}
/**
 * Make Field Nullable by default NOT NULL
 */
func(schemaField *SchemaField) Nullable()(*SchemaField){
	schemaField.FieldIsNull=true
	return schemaField
}
/**
 * Change Field Length or Size
 */
func(schemaField *SchemaField) Size(sizeValues string)(*SchemaField){
	schemaField.FieldSize=sizeValues
	return schemaField
}
/**
 * Create Autoincrement Field
 */
func(schema *Schema) Increment(name string)(*SchemaField){
	schemaField:=new(SchemaField)
	schemaField.FieldName=name
	schemaField.FieldType="Int"
	schemaField.FieldSize="11"
	schemaField.FieldIsNull=false
	schemaField.FieldIsPrimaryKey=true
	schema.autoIncrement=schemaField.FieldName
	schema.fields=append(schema.fields,schemaField)
	return schemaField
}
/**
 * Create Integer Field
 */
func(schema *Schema) Int(name string)(*SchemaField){
	schemaField:=new(SchemaField)
	schemaField.FieldName=name
	schemaField.FieldType="Int"
	schemaField.FieldSize="11"
	schemaField.FieldIsNull=false
	schema.fields=append(schema.fields,schemaField)
	return schemaField
}
/**
 * Create Varchar Field
 */
func(schema *Schema) Varchar(name string)(*SchemaField){
	schemaField:=new(SchemaField)
	schemaField.FieldName=name
	schemaField.FieldType="Varchar"
	schemaField.FieldSize="255"
	schemaField.FieldIsNull=false
	schema.fields=append(schema.fields,schemaField)
	return schemaField
}
/**
 * Create String Field alias of Varchar
 */
func(schema *Schema) String(name string)(*SchemaField){
	return schema.Varchar(name)
}
/**
 * Create Enum Field
 */
func(schema *Schema) Enum(name string)(*SchemaField){
	schemaField:=new(SchemaField)
	schemaField.FieldName=name
	schemaField.FieldType="enum"
	schemaField.FieldSize=""
	schemaField.FieldIsNull=false
	schema.fields=append(schema.fields,schemaField)
	return schemaField
}
/**
 * Generate Schema SQL
 */
func(schema *Schema) CreateSQL()string{
	var sqlQuery string
	sqlQuery+="\nCREATE TABLE IF NOT EXISTS "+schema.name+"(\n"
	var primaryKeys[]string
	var UniqueKeys[]string
	var IndexKeys[]string
	for _,field:=range(schema.fields){
		var fieldSQL string
		if field.FieldIsPrimaryKey{
			primaryKeys=append(primaryKeys,field.FieldName)
		}
		if field.FieldIsUniqueKey{
			UniqueKeys=append(UniqueKeys,field.FieldName)
		}
		if field.FieldIsIndexKey{
			IndexKeys=append(IndexKeys,field.FieldName)
		}
		fieldSQL+=field.FieldName+" "+field.FieldType+"("+field.FieldSize+")"
		if(field.FieldIsNull){
			fieldSQL+=" NULL"
		}else{
			fieldSQL+=" NOT NULL"
		}
		if field.FieldDefaultValue!=""{
			fieldSQL+=" DEFAULT '"+field.FieldDefaultValue+"'"
		}
		if schema.autoIncrement==field.FieldName{
			fieldSQL+="  AUTO_INCREMENT"
		}
		sqlQuery+="   "+fieldSQL+",\n"
	}
	if len(primaryKeys)>0{
		sqlQuery+="   PRIMARY KEY ("+strings.Join(primaryKeys, ",")+"),\n"
	}
	for _,uniqueKey:=range(UniqueKeys){
		sqlQuery+="   UNIQUE KEY "+uniqueKey+" ("+uniqueKey+"),\n"
	}
	for _,IndexKey:=range(IndexKeys){
		sqlQuery+="   KEY "+IndexKey+" ("+IndexKey+"),\n"
	}
	sqlQuery=strings.TrimRight(sqlQuery,",\n")+"\n"
	sqlQuery+=")"
	sqlQuery+="  ENGINE="+schema.engine+" DEFAULT CHARSET="+schema.charset
	if schema.autoIncrement==""{
		sqlQuery+=" ;"
	}else{
		sqlQuery+=" AUTO_INCREMENT="+strconv.Itoa(schema.autoIncrementValue)+" ;"
	}
	return sqlQuery
}
/**
 * Create Schema into Database
 */
func(schema *Schema) Create()(sql.Result, error){
	schemaSql:=schema.CreateSQL()
	return schema.db.Query(schemaSql)
}
/**
 * Drop Schema
 */
func(schema *Schema) Drop()(sql.Result, error){
	return schema.db.Query("DROP TABLE IF EXISTS "+schema.name)
}
/**
 * Rename Schema 
 */
func(schema *Schema) Rename(newTableName string)(sql.Result, error){
	return schema.db.Query("RENAME TABLE "+schema.name+" TO "+newTableName)
}