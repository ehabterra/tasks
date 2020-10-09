package api

import (
	"context"
	"strconv"
	"tasks/gen/tasks"
	storage "tasks/pkg/db"

	"github.com/dropbox/godropbox/errors"
)

// TaskManager interface
type TaskManager interface {
	List() (res tasks.StoredTaskCollection, err error)
	Show(id int) (res *tasks.StoredTask, err error)
	Add(p *tasks.Task) (err error)
	Update(p *tasks.Task) (err error)
	Remove(id int) (err error)
	Status(id int, status string) (err error)
}

// Task service example implementation.
// The example methods log the requests and return zero values.
type Task struct {
	Service TaskManager
}

// NewTask returns the tasks service implementation.
func NewTask(manager TaskManager) *Task {
	// Build and return service implementation.
	return &Task{manager}
}

// List all stored tasks
func (s *Task) List(_ context.Context, p *tasks.ListPayload) (res tasks.StoredTaskCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.List()
	if err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Show task by Email
func (s *Task) Show(ctx context.Context, p *tasks.ShowPayload) (res *tasks.StoredTask, view string, err error) {

	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.Show(p.ID)
	if err != nil {
		if errors.IsError(err, storage.ErrNotFound) {
			return nil, "", &tasks.NotFound{Message: err.Error(), ID: strconv.Itoa(p.ID)}
		}
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new task and return id.
func (s *Task) Add(_ context.Context, p *tasks.Task) (res string, err error) {
	if err = s.Service.Add(p); err != nil {
		return "", err
	}
	return res, nil
}

// Update existing task and return id.
func (s *Task) Update(_ context.Context, p *tasks.Task) (res string, err error) {
	if err = s.Service.Update(p); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Remove task from tasks data
func (s *Task) Remove(_ context.Context, p *tasks.RemovePayload) (err error) {
	return s.Service.Remove(p.ID)
}

// Status tasks by ids
func (s *Task) Status(_ context.Context, p *tasks.StatusPayload) (err error) {
	if err = s.Service.Status(p.ID, p.Status); err != nil {
		return err
	}
	return nil
}
