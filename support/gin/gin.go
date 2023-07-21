package gin

import (
	"net/http"
	"strings"

	"github.com/davidebianchi/gswagger/apirouter"

	"github.com/gin-gonic/gin"
)

// HandlerFunc is the http type handler used by Gin
type HandlerFunc gin.HandlerFunc
type Route = gin.IRoutes

type ginRouter struct {
	router *gin.Engine
}

func (r ginRouter) AddRoute(method string, path string, handler HandlerFunc) Route {
	switch method {
	case http.MethodGet:
		return r.router.GET(path, gin.HandlerFunc(handler))
	case http.MethodPost:
		return r.router.POST(path, gin.HandlerFunc(handler))
	case http.MethodPut:
		return r.router.PUT(path, gin.HandlerFunc(handler))
	case http.MethodDelete:
		return r.router.DELETE(path, gin.HandlerFunc(handler))
	default:
		return r.router.Any(path, gin.HandlerFunc(handler))
	}
}

func (r ginRouter) SwaggerHandler(contentType string, blob []byte) HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", contentType)
		c.Data(http.StatusOK, contentType, blob)
	}
}

func (r ginRouter) TransformPathToOasPath(path string) string {
	segments := strings.Split(path, "/")

	for i, segment := range segments {
		if strings.HasPrefix(segment, ":") {
			segments[i] = "{" + segment[1:] + "}"
		}
	}

	return strings.Join(segments, "/")
}

func NewRouter(router *gin.Engine) apirouter.Router[HandlerFunc, Route] {
	return &ginRouter{
		router: router,
	}
}
