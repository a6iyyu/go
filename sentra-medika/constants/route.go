// Package constants defines route constants for the application.
package constants

const (
	// Auth (Public Routes - dipakai langsung di root)
	Login    = "/auth/login"
	Logout   = "/auth/logout"
	Refresh  = "/auth/refresh"
	Register = "/auth/register"

	// Medical Records (Relative to /medical group)
	Records    = "/records"
	RecordByID = "/records/:id"
	MyRecords  = "/my-records"

	// Users (Relative to /admin group)
	Users    = "/users"
	UserByID = "/users/:id"
)