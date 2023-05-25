package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
	"github.com/vanyovan/dating-apps/internal/entity"
)

type PackageUsecase interface {
	UpdateUserPremium(ctx context.Context, param entity.UpdateUserPremium) error
	GetPackageByPackageId(ctx context.Context, packageId int) (entity.Package, error)
}

func NewPackageUsecase(db *sql.DB) PackageUsecase {
	return &Usecase{
		db: db,
	}
}

func (uc *Usecase) UpdateUserPremium(ctx context.Context, param entity.UpdateUserPremium) error {
	tx, err := uc.db.Begin()

	//package only valid for 1 month
	dateOneMonth := time.Now().AddDate(0, 1, 0)
	_, err = tx.ExecContext(ctx, "UPDATE USER set is_verified = 1, package_id = ?, package_expire = ? where username = ?", param.PackageId, dateOneMonth, param.Username)
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

func (uc *Usecase) GetPackageByPackageId(ctx context.Context, packageId int) (entity.Package, error) {
	query := "SELECT * FROM Package WHERE package_id = ?"
	row := uc.db.QueryRow(query, packageId)
	result := entity.Package{}
	err := row.Scan(&result.PackageId, &result.Name, &result.Price)
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
