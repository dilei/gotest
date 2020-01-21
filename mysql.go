package main

import (
	"fmt"

	"database/sql"                     // 这是一个抽象层包，比如区分mysql、orcal等数据库，只有这个包是连接不上mysql的，还需要搭配下面的mysql包
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
)

func main() {
	query()
}

func query() {
	db, err := sql.Open("mysql", "writeuser:ddbackend@tcp(10.255.255.22)/ProductDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	/*
	   //增加数据
	   stmt, err := db.Prepare(`INSERT student (name,age) values (?,?)`)
	   res, err := stmt.Exec("wangwu", 26)
	   id, err := res.LastInsertId()
	   fmt.Println("自增id=", id)
	   //修改数据
	   stmt, err := db.Prepare(`UPDATE student SET age=? WHERE id=?`)
	   res, err := stmt.Exec(21, 5)
	   num, err := res.RowsAffected() //影响行数
	   fmt.Println(num)
	   //删除数据
	   stmt, err := db.Prepare(`DELETE FROM student WHERE id=?`)
	   res, err := stmt.Exec(5)
	   num, err := res.RowsAffected()
	   fmt.Println(num)
	*/

	//查询数据
	rows, err := db.Query("SELECT product_id, product_name FROM Products_Core limit 3")

	//--------简单一行一行输出---start
	//    for rows.Next() { //满足条件依次下一层
	//        var id int
	//        var name string
	//        var age int
	//        rows.Columns()

	//        err = rows.Scan(&id, &name, &age)
	//        fmt.Println(id)
	//        fmt.Println(name)
	//        fmt.Println(age)
	//    }
	//--------简单一行一行输出---end

	//--------遍历放入map----start
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	//--------遍历放入map----end
}
