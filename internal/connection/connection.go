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

func InsertConnection(db *sql.DB, connection *Connection) error {
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

	if _, err := db.Exec(
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

func GetConnections(db *sql.DB) ([]Connection, error) {
	query := `select id_, hypervisor_, url_, autoconnect_, ssh_, hostname_, username_, password_ 
		from connections_;`

	rows, err := db.Query(query)
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

func DeleteConnection(db *sql.DB, id int) error {
	query := `delete from connections_ where id_ = $id`

	if _, err := db.Exec(query, sql.Named("id", id)); err != nil {
		return err
	}

	return nil
}
