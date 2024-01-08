package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Couse struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCouser(db *sql.DB) *Couse {
	return &Couse{db: db}
}

func (c *Couse) Create(name, description, categoryID string) (Couse, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses(id, name, description,category_id) VALUES ($1, $2, $3,$4)",
		id, name, description, categoryID)

	if err != nil {
		return Couse{}, err
	}

	return Couse{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Couse) FindAll() ([]Couse, error) {
	rows, err := c.db.Query("SELECT id, name, description,category_id FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	coursers := []Couse{}

	for rows.Next() {
		var id, name, description, category_id string

		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		coursers = append(coursers, Couse{ID: id, Name: name,
			Description: description, CategoryID: category_id})
	}

	return coursers, nil
}

func (c *Couse) FindByCategoryID(categoryID string) ([]Couse, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Couse{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Couse{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}
