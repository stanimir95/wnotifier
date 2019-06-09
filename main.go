package main

//TODO
// FIX ERROR HANDLING ON FLAGS
// ADD SUDO WARNING
import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"time"
)

func main() {
	commandFlow()

}

func commandFlow() {
	flags()
	for {
		writeToTmp(writeToTmp1string)

		time.Sleep(time.Duration(checkDiffTimeInterval) * time.Second)

		writeToTmp(writeToTmp2string)

		readFile("d2.txt")
		compareByteValues(readFile("d1.txt"), readFile("d2.txt"))
	}
}

// REQUIRED STRINGS
var fileToWatch string
var userDefinedCommand string
var writeToTmp1string = "d1.txt"
var writeToTmp2string = "d2.txt"

// REQUIRED INT for time.Duration()
var checkDiffTimeInterval int

func flags() {

	flag.StringVar(&fileToWatch, "f", fileToWatch, "File To Watch")
	if len(fileToWatch) < 0 {
		fmt.Println("Please select a file to watch")
	}
	flag.StringVar(&userDefinedCommand, "c", userDefinedCommand, "Command To Execute")
	flag.IntVar(&checkDiffTimeInterval, "t", 10, "Time Interval (Default is 10 seconds)")
	flag.Parse()

}

//CHECKS when was the last time the file has been modified
func dateLastModified() string {

	execute, err := exec.Command("date", "-r", fileToWatch).Output()

	if err != nil {
		fmt.Println("No Files found")
	} else {
		fmt.Println("")
	}

	//convert execute(type []byte) to type string
	//
	str := fmt.Sprintf("%s", execute)

	return str

}

// WRITES time of last modification to chosen file

func writeToTmp(file string) {

	date := dateLastModified()
	createFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	writeToFile, _ := createFile.WriteString(date)

	_ = writeToFile

	defer createFile.Close()

}

func readFile(fileName string) []byte {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return dat
}

// COMPARES THE VALUES OF BOTH TMP FILES (d1.txt, d2.txt)
// used by userDefinedFunction
func compareByteValues(a, b []byte) {
	if reflect.DeepEqual(a, b) {
		fmt.Println("Nothing to do")
	} else {
		userDefinedFunction()
	}
}

// executes user defined command
func userDefinedFunction() {

	cmd, err := exec.Command("bash", "-c", userDefinedCommand).CombinedOutput()
	if err != nil {
		fmt.Println("Output: ")
		// fmt.Println(`Please input a command within double quotes "' '" `)
	} else {
		fmt.Println("command: ", cmd)
	}

	str := fmt.Sprintf("%s", cmd)
	fmt.Print(str)

}
