package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

func parseToArray(textChannel chan string, metarChannel chan []string) {
	for text := range textChannel {
		lines := strings.Split(text, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""
		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}
		metarChannel <- metarSlice
	}
	close(metarChannel)
}

func extractWindDirection(metarChannel chan []string, windsChannel chan []string) {
	for metars := range metarChannel {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}
		windsChannel <- winds
	}
	close(windsChannel)
}

func mineWindDistribution(windsChannel chan []string, distChannel chan [8]int) {
	for winds := range windsChannel {
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windStr, 64); err == nil {
					dirIndex := int(math.Round(d/45.0)) % 8
					windDist[dirIndex]++
				}
			}
		}
	}
	distChannel <- windDist
	close(distChannel)
}

func main() {
	textChannel := make(chan string)
	metarChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultsChannel := make(chan [8]int)

	// Why implement in this way?
	// We could just put all the 3 functions parseToArray, extractWindDirection and mineWindDistribution
	// into the for loop files and process after each file, would dont have the smae peformance than using channals?
	// Well no, wouldnt have a better performance because for each file when the first function parseToArray is processing the functions
	// extractWindDirection and mineWindDistribution are awating the process finish to start them and using channels what happen is we made
	// a process pipeline where for every file in the pipe we will proccess and other proccess will continue processing without to have awating the all process finish.
	// For examples whe start at the file index 0, so we sand to the channel textChannel, the for loop files will look at the next file index 1
	// and while this happening the file index 0 ar processing, so the for loop file add the file index 1 in the channel so mean while the file index 0 is in the  last function process.
	// as you can see we almost eliminated the time awating between proccess, the time where all programa were just awaitng.

	//1. Change to array, each metar report is a separate item in the array
	go parseToArray(textChannel, metarChannel)
	//2. Extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
	go extractWindDirection(metarChannel, windsChannel)
	//3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
	go mineWindDistribution(windsChannel, resultsChannel)

	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := os.ReadDir(absPath)
	start := time.Now()
	for _, file := range files {
		dat, err := os.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(dat)
		textChannel <- text
	}
	close(textChannel)
	results := <-resultsChannel
	elapsed := time.Since(start)
	fmt.Printf("%v\n", results)
	fmt.Printf("Processing took %s\n", elapsed)
}
