package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	//Creat a task to recieve the client Data
	var newTask task
	requestBody, err := ioutil.ReadAll((r.Body))
	if err != nil {
		fmt.Fprint(w, "Write a valid task")
	}
	json.Unmarshal(requestBody, &newTask)
	//In this Line set the id on "Automatic"
	newTask.ID = len(Task) + 1
	//You add the new adquire task to the slice of tasks
	Task = append(Task, newTask)
	//You can sent new information that maight be usefull with a header
	w.Header().Set("Content-Type", "application/json")
	//Sending a status code to verified that everithing was sucesfull
	w.WriteHeader(http.StatusCreated)
	//Use the json.Encodeer
	json.NewEncoder(w).Encode(newTask)
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	/*In this example, mux.Vars(r) extracts the value of the "id" path parameter from the URL,
	which can be accessed via the vars map. This allows you to dynamically handle different
	 URLs and extract specific parameters from them.*/
	vars := mux.Vars(r)
	//Using the strconv to turn this string into a Integer
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "There has been an error idientifying the Id")
		return
	}
	for _, val := range Task {
		if val.ID == taskId {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode("The task that you asked for does exits")
		}
	}
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "That id is not valid")
	}
	for i, val := range Task {
		if taskId == val.ID {
			w.Header().Set("Content-type", "applicatio/json")
			json.NewEncoder(w).Encode("We found your task")
			//You cut the slice and then add it back to the Original 
			Task = append(Task[:i], Task[i+1:]...)
		}
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	var updatedTask task
	if err != nil {
		fmt.Fprint(w, "Invalid id")
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "The data was not Valid")
	}
	//Now use Unmarshal to asign the data in to go structs
	json.Unmarshal(reqBody, &updatedTask)

	for i, val := range Task {
		if taskId == val.ID {
			//It founds the id in the Task[]
			//You cut the slice and then add it back
			Task = append(Task[:i], Task[i+1:]...)
			//Update the task id with the Inide the user send
			updatedTask.ID = taskId
			//(Update)Now you add it back to the slice
			Task = append(Task, updatedTask)

			fmt.Fprint(w, "The task with the id %v has been updated succesfully.", taskId)
			fmt.Fprint(w, "The task is not been really Filter")
		}
	}

}
func indexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1 style="text-align:center">!Bienvenido al API!</h1>`)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/PP", indexRouter)
	router.HandleFunc("/task", getTask).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	log.Print((http.ListenAndServe(":8080", router)))
}
