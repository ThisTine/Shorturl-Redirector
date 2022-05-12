package routes

import (
	"github.com/ThisTine/shorturl-redirector/db"
	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	pathid := c.Params("pathid")
	// fmt.Println(pathid)
	var url db.Link
	db.DB.First(&url, "path", pathid)
	// // fmt.Println(url.link)
	if url.Link != "" {
		return c.Redirect("https://" + url.Link)
	}
	return c.Redirect("https://thistine.com/shorturl")
	// return c.SendString("Short url not found")
}
