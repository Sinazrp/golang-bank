package db

import (
	"database/sql"
	"github.com/sinazrp/golang-bank/util"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

// TestMain is the main entry point for running all the test cases.
// It initializes the test database connection and store the connection
// in the testDB variable and store the Queries in the testQueries variable.
var testQueries *Queries
var testDB *sql.DB

// TestMain is the main entry point for running all the test cases.
//
// It first loads a configuration from the parent directory of the current
// working directory, then connects to the configured database, and finally
// creates a new store from the database connection. The test queries are
// stored in the testQueries variable and the database connection is stored
// in the testDB variable.
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbDriver := config.DBDriver
	dbSource := config.DBSource
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	m.Run()
}
func TestWithTx(t *testing.T) {

}
func RandomAccount() CreateAccountParams {
	return CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

}
func RandomUser() CreateUserParams {
	return CreateUserParams{
		Username:       util.RandomOwner(),
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		HashedPassword: "",
	}
}
