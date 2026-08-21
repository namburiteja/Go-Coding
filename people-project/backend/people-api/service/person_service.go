package service

import (
	"context"

	"people-api/db/generated"
	"people-api/dto"
)

type PersonService struct {
	Queries *generated.Queries
}

func NewPersonService(queries *generated.Queries) *PersonService {
	return &PersonService{
		Queries: queries,
	}
}

func (s *PersonService) GetAllPeople(ctx context.Context) ([]dto.PersonDTO, error) {
	people, err := s.Queries.GetPeoplePaginated(ctx, generated.GetPeoplePaginatedParams{
		Limit:  20,
		Offset: 0,
	})

	if err != nil {
		return nil, err
	}

	result := make([]dto.PersonDTO, 0, len(people))

	for _, person := range people {
		result = append(result, dto.ToPersonDTO(person))
	}

	return result, nil
}

func (s *PersonService) GetPeoplePaginated(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]dto.PersonDTO, error) {

	people, err := s.Queries.GetPeoplePaginated(
		ctx,
		generated.GetPeoplePaginatedParams{
			Limit:  limit,
			Offset: offset,
		},
	)

	if err != nil {
		return nil, err
	}

	result := make([]dto.PersonDTO, 0, len(people))

	for _, person := range people {
		result = append(result, dto.ToPersonDTO(person))
	}

	return result, nil
}

func (s *PersonService) GetPeopleCursor(
	ctx context.Context,
	cursor string,
	limit int32,
) ([]dto.PersonDTO, error) {

	var people []generated.Person
	var err error

	if cursor == "" {
		people, err = s.Queries.GetPeopleCursorFirst(
			ctx,
			limit,
		)
	} else {
		people, err = s.Queries.GetPeopleCursorAfter(
			ctx,
			generated.GetPeopleCursorAfterParams{
				Playerid: cursor,
				Limit:    limit,
			},
		)
	}

	if err != nil {
		return nil, err
	}

	result := make([]dto.PersonDTO, 0, len(people))

	for _, person := range people {
		result = append(
			result,
			dto.ToPersonDTO(person),
		)
	}

	return result, nil
}

func (s *PersonService) GetPeopleByID(
	ctx context.Context,
	playerID string,
) (dto.PersonDTO, error) {

	person, err := s.Queries.GetPersonByID(ctx, playerID)

	if err != nil {
		return dto.PersonDTO{}, err
	}

	return dto.ToPersonDTO(person), nil
}

func (s *PersonService) CountPeople(ctx context.Context) (int64, error) {
	return s.Queries.CountPeople(ctx)
}

func (s *PersonService) SearchPeopleByName(
	ctx context.Context,
	name string,
) ([]dto.PersonDTO, error) {

	people, err := s.Queries.SearchPeopleByName(ctx, name)
	if err != nil {
		return nil, err
	}

	result := make([]dto.PersonDTO, 0, len(people))

	for _, person := range people {
		result = append(result, dto.ToPersonDTO(person))
	}

	return result, nil
}

func (s *PersonService) GetPeopleToken(
	ctx context.Context,
	playerID string,
	limit int32,
) ([]dto.PersonDTO, error) {

	var people []generated.Person
	var err error

	if playerID == "" {

		people, err = s.Queries.GetPeopleTokenFirst(
			ctx,
			limit,
		)

	} else {

		people, err = s.Queries.GetPeopleTokenAfter(
			ctx,
			generated.GetPeopleTokenAfterParams{
				Playerid: playerID,
				Limit:    limit,
			},
		)
	}

	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.PersonDTO,
		0,
		len(people),
	)

	for _, person := range people {

		result = append(
			result,
			dto.ToPersonDTO(person),
		)
	}

	return result, nil
}


func (s *PersonService) UpdatePerson(
    ctx context.Context,
    person generated.UpdatePersonParams,
) error {
    return s.Queries.UpdatePerson(ctx, person)
}


func (s *PersonService) GetPeoplePaginatedSorted(
    ctx context.Context,
    limit int32,
    offset int32,
    sortBy string,
    sortOrder string,
) ([]dto.PersonDTO, error) {

    people, err := s.Queries.GetPeoplePaginatedSorted(
        ctx,
        generated.GetPeoplePaginatedSortedParams{
            Sortby:    sortBy,
            Sortorder: sortOrder,
            Limit:     limit,
            Offset:    offset,
        },
    )

    if err != nil {
        return nil, err
    }

    result := make([]dto.PersonDTO, 0, len(people))

    for _, person := range people {
        result = append(
            result,
            dto.ToPersonDTO(person),
        )
    }

    return result, nil
}