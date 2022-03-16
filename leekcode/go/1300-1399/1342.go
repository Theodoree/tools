package _300_1399


func numberOfSteps (num int) int {
    var total int


    for num > 0 {
        if num%2 == 0 {
            num/=2
        }else{
            num-=1
        }
        total++
    }


    return  total


}
