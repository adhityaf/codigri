package main

import (
	"go-simple-chat/config"
	"go-simple-chat/models"
	"go-simple-chat/params"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db := config.ConnectDB()

	app := fiber.New()
	// Default config
	app.Use(cors.New())

	pusherSecure, err := strconv.ParseBool(os.Getenv("PUSHER_SECURE"))
	if err != nil {
		log.Fatal(err)
	}

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APPID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  pusherSecure,
	}

	app.Post("/v1/api/messages", func(c *fiber.Ctx) error {
		var data params.RequestMessage

		// parse request body
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		dataModel := models.Message{
			Username: data.Username,
			Message:  data.Message,
		}

		// insert to database
		err = db.Create(&dataModel).Error
		if err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)
		return c.Status(fiber.StatusOK).JSON(data)
	})

	app.Listen(":3000")
}
