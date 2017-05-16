package store

import (
	"sync"

	"github.com/m0a-mystudy/goa-todo/app"
)

// DB 本体
type DB struct {
	sync.Mutex
	maxTodoModelID int
	todos          map[int]*TodoModel
}

// TodoModel Payloadから作成
type TodoModel app.CreateTodoPayload

// CreateTodoPayload
// // TodoModel Todo のモデル
// type TodoModel struct {
// 	ID       int
// 	Title    string
// 	Created  time.Time
// 	Modified time.Time
// }

// NewDB は新規DB作成
func NewDB() *DB {
	return &DB{todos: map[int]*TodoModel{}}
}

// NewTodo は新規作成
func (db *DB) NewTodo() (model TodoModel) {
	db.Lock()
	defer db.Unlock()
	db.maxTodoModelID++
	id := db.maxTodoModelID
	model = TodoModel{ID: &id}
	db.todos[db.maxTodoModelID] = &model
	return
}

// SaveTodo はTodoの保存
func (db *DB) SaveTodo(model TodoModel) {
	db.Lock()
	defer db.Unlock()
	db.todos[*model.ID] = &model
}

// DeleteTodo は削除処理
func (db *DB) DeleteTodo(model TodoModel) {
	db.Lock()
	defer db.Unlock()
	delete(db.todos, *model.ID)
}

// GetTodo は取得処理
func (db *DB) GetTodo(id int) (model TodoModel, ok bool) {
	db.Lock()
	defer db.Unlock()

	todo, found := db.todos[id]
	if !found {
		return
	}
	return *todo, true
}

// GetTodos は全体取得処理
func (db *DB) GetTodos() []*TodoModel {
	db.Lock()
	defer db.Unlock()

	vs := []*TodoModel{}
	for _, v := range db.todos {
		vs = append(vs, v)
	}
	return vs
}
