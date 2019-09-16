package main

import( 
	//_ "github.com/go-sql-driver/mysql"
	//"database/sql"
	"fmt"
	//"log"
	"time"
	)


func main() {
	/*cnn, err := sql.Open("mysql", "quitSmokeDev:Proyecto098.@tcp(82.223.0.98:3306)/quite_smoke_dev")
	if err != nil {
			log.Fatal(err)
	}

	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
			log.Fatal(err)
	}

	fmt.Println(id, name)*/
	//TIME TO 20
	ch := time.After(20 * time.Second)
	defer (func() { fmt.Println("waiting"); <-ch; fmt.Println("waited") })()
	
	fmt.Println("Starting web server on port 8080")

}