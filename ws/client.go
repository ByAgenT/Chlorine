package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

// Client is a middleman structure between the websocket connection and the websocket hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// Channel of incoming messages.
	receive chan []byte

	// Indicator of a connection that is dead and should not be used.
	dead bool
}

func (c *Client) SendMessage(message []byte) {
	c.send <- message
}

func (c *Client) Deregister() {
	c.dead = true
	c.hub.unregister <- c
	_ = c.conn.Close()
}

func (c *Client) serve() {
	go c.serveRead()
	go c.serveWrite()
}

// serveWrite handle sending messages to client via websocket connection
func (c *Client) serveWrite() {
	for {
		message, ok := <-c.send
		if !ok {
			log.Printf("Stopping write goroutine for client %s due to connection close", c.conn.RemoteAddr())
			break
		}
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message: %s", message)
		}
	}
}

// serveRead handles all incoming messages from websocket connection
func (c *Client) serveRead() {
	defer c.Deregister()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		parsedMessage, err := parseIncomingMessage(c, message)
		if err != nil {
			log.Printf("Error parsing client message: %s", err)
			response, err := json.Marshal(ParseError)
			if err != nil {
				log.Fatalf("Error encoding base error response: %s", err)
			}
			c.send <- response
			continue
		}
		err = dispatchClientMessage(c.hub.dispatcher, parsedMessage)
		if err != nil {
			log.Printf("Error dispatching client action: %s", err)
			break
		}
		c.receive <- message
	}
}

func Broadcast(clients []*Client, message *Response) {
	response, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error encoding base error response: %s", err)
	}
	for _, client := range clients {
		if client.dead {
			continue
		}
		go client.SendMessage(response)
	}
}
