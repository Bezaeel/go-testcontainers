package customer

import (
	"encoding/json"
	"net/http"
	"sync"

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
		h.GET("/ask", ch.Sum)
		h.GET("/ask2", ch.SumBackground)
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

	result := ch.customerService.CreateCustomer(customer)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusCreated, result)
}

func (ch *CustomerController) Sum(c *gin.Context) {
	result := make(chan CustomerResult)
	go ch.customerService.AskTalabiAsync(result, 2)
	resp := <-result
	c.JSON(http.StatusOK, resp)
	return
}

func (ch *CustomerController) SumBackground(c *gin.Context) {
	var wg sync.WaitGroup
	go ch.customerService.AskTalabiAsyncBackground(&wg, 2)
	wg.Wait()
	c.JSON(http.StatusAccepted, "done!")
	return
}
