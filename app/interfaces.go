package app

import "github.com/gtarcea/1DevDayTalk2014/schema"

type ProjectsService interface {
	CreateProject(schema.Project) (schema.Project, error)
	AddUser(projectID string, userID string) (schema.Project, error)
	ListProjects() ([]schema.Project, error)
}

type UsersService interface {
	CreateUser(schema.User) (schema.User, error)
	GetUserByEmail(email string) (schema.User, error)
}
