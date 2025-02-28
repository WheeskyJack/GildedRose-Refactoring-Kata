package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {

	items, expectedItems := getInputs()

	runTests(t, expectedItems, items)

}

func runTests(t *testing.T, expectedItems []expItems, items []*gildedrose.Item) {
	for _, expectedItem := range expectedItems {
		gildedrose.UpdateQuality(items)
		for i, item := range items {
			if item.Name != expectedItem.items[i].Name {
				t.Errorf("For %s ON %s, Expected Name %s, got %s", expectedItem.items[i].Name, expectedItem.day, expectedItem.items[i].Name, item.Name)
			}
			if item.SellIn != expectedItem.items[i].SellIn {
				t.Errorf("For %s ON %s, Expected SellIn %d, got %d", expectedItem.items[i].Name, expectedItem.day, expectedItem.items[i].SellIn, item.SellIn)
			}
			if item.Quality != expectedItem.items[i].Quality {
				t.Errorf("For %s ON %s, Expected Quality %d, got %d", expectedItem.items[i].Name, expectedItem.day, expectedItem.items[i].Quality, item.Quality)
			}
		}
	}
}

type expItems struct {
	day   string
	items []*gildedrose.Item
}

func getInputs() ([]*gildedrose.Item, []expItems) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	var expectedItems = []expItems{
		{
			"day 1",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 9, 19},
				{"Aged Brie", 1, 1},
				{"Elixir of the Mongoose", 4, 6},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
				{"Conjured Mana Cake", 2, 5},
			},
		},
		{
			"day 2",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 8, 18},
				{"Aged Brie", 0, 2},
				{"Elixir of the Mongoose", 3, 5},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
				{"Backstage passes to a TAFKAL80ETC concert", 8, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
				{"Conjured Mana Cake", 1, 4},
			},
		},
		{
			"day 3",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 7, 17},
				{"Aged Brie", -1, 4},
				{"Elixir of the Mongoose", 2, 4},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
				{"Backstage passes to a TAFKAL80ETC concert", 7, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 2, 50},
				{"Conjured Mana Cake", 0, 3},
			},
		},
		{
			"day 4",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 6, 16},
				{"Aged Brie", -2, 6},
				{"Elixir of the Mongoose", 1, 3},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 11, 24},
				{"Backstage passes to a TAFKAL80ETC concert", 6, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
				{"Conjured Mana Cake", -1, 1},
			},
		},
		{
			"day 5",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 5, 15},
				{"Aged Brie", -3, 8},
				{"Elixir of the Mongoose", 0, 2},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
				{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
				{"Conjured Mana Cake", -2, 0},
			},
		},
		{
			"day 6",
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 4, 14},
				{"Aged Brie", -4, 10},
				{"Elixir of the Mongoose", -1, 0},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 27},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
				{"Conjured Mana Cake", -3, 0},
			},
		},
	}
	return items, expectedItems
}
