package models

// declare models

// Service represents a service in the organization
type Service struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Versions    []Version `json:"versions"`
}

// new service instance
func NewService() *Service {
	return &Service{}
}

// Services represents a collection of services
type Services []Service

// ServiceVersion represents a version of a service
type Version struct {
	ID        int    `json:"id"`
	ServiceID int    `json:"service_id"`
	Version   string `json:"version"`
}

// new version instance
func NewVersion() *Version {
	return &Version{}
}


