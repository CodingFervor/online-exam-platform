package service

import (
	"database/sql"
	"github.com/CodingFervor/online-exam-platform/internal/database"
	"github.com/CodingFervor/online-exam-platform/internal/repository"
)

type Context struct {
	repo *repository.Repository
}

func NewContext() *Context {
	return &Context{repo: repository.New(database.DB)}
}

func (c *Context) Repo() *repository.Repository { return c.repo }

func (c *Context) HealthCheck() map[string]string {
	status := map[string]string{}
	if database.DB != nil {
		if err := database.DB.Ping(); err != nil {
			status["database"] = "unhealthy: " + err.Error()
		} else {
			status["database"] = "healthy"
		}
	} else {
		status["database"] = "not connected"
	}
	return status
}
