package ctx

import "github.com/gin-gonic/gin"

const (
	UserID = "user_id"
)

type HandlersChain []HandlerFunc
type HandlerFunc func(*Context)

type Engine struct {
	*gin.Engine
}

func Default() *Engine {
	// g := gin.New()
	// g.Use(gin.Recovery())
	return &Engine{gin.Default()}
}

func (en *Engine) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandlers := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandlers = append(RHandlers, handleFunc(handle))
	}
	return &RouterGroup{en.Engine.Group(relativePath, RHandlers...)}
}
func (en *Engine) Use(middlewares ...func(c *Context)) gin.IRoutes {
	rMiddlewares := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		rMiddlewares = append(rMiddlewares, handleFunc(middleware))
	}
	return en.RouterGroup.Use(rMiddlewares...)
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) Use(middlewares ...func(c *Context)) gin.IRoutes {
	rMiddlewares := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		rMiddlewares = append(rMiddlewares, handleFunc(middleware))
	}
	return r.RouterGroup.Use(rMiddlewares...)
}
func (r *RouterGroup) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return &RouterGroup{r.RouterGroup.Group(relativePath, rHandles...)}
}

func (r *RouterGroup) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.GET(relativePath, rHandles...)
}

func (r *RouterGroup) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.POST(relativePath, rHandles...)
}

func (r *RouterGroup) PUT(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.PUT(relativePath, rHandles...)
}

func (r *RouterGroup) DELETE(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.DELETE(relativePath, rHandles...)
}

func handleFunc(handler func(c *Context)) func(c1 *gin.Context) {
	return func(c *gin.Context) {
		handler(&Context{Context: c})
	}
}
