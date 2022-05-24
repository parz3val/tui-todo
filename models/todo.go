package models

import "termtodo/value_objects"

type ToDo struct {
	ID                   int
	Title                string
	Completed            bool
	Description          string
	DueDate              string
	Priority             int
	Tags                 []string
	DateCreated          string
	PrettyRepresentation string
}

// MakePrettyToDo get pretty representation of todolist
func (t *ToDo) MakePrettyToDo() {
	// check if to do is complete
	if t.Completed == true {
		t.PrettyRepresentation = string(rune(t.ID)) + "   " + value_objects.EmojiAliases["checkbox_ticked"] + "   " + t.Title
	} else {
		t.PrettyRepresentation = string(rune(t.ID)) + "   " + value_objects.EmojiAliases["checkbox_unticked"] + "   " + t.Title
	}
}
