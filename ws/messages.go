package ws

const (
	TypeResponse  = "response"
	TypeBroadcast = "broadcast"
	StatusOK      = "ok"
	StatusError   = "error"
)

type Response struct {
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Description string                 `json:"description,omitempty"`
	Body        map[string]interface{} `json:"body,omitempty"`
}
