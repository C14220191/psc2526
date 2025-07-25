package models

type Response struct {
	StatusCode int
	Message    string // Success | Created | Deleted | error message general (conflict, fk error)
	Data       any    // inti / data yang diambil dari DB
}
