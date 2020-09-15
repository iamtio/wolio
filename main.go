package main

const configFileName string = ".wolioconfig"

// Entry -
type Entry struct {
	Name    string `json:"name"`
	HWAddr  string `json:"hwAddr"`
	UDPPort uint   `json:"port"`
}

func main() {
	app := NewApp()
	app.DrawMenu()
	app.Run()
}
