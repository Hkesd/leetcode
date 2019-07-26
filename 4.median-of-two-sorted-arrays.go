package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//lenNums1 := len(nums1)
	//lenNums2 := len(nums2)
	//allLen := lenNums1 + lenNums2
	//
	//aStart, bStart, firstNum, secondNum := 0, 0, 0, 0
	//
	//for i := 0; i <= allLen / 2; i++{
	//	firstNum = secondNum
	//
	//	if aStart < lenNums1 && (bStart >= lenNums2 || nums1[aStart] < nums2[bStart]) {
	//		secondNum = nums1[aStart]
	//		aStart++
	//	} else {
	//		secondNum = nums2[bStart]
	//		bStart++
	//	}
	//}
	//
	//if allLen & 1 == 0 {
	//	return float64(firstNum + secondNum) / 2.0
	//} else {
	//	return float64(secondNum)
	//}

	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	k := lenNums1 + lenNums2
	if k & 1 == 0 {
		return float64(getKth(nums1, 0, lenNums1 - 1, nums2, 0, lenNums2 - 1, k / 2) + getKth(nums1, 0, lenNums1 - 1, nums2, 0, lenNums2 - 1, k / 2 + 1)) * 0.5
	} else {
		return float64(getKth(nums1, 0, lenNums1 - 1, nums2, 0, lenNums2 - 1, k / 2 + 1))
	}

}

func getKth(nums1 []int, start1 int, end1 int, nums2[]int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2 + k - 1]
	}

	if k == 1 {
		return minInt(nums1[start1], nums2[start2])
	}

	i := start1 + minInt(len1, k / 2) - 1
	j := start2 + minInt(len2, k / 2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j + 1, end2, k - (j - start2 + 1))
	} else {
		return getKth(nums1, i + 1, end1, nums2, start2, end2, k - (i - start1 + 1))
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//func main()  {
//	num1 := []int{1, 3}
//	num2 := []int{2}
//
//	re := findMedianSortedArrays(num1, num2)
//
//	fmt.Println(re)
//
//}