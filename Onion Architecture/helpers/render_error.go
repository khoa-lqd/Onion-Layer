package helpers

import "github.com/gin-gonic/gin"

func RenderError(code int) gin.H {
	MESSAGE := map[int]string{
		400: "Bad Request",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		500: "System Error",
		503: "Maintenance",
	}

	output := gin.H{
		"error": gin.H{
			"code":    code,
			"message": MESSAGE[code],
		},
	}

	return output
}
