// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "Todo API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/m0a-mystudy/goa-todo/design
// --out=$(GOPATH)/src/github.com/m0a-mystudy/goa-todo
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// A Todo (default view)
//
// Identifier: application/vnd.todo+json; view=default
type Todo struct {
	Completed bool `form:"completed" json:"completed" xml:"completed"`
	// Date of creation
	Created time.Time `form:"created" json:"created" xml:"created"`
	// ID of Todo
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Date of last update
	Modified time.Time `form:"modified" json:"modified" xml:"modified"`
	Title    string    `form:"title" json:"title" xml:"title"`
}

// Validate validates the Todo media type instance.
func (mt *Todo) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if utf8.RuneCountInString(mt.Title) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 8, true))
	}
	return
}

// TodoCollection is the media type for an array of Todo (default view)
//
// Identifier: application/vnd.todo+json; type=collection; view=default
type TodoCollection []*Todo

// Validate validates the TodoCollection media type instance.
func (mt TodoCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
