package main

import (
	handler "user-management-project/app/delivery/http"
	container "user-management-project/infrastructure"
)

func main() {
	di := container.Inject()
	//delegates to start our Http Server
	handler.StartServer(di)
}
