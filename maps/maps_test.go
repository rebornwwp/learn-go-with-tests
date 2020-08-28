package maps

import "testing"

func TestSearch(t *testing.T) {
	d := map[string]string{"test": "hello world"}
	got := Search(d, "test")
	expected := "hello world"

	assertString(t, got, expected)
}

func assertString(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %s but got %s", expected, got)
	}
}

func TestDictionarySearch(t *testing.T) {

	t.Run("run with ", func(t *testing.T) {
		d := Dictionary{}
		_ = d.Add("test", "world")
		got, _ := d.Search("test")
		expected := "world"
		assertString(t, got, expected)
	})
}
