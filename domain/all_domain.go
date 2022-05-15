package domain

import (
	"MiniProject-Novellist/model"
)

type AdapterRepository interface {

	// untuk user
	CreateUsers(user model.User) error
	GetAll() []model.User
	GetOneByID(id int) (user model.User, err error)
	GetOneByEmail(email string) (user model.User, err error)
	UpdateOneByID(id int, user model.User) error
	DeleteByID(id int) error
}

type AdapterRepositoryStory interface{

	// untuk cerita
	CreateStorys(cerita model.Cerita)error
	GetAllStory() []model.Cerita
	GetOneStoryByID(id int) (cerita model.Cerita, err error)
	UpdateOneStoryByID(id int, cerita model.Cerita) error
	DeleteStoryByID(id int) error

}

type AdapterRepositoryBab interface{
	// untuk bab
	CreateBabs(bab model.Bab)error
	GetAllBab() []model.Bab
	GetOneBabByID(id int) (bab model.Bab, err error)
	UpdateOneBabByID(id int, cerita model.Bab) error
	DeleteBabsByID(id int) error
}

type AdapterRepositoryComment interface{
	// untuk komen
	CreateComment(komen model.Comment) error 
	GetAllComment() []model.Comment 
	DeleteCommentByID(id int) error 
}


type AdapterService interface {
	//untuk user
	CreateUserService(user model.User) error
	UpdateUserService(id, idToken int, user model.User) error
	GetAllUsersService() []model.User
	GetUserByID(id int) (model.User, error)
	LoginUser(email, password string) (string, int)
	DeleteByID(id int) error
}

type AdapterServiceStory interface{
	//untuk cerita
	CreateStoryService(cerita model.Cerita) error
	UpdateStoryService(id, idToken int, cerita model.Cerita) error 
	GetAllStorysService() []model.Cerita 
	GetStoryByID(id int) (model.Cerita, error) 
	DeleteStorysByID(id int) error 
}

type AdapterServiceBab interface{
	//untuk bab
	CreateBabService(bab model.Bab) error 
	UpdateBabService(id, idToken int, bab model.Bab) error 
	GetAllBabsService() []model.Bab 
	GetBabByID(id int) (model.Bab, error) 
	DeleteBabsByID(id int) error 
}

type AdapterServiceComment interface{
	//untuk komen
	CreateCommentService(komen model.Comment) error 
	GetAllCommentsService() []model.Comment 
	DeleteCommentsByID(id int) error 
}