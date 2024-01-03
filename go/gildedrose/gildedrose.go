package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

var (
	Sulfuras = "Sulfuras, Hand of Ragnaros"
	Brie     = "Aged Brie"
	Passes   = "Backstage passes to a TAFKAL80ETC concert"
)

func UpdateQuality(items []*Item) {
	for _, item := range items {

		updateQuality(item) // update Quality At the end of each day

		updateSellIn(item) // update SellIn At the end of each day

		updateQualityAfterSellInExpiry(item) // update Quality if SellIn has passed expiry At the end of each day

	}
}

func updateQuality(item *Item) {
	switch item.Name {
	case Brie:
		incrQualityByOneWithUpperLimit50(item)
	case Passes:
		incrQualityByOneWithUpperLimit50(item)
		if item.SellIn < 11 {
			incrQualityByOneWithUpperLimit50(item)
		}
		if item.SellIn < 6 {
			incrQualityByOneWithUpperLimit50(item)
		}
	case Sulfuras:
		// do nothing
	default:
		dropQualityByOneWithLowerLimit0(item)
	}
}

func updateSellIn(item *Item) {
	switch item.Name {
	case Sulfuras:
		// do nothing
	default:
		dropSellInByOne(item)
	}
	return
}

func updateQualityAfterSellInExpiry(item *Item) {
	if item.SellIn < 0 {
		switch item.Name {
		case Brie: // increases twice as fast post expiry
			incrQualityByOneWithUpperLimit50(item)
		case Passes:
			item.Quality = 0
		case Sulfuras:
			// do nothing
		default: // decreases twice as fast post expiry
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
