package stores

type EmployeeRepository interface{}

type EmployeeStore struct{}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{}
}
