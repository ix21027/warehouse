package controller

import (
	"customer_office_service/app/service"
	"customer_office_service/infrastracture/servers/nats_server"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	t.Parallel()
	natsServer := nats_server.New()
	natsServer.Run()
	defer natsServer.Stop()
	cont := New(service.NewMain(natsServer))

	tcs := []struct {
		code    int
		handler http.HandlerFunc
	}{
		{http.StatusCreated, cont.CreateUser},
		{http.StatusOK, cont.GetUserByID},
		{http.StatusOK, cont.UpdateUser},
		{http.StatusOK, cont.DeleteUser},
		{http.StatusOK, cont.GetAllGoods},
		{http.StatusOK, cont.GetGoodByID},
		{http.StatusCreated, cont.CreateOrder},
		{http.StatusOK, cont.UpdateOrder},
		{http.StatusOK, cont.GetOrderByID},
		{http.StatusOK, cont.DeleteOrder},
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
