package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"flag"
)

var (
	server         string
	port           string
	configFilePath = flag.String("configFilePath", "/work/code/gopath/src/godemo/conf/config.json", "config file path")
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Address string `json:"address"`
}

func main() {
	flag.Parse()

	config, err := initConfigFromFile(*configFilePath)
	if err != nil {
		fmt.Println("initConfigFromFile err: %s" + err.Error())
	}
	server = config.Server
	port = config.Port

	StartServer()
}

//func StartServer() {
//	lis, err := net.Listen("tcp", server+":"+port)
//	if err != nil {
//		//glog.Error("listen port: "+port+" error")
//		fmt.Println("listen port err: " + err.Error())
//	}
//
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/qry", proQry)
//
//	srv := &http.Server{
//		Addr:         server + ":" + port,
//		Handler:      mux,
//		WriteTimeout: 20 * time.Second,
//	}
//
//	srv.Serve(lis)
//}

func StartServer() {
	http.HandleFunc("/qry", proQry)
	http.ListenAndServe(server+":"+port, nil)
}

func proQry(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	id := r.Form["id"][0]

	db, err := sql.Open("mysql", "zhaoli:zhaoli@tcp(192.168.56.101:3306)/zhaoli?charset=utf8")
	if err != nil {
		fmt.Println("open db err:" + err.Error())
	}

	rows, err := db.Query("select id,name,age,sex,address from user where id=?", id)
	if err != nil {
		fmt.Println("query db err: " + err.Error())
	}

	var user_id int
	var user_name string
	var user_age int
	var user_sex string
	var address string
	var result = []User{}

	for rows.Next() {
		rows.Scan(&user_id, &user_name, &user_age, &user_sex, &address)
		tmpUser := new(User)
		tmpUser.Id = user_id
		tmpUser.Name = user_name
		tmpUser.Age = user_age
		tmpUser.Sex = user_sex
		tmpUser.Address = address

		result = append(result, *tmpUser)
	}

	data, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Println("json.marshal err: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	w.WriteHeader(200)
	//w.Write([]byte("hello go!"))
}
