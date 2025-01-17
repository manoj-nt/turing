package main

type Item struct {
	ID   int
	Name string
}

// ProcessItem is an interface for the callback
type ProcessItem interface {
	Process(*Item) (bool, error)
}
