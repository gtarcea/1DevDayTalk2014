package ws

import (
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app/services"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest/filters"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest/users"
)

func NewRegisteredServicesContainer(privateKey, publicKey []byte) *restful.Container {
	f := filters.NewJWTFilter(publicKey)
	container := restful.NewContainer()
	container.Filter(f.Filter)
	container.Add(usersWebService(privateKey, publicKey))
	return container
}

func usersWebService(privateKey, publicKey []byte) *restful.WebService {
	usersService := services.NewUsers(services.NewUsersDB())
	r := users.NewResource(usersService, privateKey)
	return r.WebService()
}
