package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type EWrap struct {
	Error E `json:"error"`
}

type E struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Context interface{} `json:"context,omitempty"`
}

type S struct {
	Data     interface{} `json:"data,omitempty"`
	DateTime time.Time   `json:"datetime"`
}

func Error(c *gin.Context, status int, data E) {
	c.AbortWithStatusJSON(status, EWrap{
		Error: data,
	})
}

func Success(c *gin.Context, status int, data S) {
	if data.DateTime.IsZero() {
		data.DateTime = time.Now()
	}
	c.JSON(status, data)
}
