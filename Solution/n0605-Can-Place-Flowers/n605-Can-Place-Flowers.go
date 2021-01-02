package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if (i == 0 || flowerbed[i - 1] == 0) && (i == len(flowerbed) - 1 || flowerbed[i + 1] == 0) && flowerbed[i] == 0 {
			n--
			if n == 0 {
				return true
			}
			flowerbed[i] = 1
		}
	}
	return false
}