package models

import (
	"errors"
	"github.com/Tijanieneye10/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = int(id)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetSingleEvent(eventId int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) DeleteEvent() error {

	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) UpdateEvent() error {

	query := `UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	where id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New(err.Error())
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}
