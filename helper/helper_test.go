package helper

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w:       httptest.NewRecorder(),
				message: "success response",
				data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessResponse(tt.args.w, tt.args.message, tt.args.data)
		})
	}
}

func TestFailedResponse(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		status int
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w:      httptest.NewRecorder(),
				status: http.StatusBadRequest,
				err:    errors.New("success failed response"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FailedResponse(tt.args.w, tt.args.status, tt.args.err)
		})
	}
}
