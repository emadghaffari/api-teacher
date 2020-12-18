package coursecontroller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	model "github.com/emadghaffari/api-teacher/model/course"
	service "github.com/emadghaffari/api-teacher/service/course"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// curse mock
type modelMock struct {
	cs         model.Course
	MockFunc   func() errors.ResError
	MockCourse func() (model.Courses, errors.ResError)
}

func (c *modelMock) Index() (model.Courses, errors.ResError) { return c.MockCourse() }
func (c *modelMock) Store() errors.ResError                  { return c.MockFunc() }
func (c *modelMock) Update() errors.ResError                 { return c.MockFunc() }
func (c *modelMock) Take() errors.ResError                   { return c.MockFunc() }
func (c *modelMock) StoreValidate() errors.ResError          { return c.MockFunc() }
func (c *modelMock) UpdateValidate() errors.ResError         { return c.MockFunc() }
func (c *modelMock) TakeValidate() errors.ResError           { return c.MockFunc() }
func (c *modelMock) Set(u *model.Course)                     { c.cs = *u }
func (c *modelMock) Get() *model.Course                      { return &c.cs }

type serviceMock struct {
	MockFunc   func() errors.ResError
	MockCourse func() (model.Courses, errors.ResError)
}

func (s *serviceMock) Index() (model.Courses, errors.ResError) { return s.MockCourse() }
func (s *serviceMock) Store() errors.ResError                  { return s.MockFunc() }
func (s *serviceMock) Update() errors.ResError                 { return s.MockFunc() }
func (s *serviceMock) Take() errors.ResError                   { return s.MockFunc() }

func init() {
	gin.SetMode(gin.TestMode)
	viper.Set("course.MinIdentitiy", 10000000)
	viper.Set("course.MaxIdentitiy", 990000000)
}

func TestTake(t *testing.T) {
	srv := serviceMock{}
	mdl := modelMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step         string
		status       int
		error        string
		mockFunc     func() errors.ResError
		validateMock func() errors.ResError
		params       string
	}{
		{
			step:         "a",
			status:       http.StatusBadRequest,
			error:        "Invalid JSON Body",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       "",
		},
		{
			step:         "b",
			status:       http.StatusOK,
			error:        "",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       `{"identitiy": "3964731541504"}`,
		},
		{
			step:         "c",
			status:       http.StatusBadRequest,
			error:        "test bad validate",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return errors.HandlerBadRequest("test bad validate") },
			params:       `{"identitiy": "3964731541504"}`,
		},
		{
			step:         "d",
			status:       http.StatusInternalServerError,
			error:        "test server validate",
			mockFunc:     func() errors.ResError { return errors.HandlerInternalServerError("test server validate", nil) },
			validateMock: func() errors.ResError { return nil },
			params:       `{"identitiy": "3964731541504"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			srv.MockFunc = tc.mockFunc
			service.Service = &srv

			mdl.MockFunc = tc.validateMock
			model.Model = &mdl

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/course/take", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			Router.Take(c)

			if tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}

			if tc.status == http.StatusInternalServerError {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}

			if tc.status == http.StatusOK {
				if !strings.Contains(w.Body.String(), "You have successfully registered for the course") {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	srv := serviceMock{}
	mdl := modelMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step         string
		status       int
		error        string
		mockFunc     func() errors.ResError
		validateMock func() errors.ResError
		params       string
	}{
		{
			step:         "a",
			status:       http.StatusBadRequest,
			error:        "Invalid JSON Body",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       "",
		},
		{
			step:         "b",
			status:       http.StatusOK,
			error:        "",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       `{"name": "courese","time": "start at 2020-12-09  end in 2021-12-09","valence": 120,"identitiy": "3767535621042"}`,
		},
		{
			step:         "c",
			status:       http.StatusBadRequest,
			error:        "test bad validate",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return errors.HandlerBadRequest("test bad validate") },
			params:       `{"name": "courese","time": "start at 2020-12-09  end in 2021-12-09","valence": 120,"identitiy": "3767535621042"}`,
		},
		{
			step:         "d",
			status:       http.StatusInternalServerError,
			error:        "test server validate",
			mockFunc:     func() errors.ResError { return errors.HandlerInternalServerError("test server validate", nil) },
			validateMock: func() errors.ResError { return nil },
			params:       `{"name": "courese","time": "start at 2020-12-09  end in 2021-12-09","valence": 120,"identitiy": "3767535621042"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			srv.MockFunc = tc.mockFunc
			service.Service = &srv

			mdl.MockFunc = tc.validateMock
			model.Model = &mdl

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/course/update", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			Router.Update(c)

			if tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}

			if tc.status == http.StatusInternalServerError {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}
		})
	}
}

func TestStore(t *testing.T) {
	srv := serviceMock{}
	mdl := modelMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step         string
		status       int
		error        string
		mockFunc     func() errors.ResError
		validateMock func() errors.ResError
		params       string
	}{
		{
			step:         "a",
			status:       http.StatusBadRequest,
			error:        "Invalid JSON Body",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       "",
		},
		{
			step:         "b",
			status:       http.StatusOK,
			error:        "",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return nil },
			params:       `{"name": "ریاضی 1","description": "class 963 Saturdays","valence": 10,"value": 1,"start_at": "16:00:00","end_at": "19:30:00"}`,
		},
		{
			step:         "c",
			status:       http.StatusBadRequest,
			error:        "test bad validate",
			mockFunc:     func() errors.ResError { return nil },
			validateMock: func() errors.ResError { return errors.HandlerBadRequest("test bad validate") },
			params:       `{"name": "ریاضی 1","description": "class 963 Saturdays","valence": 10,"value": 1,"start_at": "16:00:00","end_at": "19:30:00"}`,
		},
		{
			step:         "d",
			status:       http.StatusInternalServerError,
			error:        "test server validate",
			mockFunc:     func() errors.ResError { return errors.HandlerInternalServerError("test server validate", nil) },
			validateMock: func() errors.ResError { return nil },
			params:       `{"name": "ریاضی 1","description": "class 963 Saturdays","valence": 10,"value": 1,"start_at": "16:00:00","end_at": "19:30:00"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			srv.MockFunc = tc.mockFunc
			service.Service = &srv

			mdl.MockFunc = tc.validateMock
			model.Model = &mdl

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/course/store", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			Router.Store(c)

			if tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}

			if tc.status == http.StatusInternalServerError {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.status, w.Body.String())
				}
			}
		})
	}
}

func TestIndex(t *testing.T) {
	srv := serviceMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step       string
		status     int
		error      string
		mockCourse func() (model.Courses, errors.ResError)
		params     string
	}{
		{
			step:       "a",
			status:     http.StatusOK,
			error:      "",
			mockCourse: func() (model.Courses, errors.ResError) { return model.Courses{&model.Course{ID: 1}}, nil },
			params:     "",
		},
		{
			step:   "b",
			status: http.StatusInternalServerError,
			error:  "TEST ERROR",
			mockCourse: func() (model.Courses, errors.ResError) {
				return nil, errors.HandlerInternalServerError("TEST ERROR", fmt.Errorf("Error"))
			},
			params: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			srv.MockCourse = tc.mockCourse
			service.Service = &srv

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("GET", "/course", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			Router.Index(c)

			if w.Code == http.StatusOK && tc.step == "a" {
				if !strings.Contains(w.Body.String(), `"id":1`) {
					assert.Equal(t, `"id":1`, w.Body.String())
				}
			}

			if w.Code == tc.status && tc.step == "b" {
				if !strings.Contains(w.Body.String(), `TEST ERROR`) {
					assert.Equal(t, `TEST ERROR`, w.Body.String())
				}
			}

		})
	}
}
