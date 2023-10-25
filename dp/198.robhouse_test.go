package dp

import "testing"

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
*/

func TestMoney(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	res := Money(nums)
	t.Log(res)
}

func Money(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var maxi int
	//dp 每个元素存储打劫这间房屋的金额
	dp := make([]int, len(nums))
	//初始值
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < len(nums); i++ {
		//状态方程
		if dp[i-1] > dp[i-2]+nums[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i]
		}
		//读取最大值
		if dp[i] > maxi {
			maxi = dp[i]
		}
	}
	return maxi
}
