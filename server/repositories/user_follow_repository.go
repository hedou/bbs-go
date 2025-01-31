package repositories

import (
	"bbs-go/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var UserFollowRepository = newUserFollowRepository()

func newUserFollowRepository() *userFollowRepository {
	return &userFollowRepository{}
}

type userFollowRepository struct {
}

func (r *userFollowRepository) Get(db *gorm.DB, id int64) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userFollowRepository) Take(db *gorm.DB, where ...interface{}) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userFollowRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFollow) {
	cnd.Find(db, &list)
	return
}

func (r *userFollowRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userFollowRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserFollow, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userFollowRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFollow, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserFollow{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userFollowRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.UserFollow{})
}

func (r *userFollowRepository) Create(db *gorm.DB, t *model.UserFollow) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userFollowRepository) Update(db *gorm.DB, t *model.UserFollow) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userFollowRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserFollow{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userFollowRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserFollow{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userFollowRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserFollow{}, "id = ?", id)
}

