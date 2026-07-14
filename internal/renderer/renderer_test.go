package renderer

import (
	"testing"
)

func TestRender(t *testing.T) {
	got := "this is {{name}} bla bla {{count}}"
	want := "this is amir bla bla three"

	data := make(map[string]string)
	data["name"] = "amir"
	data["count"] = "three"

	result := Render(got, data)

	if want != result {
		t.Fatalf("got %q, want %q", result, want)
	}
}
