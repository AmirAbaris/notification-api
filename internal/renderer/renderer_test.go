package renderer

import (
	"testing"
)

func TestRender(t *testing.T) {
	got := "this is {{name}} bla bla {{count}}"
	want := "this is amir bla bla three"

	data := make(map[string]string)
	data["name"] = "amir"

	result, err := Render(got, data)
	if err != nil || result == "" {
		t.Fatalf("unexpected error: %v", err)
	}

	if want != result {
		t.Fatalf("got %q, want %q", result, want)
	}
}
