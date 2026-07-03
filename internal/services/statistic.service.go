package services

import (
	"bookmarks/internal/models"
	"bookmarks/internal/repository"
)

type StatisticService struct {
	bookmarkRepo *repository.BookmarkRepository
	tagsRepo *repository.TagRepository
}

func NewStatisticService(bookmarkRepo *repository.BookmarkRepository, tagsRepo *repository.TagRepository) *StatisticService {
	return &StatisticService{bookmarkRepo: bookmarkRepo, tagsRepo: tagsRepo}
}

const BOOKMARK_STAT_QUERY = `
	SELECT
		COUNT(*) as total_bookmarks,
		SUM(CASE WHEN is_favorite = true THEN 1 ELSE 0 END) as total_favorites,
		SUM(CASE WHEN is_archived = true THEN 1 ELSE 0 END) as total_archived
	FROM bookmarks
	WHERE user_id = ?
`

func (s *StatisticService) CalculateUserStats(userID int64) (models.UserStats, error) {
	stats := models.UserStats{}
	err := s.bookmarkRepo.CalculateBookmarkStats(userID, &stats)
	if err != nil {
		return models.UserStats{}, err
	}

	err = s.tagsRepo.CalculateTagStats(userID, &stats)
	if err != nil {
		return models.UserStats{}, err
	}

	return stats, nil
}