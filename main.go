package main

import (
	"log"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	main := newPrimitive("Player")
	sideBar := newPrimitive("Music List")

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(-7, -3).
		SetBorders(true)

	grid.AddItem(main, 0, 0, 1, 1, 0, 0, false).
		AddItem(sideBar, 0, 1, 1, 1, 0, 0, false)
	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
