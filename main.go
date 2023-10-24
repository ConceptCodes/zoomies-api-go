package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	"log"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define my types
type pet struct {
	Name    string `json:"name"`
	OwnerId int    `json:"owner_id"`
	Age     int    `json:"age"`
}

type user struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	IsVet    bool   `json:"is_vet"`
	Pets     []pet  `json:"pets"`
}

var db *sql.DB

func init() {
	var err error
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=verify-full"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func getJwtToken() string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email != "admin" || password != "password" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

	var user user

	db := getDb()

	err := db.QueryRow("SELECT username, name, age, password, is_vet FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Username, &user.Name, &user.Age, &user.Password, &user.IsVet)

	if err != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"message": "Invalid email or password"})
	}

	log.Println(user)

	c.JSON(200, gin.H{"message": "Login successful"})
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	name := c.PostForm("name")
	age := c.PostForm("age")
	email := c.PostForm("email")
	password := c.PostForm("password")

	db := getDb()

	_, err := db.Exec("INSERT INTO users (username, name, age, email, password) VALUES ($1, $2, $3, $4, $5)", username, name, age, email, password)

	if err != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"message": "Invalid email or password"})
	}

	c.JSON(200, gin.H{"message": "Registration successful"})

}

func main() {
	router := gin.Default()

	router.POST("/auth/login", login)
	router.POST("/auth/register", register)

	router.Run()
}

func getDb() *sql.DB {
	return db
}
