package models

import "sync"

type Messages struct {
	Messages []int
	Mu       sync.Mutex
}
