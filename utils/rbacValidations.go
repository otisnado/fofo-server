package utils

import (
	"github.com/gobwas/glob"
	"github.com/otisnado/nepackage/services"
)

func ValidateRolePermissions(urlPathRequested string, methodRequested string, roles string) bool {

	var g glob.Glob

	userRoles := ConvertStringToUintStruct(roles)
	for _, role := range userRoles {

		/* Get policies associated with given role, if policies not found return false */
		policies, err := services.GetPoliciesByRoleId(uint(role))
		if err != nil {
			return false
		}

		/* Compare if role has permissions on resource with the given url path */
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
	}

	return false
}
