package services

import (
	"errors"
	"todolist/features/todo"
)

type todoSrv struct {
	qry todo.TodoData
}

func New(td todo.TodoData) todo.TodoService {
	return &todoSrv{
		qry: td,
	}
}

func (ts *todoSrv) Create(newTodo todo.Core) (todo.Core, error) {
	if newTodo.Title == "" {
		return todo.Core{}, errors.New("title cannot be null")
	}
	if newTodo.ActivityGroupId == 0 {
		return todo.Core{}, errors.New("activity group id cannot be null")
	}
	newTodo.IsActive = true
	if newTodo.Priority == "" {
		newTodo.Priority = "very-high"
	}

	res, err := ts.qry.Create(newTodo)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}
func (ts *todoSrv) Update(id uint, updTodo todo.Core) (todo.Core, error) {
	// if updTodo.Title == "" || updTodo.IsActive ==  {
	// 	return todo.Core{}, errors.New("title cannot be null")
	// }

	updTodo.TodoId = id

	res, err := ts.qry.Update(updTodo)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}
func (ts *todoSrv) Delete(id uint) error {
	err := ts.qry.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
func (ts *todoSrv) GetAll(groupId uint) ([]todo.Core, error) {
	res, err := ts.qry.GetAll(groupId)
	if err != nil {
		return []todo.Core{}, err
	}

	return res, nil
}
func (ts *todoSrv) GetOne(id uint) (todo.Core, error) {
	res, err := ts.qry.GetOne(id)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}
