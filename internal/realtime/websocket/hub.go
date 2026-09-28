package realtime_websocket

type Hub struct {
	Clients    map[*Client]struct{}
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Event
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]struct{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Event),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Clients[client] = struct{}{}
		case client := <-h.unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Message)
			}
		case message := <-h.broadcast:
			for client := range h.Clients {
				client.Message <- message
			}
		}
	}
}

func (h *Hub) BroadcastToChat(chatID int, event *Event) {
	for client := range h.Clients {
		if _, ok := client.Chats[chatID]; !ok {
			continue
		}
		client.Message <- event
	}
}
