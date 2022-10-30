package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type RelationshipRepository struct {
	db *sql.DB
}

func (r *RelationshipRepository) Add(ctx context.Context, idParent, idChildren int) error {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO relationships(parent, children) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("error while prepare insert relationship query. %w", err)
	}
	defer stmt.Close()
	var args []interface{}
	args = append(args, idParent, idChildren)

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("error while exec insert relationship in database: %w", err)
	}

	return nil
}

func (r *RelationshipRepository) GetParents(ctx context.Context, id int) ([]*entity.Person, error) {

	rows, err := r.db.QueryContext(
		ctx,
		"select p.id as id, p.name as name from relationships rel join persons p on p.id = rel.parent where rel.children = ?",
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("error wihile get parents query. %w", err)
	}

	defer rows.Close()

	parents := make([]*entity.Person, 0)
	for rows.Next() {
		var (
			person entity.Person
			idDB   int64
		)

		err := rows.Scan(&idDB, &person.Name)

		if err != nil {
			return nil, fmt.Errorf("can't scan parent: %w", err)
		}

		person.ID = int(idDB)

		parents = append(parents, &person)
	}

	return parents, nil
}

func (r *RelationshipRepository) GetChildren(ctx context.Context, id int) ([]*entity.Person, error) {

	rows, err := r.db.QueryContext(
		ctx,
		"select p.id as id, p.name as name from relationships rel join persons p on p.id = rel.children where rel.parent = ?",
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("error wihile get children query. %w", err)
	}

	defer rows.Close()

	children := make([]*entity.Person, 0)
	for rows.Next() {
		var (
			person entity.Person
			idDB   int64
		)

		err := rows.Scan(&idDB, &person.Name)

		if err != nil {
			return nil, fmt.Errorf("can't scan children: %w", err)
		}

		person.ID = int(idDB)

		children = append(children, &person)
	}

	return children, nil
}

func NewRelationshipRepository(connectionString string) (*RelationshipRepository, error) {
	configureDb, err := ConfigureDb(connectionString)

	if err != nil {
		return nil, err
	}

	return &RelationshipRepository{db: configureDb}, nil
}
