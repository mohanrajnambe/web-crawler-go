package models

import (
	"sync"
)

type Crawler struct {
	Visited map[string]bool
	Mutex   sync.Mutex
}
