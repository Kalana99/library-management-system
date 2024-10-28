package services

import (
    "library-management-system/models"
    "library-management-system/repositories"
)

type IssueService struct {
    repo *repositories.IssueRepository
}

func NewIssueService(repo *repositories.IssueRepository) *IssueService {
    return &IssueService{repo: repo}
}

func (s *IssueService) IssueBook(issue *models.Issue) error {
    return s.repo.IssueBook(issue)
}

func (s *IssueService) ReturnBook(issueID uint) error {
    return s.repo.ReturnBook(issueID)
}

func (s *IssueService) GetIssuesByUserID(userID uint) ([]models.Issue, error) {
    return s.repo.GetIssuesByUserID(userID)
}
