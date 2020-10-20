// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks HTTP server types
//
// Command:
// $ goa gen tasks/design

package server

import (
	tasks "tasks/gen/tasks"
	tasksviews "tasks/gen/tasks/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "tasks" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Created date
	CreatedDate *string `form:"created_date,omitempty" json:"created_date,omitempty" xml:"created_date,omitempty"`
	// Udated date
	UpdatedDate *string `form:"updated_date,omitempty" json:"updated_date,omitempty" xml:"updated_date,omitempty"`
	// due date
	DueDate *string `form:"due_date,omitempty" json:"due_date,omitempty" xml:"due_date,omitempty"`
	// Status.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Owner.
	Owner *StoredUserRequestBody `form:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// Assignee.
	Assignee *StoredUserRequestBody `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
}

// UpdateRequestBody is the type of the "tasks" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Task *StoredTaskRequestBody `form:"task,omitempty" json:"task,omitempty" xml:"task,omitempty"`
}

// StatusRequestBody is the type of the "tasks" service "status" endpoint HTTP
// request body.
type StatusRequestBody struct {
	// ID of task
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Status.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// StoredTaskResponseCollection is the type of the "tasks" service "list"
// endpoint HTTP response body.
type StoredTaskResponseCollection []*StoredTaskResponse

// StoredTaskResponseTinyCollection is the type of the "tasks" service "list"
// endpoint HTTP response body.
type StoredTaskResponseTinyCollection []*StoredTaskResponseTiny

// ShowResponseBody is the type of the "tasks" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	ID string `form:"id" json:"id" xml:"id"`
	// Title of the task
	Title string `form:"title" json:"title" xml:"title"`
	// Description of the task
	Description string `form:"description" json:"description" xml:"description"`
	// Created date
	CreatedDate string `form:"created_date" json:"created_date" xml:"created_date"`
	// Udated date
	UpdatedDate string `form:"updated_date" json:"updated_date" xml:"updated_date"`
	// due date
	DueDate *string `form:"due_date,omitempty" json:"due_date,omitempty" xml:"due_date,omitempty"`
	// Status.
	Status string `form:"status" json:"status" xml:"status"`
	// Owner.
	Owner *StoredUserResponseBodyTiny `form:"owner" json:"owner" xml:"owner"`
	// Assignee.
	Assignee *StoredUserResponseBodyTiny `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
}

// ShowResponseBodyTiny is the type of the "tasks" service "show" endpoint HTTP
// response body.
type ShowResponseBodyTiny struct {
	ID string `form:"id" json:"id" xml:"id"`
	// Title of the task
	Title string `form:"title" json:"title" xml:"title"`
	// Assignee.
	Assignee *StoredUserResponseBodyTiny `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
	// Status.
	Status string `form:"status" json:"status" xml:"status"`
}

// ShowNotFoundResponseBody is the type of the "tasks" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing task
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredTaskResponse is used to define fields on response body types.
type StoredTaskResponse struct {
	ID string `form:"id" json:"id" xml:"id"`
	// Title of the task
	Title string `form:"title" json:"title" xml:"title"`
	// Description of the task
	Description string `form:"description" json:"description" xml:"description"`
	// Created date
	CreatedDate string `form:"created_date" json:"created_date" xml:"created_date"`
	// Udated date
	UpdatedDate string `form:"updated_date" json:"updated_date" xml:"updated_date"`
	// due date
	DueDate *string `form:"due_date,omitempty" json:"due_date,omitempty" xml:"due_date,omitempty"`
	// Status.
	Status string `form:"status" json:"status" xml:"status"`
	// Owner.
	Owner *StoredUserResponseTiny `form:"owner" json:"owner" xml:"owner"`
	// Assignee.
	Assignee *StoredUserResponseTiny `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
}

// StoredUserResponseTiny is used to define fields on response body types.
type StoredUserResponseTiny struct {
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// user role
	Role string `form:"role" json:"role" xml:"role"`
	// Is user active.
	Isactive bool `form:"isactive" json:"isactive" xml:"isactive"`
}

// StoredTaskResponseTiny is used to define fields on response body types.
type StoredTaskResponseTiny struct {
	ID string `form:"id" json:"id" xml:"id"`
	// Title of the task
	Title string `form:"title" json:"title" xml:"title"`
	// Assignee.
	Assignee *StoredUserResponseTiny `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
	// Status.
	Status string `form:"status" json:"status" xml:"status"`
}

// StoredUserResponseBodyTiny is used to define fields on response body types.
type StoredUserResponseBodyTiny struct {
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// user role
	Role string `form:"role" json:"role" xml:"role"`
	// Is user active.
	Isactive bool `form:"isactive" json:"isactive" xml:"isactive"`
}

// StoredUserRequestBody is used to define fields on request body types.
type StoredUserRequestBody struct {
	// Email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First Name of the user
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	// Last Name of user
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	// Is user active.
	Isactive *bool `form:"isactive,omitempty" json:"isactive,omitempty" xml:"isactive,omitempty"`
	// user role
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

// StoredTaskRequestBody is used to define fields on request body types.
type StoredTaskRequestBody struct {
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Created date
	CreatedDate *string `form:"created_date,omitempty" json:"created_date,omitempty" xml:"created_date,omitempty"`
	// Udated date
	UpdatedDate *string `form:"updated_date,omitempty" json:"updated_date,omitempty" xml:"updated_date,omitempty"`
	// due date
	DueDate *string `form:"due_date,omitempty" json:"due_date,omitempty" xml:"due_date,omitempty"`
	// Status.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Owner.
	Owner *StoredUserRequestBody `form:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// Assignee.
	Assignee *StoredUserRequestBody `form:"assignee,omitempty" json:"assignee,omitempty" xml:"assignee,omitempty"`
}

// NewStoredTaskResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "tasks" service.
func NewStoredTaskResponseCollection(res tasksviews.StoredTaskCollectionView) StoredTaskResponseCollection {
	body := make([]*StoredTaskResponse, len(res))
	for i, val := range res {
		body[i] = marshalTasksviewsStoredTaskViewToStoredTaskResponse(val)
	}
	return body
}

// NewStoredTaskResponseTinyCollection builds the HTTP response body from the
// result of the "list" endpoint of the "tasks" service.
func NewStoredTaskResponseTinyCollection(res tasksviews.StoredTaskCollectionView) StoredTaskResponseTinyCollection {
	body := make([]*StoredTaskResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalTasksviewsStoredTaskViewToStoredTaskResponseTiny(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "tasks" service.
func NewShowResponseBody(res *tasksviews.StoredTaskView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: *res.Description,
		CreatedDate: *res.CreatedDate,
		UpdatedDate: *res.UpdatedDate,
		DueDate:     res.DueDate,
		Status:      *res.Status,
	}
	if res.Owner != nil {
		body.Owner = marshalTasksviewsStoredUserViewToStoredUserResponseBodyTiny(res.Owner)
	}
	if res.Assignee != nil {
		body.Assignee = marshalTasksviewsStoredUserViewToStoredUserResponseBodyTiny(res.Assignee)
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "tasks" service.
func NewShowResponseBodyTiny(res *tasksviews.StoredTaskView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:     *res.ID,
		Title:  *res.Title,
		Status: *res.Status,
	}
	if res.Assignee != nil {
		body.Assignee = marshalTasksviewsStoredUserViewToStoredUserResponseBodyTiny(res.Assignee)
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "tasks" service.
func NewShowNotFoundResponseBody(res *tasks.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewListPayload builds a tasks service list endpoint payload.
func NewListPayload(view *string) *tasks.ListPayload {
	v := &tasks.ListPayload{}
	v.View = view

	return v
}

// NewShowPayload builds a tasks service show endpoint payload.
func NewShowPayload(id string, view *string) *tasks.ShowPayload {
	v := &tasks.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddTask builds a tasks service add endpoint payload.
func NewAddTask(body *AddRequestBody) *tasks.Task {
	v := &tasks.Task{
		Title:       *body.Title,
		Description: *body.Description,
		CreatedDate: *body.CreatedDate,
		UpdatedDate: *body.UpdatedDate,
		DueDate:     body.DueDate,
		Status:      *body.Status,
	}
	if body.Owner != nil {
		v.Owner = unmarshalStoredUserRequestBodyToTasksStoredUser(body.Owner)
	}
	if body.Assignee != nil {
		v.Assignee = unmarshalStoredUserRequestBodyToTasksStoredUser(body.Assignee)
	}

	return v
}

// NewUpdatePayload builds a tasks service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *tasks.UpdatePayload {
	v := &tasks.UpdatePayload{}
	v.Task = unmarshalStoredTaskRequestBodyToTasksStoredTask(body.Task)
	v.ID = id

	return v
}

// NewRemovePayload builds a tasks service remove endpoint payload.
func NewRemovePayload(id string) *tasks.RemovePayload {
	v := &tasks.RemovePayload{}
	v.ID = id

	return v
}

// NewStatusPayload builds a tasks service status endpoint payload.
func NewStatusPayload(body *StatusRequestBody) *tasks.StatusPayload {
	v := &tasks.StatusPayload{
		ID:     *body.ID,
		Status: *body.Status,
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.CreatedDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_date", "body"))
	}
	if body.UpdatedDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_date", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 200, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 5000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 5000, false))
		}
	}
	if body.CreatedDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_date", *body.CreatedDate, goa.FormatDateTime))
	}
	if body.UpdatedDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated_date", *body.UpdatedDate, goa.FormatDateTime))
	}
	if body.DueDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.due_date", *body.DueDate, goa.FormatDateTime))
	}
	if body.Status != nil {
		if !(*body.Status == "Open" || *body.Status == "Closed" || *body.Status == "Pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Open", "Closed", "Pending"}))
		}
	}
	if body.Owner != nil {
		if err2 := ValidateStoredUserRequestBody(body.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Assignee != nil {
		if err2 := ValidateStoredUserRequestBody(body.Assignee); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Task == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("task", "body"))
	}
	if body.Task != nil {
		if err2 := ValidateStoredTaskRequestBody(body.Task); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStatusRequestBody runs the validations defined on StatusRequestBody
func ValidateStatusRequestBody(body *StatusRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "Open" || *body.Status == "Closed" || *body.Status == "Pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Open", "Closed", "Pending"}))
		}
	}
	return
}

// ValidateStoredUserRequestBody runs the validations defined on
// StoredUserRequestBody
func ValidateStoredUserRequestBody(body *StoredUserRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Firstname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstname", "body"))
	}
	if body.Lastname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastname", "body"))
	}
	if body.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.email", *body.Email, ".+@.+\\..{1,6}"))
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 100, false))
		}
	}
	if body.Lastname != nil {
		if utf8.RuneCountInString(*body.Lastname) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", *body.Lastname, utf8.RuneCountInString(*body.Lastname), 100, false))
		}
	}
	if body.Role != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.role", *body.Role, "[a-z]+[a-z0-9]*"))
	}
	return
}

// ValidateStoredTaskRequestBody runs the validations defined on
// StoredTaskRequestBody
func ValidateStoredTaskRequestBody(body *StoredTaskRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.CreatedDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_date", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Owner == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("owner", "body"))
	}
	if body.UpdatedDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_date", "body"))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 200, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 5000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 5000, false))
		}
	}
	if body.CreatedDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_date", *body.CreatedDate, goa.FormatDateTime))
	}
	if body.UpdatedDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated_date", *body.UpdatedDate, goa.FormatDateTime))
	}
	if body.DueDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.due_date", *body.DueDate, goa.FormatDateTime))
	}
	if body.Status != nil {
		if !(*body.Status == "Open" || *body.Status == "Closed" || *body.Status == "Pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Open", "Closed", "Pending"}))
		}
	}
	if body.Owner != nil {
		if err2 := ValidateStoredUserRequestBody(body.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Assignee != nil {
		if err2 := ValidateStoredUserRequestBody(body.Assignee); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
