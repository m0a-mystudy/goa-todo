package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Todo API", func() {
	Title("My Todo") // Documentation title
	Description("goa study todo api")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/api")
})

var _ = Resource("todo", func() {
	DefaultMedia(Todo)
	BasePath("/todos")
	Action("list", func() {
		Routing(GET(""))
		Description("Retrieve all todos.")
		Response(OK, CollectionOf(Todo))
		Response(NotFound)
	})
	Action("show", func() {
		Routing(GET("/:todoID"))
		Description("Retrieve todo within id.")
		Params(func() {
			Param("todoID", Integer, "todo ID", func() {
				Minimum(1)
			})
		})
		Response(OK, Todo)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("create", func() {
		Routing(POST(""))
		Description("Create new todo")
		Payload(Todo)
		Response(Created, Todo)
		Response(BadRequest)
	})
	Action("update", func() {
		Routing(PUT("/:todoID"))
		Description("change todo")
		Params(func() {
			Param("todoID", Integer, "todo ID", func() {
				Minimum(1)
			})
		})
		Payload(func() {
			Member("title")
			Member("completed", Boolean)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

var Todo = MediaType("application/vnd.todo+json", func() {
	Description("A Todo")
	Reference(TodoPayload)
	Attributes(func() {
		Attribute("id", Integer, "ID of Todo", func() {
			Example(1)
		})
		Attribute("title")
		Attribute("completed")
		Attribute("created", DateTime, "Date of creation")
		Attribute("modified", DateTime, "Date of last update")
		Required("title", "completed", "created", "modified")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("completed")
		Attribute("created")
		Attribute("modified")
	})
})
var TodoPayload = Type("TodoPayload", func() {
	Attribute("title", func() {
		MinLength(8)
		Example("TodoTitleSample")
	})
	Attribute("completed", Boolean, func() {
		Default(false)
		Example(false)
	})
	Attribute("created", DateTime, func() {
		Default("1978-06-30T10:00:00+09:00")
	})
	Attribute("modified", DateTime, func() {
		Default("1978-06-30T10:00:00+09:00")
	})
})
