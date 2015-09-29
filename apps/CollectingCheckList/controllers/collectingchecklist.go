package controllers

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
)

type Hello struct {
	Message string
	id string
}

type Staff struct {
	Id int64
	Name string
}

type Config struct {
	Db RdsMysqlConfig `json:"rds_mysql"`
}

type RdsMysqlConfig struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Url  string `json:"url"`
}

func CCLHome(w http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(w,"CCLHome")
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
	absPath, _ := filepath.Abs("./apps/config.json")
	file,err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	var config Config
	json.Unmarshal(file, &config)

	db,err := sql.Open("mysql",config.Db.User+":"+config.Db.Pass+"@tcp("+config.Db.Url+")/mushroom_db")
	if err != nil {
		log.Fatal("db connection error:",err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM staff_record")
	if err != nil {
		log.Fatal("show tables error:",err)
	}
	defer rows.Close()

	staff := []Staff{}

	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("show scant staff_record error:",err)
		}
		stf := Staff{Id:id,Name:name}
		staff = append(staff,stf)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(staff)

}