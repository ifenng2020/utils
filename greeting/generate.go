package greeting

//go:generate echo $GOOS $GOPATH $GOLINE $GOPACKAGE
func hello() {
	println("Hello generate")
}
