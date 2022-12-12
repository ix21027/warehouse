package controller

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"testing"
)

func TestHandlers(t *testing.T) {
	t.Parallel()
	c := &Controller{}

	tcs := []struct {
		code    int
		handler http.HandlerFunc
	}{
		{http.StatusOK, c.CreateUser},
		{http.StatusOK, c.GetUserByID},
		{http.StatusOK, c.UpdateUser},
		{http.StatusOK, c.DeleteUser},
		{http.StatusOK, c.GetAllGoods},
		{http.StatusOK, c.GetGoodByID},
		{http.StatusCreated, c.CreateOrder},
		{http.StatusOK, c.UpdateOrder},
		{http.StatusOK, c.GetOrderByID},
		{http.StatusOK, c.DeleteOrder},
	}

	for _, tc := range tcs {
		tc := tc
		testName := fmt.Sprintf("%s", runtime.FuncForPC(reflect.ValueOf(tc.handler).Pointer()).Name())
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			request, _ := http.NewRequest(http.MethodGet, "", nil)
			response := httptest.NewRecorder()

			tc.handler(response, request)

			assert.Equal(t, tc.code, response.Code)
		})
	}
}
