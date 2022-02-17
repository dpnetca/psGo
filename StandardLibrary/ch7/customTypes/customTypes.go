package main

import (
	"fmt"

	"github.com/dpnetca/psGo/StandardLibrary/ch7/customTypes/media"
)

func main() {
	fmt.Println("My Favorite Movie")

	myFav := media.Movie{}
	myFav.Title = "Top Gun"
	myFav.Rating = media.R
	myFav.BoxOffice = 43.2

	fmt.Printf("My favorite movie is %s!!\n", myFav.Title)
	fmt.Printf("it was rated %v\n", myFav.Rating)
	fmt.Printf("it made %.2f in the Box Office\n", myFav.BoxOffice)
	myFav.Title = "Shawshank Redemption"
	fmt.Printf("%v\n", myFav)

	myFav2 := media.NewMovie("Top Gun", media.R, 43.2)
	fmt.Printf("%v\n", myFav2)

	fmt.Printf("My favorite movie is %s!!\n", myFav2.GetTitle())
	fmt.Printf("it was rated %v\n", myFav2.GetRating())
	fmt.Printf("it made %.2f in the Box Office\n", myFav2.GetBoxOffice())

	myFav2.SetTitle("Shawshank Redemption")
	fmt.Printf("%v\n", myFav2)

	var myFav3 media.Catalogable = &media.Movie3{}
	myFav3.NewMovieInterface("Top Gun", media.R, 43.2)
	fmt.Printf("%v\n", myFav3.GetTitleInterface())

}
