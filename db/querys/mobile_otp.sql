-- name: CreateNewOtp :exec
INSERT INTO mobile_otp (user_id, otp, valid_till)
VALUES ($1, $2, NOW() + INTERVAL '5 minutes');

-- name: CheckOtp :one
SELECT COUNT(*) > 0 AS is_valid
FROM mobile_otp
WHERE user_id = $1
  AND otp = $2
  AND valid_till > NOW();

-- name: GetValidOtpByMobile :one
SELECT mo.otp,u.username,u.mobile_number
FROM mobile_otp mo
JOIN users u ON mo.user_id = u.id
WHERE u.mobile_number = $1           
  AND mo.valid_till > NOW()      
  AND mo.is_used = FALSE         
LIMIT 1;                        

-- name: IsValidOtp :one
SELECT COUNT(mo.otp) >0 AS is_valid
FROM mobile_otp mo
JOIN users u ON mo.user_id = u.id
WHERE u.mobile_number = $1           
  AND mo.valid_till > NOW()      
  AND mo.is_used = FALSE         
LIMIT 1; 

-- name: MarkOtpUsed :exec
UPDATE mobile_otp
SET is_used = TRUE
WHERE otp = $1                   
  AND user_id = (
      SELECT id FROM users WHERE username = $2
  );
