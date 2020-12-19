package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
)

var DB *sql.DB

//无用的全局变量
var Result map[int]map[string]string

//无用的结构体
type info struct {
	id          string `db:"id"`
	name        string `db:"name"`
	password    string `db:"password"`
	level       int    `db:"level"`
	groups      string `db:"groups"`
	phonenumber string `db:"phonenumber"`
	photo       string `db:"photo"`
	description string `db:"description"`
	grade       string `db:"grade"`
}

//连接数据库的方法
func ConnectDB(connString string) {
	//根据获取的参数连接数据库
	db, _ := sql.Open("mysql", connString)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	fmt.Println("连接成功")
	DB = db
}

func GetAllRows(rows *sql.Rows) map[int]map[string]string {
	//返回列名称组成的slice，也就是字段名的合集
	columns, _ := rows.Columns()
	//fmt.Println(columns)
	//vals用来存放取出来的数据结果，表示一行所有列的值，后边的长度表示行数
	vals := make([][]byte, len(columns))
	//做rows,scan的参数，将扫描后的数据存储在scans中
	scans := make([]interface{}, len(columns))
	//将每一行的数据填充到[][]byte中
	for k, _ := range vals {
		//因为rows，sacns参数是指针类型，通过遍历将指针变量保存到scans切片中去，相当于buf := []byte{&id，&name}
		//fmt.Println(&vals[k])
		scans[k] = &vals[k]
	}

	result := make(map[int]map[string]string) //key作为column中的字段名，值为字段名和记录

	i := 0
	for rows.Next() {
		//将查询到的记录结果放在scans[]中
		err := rows.Scan(scans...)
		if err != nil {
			fmt.Println(err)
		}
		//因为是byte的切片，所以循环取出转换成string
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]    //字段名
			row[key] = string(v) //原来是byte类型，现在转换为string类型
		}
		//放入总的结果当中，i用来记录读取的条数
		result[i] = row
		i++
	}
	return result
}

func GetAllDB() map[int]map[string]string {
	sql1 := "SELECT * FROM user ORDER BY timestamp DESC "
	//执行SQL语句查询
	rows, err := DB.Query(sql1)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	result := GetAllRows(rows)
	return result
}

//暂时无用的方法
func GetResult() {
	result := GetAllDB()
	Result = result
}

//定时任务，或许以后用得上
func TimmingMission() *cron.Cron {
	c := cron.New()
	fmt.Println("开始了")
	c.AddFunc("*/5 * * * * ?", func() {
		GetResult()
	})
	return c
}
