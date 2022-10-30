package main

import (
	"fmt"

	"github.com/aspexp/cinema/movie"
	"github.com/aspexp/cinema/ticket"
)

func init(){
	fmt.Println("main.init")
}
func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}