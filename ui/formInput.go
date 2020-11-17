package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type FormInputWidget struct {
	row           *gtk.Box
	inputTextArea *gtk.Entry
	labelName     string
}

func NewFormInputWidget(labelName string) FormInputWidget {

	boxNew, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal(err)
	}
	labelNew, err := gtk.LabelNew(labelName)
	if err != nil {
		log.Fatal(err)
	}
	boxNew.Add(labelNew)
	entryNew, err := gtk.EntryNew()
	if err != nil {
		log.Fatal(err)
	}
	boxNew.Add(entryNew)
	return FormInputWidget{
		labelName:     labelName,
		row:           boxNew,
		inputTextArea: entryNew,
	}
}

func (widget FormInputWidget) GetRow() *gtk.Box {
	return widget.row
}

func (widget FormInputWidget) GetInputText() string {
	text, err := widget.inputTextArea.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}
