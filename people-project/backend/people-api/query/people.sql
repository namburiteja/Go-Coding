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


-- name: GetPeopleCursorFirst :many
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
LIMIT ?;


-- name: GetPeopleCursorAfter :many
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
WHERE playerID > ?
ORDER BY playerID
LIMIT ?;


-- name: GetPeopleTokenFirst :many
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
LIMIT ?;


-- name: GetPeopleTokenAfter :many
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
WHERE playerID > ?
ORDER BY playerID
LIMIT ?;

-- name: UpdatePerson :exec
UPDATE people
SET
    birthYear = COALESCE(?, birthYear),
    birthMonth = COALESCE(?, birthMonth),
    birthDay = COALESCE(?, birthDay),
    birthCountry = COALESCE(?, birthCountry),
    birthState = COALESCE(?, birthState),
    birthCity = COALESCE(?, birthCity),
    deathYear = COALESCE(?, deathYear),
    deathMonth = COALESCE(?, deathMonth),
    deathDay = COALESCE(?, deathDay),
    deathCountry = COALESCE(?, deathCountry),
    deathState = COALESCE(?, deathState),
    deathCity = COALESCE(?, deathCity),
    nameFirst = COALESCE(?, nameFirst),
    nameLast = COALESCE(?, nameLast),
    nameGiven = COALESCE(?, nameGiven),
    weight = COALESCE(?, weight),
    height = COALESCE(?, height),
    bats = COALESCE(?, bats),
    throws = COALESCE(?, throws),
    debut = COALESCE(?, debut),
    finalGame = COALESCE(?, finalGame),
    retroID = COALESCE(?, retroID),
    bbrefID = COALESCE(?, bbrefID)
WHERE playerID = ?;


-- name: GetPeoplePaginatedSorted :many
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
ORDER BY
    CASE
        WHEN sqlc.arg(sortBy) = 'nameFirst'
             AND sqlc.arg(sortOrder) = 'asc'
        THEN nameFirst
    END ASC,

    CASE
        WHEN sqlc.arg(sortBy) = 'nameFirst'
             AND sqlc.arg(sortOrder) = 'desc'
        THEN nameFirst
    END DESC,

    CASE
        WHEN sqlc.arg(sortBy) = 'birthYear'
             AND sqlc.arg(sortOrder) = 'asc'
        THEN birthYear
    END ASC,

    CASE
        WHEN sqlc.arg(sortBy) = 'birthYear'
             AND sqlc.arg(sortOrder) = 'desc'
        THEN birthYear
    END DESC,

    CASE
        WHEN sqlc.arg(sortBy) = 'height'
             AND sqlc.arg(sortOrder) = 'asc'
        THEN height
    END ASC,

    CASE
        WHEN sqlc.arg(sortBy) = 'height'
             AND sqlc.arg(sortOrder) = 'desc'
        THEN height
    END DESC,

    playerID ASC
LIMIT ? OFFSET ?;