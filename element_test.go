package argonaut

import (
	"testing"
)

func TestElement(t *testing.T) {

	item := Elem("item",
		Attr("id"),
		Attr("name"),
	)

	if item.Name != "item" {
		t.Errorf("Unexpected element name: %s\n", item.Name)
	}

	if len(item.attrs) != 2 {
		t.Fatalf("Unexpected length of attributes:", len(item.attrs))
	}

	if item.attrs[0].Name != "id" {
		t.Errorf("Unexpected attribute: %s\n", item.attrs[0].Name)
	}
	if item.attrs[1].Name != "name" {
		t.Errorf("Unexpected attribute: %s\n", item.attrs[1].Name)
	}

}
