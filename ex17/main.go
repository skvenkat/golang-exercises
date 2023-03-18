package mmport (
Dr	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
        "github.com/gofiber/fiber/v2
)

const (
	HOST = "0.0.0.0"
	PORT = 8080
)

var ADDR_STR = fmt.Sprintf("%s:%d", HOST, PORT)



func DefaultHTTP() {
	fmt.Println("=== net/http server ===")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Golang Inbuilt HTTP Server !!!"))
	})

	srv := &http.Server{
		Addr: ADDR_STR,
	}
	log.Fatal(srv.ListenAndServe())
}


func EchoServer() {
	fmt.Println("=== echo framework ===")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from echooooo!!!")
	})
	e.Logger.Fatal(e.Start(ADDR_STR))
}

func ginRouter() *gin.Engine {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Gin!!!")
	})
	return g 
}

func GinServer() {
	fmt.Println("=== gin framework ===")
	r := ginRouter()
	r.Run(ADDR_STR)
}

func GorillaMux() {
	fmt.Println("=== gorilla mux framework ===")

	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Gorilla MuX !!!"))
	}).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr: ADDR_STR,
	}
	log.Fatal(srv.ListenAndServe())
}

func FiberServer() {
    app := fiber.New()
    app.Get("/", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Hello, World!")
    })

    app.Listen(":4000")
}

func main() {
	webserver := os.Args[1]
	if webserver == "echo" {
		EchoServer()
	} else if webserver == "gin" {
		GinServer()
	} else if webserver == "gorilla" {
		GorillaMux()
        } else if webserver == "fiber" {
                FiberServer()
	} else {
		DefaultHTTP()
	}
}
