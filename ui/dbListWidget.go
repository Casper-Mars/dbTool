package ui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var MaxWidth = 3

var counter = 0

var rowIndex = 0

type DBListWidget struct {
	layout       *gtk.Grid
	checkBoxList []*gtk.CheckButton
}

func NewDBListWidget() *DBListWidget {
	gridNew, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}
	return &DBListWidget{
		layout:       gridNew,
		checkBoxList: make([]*gtk.CheckButton, 0),
	}
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
	if counter >= MaxWidth {
		counter = 0
		rowIndex++
	}
	widget.layout.Attach(label, counter, rowIndex, 1, 1)
	counter++
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
