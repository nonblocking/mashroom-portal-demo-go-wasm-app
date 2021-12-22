package main

import (
	. "fmt"
	. "github.com/siongui/godom/wasm"
	"os"
	"strconv"
	"syscall/js"
)

var count = 0
var hostElement Value
var messageBusObjName string
var c = make(chan bool)

func onButtonClick(this js.Value, inputs []js.Value) interface{} {
	messageBus := Window.Get(messageBusObjName)

	data := make(map[string]interface{})
	messageBus.Call("publish", "ping", js.ValueOf(data))
	return js.Null()
}

func onPingReceived(this js.Value, inputs []js.Value) interface{} {
	count++
	Println("Current ping count: ", count)
	receivedPingsElement := hostElement.QuerySelector(".number-pings")
	receivedPingsElement.SetInnerHTML("Received pings: " + strconv.Itoa(count))
	return js.Null()
}

func onStop(this js.Value, inputs []js.Value) interface{} {
	c <- true
	return js.Null()
}

func subscribeToPing() {
	messageBus := Window.Get(messageBusObjName)
	messageBus.Call("subscribe", "ping", js.FuncOf(onPingReceived))
}

func subscribeToStop(hostElementId string) {
	messageBus := Window.Get(messageBusObjName)
	messageBus.Call("subscribe", "stop_"+hostElementId, js.FuncOf(onStop))
}

func attachEventHandlers() {
	buttonElement := hostElement.QuerySelector("button")
	buttonElement.Set("onclick", js.FuncOf(onButtonClick))
}

func main() {
	hostElementId := os.Args[0]
	resourcesBasePath := os.Args[1]
	message := os.Args[2]
	pingButtonLabel := os.Args[3]
	messageBusObjName = os.Args[4]

	Println("Got arguments: ", hostElementId, resourcesBasePath, message, pingButtonLabel, messageBusObjName)

	hostElement = Document.GetElementById(hostElementId)
	hostElement.SetInnerHTML(createHtml(resourcesBasePath, message, pingButtonLabel))

	attachEventHandlers()

	subscribeToPing()
	subscribeToStop(hostElementId)

	<-c
	Println("Go app execution stopped")
}
