package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}

// RestQueryRes is the response shape for our db's RestQuery() method
type RestQueryRes struct {
	ID      int
	Address string
}

// RestQuery is the db query we use for our rest endpoint
func (d *Db) RestQuery() *RestQueryRes {
	rows, err := d.Query("SELECT * FROM users LIMIT 1")
	if err != nil {
		fmt.Println("RestQuery Err: ", err)
	}

	var r RestQueryRes
	for rows.Next() {
		err = rows.Scan(&r.ID, &r.Address)
		if err != nil {
			fmt.Println("rows.Next() Err: ", err)
		}
	}

	return &r
}

// GetUserByNameRes is the response shape for our db's GetUserByName() method
type GetUserByNameRes struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
	Friendly   bool   `json:"friendly"`
}

// GetUserByName is the db query we use for our graphql endpoint
func (d *Db) GetUserByName(name string) *GetUserByNameRes {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Err: ", err)
	}

	var r GetUserByNameRes
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("rows.Next() Err: ", err)
		}
	}

	return &r
}
