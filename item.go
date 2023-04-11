// Package tavern holds all the entities that are shared across packages
package tavern

import "github.com/google/uuid"

// Item is an entity that represents an item in all domains
type Item struct {
	// ID is the entity identifier
	ID          uuid.UUID
	Name        string
	Desctiption string
}
