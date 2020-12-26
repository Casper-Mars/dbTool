package ui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type DBListWidget struct {
	layout       *gtk.Grid
	checkBoxList []*gtk.CheckButton
	maxWidth     int
	counter      int
	rowIndex     int
}

func NewDBListWidget() *DBListWidget {
	gridNew, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}
	widget := DBListWidget{
		layout: gridNew,
	}
	widget.init()
	return &widget
}

func (widget *DBListWidget) resetDBList() {
	widget.init()
}

func (widget *DBListWidget) AddDB(dbName string) error {
	if len(dbName) == 0 {
		return fmt.Errorf("数据库名称不能为空")
	}
	label, err := gtk.CheckButtonNewWithLabel(dbName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if widget.counter >= widget.maxWidth {
		widget.counter = 0
		widget.rowIndex++
	}
	widget.layout.Attach(label, widget.counter, widget.rowIndex, 1, 1)
	widget.counter++
	label.Show()
	widget.checkBoxList = append(widget.checkBoxList, label)
	return nil
}

func (widget DBListWidget) GetDBList() []string {
	return widget.getDBList(false)
}

func (widget DBListWidget) getSelectedDBList() []string {
	return widget.getDBList(true)
}

func (widget DBListWidget) getDBList(selected bool) []string {
	list := widget.checkBoxList
	result := make([]string, 0)
	for _, k := range list {
		label, err := k.GetLabel()
		if err == nil {
			if selected && k.GetActive() {
				result = append(result, label)
			}
		} else {
			log.Fatal(err)
		}
	}
	return result
}

func (widget *DBListWidget) init() {
	for _, k := range widget.checkBoxList {
		widget.layout.Remove(k)
	}
	widget.maxWidth = 3
	widget.counter = 0
	widget.rowIndex = 0
	widget.checkBoxList = make([]*gtk.CheckButton, 0)
}
