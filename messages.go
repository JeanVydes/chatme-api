package main

const (
	RecordedUser 		   = "recorded_user"
)

const (
	BodyError              = "empty_body_or_invalid_json"
	JSONEmptyValuesError   = "empty_values"
	PasswordLengthError    = "password_length"
	InvalidGenderError     = "invalid_gender"
	UsernameEmailUsedError = "username_email_used"
	// Technical errors
	TextToHashError        = "hash_0001"
	SQLExecutionError      = "db_0001"
	SQLQueryError          = "db_0002"
	SQLNoRowsAffectedError = "db_0005"
)
