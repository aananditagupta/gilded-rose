package main

import "testing"

func Test_Foo(t *testing.T) {
	var items = []*Item{
		{"Aged Brie", 0, 2},
		{"Conjured", 10, 2},
		{"Sulfuras", 25, 50},
	}

	UpdateQuality(items)

	if items[0].name != "Aged Brie" {
		t.Errorf("Name: Expected %s but got %s ", "Aged Brie", items[0].name)
	}

	if items[1].sellIn != 9 {
		t.Errorf("sellIn: Expected %s but got %d ", "0", items[1].sellIn)
	}

	if items[2].quality != 49 {
		t.Errorf("Quality: Expected %s but got %d ", "49", items[2].quality)
	}
}
