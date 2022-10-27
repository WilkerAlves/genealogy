package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type PersonRepository struct {
	DB *sql.DB
}

func (r *PersonRepository) Add(ctx context.Context, name string) (*entity.Person, error) {
	stmt, err := r.DB.PrepareContext(ctx, "INSERT INTO person(name) VALUES (?)")
	if err != nil {
		return nil, errors.New("error while prepare insert person query")
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

	row := r.DB.QueryRowContext(ctx, "select id, name from person where id = ?", id)

	err := row.Scan(
		&person.Name,
		&idDB,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("person %d not found", id)
		}
		return nil, fmt.Errorf("can't scan person '%d' result row of first class: %w", id, err)
	}

	person.ID = int(idDB)

	return &person, nil
}

func (r *PersonRepository) Update(ctx context.Context, id int, name string) error {
	stmt, err := r.DB.PrepareContext(ctx, "UPDATE person SET name = ? WHERE id = ?")
	if err != nil {
		return errors.New("error while prepare update person query")
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
	stmt, err := r.DB.PrepareContext(ctx, "DELETE FROM person WHERE id = ?")
	if err != nil {
		return errors.New("error while prepare delete person query")
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

	return &PersonRepository{DB: configureDb}, nil
}
