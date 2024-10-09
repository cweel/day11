package main

import (
	"database/sql"
	"day11/common"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into student(name,province,city,enrollment)values(?,?,?,?),(?,?,?,?)")
	common.CheckErr(err)
	result, _ := stmt.Exec("小狗", "美国", "旧金山", "2010-03-23", "小猫", "日本", "东京", "2011-05-11")
	last, _ := result.LastInsertId()
	fmt.Printf("last insert=%d \n", last)
	affect, _ := result.RowsAffected()
	fmt.Printf("affect rows=%d \n", affect)
}
func replace(db *sql.DB) {
	stmt, err := db.Prepare("replace into student(name,province,city,enrollment)values(?,?,?,?),(?,?,?,?)")
	common.CheckErr(err)
	result, _ := stmt.Exec("小狗", "美国", "旧金山", "2010-03-23", "小猫", "日本", "大阪", "2011-05-11")
	last, _ := result.LastInsertId()
	fmt.Printf("last insert=%d \n", last)
	affect, _ := result.RowsAffected()
	fmt.Printf("affect rows=%d \n", affect)
}
func update(db *sql.DB) {
	stmt, err := db.Prepare("update student set score=score+? where city=?")
	common.CheckErr(err)
	result, _ := stmt.Exec(20, "朝阳")
	rows, _ := result.RowsAffected()
	fmt.Printf("affect rows=%d \n", rows)
}
func delete(db *sql.DB) {
	stmt, err := db.Prepare("delete from student where id=?")
	common.CheckErr(err)
	_, err = stmt.Exec(20)
	common.CheckErr(err)
}
func query(db *sql.DB) {
	stmt, err := db.Prepare("select name,province,city,enrollment from student where id=?")
	common.CheckErr(err)
	rows, err := stmt.Query(22)
	common.CheckErr(err)
	var name, prov, city, enroll string
	for rows.Next() {
		rows.Scan(&name, &prov, &city, &enroll)
		fmt.Printf("name=%s,province=%s,city=%s,enroll=%s\n", name, prov, city, enroll)
	}
}

func hugeInsert(db *sql.DB) {
	stmt, err := db.Prepare("insert into student(name,province,city,enrollment)values(?,?,?,?)")
	common.CheckErr(err)
	begin := time.Now()

	for i := 0; i < 100000; i++ {
		_, err := stmt.Exec("宋江"+strconv.Itoa(i+1), "山东", "梁山", "2010-12-23")
		common.CheckErr(err)
	}
	fmt.Printf("time=%d ms\n", time.Since(begin).Milliseconds())
}
func main() {
	db, err := sql.Open("mysql", "teng:5570709@/test?charset=utf8")
	common.CheckErr(err)
	//insert(db)
	//replace(db)
	//update(db)
	//delete(db)
	//query(db)
	hugeInsert(db)
}
