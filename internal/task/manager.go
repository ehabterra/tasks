package task

import (
	"tasks/gen/tasks"
	storage "tasks/pkg/db"
)

// Manager ..
type Manager struct {
	Db storage.Db
}

// NewManager ...
func NewManager(db storage.Db) *Manager {
	return &Manager{db}
}

// List ...
func (m *Manager) List() (res tasks.StoredTaskCollection, err error) {
	err = m.Db.LoadAll(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Show ...
func (m *Manager) Show(id string) (res *tasks.StoredTask, err error) {
	res = &tasks.StoredTask{}
	err = m.Db.Load(id, res)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, &tasks.NotFound{
				Message: err.Error(),
				ID:      id,
			}
		}
		return nil, err
	}
	return res, nil
}

// Add ...
func (m *Manager) Add(p *tasks.Task) (id string, err error) {
	id, err = m.Db.NewID()

	sb := tasks.StoredTask{
		ID:          id,
		Title:       p.Title,
		Description: p.Description,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
		DueDate:     p.DueDate,
		Status:      p.Status,
		Assignee:    p.Assignee,
		Owner:       p.Owner,
	}

	if err = m.Db.Save(id, &sb); err != nil {
		return "", err
	}
	return id, nil

}

// Update ...
func (m *Manager) Update(p *tasks.UpdatePayload) (err error) {

	t := p.Task
	sb := tasks.StoredTask{
		ID:          p.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DueDate:     t.DueDate,
		Status:      t.Status,
		Assignee:    t.Assignee,
		Owner:       t.Owner,
	}

	if err = m.Db.Save(p.ID, &sb); err != nil {
		return err
	}
	return nil
}

// Remove ...
func (m *Manager) Remove(id string) (err error) {
	return m.Db.Delete(id) // internal error if not nil
}

// Status ...
func (m *Manager) Status(id string, status string) (err error) {
	res := tasks.StoredTask{}

	err = m.Db.Load(id, &res)

	res.Status = status

	if err = m.Db.Save(id, &res); err != nil {
		return err // internal error
	}

	return nil
}
