package core

import (
	"github.com/gin-gonic/gin"
	"github.com/xs0910/iam/pkg/errors"
	"log"
	"net/http"
)

// Response defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
type Response struct {
	Code      int         `json:"code"`                // Code define the business error code.
	Message   string      `json:"message"`             // Message contains the detail of this message.
	Data      interface{} `json:"data"`                // Data define the business data
	Reference string      `json:"reference,omitempty"` // Reference returns the reference document which maybe useful to solve this error.
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		// TODO log需要统一一下
		log.Printf("%#+v", err)

		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
			Data:      nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}
