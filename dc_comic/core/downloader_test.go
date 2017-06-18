package core

import (
	"testing"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

func Test_downloadChapter(t *testing.T) {
	type args struct {
		dir      string
		mangaURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test download",
			args:    args{dir: "01", mangaURL: "http://readcomicbooksonline.net/reader/Wonder_Woman_2016/Wonder_Woman_2016_Issue_001"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadChapter(tt.args.dir, tt.args.mangaURL); (err != nil) != tt.wantErr {
				t.Errorf("downloadChapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDownloadEntry(t *testing.T) {
	mangaMap := utility.GetURLTemplate()
	mangaCatlog := "wonder_woman_2016"
	type args struct {
		ch          int
		comicCatlog string
		rootURL     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test download entry",
			args: args{ch: 1, comicCatlog: mangaCatlog, rootURL: mangaMap[mangaCatlog]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownloadEntry(tt.args.ch, tt.args.comicCatlog, tt.args.rootURL)
		})
	}
}
