package file

import (
	"fmt"
	"pdf/utils"

	"github.com/gin-gonic/gin"
)

func ImageResize(c *gin.Context) {
	err := utils.ImageResize(`C:\Clarenti\Data\Project\Sampling\PDF/assets/imgs/background.jpg`)
	if err != nil {
		fmt.Println(err)
	}
}
