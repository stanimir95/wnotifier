package main

//TODO
// FIX ERROR HANDLING ON FLAGS
// ADD SUDO WARNING "NEED SUDO ACCES"
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
	// TODO move to tmp folder
	writeToTmp(writeToTmp1string)

	time.Sleep(time.Duration(checkDiffTimeInterval) * time.Second)
	// TODO move to tmp folder
	writeToTmp(writeToTmp2string)

	readFile("d2.txt")
	compareByteValues(readFile("d1.txt"), readFile("d2.txt"))
}

// MUST BE STRINGS
var fileToWatch string
var userDefinedCommand string
var writeToTmp1string = "d1.txt"
var writeToTmp2string = "d2.txt"

// MUST BE INT
var checkDiffTimeInterval int

func flags() {

	flag.StringVar(&fileToWatch, "f", "", "File To Watch")
	if len(fileToWatch) < 0 {
		fmt.Println("Please select a file to watch")
	}
	flag.StringVar(&userDefinedCommand, "c", "", "Command To Execute")
	flag.IntVar(&checkDiffTimeInterval, "t", 10, "Time Interval (Default is 10 seconds)")
	flag.Parse()

}

//CHECKS last time chosen file is modified
func dateLastModified() string {

	execute, err := exec.Command("date", "-r", fileToWatch).Output()

	if err != nil {
		fmt.Println("No Files found")
	} else {
		fmt.Println("")
	}

	//convert execute to type string
	str := fmt.Sprintf("%s", execute)

	return str

}

// WRITES time to chosen file/s to compare times

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

//check for difference between "before" and "after" times in files

func readFile(fileName string) []byte {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return dat
}

// COMPARES THE VALUES OF BOTH TMP FILES
// used by userDefinedFunction
func compareByteValues(a, b []byte) {
	if reflect.DeepEqual(a, b) {
		fmt.Println("Nothing to do") // CHECK AGAIN
	} else {
		userDefinedFunction()
	}
}

// executes user defined command
func userDefinedFunction() {

	cmd, err := exec.Command("bash", "-c", userDefinedCommand).Output()
	if err != nil {
		fmt.Println(`Please input a command within double quotes "' '" `)
	} else {
		fmt.Println("command:", userDefinedCommand)
	}
	// fmt.Println("command:", userDefinedCommand)
	str := fmt.Sprintf("%s", cmd)
	fmt.Print(str)

}
