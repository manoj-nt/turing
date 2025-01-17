package main

import "fmt"

type Item struct {
	ID   int
	Name string
}

// ProcessItem is a callback function type
type ProcessItem func(item *Item) (bool, error)

// ProcessCollection processes a list of items using a callback
func ProcessCollection(items []Item, processItem ProcessItem) error {
	for _, item := range items {
		success, err := processItem(&item)
		if !success || err != nil {
			return fmt.Errorf("failed to process item %d: %v", item.ID, err)
		}
	}
	return nil
}

func main() {

}
