func maximumWealth(accounts [][]int) int {
    maxWealth :=0
    var  wealth int  
    for _, people:= range accounts{
        wealth=0
        for _ , bank:= range people {
            wealth += bank 
        }
        if(wealth>maxWealth){
            maxWealth=wealth
        }
    
    }   
     return maxWealth
}