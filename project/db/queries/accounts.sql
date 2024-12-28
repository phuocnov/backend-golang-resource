-- name: createAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES ( $1, $2, $3 )

RETURNING *;

-- name: getAccount :one
SELECT * 
FROM accounts
WHERE id = $1 
LIMIT 1;
