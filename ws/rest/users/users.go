package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app"
	"github.com/gtarcea/1DevDayTalk2014/schema"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest"
)

type usersResource struct {
	users      app.UsersService
	privateKey []byte
}

type userReq struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewResource creates a new REST resource for users.
func NewResource(users app.UsersService, privateKey []byte) *usersResource {
	return &usersResource{
		users:      users,
		privateKey: privateKey,
	}
}

// WebService creates a new restful service for users.
func (r *usersResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(rest.RouteHandler(r.getAllUsers)).
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

	ws.Route(ws.POST("/login").To(rest.RouteHandler(r.login)).
		Doc("User login").
		Reads(loginReq{}).
		Writes(schema.Auth{}))

	return ws
}

func (r *usersResource) getAllUsers(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	users, err := r.users.GetAll()
	return err, users
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
	u, err := r.users.CreateUser(req.Email, req.Fullname)
	return err, u
}

func (r *usersResource) login(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	var req loginReq
	if err := request.ReadEntity(&req); err != nil {
		return err, nil
	}

	if err := r.authenticate(req); err != nil {
		return err, nil
	}

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["ID"] = req.Username
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenStr, err := token.SignedString(r.privateKey)
	if err != nil {
		return err, nil
	}

	auth := schema.Auth{
		Username: req.Username,
		Token:    tokenStr,
	}

	return nil, &auth
}

func (r *usersResource) authenticate(req loginReq) error {
	if req.Username != "admin" || req.Password != "abc123" {
		return app.ErrNoAccess
	}
	return nil
}
