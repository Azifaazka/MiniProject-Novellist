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

func TestGetAllStory(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryStory(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `cerita`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "tanggal", "judul", "genre", "deskripsi", "iscompleted", "iduser"}).
	AddRow(1, "10-05-2019", "Dia", "Romatis", "dia, sesosok yang tidak pernah diketahui sebelumnya", "1", "1").
	AddRow(2, "15-01-2021", "Tahu", "Drama", "tahu masakan ibu", "0", "1").
	AddRow(3, "21-10-2022", "Kuntilanak 4", "Horror", "siapa pemilik ketawa itu?", "0", "2"))

	res := repo.GetAllStory()

	assert.Equal(t, res[0].Judul, "Dia")
	assert.Len(t, res, 3)
}

func TestDeleteStoryByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryStory(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteStoryByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateStoryByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryStory(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneStoryByID(1, model.Cerita{
		Judul: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}