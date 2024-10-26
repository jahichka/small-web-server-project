package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type korisnik struct {
	Plastenik string `json:"plastenik"`
	Biljka    string `json:"biljka"`
}
type Message struct {
	Data string `json:"message"`
}
type Manager struct {
	DB *sql.DB
}

func main() {
	var err error
	manager := &Manager{}
	manager.DB, err = sql.Open("mysql", "amina:password@tcp(127.0.0.1:3306)/tkm")
	if err != nil {
		panic(err)
	}
	defer manager.DB.Close()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/dashboard", manager.login)
	router.POST("/add", manager.add)
	router.Run(":8080")
}

func fetchData(username string, manager Manager) ([]korisnik, error, string) {
	var userData []korisnik
	var str string
	rows, err := manager.DB.Query("SELECT idPlasenik,biljka FROM korisnik WHERE idUser = ?", username)
	if err != nil {
		str = "DB"
		return nil, err, str
	}
	defer rows.Close()

	for rows.Next() {
		usr := korisnik{}
		err := rows.Scan(&usr.Plastenik, &usr.Biljka)
		str = "scan" + usr.Biljka + usr.Plastenik
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err, str
		}
		userData = append(userData, usr)
		str = "append"
	}

	if err := rows.Err(); err != nil {
		str = "ovo peto"
		return nil, err, str
	}

	return userData, nil, str
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func (manager *Manager) login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var dbUsername, dbPassword string
	var usr []korisnik

	err := manager.DB.QueryRow("SELECT idUser, password FROM user WHERE idUser = ?", username).Scan(&dbUsername, &dbPassword)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, &Message{"Incorrect username"})
		return
	}

	if password != dbPassword {
		c.IndentedJSON(http.StatusUnauthorized, &Message{"Incorrect password"})
		return
	}

	var str string
	usr, err, str = fetchData(dbUsername, *manager)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, &Message{str})

		return
	}

	c.IndentedJSON(http.StatusOK, usr)
}

func (manager *Manager) add(c *gin.Context) {
	type korisnik1 struct {
		Username  string `json:"username"`
		Plastenik string `json:"plastenik"`
		Biljka    string `json:"biljka"`
	}
	var usr korisnik1
	var err error
	if err := c.BindJSON(&usr); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		println(err)
		return
	}
	println("!!!", usr.Biljka)
	manager.DB.Query("INSERT INTO korisnik VALUES (?, ?, ?)",
		usr.Username, usr.Plastenik, usr.Biljka)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, usr)
}
