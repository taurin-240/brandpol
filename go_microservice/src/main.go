package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func greets(c *gin.Context) {
	name := c.Query("name")
	if name != "" {
		c.String(http.StatusOK, "Привет %s от Go!", name)
		db := connectToPostgreSQL()
		fmt.Println(save_greet(db, name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameters is failed"})
	}
}
func history(c *gin.Context) {
	db := connectToPostgreSQL()
	result := get_history(db)
	c.IndentedJSON(http.StatusOK, result)

}
func pythonGreets(c *gin.Context) {
	res := getPythonGreets()
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	db := connectToPostgreSQL()
	fmt.Println(db)

	// router
	router := gin.Default()
	router.GET("/greet/", greets)
	router.GET("/greet/history", history)
	router.GET("/python/greets", pythonGreets)

	router.Run(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
}
