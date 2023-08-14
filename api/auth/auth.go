package auth

import "github.com/gin-gonic/gin"

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		return c.Next()
	}
}

// func CheckBearerToken(token string) (claims map[string]interface{}, err error) {
// 	// Get the authorization header from the request
// 	authHeader := r.Header.Get("Authorization")
  
// 	// Check if the authorization header is present
// 	if authHeader == "" {
// 	  return nil, errors.New("Authorization header is not present")
// 	}
  
// 	// Check if the authorization header starts with "Bearer "
// 	if !strings.HasPrefix(authHeader, "Bearer ") {
// 	  return nil, errors.New("Authorization header is not in the correct format")
// 	}
  
// 	// Get the token from the authorization header
// 	token = strings.TrimPrefix(authHeader, "Bearer ")
  
// 	// Validate the token
// 	claims, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 	  // Get the public key from the issuer URL
// 	  key, err := jwks.Get(token.Header["kid"].(string))
// 	  if err != nil {
// 		return nil, err
// 	  }
  
// 	  // Verify the token signature with the public key
// 	  return key, token.Verify(key)
// 	})
  
// 	return claims, err
//   }