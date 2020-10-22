// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks service
//
// Command:
// $ goa gen tasks/design

package tasks

import (
	"context"
	tasksviews "tasks/gen/tasks/views"
)

// The tasks service performs task data.
type Service interface {
	// List all stored tasks
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	List(context.Context, *ListPayload) (res StoredTaskCollection, view string, err error)
	// Show task by ID
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Show(context.Context, *ShowPayload) (res *StoredTask, view string, err error)
	// Add new task and return ID.
	Add(context.Context, *Task) (res string, err error)
	// Update existing task and return ID.
	Update(context.Context, *UpdatePayload) (res string, err error)
	// Remove task from tasks data
	Remove(context.Context, *RemovePayload) (err error)
	// change task status by id
	Status(context.Context, *StatusPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tasks"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"list", "show", "add", "update", "remove", "status"}

// ListPayload is the payload type of the tasks service list method.
type ListPayload struct {
	// View to render
	View *string
}

// StoredTaskCollection is the result type of the tasks service list method.
type StoredTaskCollection []*StoredTask

// ShowPayload is the payload type of the tasks service show method.
type ShowPayload struct {
	// ID of task to show
	ID string
	// View to render
	View *string
}

// StoredTask is the result type of the tasks service show method.
type StoredTask struct {
	ID string
	// Title of the task
	Title string
	// Description of the task
	Description string
	// Created date
	CreatedDate string
	// Udated date
	UpdatedDate string
	// due date
	DueDate *string
	// Status.
	Status string
	// Owner.
	Owner *StoredUser
	// Assignee.
	Assignee *StoredUser
}

// Task is the payload type of the tasks service add method.
type Task struct {
	// Title of the task
	Title string
	// Description of the task
	Description string
	// Created date
	CreatedDate string
	// Udated date
	UpdatedDate string
	// due date
	DueDate *string
	// Status.
	Status string
	// Owner.
	Owner *StoredUser
	// Assignee.
	Assignee *StoredUser
}

// UpdatePayload is the payload type of the tasks service update method.
type UpdatePayload struct {
	// ID of task to show
	ID   string
	Task *StoredTask
}

// RemovePayload is the payload type of the tasks service remove method.
type RemovePayload struct {
	// ID of task to remove
	ID string
}

// StatusPayload is the payload type of the tasks service status method.
type StatusPayload struct {
	// ID of task
	ID string
	// Status.
	Status string
}

// A StoredUser describes a user retrieved by the users service.
type StoredUser struct {
	// Email of the user
	Email string
	// First Name of the user
	Firstname string
	// Last Name of user
	Lastname string
	// Is user active.
	Isactive bool
	// user role
	Role string
}

// NotFound is the type returned when attempting to show or delete a task that
// does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing task
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a task that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewStoredTaskCollection initializes result type StoredTaskCollection from
// viewed result type StoredTaskCollection.
func NewStoredTaskCollection(vres tasksviews.StoredTaskCollection) StoredTaskCollection {
	var res StoredTaskCollection
	switch vres.View {
	case "default", "":
		res = newStoredTaskCollection(vres.Projected)
	case "tiny":
		res = newStoredTaskCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredTaskCollection initializes viewed result type
// StoredTaskCollection from result type StoredTaskCollection using the given
// view.
func NewViewedStoredTaskCollection(res StoredTaskCollection, view string) tasksviews.StoredTaskCollection {
	var vres tasksviews.StoredTaskCollection
	switch view {
	case "default", "":
		p := newStoredTaskCollectionView(res)
		vres = tasksviews.StoredTaskCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredTaskCollectionViewTiny(res)
		vres = tasksviews.StoredTaskCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// NewStoredTask initializes result type StoredTask from viewed result type
// StoredTask.
func NewStoredTask(vres *tasksviews.StoredTask) *StoredTask {
	var res *StoredTask
	switch vres.View {
	case "default", "":
		res = newStoredTask(vres.Projected)
	case "tiny":
		res = newStoredTaskTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredTask initializes viewed result type StoredTask from result
// type StoredTask using the given view.
func NewViewedStoredTask(res *StoredTask, view string) *tasksviews.StoredTask {
	var vres *tasksviews.StoredTask
	switch view {
	case "default", "":
		p := newStoredTaskView(res)
		vres = &tasksviews.StoredTask{Projected: p, View: "default"}
	case "tiny":
		p := newStoredTaskViewTiny(res)
		vres = &tasksviews.StoredTask{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredTaskCollection converts projected type StoredTaskCollection to
// service type StoredTaskCollection.
func newStoredTaskCollection(vres tasksviews.StoredTaskCollectionView) StoredTaskCollection {
	res := make(StoredTaskCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredTask(n)
	}
	return res
}

// newStoredTaskCollectionTiny converts projected type StoredTaskCollection to
// service type StoredTaskCollection.
func newStoredTaskCollectionTiny(vres tasksviews.StoredTaskCollectionView) StoredTaskCollection {
	res := make(StoredTaskCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredTaskTiny(n)
	}
	return res
}

// newStoredTaskCollectionView projects result type StoredTaskCollection to
// projected type StoredTaskCollectionView using the "default" view.
func newStoredTaskCollectionView(res StoredTaskCollection) tasksviews.StoredTaskCollectionView {
	vres := make(tasksviews.StoredTaskCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredTaskView(n)
	}
	return vres
}

// newStoredTaskCollectionViewTiny projects result type StoredTaskCollection to
// projected type StoredTaskCollectionView using the "tiny" view.
func newStoredTaskCollectionViewTiny(res StoredTaskCollection) tasksviews.StoredTaskCollectionView {
	vres := make(tasksviews.StoredTaskCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredTaskViewTiny(n)
	}
	return vres
}

// newStoredTask converts projected type StoredTask to service type StoredTask.
func newStoredTask(vres *tasksviews.StoredTaskView) *StoredTask {
	res := &StoredTask{
		DueDate: vres.DueDate,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Title != nil {
		res.Title = *vres.Title
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.CreatedDate != nil {
		res.CreatedDate = *vres.CreatedDate
	}
	if vres.UpdatedDate != nil {
		res.UpdatedDate = *vres.UpdatedDate
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.Status == nil {
		res.Status = "Open"
	}
	if vres.Owner != nil {
		res.Owner = newStoredUserTiny(vres.Owner)
	}
	if vres.Assignee != nil {
		res.Assignee = newStoredUserTiny(vres.Assignee)
	}
	return res
}

// newStoredTaskTiny converts projected type StoredTask to service type
// StoredTask.
func newStoredTaskTiny(vres *tasksviews.StoredTaskView) *StoredTask {
	res := &StoredTask{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Title != nil {
		res.Title = *vres.Title
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.Status == nil {
		res.Status = "Open"
	}
	if vres.Owner != nil {
		res.Owner = newStoredUser(vres.Owner)
	}
	if vres.Assignee != nil {
		res.Assignee = newStoredUserTiny(vres.Assignee)
	}
	return res
}

// newStoredTaskView projects result type StoredTask to projected type
// StoredTaskView using the "default" view.
func newStoredTaskView(res *StoredTask) *tasksviews.StoredTaskView {
	vres := &tasksviews.StoredTaskView{
		ID:          &res.ID,
		Title:       &res.Title,
		Description: &res.Description,
		CreatedDate: &res.CreatedDate,
		UpdatedDate: &res.UpdatedDate,
		DueDate:     res.DueDate,
		Status:      &res.Status,
	}
	if res.Owner != nil {
		vres.Owner = newStoredUserViewTiny(res.Owner)
	}
	if res.Assignee != nil {
		vres.Assignee = newStoredUserViewTiny(res.Assignee)
	}
	return vres
}

// newStoredTaskViewTiny projects result type StoredTask to projected type
// StoredTaskView using the "tiny" view.
func newStoredTaskViewTiny(res *StoredTask) *tasksviews.StoredTaskView {
	vres := &tasksviews.StoredTaskView{
		ID:     &res.ID,
		Title:  &res.Title,
		Status: &res.Status,
	}
	if res.Assignee != nil {
		vres.Assignee = newStoredUserViewTiny(res.Assignee)
	}
	return vres
}

// newStoredUser converts projected type StoredUser to service type StoredUser.
func newStoredUser(vres *tasksviews.StoredUserView) *StoredUser {
	res := &StoredUser{}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	if vres.Role != nil {
		res.Role = *vres.Role
	}
	if vres.Firstname != nil {
		res.Firstname = *vres.Firstname
	}
	if vres.Lastname != nil {
		res.Lastname = *vres.Lastname
	}
	if vres.Isactive != nil {
		res.Isactive = *vres.Isactive
	}
	if vres.Isactive == nil {
		res.Isactive = true
	}
	return res
}

// newStoredUserTiny converts projected type StoredUser to service type
// StoredUser.
func newStoredUserTiny(vres *tasksviews.StoredUserView) *StoredUser {
	res := &StoredUser{}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	return res
}

// newStoredUserView projects result type StoredUser to projected type
// StoredUserView using the "default" view.
func newStoredUserView(res *StoredUser) *tasksviews.StoredUserView {
	vres := &tasksviews.StoredUserView{
		Email:     &res.Email,
		Firstname: &res.Firstname,
		Lastname:  &res.Lastname,
		Isactive:  &res.Isactive,
		Role:      &res.Role,
	}
	return vres
}

// newStoredUserViewTiny projects result type StoredUser to projected type
// StoredUserView using the "tiny" view.
func newStoredUserViewTiny(res *StoredUser) *tasksviews.StoredUserView {
	vres := &tasksviews.StoredUserView{
		Email: &res.Email,
	}
	return vres
}
