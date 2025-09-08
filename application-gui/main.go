package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(1440, 900, webview.HintNone)
	w.SetHtml("Thanks for using webview!")

	w.Navigate("http://localhost:3002")
	w.Run()
}
