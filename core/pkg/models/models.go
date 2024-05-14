package models

type VerificationMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
