package repository

import (
	"context"
	domain "go-grpc-exemple/app/domain"
	infrastructure "go-grpc-exemple/app/infrastructure"
)

type UserRepositoryImpl struct {
	Database infrastructure.Database
}

//NewUserRepository ...
func NewUserRepository(db infrastructure.Database) domain.UserRepository {
	return &UserRepositoryImpl{Database: db}
}

func (this *UserRepositoryImpl) AddUserProfile(ctx context.Context, user domain.User) (domain.User, error) {
	sqlStmt := `INSERT INTO users(id,email,first_name,last_name) VALUES($1,$2,$3,$4)`
	err := this.Database.DB.QueryRow(sqlStmt, user.ID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return domain.User{}, infrastructure.ErrNotFound
	}
	return user, nil
}
func (this *UserRepositoryImpl) FetchUsersProfiles(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User
	err = this.Database.DB.Select(&users, `select * from users`)
	if err != nil {
		return nil, infrastructure.ErrNotFound
	}
	return users, nil
}
func (this *UserRepositoryImpl) GetUserProfileByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := this.Database.DB.Get(&user, `select * from users where id = $1`, id)
	if err != nil {
		return domain.User{}, infrastructure.ErrNotFound
	}
	return user, nil
}
func (this *UserRepositoryImpl) UpdateUserProfile(ctx context.Context, user domain.User) (domain.User, error) {
	sqlStmt := `UPDATE users SET email=$1,first_name=$2,last_name=$3 WHERE id=$4`
	err := this.Database.DB.QueryRow(sqlStmt, user.Email, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return domain.User{}, infrastructure.ErrNotFound
	}
	return user, nil
}
func (this *UserRepositoryImpl) DeleteUserProfile(ctx context.Context, id string) error {
	sqlStmt := `DELETE FROM users WHERE id=$1`
	err := this.Database.DB.QueryRow(sqlStmt, id)
	if err != nil {
		return infrastructure.ErrNotFound
	}
	return nil
}
