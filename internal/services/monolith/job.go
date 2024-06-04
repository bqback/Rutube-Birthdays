package monolith

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/storages"
	"birthdays/internal/utils"
	"context"
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
	"github.com/go-co-op/gocron/v2"
)

type JobService struct {
	scheduler gocron.Scheduler
	storage   storages.IJobStorage
}

func NewJobService(storage storages.IJobStorage, scheduler gocron.Scheduler) *JobService {
	return &JobService{
		scheduler: scheduler,
		storage:   storage,
	}
}

func (s *JobService) Gather(ctx context.Context) error {
	logger, err := utils.GetLogger(ctx)
	if err != nil {
		return err
	}

	currentYear := time.Now().Year()

	jobs, err := s.storage.GetAll(ctx)
	if err != nil {
		logger.Error("Error getting jobs", "err", err.Error())
		return err
	}

	for _, job := range jobs {
		httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(job.SourceID))
		startDate := job.Date.AddDate(currentYear-job.Date.Year(), 0, 0)
		logger.Info("Job would've started at date", "date", startDate.String())
		startDate = time.Now().Add(15 * time.Second)
		_, err = s.scheduler.NewJob(
			gocron.DurationJob(
				// Для дней рождения интервал был бы год -- time.Year,
				time.Minute,
			),
			gocron.NewTask(
				func() {
					emails, err := s.storage.GetRecipientEmails(ctx, job.SourceID)
					if err != nil {
						logger.Error("Error getting recipient emails", "err", err.Error())
						return
					}
					for _, email := range emails {
						// При полноценной реализации здесь будет какое-то отправление
						logger.Info("Sending reminder", "email", email)
					}
				},
			),
			gocron.WithStartAt(
				gocron.WithStartDateTime(startDate),
			),
		)

		if err != nil {
			logger.Error("Error scheduling job for source ID", dto.UserIDKey, job.SourceID, "err", err.Error())
		}
	}

	return nil
}

func (s *JobService) Start() {
	s.scheduler.Start()
}
