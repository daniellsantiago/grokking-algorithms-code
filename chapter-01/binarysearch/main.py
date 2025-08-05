def binary_search(nums: list[int], x: int) -> int:
    l, r = 0, len(nums) - 1
    while l <= r:
        mid = (l + r) // 2
        guess = nums[mid]
        if guess == x:
            return mid
        if guess < x:
            l = mid + 1
        else:
            r = mid - 1
    return -1

print(binary_search([1, 2, 3, 4, 5, 6], 5))