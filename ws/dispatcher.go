package ws

import (
	"encoding/json"
	"fmt"
	"log"
)

// Dispatcher store available actions for a client and provide methods to pass client request to specific
// action handler and return response from action handler back to client.
type Dispatcher struct {
	actions map[string]func(message *ClientMessage) *Response
}

// NewDispatcher creates new Dispatcher instance with empty action map.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{actions: make(map[string]func(message *ClientMessage) *Response)}
}

// AttachAction add new action to dispatcher's action map.
// TODO: determine if pointer access is required here.
func (d *Dispatcher) AttachAction(actionName string, action func(message *ClientMessage) *Response) {
	d.actions[actionName] = action
}

func (d *Dispatcher) dispatch(action string, message *ClientMessage) (*Response, error) {
	if actionFunc, ok := d.actions[action]; ok {
		return actionFunc(message), nil
	}
	return nil, fmt.Errorf("dispatcher: dispatched action not found: %s", action)
}

func dispatchClientMessage(dispatcher *Dispatcher, message *ClientMessage) error {
	log.Printf("[WS ACTION] %s", message.Action)
	actionResponse, err := dispatcher.dispatch(message.Action, message)
	if err != nil {
		// TODO: consume new methods Is and Unwrap
		response, err := json.Marshal(generateActionNotFound(message.Action))
		if err != nil {
			log.Fatalf("Error encoding response for action %s: %s", message.Action, err)
		}
		message.Client.send <- response
		return nil
	}
	response, err := json.Marshal(actionResponse)
	if err != nil {
		log.Fatalf("Error encoding response for action %s: %s", message.Action, err)
	}
	message.Client.send <- response
	return nil
}
