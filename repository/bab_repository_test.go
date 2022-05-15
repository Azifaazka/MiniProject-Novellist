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

func TestGetAllBab(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryBab(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `babs`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "tanggal", "JudulBab", "konten", "idcerita"}).
	AddRow(1, "10-05-2019", "Bab 1", "Indonesia merupakan tempat indah yang diimpikan oleh semuanya.", 1).
	AddRow(2, "15-05-2019", "Bab 2", "Dia, sesosok teman masa kecilku yang mengetahui spyro",1).
	AddRow(3, "21-10-2022", "Bab 1", "'halo', ucap nenek tersebut.", 1))

	res := repo.GetAllBab()
	assert.Equal(t, res[0].IDCerita, 0)
	assert.Len(t, res, 3)
}

func TestDeleteBabByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryBab(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteBabsByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateBabByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewMysqlRepositoryBab(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneBabByID(1, model.Bab{
		Judulbab: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}