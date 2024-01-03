package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateQuality(item)

		updateSellIn(item)

		updateQualityAfterSellInChange(item)
	}
}

func updateQuality(item *Item) {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality > 0 {
			if item.Name != "Sulfuras, Hand of Ragnaros" {
				dropQualityByOne(item)
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						incrQualityByOne(item)
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						incrQualityByOne(item)
					}
				}
			}
		}
	}
}

func updateSellIn(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		dropSellInByOne(item)
	}
}

func updateQualityAfterSellInChange(item *Item) {
	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > 0 {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						dropQualityByOne(item)
					}
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			if item.Quality < 50 {
				incrQualityByOne(item)
			}
		}
	}
}

func dropQualityByOne(item *Item) {
	item.Quality = item.Quality - 1
}

func incrQualityByOne(item *Item) {
	item.Quality = item.Quality + 1
}

func dropSellInByOne(item *Item) {
	item.SellIn = item.SellIn - 1
}
