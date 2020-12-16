package studentcontroller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emadghaffari/api-teacher/model/course"
	std "github.com/emadghaffari/api-teacher/service/student"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type mockStudent struct {
	MockFunc func() (items course.Courses, err errors.ResError)
}

func (te *mockStudent) Index() (items course.Courses, err errors.ResError) {
	return te.MockFunc()
}

func TestIndexSuccess(t *testing.T) {
	mock := mockStudent{}
	mock.MockFunc = func() (items course.Courses, err errors.ResError) { return course.Courses{}, nil }
	std.Service = &mock

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
	mock := mockStudent{}
	mock.MockFunc = func() (items course.Courses, err errors.ResError) {
		return nil, errors.HandlerInternalServerError("TEST", nil)
	}
	std.Service = &mock

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock HTTP Request and it's return
	c.Request, _ = http.NewRequest("GET", "/teacher", strings.NewReader(string("")))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	// call login func
	Router.Index(c)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

}
