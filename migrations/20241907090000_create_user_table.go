package migrations

import "gofr.dev/pkg/gofr/migration"

const createTable = `CREATE TABLE IF NOT EXISTS users
(
    id  int not null auto_increment primary key,
    name    varchar(100) not null,
    email  varchar(100) not null,
    password varchar(100) not null
);`

func createTableUsers() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
