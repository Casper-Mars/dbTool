package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type FormInputWidget struct {
	row           *gtk.Grid
	inputTextArea *gtk.Entry
	labelName     string
}

func NewFormInputWidget(labelName string) FormInputWidget {
	row, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}
	labelNew, err := gtk.LabelNew(labelName)
	if err != nil {
		log.Fatal(err)
	}
	row.Attach(labelNew, 0, 0, 1, 1)
	entryNew, err := gtk.EntryNew()
	if err != nil {
		log.Fatal(err)
	}
	row.Attach(entryNew, 1, 0, 2, 1)
	return FormInputWidget{
		labelName:     labelName,
		row:           row,
		inputTextArea: entryNew,
	}
}

func (widget FormInputWidget) GetRow() *gtk.Grid {
	return widget.row
}

func (widget FormInputWidget) GetInputText() string {
	text, err := widget.inputTextArea.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (widget FormInputWidget) GetInputEntry() *gtk.Entry {
	return widget.inputTextArea
}
