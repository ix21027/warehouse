package scylla

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"strings"
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

func (s *Scylla) InsertUser(name, login, password, userType string) string {
	id := gocql.TimeUUID()
	status := "active"
	if err := s.client.Query(`
		INSERT INTO warehouse.users(id, name, login, password, status, type)
		VALUES (?, ?, ?, ?, ?, ?)`,
		&id, &name, &login, &password, &status, &userType).Exec(); err != nil {
		return fmt.Sprintf("%s", err)
	}
	return fmt.Sprintf("ID: %s, Name: %s, Login: %s, Password: %s, Type: %s, Status: %s",
		id, name, login, password, userType, status,
	)
}

func (s *Scylla) GetUserByLogin(login string) string {
	var (
		id, name, status, typ string
	)

	if err := s.client.Query(`SELECT id, name, status, type FROM warehouse.users WHERE login = ? LIMIT 1`, login).Consistency(gocql.One).
		Scan(&id, &name, &status, &typ); err != nil {
		log.Fatal(err)
	}
	return "ID: " + id + ", Name: " + name + ", Login: " + login + ", Status: " + status + ", Type: " + typ
}
func (s *Scylla) GetUserByID(id string) string {
	var (
		name, login, status string
	)
	if err := s.client.Query(`SELECT name, login, status FROM warehouse.users WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).
		Scan(&name, &login, &status); err != nil {
		return fmt.Sprintf("%s", err)
	}
	return "ID:" + id + ", Name: " + name + ", Login: " + login + ", Status: " + status
}

func (s *Scylla) GetUsersByStatus(status string) string {
	var usersID []string

	scanner := s.client.Query(`SELECT id FROM warehouse.users WHERE status = ?`, status).Iter().Scanner()
	for scanner.Next() {
		var id string
		if err := scanner.Scan(&id); err != nil {
			return fmt.Sprintf("%s", err)
		}
		fmt.Println("Get User ID:", id)
		usersID = append(usersID, id)
	}
	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
	if err := scanner.Err(); err != nil {
		return fmt.Sprintf("%s", err)
	}

	return strings.Join(usersID, ",")

}

func (s *Scylla) UpdateUserByStatus(status, id string) string {
	switch status {
	case "ban", "deleted":
	default:
		return ">UpdateUserByStatus: Invalid status (should be 'ban' or 'deleted')"
	}
	if err := s.client.Query("UPDATE warehouse.users SET status = ? WHERE id = ?", status, id).Exec(); err != nil {
		return fmt.Sprintf("%s", err)
	}
	return ""
}

func (s *Scylla) BanUserByID(id string) {
	s.UpdateUserByStatus("ban", id)
}

func (s *Scylla) DeleteUser(id string) error {
	if err := s.client.Query("DELETE FROM warehouse.users WHERE id = ?", id).Exec(); err != nil {
		return fmt.Errorf(">DeleteUser:", err)
	}
	return nil
}
