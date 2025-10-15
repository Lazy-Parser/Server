package service

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/Lazy-Parser/Server/database"
	"github.com/Lazy-Parser/Server/dto"
	"github.com/Lazy-Parser/Server/token"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
)

func AuthCreateUser(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// AuthLogInFirst First enter username, so server will send one-time password to the telegram
func AuthLogInFirst(c *gin.Context, userRepo database.UserRepo) {
	var body dto.LoginFirstDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// find in the database
	user, err := userRepo.FindByUsername(body.Username)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "error": "user not found"})
		return
	}

	// create random password and send to the tg
	passwd := makePassword()

	// save random password
	payload := map[string]interface{}{
		"password":           passwd,
		"last_password_time": time.Now().String(),
	}
	err, rowsAffected := userRepo.Update(user.ID, payload)
	if err != nil || rowsAffected == 0 {
		c.JSON(400, gin.H{"status": "error", "error": err.Error()})
		return
	}

	sendTg(passwd)

	c.JSON(200, gin.H{"status": "success"})
}
func makePassword() string {
	var pass string
	for range 9 {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		pass += strconv.FormatInt(n.Int64(), 10)
	}

	return pass
}
func sendTg(passwd string) {
	// TODO
	log.Printf("The password in tg: %s\n", passwd)
}

// AuthLogInSecond Second enter the password from the telegram
func AuthLogInSecond(c *gin.Context, userRepo database.UserRepo) {
	var body dto.LoginSecondDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "error": err.Error()})
		return
	}

	// find user and check password
	user, err := userRepo.FindByUsername(body.Username)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "error": "user not found"})
	}

	// check password
	if body.Password != user.Password {
		c.JSON(400, gin.H{"status": "error", "error": "wrong credentials"})
		return
	}

	// create tokens
	accessToken, err := token.GenerateAccessToken(user)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "error": err.Error()})
	}
	refreshToken, err := token.GenerateRefreshToken(user)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "error": err.Error()})
	}

	// save refresh: TODO
	c.SetCookie("refresh_token", refreshToken, int(token.AccessTTL.Seconds()), "/", "localhost", true, true)

	c.JSON(200, gin.H{"status": "success", "token": accessToken})
	return
}

func AuthRefresh(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AuthLogout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
