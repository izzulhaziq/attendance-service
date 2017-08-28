package resources

import "github.com/gin-gonic/gin"

func getStringParam(c *gin.Context, name string) (string, error) {
	return c.Params.ByName(name), nil
}
