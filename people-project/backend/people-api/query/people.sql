-- =========================================================
-- BASIC OFFSET PAGINATION
-- =========================================================

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


-- =========================================================
-- GET PERSON BY ID
-- =========================================================

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


-- =========================================================
-- COUNT ALL PEOPLE
-- =========================================================

-- name: CountPeople :one
SELECT COUNT(*)
FROM people;


-- =========================================================
-- SEARCH BY NAME
-- =========================================================

-- name: SearchPeopleByName :many
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
WHERE CONCAT(nameFirst, ' ', nameLast)
      LIKE CONCAT('%', ?, '%')
ORDER BY nameFirst, nameLast
LIMIT 20;


-- =========================================================
-- CURSOR PAGINATION - FIRST PAGE
-- =========================================================

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


-- =========================================================
-- CURSOR PAGINATION - AFTER CURSOR
-- =========================================================

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


-- =========================================================
-- TOKEN PAGINATION - FIRST PAGE
-- =========================================================

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


-- =========================================================
-- TOKEN PAGINATION - AFTER TOKEN
-- =========================================================

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


-- =========================================================
-- UPDATE PERSON
-- =========================================================

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


-- =========================================================
-- FILTERED + SORTED + OFFSET PAGINATION
--
-- Filters:
-- 1. Birth Country
-- 2. Birth Year From
-- 3. Birth Year To
-- 4. Height From
-- 5. Height To
-- 6. Weight From
-- 7. Weight To
-- 8. Bats
-- 9. Throws
--
-- These represent 6 filter categories:
-- Country, Birth Year Range, Height Range,
-- Weight Range, Bats, Throws
-- =========================================================

-- name: GetPeoplePaginatedSortedFiltered :many
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

WHERE
    (
        sqlc.narg(birthCountry) IS NULL
        OR birthCountry = sqlc.narg(birthCountry)
    )

    AND
    (
        sqlc.narg(birthYearFrom) IS NULL
        OR birthYear >= sqlc.narg(birthYearFrom)
    )

    AND
    (
        sqlc.narg(birthYearTo) IS NULL
        OR birthYear <= sqlc.narg(birthYearTo)
    )

    AND
    (
        sqlc.narg(heightFrom) IS NULL
        OR height >= sqlc.narg(heightFrom)
    )

    AND
    (
        sqlc.narg(heightTo) IS NULL
        OR height <= sqlc.narg(heightTo)
    )

    AND
    (
        sqlc.narg(weightFrom) IS NULL
        OR weight >= sqlc.narg(weightFrom)
    )

    AND
    (
        sqlc.narg(weightTo) IS NULL
        OR weight <= sqlc.narg(weightTo)
    )

    AND
    (
        sqlc.narg(bats) IS NULL
        OR bats = sqlc.narg(bats)
    )

    AND
    (
        sqlc.narg(throws) IS NULL
        OR throws = sqlc.narg(throws)
    )

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

    CASE
        WHEN sqlc.arg(sortBy) = 'weight'
             AND sqlc.arg(sortOrder) = 'asc'
        THEN weight
    END ASC,

    CASE
        WHEN sqlc.arg(sortBy) = 'weight'
             AND sqlc.arg(sortOrder) = 'desc'
        THEN weight
    END DESC,

    playerID ASC

LIMIT ? OFFSET ?;


-- =========================================================
-- COUNT FILTERED PEOPLE
--
-- IMPORTANT:
-- This must use EXACTLY the same filters as the
-- GetPeoplePaginatedSortedFiltered query.
-- =========================================================

-- name: CountPeopleFiltered :one
SELECT COUNT(*)
FROM people

WHERE
    (
        sqlc.narg(birthCountry) IS NULL
        OR birthCountry = sqlc.narg(birthCountry)
    )

    AND
    (
        sqlc.narg(birthYearFrom) IS NULL
        OR birthYear >= sqlc.narg(birthYearFrom)
    )

    AND
    (
        sqlc.narg(birthYearTo) IS NULL
        OR birthYear <= sqlc.narg(birthYearTo)
    )

    AND
    (
        sqlc.narg(heightFrom) IS NULL
        OR height >= sqlc.narg(heightFrom)
    )

    AND
    (
        sqlc.narg(heightTo) IS NULL
        OR height <= sqlc.narg(heightTo)
    )

    AND
    (
        sqlc.narg(weightFrom) IS NULL
        OR weight >= sqlc.narg(weightFrom)
    )

    AND
    (
        sqlc.narg(weightTo) IS NULL
        OR weight <= sqlc.narg(weightTo)
    )

    AND
    (
        sqlc.narg(bats) IS NULL
        OR bats = sqlc.narg(bats)
    )

    AND
    (
        sqlc.narg(throws) IS NULL
        OR throws = sqlc.narg(throws)
    );