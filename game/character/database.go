package character

import (
	"database/sql"
)

func getCharacter(db *sql.DB, id uint64) (character, error) {
	row := db.QueryRow("SELECT * FROM characters WHERE id = $1", id)

	var c character
	if err := row.Scan(&c.id, &c.balance, &c.reputation); err != nil {
		return character{}, err
	}

	return c, nil
}

func updateCharacter(db *sql.DB, c character) (character, error) {
	// ignoring result as its functions are not supported by pq drivers.
	_, err := db.Exec("UPDATE characters SET balance = $1, reputation = $2, updatedAt = $3 WHERE id = $4",
		c.balance, c.reputation, c.updatedAt, c.id)
	if err != nil {
		return c, err
	}

	return c, nil
}
