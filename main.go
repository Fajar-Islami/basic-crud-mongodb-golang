package main

import (
	"fmt"
	"time"

	"CRUDMongoDB/config"
	"CRUDMongoDB/src/modules/profile/model"
	"CRUDMongoDB/src/modules/profile/repository"
)

func main() {
	fmt.Println("GO mongodb")

	db, err := config.GetMongoDb()

	if err != nil {
		fmt.Println(err.Error())
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfileByID("V1", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "V2"
	p.FirstName = "Ahmad"
	p.LastName = "Fajar"
	p.Email = "ahmad@mail.com"
	p.Password = "1234"
	p.CreatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile save..")
	}

}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "V1"
	p.FirstName = "Fajar"
	p.LastName = "Islami"
	p.Email = "fajarislami@mail.com"
	p.Password = "1234"
	p.CreatedAt = time.Now()

	err := profileRepository.Update("V1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile update..")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {

	err := profileRepository.Delete("V1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted..")
	}
}

func getProfileByID(id string, profileRepository repository.ProfileRepository) {

	profile, err := profileRepository.FindById("V1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile")
		fmt.Println(profile)
	}
}

func getProfiles(profileRepository repository.ProfileRepository) {

	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}
	for _, profile := range profiles {
		fmt.Println(profile)
	}
}
