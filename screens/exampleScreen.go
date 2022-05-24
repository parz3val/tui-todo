package screens

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"termtodo/models"
)

func ToDoList(stats bool) *widgets.List {
	l := widgets.NewList()
	toDoList := []string{
		"Buy milk",
		"Buy bread",
		"Buy cheese",
		"Buy beer",
	}
	var toToDos []models.ToDo
	var counter int
	for _, toDo := range toDoList {
		toToDos = append(toToDos, models.ToDo{
			Title:     toDo,
			ID:        counter,
			Completed: stats,
		})

		counter += 1
	}
	var todoPrettys []string
	for _, todo := range toToDos {
		todo.MakePrettyToDo()
		todoPrettys = append(todoPrettys, todo.PrettyRepresentation)
	}
	l.Rows = todoPrettys
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.BorderStyle = ui.NewStyle(ui.ColorCyan)
	return l
}
