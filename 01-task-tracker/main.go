/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "task-tracker/cmd"

type Task struct {
	Title string		`json:"task"`
	DateAdded string	`json:"date"`
}

func main() {
	cmd.Execute()	
}
