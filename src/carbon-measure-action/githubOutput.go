package main

import "fmt"

func gitHubOutputVariable(gitHubString string, value string) {
	fmt.Printf("::%v::%v\n", gitHubString, value)
}

func githubDebugMessage(message string) {
	fmt.Printf("::debug::%v\n", message)
}

func githubNoticeMessage(message string) {
	fmt.Printf("::notice::%v\n", message)
}

func githubWarningMessage(message string) {
	fmt.Printf("::warning::%v\n", message)
}

func githubErrorMessage(message string) {
	fmt.Printf("::error::%v\n", message)
}
