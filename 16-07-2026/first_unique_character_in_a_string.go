func firstUniqChar(s string) int {
    n := len(s)
    hm := make(map[byte]int)
    for i:=0;i<n;i++ {
        char := s[i]
        val,exist := hm[char]
        if exist {
            hm[char] = val+1
        }else {
            hm[char] = 1
        }
    }
    for i:=0;i<n;i++ {
        if hm[s[i]]==1 {
            return i
        }
    }
    return -1
}