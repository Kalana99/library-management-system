package repositories

import (
    "library-management-system/models"
    "gorm.io/gorm"
)

type IssueRepository struct {
    DB *gorm.DB
}

func NewIssueRepository(db *gorm.DB) *IssueRepository {
    return &IssueRepository{DB: db}
}

func (r *IssueRepository) IssueBook(issue *models.Issue) error {
    return r.DB.Create(issue).Error
}

func (r *IssueRepository) ReturnBook(issueID uint) error {
    return r.DB.Model(&models.Issue{}).Where("id = ?", issueID).Update("returned", true).Error
}

func (r *IssueRepository) GetIssuesByUserID(userID uint) ([]models.Issue, error) {
    var issues []models.Issue
    err := r.DB.Where("user_id = ? AND returned = ?", userID, false).Find(&issues).Error
    return issues, err
}
