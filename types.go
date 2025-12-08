package main

// ---------------------- Error ----------------------

type InvalidBodyError struct {
	Error string `json:"error" example:"invalid body"`
}

type UserAlreadyExistsError struct {
	Error string `json:"error" example:"user already exists"`
}

type NotFoundError struct {
	Error string `json:"error" example:"not found"`
}

// ---------------------- Success ----------------------

type DeleteSuccess struct {
	Status string `json:"status" example:"deleted"`
}
