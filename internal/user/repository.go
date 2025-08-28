package user

import "gorm.io/gorm"

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Create(u *User) error { return r.db.Create(u).Error }
func (r *Repository) GetByID(id uint32) (*User, error) {
	var u User
	return &u, r.db.First(&u, id).Error
}
func (r *Repository) List() ([]User, error)  { var xs []User; return xs, r.db.Find(&xs).Error }
func (r *Repository) Update(u *User) error   { return r.db.Save(u).Error }
func (r *Repository) Delete(id uint32) error { return r.db.Delete(&User{}, id).Error }
