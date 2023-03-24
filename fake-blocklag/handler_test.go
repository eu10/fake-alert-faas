package function

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/", nil),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Handle(tt.args.w, tt.args.r)
		})
	}
}
