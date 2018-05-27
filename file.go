package logrz

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	filePath string
	useFile  bool
)

func ensureFile() {
	// detect if file exists
	_, err := os.Stat(filePath)

	// create file if not exists
	if os.IsNotExist(err) {
		err := ioutil.WriteFile(filePath, []byte(""), os.ModePerm)
		if err != nil {
			fmt.Printf("Error while creating file")
		}
	}
}

func flushFile() {
	err := ioutil.WriteFile(filePath, []byte(""), os.ModePerm)
	if err != nil {
		fmt.Printf("Error while rewriting file")
	}
}
func startFileLog() {
	var (
		err      error
		startSeq string
		f        *os.File
	)
	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Unable to open %s for logging purpose", filePath)
	}
	defer f.Close()

	startSeq = fmt.Sprintf("[ %s ]", padRight(time.Now().Format("15:04:05.999"), " ", 12))
	startSeq += " Log start"
	startSeq += "\n"
	if _, err = f.WriteString(startSeq); err != nil {
		fmt.Printf("Unable to write in file")
	}
}

func logFile(str string) {
	if !useFile {
		return
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
		//fmt.Printf("Unable to open %s for logging purpose", filePath)
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		return
		//fmt.Printf("Unable to write in file")
	}
}
