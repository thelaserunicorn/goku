package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Save inserts a new resource into the database
func (db *DB) Save(data map[string]interface{}) (int64, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal data: %w", err)
	}

	var id int64
	query := `INSERT INTO resource_table (data) VALUES ($1) RETURNING id`
	err = db.pool.QueryRow(context.Background(), query, jsonData).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to save resource: %w", err)
	}

	return id, nil
}

// GetByID retrieves a resource by its ID
func (db *DB) GetByID(id int64) (*Resource, error) {
	var resource Resource
	var jsonData []byte

	query := `SELECT id, data FROM resource_table WHERE id = $1`
	err := db.pool.QueryRow(context.Background(), query, id).Scan(&resource.ID, &jsonData)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("resource with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get resource: %w", err)
	}

	if err := json.Unmarshal(jsonData, &resource.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &resource, nil
}

// GetAll retrieves all resources from the database
func (db *DB) GetAll() ([]Resource, error) {
	rows, err := db.pool.Query(context.Background(), `SELECT id, data FROM resource_table ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("failed to get resources: %w", err)
	}
	defer rows.Close()

	var resources []Resource
	for rows.Next() {
		var resource Resource
		var jsonData []byte

		if err := rows.Scan(&resource.ID, &jsonData); err != nil {
			return nil, fmt.Errorf("failed to scan resource: %w", err)
		}

		if err := json.Unmarshal(jsonData, &resource.Data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data: %w", err)
		}

		resources = append(resources, resource)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating resources: %w", err)
	}

	return resources, nil
}

// Update modifies an existing resource by ID
func (db *DB) Update(id int64, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	query := `UPDATE resource_table SET data = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
	result, err := db.pool.Exec(context.Background(), query, jsonData, id)
	if err != nil {
		return fmt.Errorf("failed to update resource: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	return nil
}

// Delete removes a resource by ID
func (db *DB) Delete(id int64) error {
	query := `DELETE FROM resource_table WHERE id = $1`
	result, err := db.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	return nil
}
