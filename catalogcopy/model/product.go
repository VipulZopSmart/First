package model


type Product struct {
	Id    int	`json:"id"`
	Name  string	`json:"name"`
	Brand Brand		`json:"brand"`
}

