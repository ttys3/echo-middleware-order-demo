package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC | log.Lmicroseconds)
	log.Printf("http server starting...")
	e := echo.New()
	e.HideBanner = true

	log.Printf("use Middleware1")
	e.Use(Middleware1())
	log.Printf("use Middleware2")
	e.Use(Middleware2())
	log.Printf("use Middleware3")
	e.Use(Middleware3())
	e.GET("/", func(c echo.Context) error {
		time.Sleep(50 * time.Millisecond)
		log.Print("::handler exec")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

/*
2023/03/22 22:13:00 Middleware3 pre
2023/03/22 22:13:00 Middleware2 pre
2023/03/22 22:13:00 Middleware1 pre

2023/03/22 22:13:00 Middleware1 post
2023/03/22 22:13:00 Middleware2 post
2023/03/22 22:13:00 Middleware3 post
2023/03/22 22:13:00 handler 2023-03-22 22:13:00.938538351 +0800 HKT m=+1.285199382
*/

func Middleware1() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// applyMiddleware() 的时候执行
		log.Printf("Middleware1 wrap")
		return func(c echo.Context) error {
			tstart := time.Now()
			log.Printf("::Middleware1 begin exec")
			time.Sleep(100 * time.Millisecond)
			c.Response().Header().Set("from-middleware1", fmt.Sprintf("::middleware1: %v", time.Now()))
			ret := next(c)
			cost := time.Since(tstart).Milliseconds()
			log.Printf("Middleware1 took %v", cost)
			c.Response().Header().Set("from-middleware1-cost", fmt.Sprintf("%v", cost))
			return ret
		}
	}
}

func Middleware2() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// applyMiddleware() 的时候执行
		log.Printf("Middleware2 wrap")
		return func(c echo.Context) error {
			tstart := time.Now()
			log.Printf("::Middleware2 begin exec")
			time.Sleep(200 * time.Millisecond)
			log.Printf("Middleware2: get header from middleware1: %v", c.Response().Header().Get("from-middleware1"))
			ret := next(c)
			log.Printf("Middleware2 took %v", time.Since(tstart).Milliseconds())
			return ret
		}
	}
}

func Middleware3() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// applyMiddleware() 的时候执行
		log.Printf("Middleware3 wrap")
		return func(c echo.Context) error {
			log.Printf("::Middleware3 begin exec")
			startTime := time.Now()
			time.Sleep(400 * time.Millisecond)
			ret := next(c)
			log.Printf("Middleware3 took %v", time.Since(startTime).Milliseconds())
			return ret
		}
	}
}

/*
	e.findRouter(r.Host).Find(r.Method, GetPath(r), c)
	h = c.Handler()
	h = applyMiddleware(h, e.middleware...)
*/

// h 是最后一个执行的, 它是 handler
func applyMiddleware(h echo.HandlerFunc, middleware ...echo.MiddlewareFunc) echo.HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

var handler = func(c echo.Context) error {
	log.Printf("handler %v", time.Now())
	return c.String(http.StatusOK, "Hello, World!")
}

var m3 = func() echo.HandlerFunc {
	log.Printf("Middleware3 wrap")
	// time.Sleep(810 * time.Millisecond)
	return func(c echo.Context) error {
		log.Printf("Middleware3 begin exec")
		// time.Sleep(820 * time.Millisecond)
		return handler(c)
	}
}

var m2 = func() echo.HandlerFunc {
	log.Printf("Middleware2 wrap")
	// time.Sleep(810 * time.Millisecond)
	return func(c echo.Context) error {
		log.Printf("Middleware2 begin exec")
		// time.Sleep(820 * time.Millisecond)
		return m3()(c)
	}
}

/*
m3 {
next: {
m2 {
next: {
m1 {
}
}
}
}
}
*/
