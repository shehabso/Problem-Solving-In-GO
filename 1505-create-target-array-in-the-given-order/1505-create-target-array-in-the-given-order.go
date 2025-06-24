func createTargetArray(nums []int, index []int) []int {
    result := [] int {}
    for k, v := range index {
        if len(result) == index[k] {
            result=append(result,nums[index[k]])
        }else {
            result=append(result,0) // add zero to the end 
            copy(result[v+1:],result[v:])  // shift to the right 
            result[index[k]]=nums[k]
        }
    }
    return result ;
}