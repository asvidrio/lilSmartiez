package main

import (
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func showLoginPage(c *gin.Context) {

}

func login(c *gin.Context) { /*
		var user LoginInfo
		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"[ERROR] ": err})
			return
		}

		hashed_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		user.Password = string(hashed_password)

		users := findAllUsers(c)

		found := false
		var foundUser User
		for i := 0; i < len(users); i++ {
			if (users[i].Username == user.Username) && (users[i].Password == user.Password) {
				found = true
				foundUser = *users[i]
				break
			} else {
				foundUser = *users[0]
			}
		}

		if found {
			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    foundUser.Id.String(),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //expires after 1 day
			})

			// If the user is created, set the token in a cookie and log the user in
			token, err := claims.SignedString([]byte(getSecretKey()))

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
			}
			log.Println("token", token)
			c.SetCookie("token", token, 3600*24, "", "", true, true) //MaxAge: 24 hours
			c.Set("is_logged_in", true)

			c.JSON(http.StatusOK, gin.H{
				"message":  "created new account",
				"Email":    foundUser.Email,
				"Username": foundUser.Username,
				"Password": foundUser.Password})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "user not found"})
		}*/
}
