package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	
	"golang_projects/iman-task/service1/repository"
	"github.com/jmoiron/sqlx"
)

const (
	postsTableName          = "posts"
	
)

type Entity struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Body   string `db:"body"`
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreatePosts(ctx context.Context, postsBatch <-chan []repository.Post) <-chan error {
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		err := r.transact(func(tx *sqlx.Tx) error {
			for {
				select {
				case posts, open := <-postsBatch:
					if !open {
						return nil
					}

					if err := createPosts(ctx, tx, posts); err != nil {
						fmt.Errorf("error while creating posts: %v", err)
						return err
					}

				case <-ctx.Done():
					fmt.Errorf("canceled create posts, error: %v", ctx.Err())
					return ctx.Err()
				}
			}
		})

		errChan <- err
	}()

	return errChan
}

func (r *PostgresRepository) SetDownloadStatus(ctx context.Context, success bool, downloadErr error) error {
	return setDownloadStatus(ctx, r.db, success, downloadErr)
}

func (r *PostgresRepository) GetDownloadStatus(ctx context.Context) (success bool, errMsg string, err error) {
	return getDownloadStatus(ctx, r.db)
}

func (r *PostgresRepository) transact(txFunc func(*sqlx.Tx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) 
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit() 
		}
	}()

	return txFunc(tx)
}

func createPosts(ctx context.Context, tx *sqlx.Tx, posts []repository.Post) error {
	for _, p := range posts {
		if err := createPost(ctx, tx, toPostModel(p)); err != nil {
			return err
		}
	}

	return nil
}

func createPost(ctx context.Context, tx *sqlx.Tx, pm Entity) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (id, user_id, title, body)
		VALUES ($1, $2, $3, $4)`, postsTableName,
	)
	_, err := tx.ExecContext(ctx, query, pm.ID, pm.UserID, pm.Title, pm.Body)

	return err
}

func setDownloadStatus(ctx context.Context, db *sqlx.DB, success bool, downloadErr error) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (success, error) VALUES ($1, $2)`,
	)
	var errWrapper sql.NullString
	if downloadErr != nil {
		errWrapper = sql.NullString{
			String: downloadErr.Error(),
			Valid:  true,
		}
	}

	_, err := db.ExecContext(ctx, query, success, errWrapper)

	return err
}

func getDownloadStatus(ctx context.Context, db *sqlx.DB) (success bool, errMsg string, err error) {
	query := fmt.Sprintf(`SELECT success, error FROM %s ORDER BY created_at ASC LIMIT 1`)
	var errWrapper sql.NullString

	err = db.QueryRowxContext(ctx, query).Scan(&success, &errWrapper)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, "", err
		}

		return false, "", err
	}

	return success, errWrapper.String, nil
}

func toPostModel(p repository.Post) Entity {
	return Entity{
		ID:     p.ID,
		UserID: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}
}
