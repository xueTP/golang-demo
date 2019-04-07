package model

import (
	"bou.ke/monkey"
	"golang-demo/video-go/video-api/util"
	"reflect"
	"testing"
	"time"
)

// func TestMain(m *testing.M) {
// 	// 清数据
// 	TruncateTable(UserTableName)
// 	TruncateTable(VideoTableName)
// 	m.Run()
// }

func Test_UserRunOrder(t *testing.T) {
	// 模拟time.Now()
	pg := monkey.Patch(time.Now, func() time.Time { return time.Unix(12345678, 0) })
	defer pg.Unpatch()
	// 清数据
	TruncateTable(UserTableName)
	t.Run("addUser", TestUserModel_AddUser)
	t.Run("getUserInfo", TestUserModel_GetUserInfo)
}

func TestUserModel_AddUser(t *testing.T) {
	type args struct {
		user UserTable
	}
	tests := []struct {
		name    string
		this    UserModel
		args    args
		want    int32
		wantErr bool
	}{
		{name: "t1", args: args{user: UserTable{UName: "小米", PWD: util.GetMD5("123456"), CreateTime: time.Now()}}},
	}
	for _, tt := range tests {
		got, err := tt.this.AddUser(tt.args.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. UserModel.AddUser() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got <= tt.want {
			t.Errorf("%q. UserModel.AddUser() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestUserModel_GetUserInfo(t *testing.T) {
	type args struct {
		user UserTable
	}
	tests := []struct {
		name    string
		this    UserModel
		args    args
		want    UserTable
		wantErr bool
	}{
		{
			name: "t1", args: args{user: UserTable{UID: 1}},
			want: UserTable{UID: 1, UName: "小米", PWD: util.GetMD5("123456"), CreateTime: time.Now().UTC()},
		},
	}
	for _, tt := range tests {
		got, err := tt.this.GetUserInfo(tt.args.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. UserTable.GetUserInfo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. UserTable.GetUserInfo() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
