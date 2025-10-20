package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Message struct {
	RoomID string
	Data   []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients to all.
	broadcast chan []byte

	// Inbound messages from the clients to room.
	broadcastToRoom chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Registered clients by room
	rooms map[string]map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		broadcast:       make(chan []byte),
		broadcastToRoom: make(chan *Message),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		clients:         make(map[*Client]bool),
		rooms:           make(map[string]map[*Client]bool),
	}
}

func (h *Hub) printAllRoomAndClient() {
	fmt.Println("----ALL CLIENT----")
	for c, _ := range h.clients {
		fmt.Println(c)
	}

	fmt.Println("----ALL ROOM----")
	for roomID, clients := range h.rooms {
		fmt.Println("---", roomID, "---")
		for c, _ := range clients {
			fmt.Println(c)
		}
	}
}

func (h *Hub) BroadcastToAll(data interface{}) {
	jsonByte, _ := json.Marshal(data)
	message := bytes.TrimSpace(bytes.Replace(jsonByte, newline, space, -1))
	h.broadcast <- message
}

func (h *Hub) BroadcastToRoom(room string, data interface{}) {
	jsonByte, _ := json.Marshal(data)
	message := bytes.TrimSpace(bytes.Replace(jsonByte, newline, space, -1))
	h.broadcastToRoom <- &Message{RoomID: room, Data: message}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			//check room
			room := h.rooms[client.roomID]
			if room == nil {
				//the first user in the room
				room = make(map[*Client]bool)
				h.rooms[client.roomID] = room
			}
			room[client] = true
			h.printAllRoomAndClient()
		case client := <-h.unregister:
			//delete from room
			room := h.rooms[client.roomID]
			if room != nil {
				delete(room, client)
				if len(room) == 0 {
					//delete the room
					delete(h.rooms, client.roomID)
				}
			}
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.broadcastToRoom:
			room := h.rooms[message.RoomID]
			if room == nil {
				for client := range room {
					select {
					case client.send <- message.Data:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}

		}

	}
}
