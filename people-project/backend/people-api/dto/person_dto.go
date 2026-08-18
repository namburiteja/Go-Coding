package dto

import (
	"time"

	"people-api/db/generated"
)

type PersonDTO struct {
	PlayerID     string     `json:"playerID"`
	BirthYear    *int16     `json:"birthYear,omitempty"`
	BirthMonth   *int16     `json:"birthMonth,omitempty"`
	BirthDay     *int16     `json:"birthDay,omitempty"`
	BirthCountry *string    `json:"birthCountry,omitempty"`
	BirthState   *string    `json:"birthState,omitempty"`
	BirthCity    *string    `json:"birthCity,omitempty"`
	DeathYear    *int16     `json:"deathYear,omitempty"`
	DeathMonth   *int16     `json:"deathMonth,omitempty"`
	DeathDay     *int16     `json:"deathDay,omitempty"`
	DeathCountry *string    `json:"deathCountry,omitempty"`
	DeathState   *string    `json:"deathState,omitempty"`
	DeathCity    *string    `json:"deathCity,omitempty"`
	NameFirst    *string    `json:"nameFirst,omitempty"`
	NameLast     *string    `json:"nameLast,omitempty"`
	NameGiven    *string    `json:"nameGiven,omitempty"`
	Weight       *int16     `json:"weight,omitempty"`
	Height       *int16     `json:"height,omitempty"`
	Bats         *string    `json:"bats,omitempty"`
	Throws       *string    `json:"throws,omitempty"`
	Debut        *time.Time `json:"debut,omitempty"`
	FinalGame    *time.Time `json:"finalGame,omitempty"`
	RetroID      *string    `json:"retroID,omitempty"`
	BBRefID      *string    `json:"bbrefID,omitempty"`
}

func ToPersonDTO(p generated.Person) PersonDTO {
	var birthYear *int16
	if p.Birthyear.Valid {
		birthYear = &p.Birthyear.Int16
	}

	var birthMonth *int16
	if p.Birthmonth.Valid {
		birthMonth = &p.Birthmonth.Int16
	}

	var birthDay *int16
	if p.Birthday.Valid {
		birthDay = &p.Birthday.Int16
	}

	var birthCountry *string
	if p.Birthcountry.Valid {
		birthCountry = &p.Birthcountry.String
	}

	var birthState *string
	if p.Birthstate.Valid {
		birthState = &p.Birthstate.String
	}

	var birthCity *string
	if p.Birthcity.Valid {
		birthCity = &p.Birthcity.String
	}

	var deathYear *int16
	if p.Deathyear.Valid {
		deathYear = &p.Deathyear.Int16
	}

	var deathMonth *int16
	if p.Deathmonth.Valid {
		deathMonth = &p.Deathmonth.Int16
	}

	var deathDay *int16
	if p.Deathday.Valid {
		deathDay = &p.Deathday.Int16
	}

	var deathCountry *string
	if p.Deathcountry.Valid {
		deathCountry = &p.Deathcountry.String
	}

	var deathState *string
	if p.Deathstate.Valid {
		deathState = &p.Deathstate.String
	}

	var deathCity *string
	if p.Deathcity.Valid {
		deathCity = &p.Deathcity.String
	}

	var nameFirst *string
	if p.Namefirst.Valid {
		nameFirst = &p.Namefirst.String
	}

	var nameLast *string
	if p.Namelast.Valid {
		nameLast = &p.Namelast.String
	}

	var nameGiven *string
	if p.Namegiven.Valid {
		nameGiven = &p.Namegiven.String
	}

	var weight *int16
	if p.Weight.Valid {
		weight = &p.Weight.Int16
	}

	var height *int16
	if p.Height.Valid {
		height = &p.Height.Int16
	}

	var bats *string
	if p.Bats.Valid {
		bats = &p.Bats.String
	}

	var throws *string
	if p.Throws.Valid {
		throws = &p.Throws.String
	}

	var debut *time.Time
	if p.Debut.Valid {
		debut = &p.Debut.Time
	}

	var finalGame *time.Time
	if p.Finalgame.Valid {
		finalGame = &p.Finalgame.Time
	}

	var retroID *string
	if p.Retroid.Valid {
		retroID = &p.Retroid.String
	}

	var bbrefID *string
	if p.Bbrefid.Valid {
		bbrefID = &p.Bbrefid.String
	}

	return PersonDTO{
		PlayerID:     p.Playerid,
		BirthYear:    birthYear,
		BirthMonth:   birthMonth,
		BirthDay:     birthDay,
		BirthCountry: birthCountry,
		BirthState:   birthState,
		BirthCity:    birthCity,
		DeathYear:    deathYear,
		DeathMonth:   deathMonth,
		DeathDay:     deathDay,
		DeathCountry: deathCountry,
		DeathState:   deathState,
		DeathCity:    deathCity,
		NameFirst:    nameFirst,
		NameLast:     nameLast,
		NameGiven:    nameGiven,
		Weight:       weight,
		Height:       height,
		Bats:         bats,
		Throws:       throws,
		Debut:        debut,
		FinalGame:    finalGame,
		RetroID:      retroID,
		BBRefID:      bbrefID,
	}
}
