package utils

import "sync/atomic"

type SyncValueBool struct {
	innerValue atomic.Value
}

func NewSyncValueBool() *SyncValueBool {
	v := SyncValueBool{}
	v.innerValue.Store(false)
	return &v
}

func (v *SyncValueBool) Set(value bool) {
	v.innerValue.Store(value)
}

func (v *SyncValueBool) Get() bool {
	return v.innerValue.Load().(bool)
}
