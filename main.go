package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "List"

	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
				case "q", "<C-c>":
					return
				case "j", "<Down>":
					l.ScrollDown()
				case "k", "<Up>":
					l.ScrollUp()
			}
		case <-ticker:

			l.Rows = []string{
				time.Now().String(),
				time.Now().String(),
			}
			ui.Render(l)
		}
	}

	//previousKey := ""
	//uiEvents := ui.PollEvents()
	//
	//for {
	//	e := <-uiEvents
	//	switch e.ID {
	//	case "q", "<C-c>":
	//		return
	//	case "j", "<Down>":
	//		l.ScrollDown()
	//	case "k", "<Up>":
	//		l.ScrollUp()
	//	case "<C-d>":
	//		l.ScrollHalfPageDown()
	//	case "<C-u>":
	//		l.ScrollHalfPageUp()
	//	case "<C-f>":
	//		l.ScrollPageDown()
	//	case "<C-b>":
	//		l.ScrollPageUp()
	//	case "g":
	//		if previousKey == "g" {
	//			l.ScrollTop()
	//		}
	//	case "<Home>":
	//		l.ScrollTop()
	//	case "G", "<End>":
	//		l.ScrollBottom()
	//	}
	//
	//	if previousKey == "g" {
	//		previousKey = ""
	//	} else {
	//		previousKey = e.ID
	//	}
	//
	//	ui.Render(l)
	//}
}