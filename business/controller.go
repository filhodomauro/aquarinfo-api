package business

import "github.com/gin-gonic/gin"
import "net/http"

type business struct {
	City string `json:"city"`
	Name string `json:"name"`
}

func List(c *gin.Context) {
	business1 := business{}
	business1.City = "Barueri - SP"
	business1.Name = "Petz"

	business2 := business{}
	business2.City = "Barueri - SP"
	business2.Name = "Pet Love"
	c.JSON(http.StatusOK, []business{business1, business2})
}
