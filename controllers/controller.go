package controllers

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/sirawong/go-fiber-app/config"
	"github.com/sirawong/go-fiber-app/db"
	"github.com/sirawong/go-fiber-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mock data
func FakerData(c *fiber.Ctx) error {
	collection := db.MI.DB.Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	for i := 0; i < 50; i++ {
		_, err := collection.InsertOne(ctx, models.Product{
			Title:       faker.Word(),
			Description: faker.Paragraph(),
			Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", faker.UUIDDigit()),
			Price:       rand.Intn(90) + 10,
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "failed created",
				"error":   err,
			})
		}
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

// search data
func GetData(c *fiber.Ctx) error {
	collection := db.MI.DB.Collection(config.NewFlags.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var products []models.Product

	// filter search
	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"title": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	// sort
	if sort := c.Query("sort"); sort != "" {
		if sort == "asc" {
			findOptions.SetSort(bson.D{{"price", 1}})
		} else if sort == "desc" {
			findOptions.SetSort(bson.D{{"price", -1}})
		}
	}

	// page
	page, _ := strconv.Atoi(c.Query("page", "1"))
	var perPage int64 = 9

	total, _ := collection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * perPage) // 0-9, 9-18
	findOptions.SetLimit(perPage)                    // 9

	// get in db
	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var prodoct models.Product
		cursor.Decode(&prodoct)
		products = append(products, prodoct)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / perPage)),
	})
}
