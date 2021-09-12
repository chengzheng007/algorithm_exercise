func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // O(m + n)的解法比较直观，直接merge两个数组，然后求第k大的元素。
    // 不过我们仅仅需要第k大的元素，是不需要“排序”这么复杂的操作的。
    // 可以用一个计数器，记录当前已经找到第m大的元素了。同时我们使用两
    // 个指针pA和pB，分别指向A和B数组的第一个元素，使用类似于
    // merge sort的原理，如果数组A当前元素小，那么pA++，同时m++；如果
    // 数组B当前元素小，那么pB++，同时n++。最终当m等于k的时候，就得到
    // 了我们的答案，O(k)时间，O(1)空间。但是，当k很接近m+n的时候，这个
    // 方法还是O(m+n)的。
    // 有没有更好的方案呢？我们可以考虑从k入手。如果我们每次都能够删除一
    // 个一定在第k大元素之前的元素，那么我们需要进行k次。但是如果每次我
    // 们都删除一半呢？由于A和B都是有序的，我们应该充分利用这里面的信息，
    // 类似于二分查找，也是充分利用了“有序”。
    // 假设A和B的元素个数都大于k/2，我们将A的第k/2个元素（即A[k/2-1]）和B
    // 的第k/2个元素（即 B[k/2-1]）进行比较，有以下三种情况（为了简化这里
    // 先假设k为偶数，所得到的结论对于k是奇数也是成立的）：
    // • A[k/2-1] == B[k/2-1]
    // • A[k/2-1] > B[k/2-1]
    // • A[k/2-1] < B[k/2-1]
    // 如果A[k/2-1] < B[k/2-1]，意味着A[0] 到A[k/2-1]的肯定在A∪B（A
    // 并B）的top k元素的范围内，换句话说，A[k/2-1]不可能大于A∪B的第k大
    // 元素。留给读者证明。
    // 因此，我们可以放心的删除A数组的这k/2个元素。同理，当A[k/2-1] > 
    // B[k/2-1]时，可以删除B数组的k/2个元素。
    // 当A[k/2-1] == B[k/2-1]时，说明找到了第k大的元素，直接返回
    // A[k/2-1] 或 B[k/2-1]即可。
    // 因此，我们可以写一个递归函数。那么函数什么时候应该终止呢？
    // • 当A或B是空时，直接返回B[k-1]或A[k-1]；
    // • 当k=1时，返回min(A[0], B[0])； 
    // • 当A[k/2-1] == B[k/2-1]时，返回A[k/2-1]或B[k/2-1]
    m := len(nums1)
    n := len(nums2)
    total := m+n
    if total & 0x01 == 1 { // 数字总数为奇数个
        return find_kth(nums1, m, nums2, n, total/2+1)
    } else {
        return (find_kth(nums1, m, nums2, n, total/2)+find_kth(nums1, m, nums2, n, total/2+1))/2
    }
}

func find_kth(nums1 []int, m int, nums2 []int, n, k int) float64 {
    // 转为m<n处理
    if m > n {
        return find_kth(nums2, n, nums1, m, k)
    }
    if m == 0 {
        return float64(nums2[k-1])
    }
    if k == 1 {
        if nums1[0] < nums2[0] {
            return float64(nums1[0])
        }
        return float64(nums2[0])
    }
    
    pa := k/2
    if pa > m {
        pa = m
    }
    pb := k - pa
    if nums1[pa-1] < nums2[pb-1] {
        // 剔除nums1的前pa-1部分，变成找第k-pa大小的元素
        return find_kth(nums1[pa:], m-pa, nums2, n, k-pa)
    } else if nums1[pa-1] > nums2[pb-1] {
        // 剔除nums2的前pb-1部分，变成找第k-pb大小的元素
        return find_kth(nums1, m, nums2[pb:], n-pb, k-pb)
    }
    
    return float64(nums1[pa-1])
}
