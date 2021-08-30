func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    // 对于负指数优化
    if n < 0 {
        n = -n
        x = 1/x
    }
    
    // 偶数次幂
    if n%2 == 0 {
        return myPow(x*x, n/2)
    }
    
    // 奇数次幂
    return x*myPow(x*x, n/2)
}
