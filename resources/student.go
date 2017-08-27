package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/izzulhaziq/profile-service/models"
)

var students map[string]Student

func init() {
	students["1"] = models.Student{
		ID:        "1",
		Name:      "Michelle",
		StudentID: "6441",
		Batch:     "2003",
	}

	students["2"] = models.Student{
		ID:        "2",
		Name:      "Yen Peng",
		StudentID: "3441",
		Batch:     "2001",
	}
}

// StudentResource defines the /student rest api endpoint
type StudentResource struct {
}

// NewStudentResource creates a new endpoint and register to Gin
func NewStudentResource(e *gin.Engine) {
	u := StudentResource{}

	// Setup Routes
	e.GET("/student", u.getAllStudents)
	e.GET("/student/:id", u.getStudentByID)
}

func (r *StudentResource) getAllStudents(c *gin.Context) {

	c.JSON(200, students)
}

func (r *StudentResource) getStudentByID(c *gin.Context) {
	id, err := getStringParam(c, "id")
	if err != nil {
		c.JSON(400, "could not parse id")
		return
	}

	student, err := students[id]
	if err != nil {
		c.JSON(404, "Not Found")
		return
	}

	c.JSON(200, student)
}
