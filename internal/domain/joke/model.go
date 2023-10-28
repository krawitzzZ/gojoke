package joke

import (
	"time"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Model struct {
	ID            uuid.UUID `json:"id"`
	Body          string    `json:"body"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

func (m Model) TextShort() string { return text.Snip(m.Body, 77, "...") }
func (m Model) TextFull() string  { return m.Body }
