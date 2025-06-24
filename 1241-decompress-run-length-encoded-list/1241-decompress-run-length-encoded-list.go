func decompressRLElist(nums []int) []int {
    result := [] int {}
    for i:=0; i< len(nums);i+=2 {
        for count :=0; count < nums[i]; count++{
            result = append(result,nums[i+1]);
        }
    }

    return result 
}