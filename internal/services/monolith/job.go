package monolith

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/storages"
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
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Gather"))
	oplog := httplog.LogEntry(ctx)

	var err error

	jobs, err := s.storage.GetAll(ctx)
	if err != nil {
		oplog.Error("Error getting jobs", "err", err.Error())
		return err
	}

	for _, job := range jobs {
		httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(job.SourceID))
		_, err = s.scheduler.NewJob(
			gocron.DurationJob(
				// Для дней рождения интервал был бы год -- time.Year,
				time.Minute,
			),
			gocron.NewTask(
				func() {
					emails, err := s.storage.GetRecipientEmails(ctx, job.SourceID)
					if err != nil {
						oplog.Error("Error getting recipient emails", "err", err.Error())
						return
					}
					for _, email := range emails {
						// При полноценной реализации здесь будет какое-то отправление
						oplog.Info("Sending reminder", "email", email)
					}
				},
			),
		)

		if err != nil {
			oplog.Error("Error scheduling job for source ID", dto.UserIDKey, job.SourceID, "err", err.Error())
		}
	}
	return err
}
