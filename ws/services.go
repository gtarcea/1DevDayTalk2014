package ws

import (
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app/services"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest/users"
)

func NewRegisteredServicesContainer() *restful.Container {
	container := restful.NewContainer()
	container.Add(usersWebService())
	return container
}

func usersWebService() *restful.WebService {
	usersService := services.NewUsers(services.NewUsersDB())
	r := users.NewResource(usersService)
	return r.WebService()
}
