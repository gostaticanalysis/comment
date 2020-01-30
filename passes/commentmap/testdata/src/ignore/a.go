package a

func f() {
	//lint:ignore check test
	var _ = "" // want "ignore"

	//lint:ignore check1,check,check2 test
	var _ = "" // want "ignore"

	// lint:ignore check test
	var _ = "" // want "ignore"

	//lint:ignore check multiple words in reason
	var _ = "" // want "ignore"

	//lint:ignore check1 test
	var _ = "" // not ignore

	//lint:ignore check1
	var _ = "" // not ignore
}
