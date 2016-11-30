package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"net/http"
	"os"
)

var msg = "<html><body>press here <form action=\"/green\"><input type=\"submit\" value=\"green\"></form><form action=\"/red\"><input type=\"submit\" value=\"red\"></form></body></html>"
var greenmsg = "<html><body>Green Pressed<br/><form action=\"/green\"><input type=\"submit\" value=\"green\"></form><form action=\"/red\"><input type=\"submit\" value=\"red\"></form></body></html>"
var redmsg = "<html><body>Red Pressed<br/><form action=\"/green\"><input type=\"submit\" value=\"green\"></form><form action=\"/red\"><input type=\"submit\" value=\"red\"></form></body></html>"

var (
	// Use mcu pin 18, corresponds to physical pin 23 on the pi
	// Use mcu pin 19, corresponds to physical pin 24 on the pi
	greenpin = rpio.Pin(23)
	redpin   = rpio.Pin(24)
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, msg)
}

func Green(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greenmsg)
	fmt.Println("green pressed")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	redpin.Low()
        greenpin.High()

}
func Red(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, redmsg)
	fmt.Println("red pressed")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	redpin.High()
        greenpin.Low()

}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/green", Green)
	http.HandleFunc("/red", Red)
	http.ListenAndServe(":8080", nil)
}
