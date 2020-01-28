package authenticator

// import (
// 	"time"

// 	jwt "github.com/appleboy/gin-jwt/v2"
// 	"github.com/gin-gonic/gin"
// 	Models "github.com/golang-crud/model"
// )

// func AuthenticatorJwt() {
// 	authMiddleware := &jwt.GinJWTMiddleware{
// 		Realm:      "test zone",
// 		Key:        []byte("secret key"),
// 		Timeout:    time.Hour,
// 		MaxRefresh: time.Hour,
// 		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
// 			var user Models.User

// 			err := Models.Authenticator(&user)

// 			if err != nil {
// 				return
// 			}

// 		},
// 		Authorizator: func(userId string, c *gin.Context) bool {
// 			if userId == "admin" {
// 				return true
// 			}

// 			return false
// 		},
// 		Unauthorized: func(c *gin.Context, code int, message string) {
// 			c.JSON(code, gin.H{
// 				"code":    code,
// 				"message": message,
// 			})
// 		},
// 		// TokenLookup is a string in the form of "<source>:<name>" that is used
// 		// to extract token from the request.
// 		// Optional. Default value "header:Authorization".
// 		// Possible values:
// 		// - "header:<name>"
// 		// - "query:<name>"
// 		// - "cookie:<name>"
// 		TokenLookup: "header:Authorization",
// 		// TokenLookup: "query:token",
// 		// TokenLookup: "cookie:token",

// 		// TokenHeadName is a string in the header. Default value is "Bearer"
// 		TokenHeadName: "Bearer",

// 		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
// 		TimeFunc: time.Now,
// 	}
// }
