package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main () {
	app := InitApp()
	// Start the Fiber application
	log.Fatal(app.Listen(":3000"))

}

func InitApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/getvalue/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	app.Get("/getname/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name", "Guest")
		return c.SendString("Hello, " + name + "!")
	})

	app.Get("/getall/*", func(c *fiber.Ctx) error {
		return c.SendString("All parameters: " + c.Params("*"))
	})

	app.Static("/static", "./public")


	// Lecture 2: 
	// // Grouping routes with a prefix

	apiL1 := app.Group("/api")  // /api

	apiL1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Home")
	})   // /api

	v1 := apiL1.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API v1")
	})   // /api/v1
	

	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("API v1 List")
	})          // /api/v1/list
	v1.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("API v1 User")
	})          // /api/v1/user

	v2 := apiL1.Group("/v2")

	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API v2")
	})   // /api/v2

	v2.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("API v2 List")
	})          // /api/v2/list
	v2.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("API v2 User")
	})          // /api/v2/user

	// Routes
	app.Route("/route_example", func(api fiber.Router) {
      api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("get all")
		})
		api.Post("/", func(c *fiber.Ctx) error {
			return c.SendString("create")
		})
		api.Get("/:id", func(c *fiber.Ctx) error {
			return c.SendString("get " + c.Params("id"))
		})
		api.Put("/:id", func(c *fiber.Ctx) error {
			return c.SendString("update " + c.Params("id"))
		})
		api.Delete("/:id", func(c *fiber.Ctx) error {
			return c.SendString("delete " + c.Params("id"))
		})
		api.Patch("/:id", func(c *fiber.Ctx) error {
			return c.SendString("patch " + c.Params("id"))
		})
	})

	// Shut Down Handler
	app.Get("/shutdown", func(c *fiber.Ctx) error {
		log.Println("Shutting down the server...")
		if c.IP() == "127.0.0.1" {
			return app.Shutdown()
		}
		return c.SendString("Shutdown is not allowed from this IP address.")
	})

	// TEST
	app.Get("/test-token", func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	// Lecture 3: 
	// // CTX

	app.Get("l3/:name/:product", func(c *fiber.Ctx) error {
		
		params := c.AllParams()

		name := c.Params("name")
		product := c.Params("product")

		fmt.Println(params["name"])
		fmt.Println(params["product"])

		fmt.Println(name)
		fmt.Println(product)

		return c.SendStatus(fiber.StatusOK)
	})

	// Download

	app.Get("l3/download", func(c *fiber.Ctx) error {
		
		return c.Download("./files/samplefile.txt")
	})

	// Body Parser
	// Field names should start with an uppercase letter
	
	app.Post("l3/bodyparser", func(c *fiber.Ctx) error {
		type Person struct {
			Name string `json:"name"`
			Pass string `json:"pass"`
		}
		newId := 123453421
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}


		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name": p.Name,
			"pass": p.Pass,
			"id": newId,
		})

	})

	app.Post("l3/bodyparser1", func(c *fiber.Ctx) error {
		type Person struct {
			Name string `json:"name"`
			Pass string `json:"pass"`
		}
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}


		return c.Status(fiber.StatusOK).JSON(p)

	})

	// L4:
	// Cookie
	app.Get("/l4/cookie", func(c *fiber.Ctx) error {
		// Create cookie
		cookieAuth := new(fiber.Cookie)
		cookieAuth.Name = "john"
		cookieAuth.Value = "doe"
		cookieAuth.Expires = time.Now().Add(24 * time.Hour)

		cookieTheme := new(fiber.Cookie)
		cookieTheme.Name = "app_theme"
		cookieTheme.Value = "dark"

		// Set cookie
		c.Cookie(cookieAuth)
		c.Cookie(cookieTheme)
		
		return c.SendStatus(fiber.StatusOK)

	})

	// access the cookies sent by browser
	app.Get("/l4/cookie/checkout", func(c *fiber.Ctx) error {
		fmt.Println("username", c.Cookies("username", "no user cookie found"))
		fmt.Println("app_theme", c.Cookies("app_theme", "no app_theme cookie found"))
		
		return c.SendStatus(fiber.StatusOK)

	})

	// delete cookies (when logout)
	app.Get("l4/cookie/logout", func(c *fiber.Ctx) error {

		// Clears all cookies:
		c.ClearCookie()

		// Expire specific cookie by name:
		c.ClearCookie("app_theme")

		// Expire multiple cookies by names:
  		c.ClearCookie("token", "session", "track_id", "version")
		
		return c.SendStatus(fiber.StatusOK)

	})



	app.Get("l4/cookie/parse", func(c *fiber.Ctx) error {
		// Field names should start with an uppercase letter
		type Person struct {
			Name     string  `cookie:"name"`
			Age      int     `cookie:"age"`
			Job      bool    `cookie:"job"`
			Apptheme    string  `cookie:"app_theme"`
		}
		p := new(Person)

		if err := c.CookieParser(p); err != nil {
			return err
		}

		log.Println(p.Name)     // Joseph
		log.Println(p.Age)      // 23
		log.Println(p.Job)      // true
		log.Println(p.Apptheme)

		return c.SendStatus(fiber.StatusOK)
	})

	// Lecture 5
	// Params

	app.Get("l5/params/:name", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	app.Get("l5/params1/:name/:id", func(c *fiber.Ctx) error {

		name := c.Params("name")
		// id, _ = strconv.Atoi(c.Params("id"))
		id, _ := c.ParamsInt("id")


		return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"name": name,
				"id": id,
			})
	})

	// Params Parser
	app.Get("/l5/paramsparser/:name/:age/:job", func(c *fiber.Ctx) error {
		// Field names should start with an uppercase letter
		type Person struct {
			Name     string  `params:"name"`
			Age      int64     `params:"age"`
			Job      string    `params:"job"`
		}

		p := new(Person)
		if err := c.ParamsParser(p); err != nil {
			return err
		}

		log.Println(p.Name)     // Joseph
		log.Println(p.Age)      // 23
		log.Println(p.Job)      // true

		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name": p.Name,
			"age": p.Age,
			"job": p.Job,
		})
	})

	// Queries
	// GET http://example.com/?name=alex&want_pizza=false&id=
	app.Get("/l5/queries", func(c *fiber.Ctx) error {
		m := c.Queries()
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name": m["name"],
			"want_pizza": m["want_pizza"],
			"id": m["id"],
		})
	})

	// Query
	app.Get("/l5/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		want_pizza := c.QueryBool("want_pizza")
		id := c.QueryInt("id")
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name": name,
			"want_pizza": want_pizza,
			"id": id,
		})
	})

	// Query Parser


	app.Get("/l5/queryparser", func(c *fiber.Ctx) error {

		// Field names should start with an uppercase letter
		type Person struct {
			Name     string     `query:"name"`
			Pass     string     `query:"pass"`
			Products []string   `query:"products"`
		}
		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name": p.Name,
			"pass": p.Pass,
			"products": p.Products,
		})
	})

	// Lecture 6
	// Form Value, 
	app.Post("/l6/login", func(c *fiber.Ctx) error {
		// Get first value from form field "name":
		username := c.FormValue("username")
		password := c.FormValue("password")
		// => "john" or "" if not exist

		if username != "john" || password != "john123" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"message": "You are not authorized!",
				})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "welcome!",
		})
	})

	// file upload
	app.Post("/l6/upload", func(c *fiber.Ctx) error {
		file, _ := c.FormFile("document")

		return c.SaveFile(file, "/Users/sunilvytla/projects/Golang/fiber/01-fiber/files/uploaded.txt")
	})


	app.Post("/l6/uploads", func(c *fiber.Ctx) error {
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
				// => *multipart.Form

				// Get all files from "documents" key:
				files := form.File["documents"]
				// => []*multipart.FileHeader

				// Loop through files:
				for _, file := range files {
					// name := fmt.Sprintf(file.Filename, file.Size, time.Now().String())
					// log.Println(name)
					// fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
					// => "tutorial.pdf" 360641 "application/pdf"
					filePath := fmt.Sprintf("/Users/sunilvytla/projects/Golang/fiber/01-fiber/files/%s", file.Filename)
					// Save the files to disk:
					if err := c.SaveFile(file, filePath); err != nil {
						return err
					}
				}
				return err
		}
		return c.SendStatus(fiber.StatusOK)
	})

	type keyType struct{}
	var userKey keyType
	// This is called before the below endpoints are called.
	app.Use(func(c *fiber.Ctx) error {
		log.Println("app use func")
		c.Locals(userKey, "admin")
		return c.Next()
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		user, ok := c.Locals(userKey).(string) // Retrieves the data stored under the key and performs a type assertion
		if ok && user == "admin" {
			return c.Status(fiber.StatusOK).SendString("Welcome, admin!")
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	app.Get("/admins", func(c *fiber.Ctx) error {
		user, ok := c.Locals(userKey).(string) // Retrieves the data stored under the key and performs a type assertion
		if ok && user == "admin" {
			return c.Status(fiber.StatusOK).SendString("Welcome, admins!")
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	app.Get("/l7/getname/:name", func(c *fiber.Ctx) error {
		name := []byte(c.Params("name"))
		return c.Send(name)
	})

	app.Get("/l7/sendfile", func(c *fiber.Ctx) error {
		return c.SendFile("./files/uploaded.txt")
	})

	app.Get("/l7/sendStatus", func(c *fiber.Ctx) error {
		c.SendString("hello world")
		return c.SendStatus(fiber.StatusUnsupportedMediaType)
	})

	// SendString (this automatically sends status code 200 by default if we dont send any status code back)
	app.Get("/l7/sendString", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Get("/l7/sendStringclone", func(c *fiber.Ctx) error {
		c.SendString("hello world")
		return c.SendStatus(fiber.StatusOK)
	})

	// we can chain Status with different things like below.
	app.Get("/l7/statuschaining1", func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	})

	app.Get("/l7/statuschaining2", func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./files/uploaded.txt")
	})

	app.Get("/l7/statuschaining3", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return nil
	})

	return app
}