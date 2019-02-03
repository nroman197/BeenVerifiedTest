-- These are the required queries for each function
-- Just to help to understand the code

-- get songs by title
SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length
FROM Songs AS S
INNER JOIN Genres AS G
ON G.ID = S.genre
WHERE S.song LIKE '%"+ title + "%'
ORDER BY S.song ASC;

-- get songs by artist
SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length
FROM Songs AS S
INNER JOIN Genres AS G
ON G.ID = S.genre
WHERE S.artist LIKE '%"+ artist + "%'
ORDER BY S.artist ASC;

-- get songs by genre
SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length
FROM Songs AS S
INNER JOIN Genres AS G
ON G.ID = S.genre
WHERE G.name LIKE '%"+ genre + "%'
ORDER BY G.name ASC;

-- get songs by length (min and max)
SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length
FROM Songs AS S
INNER JOIN Genres AS G ON G.ID = S.genre
WHERE S.length BETWEEN  minLength  AND  maxLength
ORDER BY S.length ASC;

-- get number of songs and total length by genre
SELECT G.name AS Genre, COUNT(1) AS NumberofSongs, SUM(S.length) AS TotalLength
FROM Songs AS S
INNER JOIN Genres AS G
ON G.ID = S.genre
GROUP BY G.name
ORDER BY G.name;
