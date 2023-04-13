package data

import (
	"errors"
	"log"
	"todolist/features/activity"

	"gorm.io/gorm"
)

type ActivityQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.ActivityData {
	return &ActivityQry{
		db: db,
	}
}

func (aq *ActivityQry) Create(newAct activity.Core) (activity.Core, error) {
	act := CoreToData(newAct)

	err := aq.db.Create(&act).Error
	if err != nil {
		log.Println("error creating act: ", err)
	}

	return ToCore(act), nil
}

func (aq *ActivityQry) Update(updAct activity.Core) (activity.Core, error) {
	a := Activity{}

	aff := aq.db.Where("activity_id", updAct.ActivityId).First(&a).RowsAffected
	if aff == 0 {
		log.Println("id not found")
		return activity.Core{}, errors.New("id not found")
	}

	a.Title = updAct.Title

	err := aq.db.Updates(&a).Error
	if err != nil {
		log.Println("error updating activity: ", err)
		return activity.Core{}, err
	}

	return ToCore(a), nil
}

func (aq *ActivityQry) Delete(id uint) error {
	act := Activity{ActivityId: id}
	aff := aq.db.Where("activity_id", id).First(&act).RowsAffected
	if aff == 0 {
		log.Println("id not found")
		return errors.New("id not found")
	}

	err := aq.db.Where("activity_id", id).Delete(&act).Error
	if err != nil {
		log.Println("error deleting activity: ", err)
		return err
	}

	return nil
}
func (aq *ActivityQry) GetAll() ([]activity.Core, error) {
	activities := []Activity{}
	err := aq.db.Raw("SELECT * FROM activities").Scan(&activities).Error
	if err != nil {
		log.Println("error getting activities: ", err)
		return []activity.Core{}, err
	}

	coreActs := []activity.Core{}
	for _, v := range activities {
		coreActs = append(coreActs, ToCore(v))
	}

	return coreActs, nil
}
func (aq *ActivityQry) GetOne(id uint) (activity.Core, error) {
	act := Activity{ActivityId: id}
	aff := aq.db.First(&act).RowsAffected
	if aff == 0 {
		log.Println("activity id is not found")
		return activity.Core{}, errors.New("activity id is not found")
	}

	return ToCore(act), nil
}
