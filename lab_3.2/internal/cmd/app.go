package cmd

import (
	cont "github.com/dzhichko/Contains_module/Contains_function"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) SearchSubstr(directory string, substr string) (bool, error) {
	return cont.Contains(directory, substr)
}
