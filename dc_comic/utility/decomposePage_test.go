package utility

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_decomposePictureURL(t *testing.T) {
	timeoutRequest := http.Client{Timeout: time.Second * 30}
	mangaURL := "http://readcomicbooksonline.net/reader/Wonder_Woman_2016/Wonder_Woman_2016_Issue_001"
	response, err := timeoutRequest.Get(mangaURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteData, _ := ioutil.ReadAll(response.Body)
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
			args: args{page: string(byteData)},
			want: "http://readcomicbooksonline.net/reader/mangas/Wonder Woman 2016/Wonder Woman 2016 Issue 001/jbnythrrp-02-001.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decomposePictureURL(tt.args.page); got != tt.want {
				t.Errorf("decomposePictureURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
