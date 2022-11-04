package main

import (
	"fmt"
	"os"

	"github.com/polatyener-dev/golang_envfile/utils"
)

func init() {
	utils.LoadEnvVariable()
}

func main() {

	mysqlDb := os.Getenv("MYSQL_DB")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASS")

	fmt.Printf("DBHOST:=%s\nDBNAME=%s\nDBUSER=%s\nDBPASS=%s", mysqlHost, mysqlDb, mysqlUser, mysqlPassword)

}
