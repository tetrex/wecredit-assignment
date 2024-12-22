-- name: CreateNewOtp :exec
INSERT INTO mobile_otp (user_id, otp, valid_till)
VALUES ($1, $2, NOW() + INTERVAL '5 minutes');

-- name: CheckOtp :exec
SELECT COUNT(*) > 0 AS is_valid
FROM mobile_otp
WHERE user_id = $1
  AND otp = $2
  AND valid_till > NOW();
