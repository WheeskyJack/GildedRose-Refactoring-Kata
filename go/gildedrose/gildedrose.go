package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		updateQuality(item) // update Quality At the end of each day

		updateSellIn(item) // update SellIn At the end of each day

		updateQualityAfterSellInExpiry(item) // update Quality if SellIn has passed expiry At the end of each day

	}
}

func updateQuality(item *Item) {
	switch item.Name {
	case "Aged Brie":
		incrQualityByOneWithUpperLimit50(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		incrQualityByOneWithUpperLimit50(item)
		if item.SellIn < 11 {
			incrQualityByOneWithUpperLimit50(item)
		}
		if item.SellIn < 6 {
			incrQualityByOneWithUpperLimit50(item)
		}
	case "Sulfuras, Hand of Ragnaros":
		// do nothing
	default:
		dropQualityByOneWithLowerLimit0(item)
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
			incrQualityByOneWithUpperLimit50(item)
		case "Backstage passes to a TAFKAL80ETC concert":
			item.Quality = 0
		case "Sulfuras, Hand of Ragnaros":
			// do nothing
		default:
			dropQualityByOneWithLowerLimit0(item)
		}
	}
}

func dropQualityByOneWithLowerLimit0(item *Item) {
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
}

func incrQualityByOneWithUpperLimit50(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func dropSellInByOne(item *Item) {
	item.SellIn = item.SellIn - 1
}
