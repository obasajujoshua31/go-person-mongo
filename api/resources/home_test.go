package resources

import (
	"net/http"
	"testing"
)

func TestWelcomeToAPI(t *testing.T) {
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
			name: "it should print welcome to api",
			args: args{
				w: 
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WelcomeToAPI(tt.args.w, tt.args.r)
		})
	}
}
