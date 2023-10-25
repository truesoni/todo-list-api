package main

import (
     "net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

 type task struct{
	ID		string	
	Item		string	
	Completed bool		
 }


 var tasks = []task{
	{ID: "1", Item: "Drink Water", Completed: true},
	{ID: "2", Item: "Clean Room", Completed: true},
	{ID: "3", Item: "Hygiene Yourself", Completed: true},
	{ID: "4", Item: "Practice Gratitude", Completed: true},
	{ID: "5", Item: "Read Book", Completed: true},
	{ID: "6", Item: "Practice Coding", Completed: true},
	{ID: "7", Item: "Read Documentation", Completed: true},

 }

 func gettasks(context *gin.Context){

	context.IndentedJSON(http.StatusOK, tasks)
 }
 
 func addtask(context *gin.Context){
	var newTask task
	
	if err := context.BindJSON(&newTask); err != nil{
		return
	}

	tasks = append(tasks, newTask)
	context.IndentedJSON(http.StatusCreated, newTask)
 }

func gettaskById(id string) (*task, error){
	 for i, t := range tasks{
		if t.ID == id {
			return &tasks[i], nil
		}
	 }

	 return nil, errors.New("tasks not exist")
}	

func gettask(context *gin.Context){
	id := context.Param("id")
	task, err := gettaskById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "task does not exist"})
		return 
	}

	context.IndentedJSON(http.StatusOK, task)
}

func toggletaskProgress(context *gin.Context ) {
	id := context.Param("id")
	task, err := gettaskById(id)
	
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "task does not exist"})
		return 
	}
	task.Completed = !task.Completed
	context.IndentedJSON(http.StatusOK, task)
}

 func main(){
	// creating a server
	router := gin.Default()

	// HTTP Verb - GET (*)
	router.GET("/tasks", gettasks)
	
	// HTTP Verb - GET (specific)
	router.GET("/tasks/:id", gettask)

	// HTTP Verb - PATCH(to change in specific attribute of an object)

	router.PATCH("/tasks/:id", toggletaskProgress)
	// HTTP Verb - POST
	router.POST("/tasks", addtask)

	// To run our server 
	router.Run("localhost:3000")
 }