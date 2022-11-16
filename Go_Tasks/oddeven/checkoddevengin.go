package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Check(ch chan int, a *gin.Context) {
	val := <-ch
	switch {
	case val%2 == 0:
		fmt.Println(" Even number")
		a.JSON(200, gin.H{
			"response": " Even number",
		})
	case val%2 != 0:
		fmt.Println(" Odd number")
		a.JSON(200, gin.H{
			"response": " Odd number",
		})
	default:
		fmt.Println("")
	}
}
func getParam(ch *gin.Context) {
	channel := make(chan int)
	go Check(channel, ch)
	id := ch.Param("value")
	val, _ := strconv.Atoi(id)
	channel <- val
	// select {
	// case channel <- val:
	// default:
	// }
	// fmt.Println(<-channel)
	defer close(channel)
}
func main() {
	r := gin.Default()
	r.GET("/:value", getParam)
	r.Run(":8080")
}
