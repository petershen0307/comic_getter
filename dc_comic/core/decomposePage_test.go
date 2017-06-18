package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

func Test_decomposePictureURL(t *testing.T) {
	pageDOM, _ := getPage("http://readcomicbooksonline.net/reader/Wonder_Woman_2016/Wonder_Woman_2016_Issue_001")
	type args struct {
		page string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// TODO: Add test cases.
			name: "one page",
			args: args{page: pageDOM},
			want: "http://readcomicbooksonline.net/reader/mangas/Wonder Woman 2016/Wonder Woman 2016 Issue 001/jbnythrrp-02-001.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := decomposePictureURL(tt.args.page); got != tt.want {
				t.Errorf("decomposePictureURL() = %v, want %v", got, tt.want)
				t.Errorf("%v", err)
			}
		})
	}
}

func Test_decomposeChapterMaxPage(t *testing.T) {
	pageDOM, _ := getPage("http://readcomicbooksonline.net/reader/Wonder_Woman_2016/Wonder_Woman_2016_Issue_001")
	type args struct {
		page string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "get max page test",
			args:    args{page: pageDOM},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decomposeChapterMaxPage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("decomposeChapterMaxPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decomposeChapterMaxPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decomposeAllChapter(t *testing.T) {
	mangaMap := utility.GetURLTemplate()
	mangaCatlog := "wonder_woman_2016"
	pageDetail, _ := getPage(mangaMap[mangaCatlog])
	type args struct {
		rootPageDetail string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "all chapter", args: args{rootPageDetail: pageDetail}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decomposeAllChapter(tt.args.rootPageDetail)
			fmt.Println(len(got))
			fmt.Println(strings.Join(got, "\n"))
		})
	}
}
