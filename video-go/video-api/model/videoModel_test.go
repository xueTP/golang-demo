package model

import (
	"bou.ke/monkey"
	"reflect"
	"testing"
	"time"
)

func Test_VideoRunOrder(t *testing.T) {
	// 模拟time.Now()
	pg := monkey.Patch(time.Now, func() time.Time { return time.Unix(12345678, 0) })
	defer pg.Unpatch()
	// 清数据
	TruncateTable(VideoTableName)
	t.Run("addVideo", TestVideoModel_AddVideo)
	t.Run("getContain", TestVideoModel_getCondition)
	t.Run("getOtherSql", TestVideoModel_getOtherSql)
	t.Run("getVideoInfo", TestVideoModel_GetVideoInfo)
	t.Run("getVideoList", TestVideoModel_GetVideoList)
	t.Run("deleteVideo", TestVideoModel_DeleteVideo)
}

func TestVideoModel_AddVideo(t *testing.T) {
	type args struct {
		video VideoTable
	}
	tests := []struct {
		name    string
		this    VideoModel
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "t1", args: args{video: VideoTable{
				VName: "小米发布会", VDesc: "上海体育馆，小米发布会", Path: "2019-03-30/xiaomi.jpg",
			}},
		},
	}
	for _, tt := range tests {
		this := VideoModel{}
		got, err := this.AddVideo(tt.args.video)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. VideoModel.AddVideo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got <= tt.want {
			t.Errorf("%q. VideoModel.AddVideo() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVideoModel_getCondition(t *testing.T) {
	type args struct {
		params *GetVideoParam
	}
	tests := []struct {
		name  string
		this  VideoModel
		args  args
		want  string
		want1 []interface{}
	}{
		{
			name: "t1", args: args{params: &GetVideoParam{}},
			want: "1 = 1",
		},
	}
	for _, tt := range tests {
		this := VideoModel{}
		got, got1 := this.getCondition(tt.args.params)
		if got != tt.want {
			t.Errorf("%q. VideoModel.getCondition() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. VideoModel.getCondition() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestVideoModel_GetVideoList(t *testing.T) {
	type args struct {
		params GetVideoParam
	}
	tests := []struct {
		name    string
		this    VideoModel
		args    args
		want    []VideoTable
		wantErr bool
	}{
		{
			name: "t1", args: args{params: GetVideoParam{Limit: 1}},
			want: []VideoTable{{VID: 1, VName: "小米发布会", VDesc: "上海体育馆，小米发布会", Path: "2019-03-30/xiaomi.jpg", CreateTime: time.Now().UTC(), UpdateTime: time.Now().UTC()}},
		},
	}
	for _, tt := range tests {
		this := VideoModel{}
		got, err := this.GetVideoList(tt.args.params)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. VideoModel.GetVideoList() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. VideoModel.GetVideoList() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVideoModel_getOtherSql(t *testing.T) {
	type args struct {
		params GetVideoParam
	}
	tests := []struct {
		name string
		this VideoModel
		args args
		want string
	}{
		{name: "t1", args: args{params: GetVideoParam{Limit: 1, Offset: 1}}, want: " ORDER BY vid desc LIMIT 1 OFFSET 1"},
	}
	for _, tt := range tests {
		this := VideoModel{}
		if got := this.getOtherSql(tt.args.params); got != tt.want {
			t.Errorf("%q. VideoModel.getOtherSql() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVideoModel_GetVideoInfo(t *testing.T) {
	type args struct {
		params GetVideoParam
	}
	tests := []struct {
		name    string
		this    VideoModel
		args    args
		want    VideoTable
		wantErr bool
	}{
		{name: "t1", args: args{params: GetVideoParam{}}, want: VideoTable{
			VID: 1, VName: "小米发布会", VDesc: "上海体育馆，小米发布会", Path: "2019-03-30/xiaomi.jpg", CreateTime: time.Now().UTC(), UpdateTime: time.Now().UTC(),
		}},
	}
	for _, tt := range tests {
		this := VideoModel{}
		got, err := this.GetVideoInfo(tt.args.params)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. VideoModel.GetVideoInfo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. VideoModel.GetVideoInfo() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVideoModel_DeleteVideo(t *testing.T) {
	type args struct {
		params GetVideoParam
	}
	tests := []struct {
		name    string
		this    VideoModel
		args    args
		want    int32
		wantErr bool
	}{
		{name: "t1", args: args{params: GetVideoParam{VideoTable: VideoTable{VID: 1}}}, want: 0},
	}
	for _, tt := range tests {
		this := VideoModel{}
		got, err := this.DeleteVideo(tt.args.params)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. VideoModel.DeleteVideo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got <= tt.want {
			t.Errorf("%q. VideoModel.DeleteVideo() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
