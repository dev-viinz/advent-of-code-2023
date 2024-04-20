package main

type SeedToSoil struct {
	seed int
	soil int
}
type SoilToFertilizer struct {
	soil       int
	fertilizer int
}

type FertilizerToWater struct {
	fertilizer int
	water      int
}

type WaterToLight struct {
	water int
	light int
}

type LightToTemperature struct {
	light       int
	temperature int
}

type TemperatureToHumidity struct {
	temperature int
	humidity    int
}

type HumidityToLocation struct {
	humidity int
	location int
}
