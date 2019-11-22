package ws

import (
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
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()

	go c.serveRead()
	go c.serveWrite()
}

func (c *Client) serveWrite() {
	for {
		message := <-c.send
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)
	}
}

func (c *Client) serveRead() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		log.Printf("Message from websocket: %s", message)
		c.receive <- message
	}
}
