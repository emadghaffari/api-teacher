package coursecontroller

import (
	"fmt"
	"net/http"

	model "github.com/emadghaffari/api-teacher/model/course"
	service "github.com/emadghaffari/api-teacher/service/course"
	"github.com/emadghaffari/api-teacher/utils/random"
	"github.com/spf13/viper"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router courses = &course{}
)

// courses interface
type courses interface {
	Index(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Take(c *gin.Context)
}

// course struct
type course struct{}

// get all courses
func (u *course) Index(c *gin.Context) {
	// create a new User
	ts, resErr := service.Service.Index()
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"error": resErr.Message()})
		return
	}
	c.JSON(http.StatusOK, ts)
}

// Store new course
func (u *course) Store(c *gin.Context) {
	cs := model.Course{}

	if err := u.bind(c, &cs); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	if err := cs.StoreValidate(); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	cs.Identitiy = fmt.Sprintf("%d", random.Rand(viper.GetInt("course.MinIdentitiy"), viper.GetInt("course.MaxIdentitiy")))

	// create a new Course
	if err := service.Service.Store(&cs); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cs})
}

// Store new course
func (u *course) Update(c *gin.Context) {
	cs := model.Course{}

	if err := u.bind(c, &cs); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	if err := cs.UpdateValidate(); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	// create a new Course
	if err := service.Service.Update(&cs); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cs})
}

// take a course
func (u *course) Take(c *gin.Context) {

}

func (u *course) bind(c *gin.Context, cs *model.Course) errors.ResError {
	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&cs); err != nil {
		return errors.HandlerBadRequest("Invalid JSON Body.")

	}
	return nil
}
