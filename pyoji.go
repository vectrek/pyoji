package pyoji

var COVER = make(map[string]bool)
var ACTIVE = false

func Got(name string) {
	if !ACTIVE {
		return
	}
	COVER[name] = true
}

func Count() int {
	count := 0
	for _ = range COVER {
		count += 1
	}
	return count
}

func Reset() {
	COVER = make(map[string]bool)
}