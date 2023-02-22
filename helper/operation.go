package helper

func SumCalc(arr []float32) float32 {
	total := float32(0)
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return float32(total)
}

func Multiply(arr []float32) float32 {
	total := float32(arr[0])
	for i := 1; i < len(arr); i++ {
		total *= arr[i]
	}
	return float32(total)
}

func Divide(arr []float32) float32 {
	total := float32(arr[0])
	for i := 1; i < len(arr); i++ {
		total /= float32(arr[i])
	}
	return float32(total)
}
func Subtract(arr []float32) float32 {
	total := float32(arr[0])
	for i := 1; i < len(arr); i++ {
		total -= arr[i]
	}
	return float32(total)
}
