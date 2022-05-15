package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "username", "password"}).
	AddRow(1, "Azka", "Zwyne@gmail.com", "Ohzyfa", "123456").
	AddRow(2, "Asta", "Gryani@gmail.com", "ReyST", "123321").
	AddRow(3, "Kimi", "Kokumi@gmail.com", "Allboutkimi", "kimi12"))

	res := repo.GetAll()
	assert.Equal(t, res[0].Name, "Azka")
	assert.Len(t, res, 3)
}

func TestDeleteByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneByID(1, model.User{
		Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}