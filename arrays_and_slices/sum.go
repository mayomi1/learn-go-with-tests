package main

// Sum numbers
func Sum(numbers []int) (sum int) {

  for _, number := range numbers {
	  sum += number
  }
  return
}

// SumAll slices
func SumAll(numbersToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numbersToSum {
    sums = append(sums, Sum(numbers))
  }
  return sums
}

// SumAllTails tails
func SumAllTails(numbersToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numbersToSum {
    if len(numbers) == 0 {
      sums = append(sums, 0)
    } else {
      tail := numbers[1:]
      sums = append(sums, Sum(tail))
    }
  
  } 

  return sums
}