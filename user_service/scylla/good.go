package scylla

import (
	"github.com/gocql/gocql"
	"log"
)

type Good struct {
	ID          gocql.UUID `cql:"id"`
	Name        string     `cql:"name"`
	Image       string     `cql:"image"`
	Description string     `cql:"description"`
	Status      string     `cql:"status"`
	Price       int        `cql:"price"`
}

func (s *Scylla) createGoodsTable() error {
	q := `CREATE TABLE IF NOT EXISTS warehouse.goods(
			id uuid,
			name text,
			image text,
			description text,
			status text,
			price int,
			PRIMARY KEY (id));`

	return s.Query(q)
}

func (s *Scylla) createGoodsByNameIndex() error {
	q := "CREATE INDEX IF NOT EXISTS goods_by_name ON warehouse.goods(name);"
	return s.Query(q)
}

func (s *Scylla) CreateGood(name, desc, img string, price int) *Good {
	g := &Good{Name: name, Description: desc, Image: img, Price: price}
	g.ID = gocql.TimeUUID()
	g.Status = "created"
	if err := s.client.Query(`
		INSERT INTO warehouse.goods(id, name, image, description, status, price) 
		VALUES (?, ?, ?, ?, ?, ?)`,
		g.ID, g.Name, g.Image, g.Description, g.Status, g.Price).Exec(); err != nil {
		log.Fatal(err)
	}
	return g
}

func (s *Scylla) GetGoodByID(id gocql.UUID) *Good {
	g := &Good{}
	if err := s.client.Query(`SELECT * FROM warehouse.goods WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).
		Scan(&g.ID, &g.Name, &g.Description, &g.Image, &g.Price, &g.Status); err != nil {
		log.Fatal(err)
	}
	return g
}

func (s *Scylla) GetGoodByName(name string) *Good {
	g := &Good{}
	if err := s.client.Query(`SELECT * FROM warehouse.goods WHERE name = ? LIMIT 1`, name).Consistency(gocql.One).
		Scan(&g.ID, &g.Name, &g.Description, &g.Image, &g.Price, &g.Status); err != nil {
		log.Fatal(err)
	}
	return g
}

func (s *Scylla) UpdateGoodByStatus(status string, id gocql.UUID) {
	if err := s.client.Query("UPDATE warehouse.goods SET status = ? WHERE id = ?", status, id).Exec(); err != nil {
		log.Fatal(">UpdateGoodByStatus:", err)
	}
}

func (s *Scylla) UpdateGoodStatusToDeleted(id gocql.UUID) {
	if err := s.client.Query("UPDATE warehouse.goods SET status = ? WHERE id = ?", "deleted", id).Exec(); err != nil {
		log.Fatal(">UpdateGoodByStatus:", err)
	}
}

func (s *Scylla) DeleteGood(id string) {
	if err := s.client.Query("DELETE FROM warehouse.goods WHERE id = ?", id).Exec(); err != nil {
		log.Fatal(">DeleteUser:", err)
	}
}
