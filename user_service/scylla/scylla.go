package scylla

import (
	"errors"
	"github.com/gocql/gocql"
	"log"
	"time"
)

type Scylla struct {
	client *gocql.Session
}

func (s *Scylla) connectToScylla() error {
	if s.client != nil {
		s.client.Close()
	}

	var session *gocql.Session
	var err error

	for i := 0; i < 3; i++ {
		cluster := gocql.NewCluster("scylla_node1")
		session, err = cluster.CreateSession()
		if err == nil {
			log.Println("Scylla Session Created")
			break
		}

		log.Println("ERROR_CONNECT_RECONNECTING")
		time.Sleep(6 * time.Second)
	}

	if session == nil {
		return errors.New("ERROR_CONNECT_SCYLLA")
	}

	s.client = session
	return nil
}

func createKeyspace() string {
	return `CREATE KEYSPACE IF NOT EXISTS warehouse
			WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1 };`
}

func CreateAndConnect() *Scylla {
	s := &Scylla{}
	if err := s.connectToScylla(); err != nil {
		panic(err)
	}
	if err := s.Query(createKeyspace()); err != nil {
		log.Fatal(err)
	}
	if err := s.createAllTables(); err != nil {
		log.Fatal(err)
	}
	if err := s.createAllIndexes(); err != nil {
		log.Fatal(err)
	}
	return s
}

func (s *Scylla) createAllTables() error {
	if err := s.createUserTable(); err != nil {
		return err
	}
	if err := s.createGoodsTable(); err != nil {
		return err
	}
	return nil
}

func (s *Scylla) createAllIndexes() error {
	if err := s.createUsersByStatusIndex(); err != nil {
		return err
	}
	if err := s.createUsersByLoginIndex(); err != nil {
		return err
	}
	if err := s.createGoodsByNameIndex(); err != nil {
		return err
	}
	return nil
}

func (s *Scylla) Close() {
	s.client.Close()
}

func (s *Scylla) Query(q string) error {
	if err := s.client.Query(q).Exec(); err != nil {
		return err
	}
	return nil
}
