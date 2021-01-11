package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type Welcome struct {
	Name string
	Time string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST %s", r.URL.Path[1:])
}

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", nil)
	})
	user := os.Getenv("QOVERY_DATABASE_TEST_DB_USERNAME")
	pass := os.Getenv("QOVERY_DATABASE_TEST_DB_PASSWORD")
	host := os.Getenv("QOVERY_DATABASE_TEST_DB_HOST")
	connectionString := fmt.Sprintf("%s:%s@%s:3306/test_db",user,pass,host)
	cnn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)

	router.Run()
}
