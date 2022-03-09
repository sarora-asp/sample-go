package todosvc

import (
	"context"
	"fmt"
	"sample/twirp/rpc/todo"

	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var todos []*todo.Todo

type todoServiceProvicer struct{}

func New() *todoServiceProvicer {
	return new(todoServiceProvicer)
}

func (*todoServiceProvicer) CreateTodo(ctx context.Context, td *todo.Todo) (*todo.Response, error) {

	t := todo.Todo{
		TaskId:    td.TaskId,
		Desc:      td.Desc,
		Title:     td.Title,
		Completed: false,
		CreatedAt: timestamppb.Now(),
	}

	todos = append(todos, &t)
	fmt.Println(todos)
	return &todo.Response{
		Code:    201,
		Success: true,
		Msg:     "Hellyeah! THe code has been created",
		Todo:    &t,
	}, nil
}

func (*todoServiceProvicer) GetTodos(ctx context.Context, req *todo.Empty) (*todo.Response, error) {
	return &todo.Response{
		Success: true,
		Code:    200,
		Todos:   todos,
	}, nil
}

func (*todoServiceProvicer) GetTodo(ctx context.Context, req *todo.Request) (*todo.Response, error) {
	if req.TaskId == 0 {
		return nil, twirp.NewError(twirp.InvalidArgument, "task_id is mandatory!")
	}
	if len(todos) > 0 {
		for _, td := range todos {
			if req.TaskId == td.TaskId {
				return &todo.Response{
					Code:    200,
					Success: true,
					Msg:     "Fetched todos successfully",
					Todo:    td,
				}, nil
			}
		}
	}
	return &todo.Response{
		Code:    400,
		Success: true,
		Msg:     "No todo found",
		Todo:    nil,
	}, nil

}

func (*todoServiceProvicer) DeleteTodo(ctx context.Context, req *todo.Request) (*todo.Response, error) {
	if req.TaskId == 0 {
		return nil, twirp.NewError(twirp.InvalidArgument, "task_id is mandatory!")
	}
	tl := len(todos)
	if tl > 0 {
		for i, td := range todos {
			if req.TaskId == td.TaskId {
				todos = append(todos[:i], todos[i+1:]...)
			}
		}
	}

	if tl == len(todos) {
		return &todo.Response{
			Code:    400,
			Success: false,
			Msg:     "Unable to delete the task",
			Todo:    nil,
		}, nil
	}
	return &todo.Response{
		Code:    200,
		Success: true,
		Msg:     "Deleted successfully",
		Todo:    nil,
	}, nil
}

func (*todoServiceProvicer) UpdateTodo(ctx context.Context, req *todo.Request) (*todo.Response, error) {
	if req.TaskId == 0 {
		return nil, twirp.NewError(twirp.InvalidArgument, "task_id is mandatory!")
	}
	var t *todo.Todo
	for _, td := range todos {
		if td.TaskId == req.TaskId {
			t = td
			if td.Title != req.Title {
				td.TaskId = req.TaskId
			} else if td.Completed != req.Completed {
				td.Completed = req.Completed
			} else if td.Desc != req.Desc {
				td.Desc = req.Desc
			}
		}
	}
	return &todo.Response{
		Code:    200,
		Success: true,
		Msg:     "Updated succesfsfully",
		Todo:    t,
	}, nil
}
