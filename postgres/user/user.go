package User

import (
	"context"
	"database/sql"
	"fmt"
	"rekber/ierr"
	"rekber/internal/user"
	"rekber/postgres/model"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func (u Repository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (user.User, error) {
	var usr model.User
	if err := u.db.GetContext(ctx, &usr, "SELECT * FROM users WHERE phone_number = $1", phoneNumber); err != nil {
		if err == sql.ErrNoRows {
			return user.User{}, ierr.UserNotFound{PhoneNumber: phoneNumber}
		}

		return user.User{}, fmt.Errorf("failed to query from database: %w", err)
	}

	return user.User{
		ID:                    usr.ID,
		PhoneNumber:           usr.PhoneNumber,
		Name:                  usr.Name,
		PhoneNumberVerifiedAt: usr.PhoneNumberVerifiedAt,
		CreatedAt:             usr.CreatedAt,
	}, nil
}

func (u Repository) Save(ctx context.Context, user user.User) error {
	tx := u.db.MustBegin()

	userModel := model.User{
		ID:                    user.ID,
		Name:                  user.Name,
		PhoneNumber:           user.PhoneNumber,
		PhoneNumberVerifiedAt: user.PhoneNumberVerifiedAt,
		CreatedAt:             user.CreatedAt,
	}

	_, err := tx.NamedExecContext(ctx, "INSERT INTO users (id, name, phone_number, phone_number_verified_at, created_at) VALUES (:id, :name, :phone_number, :phone_number_verified_at, :created_at)", userModel)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert user: %w", err)
	}

	tx.Commit()
	return nil
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}
