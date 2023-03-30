package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because it has no identifier and is immutable
type Transaction struct {
	amount     int
	from       uuid.UUID
	to         uuid.UUID
	created_at time.Time
}
