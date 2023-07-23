package models

type CachedPair struct {
	value          interface{}
	ExpirationTime int64
}
