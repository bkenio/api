package routes

import (
	"fmt"

	"github.com/bken-io/api/api/db"
	"github.com/bken-io/api/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/teris-io/shortid"
)

// GetVideo returns a video
func GetVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var video models.Video
	result := db.Where("id = ? and visibility = 'public'", id).Find(&video)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	if video.ID == "" {
		return c.SendStatus(404)
	}

	video.URL = fmt.Sprintf("https://cdn.bken.io/v/%s/hls/master.m3u8", video.ID)
	return c.JSON(video)
}

// GetVideos returns all videos
func GetVideos(c *fiber.Ctx) error {
	db := db.DBConn
	var videos []models.Video
	db.Find(&videos).Where("visibility = 'public'")
	return c.JSON(videos)
}

// CreateVideo creates a new video
func CreateVideo(c *fiber.Ctx) error {
	db := db.DBConn
	video := new(models.Video)

	if err := c.BodyParser(video); err != nil {
		return c.Status(400).SendString("video input failed unmarshalling")
	}

	id, sidErr := shortid.Generate()
	if sidErr != nil {
		return c.Status(500).SendString("failed to create video short id")
	}

	video.ID = id
	db.Create(&video)
	return c.JSON(video)
}

// DeleteVideo creates a new video
func DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn

	var video models.Video
	db.Where("id = ?", id).Find(&video)

	fmt.Println(video)

	if video.Title == "" {
		return c.SendStatus(404)
	}

	db.Delete(&video)
	return c.SendString("Video successfully deleted")
}