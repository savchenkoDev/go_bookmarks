package repository

import (
	"errors"
	
	"bookmarks/internal/models"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) GetTagByIDAndUserID(id int64, userID int64) (models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&tag).Error
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (r *TagRepository) GetTagsByUserID(userID int64) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Where("user_id = ?", userID).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) Create(t models.TagRequest) (models.Tag, error) {
	var tag models.Tag
	tag = models.Tag{
		UserID: t.UserID,
		Name: t.Name,
	}
	err := r.db.Create(&tag).Error
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (r *TagRepository) Update(userID int64, id int64, tr models.TagUpdateRequest) (models.Tag, error) {
	if tr.Name == nil {
		return models.Tag{}, errors.New("nothing to update")
	}

	t, err := r.GetTagByIDAndUserID(id, userID)
	if err != nil {
		return models.Tag{}, err
	}
  result := r.db.Model(&t).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"name": *tr.Name,
		})
	if result.Error != nil {
			return models.Tag{}, result.Error
	}

	return r.GetTagByIDAndUserID(id, userID)
}

func (r *TagRepository) Delete(userID int64, id int64) error {
	err := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Tag{}).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&models.Tag{}, id).Error
}