package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldCallPing(t *testing.T) {
	
	// Setup
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("Error '%s' was not expected while opening stub database connection.",err)
	}

	defer db.Close()
	mock.ExpectPing()

	// Act
	err = Ping(db)
	if err != nil {
		t.Fatalf("Error occurred while executing ping: '%s'",err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Unmet expectations: %s",err)
	}

}