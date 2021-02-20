package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)


type MySQLOps struct {
	mysql *sql.DB
}

type MySQLConfig struct {
	username string
	password string
	network  string
	server   string
	port     int
	database string
}

func GetOps(config MySQLConfig) *MySQLOps {
	connection := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.username, config.password, config.network, config.server, config.port, config.database)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(fmt.Errorf("MySQL Ops Init Error, %s \n", err))
	}
	return &MySQLOps{
		mysql: db,
	}
}


func (ops *MySQLOps) query(qSql string, qType interface{})  {
	rt := reflect.TypeOf(qType)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fmt.Printf("field's name: %s, type: % v\n", field.Name, field.Type)
	}

	result, err := ops.mysql.Query(qSql)
	if err != nil {
		fmt.Println("query error: %s \n", err)
	}


	result.Columns()


	result.Close();

}

func (ops *MySQLOps) delete(dSql string) int {


	return 0
}


func (ops *MySQLOps) insert(iSql string) int {

	return 0
}


func (ops *MySQLOps) update(uSql string) int {

	return 0
}







//
///**
//初始化MySQL Ops
//*/
//func init() {
//	viper.SetConfigName("database")   // 设置配置名称
//	viper.SetConfigType("yaml")       // 设置配置类型
//	viper.AddConfigPath("./database") // 从项目根路径查找
//
//	err := viper.ReadInConfig() // 查找并读取配置文件
//	if err != nil {
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//
//	username := viper.Get("mysql.username").(string)
//	password := viper.Get("mysql.password").(string)
//	network := viper.Get("mysql.network").(string)
//	server := viper.Get("mysql.server").(string)
//	port := viper.Get("mysql.port").(string)
//	database := viper.Get("mysql.database").(string)
//	connection := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database)
//}

func main() {
	fmt.Println("111")
}

//
//
//
//
//func main() {
//	fmt.Println("Go MySQL Tutorial")
//
//	// Open up our database connection.
//	// I've set up a database on my local machine using phpmyadmin.
//	// The database is called testDb
//	db, err := sql.Open("mysql", ":@tcp(:)/")
//
//	// if there is an error opening the connection, handle it
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// defer the close till after the main function has finished
//	// executing
//	defer db.Close()
//
//	// Execute the query
//	rows, err := db.Query("SELECT * FROM data_market_leads_assigin_info limit 200")
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//
//	// Get column names
//	columns, err := rows.Columns()
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//
//	// Make a slice for the values
//	values := make([]sql.RawBytes, len(columns))
//
//	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
//	// references into such a slice
//	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
//	scanArgs := make([]interface{}, len(values))
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//
//	// Fetch rows
//	for rows.Next() {
//		// get RawBytes from data
//		err = rows.Scan(scanArgs...)
//		if err != nil {
//			panic(err.Error()) // proper error handling instead of panic in your app
//		}
//
//		// Now do something with the data.
//		// Here we just print each column as a string.
//		var value string
//		for i, col := range values {
//			// Here we can check if the value is nil (NULL value)
//			if col == nil {
//				value = "NULL"
//			} else {
//				value = string(col)
//			}
//			fmt.Println(columns[i], ": ", value)
//		}
//		fmt.Println("-----------------------------------")
//	}
//	if err = rows.Err(); err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//}
