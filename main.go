package main

import (
	"strconv"

	"github.com/rudolfoborges/gocrud/migrations"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	app := gofr.New()
	app.Migrate(migrations.All())

	app.GET("/users", func(c *gofr.Context) (interface{}, error) {
		query := `
            select * from users order by name
        `

		data := make([]User, 0)

		c.SQL.Select(c, &data, query)

		return data, nil
	})

	app.GET("/users/{id}", func(c *gofr.Context) (interface{}, error) {
		id, _ := strconv.Atoi(c.PathParam("id"))

		var count int

		c.SQL.Select(c, &count, "select count(*) from users where id = ?", id)

		if count == 0 {
			return nil, http.ErrorEntityNotFound{
				Name: "User",
			}
		}

		query := `
            select * from users where id = ?
        `

		var data User

		c.SQL.Select(c, &data, query, id)

		return data, nil
	})

	app.POST("/users", func(c *gofr.Context) (interface{}, error) {
		var user User

		err := c.Request.Bind(&user)
		if err != nil {
			return nil, err
		}

		query := `
            insert into users (name, email, password) values (?, ?, ?)
        `

		_, err = c.SQL.Exec(query, user.Name, user.Email, user.Password)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	app.DELETE("/users/{id}", func(c *gofr.Context) (interface{}, error) {
		id, _ := strconv.Atoi(c.PathParam("id"))

		query := `
            delete from users where id = ?
        `

		_, err := c.SQL.Exec(query, id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	app.Run()
}
