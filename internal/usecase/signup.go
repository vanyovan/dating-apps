package usecase

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
	"github.com/vanyovan/dating-apps/internal/entity"
)

type SignUpUsecase interface {
	CreateSignUp(ctx context.Context, param entity.CreateSignUpParam) error
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type Usecase struct {
	db *sql.DB
}

func NewSignUpUsecase(db *sql.DB) SignUpUsecase {
	return &Usecase{
		db: db,
	}
}

func (uc *Usecase) CreateSignUp(ctx context.Context, param entity.CreateSignUpParam) error {
	tx, err := uc.db.Begin()
	_, err = tx.ExecContext(ctx, "INSERT INTO User (username, password, email) VALUES (?, ?, ?)", param.Username, param.Password, param.Email)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (uc *Usecase) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	// rows, err := uc.db.Query("SELECT * FROM User WHERE USERNAME = ?", username)
	// if err != nil {
	// 	return entity.User{}, fmt.Errorf("failed to get user")
	// }
	// defer rows.Close()
	// var users []entity.User
	// for rows.Next() {
	// 	var user entity.User
	// 	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, user)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
	query := "SELECT * FROM User WHERE USERNAME = ?"
	row := uc.db.QueryRow(query, username)
	result := entity.User{}
	err := row.Scan(&result.Username, &result.ID, &result.Email, &result.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Failed to retrieve row:", err)
		}
		return result, err
	}
	return result, nil
}
