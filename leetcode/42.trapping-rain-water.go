func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	// broute force
	//     var ans int
	//     for i := 1; i < n; i++ {
	//         leftMax := 0
	//         rightMax := 0
	//         for j := i; j >= 0; j-- {
	//             if height[j] > leftMax {
	//                 leftMax = height[j]
	//             }
	//         }
	//         for j := i; j < n; j++ {
	//             if height[j] > rightMax {
	//                 rightMax = height[j]
	//             }
	//         }

	//         if leftMax < rightMax && leftMax - height[i] > 0 {
	//             ans += leftMax - height[i]
	//         } else if rightMax < leftMax && rightMax - height[i] > 0 {
	//             ans += rightMax - height[i]
	//         }
	//     }

	// dp
	//     var ans int
	//     leftMax := make([]int, n)
	//     rightMax := make([]int, n)
	//     // leftMax[i]记录从左往右数时，截止下标i处的最大水柱高度值
	//     leftMax[0] = height[0]
	//     for i := 1; i < n; i++ {
	//         if height[i] > leftMax[i-1] {
	//             leftMax[i] = height[i]
	//         } else {
	//             leftMax[i] = leftMax[i-1]
	//         }
	//     }
	//	   // rightMax[i]记录从右往左数时，截止下标i处的最大水柱高度值
	//     rightMax[n-1] = height[n-1]
	//     for i := n-2; i >= 0; i-- {
	//         if height[i] > rightMax[i+1] {
	//             rightMax[i] = height[i]
	//         } else {
	//             rightMax[i] = rightMax[i+1]
	//         }
	//     }

	//     // 核心思想：每个下标i处的能装的水，等于截止到i处，左边最大高度值（包括height[i]）
	//     // 和右边最大高度值（包括height[i]）中较小的那一个
	//     // 减去自身高度height[i]所得结果
	//     for i := 1; i < n-1; i++ {
	//         if leftMax[i] < rightMax[i] && leftMax[i] > height[i] {
	//             ans += leftMax[i] - height[i]
	//         } else if rightMax[i] < leftMax[i] && rightMax[i] > height[i] {
	//             ans += rightMax[i] - height[i]
	//         }
	//     }

	// 双指针
	// 从dp解法可以看出，如果leftMax[i] < rightMax[i]，实际是以leftMax[i]最为最小高度，也就是最小高度肯定在下标i的左边出现，同理对于rightMax[i] < leftMax[i]的，最小高度肯定在右边，不用管左边了
	l := 0
	r := n - 1              // 左右下标
	var lmax, rmax, ans int // 左右方向上的的最大高度值
	for l <= r {
		if height[l] > lmax {
			lmax = height[l]
		}
		if height[r] > rmax {
			rmax = height[r]
		}

		// ans += min(lmax, rmax) - height[i]
		if lmax < rmax {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
		}
	}

	return ans
}