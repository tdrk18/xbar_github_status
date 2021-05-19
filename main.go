package main

func main() {
	data := getStatus()
	showTopInfo(data.Status)
	showStatus(data)
	showLastUpdate(data.Page)
}
