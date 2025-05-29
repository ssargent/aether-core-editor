package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"connectrpc.com/connect"
	gamev1 "github.com/ssargent/aether-core-editor/gen/game/v1"
	"github.com/ssargent/aether-core-editor/internal/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AttributeService implements the AttributeService Connect RPC service
type AttributeService struct {
	db *database.DB
}

// NewAttributeService creates a new AttributeService
func NewAttributeService(db *database.DB) *AttributeService {
	return &AttributeService{db: db}
}

// CreateAttribute creates a new attribute
func (s *AttributeService) CreateAttribute(
	ctx context.Context,
	req *connect.Request[gamev1.CreateAttributeRequest],
) (*connect.Response[gamev1.AttributeResponse], error) {
	now := time.Now()

	query := `
		INSERT INTO game.attributes (name, description, created_at, updated_at) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at`

	var id int64
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query,
		req.Msg.Name,
		req.Msg.Description,
		now,
		now,
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create attribute: %w", err))
	}

	attribute := &gamev1.Attribute{
		Id:          id,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Metadata: &gamev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		},
	}

	return connect.NewResponse(&gamev1.AttributeResponse{
		Attribute: attribute,
	}), nil
}

// GetAttribute retrieves an attribute by ID
func (s *AttributeService) GetAttribute(
	ctx context.Context,
	req *connect.Request[gamev1.GetAttributeRequest],
) (*connect.Response[gamev1.AttributeResponse], error) {
	query := `
		SELECT id, name, description, created_at, updated_at 
		FROM game.attributes 
		WHERE id = $1`

	var attribute gamev1.Attribute
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query, req.Msg.Id).Scan(
		&attribute.Id,
		&attribute.Name,
		&attribute.Description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("attribute not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get attribute: %w", err))
	}

	attribute.Metadata = &gamev1.Metadata{
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}

	return connect.NewResponse(&gamev1.AttributeResponse{
		Attribute: &attribute,
	}), nil
}

// UpdateAttribute updates an existing attribute
func (s *AttributeService) UpdateAttribute(
	ctx context.Context,
	req *connect.Request[gamev1.UpdateAttributeRequest],
) (*connect.Response[gamev1.AttributeResponse], error) {
	now := time.Now()

	query := `
		UPDATE game.attributes 
		SET name = $2, description = $3, updated_at = $4 
		WHERE id = $1 
		RETURNING created_at, updated_at`

	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query,
		req.Msg.Id,
		req.Msg.Name,
		req.Msg.Description,
		now,
	).Scan(&createdAt, &updatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("attribute not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update attribute: %w", err))
	}

	attribute := &gamev1.Attribute{
		Id:          req.Msg.Id,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Metadata: &gamev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		},
	}

	return connect.NewResponse(&gamev1.AttributeResponse{
		Attribute: attribute,
	}), nil
}

// DeleteAttribute deletes an attribute and returns the deleted entity
func (s *AttributeService) DeleteAttribute(
	ctx context.Context,
	req *connect.Request[gamev1.DeleteAttributeRequest],
) (*connect.Response[gamev1.DeleteResponse], error) {
	// First get the attribute to return it
	query := `
		SELECT id, name, description, created_at, updated_at 
		FROM game.attributes 
		WHERE id = $1`

	var attribute gamev1.Attribute
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query, req.Msg.Id).Scan(
		&attribute.Id,
		&attribute.Name,
		&attribute.Description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("attribute not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get attribute for deletion: %w", err))
	}

	// Delete the attribute
	deleteQuery := `DELETE FROM game.attributes WHERE id = $1`
	_, err = s.db.ExecContext(ctx, deleteQuery, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to delete attribute: %w", err))
	}

	attribute.Metadata = &gamev1.Metadata{
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}

	return connect.NewResponse(&gamev1.DeleteResponse{
		Success: true,
		Message: fmt.Sprintf("Attribute '%s' deleted successfully", attribute.Name),
	}), nil
}

// ListAttributes lists attributes with optional filtering
func (s *AttributeService) ListAttributes(
	ctx context.Context,
	req *connect.Request[gamev1.ListAttributesRequest],
) (*connect.Response[gamev1.ListAttributesResponse], error) {
	// Start building the query
	query := `
		SELECT id, name, description, created_at, updated_at 
		FROM game.attributes`

	args := []interface{}{}
	conditions := []string{}

	// Add name filter if provided
	if req.Msg.Name != "" {
		conditions = append(conditions, "name ILIKE $"+fmt.Sprintf("%d", len(args)+1))
		args = append(args, "%"+req.Msg.Name+"%")
	}

	// Add WHERE clause if we have conditions
	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			query += " AND " + conditions[i]
		}
	}

	// Add ordering
	query += " ORDER BY name ASC"

	// Execute query
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list attributes: %w", err))
	}
	defer rows.Close()

	var attributes []*gamev1.Attribute
	for rows.Next() {
		var attr gamev1.Attribute
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&attr.Id,
			&attr.Name,
			&attr.Description,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to scan attribute: %w", err))
		}

		attr.Metadata = &gamev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}

		attributes = append(attributes, &attr)
	}

	if err = rows.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error iterating attributes: %w", err))
	}

	// Since the current protobuf doesn't include pagination in request,
	// we'll return simple pagination info
	pagination := &gamev1.PaginationResponse{
		TotalCount:  int32(len(attributes)),
		PageCount:   1,
		CurrentPage: 1,
		PageSize:    int32(len(attributes)),
	}

	return connect.NewResponse(&gamev1.ListAttributesResponse{
		Attributes: attributes,
		Pagination: pagination,
	}), nil
}
