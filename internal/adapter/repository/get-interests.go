package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetInterests() ([]domain.GetInterestsResponse, error) {
	var result []domain.GetInterestsResponse

	query := `
		SELECT
			interest_id,
			image_url AS interest_image_url
			name AS interest_name
		FROM interests
		ORDER BY interest_id;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var interest domain.GetInterestsResponse
		err := rows.Scan(
			&interest.InterestId,
			&interest.InterestImageUrl,
			&interest.InterestName,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, interest)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	return result, nil
}
