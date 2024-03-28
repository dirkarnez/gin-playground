package main

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type Nested struct {
	Content string                `form:"content"`
	Image   *multipart.FileHeader `form:"image"`
}

type Form struct {
	NestedList Nested `form:"nested"`
	S          string `form:"s"`
}

func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		var form Form
		err := c.ShouldBind(&form)
		if err != nil {
			fmt.Printf("%#v\n", err)
			return
		}
		fmt.Printf("324234")
	})

	r.Run(":3000")
}
