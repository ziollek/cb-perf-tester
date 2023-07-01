package interfaces

import "time"

type KV interface {
	Upsert(string, interface{}, time.Duration) error
	Get(string, interface{}) error
	LookupIn(string, []string, ...interface{}) error
}
