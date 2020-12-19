package helper

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	tests := []struct {
		step    string
		error   string
		params  string
		cluster interface{}
	}{
		{
			step:  "a",
			error: "",
			params: `{"password": "123456798","name":"emad","lname":"ghaffari","role": {	"name": "student"}}`,
			cluster: user.User{},
		},
		{
			step:    "b",
			error:   "Invalid JSON Body.",
			params:  ``,
			cluster: user.User{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			err := Bind(c, tc.cluster)

			if err != nil && !strings.Contains(err.Message(), tc.error) {
				assert.Equal(t, err.Message(), tc.error)
			}

		})
	}
}
