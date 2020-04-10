//func printToNum(num int) {
//    for i := 1; i <= num; i++ {
//        fmt.Print(i)
//        fmt.Print(" ")
//        if i % 10 == 0 {
//            fmt.Println("\n")
//        }
//    }
//}
//
//func getArraySum(array [5]float64) float64 {
//    var sum float64 = 0 
//
//    for i := 0; i < len(array); i++ {
//        sum += array[i]
//    }
//
//    return sum
//}
//
//func getArrayAverage(sum float64, length int) float64 {
//    return sum / float64(length)
//}
//
//func printInt(value int) {
//    fmt.Println(value)
//}
//
//func printFloat64(value float64) {
//    fmt.Println(value)
//}
//
//func sum(args ...int) int {
//    total := 0
//    for _, num := range args {
//        total += num
//    }
//
//    return total
//}
//
//func makeTopGun(name *string) {
//    fmt.Println("in makeTopGun")
//    *name = "Maverik"
//}
//
//type Pilot struct {
//    firstName string
//    lastName string
//    callsign string
//    aircraft string
//}
//
//func main() {
//    var age int = 21
//    var speed float64 = 100.5
//
//    printInt(age)
//    printFloat64(speed)
//    printToNum(100)
//
//    scores := [5]float64{9, 1.5, 4.5, 7, 8}
//
//    var scores_sum float64 = getArraySum(scores)
//    fmt.Println(scores_sum)
//
//    var scores_length = len(scores)
//    var average float64 = getArrayAverage(scores_sum, scores_length);
//
//    fmt.Println(average)
//
//    pilotAges := make(map[string]int)
//
//    pilotAges["maverik"] = 30
//    pilotAges["goose"] = 27
//    fmt.Println(pilotAges)
//
//    fmt.Println(sum(1, 2, 3, 4))
//
//    var name string = "Goose"
//    fmt.Println(name)
//    makeTopGun(&name)
//    fmt.Println(name)
//
//    p = Pilot{firstName: "Pete", lastName: "Long", callsign: "KAWWWW", aircraft: "Bomber"}
//}

