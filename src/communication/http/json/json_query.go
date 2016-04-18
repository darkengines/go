package json
import "encoding/json"

type JsonQuery struct {
	ServiceName string
	RawParameter json.RawMessage
}