package media

import "strings"

type Movie struct {
	Title     string
	Rating    Rating
	BoxOffice float32
}

type Rating string

const (
	R = "R (Restricted)"
	G = "G (General)"
)

type Movie2 struct {
	title     string
	rating    Rating
	boxOffice float32
}

func NewMovie(title string, rating Rating, boxOffice float32) Movie2 {
	return Movie2{
		title:     title,
		rating:    rating,
		boxOffice: boxOffice,
	}
}

func (m *Movie2) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie2) SetTitle(newTitle string) {
	m.title = newTitle
}

func (m *Movie2) GetRating() Rating {
	return m.rating
}
func (m *Movie2) SetRating(newRating Rating) {
	m.rating = newRating
}

func (m *Movie2) GetBoxOffice() float32 {
	return m.boxOffice
}
func (m *Movie2) SetBoxOffice(newBoxOffice float32) {
	m.boxOffice = newBoxOffice
}

type Catalogable interface {
	NewMovieInterface(title string, rating Rating, boxOffice float32)
	GetTitleInterface() string
	GetRatingInterface() Rating
	GetBoxOfficeInterface() float32
	SetTitleInterface(newTitle string)
	SetRatingInterface(newRating Rating)
	SetBoxOfficeInterface(newBoxOffice float32)
}

type Movie3 struct {
	title     string
	rating    Rating
	boxOffice float32
}

func (m *Movie3) GetTitleInterface() string {
	return strings.ToTitle(m.title)
}

func (m *Movie3) SetTitleInterface(newTitle string) {
	m.title = newTitle
}

func (m *Movie3) GetRatingInterface() Rating {
	return m.rating
}
func (m *Movie3) SetRatingInterface(newRating Rating) {
	m.rating = newRating
}

func (m *Movie3) GetBoxOfficeInterface() float32 {
	return m.boxOffice
}
func (m *Movie3) SetBoxOfficeInterface(newBoxOffice float32) {
	m.boxOffice = newBoxOffice
}

func (m *Movie3) NewMovieInterface(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}
