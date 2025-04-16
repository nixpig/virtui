package connection

import (
	"database/sql"
)

type Connection struct {
	ID          int
	URL         string
	Autoconnect bool
}

type ConnectionStore interface {
	HasConnections() (bool, error)
	InsertConnection(connection *Connection) error
	GetConnections() ([]Connection, error)
	GetConnectionByID(id int) (*Connection, error)
	DeleteConnectionByID(id int) error
}

type ConnectionStoreImpl struct {
	db *sql.DB
}

func NewConnectionStoreImpl(db *sql.DB) ConnectionStore {
	return &ConnectionStoreImpl{db}
}

func (c ConnectionStoreImpl) HasConnections() (bool, error) {
	query := `select count(*) from connections_;`

	var count int
	row := c.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (c ConnectionStoreImpl) InsertConnection(connection *Connection) error {
	query := `insert into connections_ (
		url_,
		autoconnect_
	) values (
		$url,
		$autoconnect
	);`

	if _, err := c.db.Exec(
		query,
		sql.Named("url", connection.URL),
		sql.Named("autoconnect", connection.Autoconnect),
	); err != nil {
		return err
	}

	return nil
}

func (c ConnectionStoreImpl) GetConnections() ([]Connection, error) {
	query := `select id_, url_, autoconnect_ from connections_;`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	var connections []Connection
	for rows.Next() {
		var connection Connection
		if err := rows.Scan(
			&connection.ID,
			&connection.URL,
			&connection.Autoconnect,
		); err != nil {
			return nil, err
		}

		connections = append(connections, connection)
	}

	return connections, nil
}

func (c ConnectionStoreImpl) GetConnectionByID(id int) (*Connection, error) {
	query := `select id_, url_, autoconnect_ from connections_ where id_ = $id;`

	row := c.db.QueryRow(query, sql.Named("id", id))

	var connection Connection
	if err := row.Scan(
		&connection.ID,
		&connection.URL,
		&connection.Autoconnect,
	); err != nil {
		return nil, err
	}

	return &connection, nil
}

func (c ConnectionStoreImpl) DeleteConnectionByID(id int) error {
	query := `delete from connections_ where id_ = $id;`

	if _, err := c.db.Exec(query, sql.Named("id", id)); err != nil {
		return err
	}

	return nil
}
