package productController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"items/dto"
	service "items/services"
	client "items/services/repositories"
	"net/http"
)

var (
	itemService = service.NewItemServiceImpl(
		client.NewItemInterface("localhost", 27017, "items"),
		client.NewMemcachedInterface("localhost", 11211),
		client.NewQueueClient("user", "password", "localhost", 5672),
		client.NewSolrClient("localhost", 8983, "items"),
	)
)

func GetItemById(c *gin.Context) {
	var itemDto dto.ItemDto
	id := c.Param("item_id")
	itemDto, err := itemService.GetItemById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemDto)
}

func InsertItem(c *gin.Context) {
	var itemDto dto.ItemDto
	err := c.BindJSON(&itemDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto, er := itemService.InsertItem(itemDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemDto)
}

func QueueItems(c *gin.Context) {
	var itemsDto dto.ItemsDto
	err := c.BindJSON(&itemsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := itemService.QueueItems(itemsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
}
