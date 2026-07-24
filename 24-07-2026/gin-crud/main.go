package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

var employees = []Employee{
	{
		ID:     1,
		Name:   "Rahul",
		Age:    25,
		Salary: 50000,
	},
	{
		ID:     2,
		Name:   "Teja",
		Age:    22,
		Salary: 60000,
	},
}

func main() {

	router := gin.Default()

	// Home Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin CRUD API",
		})
	})

	// GET All Employees
	router.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	// GET Employee By ID
	router.GET("/employees/:id", func(c *gin.Context) {

		id := c.Param("id")

		for _, emp := range employees {

			if strconv.Itoa(emp.ID) == id {

				c.JSON(http.StatusOK, emp)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Employee Not Found",
		})

	})

	// CREATE Employee
	router.POST("/employees", func(c *gin.Context) {

		var emp Employee

		if err := c.BindJSON(&emp); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		employees = append(employees, emp)

		c.JSON(http.StatusCreated, emp)

	})

	// UPDATE Employee
	router.PUT("/employees/:id", func(c *gin.Context) {

		id := c.Param("id")

		var updatedEmp Employee

		if err := c.BindJSON(&updatedEmp); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		for i, emp := range employees {

			if strconv.Itoa(emp.ID) == id {

				employees[i] = updatedEmp

				c.JSON(http.StatusOK, updatedEmp)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Employee Not Found",
		})

	})

	// DELETE Employee
	router.DELETE("/employees/:id", func(c *gin.Context) {

		id := c.Param("id")

		for i, emp := range employees {

			if strconv.Itoa(emp.ID) == id {

				employees = append(employees[:i], employees[i+1:]...)

				c.JSON(http.StatusOK, gin.H{
					"message": "Employee Deleted Successfully",
				})

				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Employee Not Found",
		})

	})

	// Start Server
	router.Run(":9090")
}