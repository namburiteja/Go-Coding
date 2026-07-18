package main
import (
	"os"
)
func main() {
	data := "Neelo valapu anuvule enni ani - neutron,electron nee kannulloney motham undani"
	file,err := os.OpenFile(
		"read.txt",
		os.O_APPEND | os.O_CREATE | os.O_WRONLY,
		0644,
	)
	if err!=nil {
		panic(err)
		return
	}

	defer file.Close()
	file.WriteString("\n")
	file.WriteString(data)
}