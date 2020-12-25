package worker

import (
	"github.com/Casper-Mars/dbTool/ui"
	"github.com/gotk3/gotk3/gtk"
)

type Worker interface {
	GetFace() gtk.IWidget
}

type AbstractWorker struct {
	face ui.Face
}

func (receiver AbstractWorker) GetFace() gtk.IWidget {
	return receiver.face.GetContent()
}
