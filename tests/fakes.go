package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-clicker/app/common"
	"net/http/httptest"
)

// Fakes -
type Fakes struct {
	DB       *gorm.DB
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
}

// CreateFakes -
func CreateFakes() Fakes {
	db := createFakeDB()
	r, c := createFakeContext(db)

	return Fakes{
		DB:       db,
		Recorder: r,
		Context:  c,
	}
}

// Creates a fake db connection, ignoring errors
func createFakeDB() *gorm.DB {
	// Create fakes
	conn, _, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", conn)

	// Return only the gorm DB
	return db
}

// Creates a fake gin context
func createFakeContext(db *gorm.DB) (*httptest.ResponseRecorder, *gin.Context) {
	// Create fakes
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	// Set context parameters
	c.Set(common.KContextDB, db)

	// Return
	return r, c
}
