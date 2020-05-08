package ws

import "encoding/json"

type ClientMessage struct {
	Client *Client
	Action string
	Params map[string]interface{}
}

type RawMessage struct {
	Action string
	Params map[string]interface{}
}

func parseIncomingMessage(client *Client, message []byte) (*ClientMessage, error) {
	rawMessage := &RawMessage{}
	err := json.Unmarshal(message, &rawMessage)
	if err != nil {
		return nil, err
	}
	return &ClientMessage{
		Client: client,
		Action: rawMessage.Action,
		Params: rawMessage.Params,
	}, nil

}
