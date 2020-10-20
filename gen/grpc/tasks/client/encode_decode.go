// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks gRPC client encoders and decoders
//
// Command:
// $ goa gen tasks/design

package client

import (
	"context"
	taskspb "tasks/gen/grpc/tasks/pb"
	tasks "tasks/gen/tasks"
	tasksviews "tasks/gen/tasks/views"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "tasks" service "list"
// endpoint.
func BuildListFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*taskspb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &taskspb.ListRequest{}, opts...)
	}
}

// EncodeListRequest encodes requests sent to tasks list endpoint.
func EncodeListRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.ListPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "list", "*tasks.ListPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewListRequest(), nil
}

// DecodeListResponse decodes responses from the tasks list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*taskspb.StoredTaskCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "list", "*taskspb.StoredTaskCollection", v)
	}
	res := NewListResult(message)
	vres := tasksviews.StoredTaskCollection{Projected: res, View: view}
	if err := tasksviews.ValidateStoredTaskCollection(vres); err != nil {
		return nil, err
	}
	return tasks.NewStoredTaskCollection(vres), nil
}

// BuildShowFunc builds the remote method to invoke for "tasks" service "show"
// endpoint.
func BuildShowFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Show(ctx, reqpb.(*taskspb.ShowRequest), opts...)
		}
		return grpccli.Show(ctx, &taskspb.ShowRequest{}, opts...)
	}
}

// EncodeShowRequest encodes requests sent to tasks show endpoint.
func EncodeShowRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.ShowPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "show", "*tasks.ShowPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewShowRequest(payload), nil
}

// DecodeShowResponse decodes responses from the tasks show endpoint.
func DecodeShowResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*taskspb.ShowResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "show", "*taskspb.ShowResponse", v)
	}
	res := NewShowResult(message)
	vres := &tasksviews.StoredTask{Projected: res, View: view}
	if err := tasksviews.ValidateStoredTask(vres); err != nil {
		return nil, err
	}
	return tasks.NewStoredTask(vres), nil
}

// BuildAddFunc builds the remote method to invoke for "tasks" service "add"
// endpoint.
func BuildAddFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*taskspb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &taskspb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to tasks add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.Task)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "add", "*tasks.Task", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the tasks add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "add", "*taskspb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildUpdateFunc builds the remote method to invoke for "tasks" service
// "update" endpoint.
func BuildUpdateFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*taskspb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &taskspb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to tasks update endpoint.
func EncodeUpdateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.UpdatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "update", "*tasks.UpdatePayload", v)
	}
	return NewUpdateRequest(payload), nil
}

// DecodeUpdateResponse decodes responses from the tasks update endpoint.
func DecodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.UpdateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "update", "*taskspb.UpdateResponse", v)
	}
	res := NewUpdateResult(message)
	return res, nil
}

// BuildRemoveFunc builds the remote method to invoke for "tasks" service
// "remove" endpoint.
func BuildRemoveFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Remove(ctx, reqpb.(*taskspb.RemoveRequest), opts...)
		}
		return grpccli.Remove(ctx, &taskspb.RemoveRequest{}, opts...)
	}
}

// EncodeRemoveRequest encodes requests sent to tasks remove endpoint.
func EncodeRemoveRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.RemovePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "remove", "*tasks.RemovePayload", v)
	}
	return NewRemoveRequest(payload), nil
}

// BuildStatusFunc builds the remote method to invoke for "tasks" service
// "status" endpoint.
func BuildStatusFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Status(ctx, reqpb.(*taskspb.StatusRequest), opts...)
		}
		return grpccli.Status(ctx, &taskspb.StatusRequest{}, opts...)
	}
}

// EncodeStatusRequest encodes requests sent to tasks status endpoint.
func EncodeStatusRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.StatusPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "status", "*tasks.StatusPayload", v)
	}
	return NewStatusRequest(payload), nil
}