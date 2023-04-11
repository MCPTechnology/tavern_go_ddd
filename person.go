// Package tavern holds all the entities that are shared across packages
package tavern

import "github.com/google/uuid"

// Person is an entity that represents a person in all domains
type Person struct {
	// ID is the entity identifier
	ID   uuid.UUID
	Name string
	Age  int
}
