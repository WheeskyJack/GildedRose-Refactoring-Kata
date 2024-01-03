package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					dropQualityByOne(items[i])
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							incrQualityByOne(items[i])
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							incrQualityByOne(items[i])
						}
					}
				}
			}
		}

		item := items[i]

		changeSellIn(item)

		updateQualityAfterSellInChange(item)
	}
}

func changeSellIn(item *Item) {
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
