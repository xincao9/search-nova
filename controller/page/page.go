package page

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"search-nova/internal/constant"
	"search-nova/internal/util"
	"search-nova/service/page"
)

func Route(engine *gin.Engine) {
	engine.GET("/page", func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			util.RenderJSON(c, http.StatusBadRequest, "")
			return
		}
		r, err := page.P.Query(text)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, r)
	})
}
