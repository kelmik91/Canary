package canary

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Canary
	}{
		{
			name: "Success",
			args: args{str: "147.45.146.110 - - [20/Apr/2024:00:00:02 +0300] \"GET / HTTP/1.0\" 404 10336 \"-\" \"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36\""},
			want: Canary{
				Method:     "GET",
				URI:        "/",
				StatusCode: "404",
				StrLog:     "147.45.146.110 - - [20/Apr/2024:00:00:02 +0300] \"GET / HTTP/1.0\" 404 10336 \"-\" \"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36\"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
