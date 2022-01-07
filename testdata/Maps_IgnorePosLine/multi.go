-- a.go --
package a

func A() {}
-- b.go --
package a

//lint:ignore test-check reason
func B_() {
}
