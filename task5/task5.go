package main 

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
)

var database *sql.DB

type Message struct{
	Id uint16 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Images string `json:"images"`
	Cost uint16 `json:"cost"`
	Timestamp string `json:"timestamp"`
}

func create_page(w http.ResponseWriter, r *http.Request){
	temp, err:= template.ParseFiles("templates/create_page.html")
	if err != nil{log.Fatal(err)}

	if r.Method == "POST"{
		name:= r.FormValue("name")
		description:= r.FormValue("description")
		images:= r.FormValue("images")
		cost:= r.FormValue("cost")
		imagesCount := len(strings.Split(images, ","))

		if name == "" || description == "" || images == "" || cost == ""{
			err:= make(map[string]string)
			err["error"] = "One or more fields were skipped"
			js, _ := json.Marshal(err)
			w.WriteHeader(400)
			w.Write(js)
		}else{
			if len(name) > 200{
				err:= make(map[string]string)
				err["error"] = "name - The length of the string exceeds 200 characters"
				js, _ := json.Marshal(err)
				w.Write(js)
				w.WriteHeader(400)
			}else if len(description) > 1000{
				err:= make(map[string]string)
				err["error"] = "description - The length of the string exceeds 1000 characters"
				js, _ := json.Marshal(err)
				w.WriteHeader(400)
				w.Write(js)
			}else if imagesCount > 3{
				err:= make(map[string]string)
				err["error"] = "images - You have specified more than three links to the photo (check the commas)"
				js, _ := json.Marshal(err)
				w.WriteHeader(400)
				w.Write(js)
			}else{
				insert, err:= database.Exec("INSERT INTO messages (name, description, images, cost, timestamp) VALUES (?, ?, ?, ?, NOW())", name, description, images, cost)
				if err != nil{log.Fatal(err)}
				id, _:= insert.LastInsertId()
				resp:= make(map[string]int64)
				resp["id"] = id
				js, _ := json.Marshal(resp)
				w.WriteHeader(301)
				w.Write(js)
			}
		}
	}else if r.Method == "GET"{
		w.WriteHeader(200)
		temp.Execute(w, nil)
	}
}

func index_page(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var err error
	var res *sql.Rows

	if r.URL.Query().Get("sort") == "date"{
		res, err = database.Query("SELECT * FROM messages ORDER BY timestamp")
		if err != nil{log.Fatal(err)}
	}else if r.URL.Query().Get("sort") == "undate"{
		res, err = database.Query("SELECT * FROM messages ORDER BY timestamp DESC")
		if err != nil{log.Fatal(err)}
	}else if r.URL.Query().Get("sort") == "cost"{
		res, err = database.Query("SELECT * FROM messages ORDER BY cost")
		if err != nil{log.Fatal(err)}
	}else if r.URL.Query().Get("sort") == "uncost"{
		res, err = database.Query("SELECT * FROM messages ORDER BY cost DESC")
		if err != nil{log.Fatal(err)}
	}else{
		res, err = database.Query("SELECT * FROM messages")
		if err != nil{log.Fatal(err)}		
	}

	var messages []Message
	for res.Next(){
		var msg Message
		err = res.Scan(&msg.Id, &msg.Name, &msg.Description, &msg.Images, &msg.Cost, &msg.Timestamp)
		if err != nil{log.Fatal(err)}
		messages = append(messages, msg)
	}
	js, _ := json.Marshal(messages)
	w.WriteHeader(200)
	w.Write(js)
}

func getById(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	res, err:= database.Query(fmt.Sprintf("SELECT * FROM messages WHERE id = %s", vars["id"]))
	if err != nil{log.Fatal(err)}

	var messages []Message
	for res.Next(){
		var msg Message
		err = res.Scan(&msg.Id, &msg.Name, &msg.Description, &msg.Images, &msg.Cost, &msg.Timestamp)
		if err != nil{log.Fatal(err)}

		img:= strings.Split(msg.Images, ",")
		msg.Images = img[0]
		messages = append(messages, msg)
	}
	js, _ := json.Marshal(messages)
	w.WriteHeader(200)
	w.Write(js)
}

func handleRequest(){
	router:= mux.NewRouter()
	router.HandleFunc("/id/{id:[0-9]+}", getById)
	router.HandleFunc("/", index_page)
	router.HandleFunc("/create", create_page)
	http.Handle("/", router)
	http.ListenAndServe(":88", nil)
}

func main(){
	db, err:= sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil{log.Fatal(err)}

	database = db
	defer db.Close()

	fmt.Println("Server is lestening...")
	handleRequest()
}
