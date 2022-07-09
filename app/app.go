package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// App struct
type App struct {
	ctx           context.Context
	GithubService *GHService
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}
	client := &http.Client{Timeout: 3 * time.Second}
	app.GithubService = NewGHService(client, "https://api.github.com")
	return app

}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
func (a *App) GetGHUser(name string) (*User, error) {
	api, err := a.GithubService.GetUser(name)
	if err != nil {
		return nil, err
	}
	return api, nil
}
