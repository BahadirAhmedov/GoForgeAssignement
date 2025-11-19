package postgres

import (
	"fmt"
	"sort"

	"goforge/internal/domain/models"
)


func (s *Storage) SaveNumber(num models.Number) ([]int, error) {
	const op = "storage.postgres.SaveNumber"

    _, err := s.db.Exec("INSERT INTO numbers(value) VALUES($1)", num.Value)
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }

	rows, err := s.db.Query("SELECT value FROM numbers")
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }
    defer rows.Close()

    var allNumbers []int
    for rows.Next() {
        var v int
        if err := rows.Scan(&v); err != nil {
            return nil, fmt.Errorf("%s: %w", op, err)
        }
        allNumbers = append(allNumbers, v)
    }

    sort.Ints(allNumbers)

	return allNumbers, nil
}
