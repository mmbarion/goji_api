package controllers

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
)

type Hello struct {
	Message string
	id string
}

func CCLHome(w http.ResponseWriter,req *http.Request) {
	//fmt.Fprintf(w,"CCLHome")
	w.Header().Set("Content-Type","application/json")

	hello := Hello{
		Message: "Hello, World!",
		id: "1",
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(hello)
}

func AddCheckList(w http.ResponseWriter, r *http.Request) {
	pp.Print(r)
}

func GetStaff(w http.ResponseWriter,req *http.Request) {
	db,err := sql.Open("mysql","mushroom_db_user:mmn39504@tcp(mushroominstance.cxneilncg8ry.ap-northeast-1.rds.amazonaws.com:3306)/mushroom_db")
	if err != nil {
		log.Fatal("db connection error:",err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM staff_record")
	if err != nil {
		log.Fatal("show tables error:",err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("show scant staff_record error:",err)
		}
		fmt.Println(id,name)
	}

}