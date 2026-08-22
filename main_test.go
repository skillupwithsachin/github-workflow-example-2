package main
import "testing"
func TestGreet(t *testing.T) {
    if Greet("Ada") != "Hello, Ada" { t.Fatal("wrong greeting") }
}
