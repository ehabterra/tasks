package task

import (
	"fmt"
	"reflect"
	"strconv"
	"tasks/gen/tasks"
	"tasks/mocks"
	storage "tasks/pkg/db"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestManager_Status(t *testing.T) {
	taskMock := &mocks.Db{}
	currentID := 0

	taskData := []*tasks.Task{
		{
			Title:       "Add tasks",
			Description: "Add some tasks for test purposes.",
			CreatedDate: "09/16/2020T18:44:00Z",
			UpdatedDate: "09/16/2020T18:44:00Z",
			Status:      "Pending",
		},
		{
			Title:       "Add users",
			Description: "Add some users for test purposes.",
			CreatedDate: "09/11/2020T18:00:00Z",
			UpdatedDate: "09/12/2020T18:00:00Z",
			Status:      "Open",
		},
	}
	var sp []*tasks.StoredTask

	for _, task := range taskData {
		sp = append(sp, &tasks.StoredTask{
			Title:       task.Title,
			Description: task.Description,
			CreatedDate: task.CreatedDate,
			UpdatedDate: task.UpdatedDate,
			Status:      task.Status,
		})
	}

	taskMock.On("Save", mock.AnythingOfType("string"), mock.AnythingOfType("*tasks.StoredTask")).Return(nil)
	taskMock.On("Load", mock.AnythingOfType("string"), &tasks.StoredTask{}).Return(func(id string, res interface{}) error {
		data := res.(*tasks.StoredTask)
		for _, task := range sp {
			if task.ID == id {
				data.Title = task.Title
				data.Description = task.Description
				data.CreatedDate = task.CreatedDate
				data.UpdatedDate = task.UpdatedDate
				data.Status = task.Status
				break
			}
		}
		return nil
	})

	taskMock.On("NewID").Return(func() (string, error) {
		currentID++
		return strconv.Itoa(currentID), nil
	})

	type fields struct {
		Db storage.Db
	}
	type args struct {
		id     string
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Status",
			fields{taskMock},
			args{"1", "Open"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			if err := m.Status(tt.args.id, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Status() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_Add(t *testing.T) {
	taskMock := &mocks.Db{}

	wantDes := "1"
	taskMock.On("NewID").Return(wantDes, nil)

	date := "09/16/2020T18:44:00Z"

	task := &tasks.Task{
		Title:       "Add tasks",
		Description: "Add some tasks for test purposes.",
		CreatedDate: date,
		UpdatedDate: date,
		DueDate:     &date,
		Status:      "Pending",
		Assignee: &tasks.StoredUser{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  true,
		},
		Owner: &tasks.StoredUser{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  true,
		},
	}
	// wantDes, err := taskMock.NewID()

	// if err != nil {
	// 	t.Errorf("Add() error = %v", err)
	// }

	sp := &tasks.StoredTask{
		ID:          wantDes,
		Title:       task.Title,
		Description: task.Description,
		CreatedDate: task.CreatedDate,
		UpdatedDate: task.UpdatedDate,
		DueDate:     task.DueDate,
		Status:      task.Status,
		Assignee:    task.Assignee,
		Owner:       task.Owner,
	}

	taskMock.On("Save", mock.AnythingOfType("string"), sp).Return(nil)

	type fields struct {
		Db storage.Db
	}
	type args struct {
		p *tasks.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Add",
			fields{taskMock},
			args{task},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			id, err := m.Add(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			if id != wantDes {
				t.Errorf("Add() id = %v, wantRes %v", id, wantDes)
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	taskMock := &mocks.Db{}

	taskData := tasks.StoredTaskCollection{
		{
			Title:       "Add tasks",
			Description: "Add some tasks for test purposes.",
			CreatedDate: "09/16/2020T18:44:00Z",
			UpdatedDate: "09/16/2020T18:44:00Z",
			Status:      "Pending",
		},
		{
			Title:       "Add users",
			Description: "Add some users for test purposes.",
			CreatedDate: "09/11/2020T18:00:00Z",
			UpdatedDate: "09/12/2020T18:00:00Z",
			Status:      "Open",
		},
	}
	var res tasks.StoredTaskCollection

	taskMock.On("LoadAll", &res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*tasks.StoredTaskCollection)
		*arg = append(*arg, taskData...)
		fmt.Printf("value %v, type %T \n", arg, arg)
	})

	type fields struct {
		Db storage.Db
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes tasks.StoredTaskCollection
		wantErr bool
	}{
		{
			"List",
			fields{taskMock},
			taskData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			gotRes, err := m.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("List() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestManager_Remove(t *testing.T) {
	taskMock := mocks.Db{}
	email := "ehab@test.com"
	taskMock.On("Delete", email).Return(nil)

	type fields struct {
		Db storage.Db
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Remove",
			fields{&taskMock},
			args{email},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			if err := m.Remove(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_Show(t *testing.T) {
	taskMock := &mocks.Db{}

	taskData := tasks.StoredTask{
		Title:       "Add tasks",
		Description: "Add some tasks for test purposes.",
		CreatedDate: "09/16/2020T18:44:00Z",
		UpdatedDate: "09/16/2020T18:44:00Z",
		Status:      "Pending",
	}

	var res tasks.StoredTask

	taskMock.On("Load", taskData.ID, &res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*tasks.StoredTask)
		*arg = taskData
		fmt.Printf("value %v, type %T \n", arg, arg)
	})

	type fields struct {
		Db storage.Db
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *tasks.StoredTask
		wantErr bool
	}{
		{
			"Show",
			fields{taskMock},
			args{taskData.ID},
			&taskData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			gotRes, err := m.Show(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Show() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestManager_Update(t *testing.T) {
	taskMock := &mocks.Db{}

	wantID := "1"
	date := "09/16/2020T18:44:00Z"

	task := &tasks.Task{
		Title:       "Add tasks",
		Description: "Add some tasks for test purposes.",
		CreatedDate: date,
		UpdatedDate: date,
		DueDate:     &date,
		Status:      "Pending",
		Assignee: &tasks.StoredUser{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  true,
		},
		Owner: &tasks.StoredUser{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  true,
		},
	}

	sp := &tasks.StoredTask{
		ID:          wantID,
		Title:       task.Title,
		Description: task.Description,
		CreatedDate: task.CreatedDate,
		UpdatedDate: task.UpdatedDate,
		DueDate:     task.DueDate,
		Status:      task.Status,
		Assignee:    task.Assignee,
		Owner:       task.Owner,
	}

	taskMock.On("Save", mock.AnythingOfType("string"), sp).Return(nil)

	type fields struct {
		Db storage.Db
	}
	type args struct {
		p *tasks.UpdatePayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Update",
			fields{taskMock},
			args{
				&tasks.UpdatePayload{
					ID:   "1",
					Task: sp,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			if err := m.Update(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewManager(t *testing.T) {
	taskMock := &mocks.Db{}
	type args struct {
		db storage.Db
	}
	tests := []struct {
		name string
		args args
		want *Manager
	}{
		{
			"NewManager",
			args{taskMock},
			&Manager{taskMock},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewManager(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
