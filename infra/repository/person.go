package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type PersonRepository struct {
	db *sql.DB
}

func (r *PersonRepository) Add(ctx context.Context, name string) (*entity.Person, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO persons(name) VALUES (?)")
	if err != nil {
		return nil, fmt.Errorf("error while prepare insert person query. %w", err)
	}
	defer stmt.Close()
	var args []interface{}
	args = append(args, name)

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("error while exec insert person in database: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error for get lastId: %w", err)
	}

	return &entity.Person{
		ID:   int(id),
		Name: name,
	}, nil
}

func (r *PersonRepository) Get(ctx context.Context, id int) (*entity.Person, error) {
	var (
		person entity.Person
		idDB   int64
	)

	row := r.db.QueryRowContext(ctx, "select id, name from persons where id = ?", id)

	err := row.Scan(
		&idDB,
		&person.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("person %d not found", id)
		}
		return nil, fmt.Errorf("can't scan person '%d': %w", id, err)
	}

	person.ID = int(idDB)

	return &person, nil
}

func (r *PersonRepository) Update(ctx context.Context, id int, name string) error {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE persons SET name = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("error while prepare update person query: %w", err)
	}
	defer stmt.Close()
	var args []interface{}
	args = append(args, name, id)

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("error while exec update person in database: %w", err)
	}

	return nil
}

func (r *PersonRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM persons WHERE id = ?")
	if err != nil {
		return fmt.Errorf("error while prepare delete person query: %w", err)
	}
	defer stmt.Close()
	var args []interface{}
	args = append(args, id)

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("error while exec delete person in database: %w", err)
	}

	return nil
}

func NewPersonRepository(connectionString string) (*PersonRepository, error) {
	configureDb, err := ConfigureDb(connectionString)

	if err != nil {
		return nil, err
	}

	return &PersonRepository{db: configureDb}, nil
}
