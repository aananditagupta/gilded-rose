package main

const AGEDBRIE = "Aged Brie"
const BACKSTAGEPASS = "Backstage passes to a TAFKAL80ETC concert"
const SULFARES = "Sulfuras, Hand of Ragnaros"
const CONJURED = "Conjured"

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].name != AGEDBRIE && items[i].name != BACKSTAGEPASS {
			if items[i].name != SULFARES && items[i].name != CONJURED {
				// decrease quality by 1 by none match
				items[i].quality = decreaseInQuality(items[i].quality, 1)
			} else if items[i].name != SULFARES {
				// in the case of conjured items -> the quality decreases twice as fast as normal items
				items[i].quality = decreaseInQuality(items[i].quality, 2)
			}
		} else {
			if items[i].quality < 50 {
				// if name matches any of the above then increase quality
				items[i].quality = items[i].quality + 1
				if items[i].name == BACKSTAGEPASS {
					// Quality increases by 2 when there are 10 days or less
					if items[i].sellIn < 11 {
						items[i].quality = increaseInQuality(items[i].quality)
					}
					//  Quality increases by 3 when there are 5 days or less
					if items[i].sellIn < 6 {
						items[i].quality = increaseInQuality(items[i].quality)
					}
				}
			}
		}

		if items[i].name != SULFARES {
			items[i].sellIn = decreaseInDays(items[i].sellIn)
		}

		if items[i].sellIn < 0 {
			if items[i].name != AGEDBRIE {
				if items[i].name != BACKSTAGEPASS && items[i].name != CONJURED {
					if items[i].name != SULFARES {
						items[i].quality = decreaseInQuality(items[i].quality, 1)
					}
				} else if items[i].name != CONJURED {
					// if sellin date has passed then quality becomes zero including in the case of backstage pass
					items[i].quality = items[i].quality - items[i].quality
				} else {
					// if sellin date has passed then the quality decreases twice as fast as normal items
					items[i].quality = decreaseInQuality(items[i].quality, 2)
				}
			} else {
				// in the case of aged vrie - quality increases once the sell in date has crossed
				items[i].quality = increaseInQuality(items[i].quality)
			}
		}
	}

}

// Break down the big function into smaller functions that do just one thing
func increaseInQuality(quality int) int {
	if quality < 50 {
		quality = quality + 1
	}
	return quality
}

func decreaseInQuality(quality int, decreaseBy int) int {
	if quality > 0 {
		quality = quality - decreaseBy
	}
	return quality
}

func decreaseInDays(sellInDays int) int {
	sellInDays = sellInDays - 1
	return sellInDays
}
