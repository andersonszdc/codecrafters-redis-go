package main

import "time"

type Storage struct {
	data map[string]ValueWithExpiry
}

type ValueWithExpiry struct {
	value     string
	expiresAt time.Time
}

func (v ValueWithExpiry) IsExpired() bool {
	if v.expiresAt.IsZero() {
		return false
	}

	return v.expiresAt.Before(time.Now())
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]ValueWithExpiry),
	}
}

func (kv *Storage) Get(key string) (string, bool) {
	valueWithExpiry, ok := kv.data[key]
	if !ok {
		return "", false
	}

	if valueWithExpiry.IsExpired() {
		delete(kv.data, key)
		return "", false
	}

	return valueWithExpiry.value, true
}

func (kv *Storage) Set(key string, value string) {
	kv.data[key] = ValueWithExpiry{value: value}
}

func (kv *Storage) SetWithExpiry(key string, value string, expiry time.Duration) {
	kv.data[key] = ValueWithExpiry{
		value:     value,
		expiresAt: time.Now().Add(expiry),
	}
}
