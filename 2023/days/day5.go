package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"

	"github.com/joshfinnie/advent-of-code/2023/utils"
)

type Almanac struct {
	Seeds                    []int
	SeedToSoilMap            EntryList
	SoilToFertilizerMap      EntryList
	FertilizerToWaterMap     EntryList
	WaterToLightMap          EntryList
	LightToTemperatureMap    EntryList
	TemperatureToHumidityMap EntryList
	HumidityToLocationMap    EntryList
}

func (a Almanac) seedToLocation(seed int) int {
	soil := a.SeedToSoilMap.dest(seed)
	fert := a.SoilToFertilizerMap.dest(soil)
	water := a.FertilizerToWaterMap.dest(fert)
	light := a.WaterToLightMap.dest(water)
	temp := a.LightToTemperatureMap.dest(light)
	humid := a.TemperatureToHumidityMap.dest(temp)
	return a.HumidityToLocationMap.dest(humid)
}

type EntryList struct {
	entries []Entry
}

func (s EntryList) dest(src int) int {
	for _, e := range s.entries {
		if src >= e.src && src < e.src+e.len {
			return e.dest + (src - e.src)
		}
	}
	return src
}

type Entry struct {
	src  int
	dest int
	len  int
}

func parseDay5(file io.Reader) Almanac {
	scanner := bufio.NewScanner(file)
	data := Almanac{}
	for scanner.Scan() {
		// get seeds only
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		data.Seeds = utils.ParseIntArray(strings.Fields(line[7:]))
		break
	}

	if !scanner.Scan() || scanner.Text() != "" {
		panic("AHH")
	}

	var seedToSoilMap EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "seed-to-soil map:" {
			continue
		}
		if line == "" {
			data.SeedToSoilMap = seedToSoilMap
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		seedToSoilMap.entries = append(seedToSoilMap.entries, entry)
	}

	var soilToFertilizerMap EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "soil-to-fertilizer map:" {
			continue
		}
		if line == "" {
			data.SoilToFertilizerMap = soilToFertilizerMap
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		soilToFertilizerMap.entries = append(soilToFertilizerMap.entries, entry)
	}

	var fertilizerToWater EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "fertilizer-to-water map:" {
			continue
		}
		if line == "" {
			data.FertilizerToWaterMap = fertilizerToWater
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		fertilizerToWater.entries = append(fertilizerToWater.entries, entry)
	}

	var waterToLight EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "water-to-light map:" {
			continue
		}
		if line == "" {
			data.WaterToLightMap = waterToLight
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		waterToLight.entries = append(waterToLight.entries, entry)
	}

	var lightToTemperature EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "light-to-temperature map:" {
			continue
		}
		if line == "" {
			data.LightToTemperatureMap = lightToTemperature
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		lightToTemperature.entries = append(lightToTemperature.entries, entry)
	}

	var temperatureToHumidity EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "temperature-to-humidity map:" {
			continue
		}
		if line == "" {
			data.TemperatureToHumidityMap = temperatureToHumidity
			break
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		temperatureToHumidity.entries = append(temperatureToHumidity.entries, entry)
	}

	var humidityToLocation EntryList
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "humidity-to-location map:" {
			continue
		}
		entryArray := utils.ParseIntArray(strings.Fields(line))
		entry := Entry{dest: entryArray[0], src: entryArray[1], len: entryArray[2]}
		humidityToLocation.entries = append(humidityToLocation.entries, entry)
	}

	data.HumidityToLocationMap = humidityToLocation

	return data
}

func Day05A(file io.Reader) string {
	almanac := parseDay5(file)
	lowest := -1
	for _, seed := range almanac.Seeds {
		loc := almanac.seedToLocation(seed)
		if lowest < 0 || loc < lowest {
			lowest = loc
		}
	}

	return strconv.Itoa(lowest)
}

func Day05B(file io.Reader) string {
	almanac := parseDay5(file)
	lowest := -1
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < len(almanac.Seeds); i += 2 {
		wg.Add(1)
		start := almanac.Seeds[i]
		length := almanac.Seeds[i+1]

		go func() {
			defer wg.Done()
			roLowest := -1

			for i := 0; i < length; i++ {
				seed := start + i
				location := seed

				loc := almanac.seedToLocation(location)

				if roLowest < 0 || loc < roLowest {
					roLowest = loc
				}
			}

			mu.Lock()
			defer mu.Unlock()

			if lowest < 0 || roLowest < lowest {
				lowest = roLowest
			}
		}()
	}

	wg.Wait()

	return strconv.Itoa(lowest)
}
