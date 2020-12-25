package ui

import "github.com/gotk3/gotk3/gtk"

type Face interface {
	GetContent() gtk.IWidget
}
