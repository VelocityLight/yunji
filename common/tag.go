package common

type Tag struct {
	Name string `json:"name" db:"resource_tags_user_component"`
}

type UsedByTag struct {
	Name string `json:"name" db:"resource_tags_user_usedby"`
}
