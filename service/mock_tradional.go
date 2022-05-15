package service

import "MiniProject-Novellist/model"

type repoMockTraditional struct {f func(id int, user model.User) error}
type repoMockTraditionalStory struct {f func(id int, cerita model.Cerita) error}
type repoMockTraditionalBab struct {f func(id int, bab model.Bab) error}
type repoMockTraditionalComment struct{f func(id int, komen model.Comment) error}

func (r *repoMockTraditional) CreateUsers(user model.User) error {panic("implement")}
func (r *repoMockTraditional) GetAll() []model.User {panic("implement")}
func (r *repoMockTraditional) GetOneByID(id int) (user model.User, err error) {panic("implement")}
func (r *repoMockTraditional) GetOneByEmail(email string) (user model.User, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateOneByID(id int, user model.User) error {return r.f(id, user)}
func (r *repoMockTraditional) DeleteByID(id int) error {panic("implement")}

func (r *repoMockTraditionalStory) CreateStorys(cerita model.Cerita) error {panic("implement")}
func (r *repoMockTraditionalStory) GetAllStory() []model.Cerita {panic("implement")}
func (r *repoMockTraditionalStory) GetOneStoryByID(id int) (cerita model.Cerita, err error) {panic("implement")}
func (r *repoMockTraditionalStory) UpdateOneStoryByID(id int, cerita model.Cerita) error {return r.f(id, cerita)}
func (r *repoMockTraditionalStory) DeleteStoryByID(id int) error {panic("implement")}

func (r *repoMockTraditionalBab) CreateBabs(bab model.Bab) error {panic("implement")}
func (r *repoMockTraditionalBab) GetAllBab() []model.Bab {panic("implement")}
func (r *repoMockTraditionalBab) GetOneBabByID(id int) (bab model.Bab, err error) {panic("implement")}
func (r *repoMockTraditionalBab) UpdateOneBabByID(id int, bab model.Bab) error {return r.f(id, bab)}
func (r *repoMockTraditionalBab) DeleteBabsByID(id int) error {panic("implement")}

func (r *repoMockTraditionalComment) CreateCommit(komen model.Comment) error {panic("implement")}
func (r *repoMockTraditionalComment) GetAllCommit() []model.Comment {panic("implement")}
func (r *repoMockTraditionalComment) DeleteCommitByID(id int) error {panic("implement")}