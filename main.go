package main

func main() {

	// db before the server
	dbUrl := "postgres://postgres:postgres@localhost:5435/restgolang"
	db, err := getDB(dbUrl)
	if err != nil {
		panic("Failed to connect to the db")
	}

	addr := ":3000"
	s := NewAPIServer(addr, db)
	s.Run()
}
