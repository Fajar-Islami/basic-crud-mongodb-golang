package repository

import (
	"CRUDMongoDB/src/modules/profile/model"
)

// ProfileRepository interface
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindById(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
