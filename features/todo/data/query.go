package data

import (
	"errors"
	"log"
	"todolist/features/todo"

	"gorm.io/gorm"
)

type todoQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.TodoData {
	return &todoQry{
		db: db,
	}
}

func (tq *todoQry) Create(newTodo todo.Core) (todo.Core, error) {
	tod := CoreToData(newTodo)

	err := tq.db.Create(&tod).Error
	if err != nil {
		log.Println("error creating todo: ", err)
	}

	return ToCore(tod), nil
}
func (tq *todoQry) Update(updTodo todo.Core) (todo.Core, error) {
	t := Todo{}
	tod := CoreToData(updTodo)

	aff := tq.db.Where("todo_id", tod.TodoId).First(&t).RowsAffected
	if aff == 0 {
		log.Println("id not found")
		return todo.Core{}, errors.New("id not found")
	}

	err := tq.db.Updates(&tod).Error
	if err != nil {
		log.Println("error updating todo: ", err)
		return todo.Core{}, err
	}

	return ToCore(tod), nil
}
func (tq *todoQry) Delete(id uint) error {
	tod := Todo{TodoId: id}
	aff := tq.db.Where("todo_id", id).First(&tod).RowsAffected
	if aff == 0 {
		log.Println("id not found")
		return errors.New("id not found")
	}

	err := tq.db.Where("activity_id", id).Delete(&tod).Error
	if err != nil {
		log.Println("error deleting todo: ", err)
		return err
	}

	return nil
}
func (tq *todoQry) GetAll(groupId uint) ([]todo.Core, error) {
	todos := []Todo{}
	if groupId != 0 {
		err := tq.db.Raw("SELECT * FROM todos WHERE activity_group_id = ?", groupId).Scan(&todos).Error
		if err != nil {
			log.Println("error getting todos: ", err)
			return []todo.Core{}, err
		}
	} else {
		err := tq.db.Raw("SELECT * FROM todos").Scan(&todos).Error
		if err != nil {
			log.Println("error getting todos: ", err)
			return []todo.Core{}, err
		}
	}
	coreTodos := []todo.Core{}
	for _, v := range todos {
		coreTodos = append(coreTodos, ToCore(v))
	}

	return coreTodos, nil
}
func (tq *todoQry) GetOne(id uint) (todo.Core, error) {
	tod := Todo{TodoId: id}
	aff := tq.db.First(&tod).RowsAffected
	if aff == 0 {
		log.Println("activity id is not found")
		return todo.Core{}, errors.New("activity id is not found")
	}

	return ToCore(tod), nil
}
