// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks gRPC client types
//
// Command:
// $ goa gen tasks/design

package client

import (
	taskspb "tasks/gen/grpc/tasks/pb"
	tasks "tasks/gen/tasks"
	tasksviews "tasks/gen/tasks/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "tasks" service.
func NewListRequest() *taskspb.ListRequest {
	message := &taskspb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the "tasks"
// service from the gRPC response type.
func NewListResult(message *taskspb.StoredTaskCollection) tasksviews.StoredTaskCollectionView {
	result := make([]*tasksviews.StoredTaskView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &tasksviews.StoredTaskView{
			ID:          &val.Id,
			Title:       &val.Title,
			Description: &val.Description,
			CreatedDate: &val.CreatedDate,
			UpdatedDate: &val.UpdatedDate,
			Status:      &val.Status,
		}
		if val.DueDate != "" {
			result[i].DueDate = &val.DueDate
		}
		if val.Owner != nil {
			result[i].Owner = protobufTaskspbStoredUserToTasksviewsStoredUserView(val.Owner)
		}
		if val.Assignee != nil {
			result[i].Assignee = protobufTaskspbStoredUserToTasksviewsStoredUserView(val.Assignee)
		}
	}
	return result
}

// NewShowRequest builds the gRPC request type from the payload of the "show"
// endpoint of the "tasks" service.
func NewShowRequest(payload *tasks.ShowPayload) *taskspb.ShowRequest {
	message := &taskspb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the "tasks"
// service from the gRPC response type.
func NewShowResult(message *taskspb.ShowResponse) *tasksviews.StoredTaskView {
	result := &tasksviews.StoredTaskView{
		ID:          &message.Id,
		Title:       &message.Title,
		Description: &message.Description,
		CreatedDate: &message.CreatedDate,
		UpdatedDate: &message.UpdatedDate,
		Status:      &message.Status,
	}
	if message.DueDate != "" {
		result.DueDate = &message.DueDate
	}
	if message.Owner != nil {
		result.Owner = protobufTaskspbStoredUserToTasksviewsStoredUserView(message.Owner)
	}
	if message.Assignee != nil {
		result.Assignee = protobufTaskspbStoredUserToTasksviewsStoredUserView(message.Assignee)
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "tasks" service from the gRPC error response type.
func NewShowNotFoundError(message *taskspb.ShowNotFoundError) *tasks.NotFound {
	er := &tasks.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "tasks" service.
func NewAddRequest(payload *tasks.Task) *taskspb.AddRequest {
	message := &taskspb.AddRequest{
		Title:       payload.Title,
		Description: payload.Description,
		CreatedDate: payload.CreatedDate,
		UpdatedDate: payload.UpdatedDate,
		Status:      payload.Status,
	}
	if payload.DueDate != nil {
		message.DueDate = *payload.DueDate
	}
	if payload.Owner != nil {
		message.Owner = svcTasksStoredUserToTaskspbStoredUser(payload.Owner)
	}
	if payload.Assignee != nil {
		message.Assignee = svcTasksStoredUserToTaskspbStoredUser(payload.Assignee)
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "tasks"
// service from the gRPC response type.
func NewAddResult(message *taskspb.AddResponse) string {
	result := message.Field
	return result
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "tasks" service.
func NewUpdateRequest(payload *tasks.UpdatePayload) *taskspb.UpdateRequest {
	message := &taskspb.UpdateRequest{
		Id: payload.ID,
	}
	if payload.Task != nil {
		message.Task = svcTasksStoredTaskToTaskspbStoredTask(payload.Task)
	}
	return message
}

// NewUpdateResult builds the result type of the "update" endpoint of the
// "tasks" service from the gRPC response type.
func NewUpdateResult(message *taskspb.UpdateResponse) string {
	result := message.Field
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "tasks" service.
func NewRemoveRequest(payload *tasks.RemovePayload) *taskspb.RemoveRequest {
	message := &taskspb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewStatusRequest builds the gRPC request type from the payload of the
// "status" endpoint of the "tasks" service.
func NewStatusRequest(payload *tasks.StatusPayload) *taskspb.StatusRequest {
	message := &taskspb.StatusRequest{
		Id:     payload.ID,
		Status: payload.Status,
	}
	return message
}

// ValidateStoredTaskCollection runs the validations defined on
// StoredTaskCollection.
func ValidateStoredTaskCollection(message *taskspb.StoredTaskCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredTask(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
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

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *taskspb.ShowResponse) (err error) {
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
