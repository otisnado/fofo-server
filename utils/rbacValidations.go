package utils

import (
	"github.com/gobwas/glob"
	"github.com/otisnado/fofo-server/services"
)

func ValidateRolePermissions(urlPathRequested string, methodRequested string, roleId uint) bool {

	var g glob.Glob

	/* Get policies associated with given roleId, if policies not found return false */
	policies, err := services.GetPoliciesByRoleId(roleId)
	if err != nil {
		return false
	}

	/* Compare if roleId has permissions on resource with the given url path */
	for _, policy := range policies {
		g = glob.MustCompile(policy.Path)
		if g.Match(urlPathRequested) {
			authorizedMethods := ConvertStringToStruct(policy.AuthorizedMethods)
			for _, method := range authorizedMethods {
				g = glob.MustCompile(method)
				if g.Match(methodRequested) {
					return true
				}
			}

		}
	}

	return false
}
