package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/michaelCHU95/auth-hub/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HandleSignup(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Get the user info from request body
		body := new(SignupRequest)

		body_err := c.Bind(&body)
		if body_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"error":   body_err.Error(),
			})
			return
		}

		// Hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed at Password Hashing",
				"error":   err.Error(),
			})

			return
		}

		// Create User
		new_user := models.User{
			First_Name: body.First_Name,
			Last_Name:  body.Last_Name,
			Email:      body.Email,
			Password:   string(hash),
		}

		result := db.Create(&new_user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Transaction Failed",
				"error":   result.Error.Error(),
			})

			return
		}

		// Response
		c.JSON(http.StatusOK, gin.H{})
	}

	return gin.HandlerFunc(fn)
}

func HandleLogin(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Handling request body in JSON
		body := new(LoginRequest)

		err := c.Bind(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"error":   err.Error(),
			})
		}

		// Lookup User from Database
		var user models.User
		db.Find(&user, "email = ?", body.Email)

		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Email/Password",
			})
			return
		}

		// Validate the password
		pwd_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if pwd_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Password",
				"error":   pwd_err.Error(),
			})

			return
		}

		// Create & Signing a JWT Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			// Expired in one month
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenStr, token_err := token.SignedString([]byte(os.Getenv("SECRET")))

		if token_err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create JWT Token",
				"error":   token_err.Error(),
			})
			return
		}

		// Send JWT Token back to Client
		c.JSON(http.StatusOK, gin.H{
			"token": tokenStr,
		})
	}

	return gin.HandlerFunc(fn)
}

func HandlePasswordReset(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {}

	return gin.HandlerFunc(fn)
}

func HandleReprovision(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {}

	return gin.HandlerFunc(fn)
}
