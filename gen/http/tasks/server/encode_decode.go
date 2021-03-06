// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks HTTP server encoders and decoders
//
// Command:
// $ goa gen tasks/design

package server

import (
	"context"
	"io"
	"net/http"
	tasks "tasks/gen/tasks"
	tasksviews "tasks/gen/tasks/views"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the tasks
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(tasksviews.StoredTaskCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewStoredTaskResponseCollection(res.Projected)
		case "tiny":
			body = NewStoredTaskResponseTinyCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the tasks list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			view *string
			err  error
		)
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(view)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the tasks
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*tasksviews.StoredTask)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "tiny":
			body = NewShowResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the tasks show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id, view)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show tasks
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*tasks.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddResponse returns an encoder for responses returned by the tasks add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the tasks add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddTask(&body)

		return payload, nil
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the tasks
// update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the tasks update
// endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewUpdatePayload(&body, id)

		return payload, nil
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the tasks
// remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the tasks remove
// endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewRemovePayload(id)

		return payload, nil
	}
}

// EncodeStatusResponse returns an encoder for responses returned by the tasks
// status endpoint.
func EncodeStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeStatusRequest returns a decoder for requests sent to the tasks status
// endpoint.
func DecodeStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StatusRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStatusRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStatusPayload(&body)

		return payload, nil
	}
}

// marshalTasksviewsStoredTaskViewToStoredTaskResponse builds a value of type
// *StoredTaskResponse from a value of type *tasksviews.StoredTaskView.
func marshalTasksviewsStoredTaskViewToStoredTaskResponse(v *tasksviews.StoredTaskView) *StoredTaskResponse {
	res := &StoredTaskResponse{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: *v.Description,
		CreatedDate: *v.CreatedDate,
		UpdatedDate: *v.UpdatedDate,
		DueDate:     v.DueDate,
		Status:      *v.Status,
	}
	if v.Owner != nil {
		res.Owner = marshalTasksviewsStoredUserViewToStoredUserResponseTiny(v.Owner)
	}
	if v.Assignee != nil {
		res.Assignee = marshalTasksviewsStoredUserViewToStoredUserResponseTiny(v.Assignee)
	}

	return res
}

// marshalTasksviewsStoredUserViewToStoredUserResponseTiny builds a value of
// type *StoredUserResponseTiny from a value of type *tasksviews.StoredUserView.
func marshalTasksviewsStoredUserViewToStoredUserResponseTiny(v *tasksviews.StoredUserView) *StoredUserResponseTiny {
	res := &StoredUserResponseTiny{
		Email: *v.Email,
	}

	return res
}

// marshalTasksviewsStoredTaskViewToStoredTaskResponseTiny builds a value of
// type *StoredTaskResponseTiny from a value of type *tasksviews.StoredTaskView.
func marshalTasksviewsStoredTaskViewToStoredTaskResponseTiny(v *tasksviews.StoredTaskView) *StoredTaskResponseTiny {
	res := &StoredTaskResponseTiny{
		ID:     *v.ID,
		Title:  *v.Title,
		Status: *v.Status,
	}
	if v.Assignee != nil {
		res.Assignee = marshalTasksviewsStoredUserViewToStoredUserResponseTiny(v.Assignee)
	}

	return res
}

// marshalTasksviewsStoredUserViewToStoredUserResponseBodyTiny builds a value
// of type *StoredUserResponseBodyTiny from a value of type
// *tasksviews.StoredUserView.
func marshalTasksviewsStoredUserViewToStoredUserResponseBodyTiny(v *tasksviews.StoredUserView) *StoredUserResponseBodyTiny {
	res := &StoredUserResponseBodyTiny{
		Email: *v.Email,
	}

	return res
}

// unmarshalStoredUserRequestBodyToTasksStoredUser builds a value of type
// *tasks.StoredUser from a value of type *StoredUserRequestBody.
func unmarshalStoredUserRequestBodyToTasksStoredUser(v *StoredUserRequestBody) *tasks.StoredUser {
	if v == nil {
		return nil
	}
	res := &tasks.StoredUser{
		Email:     *v.Email,
		Firstname: *v.Firstname,
		Lastname:  *v.Lastname,
		Role:      *v.Role,
	}
	if v.Isactive != nil {
		res.Isactive = *v.Isactive
	}
	if v.Isactive == nil {
		res.Isactive = true
	}

	return res
}

// unmarshalStoredTaskRequestBodyToTasksStoredTask builds a value of type
// *tasks.StoredTask from a value of type *StoredTaskRequestBody.
func unmarshalStoredTaskRequestBodyToTasksStoredTask(v *StoredTaskRequestBody) *tasks.StoredTask {
	res := &tasks.StoredTask{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: *v.Description,
		CreatedDate: *v.CreatedDate,
		UpdatedDate: *v.UpdatedDate,
		DueDate:     v.DueDate,
		Status:      *v.Status,
	}
	res.Owner = unmarshalStoredUserRequestBodyToTasksStoredUser(v.Owner)
	if v.Assignee != nil {
		res.Assignee = unmarshalStoredUserRequestBodyToTasksStoredUser(v.Assignee)
	}

	return res
}
