package services

import (
	"errors"
	"todolist/features/activity"
)

type activitySrv struct {
	qry activity.ActivityData
}

func New(ad activity.ActivityData) activity.ActivityService {
	return &activitySrv{
		qry: ad,
	}
}

func (as *activitySrv) Create(newAct activity.Core) (activity.Core, error) {
	if newAct.Title == "" {
		return activity.Core{}, errors.New("title cannot be null")
	}

	res, err := as.qry.Create(newAct)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}

func (as *activitySrv) Update(id uint, updAct activity.Core) (activity.Core, error) {
	if updAct.Title == "" {
		return activity.Core{}, errors.New("title cannot be null")
	}

	updAct.ActivityId = id

	res, err := as.qry.Update(updAct)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}
func (as *activitySrv) Delete(id uint) error {
	err := as.qry.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
func (as *activitySrv) GetAll() ([]activity.Core, error) {
	res, err := as.qry.GetAll()
	if err != nil {
		return []activity.Core{}, err
	}

	return res, nil
}
func (as *activitySrv) GetOne(id uint) (activity.Core, error) {
	res, err := as.qry.GetOne(id)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}
