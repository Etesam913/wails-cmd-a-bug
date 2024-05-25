package main

import "github.com/wailsapp/wails/v3/pkg/application"

type GreetService struct{}

func (g *GreetService) Greet(name string) []string {
	filePaths, _ := application.OpenFileDialog().
		CanChooseFiles(true).
		CanCreateDirectories(true).
		ShowHiddenFiles(true).
		PromptForMultipleSelection()

	return filePaths
}
