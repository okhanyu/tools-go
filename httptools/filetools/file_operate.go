package filetools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func fileRead(path string) (string, error) {
	start := time.Now()
	fi, _ := os.Open(path)
	read, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	return string(read), nil
}
