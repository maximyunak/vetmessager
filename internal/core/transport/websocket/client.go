package core_websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Message    chan *Event
	ID         int
	Chats      map[int]struct{}
}

type Event struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	ChatID  int      `json:"chat_id"`
	UserID  int      `json:"user_id"`
	Message *Message `json:"message"`
}

type Message struct {
	ID      int     `json:"id"`
	Content *string `json:"content"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Connection.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		c.Connection.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.Connection.Close()
	}()

	for {
		_, message, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var event Event

		if err := json.Unmarshal(message, &event); err != nil {
			log.Printf("error: %v", err)
			continue
		}
		event.Payload.UserID = c.ID
		if _, ok := c.Chats[event.Payload.ChatID]; !ok {
			continue
		}

		hub.BroadcastToChat(event.Payload.ChatID, &event)
	}
}
