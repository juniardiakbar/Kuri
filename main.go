package main

import (
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func caesar(c *gin.Context) {
	pl := c.Query("plaintext")
	ci := c.Query("ciphertext")
	key := c.Query("key")
	k, err := strconv.Atoi(key)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "key is not a number",
		})
	} else {
		if pl != "" {
			cipher := ""
			for i := 0; i < len(pl); i++ {
				x := 0
				if int(pl[i]) > 64 && int(pl[i]) < 91 {
					x = (int(pl[i]) + k - int('A')) % 26
					x += int('A')
				} else {
					x = (int(pl[i]) + k - int('a')) % 26
					x += int('a')
				}
				cipher += string(x)
			}
			c.JSON(200, gin.H{
				"ciphertext": cipher,
			})
		} else if ci != "" {
			plain := ""
			for i := 0; i < len(ci); i++ {
				x := 0
				if int(ci[i]) > 64 && int(ci[i]) < 91 {
					x = (int(ci[i]) - k - int('A') + 260) % 26
					x += int('A')
				} else {
					x = (int(ci[i]) - k - int('a') + 260) % 26
					x += int('a')
				}
				plain += string(x)
			}
			c.JSON(200, gin.H{
				"plaintext": plain,
			})
		}
	}
}

func generateKey(str, key string) string {
	res := ""
	if len(str) < len(key) {
		temp := strings.Split(key, "")
		for i := 0; i < len(str); i++ {
			res += temp[i%len(temp)]
		}
		return res
	} else if len(str) == len(key) {
		return key
	} else {
		temp := strings.Split(key, "")
		for i := 0; i < len(str)-len(key); i++ {
			res += temp[i%len(temp)]
		}
		return key + res
	}
}

func vigenere(c *gin.Context) {
	pl := c.Query("plaintext")
	ci := c.Query("ciphertext")
	k := c.Query("key")
	if pl != "" {
		key := generateKey(pl, k)
		cipher := ""
		for i := 0; i < len(pl); i++ {
			x := 0
			if int(pl[i]) > 64 && int(pl[i]) < 91 {
				x = (int(pl[i]) + int(key[i]) - int('A')) % 26
				x += int('A')
			} else {
				x = (int(pl[i]) + int(key[i]) - int('a')) % 26
				x += int('a')
			}
			cipher += string(x)
		}
		c.JSON(200, gin.H{
			"ciphertext": cipher,
		})
	} else if ci != "" {
		key := generateKey(ci, k)
		plain := ""
		for i := 0; i < len(ci); i++ {
			x := 0
			if int(ci[i]) > 64 && int(ci[i]) < 91 {
				x = (int(ci[i]) - int(key[i]) - int('A') + 260) % 26
				x += int('A')
			} else {
				x = (int(ci[i]) - int(key[i]) - int('a') + 260) % 26
				x += int('a')
			}
			plain += string(x)
		}
		c.JSON(200, gin.H{
			"plaintext": plain,
		})
	}
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST"}
	r.Use(cors.New(config))
	r.POST("/caesarcipher", caesar)
	r.POST("/vigenerecipher", vigenere)
	r.Run()
}
