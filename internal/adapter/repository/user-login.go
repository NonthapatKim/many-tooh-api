package repository

import (
	"database/sql"
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func (r *repository) UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error) {
	var result domain.UserLoginResult

	err := r.db.QueryRow(
		`SELECT 
			user_id, 
			password
		FROM users
		WHERE email = ?`,
		req.Email,
	).Scan(
		&result.UserId,
		&result.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("user not found")
		}
		return result, err
	}

	if bcrypt.CompareHashAndPassword([]byte(*result.Password), []byte(*req.Password)) != nil {
		return result, errors.New("incorrect password")
	}

	result.Password = nil

	result = domain.UserLoginResult{
		UserId: result.UserId,
	}

	return result, nil
}
