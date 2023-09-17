package customer

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService CustomerService
}

func NewCustomerController(controller *gin.Engine) {
	ch := &CustomerController{
		customerService: NewCustomerService(),
	}

	h := controller.RouterGroup.Group("/customers")
	{
		h.POST("/", ch.CreateCustomer)
	}
}

func (ch *CustomerController) CreateCustomer(c *gin.Context) {
	// var dto dtos.CustomerDTO
	var customer Customer

	err := json.NewDecoder(c.Request.Body).Decode(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	result, err := ch.customerService.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, result)
}
