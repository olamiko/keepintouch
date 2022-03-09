package dbutils_test

import (
	//"fmt"
	"github.com/olamiko/keepintouch/dbutils"
	"os"
	"testing"
)

func testSetup() {
	dbutils.DB.Init("./test.db", "test")
}

func testTeardown() {
	os.Remove("./test.db")
}

func TestDBInit(t *testing.T) {
	db, _ := dbutils.DB.Init("./test.db", "test")
	if db == nil {
		t.Errorf("d")
	}
}
