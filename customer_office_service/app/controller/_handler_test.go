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
	c := &HttpController{}

	tcs := []struct {
		code    int
		handler http.HandlerFunc
	}{
		{http.StatusOK, c.Service.UserService.CreateUser},
		{http.StatusOK, c.Service.UserService.GetUserByID},
		{http.StatusOK, c.Service.UserService.UpdateUser},
		{http.StatusOK, c.Service.UserService.DeleteUser},
		{http.StatusOK, c.Service.GoodService.GetAllGoods},
		{http.StatusOK, c.Service.GoodService.GetGoodByID},
		{http.StatusCreated, c.Service.OrderService.CreateOrder},
		{http.StatusOK, c.Service.OrderService.UpdateOrder},
		{http.StatusOK, c.Service.OrderService.GetOrderByID},
		{http.StatusOK, c.Service.OrderService.DeleteOrder},
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
