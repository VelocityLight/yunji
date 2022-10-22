package common

type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UsedByTag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
