package projects

import (
	"github.com/emicklei/go-restful"
	"github.com/gtarcea/1DevDayTalk2014/app"
	"github.com/gtarcea/1DevDayTalk2014/schema"
	"github.com/gtarcea/1DevDayTalk2014/ws/rest"
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
		Writes([]schema.Project{}))

	ws.Route(ws.POST("").To(rest.RouteHandler(r.createProject)).
		Doc("Creates a new project").
		Reads(schema.Project{}).
		Writes(schema.Project{}))

	ws.Route(ws.POST("/{project-id}/user/{user-id}").
		To(rest.RouteHandler(r.addUser)).
		Doc("Adds user to project").
		Param(ws.PathParameter("project-id", "ID of project").DataType("string")).
		Param(ws.PathParameter("user-id", "ID of user").DataType("string")).
		Reads(schema.User{}).
		Writes(schema.Project{}))
	return ws
}

func (r *projectsResource) allProjects(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	projects, err := r.projects.ListProjects()

	if err != nil {
		return err, nil
	}

	return nil, projects
}

func (r *projectsResource) createProject(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	var project schema.Project

	if err := request.ReadEntity(&project); err != nil {
		return err, nil
	}

	p, err := r.projects.CreateProject(project)
	if err != nil && err != app.ErrExists {
		return err, nil
	}

	return nil, &p
}

func (r *projectsResource) addUser(request *restful.Request, response *restful.Response, user schema.User) (error, interface{}) {
	userID := request.PathParameter("user-id")
	projectID := request.PathParameter("project-id")

	if userID == "" || projectID == "" {
		// required ids not sent
		return app.ErrInvalid, nil
	}

	project, err := r.projects.AddUser(projectID, userID)
	if err != nil && err != app.ErrExists {
		return err, nil
	}

	return nil, &project
}
