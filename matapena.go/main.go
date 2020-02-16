package main

import (
	"app/etc/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		if args[0] == "-h" {
			cmdH()
		} else if args[0] == "-start" {
			cmdStart("local")
		} else {
			println("Invalid Arguments [" + args[0] + "]!. Check help with -h")
		}
	} else if len(args) == 2 {
		if args[0] == "-start" {
			if args[1] == "local" || args[1] == "dev" || args[1] == "prod" {
				cmdStart(args[1])
			} else {
				println("Invalid Arguments [" + args[1] + "]!. Check help with -h")
			}
		} else {
			println("Invalid Arguments [" + args[0] + " + " + args[1] + "]!. Check help with -h")
		}
	} else {
		println("Invalid Arguments!. Check help with -h")
	}
}

func cmdH() {
	println("--------------------------------------")
	println("|               Tools                |")
	println("--------------------------------------")
	println("|  -h                 Help           |")
	println("|  -start             Run Local Env  |")
	println("|  -start <env>       Run Env        |")
	println("|                                    |")
	println("--------------------------------------")
	println("|            Environment             |")
	println("--------------------------------------")
	println("|   local             Local          |")
	println("|   dev               Development    |")
	println("|   prod              Production     |")
	println("--------------------------------------")
}

func cmdStart(evn string) {
	setupConfig(evn)
	db := setupDbConnection()
	router.InitRoute(db)
	defer db.Close()
}

func setupConfig(arg string) {
	var configFile string
	switch arg {
	case "local":
		{
			configFile = "config.json"
		}
	case "dev":
		{
			configFile = "config-dev.json"
		}
	case "prod":
		{
			configFile = "config-prod.json"
		}
	}

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	println("Config Env : " + arg)
	println("Config File : " + configFile)
	println("Config Port : " + viper.GetString("server.address"))
}

func setupDbConnection() *gorm.DB {
	println("setupDbConnection")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbname := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open("mysql", connectionString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	return db
}
