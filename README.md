# GoMysql #

GoMysql is a Google Go Language Based Database Wraper.Using this Package you can easily intract with mysql database from Go lang.

[![Build Status](https://api.travis-ci.org/biswarupadhikari/gomysql.svg?branch=master)](https://travis-ci.org/biswarupadhikari/gomysql)

## How to Install GoMysql

```
go get -u github.com/biswarupadhikari/gomysql
```

## How to Test GoMysql

```
#Install GoMysql
go get -u github.com/biswarupadhikari/gomysql
#Run Tests
go test github.com/biswarupadhikari/gomysql
```
## Connect To Mysql database

```go
package main
import (
	"github.com/biswarupadhikari/gomysql"
	"log"
)

//Connect Using Default PORT 3306
db,err := gomysql.Connect("localhost", "DBUsername", "DBPassword", "DBName")
if err!=nil{
	log.Fatal("Failed to Connect to Database")
}
/**
 * Connect Using Alternative PORT AND Custom HOST
 * db,err := gomysql.Connect("example.com", "DBUsername", "DBPassword", "DBName", "9595")
 */
```

## Select Record From DataBase Table


```go
users,err:=db.Select("*").From("users").Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select * FROM users
```

## Select Record Using Specific Columns


```go
users,err:=db.Select("id,username").From("users").Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select id,username FROM users
```

## How to Use where Condition


```go
users,err:=db.Select("id,username").From("users").Where("id","=",157).Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select id,username FROM users WHERE id=157
```

## How to Use Multiple Where Condition


```go
users,err:=db.Select("id,username").From("users").Where("username","=","userone").Where("password","=","secret").Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select id,username FROM users WHERE username="userone" AND password="secret"
```
## How to use AND And OR Condition


```go
users,err:=db.Select("id,username").From("users").Where("role","=","administrator").ORWhere("role","=","superadmin").Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select id,username FROM users WHERE role="administrator" OR role="superadmin"
```

## How to Use Raw Where Condition


```go
users,err:=db.Select("id,username").From("users").RawWhere("role=? OR role=?","administrator","superadmin").Get()
if err!=nil{
	log.Fatal(err)
}
for i:=0;i<len(users);i++{
	user:=users[i]
	log.Println("Id => ",user["id"]," || Username ",user["username"])
}
//SQL OUTPUT => Select id,username FROM users WHERE role="administrator" OR role="superadmin"
```


## How to Insert Data to Database Table

```go
data := make(map[string]interface{})
data["username"] = "biswarupadhikari"
data["password"] = "mysecretpass"
data["age"] = 27
db.Table("users").Insert(data)
log.Println("Record Inserted")
//SQL OUTPUT => INSERT INTO users(username,password,age) VALUES("biswarupadhikari","mysecretpass",27)
```

## Alternative Syntax For Insert Data

```go
db.Table("users").InsertSQL(map[string]interface{}{"username":"biswarupadhikari","password":"mysecretpass","age":27})
log.Println("Record Inserted")
//SQL OUTPUT => INSERT INTO users(username,password,age) VALUES("biswarupadhikari","mysecretpass",27)
```

## How to Get Insert ID After Inserting Data

```go
data := make(map[string]interface{})
data["username"] = "biswarupadhikari"
data["password"] = "mysecretpass"
data["age"] = 27
db.Table("users").Insert(data)
result,err := query.Insert(data)
if err!=nil{
	log.Fatal(err)
}
newId, _ := result.LastInsertId()
log.Println("Last Insert Id Is", newId)
//SQL OUTPUT => INSERT INTO users(username,password,age) VALUES("biswarupadhikari","mysecretpass",27)
```

## Update Database Record 

```go
data := make(map[string]interface{})
data["username"] = "mynewusername"
data["password"] = "new secret pass"
db.Table("users").Where("id", "=", 158).Update(data)
log.Println("Record Updated")
//SQL OUTPUT => UPDATE users SET username="mynewusername",password="new secret pass" WHERE id= 158
```

## Update Record Using Multiple Condition

```go
data := make(map[string]interface{})
data["username"] = "mynewusername"
data["password"] = "new secret pass"
db.Table("users").Where("id", "=", 158).Where("role_id","=",15).Update(data)
log.Println("Record Updated")
//SQL OUTPUT => UPDATE users SET username="mynewusername",password="new secret pass" WHERE id= 158  AND role_id=15
```


## Get Affected Rows after Update

```go
data := make(map[string]interface{})
data["password"] = "new Password"
result,err:=db.Table("users").Where("id", ">", 2).Update(data)
if err!=nil{
	log.Fatal(err)
}
affectedRows,_:=result.RowsAffected()
log.Println("Affected Rows",affectedRows)
//SQL OUTPUT => UPDATE users SET password="new Password" WHERE id > 2
```

## Delete Record From Database Table

```go
db.Table("users").Where("id", "=",158).Delete()
log.Println("Record Deleted")
//SQL OUTPUT => DELETE FROM users WHERE id > 158
```

## Get Affected Rows after Delete

```go
result,err:=db.Table("users").Where("id", ">", 2).Delete()
if err!=nil{
	log.Fatal(err)
}
affectedRows,_:=result.RowsAffected()
log.Println("Affected Rows",affectedRows)
//SQL OUTPUT => DELETE FROM users WHERE id > 158
```


## Run Custom Query


```go
result,err := db.Query("<YOUR CUSTOM SQL QUERY>")
if err!=nil{
	log.Fatal(err)
}
//Do Something with result
```

## Run Custom Query Insert Data Process 1

```go
result,err := db.Query("INSERT INTO users(username,password,age) VALUES("biswarupadhikari","mysecretpass",27)")
if err!=nil{
	log.Fatal(err)
}
newId, _ := result.LastInsertId()
log.Println("Last Insert Id Is", newId)
```
## Run Custom Query Insert Data Process 2

```go
result,err := db.Query("INSERT INTO users(username,password,age) VALUES(?,?,?)","biswarupadhikari","mysecretpass",27)
if err!=nil{
	log.Fatal(err)
}
newId, _ := result.LastInsertId()
log.Println("Last Insert Id Is", newId)
```

## Display Table Structure OR Display Rows using QueryRows

```go
rows,err:=db.QueryRows("DESC test_users")
if err!=nil{
	log.Fatal("Table Not Exist")
}
for rows.Next(){
	var Field string
	var Type string
	var Null string
	var Key string
	var Default string
	var Extra string
	err=rows.Scan(&Field,&Type,&Null,&Key,&Default,&Extra)
	log.Println(Field,Type,Null,Key,Default,Extra)
}
```