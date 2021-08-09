package tests

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-clicker/app/common"
	"go-clicker/app/game"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Fakes

var db *sql.DB
var dbmock sqlmock.Sqlmock
var rr *httptest.ResponseRecorder
var c *gin.Context

// Set up & Tear down

func setUp() {
	// Create fake db
	db, dbmock, _ = sqlmock.New()
	gdb, _ := gorm.Open("sqlite3", db)

	// Create fake recorder & gin context
	rr = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(rr)

	// Create error builder
	bundle := common.CreateLocalizationBundle()
	l := i18n.NewLocalizer(bundle, "en", "en")
	b := game.ErrorBuilder{L: l}

	// Set context parameters
	c.Set(common.KContextDB, gdb)
	c.Set(common.KContextErrorBuilder, b)
}

// Tests

func Test_CreateGame_WithInsertFailure(t *testing.T) {
	setUp()

	// DB Expectations
	dbmock.ExpectBegin()
	dbmock.
		ExpectExec(`INSERT INTO "games"`).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnError(fmt.Errorf("Unit Test Insert error"))
	dbmock.ExpectRollback()

	// Call create endpoint
	game.Create(c)

	// we make sure that all db expectations were met
	if err := dbmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"message":"Failed to create game"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_CreateGame_Successful(t *testing.T) {
	setUp()

	// DB Expectations
	dbmock.ExpectBegin()
	dbmock.
		ExpectExec(`INSERT INTO "games"`).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbmock.ExpectCommit()

	// Call create endpoint
	game.Create(c)

	// we make sure that all db expectations were met
	if err := dbmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func Test_RetrieveGame_WithSelectFailure(t *testing.T) {
	setUp()

	// DB Expectations
	dbmock.
		ExpectQuery("^SELECT .*").
		WithArgs("23-EXT-ID").
		WillReturnError(fmt.Errorf("Unit Test Select error"))

	// Call retrieve endpoint
	c.Params = []gin.Param{gin.Param{Key: "external_id", Value: "23-EXT-ID"}}
	game.Retrieve(c)

	// we make sure that all db expectations were met
	if err := dbmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check the status code is what we expect.
	expectedStatus := http.StatusNotFound
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	// Check the response body is what we expect.
	expected := `{"message":"Game not found"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_RetrieveGame_WithSelectSuccess(t *testing.T) {
	setUp()

	// DB Expectations
	dbmock.
		ExpectQuery("^SELECT .*").
		WithArgs("23-EXT-ID").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "external_id", "status", "score"}).
				AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "created", "10"),
		)

	// Call retrieve endpoint
	c.Params = []gin.Param{gin.Param{Key: "external_id", Value: "23-EXT-ID"}}
	game.Retrieve(c)

	// we make sure that all db expectations were met
	if err := dbmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check the status code is what we expect.
	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}
}
