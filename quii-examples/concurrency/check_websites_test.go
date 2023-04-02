package concurrency

import (
	"testing"
	"reflection"
)

func mockWebsiteChecker(url string) bool {
	if url == "areyou://sureyouare.mtls" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"http://google.com",
		"http://yahoo.com",
		"areyou://sureyouare.mtls",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://yahoo.com": true,
		"areyou://sureyouare.mtls": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
