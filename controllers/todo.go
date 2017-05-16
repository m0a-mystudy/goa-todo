package controllers

import (
	"github.com/goadesign/goa"
	"github.com/m0a-mystudy/goa-todo/app"
	"github.com/m0a-mystudy/goa-todo/store"
)

// TodoController implements the todo resource.
type TodoController struct {
	*goa.Controller
	db *store.DB
}

func ToTodoMedia(a *store.TodoModel) *app.Todo {
	ret := app.Todo(*a)
	return &ret
}

// NewTodoController creates a todo controller.
func NewTodoController(service *goa.Service, db *store.DB) *TodoController {
	return &TodoController{
		Controller: service.NewController("TodoController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *TodoController) Create(ctx *app.CreateTodoContext) error {
	newTodo := c.db.NewTodo()
	payload := ctx.Payload
	saveTodo := store.TodoModel(*payload)
	saveTodo.ID = newTodo.ID
	c.db.SaveTodo(saveTodo)
	return ctx.Created(ToTodoMedia(&saveTodo))
}

// List runs the list action.
func (c *TodoController) List(ctx *app.ListTodoContext) error {
	todos := c.db.GetTodos()
	res := app.TodoCollection{}
	for _, todo := range todos {
		res = append(res, ToTodoMedia(todo))
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TodoController) Show(ctx *app.ShowTodoContext) error {
	if todo, ok := c.db.GetTodo(ctx.TodoID); ok {
		res := ToTodoMedia(&todo)
		return ctx.OK(res)
	}
	return ctx.NotFound()
}

// Update runs the update action.
func (c *TodoController) Update(ctx *app.UpdateTodoContext) error {
	if todo, ok := c.db.GetTodo(ctx.TodoID); ok {
		payload := ctx.Payload
		if payload.Completed != nil {
			todo.Completed = *payload.Completed
		}
		if payload.Title != nil {
			todo.Title = *payload.Title
		}
		c.db.SaveTodo(todo)
		return ctx.NoContent()
	}
	return ctx.NotFound()
}
