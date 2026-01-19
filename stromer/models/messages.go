package models

import "sync"

type Messages struct {
	Messages map[any]bool
	Mu       sync.Mutex
}
