package domain

type CheckExistsRequest struct {
	Table  string
	Column string
	Id     *string
}

type CheckExistsResponse struct {
	Exists bool
}
