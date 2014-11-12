package ws

import (
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app/services"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest/filters"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest/users"
)

// NewRegisteredServicesContainer creates a new container of web services. It
// also connects up a JWT Filter to the container to authenticate the user.
func NewRegisteredServicesContainer(privateKey, publicKey []byte) *restful.Container {
	f := filters.NewJWTFilter(publicKey, "/users/login")
	container := restful.NewContainer()
	container.Filter(f.Filter)
	container.Add(usersWebService(privateKey, publicKey))
	return container
}

// usersWebService creates a new users REST end point.
func usersWebService(privateKey, publicKey []byte) *restful.WebService {
	usersService := services.NewUsers(services.NewUsersDB())
	r := users.NewResource(usersService, privateKey)
	return r.WebService()
}
