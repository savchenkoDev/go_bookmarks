package repository

import (
	"bookmarks/internal/errors"
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
		return models.Tag{}, errors.NewError(err)
	}
	return tag, nil
}

func (r *TagRepository) GetTagsByUserID(userID int64) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Where("user_id = ?", userID).Find(&tags).Error
	if err != nil {
		return nil, errors.NewError(err)
	}
	return tags, nil
}

func (r *TagRepository) Create(t models.TagRequest) (models.Tag, error) {
	tag := models.Tag{
		UserID: t.UserID,
		Name:   t.Name,
	}
	err := r.db.Create(&tag).Error
	if err != nil {
		return models.Tag{}, errors.NewError(err)
	}
	return tag, nil
}

func (r *TagRepository) Update(userID int64, id int64, tr models.TagUpdateRequest) (models.Tag, error) {
	if tr.Name == nil {
		return models.Tag{}, errors.RecordInvalidError()
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
		return models.Tag{}, errors.NewError(result.Error)
	}

	return r.GetTagByIDAndUserID(id, userID)
}

func (r *TagRepository) Delete(userID int64, id int64) error {
	err := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Tag{}).Error
	if err != nil {
		return errors.NewError(err)
	}
	err = r.db.Delete(&models.Tag{}, id).Error
	if err != nil {
		return errors.NewError(err)
	}
	return nil
}

const TAG_STAT_QUERY = `
	SELECT
		COUNT(id) as count
	FROM tags
	WHERE user_id = ?
`

func (r *TagRepository) CalculateTagStats(userID int64, stats *models.UserStats) error {
	var count int64
	err := r.db.Raw(TAG_STAT_QUERY, userID).Scan(&count).Error
	if err != nil {
		return errors.NewError(err)
	}
	stats.Tags = count
	return nil
}
