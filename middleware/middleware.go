package middleware

import (
	"github.com/AKAZJAYA/blogbackend/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx)error {

	cookie:=c.Cookies("jwt")

	if _,err := util.Parsejwt(cookie); err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"unAuthenticated",
		})
	}
	return c.Next()
}