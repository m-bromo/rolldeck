-- name: Save :exec
INSERT INTO
    users (id, username, email, password)
VALUES
    ($1, $2, $3, $4);
