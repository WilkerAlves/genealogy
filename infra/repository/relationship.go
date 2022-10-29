package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func NewRelationshipRepository(connectionString string) (*RelationshipRepository, error) {
	configureDb, err := ConfigureDb(connectionString)

	if err != nil {
		return nil, err
	}

	return &RelationshipRepository{db: configureDb}, nil
}
