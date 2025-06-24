
func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies:= candies[0]
    for i, candy := range candies{
        if(candy>maxCandies){
            maxCandies = candies[i]
        }
    }
     result :=make([]bool,len(candies))
     for i, candy := range candies { 
        result[i] =(candy+extraCandies) >= maxCandies 
     }
     return result 

     }
