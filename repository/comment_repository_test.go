package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllComment(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryComment(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "deskripsi", "idcerita", "iduser"}).
	AddRow(1, "Cerita ini bagus sekali!", "1", "1").
	AddRow(2, "Keren!", "2", "2").
	AddRow(3, "kumenunggu novel cetaknya!", "1", "2"))

	res := repo.GetAllComment()
	assert.Equal(t, res[0].IDCerita, 0)
	assert.Len(t, res, 3)
}

func TestDeleteCommentByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryComment(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteCommentByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}