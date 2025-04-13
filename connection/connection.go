package connection

import "database/sql"

type Connection struct {
	ID          int
	Hypervisor  string
	URL         string
	Autoconnect bool
	SSH         bool
	Hostname    string
	Username    string
	Password    string
}

type ConnectionRepository interface {
	HasConnections() (bool, error)
	InsertConnection(connection *Connection) error
	GetConnections() ([]Connection, error)
	DeleteConnection(id int) error
}

type ConnectionRepositoryImpl struct {
	db *sql.DB
}

func NewConnectionRepositoryImpl(db *sql.DB) *ConnectionRepositoryImpl {
	return &ConnectionRepositoryImpl{db}
}

func (c ConnectionRepositoryImpl) HasConnections() (bool, error) {
	query := `select count * from connections_`

	var count int
	row := c.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (c ConnectionRepositoryImpl) InsertConnection(connection *Connection) error {
	query := `insert into connections_ (
		hypervisor_,
		url_,
		autoconnect_, 
		ssh_,
		hostname_,
		username_,
		password_
	) values (
		$hypervisor,
		$url,
		$autoconnect,
		$ssh,
		$hostname,
		$username,
		$password
	);`

	if _, err := c.db.Exec(
		query,
		sql.Named("hypervisor", connection.Hypervisor),
		sql.Named("url", connection.URL),
		sql.Named("autoconnect", connection.Autoconnect),
		sql.Named("ssh", connection.SSH),
		sql.Named("hostname", connection.Hostname),
		sql.Named("username", connection.Username),
		sql.Named("password", connection.Password),
	); err != nil {
		return err
	}

	return nil
}

func (c ConnectionRepositoryImpl) GetConnections() ([]Connection, error) {
	query := `select id_, hypervisor_, url_, autoconnect_, ssh_, hostname_, username_, password_ 
		from connections_;`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	var connections []Connection
	for rows.Next() {
		var connection Connection
		if err := rows.Scan(&connection); err != nil {
			return nil, err
		}

		connections = append(connections, connection)
	}

	return connections, nil
}

func (c ConnectionRepositoryImpl) DeleteConnection(id int) error {
	query := `delete from connections_ where id_ = $id`

	if _, err := c.db.Exec(query, sql.Named("id", id)); err != nil {
		return err
	}

	return nil
}
