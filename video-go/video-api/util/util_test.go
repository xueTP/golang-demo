package util

import (
	"fmt"
	"testing"
)

func Test_getMachineGuid(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "t1", want: "93d959f7-86e3-4a60-bd52-378acd23e9fc"},
	}
	for _, tt := range tests {
		got, err := getMachineGuid()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getMachineGuid() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. getMachineGuid() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetRandStr(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{len: 8}},
	}
	for _, tt := range tests {
		got := GetRandStr(tt.args.len)
		fmt.Println("got: ", got)
		if len(got) != tt.args.len {
			t.Errorf("%q. GetRandStr() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetUUid(t *testing.T) {
	type args struct {
		userPrefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{userPrefix: "xuezd"}},
	}
	for _, tt := range tests {
		if got := GetUUid(tt.args.userPrefix); got != tt.want {
			t.Errorf("%q. GetUUid() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
