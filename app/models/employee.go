package models

// Employee model to standardize the data
type employee struct {
	Name   string
	Salary string
	Email  string
	Id     string
}

// Convert the given struct to string array
func (e employee) ToSlice() []string  {
	return []string{
		e.Name,
		e.Salary,
		e.Email,
		e.Id,
	}
}
