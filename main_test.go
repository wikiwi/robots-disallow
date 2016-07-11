package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func getFreePort() string {
	l, _ := net.Listen("tcp", "localhost:0")
	defer l.Close()
	return strings.Split(l.Addr().String(), ":")[1]
}

func runServer() {
	main()
	panic("main routine exited")
}

func TestServe(t *testing.T) {
	addr := "localhost:" + getFreePort()
	os.Args = []string{"robots-disallow", "--listen", addr}
	go runServer()

	time.Sleep(time.Second)

	resp, err := http.Get("http://" + addr)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	content := string(bytes)
	if content != "User-agent: *\nDisallow: /\n" {
		t.Fatalf("unexpected content %s", content)
	}
}

func TestVersion(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	output := make(chan string)
	go func() {
		bytes, _ := ioutil.ReadAll(r)
		output <- strings.TrimSpace(string(bytes))
	}()

	os.Args = []string{"robots-disallow", "--version"}
	main()
	w.Close()
	os.Stdout = stdout

	content := <-output
	if content != version {
		t.Fatalf("%q != %q", content, version)
	}
}
