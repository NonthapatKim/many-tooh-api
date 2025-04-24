package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) CheckExists(req domain.CheckExistsRequest) (domain.CheckExistsResponse, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = ?)", req.Table, req.Column)

	err := r.db.QueryRow(query, req.Id).Scan(&exists)
	if err != nil {
		return domain.CheckExistsResponse{}, fmt.Errorf("error checking %s: %w", req.Column, err)
	}

	return domain.CheckExistsResponse{
		Exists: exists,
	}, nil
}
