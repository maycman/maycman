package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	currentDate := t.Format("2006-01-02 3:4:5 pm")
	thisYear := time.Now().Year()
	startTimeOfThisYear := time.Date(thisYear, time.January, 0, 0, 0, 0, 0, time.UTC).UnixNano()
	endTimeOfThisYear := time.Date(thisYear, time.December, 31, 23, 59, 59, 0, time.UTC).UnixNano()

	progressOfThisYear := float64(time.Now().UnixNano()-startTimeOfThisYear) /
		float64(endTimeOfThisYear-startTimeOfThisYear)

	progressBarCapacity := 30.0
	passedProgressBarIndex := int(math.Floor(progressOfThisYear * progressBarCapacity))

	passedProgressBar := ""
	for i := 0; i < passedProgressBarIndex; i++ {
		passedProgressBar = passedProgressBar + "ā "
	}

	leftProgressBar := ""
	for i := 0; i < int(progressBarCapacity)-int(passedProgressBarIndex); i++ {
		leftProgressBar = leftProgressBar + "ā”"
	}

	file, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	README := "### Hi there š\n\n" +

		"This is Michael (Mike), I'm currently working as Technical Analyst @ Infosys Mexico.\n\n" +

		"About me:\n\n" +

		"- š­ Iām currently working on self side projects.\n" +
		"- š± Iām currently learning JS FrontEnd Dev, IT Security, Go Programming language.\n" +
		"- š¬ Ask me about comics, cars, tech, and Info Security.\n" +

		"- ā” Fun fact: I'm usually listen Cumbias music genre while I'm driving.\n\n" +
		"ā³ Year progress  [" + passedProgressBar + leftProgressBar + "]  " + strconv.FormatFloat(progressOfThisYear*100, 'f', 2, 64) + " %\n\n" +

		"![Top Langs](https://github-readme-stats.vercel.app/api/top-langs/?username=maycman&layout=compact&langs_count=10&theme=dark&hide=html,css)\n\n" +

		"\n\nLast updated: " + currentDate

	_, err = file.WriteString(README)
	if err != nil {
		log.Fatalf("Failed write to file: %s", err)
	}
}
