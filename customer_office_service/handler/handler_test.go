package handler

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

	tcs := []struct {
		code    int
		handler http.HandlerFunc
	}{
		{http.StatusOK, CreateUser},
		{http.StatusOK, GetUserByID},
		{http.StatusOK, UpdateUser},
		{http.StatusOK, DeleteUser},
		{http.StatusOK, GetAllGoods},
		{http.StatusOK, GetGoodByID},
		{http.StatusCreated, CreateOrder},
		{http.StatusOK, UpdateOrder},
		{http.StatusOK, GetOrderByID},
		{http.StatusOK, DeleteOrder},
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
