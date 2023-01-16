package scylla

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type User struct {
	ID       gocql.UUID `cql:"id"`
	Name     string     `cql:"name"`
	Login    string     `cql:"login"`
	Password string     `cql:"password"`
	Status   string     `cql:"status"`
	Type     string     `cql:"type"`
}

func (s *Scylla) createUserTable() error {
	q := `CREATE TABLE IF NOT EXISTS warehouse.users(
			id uuid,
			name text,
			login text,
			password text,
			status text,
			type text,
			PRIMARY KEY (id));`

	return s.Query(q)
}

func (s *Scylla) createUsersByStatusIndex() error {
	q := "CREATE INDEX IF NOT EXISTS users_by_status ON warehouse.users(status);"
	return s.Query(q)
}

func (s *Scylla) createUsersByLoginIndex() error {
	q := "CREATE INDEX IF NOT EXISTS users_by_login ON warehouse.users(login);"
	return s.Query(q)
}

func (s *Scylla) CreateUser(name, login, password, userType string) *User {
	u := &User{Name: name, Login: login, Password: password, Type: userType}
	u.ID = gocql.TimeUUID()
	u.Status = "active"
	if err := s.client.Query(`
		INSERT INTO warehouse.users(id, name, login, password, status, type)
		VALUES (?, ?, ?, ?, ?, ?)`,
		u.ID, u.Name, u.Login, u.Password, u.Status, u.Type).Exec(); err != nil {
		log.Fatal(err)
	}
	return u
}

func (s *Scylla) GetUserByLogin(login string) *User {
	u := &User{}
	if err := s.client.Query(`SELECT * FROM warehouse.users WHERE login = ? LIMIT 1`, login).Consistency(gocql.One).
		Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Status, &u.Type); err != nil {
		log.Fatal(err)
	}
	return u
}
func (s *Scylla) GetUserByID(id gocql.UUID) *User {
	u := &User{}
	if err := s.client.Query(`SELECT * FROM warehouse.users WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).
		Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Status, &u.Type); err != nil {
		log.Fatal(err)
	}
	return u
}

func (s *Scylla) GetUsersByStatus(status string) []*User {
	var users []*User

	scanner := s.client.Query(`SELECT * FROM warehouse.users WHERE status = ?`, status).Iter().Scanner()
	for scanner.Next() {
		u := &User{}

		if err := scanner.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Status, &u.Type); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Get User:", u)
		users = append(users, u)
	}
	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (s *Scylla) UpdateUserByStatus(status string, id gocql.UUID) {
	switch status {
	case "ban", "deleted":
	default:
		log.Fatal(">UpdateUserByStatus: Invalid status (should be 'ban' or 'deleted')")
	}
	if err := s.client.Query("UPDATE warehouse.users SET status = ? WHERE id = ?", status, id).Exec(); err != nil {
		log.Fatal(">UpdateUserByStatus:", err)
	}
}

func (s *Scylla) BanUserByID(id gocql.UUID) {
	s.UpdateUserByStatus("ban", id)
}

func (s *Scylla) DeleteUser(id string) {
	if err := s.client.Query("DELETE FROM warehouse.users WHERE id = ?", id).Exec(); err != nil {
		log.Fatal(">DeleteUser:", err)
	}
}
