package telegram

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_urlFor(t *testing.T) {
	type fields struct {
		baseURL    string
		botToken   string
		httpClient *http.Client
	}
	type args struct {
		apiMethod string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantErr    bool
		errMessage string
	}{
		{
			name: "No method",
			fields: fields{
				baseURL:    "foo",
				botToken:   "bar",
				httpClient: &http.Client{},
			},
			args: args{
				apiMethod: "",
			},
			want:       "",
			wantErr:    true,
			errMessage: "apiMethod should not be empty",
		},
		{
			name: "Method is not empty",
			fields: fields{
				baseURL:    "foo",
				botToken:   "bar",
				httpClient: &http.Client{},
			},
			args: args{
				apiMethod: "baz",
			},
			want:    "foobar/baz",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				botToken:   tt.fields.botToken,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.urlFor(tt.args.apiMethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.urlFor() got error = \"%v\"", err)
				return
			}
			if tt.wantErr && (err != nil) && err.Error() != tt.errMessage {
				t.Errorf("Client.urlFor() error = \"%v\", want \"%v\"", err, tt.errMessage)
				return
			}
			if got != tt.want {
				t.Errorf("Client.urlFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		baseURL    string
		botToken   string
		httpClient *http.Client
	}
	tests := []struct {
		name       string
		args       args
		want       *Client
		wantErr    bool
		errMessage string
	}{
		{
			name: "Empty baseURL",
			args: args{
				baseURL:    "",
				botToken:   "foo",
				httpClient: &http.Client{},
			},
			want:       nil,
			wantErr:    true,
			errMessage: "baseURL should not be empty",
		},
		{
			name: "Empty botToken",
			args: args{
				baseURL:    "bar",
				botToken:   "",
				httpClient: &http.Client{},
			},
			want:       nil,
			wantErr:    true,
			errMessage: "botToken should not be empty",
		},
		{
			name: "No httpClient",
			args: args{
				baseURL:    "foo",
				botToken:   "bar",
				httpClient: nil,
			},
			want:       nil,
			wantErr:    true,
			errMessage: "httpClient should not be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.baseURL, tt.args.botToken, tt.args.httpClient)
			if ((err != nil) == tt.wantErr) && (err.Error() != tt.errMessage) {
				t.Errorf("NewClient() error = \"%v\", want \"%v\"", err, tt.errMessage)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
