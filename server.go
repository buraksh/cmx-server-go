// Copyright 2022 Burak Şahin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "CMX Server")
	})

	r.GET("/api/config/v1/maps", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		data, err := os.ReadFile("./data/map.json")

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Could not read map file",
			})
			return
		}

		c.String(http.StatusOK, string(data))
	})

	r.GET("/api/config/v1/heterarchy/allUserLevels", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		data, err := os.ReadFile("./data/heterarchy.json")

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Could not read heterarchy file",
			})
			return
		}

		c.String(http.StatusOK, string(data))
	})

	r.GET("/api/config/v1/maps/imagesource/:imageName", func(c *gin.Context) {
		c.File("./data/" + c.Param("imageName"))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
