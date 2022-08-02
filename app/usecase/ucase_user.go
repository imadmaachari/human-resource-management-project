package usecase

import (
	"context"
	domain "go-grpc-exemple/app/domain"

	"github.com/satori/uuid"
)

type UserUseCaseImpl struct {
	userRepository domain.UserRepository
}

//NewUserRepository ...
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUseCaseImpl{userRepository: userRepo}
}
func (this *UserUseCaseImpl) AddUserProfile(ctx context.Context, user domain.User) (domain.User, error) {
	id := uuid.NewV4()
	user.ID = id.String()
	user, err := this.userRepository.AddUserProfile(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (this *UserUseCaseImpl) FetchUsersProfiles(ctx context.Context) (res []domain.User, err error) {
	users, err := this.userRepository.FetchUsersProfiles(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (this *UserUseCaseImpl) GetUserProfileByID(ctx context.Context, id string) (domain.User, error) {
	user, err := this.userRepository.GetUserProfileByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (this *UserUseCaseImpl) UpdateUserProfile(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := this.userRepository.UpdateUserProfile(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (this *UserUseCaseImpl) DeleteUserProfile(ctx context.Context, id string) error {
	err := this.userRepository.DeleteUserProfile(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
