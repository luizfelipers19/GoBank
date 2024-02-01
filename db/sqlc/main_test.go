package db

var testQueries *Queries

const {
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/gobank?sslmode=disable"
}

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
}
