package mock

import (
	"gwiapi/app/helpers"
	"gwiapi/app/models"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) error {
	// Participants Seeding
	err := seedParticipants(db)
	if err != nil {
		return err
	}

	// Users Seeding
	err = seedUsers(db)
	if err != nil {
		return err
	}

	// Assets Seeding
	err = seedAssets(db)
	if err != nil {
		return err
	}
	return nil
}

func seedParticipants(db *gorm.DB) error {

	//var err error
	numberOfParticipants := 5000

	// hoursSpentOnSocialMedia
	hoursMin := 0.0
	hoursMax := 12.0
	hoursPeak := 2.0

	// Age Range
	ageMin := 15
	ageMax := 85

	// Countries
	country := models.Country{}
	countries := country.GetAllCountries()
	countriesLen := len(countries)

	participants := []models.Participant{}
	for i := 1; i <= numberOfParticipants; i++ {
		participants = append(participants, models.Participant{
			Gender:                  helpers.StringTernary(helpers.GenerateRandomBool() == true, "m", "f"),
			HoursSpentOnSocialDaily: uint8(helpers.GenerateNormalDistribution(hoursPeak, hoursMin, hoursMax)),
			Age:                     uint8(rand.Intn(ageMax-ageMin) + ageMin),
			CountryCode:             countries[rand.Intn(countriesLen)].Alpha2Code,
			CreatedAt:               time.Now(),
		})
	}

	var participants2 []interface{}
	for i := 1; i <= numberOfParticipants; i++ {
		participants2 = append(participants2, models.Participant{
			Gender:                  helpers.StringTernary(helpers.GenerateRandomBool() == true, "m", "f"),
			HoursSpentOnSocialDaily: uint8(helpers.GenerateNormalDistribution(hoursPeak, hoursMin, hoursMax)),
			Age:                     uint8(rand.Intn(ageMax-ageMin) + ageMin),
			CountryCode:             countries[rand.Intn(countriesLen)].Alpha2Code,
			CreatedAt:               time.Now(),
		})
	}

	// tStart := time.Now()
	// for _, p := range participants {
	// 	err := db.Model(&models.Participant{}).Save(&p).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// fmt.Println("Method 1 ", time.Since(tStart))

	// tStart = time.Now()
	// var wg sync.WaitGroup
	// for _, p := range participants {
	// 	wg.Add(1)
	// 	go func(p models.Participant, db *gorm.DB) {
	// 		defer wg.Done()
	// 		_ = db.Model(&models.Participant{}).Save(&p).Error
	// 	}(p, db)
	// }
	// wg.Wait()
	// fmt.Println("Method 2 ", time.Since(tStart))

	//tStart = time.Now()
	dbGorm := helpers.GormStruct{DB: db}
	dbGorm.BatchInsert(participants2)
	// fmt.Println("Method 3 ", time.Since(tStart))

	return nil
}

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Email:    "user1@example.com",
			Password: "secretpassword1",
		},
		{
			Email:    "user2@example.com",
			Password: "secretpassword2",
		},
	}
	for _, u := range users {
		u.Prepare()
		u.Validate("")
		err := db.Model(&models.User{}).Create(&u).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedAssets(db *gorm.DB) error {

	charts := []models.Chart{
		models.Chart{
			GroupedMetric: "age",
		},
		models.Chart{
			GroupedMetric: "country_code",
		},
	}
	for _, c := range charts {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := c.CreateAssetChart(db, uid)
		if err != nil {
			return err
		}
	}

	insights := []models.Insight{
		models.Insight{
			Gender:          "f",
			CountryCode:     "",
			AgeFrom:         20,
			AgeTo:           25,
			HoursComparator: ">",
			HoursReference:  1,
		},
		models.Insight{
			Gender:          "m",
			CountryCode:     "US",
			AgeFrom:         30,
			AgeTo:           35,
			HoursComparator: "<=",
			HoursReference:  1,
		},
		models.Insight{
			Gender:          "f",
			CountryCode:     "US",
			HoursComparator: "<=",
			HoursReference:  1,
		},
		models.Insight{
			Gender:          "f",
			CountryCode:     "",
			AgeFrom:         20,
			HoursComparator: ">",
			HoursReference:  4,
		},
		models.Insight{
			Gender:          "m",
			CountryCode:     "",
			HoursComparator: ">",
			HoursReference:  10,
		},
		models.Insight{
			Gender:          "",
			CountryCode:     "GR",
			AgeTo:           50,
			HoursComparator: ">",
			HoursReference:  1,
		},
	}
	for _, i := range insights {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := i.CreateAssetInsight(db, uid)
		if err != nil {
			return err
		}
	}

	audiences := []models.Audience{
		models.Audience{
			Gender:      "m",
			CountryCode: "",
			AgeFrom:     20,
			AgeTo:       25,
		},
		models.Audience{
			Gender:      "f",
			CountryCode: "",
			AgeFrom:     20,
			AgeTo:       25,
		},
		models.Audience{
			Gender:      "",
			CountryCode: "GR",
			AgeFrom:     25,
			AgeTo:       30,
		},
		models.Audience{
			Gender:      "",
			CountryCode: "US",
			AgeFrom:     25,
			AgeTo:       30,
		},
		models.Audience{
			Gender:      "",
			CountryCode: "GR",
			AgeFrom:     30,
			AgeTo:       35,
		},
		models.Audience{
			Gender:      "f",
			CountryCode: "GR",
			AgeFrom:     30,
			AgeTo:       35,
		},
		models.Audience{
			AgeFrom:     30,
			AgeTo:       40,
			CountryCode: "US",
			Gender:      "",
		},
	}
	for _, i := range audiences {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := i.CreateAssetAudience(db, uid)
		if err != nil {
			return err
		}
	}

	return nil

}
