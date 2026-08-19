-- name: GetPeoplePaginated :many
SELECT
    playerID,
    birthYear,
    birthMonth,
    birthDay,
    birthCountry,
    birthState,
    birthCity,
    deathYear,
    deathMonth,
    deathDay,
    deathCountry,
    deathState,
    deathCity,
    nameFirst,
    nameLast,
    nameGiven,
    weight,
    height,
    bats,
    throws,
    debut,
    finalGame,
    retroID,
    bbrefID
FROM people
ORDER BY playerID
LIMIT ? OFFSET ?;


-- name: GetPersonByID :one
SELECT
    playerID,
    birthYear,
    birthMonth,
    birthDay,
    birthCountry,
    birthState,
    birthCity,
    deathYear,
    deathMonth,
    deathDay,
    deathCountry,
    deathState,
    deathCity,
    nameFirst,
    nameLast,
    nameGiven,
    weight,
    height,
    bats,
    throws,
    debut,
    finalGame,
    retroID,
    bbrefID
FROM people
WHERE playerID = ?;



-- name: CountPeople :one
SELECT COUNT(*)
FROM people;


-- name: SearchPeopleByName :many
SELECT *
FROM people
WHERE CONCAT(nameFirst, ' ', nameLast) LIKE CONCAT('%', ?, '%')
ORDER BY nameFirst, nameLast
LIMIT 20;

