package model

type Cerita struct {
	ID		 	int 		`json:"id" gorm:"primaryKey"`
	Tanggal 	timestamp  	`json:"tanggal"`
	Judul 		string 		`json:"judul"`
	Genre		string 		`json:"genre`
	Deskripsi 	string  	`json:"deskripsi"`
	IsCompleted bool 		`json:"iscompleted"`
	ID_User 	User 		`json:"id_user" gorm:"foreignkey:"name.User"`
}