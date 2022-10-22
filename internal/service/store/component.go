package store

import "yunji/utils/sql"

type ComponentService struct {
	db *sql.Database
}

func NewComponentService(db *sql.Database) *ComponentService {
	return &ComponentService{db}
}

func (s *ComponentService) Update() {}

func (s *ComponentService) Get() {}

func (s *ComponentService) Delete() {}

func (s *ComponentService) List() {}

func (s *ComponentService) Create() {}
