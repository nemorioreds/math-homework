package main
import "fmt"
func getRandomInt(min int, max int) int {
return min + (max - min) * rand.Float64()
}
func generateCode() []int {
code := make([]int, 0)
for i := 1; i <= 5; i++ {
	if getRandomInt(0, 2) == 0 {
		code = append(code, i)
	} else {
		code = append(code, -i)
	}
}
return code
}
func main() {
fmt.Println(generateCode())
}