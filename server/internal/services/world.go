package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"connectrpc.com/connect"
	aethercorev1 "github.com/ssargent/aether-core-editor/gen/aethercore/v1"
	"github.com/ssargent/aether-core-editor/internal/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// WorldService implements the WorldService Connect RPC service
type WorldService struct {
	db *database.DB
}

// NewWorldService creates a new WorldService
func NewWorldService(db *database.DB) *WorldService {
	return &WorldService{db: db}
}

// CreateWorld creates a new world
func (s *WorldService) CreateWorld(
	ctx context.Context,
	req *connect.Request[aethercorev1.CreateWorldRequest],
) (*connect.Response[aethercorev1.WorldResponse], error) {
	now := time.Now()

	query := `
		INSERT INTO worlds (code, name, description, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, created_at, updated_at`

	var id int64
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query,
		req.Msg.Code,
		req.Msg.Name,
		req.Msg.Description,
		now,
		now,
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create world: %w", err))
	}

	world := &aethercorev1.World{
		Id:          id,
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Metadata: &aethercorev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		},
	}

	return connect.NewResponse(&aethercorev1.WorldResponse{
		World: world,
	}), nil
}

// GetWorld retrieves a world by ID
func (s *WorldService) GetWorld(
	ctx context.Context,
	req *connect.Request[aethercorev1.GetWorldRequest],
) (*connect.Response[aethercorev1.WorldResponse], error) {
	query := `
		SELECT id, code, name, description, created_at, updated_at 
		FROM worlds 
		WHERE id = $1`

	var world aethercorev1.World
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query, req.Msg.Id).Scan(
		&world.Id,
		&world.Code,
		&world.Name,
		&world.Description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("world not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get world: %w", err))
	}

	world.Metadata = &aethercorev1.Metadata{
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}

	return connect.NewResponse(&aethercorev1.WorldResponse{
		World: &world,
	}), nil
}

// UpdateWorld updates an existing world
func (s *WorldService) UpdateWorld(
	ctx context.Context,
	req *connect.Request[aethercorev1.UpdateWorldRequest],
) (*connect.Response[aethercorev1.WorldResponse], error) {
	now := time.Now()

	query := `
		UPDATE worlds 
		SET code = $2, name = $3, description = $4, updated_at = $5 
		WHERE id = $1 
		RETURNING created_at, updated_at`

	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, query,
		req.Msg.Id,
		req.Msg.Code,
		req.Msg.Name,
		req.Msg.Description,
		now,
	).Scan(&createdAt, &updatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("world not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update world: %w", err))
	}

	world := &aethercorev1.World{
		Id:          req.Msg.Id,
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Metadata: &aethercorev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		},
	}

	return connect.NewResponse(&aethercorev1.WorldResponse{
		World: world,
	}), nil
}

// DeleteWorld deletes a world by ID
func (s *WorldService) DeleteWorld(
	ctx context.Context,
	req *connect.Request[aethercorev1.DeleteWorldRequest],
) (*connect.Response[aethercorev1.DeleteResponse], error) {
	// First get the world to return it
	getQuery := `
		SELECT id, code, name, description, created_at, updated_at 
		FROM worlds 
		WHERE id = $1`

	var world aethercorev1.World
	var createdAt, updatedAt time.Time

	err := s.db.QueryRowContext(ctx, getQuery, req.Msg.Id).Scan(
		&world.Id,
		&world.Code,
		&world.Name,
		&world.Description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("world not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get world: %w", err))
	}

	// Now delete it
	deleteQuery := `DELETE FROM worlds WHERE id = $1`
	_, err = s.db.ExecContext(ctx, deleteQuery, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to delete world: %w", err))
	}

	world.Metadata = &aethercorev1.Metadata{
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}

	return connect.NewResponse(&aethercorev1.DeleteResponse{
		Success: true,
		Message: fmt.Sprintf("World '%s' deleted successfully", world.Name),
	}), nil
}

// ListWorlds lists worlds with pagination
func (s *WorldService) ListWorlds(
	ctx context.Context,
	req *connect.Request[aethercorev1.ListWorldsRequest],
) (*connect.Response[aethercorev1.ListWorldsResponse], error) {
	page := int32(1)
	pageSize := int32(50)

	if req.Msg.Pagination != nil {
		if req.Msg.Pagination.Page > 0 {
			page = req.Msg.Pagination.Page
		}
		if req.Msg.Pagination.PageSize > 0 {
			pageSize = req.Msg.Pagination.PageSize
		}
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// Build query with optional filters
	whereClause := ""
	args := []interface{}{}
	argIndex := 1

	if req.Msg.Code != "" {
		whereClause += fmt.Sprintf(" AND code ILIKE $%d", argIndex)
		args = append(args, "%"+req.Msg.Code+"%")
		argIndex++
	}

	if req.Msg.Name != "" {
		whereClause += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+req.Msg.Name+"%")
		argIndex++
	}

	if whereClause != "" {
		whereClause = "WHERE" + whereClause[4:] // Remove leading " AND"
	}

	// Get total count
	var totalCount int32
	countQuery := "SELECT COUNT(*) FROM worlds " + whereClause
	err := s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to count worlds: %w", err))
	}

	// Get worlds
	query := fmt.Sprintf(`
		SELECT id, code, name, description, created_at, updated_at 
		FROM worlds %s
		ORDER BY created_at DESC 
		LIMIT $%d OFFSET $%d`, whereClause, argIndex, argIndex+1)

	args = append(args, pageSize, offset)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list worlds: %w", err))
	}
	defer rows.Close()

	var worlds []*aethercorev1.World
	for rows.Next() {
		var world aethercorev1.World
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&world.Id,
			&world.Code,
			&world.Name,
			&world.Description,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to scan world: %w", err))
		}

		world.Metadata = &aethercorev1.Metadata{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}

		worlds = append(worlds, &world)
	}

	if err := rows.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to iterate worlds: %w", err))
	}

	pageCount := (totalCount + pageSize - 1) / pageSize

	return connect.NewResponse(&aethercorev1.ListWorldsResponse{
		Worlds: worlds,
		Pagination: &aethercorev1.PaginationResponse{
			TotalCount:  totalCount,
			PageCount:   pageCount,
			CurrentPage: page,
			PageSize:    pageSize,
		},
	}), nil
}
