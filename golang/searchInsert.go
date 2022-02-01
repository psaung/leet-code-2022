package main

func searchInsert(nums []int, target int) int {
    i :=0
    for i < len(nums){
        if nums[i]<target{
            i++
            continue;
        } else if nums[i]>target{ 
            break;
        } else if nums[i]==target{
            break;
        }
    }
    return i
}