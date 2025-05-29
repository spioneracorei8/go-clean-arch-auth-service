package constants

const (
	INVALID_REQUEST_BODY = "invalid request body"
)

const (
	USERNAME_OR_PASSWORD_WAS_WRONG  = "username or password was wrong"
	PASSWORD_PATTERN_DOES_NOT_MATCH = "password pattern does not match"

	PASSWORD_REGEX                           = `[A-Za-z0-9!@#$%^&*()_+|~/={}\[\]:;'<>?,.\-"]{8,}`
	USERNAME_PATTERN_DOES_NOT_MATCH          = "username pattern does not match"
	PASSWORD_NOT_MATCH_WITH_CONFIRM_PASSWORD = "new_password not match with confirm_new_password"
	PASSWORD_NOT_MATCH_WITH_ACCOUNT          = "old_password not match with current password"
	PASSWORD_AMOUNT                          = 8
)

const (
	MSG_ERROR_NOT_FOUND_CONTENT      = "content not found"
	MSG_ERROR_INVALID_STATUS         = "invalid status"
	MSG_ERROR_INVALID_TYPE           = "invalid type"
	MSG_ERROR_CANNOT_VERIFY          = "cannot verify"
	MSG_ERROR_ALREADY_VERIFY         = "already verify"
	MSG_ERROR_INVALID_AUTH           = "username or password incorrect"
	MSG_ERROR_USER_INACTIVE          = "user status inactive"
	MSG_ERROR_INVALID_EMAIL          = "invalid email"
	MSG_ERROR_INVALID_OTP            = "invalid otp"
	MSG_ERROR_USER_EXIT              = "user exit"
	MSG_NOT_FOUND_ERROR              = "not found"
	MSG_ERROR_HEARDER_REQUIRED       = "header must be required"
	MSG_ERROR_ID_CARD_NUMBER_LENGTH  = "length of id card number is not 13"
	MSG_ERROR_INVALID_ID_CARD_NUMBER = "id card number is invalid"
	MSG_ERROR_INVALID_PASSWORD       = "invalid password"
	MSG_ERROR_USER_NOT_EXIST         = "user not exist"
	MSG_ERROR_UNKNOWN_USER           = "unknown user"
)

const (
	KEY_ERROR_NOT_FOUND_CONTENT = "content_not_found"
	KEY_SERVICE_ERROR           = "internal_error"
	KEY_CONFLICT_ERROR          = "conflict_error"
	KEY_NOT_FOUND_ERROR         = "not_found_error"
)

const (
	STATUS_STAFF_OFF = "OFF"
)
