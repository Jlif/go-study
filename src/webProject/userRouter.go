package main

import (
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) {
	e.GET("/user", userHandler)
	e.POST("/login", loginHandler)
}
