package services

import "context"

type IJobService interface {
	Gather(context.Context) error
	Start()
}
