package screens

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"termtodo/value_objects"
)

var titleText = fmt.Sprintf(
	"%s TERMINAL TO DO APP %s",
	value_objects.EmojiAliases["innocent"],
	value_objects.EmojiAliases["innocent"],
)

var descriptionText = fmt.Sprintf(
	"Press N/n to create, select number to edit/delete, vim bindings for navigation",
)

func titleBar() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = titleText
	p.Text = descriptionText
	return p
}

func contentBox() *widgets.List {
	return ToDoList(false)
}
func Home() {
	// Initlize the ui
	if err := ui.Init(); err != nil {
		panic(err)
	}
	// Title bar
	titleBar_ := titleBar()
	titleBar_.PaddingLeft = 0
	fmt.Printf("%v", titleBar_)
	//titleBar_.SetRect(0, 0, 100, 100)
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	// GRID Layout
	//TITLE BOX
	// TITLE BOX
	// TITLE BOX
	contentBox_ := contentBox()
	pieChart_ := BarProgress()

	grid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, titleBar_),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/3, contentBox_),
			ui.NewCol(1.0/3, ToDoList(false)),
			ui.NewCol(1.0/3, ToDoList(true)),
		),
		ui.NewRow(1.0,
			ui.NewCol(1.0/3, pieChart_),
		),
	)
	ui.Render(grid)

	// event loop for the ui
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
	defer ui.Close()
}
