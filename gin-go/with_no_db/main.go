package main

import (
	"log"
	// "strconv"
	// "os/exec"
	// "context"
	// "database/sql"

	"github.com/gin-gonic/gin"
)

var users []map[string]string // slice of maps

//check if user exists
func user_Exists_in_slice(user map[string]string) bool {
	for _, u := range users {
		if u["id"] == user["id"] || u["name"] == user["name"] {
			return true
		}
	}
	return false
}


// main
func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		log.Println("Server is running")
		c.JSON(200, gin.H{"message": "Server is running", "status": "success", "status_code": 200})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"users": users, "message": "Users retrieved successfully", "status": "success", "status_code": 200})
	})

	router.POST("/users", func(c *gin.Context) {
		var new_user map[string]string
		if err := c.BindJSON(&new_user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request", "status": "error", "status_code": 400})
			return
		}

		if user_Exists_in_slice(new_user){
			c.JSON(409, gin.H{"message": "User already exists", "status": "error", "status_code": 409})
			return
		}

		users = append(users, new_user)
		c.JSON(201, gin.H{"user": new_user, "message": "User added successfully", "status": "success", "status_code": 201})
		
	})

	//setting port
	router.Run(":3000")
}
