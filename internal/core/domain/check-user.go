package domain

type CheckUserRequest struct {
	Email string
}

type CheckUserResult struct {
	Exists bool
}
