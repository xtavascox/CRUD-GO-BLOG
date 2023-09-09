package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_blog/database"
	"go_blog/model"
	"log"
)

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "BlogList",
	}

	db := database.DBConnection

	var records []model.Blog
	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)

}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add Blog",
	}

	var record struct {
		Title string `json:"title"`
		Post  string `json:"post"`
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in BlogCreate parsing request. ")
		context["statusText"] = "Error"
		context["msg"] = "Something went wrong. "
		return c.JSON(err)
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error in BlogCreate generating UUID. ")
		context["statusText"] = "Error"
		context["msg"] = "Something went wrong. "
		return c.JSON(err)

	}
	blog := &model.Blog{
		ID:    id,
		Title: record.Title,
		Post:  record.Post,
	}
	result := database.DBConnection.Create(&blog)
	if result.Error != nil {
		log.Println("Error in BlogCreate inserting record. ")
		context["msg"] = "Something went wrong. "
		context["statusText"] = "Error"

		return c.JSON(result.Error)
	}
	context["msg"] = "Blog Created Successfully"
	context["data"] = record
	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog for the given ID ",
	}

	id := c.Params("id")
	var record model.Blog

	if err := database.DBConnection.First(&record, "id=?", id).Error; err != nil {
		log.Println("Error in BlogUpdate fetching record. ")
		context["statusText"] = "Error"
		context["msg"] = "Record not found. "
		return c.JSON(err)
	}

	//if record.ID == uuid.Nil {
	//	log.Println("Error in BlogUpdate record not found. ")
	//	context["statusText"] = "Error"
	//	context["msg"] = "Record not found. "
	//	c.Status(404)
	//	return c.JSON(context)
	//}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in BlogUpdate parsing request. ")
		context["statusText"] = "Error"
		context["msg"] = "Something went wrong. "
		return c.JSON(err)
	}
	result := database.DBConnection.Save(&record)

	if result.Error != nil {
		log.Println("Error in BlogUpdate updating record. ")
		context["msg"] = "Something went wrong. "
		context["statusText"] = "Error"
		return c.JSON(result.Error)
	}
	context["msg"] = "Blog Updated Successfully"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete Blog for the given ID ",
	}

	id := c.Params("id")
	log.Println("id: ", id)
	var record model.Blog

	if err := database.DBConnection.First(&record, "id=?", id).Error; err != nil {
		log.Println("Error in BlogUpdate fetching record. ")
		context["statusText"] = "Error"
		context["msg"] = "Record not found. "
		c.Status(404)
		return c.JSON(context)
	}

	result := database.DBConnection.Delete(record)

	if result.Error != nil {
		log.Println("Error in BlogDelete deleting record. ")
		context["msg"] = "Something went wrong. "
		context["statusText"] = "Error"
		c.Status(500)
		return c.JSON(context)
	}

	c.Status(200)
	return c.JSON(context)
}
