package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const startMessage = `[107;40m[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;m([38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;032m([38;5;039m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;039m([38;5;039m([38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;032m([38;5;039m([38;5;039m([38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;m([38;5;m/
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;032m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;038m([38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;039m([38;5;039m([38;5;032m/[38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;015m@[38;5;221m%[38;5;221m#[38;5;221m#[38;5;222m%[38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m(
	[38;5;m [38;5;m [38;5;m [38;5;m([38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m,[38;5;m&[38;5;m*[38;5;m#[38;5;m#[38;5;m [38;5;m [38;5;m [38;5;m.[38;5;m&[38;5;m([38;5;m [38;5;m [38;5;m [38;5;m,[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m/[38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;249m%[38;5;074m([38;5;032m([38;5;032m([38;5;221m#[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m#[38;5;032m([38;5;032m/[38;5;032m/[38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;153m&[38;5;015m@[38;5;221m#[38;5;239m,[38;5;015m@[38;5;239m*[38;5;239m*[38;5;221m#[38;5;038m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;m.
	[38;5;m [38;5;m&[38;5;m&[38;5;m&[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m&[38;5;m&[38;5;m&[38;5;m [38;5;m [38;5;m [38;5;m&[38;5;m([38;5;m,[38;5;m#[38;5;m#[38;5;m#[38;5;m#[38;5;081m#[38;5;081m#[38;5;081m#[38;5;159m&[38;5;123m&[38;5;081m#[38;5;081m#[38;5;081m#[38;5;081m#[38;5;081m%[38;5;123m%[38;5;159m&[38;5;159m&[38;5;159m&[38;5;159m&[38;5;159m&[38;5;152m&[38;5;250m%[38;5;250m%[38;5;250m%[38;5;250m%[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m#[38;5;221m%[38;5;221m%[38;5;221m%[38;5;221m%[38;5;032m/[38;5;038m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;015m@[38;5;015m@[38;5;222m&[38;5;221m#[38;5;221m#[38;5;229m&[38;5;015m@[38;5;032m([38;5;032m([38;5;032m([38;5;039m([38;5;032m([38;5;032m([38;5;032m([38;5;m([38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m*[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;m#[38;5;m#[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m#[38;5;m#[38;5;m#[38;5;m#[38;5;m#[38;5;081m#[38;5;081m#[38;5;081m#[38;5;081m#[38;5;081m#[38;5;247m#[38;5;247m#[38;5;247m#[38;5;247m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;220m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;214m#[38;5;185m#[38;5;032m([38;5;032m([38;5;032m([38;5;032m([38;5;211m%[38;5;211m%[38;5;032m([38;5;032m([38;5;032m([38;5;039m([38;5;032m([38;5;038m([38;5;038m([38;5;254m@[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;m&[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;239m*[38;5;032m([38;5;039m([38;5;032m([38;5;239m*[38;5;239m*[38;5;239m*[38;5;060m*[38;5;039m([38;5;032m([38;5;039m([38;5;176m%[38;5;211m%[38;5;211m%[38;5;211m%[38;5;218m&[38;5;254m@[38;5;254m&[38;5;254m@[38;5;254m&[38;5;254m@[38;5;254m&[38;5;254m@[38;5;254m&[38;5;254m@[38;5;m&[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;032m([38;5;032m([38;5;060m*[38;5;239m*[38;5;239m*[38;5;188m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;218m%[38;5;255m@[38;5;169m#[38;5;169m#[38;5;211m%[38;5;211m%[38;5;211m%[38;5;211m%[38;5;211m%[38;5;m%[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m@[38;5;m.[38;5;m [38;5;m.[38;5;m.[38;5;m [38;5;m [38;5;m [38;5;m#[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m*[38;5;m [38;5;m [38;5;m [38;5;m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;039m([38;5;032m([38;5;039m([38;5;239m*[38;5;239m*[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m@[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m@[38;5;015m@[38;5;015m@[38;5;169m#[38;5;169m#[38;5;169m#[38;5;211m%[38;5;211m%[38;5;211m%[38;5;m%[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m&[38;5;m#[38;5;m,[38;5;045m#[38;5;m&[38;5;m/[38;5;m [38;5;m#[38;5;m%[38;5;m [38;5;123m&[38;5;m&[38;5;123m&[38;5;m%[38;5;045m#[38;5;m#[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m&[38;5;254m@[38;5;254m&[38;5;254m&[38;5;254m&[38;5;m&[38;5;m,[38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m [38;5;m
	[0m`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello - you've requested %s\n", r.URL.Path)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	bindAddr := fmt.Sprintf(":%s", port)

	lines := strings.Split(startMessage, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Printf("==> Server listening at %s", bindAddr)

	err := http.ListenAndServe(bindAddr, nil)
	if err != nil {
		panic(err)
	}
}
