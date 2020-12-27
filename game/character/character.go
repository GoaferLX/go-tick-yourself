package character

import (
	"database/sql"
	"time"
)

type character struct {
	id         uint64
	balance    uint64
	reputation int16
	ticks      uint16
	updatedAt  time.Time
}

func (c character) updateTicks() character {
	elapsed := time.Since(c.updatedAt).Hours()
	c.ticks += uint16(elapsed)
	return c
}

// Save a character to storage - character may have been updated via http request or
// internal operations i.e. reputation changes caused by completion of a mission.
func Save(db *sql.DB, c character) error {
	// TODO: Validate character fields
	c = c.updateTicks()
	c.updatedAt = time.Now()
	// commit to storage.
	c, err := updateCharacter(db, c)
	if err != nil {
		return err
	}
	return nil
}
