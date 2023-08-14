package auth

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"time"

	"application/interceptor"
	"application/logger"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func AuthenticationMiddleware() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("APPLICATION_DOMAIN") + "/")
	if err != nil {
		logger.ThrowErrorLog("Failed to parse the issuer url")
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_USER_GET_API")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		logger.ThrowErrorLog("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		logger.ThrowErrorLog("Encountered error while validating JWT")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

func CheckPermission(permissions []string, CheckAllPermissions bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		// these are the permissions which i have specified in the backend 
		// so that any user who wishes to use these has to have these roles attached to him.
		if len(permissions) != 0 {
			if !claims.IsAuthorized(permissions, CheckAllPermissions) {
				logger.ThrowErrorLog("insufficient permissions")
				interceptor.SendErrRes(c, "unathorized request", "insufficient permissions", http.StatusUnauthorized)
				return
			}
		}
		c.Next()
	}

}
