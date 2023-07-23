package models

import "sync"

type LocalCache struct {
	pairs map[string]CachedPair
	mute  sync.Mutex
	wg    sync.WaitGroup
	quit  chan bool
}
