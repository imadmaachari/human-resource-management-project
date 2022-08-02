package infrastructure

import (
	booksrepo "go-starling-middleware/app/repository/pg/book"
	booksservice "go-starling-middleware/app/usecase/book"
	domain "go-starling-middleware/domain/book"
	"go-starling-middleware/infrastructure/utils"
)

//Declare services/repos/components here
type Container struct {
	Services Services
}

type Services struct {
	BookService domain.BookUsecase
}

//Inject represent the starter of our IoC container
//Here , all the necessary structs/functions the we
//need to buid the project will be injected
func Inject() Container {
	//Init Databases
	bdbs := utils.NewDatabase()
	//Init Repositories
	br := booksrepo.NewBookRepository(bdbs)
	//Init Services
	bs := booksservice.NewBookUsecase(br)

	services := Services{
		BookService: bs,
	}

	return Container{
		Services: services,
	}
}
