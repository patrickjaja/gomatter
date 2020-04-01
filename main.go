package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	//ToDO lastActivity memoryCache
	//Move to addCache/addRow
	//lastActivities :=map[string]string{"user_id": "last_activity"}
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "List"

	//ToDo extract to helper func
	//ToDo move credentials to env vars
	requestBody := []byte(`{"login_id":"nxs_schoenfeld","password":"xxxxxx"}`)
	resp, err := http.Post("https://mattermost.nxs360.com/api/v4/users/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//ToDo use to get last activity
	//body, err :=ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(string(body))

	bearer := ""
	for k, v := range resp.Header {
		if k == "Token" {
			bearer = v[0] //Error Handling -> no session aquired
		}
	}

	//log.Println(bearer)
	l.Rows = []string{
		"Bearer Token:" + bearer,
	}

	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 100, 8)

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
				"Bearer Token:" + bearer,
			}
			ui.Render(l) //redraw ui
		}
	}
}
