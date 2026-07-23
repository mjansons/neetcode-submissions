func calPoints(operations []string) int {

scores := []int{}
sum := 0
	for _, v := range operations {
		switch v {
			case "+":
				num := scores[len(scores)-1]+scores[len(scores)-2]
				scores = append(scores, num)
				sum+=num
			case "D":
				num := scores[len(scores)-1]*2
				scores = append(scores, num)
				sum+=num
			case "C":
				num:= scores[len(scores)-1]
				scores = scores[:len(scores)-1]
				sum-=num

			default:
				num, _ := strconv.Atoi(v)
				scores = append(scores, num)
				sum+=num

	}
 }
 return sum
}
