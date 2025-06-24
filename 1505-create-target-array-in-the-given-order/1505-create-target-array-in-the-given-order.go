func createTargetArray(nums []int, index []int) []int {
    result := [] int {}
    for k, v := range index {
        if len(result) == index[k] {
            result=append(result,nums[index[k]])
        }else {
            result=append(result,0)
            copy(result[v+1:],result[v:])
            result[index[k]]=nums[k]
        }
    }
    return result ;
}