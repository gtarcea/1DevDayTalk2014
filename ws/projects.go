package projects

import (
	"github.com/emicklei/go-restful"
	"github.com/materials-commons/mcfs/common/schema"
	"github.com/materials-commons/mcfs/mcfsd/app"
	"github.com/materials-commons/mcfs/mcfsd/interfaces/ws/rest"
)

type projectsResource struct {
	projects app.ProjectsService
}

type User struct {
	ID       string `json:"id"`
	EMail    string `json:"email"`
	Fullname string `json:"fullname"`
}

type Project struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

var projects = []Project{
	Project{
		ID:   "1",
		Name: "Project 1",
		Users: []User{
			User{
				ID:       "1",
				EMail:    "bob@somewhere.com",
				Fullname: "Bob Biff",
			},
		},
	},
}

// NewResource creates a new REST resource for projects.
func NewResource(projects app.ProjectsService) *projectsResource {
	return &projectsResource{
		projects: projects,
	}
}

// WebService creates a new restful service for projects.
func (r *projectsResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/projects").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(rest.RouteHandler(r.allProjects)).
		Doc("List all projects").
		Writes([]Project{}))

	ws.Route(ws.POST("").To(rest.RouteHandler(r.createProject)).
		Doc("Creates a new project").
		Reads(Project{}).
		Writes(Project{}))

	ws.Route(ws.POST("/{project-id}/user/{user-id}").To(r.addUser).
		Doc("Adds user to project").
		Param(ws.PathParameter("project-id", "ID of project").DataType("string")).
		Param(ws.PathParameter("user-id", "ID of user").DataType("string")).
		Reads(User{}).
		Writes(Project{}))
	return ws
}

func (r *projectsResource) allProjects(request *restful.Request, response *restful.Response, user schema.User) {

}

func (s *projectsResource) addProject(request *restful.Request, response *restful.Response, user schema.User) {

}

func (s *projectsResource) addUser(request *restful.Request, response *restful.Response, user schema.User) {

}
