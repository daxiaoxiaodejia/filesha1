package main

import (
	"filesha1"
	"filesha1/pkg/log"
)

var (
	fileSha1 *filesha1.FileSha1
	err      error
)

func main() {
	initialize()
	log.Debug("Starting filesha1")

	fileSha1, err = filesha1.NewFileSha1(`{"root":"file","outputFile":"1.txt","exclude":["nihao","hahahah","hahahah/"]}`)
	if err != nil {
		log.Debug(err.Error())
	}
	fileSha1.HandleFilelist()

	fileSha1, err = filesha1.NewFileSha1(`{"root":"file","outputFile":"2.txt","exclude":["pkg","/nism","/*base/"]}`)
	if err != nil {
		log.Debug(err.Error())
	}
	fileSha1.HandleFilelist()

	fileSha1, err = filesha1.NewFileSha1(`{"root":"../../file/","outputFile":"3.txt","exclude":["*.go","pkg","*.log","controllers/"]}`)
	if err != nil {
		log.Debug(err.Error())
	}
	fileSha1.HandleFilelist()
}

//初始化
func initialize() {
	log.NewLogger(0, "console", `{"level": 0, "formatting":true}`)
}
