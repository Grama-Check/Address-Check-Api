-- name: getPerson :one
SELECT * FROM persons
WHERE nic = $1
LIMIT 1;

-- name: CreatePerson :one
INSERT INTO persons (
    nic,
    address,
    name
) VALUES (
    $1,
    $2,
    $3)
RETURNING *;