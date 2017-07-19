package main

import "github.com/gin-gonic/gin"
import "net/http"

type Business struct {
	Id int `json:"name"`
	City string `json:"city"`
	Name string `json:"name"`
}

var businesses []Business

func List(c *gin.Context) {
	city := c.Param("city")
	c.JSON(http.StatusOK, findBusinessBy(city))
}

func Save(c *gin.Context) {
	var business Business
	c.BindJSON(&business)
	save(business)
	//TODO: add resource id
	c.JSON(http.StatusCreated, gin.H{"location":"/businesses/"})
}

//TODO: implement db save
func save(business Business) {
	business.Id = len(businesses)
	businesses[business.Id] = business
}

//TODO: implement db search
func findBusinessBy(city string) []Business {
	found := []Business{}
	i := -1
	for _, business := range businesses {
		if business.City == city {
			found[i+1] = business
		}
	}
	return found
}
