package tools

func SafeRun(g func()) {
	defer func() {
		if e := recover(); e != nil {
			println("caught a panic", e)
		}
	}()
	g()
}
