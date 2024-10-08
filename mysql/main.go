package main

import (
	"database/sql"
	"day11/common"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func update(db *sql.DB) {
	res, err := db.Exec("update student set name='李四',province='北京',city='海淀' where id=2")
	common.CheckErr(err)
	last, _ := res.LastInsertId()
	fmt.Printf("lastInsertId=%d\n", last)
	rows, _ := res.RowsAffected()
	fmt.Printf("Rows Affected =%d\n", rows)
}
func insert(db *sql.DB) {
	res, err := db.Exec("insert into student(name,province,city,enrollment)values('大明','北京','朝阳','2021-06-22'),('小张','上海','浦东','2023-06-21')")
	common.CheckErr(err)
	last, _ := res.LastInsertId()
	fmt.Printf("lastInsertId=%d\n", last)
	rows, _ := res.RowsAffected()
	fmt.Printf("Rows Affected =%d\n", rows)
}
func replace(db *sql.DB) {
	res, err := db.Exec("replace into student(name,province,city,enrollment)values('大明','北京南','朝阳','2021-06-22'),('小张','上海东','浦东','2023-06-21')")
	common.CheckErr(err)
	last, _ := res.LastInsertId()
	fmt.Printf("lastInsertId=%d\n", last)
	rows, _ := res.RowsAffected()
	fmt.Printf("Rows Affected =%d\n", rows)
}
func delete(db *sql.DB) {
	result, err := db.Exec("delete from student where id>15")
	common.CheckErr(err)
	last, _ := result.LastInsertId()
	fmt.Printf("last=%d \n", last)
	rows, _ := result.RowsAffected()
	fmt.Printf("affect rows=%d\n", rows)
}
func query(db *sql.DB) {
	rows, err := db.Query("select id,name,city,score from student where id>2")
	common.CheckErr(err)
	var id int
	var score float32
	var name, city string
	for rows.Next() {
		rows.Scan(&id, &name, &city, &score)
		fmt.Printf("id=%d,name=%s,city=%s,score=%.2f \n", id, name, city, score)
	}
}
func main() {
	db, err := sql.Open("mysql", "teng:5570709@(chent7.top:3306)/test")
	common.CheckErr(err)
	//insert(db)
	update(db)
	//replace(db)
	query(db)
	delete(db)
}
