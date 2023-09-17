package customer

import (
	"context"
	"fmt"
	"go-testcontainers/database"
	testutils "go-testcontainers/test-utils"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IntTestSuite struct {
	suite.Suite
	db              *gorm.DB
	customerService CustomerService
}

func TestIntTestSuite(t *testing.T) {
	suite.Run(t, &IntTestSuite{})
}

func (its *IntTestSuite) TestCreateCustomer() {
	var customer Customer
	customer.Email = "talabi@mail.com"
	customer.Name = "Talabi"

	actual, err := its.customerService.CreateCustomer(customer)

	its.Nil(err)
	its.Equal(int64(1), actual.ID)
}

func (its *IntTestSuite) SetupSuite() {
	its.T().Log("setting up database")
	ctx := context.Background()
	_, err := testutils.CreateContainer(ctx)
	if err != nil {
		its.FailNowf("failed to establish database connection", err.Error())
	}

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", testutils.Pport, "", "", "mytest")
	db, err := gorm.Open(postgres.Open(postgresqlDbInfo), &gorm.Config{})
	if err != nil {
		its.FailNowf("failed to establish database connection", err.Error())
	}

	database.DbContext = db

	setupDatabase(its, db)

	its.db = db
	its.customerService = NewCustomerService()
}

func (its *IntTestSuite) TearDownSuite() {
	tearDownDatabase(its)
}

func (its *IntTestSuite) TearDownTest() {
	cleanTable(its)
}

func setupDatabase(its *IntTestSuite, db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}

func seedTestTable(its *IntTestSuite, db *gorm.DB) {
	tx := db.Exec(`INSERT INTO customers (email, name) VALUES ($1, $2)`, "talabi@mail.com", "Talabi")
	if tx.Error != nil {
		its.FailNowf("failed to seed table", tx.Error.Error())
	}
}

func cleanTable(its *IntTestSuite) {
	tx := its.db.Exec(`DELETE FROM customers`)
	if tx.Error != nil {
		its.FailNowf("failed to clean table", tx.Error.Error())
	}
}

func tearDownDatabase(its *IntTestSuite) {
	its.T().Log("tearing down database")
	tx := its.db.Exec(`DROP TABLE customers`)
	if tx.Error != nil {
		its.FailNowf("failed to drop table", tx.Error.Error())
	}
	
	err := testutils.TearDown()
	if err != nil {
		its.FailNowf("unable to close database", err.Error())
	}
}
