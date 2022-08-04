package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func createServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS()) // 추후에 웹 페이지에서 요청을 날릴 수 있도록 CORS 헤더 추가
	e.GET("/note", func(c echo.Context) error {
		id := c.QueryParam("from")
		getOneNote(id)
		fmt.Println(id)
		// if err != nil {
		// 	c.Error(err)
		// }
		return nil
	})

	e.POST("/note/:user", func(c echo.Context) error {
		user := c.Param("user")
		body := &note{}
		if err := c.Bind(body); err != nil {
			return err
		}

		note, err := createNote(user, *body)
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusCreated, note)
	})

	e.GET("/note/:user", func(c echo.Context) error {
		user := c.Param("user")
		id := c.QueryParam("from")
		notes, err := getNotes(user, id)

		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, notes)
	})

	// e.GET("/note/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	notes, err := getOneNote(id)

	// 	if err != nil {
	// 		c.Error(err)
	// 	}
	// 	return c.JSON(http.StatusOK, notes)
	// })

	e.PUT("/note/:user/:id", func(c echo.Context) error {
		user := c.Param("user")
		id := c.Param("id")
		body := &note{}
		if err := c.Bind(body); err != nil {
			return err
		}

		note, err := updateNote(user, id, *body)
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusCreated, note)
	})

	e.DELETE("/note/:user/:id", func(c echo.Context) error {
		user := c.Param("user")
		id := c.Param("id")
		if err := deleteNote(user, id); err != nil {
			c.Error(err)
		}
		return c.NoContent(http.StatusAccepted)
	})
	return e
}
