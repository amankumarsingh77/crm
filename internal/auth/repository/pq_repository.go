package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/amankumarsingh77/cmr/pkg/utils"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) auth.Repository {
	return &authRepo{
		db: db,
	}
}

func (a *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}
	err := a.db.QueryRowxContext(
		ctx,
		createUser,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Phone,
		&user.IsActive,
	).StructScan(u)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return u, nil
}

func (a *authRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := a.db.GetContext(
		ctx,
		u,
		updateUser,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Role,
		&user.Phone,
		&user.IsActive,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to update user : %v", err)
	}
	return u, nil
}

func (a *authRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	result, err := a.db.ExecContext(
		ctx,
		deleteUserQuery,
		userID,
	)
	if err != nil {
		return fmt.Errorf("failed to delete user %v : ", err)
	}
	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rowsaffected %v", err)
	}
	if rowsEffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (a *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	u := &models.User{}
	if err := a.db.QueryRowxContext(
		ctx,
		getUserQuery,
		userID,
	).StructScan(u); err != nil {
		return nil, fmt.Errorf("failed to get user : %v", err)
	}
	return u, nil
}

// not sure if I must implement this
func (a *authRepo) FindByName(ctx context.Context, name string, query *utils.Pagination) (*models.UsersList, error) {
	return nil, nil
}

func (a *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}

	if err := a.db.QueryRowxContext(
		ctx,
		getUserByEmail,
		&user.Email,
	).StructScan(u); err != nil {
		return nil, fmt.Errorf("failed to get user :%v", err)
	}
	return u, nil
}

func (a *authRepo) GetUsers(ctx context.Context, pq *utils.Pagination) (*models.UsersList, error) {
	var totalCount int
	if err := a.db.GetContext(
		ctx,
		&totalCount,
		getTotal,
	); err != nil {
		return nil, fmt.Errorf("faile to get users : %v", err)
	}
	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetPage()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			Users:      make([]*models.User, 0),
		}, nil
	}
	var users = make([]*models.User, 0, pq.GetSize())
	if err := a.db.SelectContext(
		ctx,
		&users,
		getUsers,
		pq.GetOrderBy(),
		pq.GetOffset(),
		pq.GetLimit(),
	); err != nil {
		return nil, fmt.Errorf("failed to get users %v ", err)
	}
	return &models.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetPage()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		Users:      users,
	}, nil
}
