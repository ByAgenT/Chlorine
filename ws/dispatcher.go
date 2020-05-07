package ws

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dispatcher struct {
	actions map[string]func(message *ClientMessage) *Response
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{actions: make(map[string]func(message *ClientMessage) *Response)}
}

// TODO: star vs no-star
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
