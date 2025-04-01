package main

import (
	"log"
	"strconv"
	// "os/exec"
	// "context"
	"database/sql"

	"github.com/gin-gonic/gin"

	// "github.com/jackc/pgx/v4"
	_ "github.com/mattn/go-sqlite3"
)

//postgresql://postgres:spookie@localhost:5432/postgres
// var db *pgx.Conn
// func initDB() {
// 	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:spookie@localhost:5432/postgres")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db = conn
// }

//sqlite3
var db *sql.DB
func initDB(){
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func execute_and_serialize(query string) ([]map[string]string, error) {
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []map[string]string
    for rows.Next() {
        var id int64
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            return nil, err
        }
        results = append(results, map[string]string{"id": strconv.FormatInt(id, 10), "name": name})
    }

    return results, nil
}

func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, name FROM user")

	if err != nil {
		c.JSON(500, gin.H{"error": "DB error"});
		return
	}

	var users []map[string]string
	for rows.Next(){
		var id int64
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			c.JSON(500, gin.H{"error": "DB error"});
			return
		}
		users = append(users, map[string]string{"id": strconv.FormatInt(id, 10), "name": name})
	}
	c.JSON(200, gin.H{"users": users, "message": "Users retrieved successfully", "status": "success", "status_code": 200})
}

// add user to the database and return all users using sql query
func addUser(c *gin.Context) {
	var newUser map[string]string
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if userExists(newUser){
		c.JSON(409, gin.H{"error": "User already exists"})
		return
	}

	_, err := db.Exec("INSERT INTO user (id, name) VALUES (?,?)", newUser["id"], newUser["name"])
	if err != nil {
		c.JSON(500, gin.H{"error": "DB error"})
		return
	}

	results, err := execute_and_serialize("SELECT * FROM user")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "DB error"})
		return
	}

	c.JSON(201, gin.H{"users": results, "message": "User added successfully", "status": "success", "status_code": 201})
}



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


// check if user exists in the database
func userExists(user map[string]string) bool {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM user WHERE id = ? OR name = ?", user["id"], user["name"]).Scan(&count)
    if err != nil {
        log.Println(err)
        return false
    }
    return count > 0
}





// main
func main(){
	router := gin.Default()
	router.Use(AuthMiddleware, LogMiddleware) // middlewares applied globally
	// router.Use is similar to (@app.before_request in flask)

	initDB()
	defer db.Close()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running", "status": "success", "status_code": 200})
	})

	router.GET("/users", getUsers)
	router.POST("/users", addUser)

	//setting port
	router.Run(":3000")
}