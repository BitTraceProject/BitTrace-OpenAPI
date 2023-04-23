package tests

import (
	"testing"
	"time"

	"github.com/BitTraceProject/BitTrace-OpenAPI/server"
	"github.com/BitTraceProject/BitTrace-Types/pkg/config"
)

var (
	s *server.OpenAPIServer
)

func init() {
	s = server.InitOpenAPIServer(config.DatabaseConfig{
		Address:  "master.collector.bittrace.proj:33062",
		Username: "openapi",
		Password: "admin",
	})
}

func TestPeer(t *testing.T) {
	peers, err := server.QueryPeer()
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	t.Logf("peers:%+v", peers)
}

func TestRaw(t *testing.T) {
	now := time.Now()
	tables := make([]string, 0)
	db, err := s.Raw("SHOW TABLES")
	if err != nil {
		t.Fatal(err)
	}
	db = db.Scan(&tables)
	if db.Error != nil {
		t.Fatal(db.Error)
	}
	t.Logf("%+v, cost:%s", tables, time.Until(now))

	db, err = s.Raw("SHOW TABLES")
	if err != nil {
		t.Fatal(err)
	}
	db = db.Scan(&tables)
	if db.Error != nil {
		t.Fatal(db.Error)
	}
	t.Logf("%+v, cost:%s", tables, time.Until(now))

	db, err = s.Raw("SHOW TABLES")
	if err != nil {
		t.Fatal(err)
	}
	db = db.Scan(&tables)
	if db.Error != nil {
		t.Fatal(db.Error)
	}
	t.Logf("%+v, cost:%s", tables, time.Until(now))

	db, err = s.Raw("SHOW TABLES")
	if err != nil {
		t.Fatal(err)
	}
	db = db.Scan(&tables)
	if db.Error != nil {
		t.Fatal(db.Error)
	}
	t.Logf("%+v, cost:%s", tables, time.Until(now))
}
