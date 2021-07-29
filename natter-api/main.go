package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/natter?multiStatements=true")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	db = conn

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/spaces", createSpace).Methods("POST")

	log.Fatal(http.ListenAndServe(":4567", router))
}

type Space struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Uri   string `json:"uri,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func createSpace(w http.ResponseWriter, r *http.Request) {
	// parse JSON body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var space Space
	_ = json.Unmarshal(reqBody, &space)

	// validation
	if len(space.Name) > 255 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "spaces name to long"})
		return
	}

	match, _ := regexp.Match("[a-zA-Z] [a-zA-Z0-9] {1,29}", []byte(space.Owner))
	if !match {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid username" + space.Owner})
		return
	}

	// perform database transaction
	log.Println("INSERT INTO spaces(name, owner) VALUES('" + space.Name + "','" + space.Owner + "')")
	result, err := db.Exec("INSERT INTO spaces(name, owner) VALUES(?,?)", space.Name, space.Owner)
	if err != nil {
		log.Println(fmt.Errorf("insert error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	id, _ := result.LastInsertId()
	space.Uri = fmt.Sprintf("/spaces/%d", id)

	// Response back
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(space)
}
