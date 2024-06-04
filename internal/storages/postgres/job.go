package postgresql

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"birthdays/internal/utils"
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgJobStorage struct {
	db *sqlx.DB
}

func NewJobStorage(db *sqlx.DB) *PgJobStorage {
	return &PgJobStorage{
		db: db,
	}
}

func (s *PgJobStorage) GetAll(ctx context.Context) ([]*entities.Job, error) {
	logger, err := utils.GetLogger(ctx)
	if err != nil {
		return nil, err
	}

	query, args, err := squirrel.
		Select(userJobSelectFields...).
		From(userTable).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Query built")

	var users []*dto.JobUser
	if err := s.db.Select(&users, query, args...); err != nil {
		logger.Debug("User select failed", "err", err.Error())
		return nil, apperrors.ErrUserNotSelected
	}
	logger.Debug("Query executed")

	if len(users) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	jobs := make([]*entities.Job, len(users))
	for i, user := range users {
		jobs[i] = &entities.Job{
			Date:       user.DOB,
			SourceID:   user.ID,
			SourceName: user.Name + user.Surname,
		}
	}
	logger.Debug("Job list built")

	return jobs, nil
}

func (s *PgJobStorage) Add(ctx context.Context, job entities.Job) error {
	return nil
}

func (s *PgJobStorage) GetRecipientEmails(ctx context.Context, sourceID uint64) ([]string, error) {
	logger, err := utils.GetLogger(ctx)
	if err != nil {
		return nil, err
	}

	query, args, err := squirrel.
		Select(userJobEmailSelectFields...).
		From(userTable).
		LeftJoin(notificationTable + " ON " + userIdField + " = " + idSubscriberField).
		Where(squirrel.Eq{idSourceField: sourceID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Query built")

	var emails []string
	if err := s.db.Select(&emails, query, args...); err != nil {
		logger.Debug("Email select failed", "err", err.Error())
		return nil, apperrors.ErrEmailsNotSelected
	}
	logger.Debug("Query executed")

	if len(emails) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	return emails, nil
}
