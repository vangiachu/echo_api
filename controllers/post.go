package controllers

import (
    "echo_api/models"
    "echo_api/database"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
    userID := c.Get("userID").(uint)
    post := new(models.Post)
    if err := c.Bind(post); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }
    post.UserID = userID
    database.DB.Create(post)
    return c.JSON(http.StatusCreated, post)
}

func GetPosts(c echo.Context) error {
    userID := c.Get("userID").(uint)
    var posts []models.Post
    database.DB.Where("user_id = ?", userID).Find(&posts)
    return c.JSON(http.StatusOK, posts)
}

func UpdatePost(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var post models.Post
    database.DB.First(&post, id)
    if post.ID == 0 {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Post not found"})
    }
    userID := c.Get("userID").(uint)
    if post.UserID != userID {
        return c.JSON(http.StatusForbidden, echo.Map{"error": "Unauthorized"})
    }
    c.Bind(&post)
    database.DB.Save(&post)
    return c.JSON(http.StatusOK, post)
}

func DeletePost(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var post models.Post
    database.DB.First(&post, id)
    if post.ID == 0 {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Post not found"})
    }
    userID := c.Get("userID").(uint)
    if post.UserID != userID {
        return c.JSON(http.StatusForbidden, echo.Map{"error": "Unauthorized"})
    }
    database.DB.Delete(&post)
    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted"})
}
