// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks gRPC server types
//
// Command:
// $ goa gen tasks/design

package server

import (
	taskspb "tasks/gen/grpc/tasks/pb"
	tasks "tasks/gen/tasks"
	tasksviews "tasks/gen/tasks/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewListPayload builds the payload of the "list" endpoint of the "tasks"
// service from the gRPC request type.
func NewListPayload(view *string) *tasks.ListPayload {
	v := &tasks.ListPayload{}
	v.View = view
	return v
}

// NewStoredTaskCollection builds the gRPC response type from the result of the
// "list" endpoint of the "tasks" service.
func NewStoredTaskCollection(result tasksviews.StoredTaskCollectionView) *taskspb.StoredTaskCollection {
	message := &taskspb.StoredTaskCollection{}
	message.Field = make([]*taskspb.StoredTask, len(result))
	for i, val := range result {
		message.Field[i] = &taskspb.StoredTask{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Title != nil {
			message.Field[i].Title = *val.Title
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.CreatedDate != nil {
			message.Field[i].CreatedDate = *val.CreatedDate
		}
		if val.UpdatedDate != nil {
			message.Field[i].UpdatedDate = *val.UpdatedDate
		}
		if val.DueDate != nil {
			message.Field[i].DueDate = *val.DueDate
		}
		if val.Status != nil {
			message.Field[i].Status = *val.Status
		}
		if val.Status == nil {
			message.Field[i].Status = "Open"
		}
		if val.Owner != nil {
			message.Field[i].Owner = svcTasksviewsStoredUserViewToTaskspbStoredUser(val.Owner)
		}
		if val.Assignee != nil {
			message.Field[i].Assignee = svcTasksviewsStoredUserViewToTaskspbStoredUser(val.Assignee)
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "tasks"
// service from the gRPC request type.
func NewShowPayload(message *taskspb.ShowRequest, view *string) *tasks.ShowPayload {
	v := &tasks.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewShowResponse builds the gRPC response type from the result of the "show"
// endpoint of the "tasks" service.
func NewShowResponse(result *tasksviews.StoredTaskView) *taskspb.ShowResponse {
	message := &taskspb.ShowResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.Title != nil {
		message.Title = *result.Title
	}
	if result.Description != nil {
		message.Description = *result.Description
	}
	if result.CreatedDate != nil {
		message.CreatedDate = *result.CreatedDate
	}
	if result.UpdatedDate != nil {
		message.UpdatedDate = *result.UpdatedDate
	}
	if result.DueDate != nil {
		message.DueDate = *result.DueDate
	}
	if result.Status != nil {
		message.Status = *result.Status
	}
	if result.Status == nil {
		message.Status = "Open"
	}
	if result.Owner != nil {
		message.Owner = svcTasksviewsStoredUserViewToTaskspbStoredUser(result.Owner)
	}
	if result.Assignee != nil {
		message.Assignee = svcTasksviewsStoredUserViewToTaskspbStoredUser(result.Assignee)
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "tasks" service.
func NewShowNotFoundError(er *tasks.NotFound) *taskspb.ShowNotFoundError {
	message := &taskspb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "tasks"
// service from the gRPC request type.
func NewAddPayload(message *taskspb.AddRequest) *tasks.Task {
	v := &tasks.Task{
		Title:       message.Title,
		Description: message.Description,
		CreatedDate: message.CreatedDate,
		UpdatedDate: message.UpdatedDate,
		Status:      message.Status,
	}
	if message.DueDate != "" {
		v.DueDate = &message.DueDate
	}
	if message.Owner != nil {
		v.Owner = protobufTaskspbStoredUserToTasksStoredUser(message.Owner)
	}
	if message.Assignee != nil {
		v.Assignee = protobufTaskspbStoredUserToTasksStoredUser(message.Assignee)
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "tasks" service.
func NewAddResponse(result string) *taskspb.AddResponse {
	message := &taskspb.AddResponse{}
	message.Field = result
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the "tasks"
// service from the gRPC request type.
func NewUpdatePayload(message *taskspb.UpdateRequest) *tasks.UpdatePayload {
	v := &tasks.UpdatePayload{
		ID: message.Id,
	}
	if message.Task != nil {
		v.Task = protobufTaskspbStoredTaskToTasksStoredTask(message.Task)
	}
	return v
}

// NewUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "tasks" service.
func NewUpdateResponse(result string) *taskspb.UpdateResponse {
	message := &taskspb.UpdateResponse{}
	message.Field = result
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the "tasks"
// service from the gRPC request type.
func NewRemovePayload(message *taskspb.RemoveRequest) *tasks.RemovePayload {
	v := &tasks.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "tasks" service.
func NewRemoveResponse() *taskspb.RemoveResponse {
	message := &taskspb.RemoveResponse{}
	return message
}

// NewStatusPayload builds the payload of the "status" endpoint of the "tasks"
// service from the gRPC request type.
func NewStatusPayload(message *taskspb.StatusRequest) *tasks.StatusPayload {
	v := &tasks.StatusPayload{
		ID:     message.Id,
		Status: message.Status,
	}
	return v
}

// NewStatusResponse builds the gRPC response type from the result of the
// "status" endpoint of the "tasks" service.
func NewStatusResponse() *taskspb.StatusResponse {
	message := &taskspb.StatusResponse{}
	return message
}

// ValidateStoredTask runs the validations defined on StoredTask.
func ValidateStoredTask(message *taskspb.StoredTask) (err error) {
	if message.Owner == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("owner", "message"))
	}
	if utf8.RuneCountInString(message.Title) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.title", message.Title, utf8.RuneCountInString(message.Title), 200, false))
	}
	if utf8.RuneCountInString(message.Description) > 5000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 5000, false))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_date", message.CreatedDate, goa.FormatDateTime))

	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_date", message.UpdatedDate, goa.FormatDateTime))

	err = goa.MergeErrors(err, goa.ValidateFormat("message.due_date", message.DueDate, goa.FormatDateTime))

	if !(message.Status == "Open" || message.Status == "Closed" || message.Status == "Pending") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.status", message.Status, []interface{}{"Open", "Closed", "Pending"}))
	}
	if message.Owner != nil {
		if err2 := ValidateStoredUser(message.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Assignee != nil {
		if err2 := ValidateStoredUser(message.Assignee); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredUser runs the validations defined on StoredUser.
func ValidateStoredUser(message *taskspb.StoredUser) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.email", message.Email, ".+@.+\\..{1,6}"))
	if utf8.RuneCountInString(message.Firstname) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.firstname", message.Firstname, utf8.RuneCountInString(message.Firstname), 100, false))
	}
	if utf8.RuneCountInString(message.Lastname) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.lastname", message.Lastname, utf8.RuneCountInString(message.Lastname), 100, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.role", message.Role, "[a-z]+[a-z0-9]*"))
	return
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *taskspb.AddRequest) (err error) {
	if utf8.RuneCountInString(message.Title) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.title", message.Title, utf8.RuneCountInString(message.Title), 200, false))
	}
	if utf8.RuneCountInString(message.Description) > 5000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 5000, false))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_date", message.CreatedDate, goa.FormatDateTime))

	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_date", message.UpdatedDate, goa.FormatDateTime))

	if message.DueDate != "" {
		err = goa.MergeErrors(err, goa.ValidateFormat("message.due_date", message.DueDate, goa.FormatDateTime))
	}
	if !(message.Status == "Open" || message.Status == "Closed" || message.Status == "Pending") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.status", message.Status, []interface{}{"Open", "Closed", "Pending"}))
	}
	if message.Owner != nil {
		if err2 := ValidateStoredUser(message.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Assignee != nil {
		if err2 := ValidateStoredUser(message.Assignee); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequest runs the validations defined on UpdateRequest.
func ValidateUpdateRequest(message *taskspb.UpdateRequest) (err error) {
	if message.Task == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("task", "message"))
	}
	if message.Task != nil {
		if err2 := ValidateStoredTask(message.Task); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStatusRequest runs the validations defined on StatusRequest.
func ValidateStatusRequest(message *taskspb.StatusRequest) (err error) {
	if !(message.Status == "Open" || message.Status == "Closed" || message.Status == "Pending") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.status", message.Status, []interface{}{"Open", "Closed", "Pending"}))
	}
	return
}

// svcTasksviewsStoredUserViewToTaskspbStoredUser builds a value of type
// *taskspb.StoredUser from a value of type *tasksviews.StoredUserView.
func svcTasksviewsStoredUserViewToTaskspbStoredUser(v *tasksviews.StoredUserView) *taskspb.StoredUser {
	res := &taskspb.StoredUser{}
	if v.Email != nil {
		res.Email = *v.Email
	}
	if v.Firstname != nil {
		res.Firstname = *v.Firstname
	}
	if v.Lastname != nil {
		res.Lastname = *v.Lastname
	}
	if v.Isactive != nil {
		res.Isactive = *v.Isactive
	}
	if v.Role != nil {
		res.Role = *v.Role
	}
	if v.Isactive == nil {
		res.Isactive = true
	}

	return res
}

// protobufTaskspbStoredUserToTasksviewsStoredUserView builds a value of type
// *tasksviews.StoredUserView from a value of type *taskspb.StoredUser.
func protobufTaskspbStoredUserToTasksviewsStoredUserView(v *taskspb.StoredUser) *tasksviews.StoredUserView {
	res := &tasksviews.StoredUserView{
		Email:     &v.Email,
		Firstname: &v.Firstname,
		Lastname:  &v.Lastname,
		Isactive:  &v.Isactive,
		Role:      &v.Role,
	}

	return res
}

// protobufTaskspbStoredUserToTasksStoredUser builds a value of type
// *tasks.StoredUser from a value of type *taskspb.StoredUser.
func protobufTaskspbStoredUserToTasksStoredUser(v *taskspb.StoredUser) *tasks.StoredUser {
	if v == nil {
		return nil
	}
	res := &tasks.StoredUser{
		Email:     v.Email,
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Isactive:  v.Isactive,
		Role:      v.Role,
	}

	return res
}

// svcTasksStoredUserToTaskspbStoredUser builds a value of type
// *taskspb.StoredUser from a value of type *tasks.StoredUser.
func svcTasksStoredUserToTaskspbStoredUser(v *tasks.StoredUser) *taskspb.StoredUser {
	if v == nil {
		return nil
	}
	res := &taskspb.StoredUser{
		Email:     v.Email,
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Isactive:  v.Isactive,
		Role:      v.Role,
	}

	return res
}

// protobufTaskspbStoredTaskToTasksStoredTask builds a value of type
// *tasks.StoredTask from a value of type *taskspb.StoredTask.
func protobufTaskspbStoredTaskToTasksStoredTask(v *taskspb.StoredTask) *tasks.StoredTask {
	res := &tasks.StoredTask{
		ID:          v.Id,
		Title:       v.Title,
		Description: v.Description,
		CreatedDate: v.CreatedDate,
		UpdatedDate: v.UpdatedDate,
		Status:      v.Status,
	}
	if v.DueDate != "" {
		res.DueDate = &v.DueDate
	}
	if v.Owner != nil {
		res.Owner = protobufTaskspbStoredUserToTasksStoredUser(v.Owner)
	}
	if v.Assignee != nil {
		res.Assignee = protobufTaskspbStoredUserToTasksStoredUser(v.Assignee)
	}

	return res
}

// svcTasksStoredTaskToTaskspbStoredTask builds a value of type
// *taskspb.StoredTask from a value of type *tasks.StoredTask.
func svcTasksStoredTaskToTaskspbStoredTask(v *tasks.StoredTask) *taskspb.StoredTask {
	res := &taskspb.StoredTask{
		Id:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		CreatedDate: v.CreatedDate,
		UpdatedDate: v.UpdatedDate,
		Status:      v.Status,
	}
	if v.DueDate != nil {
		res.DueDate = *v.DueDate
	}
	if v.Owner != nil {
		res.Owner = svcTasksStoredUserToTaskspbStoredUser(v.Owner)
	}
	if v.Assignee != nil {
		res.Assignee = svcTasksStoredUserToTaskspbStoredUser(v.Assignee)
	}

	return res
}
