package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/helpers"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model"
	"net/http"
	"totoval/app/models"
)

type User struct{
	controller.BaseController
}

func (*User) Info(c *gin.Context) {
	userID := helpers.AuthClaimsID(c)
	user := models.User{
		ID: &userID,
	}

	if err := model.H.First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user.Password = nil
	c.JSON(http.StatusOK, gin.H{"data":user})
	return
}

func (*User) AllUser(c *gin.Context){
	user := &models.User{}
	outArr, err := user.ObjArr([]model.Filter{}, []model.Sort{}, 0, false)
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":outArr.([]models.User)})
	return
}

func (*User) PaginateUser(c *gin.Context){
	user := &models.User{}
	pagination, err := user.ObjArrPaginate(c, 25, []model.Filter{}, []model.Sort{}, 0, false)
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"item": pagination.ItemArr(), "totalPage":pagination.LastPage(), "currentPage":pagination.CurrentPage(), "count":pagination.Count(), "total":pagination.Total()}})
	return
}

func (*User) Update(c *gin.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := model.H.First(&user, false); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	name := "t2222es123t"
	modifyUser := models.User{
		Name: &name,
	}
	if err := model.H.Save(&user, modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":user})
	return

	//model.Transaction(func() {
	//	fmt.Println(id)
	//	panic(123)
	//}, 3)
}
func (*User) Delete(c *gin.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := model.H.Delete(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
	return
}
func (*User) DeleteTransaction(c *gin.Context) {
	defer func(){ // handle transaction error
		if err := recover(); err != nil{
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.(error).Error()})
			return
		}
	}()

	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	model.Transaction(func(h *model.Helper) {
		user.SetTX(h.DB()) // important
		if err := model.H.Delete(&user, false); err != nil {
			panic(err)
		}
	}, 1)

	c.JSON(http.StatusOK, gin.H{"data": true})
	return
}
func (*User) Restore(c *gin.Context) {
	var id uint
	id = 14
	modifyUser := models.User{
		ID: &id,
	}

	if err := model.H.Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
	return
}
