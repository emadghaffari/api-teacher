package teachercontroller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emadghaffari/api-teacher/model/course"
	tech "github.com/emadghaffari/api-teacher/service/teacher"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type mockTeacher struct {
	MockFunc func() (items course.Courses, err errors.ResError)
}

func (te *mockTeacher) Index() (items course.Courses, err errors.ResError) {
	return te.MockFunc()
}

func TestIndexSuccess(t *testing.T) {
	mock := mockTeacher{}
	mock.MockFunc = func() (items course.Courses, err errors.ResError) { return course.Courses{}, nil }
	tech.Service = &mock

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock HTTP Request and it's return
	c.Request, _ = http.NewRequest("GET", "/teacher", strings.NewReader(string("")))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	// call login func
	Router.Index(c)

	assert.Equal(t, w.Code, http.StatusOK)

}

func TestIndexFail(t *testing.T) {
	mock := mockTeacher{}
	mock.MockFunc = func() (items course.Courses, err errors.ResError) {
		return nil, errors.HandlerInternalServerError("TEST", nil)
	}
	tech.Service = &mock

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock HTTP Request and it's return
	c.Request, _ = http.NewRequest("GET", "/teacher", strings.NewReader(string("")))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	// call login func
	Router.Index(c)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

}
