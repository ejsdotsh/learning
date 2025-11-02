package binarysearch


func SearchInts(list []int, key int) int {
    low := 0
    mid := 0
    end := len(list) - 1

    for low <= end {
        mid = (low + end) >> 1
        switch {
            case list[mid] > key:
            	end = mid - 1
            case list[mid] < key:
            	low = mid + 1
            default:
            	return mid
        }
    }

    return -1
}
