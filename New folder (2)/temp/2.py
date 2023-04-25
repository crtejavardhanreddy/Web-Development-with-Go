class Solution:
    def search(self,nums, target):
        Li=0
        Ri=len(nums)-1
       
        while Li<=Ri:
            # print(1)
            mid=(Li+Ri)//2
            if nums[mid]==target:
                return mid
            elif nums[mid]<target:
                Li=mid+1
            else:
                Ri=mid-1
        return -1
    

s = Solution()
print(s.search([1,2,3,4,56],5))