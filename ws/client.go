package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// Channel of incoming messages
	receive chan []byte
}

func (c *Client) serve() {
	go c.serveRead()
	go c.serveWrite()
}

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

func (c *Client) serveRead() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()

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
