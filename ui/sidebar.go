package sidebar

import (
	"github.com/rivo/tview"
)

func Siderbar(songs []string) tview.Primitive {
	heading := tview.NewTextView().
		SetText("Song's list").
		SetTextAlign(tview.AlignCenter)

	var content tview.Primitive

	if len(songs) == 0 {
		content = noSongsFoundView()
	} else {
		content = songList(songs)
	}

	sb := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(heading, 3, 1, false).
		AddItem(content, 0, 3, false)

	return sb

}

func songList(songs []string) *tview.List {
	var list *tview.List = tview.NewList()

	for _, s := range songs {
		list.AddItem(s, "", 0, nil)
	}

	return list
}

func noSongsFoundView() *tview.TextView {
	var tv *tview.TextView = tview.NewTextView().
		SetText("No songs found").
		SetTextAlign(tview.AlignCenter)

	return tv
}
