package model

type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
