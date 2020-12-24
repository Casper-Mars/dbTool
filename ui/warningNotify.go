package ui

import "github.com/gotk3/gotk3/gtk"

func ShowWarning(window *gtk.Window, info string) {
	dialogNew := gtk.MessageDialogNew(
		window,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_WARNING,
		gtk.BUTTONS_OK,
		info,
	)
	dialogNew.Run()
	dialogNew.Destroy()
}
