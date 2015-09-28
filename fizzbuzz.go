package main

import "fmt"

func main() {
	var fizz func() string
	var topFizz func() string
	fizz = func() string {
		fizz = func() string {
			fizz = func() string {
				fizz = topFizz
				return "fizz"
			}
			return ""
		}
		return ""
	}
	topFizz = fizz

	var buzz func() string
	var topBuzz func() string
	buzz = func() string {
		buzz = func() string {
			buzz = func() string {
				buzz = func() string {
					buzz = func() string {
						buzz = topBuzz
						return "buzz"
					}
					return ""
				}
				return ""
			}
			return ""
		}
		return ""
	}
	topBuzz = buzz

	for i := 1; i < 100; i++ {
		fmt.Printf("%d: %s%s\n", i, fizz(), buzz())
	}
}
