package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// middleware
func AuthMiddleware(c *gin.Context) {
	if c.Request.URL.Path != "/" && c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		// c.AbortWithStatusJSON is similar to (abort in flask)
		return
	}
	c.Next() // proceed to the next middleware
}

func LogMiddleware(c *gin.Context) {
	log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path) //%s %s is used to print the request method and path
	c.Next()
	log.Printf("Response: %d", c.Writer.Status())
}


var users []map[string] string // slice of maps

// check if user exists
func userExists(newUser map[string]string) bool{
	for _, user := range users{ // "_" is used to ignore the index
		if user["id"] == newUser["id"] || user["name"] == newUser["name"]{
			return true
		}
	}
	return false
}


func main(){
	router := gin.Default()
	router.Use(AuthMiddleware, LogMiddleware) // middlewares applied globally
	// router.Use is similar to (@app.before_request in flask)

	router.GET("/", func(c *gin.Context) {
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

		if userExists(new_user){
			c.JSON(409, gin.H{"message": "User already exists", "status": "error", "status_code": 409})
			return
		}

		users = append(users, new_user)
		c.JSON(201, gin.H{"user": new_user, "message": "User added successfully", "status": "success", "status_code": 201})
		
	})

	router.Run()  // :8080 default port
}















// // route to get all users

// func getUsers(c *gin.Context){
// 	router := gin.Default()
// 	router.GET("/users", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "Users retrieved successfully", "status": "success", "status_code": 200})
// 	})

// 	router.Run()

// }