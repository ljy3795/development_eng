// https://stackoverflow.com/questions/41483970/how-to-return-html-on-gin

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


func homePage(c *gin.Context){
	c.JSON(200, gin.H{
		"message" : "Hello",
	})
}

func postHomePage(c *gin.Context){
	// c.JSON(200, gin.H{
	// 	"message": "POST HOMEPAGE",
	// })
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name") // name == elliot
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name" : name,
		"age" : age,
	})
}

func PathParams(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name" : name,
		"age" : age,
	})
}

func main() {
	// fmt.Println("hello world")

	r := gin.Default()


	// 아래가 각각 end-point
	r.GET("/", homePage)
	r.POST("/", postHomePage)

	// http://localhost:8080/query/?name=elliot&age=25
	r.GET("/query", QueryStrings) // /query?name=elliot&age=24
	r.GET("/path/:name/:age",PathParams) // /path/elliot/24/

	// r.GET("/", func(c *gin.Context){
	// 	c.JSON(200, gin.H{
	// 		"message" : "Hello_Work",
	// 	})
	// })
	r.Run()
}