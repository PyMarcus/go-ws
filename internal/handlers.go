package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Home loads the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "home.jet", nil)
	if err != nil {
		log.Println(fmt.Printf("Error in Home after try renderTemplate: %s", err.Error()))
		return
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)

	if err != nil {
		log.Println(fmt.Printf("Error in WsEndpoint %s", err.Error()))
		return
	}
	var response WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(fmt.Printf(" Error in WsEndpoint %s", err.Error()))
		return
	}
}

// renderTemplate renders a jet template
func renderTemplate(w http.ResponseWriter, name string, data jet.VarMap) error {
	view, err := views.GetTemplate(name)
	if err != nil {
		log.Println(fmt.Printf("Error in renderTemplate %s", err.Error()))
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(fmt.Printf("Error in renderTemplate %s", err.Error()))
		return err
	}

	return nil
}
