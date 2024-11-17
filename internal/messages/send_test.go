package messages

import (
	"reflect"
	"testing"
)

func TestMessage_Decode(t *testing.T) {
	type fields struct {
		Site string
		Log  string
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Message
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Success",
			fields: fields{},
			args: args{
				body: []byte(`{"site": "test", "log": "test"}`),
			},
			want: Message{
				Site: "test",
				Log:  "test",
			},
			wantErr: false,
		},
		{
			name:    "Err, Empty Body",
			fields:  fields{},
			args:    args{},
			want:    Message{},
			wantErr: true,
		},
		{
			name:   "Empty Log",
			fields: fields{},
			args: args{
				body: []byte(`{"site": "test"}`),
			},
			want: Message{
				Site: "test",
			},
			wantErr: false,
		},
		{
			name:   "Empty Site",
			fields: fields{},
			args: args{
				body: []byte(`{"log": "test"}`),
			},
			want: Message{
				Log: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				Site: tt.fields.Site,
				Log:  tt.fields.Log,
			}
			got, err := m.Decode(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Encode(t *testing.T) {
	type fields struct {
		Site string
		Log  string
	}
	type args struct {
		site string
		log  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				Site: "test",
				Log:  "test",
			},
			args:    args{},
			want:    []byte(`{"site":"test","log":"test"}`),
			wantErr: false,
		},
		{
			name:    "Err",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Empty Site",
			fields: fields{},
			args: args{
				log: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Empty Log",
			fields: fields{},
			args: args{
				site: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				Site: tt.fields.Site,
				Log:  tt.fields.Log,
			}
			got, err := m.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
