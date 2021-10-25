package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// card represents data about a record card.
type card struct {
	ID    string `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
}

// cards slice to seed record card data.
var cards = []card{
	{ID: "1", Front: "Blue Train", Back: "John Coltrane"},
}

func main() {
	router := gin.Default()
	router.GET("/cards", getCards)
	router.GET("/cards/:id", getCardByID)
	router.POST("/cards", postCards)

	router.Run("localhost:8080")
}

// getCards responds with the list of all cards as JSON.
func getCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

// postCards adds an card from JSON received in the request body.
func postCards(c *gin.Context) {
	var newCard card

	// Call BindJSON to bind the received JSON to
	// newCard.
	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	// Add the new card to the slice.
	cards = append(cards, newCard)
	c.IndentedJSON(http.StatusCreated, newCard)
}

// getCardByID locates the card whose ID value matches the id
// parameter sent by the client, then returns that card as a response.
func getCardByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of cards, looking for
	// an card whose ID value matches the parameter.
	for _, a := range cards {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "card not found"})
}
