package store

import "yunji/utils/sql"

type TagService struct {
	db *sql.Database
}

func NewTagService(db *sql.Database) *TagService {
	return &TagService{db}
}

func (s *TagService) Update() {}

func (s *TagService) Get() {}

func (s *TagService) Delete() {}

func (s *TagService) List() {}

func (s *TagService) Create() {}
