// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"encoding/json"
)

type Message struct {
	ID      int64
	Content json.RawMessage
}
