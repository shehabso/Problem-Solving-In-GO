func shuffle(nums []int, n int) []int {
    result := [] int {}
    for i:=0;i<len(nums);i++ {
        if(i+n<len(nums)){
        result =append(result,nums[i]);
        result =append(result,nums[i+n]);
        }
    }  
    return result 
}