package users

import (
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app"
	"github.com/gtarcea/1DevDayTalk2014/schema"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest"
)

type usersResource struct {
	users app.UsersService
}

type userReq struct {
	email    string
	fullname string
}

// NewResource creates a new REST resource for users.
func NewResource(users app.UsersService) *usersResource {
	return &usersResource{
		users: users,
	}
}

// WebService creates a new restful service for users.
func (r *usersResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(rest.RouteHandlder(r.getAllUsers)).
		Doc("Retrieves all users").
		Writes([]schema.User{}))

	ws.Route(ws.GET("{email}").To(rest.RouteHandler(r.getUser)).
		Doc("Retrieves a user by email address").
		Param(ws.PathParameter("email", "email address of existing user").DataType("string")).
		Writes(schema.User{}))

	ws.Route(ws.POST("").To(rest.RouteHandler(r.createUser)).
		Doc("Creates a new user account").
		Reads(userReq{}).
		Writes(schema.User{}))

	return ws
}

func (r *usersResource) getUser(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	email := request.PathParameter("email")

	if email == "" {
		return app.ErrInvalid, nil
	}

	user, err := r.users.GetUserByEmail(email)
	if err != nil {
		return err, nil
	}

	return nil, &user
}

func (r *usersResource) createUser(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	var req userReq
	if err := request.ReadEntity(&req); err != nil {
		return err, nil
	}

	u, err := r.users.CreateUser(req.email, req.fullname)
	return err, u
}

func (r *usersResource) getAllUsers(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	return nil, nil
}
