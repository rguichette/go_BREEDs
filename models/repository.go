package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
}

type mysqlRepo struct {
	DB *sql.DB
}

func newMySqlRepo(conn *sql.DB) Repository {
	return &mysqlRepo{
		DB: conn,
	}
}

type testRepo struct {
	DB *sql.DB
}

func newTestRepo(conn *sql.DB) Repository {
	return &testRepo{
		DB: nil,
	}
}
