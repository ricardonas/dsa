// 27. Remove Element - Two pointers
func removeElement(nums []int, val int) int {
    write := 0

    for read := 0; read < len(nums); read++ {
        if nums[read] != val {
            nums[write] = nums[read]
            write++
        }
    } 

    return write
}
