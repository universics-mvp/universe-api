package pkg

type RoleValidator struct{}

func (rv *RoleValidator) Validate(role string, requiredRole string) bool {
	return role == requiredRole
}
