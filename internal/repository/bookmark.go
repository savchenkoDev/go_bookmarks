package repository

import (
	"strings"
	"bookmarks/internal/models"

	"gorm.io/gorm"
)

type BookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) GetBookmarkByIDAndUserID(id int64, userID int64) (models.Bookmark, error) {
	var b models.Bookmark
	err := r.db.Preload("Tags").Where("id = ? AND user_id = ?", id, userID).First(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

var sortFields = map[string]string{
        "created_at": "created_at",
        "updated_at": "updated_at",
        "title":      "title",
    }

func (r *BookmarkRepository) ListByUserID(userID int64, p models.BookmarkListParams) (models.PaginatedBookmarks, error) {
	query := r.db.Model(&models.Bookmark{}).Where("bookmarks.user_id = ?", userID)

	if p.IsFavorite != nil {
        query = query.Where("is_favorite = ?", *p.IsFavorite)
    }
    if p.IsArchived != nil {
        query = query.Where("is_archived = ?", *p.IsArchived)
    }
    if p.Query != "" {
        like := "%" + p.Query + "%"
        query = query.Where("title ILIKE ? OR url ILIKE ? OR description ILIKE ?", like, like, like)
    }
    if p.Tag != "" {
        query = query.Joins("JOIN bookmark_tags ON bookmark_tags.bookmark_id = bookmarks.id").
            Joins("JOIN tags ON tags.id = bookmark_tags.tag_id").
            Where("tags.name = ? AND tags.user_id = ?", p.Tag, userID)
    }

		var total int64
    if err := query.Count(&total).Error; err != nil {
        return models.PaginatedBookmarks{}, err
    }

		sortField, ok := sortFields[p.Sort]
		if !ok {
			sortField = "created_at"
		}
		order := "DESC"
    if strings.ToLower(p.Order) == "asc" {
        order = "ASC"
    }
		offset := (p.Page - 1) * p.PerPage

		var bookmarks []models.Bookmark
		err := query.Preload("Tags").Order(sortField + " " + order).Limit(p.PerPage).
		  Offset(offset).Find(&bookmarks).Error
    
		if err != nil {
      return models.PaginatedBookmarks{}, err
    }

		data := make([]models.BookmarkResponse, len(bookmarks))
    for i, b := range bookmarks {
        data[i] = b.ToResponse()
    }
    totalPages := int(total) / p.PerPage
    if int(total)%p.PerPage > 0 {
        totalPages++
    }

		return models.PaginatedBookmarks{
      Data:       data,
      Total:      total,
      Page:       p.Page,
      PerPage:    p.PerPage,
      TotalPages: totalPages,
    }, nil
		
}

func (r *BookmarkRepository) Create(userID int64, br models.BookmarkRequest) (models.Bookmark, error) {
	var b models.Bookmark
	b = models.Bookmark{
		UserID: userID,
		Title: br.Title,
		URL: br.URL,
		Description: br.Description,	
		IsFavorite: br.IsFavorite,
		IsArchived: br.IsArchived,
	}
	err := r.db.Create(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

func (r *BookmarkRepository) Delete(userID int64, id int64) error {
	err := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Bookmark{}).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&models.Bookmark{}, id).Error
}

func (r *BookmarkRepository) Update(userID int64, id int64, br models.BookmarkUpdateRequest) (models.Bookmark, error) {
	var b models.Bookmark
	err := r.db.Model(&b).Where("user_id = ? AND id = ?", userID, id).Updates(br).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	err = r.db.Save(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

func (r *BookmarkRepository) ToggleArchive(id int64) (models.Bookmark, error) {
	var b models.Bookmark
	err := r.db.Model(&b).Where("id = ?", id).First(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}

	err = r.db.Model(&b).Where("id = ?", id).Update("is_archived", !b.IsArchived).Error
	if err != nil {
		return models.Bookmark{}, err
	}

	return b, nil
}

func (r *BookmarkRepository) AttachTagToBookmark(userID int64, bookmarkID int64, tagID int64) (models.BookmarkTag, error ){
	bt := models.BookmarkTag{
		BookmarkID: bookmarkID,
		TagID: tagID,
	}
	err := r.db.Create(&bt).Error
	if err != nil {
		return models.BookmarkTag{}, err
	}

	err = r.db.
			Preload("Bookmark").
			Preload("Tag").
			First(&bt, bt.ID).Error
	if err != nil {
		return models.BookmarkTag{}, err
	}
	return bt, nil
}

const BOOKMARK_STAT_QUERY = `
	SELECT
		COUNT(id) as total_bookmarks,
		COUNT(id) FILTER (WHERE is_favorite = true) as total_favorites,
		COUNT(id) FILTER (WHERE is_archived = true) as total_archived
	FROM bookmarks
	WHERE user_id = ?
`
func (r *BookmarkRepository) CalculateBookmarkStats(userID int64, stats *models.UserStats) error {
	err := r.db.Raw(BOOKMARK_STAT_QUERY, userID).Scan(stats).Error
	if err != nil {
		return err
	}
	
	return nil
}