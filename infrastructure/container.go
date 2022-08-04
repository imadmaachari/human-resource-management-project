package infrastructure

import (
	domain "user-management-project/app/domain"
	repositories "user-management-project/app/repository"
	services "user-management-project/app/usecase"
	"user-management-project/infrastructure/utils"
)

//Declare services/repos/components here
type Container struct {
	Services Services
}

type Services struct {
	UserService domain.UserUsecase
}

//Inject represent the starter of our IoC container
//Here , all the necessary structs/functions the we
//need to buid the project will be injected
func Inject() Container {
	//Init Databases
	bdbs := utils.NewDatabase()
	//Init Repositories
	ur := repositories.NewUserRepository(bdbs)
	//Init Services
	us := services.NewUserUseCase(ur)

	services := Services{
		UserService: us,
	}

	return Container{
		Services: services,
	}
}
