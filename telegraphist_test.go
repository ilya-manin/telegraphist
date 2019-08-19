package telegraphist

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/xamut/telegraphist/telegram"
)

func TestNew(t *testing.T) {
	type args struct {
		config *Config
	}
	type test struct {
		name    string
		args    args
		want    *telegram.Client
		wantErr bool
	}

	tests := make([]test, 0)

	c1, _ := telegram.NewClient("https://example.com", "foo", &http.Client{Timeout: 10 * time.Second})
	tests = append(tests, test{
		name: "Configured with baseUrl",
		args: args{
			config: &Config{
				BotToken: "foo",
				BaseURL:  "https://example.com",
			},
		},
		want:    c1,
		wantErr: false,
	})

	c2, _ := telegram.NewClient("https://api.telegram.org/bot", "bar", &http.Client{Timeout: 100 * time.Minute})
	tests = append(tests, test{
		name: "Configured with timeout",
		args: args{
			config: &Config{
				BotToken: "bar",
				Timeout:  100 * time.Minute,
			},
		},
		want:    c2,
		wantErr: false,
	})

	c3, _ := telegram.NewClient("https://api.telegram.org/bot", "baz", &http.Client{Timeout: 10 * time.Second})
	tests = append(tests, test{
		name: "Configured with timeout",
		args: args{
			config: &Config{
				BotToken: "baz",
			},
		},
		want:    c3,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
