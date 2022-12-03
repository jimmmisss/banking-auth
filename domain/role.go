package domain

import "strings"

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (r RolePermissions) IsAuthorizedFor(role, routeName string) bool {
	perms := r.rolePermissions[role]
	for _, r := range perms {
		if r == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}
