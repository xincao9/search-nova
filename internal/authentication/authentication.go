package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"search-nova/internal/constant"
	"search-nova/internal/util"
	userService "search-nova/service/user"
	"time"
)

func Authentication(c *gin.Context) {
	t, err := c.Cookie(constant.Token) // 请求必须携带token
	if err != nil {
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, "the request must carry user token")
		return
	}
	u, err := userService.U.GetUserByToken(t)
	if err != nil || u == nil || u.Expire.Before(time.Now()) { // 会话对象是否过期
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, "session expired or nonexistent session")
		return
	}
	c.Set(constant.SessionUser, u) // 设置本地会话
	if c.Request.Method == "GET" {
		return
	} else if c.Request.Method == "DELETE" || c.Request.Method == "POST" || c.Request.Method == "PUT" {

	}
}
