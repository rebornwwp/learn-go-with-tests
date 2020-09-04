package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "hhtp://hello"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"hhtp://hello",
	}
	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"hhtp://hello":               false,
	}
	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
