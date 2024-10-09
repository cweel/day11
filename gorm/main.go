package main

import (
	"day11/common"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	Province   string
	City       string
	Address    string    `gorm:"column:addr"`
	Score      float32   `gorm:"column:score"`
	Enrollment time.Time `gorm:"column:enrollment;type:date"`
}

func (Student) TableName() string {
	return "student"
}
func query(db *gorm.DB) {
	var student Student
	db.Where("city=?", "北京").First(&student) //返回一条记录
	fmt.Println(student.Name)
	//返回多条记录
	var students []Student
	db.Where("city=?", "华强北").Find(&students)
	for _, ele := range students {
		fmt.Printf("id=%d , name=%s\n", ele.Id, ele.Name)
	}

	db.Where("city in ?", []string{"滕州", "海淀"}).Find(&students)
	for _, ele := range students {
		fmt.Printf("id=%d , name=%s \n", ele.Id, ele.Name)
	}
	student = Student{}
	db.First(&student, 1)
	println(student.Name)
}
func update(db *gorm.DB) {
	db.Model(&Student{}).Where("city=?", "北京").Update("score", 10)
	db.Model(&Student{}).Where("city=?", "北京").Updates(map[string]interface{}{"score": 3, "addr": "海淀区"})
}

// 插入一条记录
func create(db *gorm.DB) {
	student := Student{Name: "光绪9", Province: "北京", City: "北京", Score: 38, Enrollment: time.Now()}
	db.Create(&student)
	//插入两条记录
	students := []Student{{Name: "无极", Province: "河南", Enrollment: time.Now()}, {Name: "有道", Province: "广东", Enrollment: time.Now()}}
	db.Create(students)
}
func delete(db *gorm.DB) {
	db.Where("city in ?", []string{"济南", "朝阳"}).Delete(&Student{})
	db.Delete(&Student{}, 120867)
	db.Delete(&Student{}, []int{120869, 120870})
}

// 事务
func transaction(db *gorm.DB) {
	var student Student
	tx := db.Begin()
	for i := 0; i < 10; i++ {
		student = Student{Name: "无情" + strconv.Itoa(i), Province: "北京", Enrollment: time.Now()}
		tx.Create(&student)
	}
	tx.Commit()
}
func main() {
	dsn := "teng:5570709@tcp(chent7.top:3306)/test?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	common.CheckErr(err)
	//query(db)
	//update(db)
	//create(db)
	//delete(db)
	transaction(db)
}
