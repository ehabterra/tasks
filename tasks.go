package tasksapi

import (
	"context"
	"log"
	tasks "tasks/gen/tasks"
)

// tasks service example implementation.
// The example methods log the requests and return zero values.
type taskssrvc struct {
	logger *log.Logger
}

// NewTasks returns the tasks service implementation.
func NewTasks(logger *log.Logger) tasks.Service {
	return &taskssrvc{logger}
}

// List all stored tasks
func (s *taskssrvc) List(ctx context.Context, p *tasks.ListPayload) (res tasks.StoredTaskCollection, view string, err error) {
	view = "default"
	s.logger.Print("tasks.list")
	return
}

// Show task by ID
func (s *taskssrvc) Show(ctx context.Context, p *tasks.ShowPayload) (res *tasks.StoredTask, view string, err error) {
	res = &tasks.StoredTask{}
	view = "default"
	s.logger.Print("tasks.show")
	return
}

// Add new task and return ID.
func (s *taskssrvc) Add(ctx context.Context, p *tasks.Task) (res int, err error) {
	s.logger.Print("tasks.add")
	return
}

// Update existing task and return ID.
func (s *taskssrvc) Update(ctx context.Context, p *tasks.UpdatePayload) (res int, err error) {
	s.logger.Print("tasks.update")
	return
}

// Remove task from tasks data
func (s *taskssrvc) Remove(ctx context.Context, p *tasks.RemovePayload) (err error) {
	s.logger.Print("tasks.remove")
	return
}

// change task status by id
func (s *taskssrvc) Status(ctx context.Context, p *tasks.StatusPayload) (err error) {
	s.logger.Print("tasks.status")
	return
}
