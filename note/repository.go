package note

import "gorm.io/gorm"

type Repository interface {
	Save(note Note) (Note, error)
	All(userId int) ([]Note, error)
	First(id int) (Note, error)
	Update(note Note) (Note, error)
	Delete(note Note) (Note, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(note Note) (Note, error) {
	err := r.db.Create(&note).Error
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *repository) All(userId int) ([]Note, error) {
	var note []Note
	if userId == 0 {
		err := r.db.Find(&note).Error
		if err != nil {
			return note, err
		}
	} else {
		err := r.db.Where("user_id = ?", userId).Find(&note).Error
		if err != nil {
			return note, err
		}
	}

	return note, nil
}

func (r *repository) First(id int) (Note, error) {
	var note Note
	err := r.db.Where("id = ?", id).Find(&note).Error
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *repository) Update(note Note) (Note, error) {
	err := r.db.Save(&note).Error
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *repository) Delete(note Note) (Note, error) {
	err := r.db.Delete(&note).Error
	if err != nil {
		return note, err
	}

	return note, nil
}
