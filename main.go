package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []task

var Task = allTask{
	{ID: 1, Name: "Johan", Content: "This is the content"},
	{ID: 2, Name: "Josue", Content: " The second content"},
}
func getTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Task)
}
func createTask(w http.ResponseWriter, r *http.Request){
	//Creat a task to recieve the client Data  
	var newTask task
	requestBody ,err:=ioutil.ReadAll((r.Body))	
	if err!= nil{
		fmt.Fprint(w,"Write a valid task")
	}
	json.Unmarshal(requestBody,&newTask)
	//In this Line set the id on "Automatic"
	newTask.ID = len(Task)-1
	//You add the new adquire task to the slice of tasks
	Task =append(Task,newTask)	
	//You can sent new information that maight be usefull with a header
	w.Header().Set("Content-Type", "application/json")
	//Sending a status code to verified that everithing was sucesfull
	w.WriteHeader(http.StatusCreated)
	//Use the json.Encodeer 
	json.NewEncoder(w).Encode(newTask)

}
func indexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1 style="text-align:center">!Bienvenido al API!</h1>`)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/PP", indexRouter)
	router.HandleFunc("/task",getTask)
	log.Print((http.ListenAndServe(":8080", router)))

}
