package routes

import (
	"OnlineJudge/app/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine)  {

	// api
	api := router.Group("/api")
	{
		api.GET("/", controller.Index)

		customer := api.Group("/customer")
		{
			customer.POST("/getAllCustomer", controller.FindAllCustomer)
			customer.POST("/addCustomer", controller.AddCustomer)
		}

		staff := api.Group("/staff")
		{
			staff.POST("/getAllStaff", controller.FindAllStaff)
			staff.POST("/addStaff", controller.AddStaff)
		}

		garage := api.Group("/garage")
		{
			garage.POST("/getAllGarage", controller.GetAllGarage)
			garage.GET("/getNoFullGarage", controller.GetNoFullGarage)
			garage.POST("/addGarage", controller.AddGarage)
		}

		model := api.Group("/model")
		{
			model.POST("/getAllModel", controller.GetAllModel)
			model.POST("/addModel", controller.AddModel)
			model.POST("/editModelLevel", controller.EditModelLevel)
		}

		car := api.Group("/car")
		{
			car.POST("/getAllCar", controller.GetAllCar)
			car.POST("/addCar", controller.AddCar)
		}

		level := api.Group("/level")
		{
			level.POST("/getAllLevel", controller.GetAllLevel)
			level.POST("/addLevel", controller.AddLevel)
		}

		bill := api.Group("/bill")
		{
			bill.POST("/getAllRent", controller.GetAllRent)
			bill.POST("/getAllTaken", controller.GetAllTaken)
		}
	}
	router.StaticFS("/public", http.Dir("./web/"))
}