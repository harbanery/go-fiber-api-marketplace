package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}

func JWTMiddleware() fiber.Handler {
	secretKey := os.Getenv("SECRETKEY")
	return func(c *fiber.Ctx) error {
		tokenString := ExtractToken(c)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":     "unauthorized",
				"statusCode": 401,
				"message":    "Token string unauthorized",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":     "unauthorized",
				"statusCode": 401,
				"message":    "Token unauthorized",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("user", claims)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":     "unauthorized",
				"statusCode": 401,
				"message":    "Token claims unauthorized",
			})
		}

		return c.Next()
	}
}

// func Authorize(requiredRole string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		user := c.Locals("user").(jwt.MapClaims)
// 		role := user["role"].(string)

// 		if requiredRole != "" && role != requiredRole {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"status":     "forbidden",
// 				"statusCode": 403,
// 				"message":    "Forbidden: incorrect role",
// 			})
// 		}

// 		return c.Next()
// 	}
// }

func UserLocals(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(jwt.MapClaims)

	return user
}

// func JWTValidation(requiredRole string, c *fiber.Ctx) error {
// 	user := c.Locals("user").(jwt.MapClaims)
// 	role := user["role"].(string)

// 	if requiredRole != "" && role != requiredRole {
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"status":     "forbidden",
// 			"statusCode": 403,
// 			"message":    "Forbidden: incorrect role",
// 		})
// 	}

// 	return nil
// }
