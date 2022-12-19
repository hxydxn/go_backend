-- name: GetCreditBySSRN :one
SELECT * FROM credits WHERE ssrn = $1;
