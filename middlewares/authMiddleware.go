package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/auth"
	"github.com/otisnado/fofo-server/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		/* Validate that a JWT exists in Authorization header, if not return a 401 error message */
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		/* Get claims in JWT, if found an error return a 401 error message */
		jwtClaims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		/* Get roles in JWT provided by user in Authorization header */
		userRole := jwtClaims.Role
		authorized := false

		/* Compare path requested with path in AuthorizedRoles */
		pathRequested := c.Request.URL.Path
		methodRequested := c.Request.Method
		authorized = utils.ValidateRolePermissions(pathRequested, methodRequested, userRole)

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Your user's role is not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
