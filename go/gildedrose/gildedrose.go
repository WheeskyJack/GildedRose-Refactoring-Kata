package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateQuality(item) // update Quality At the end of each day

		updateSellIn(item) // update SellIn At the end of each day

		updateQualityAfterSellInExpiry(item) // update Quality based on SellIn has passed expiry At the end of each day
	}
}

func updateQuality(item *Item) {
	switch item.Name {
	case "Aged Brie":
		if item.Quality < 50 {
			incrQualityByOne(item)
		}
	case "Backstage passes to a TAFKAL80ETC concert":
		if item.Quality < 50 {
			incrQualityByOne(item)
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
	case "Sulfuras, Hand of Ragnaros":
		// do nothing
	default:
		if item.Quality > 0 {
			dropQualityByOne(item)
		}
	}
}

func updateSellIn(item *Item) {
	switch item.Name {
	case "Sulfuras, Hand of Ragnaros":
		// do nothing
	default:
		dropSellInByOne(item)
	}
	return
}

func updateQualityAfterSellInExpiry(item *Item) {
	if item.SellIn < 0 {
		switch item.Name {
		case "Aged Brie":
			if item.Quality < 50 {
				incrQualityByOne(item)
			}
		case "Backstage passes to a TAFKAL80ETC concert":
			item.Quality = item.Quality - item.Quality
		case "Sulfuras, Hand of Ragnaros":
			// do nothing
		default:
			if item.Quality > 0 {
				dropQualityByOne(item)
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
