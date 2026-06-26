package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloHandler(t *testing.T) {
	type want struct {
		code        int
		response    map[string]string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Positive test (200 OK)",
			want: want{
				code:        http.StatusOK,
				response:    map[string]string{"message": "hello, world!"},
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello", nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.want.code, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

			wantBody, err := json.Marshal(tt.want.response)
			if err != nil {
				t.Fatal(err)
			}

			assert.JSONEq(t, string(wantBody), string(resBody))
		})
	}
}
