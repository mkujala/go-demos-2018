// find second lowest int in slice
// remove dublicates from slice
// sort + revers sort slice
// merge two sorted slices

package main

import (
	"fmt"
	"sort"
)

type intMap map[int]struct{}

func main() {
	slice1 := []int{1, 2, 3, 4, 1, 2, 3, 4}
	slice2 := []int{-9, 0, -2, 7}
	slice3 := []int{1, 2, 1, 4, 1, 2, -9, 0, -2, 7, 1115, 837, 1522, 320, 880, 1932, 1153, 1403, 1340, 1447, 326, 1212, 964, 945, 984, 1331, 1251, 785, 271, 833, 1467, 423, 115, 562, 1798, 522, 612, 321, 1846, 985, 822, 302, 186, 733, 380, 1016, 108, 458, 1900, 138, 1138, 483, 1125, 519, 1360, 450, 1928, 702, 806, 1266, 1333, 1289, 5, 661, 296, 763, 1866, 151, 1432, 611, 157, 401, 1606, 1050, 174, 1569, 1292, 1313, 1571, 1897, 1215, 201, 1076, 1411, 683, 1823, 1171, 1820, 191, 32, 826, 457, 1595, 1173, 1285, 1533, 41, 680, 1165, 362, 659, 1925, 1296, 1941, 922, 1442, 234, 896, 1482, 473, 930, 267, 750, 1294, 285, 1787, 1487, 720, 1760, 765, 1997, 1503, 558, 1106, 1892, 48, 119, 266, 316, 1805, 1785, 1304, 1744, 422, 730, 807, 1062, 535, 938, 1297, 501, 1150, 252, 1964, 1375, 1414, 231, 504, 1963, 1325, 650, 1524, 1977, 1255, 1815, 249, 445, 1792, 1083, 1904, 606, 1114, 1052, 185, 1868, 1260, 1535, 974, 1649, 987, 397, 942, 1342, 1045, 1259, 370, 1745, 274, 1678, 842, 1199, 1973, 149, 875, 444, 1556, 1618, 443, 1211, 1586, 1853, 609, 1316, 16, 1709, 1810, 726, 1929, 1335, 1132, 584, 634, 1717, 1190, 1671, 1627, 619, 882, 469, 1579, 663, 745, 1358, 334, 789, 1449, 337, 600, 728, 490, 724, 1712, 1021, 718, 250, 514, 1385, 1565, 502, 1653, 1563, 594, 1642, 778, 1986, 603, 966, 432, 1095, 796, 1994, 1299, 1983, 844, 418, 1180, 1303, 1392, 804, 1066, 1044, 592, 222, 1143, 118, 729, 1288, 1604, 1703, 170, 1657, 1134, 1509, 1440, 685, 1640, 402, 1366, 589, 1030, 293, 247, 125, 1265, 937, 1311, 541, 1283, 292, 103, 1807, 303, 1806, 1508, 96, 1146, 228, 1073, 1072, 1949, 1126, 1660, 1070, 1516, 478, 899, 1077, 233, 1818, 1568, 622, 671, 1302, 141, 133, 1093, 530, 1927, 719, 53, 197, 278, 82, 34, 1905, 764, 1033, 1938, 830, 1631, 1400, 1444, 240, 684, 576, 494, 868, 1324, 1436, 1961, 1931, 1060, 1771, 1361, 795, 694, 1670, 1071, 283, 27, 799, 1659, 1433, 213, 238, 353, 1263, 427, 815, 1794, 980, 1731, 1975, 1779, 1116, 1510, -664, 713, 859, 1764, 279, 559, 1381, 1354, 1541, 1811, 673, 1976, 758, 566, 1176, 1488, 1517, 371, 1291, 1268, 256, 754, 344, 288, 1935, 291, 372, 356, 993, 1705, 31, 131, 1690, 1101, 400, 1200, 1100, 1742, 1647, 244, 629, 1936, 1592, 1387, 839, 1639, 1491, 1786, 1431, 52, 33, 63, 237, 1246, 1857, 1090, 251, 1404, 1969, 901, 670, 672, 1476, 1264, 856, 678, 1459, 269, 932, 823, 1091, 568, 707, 1797, 1608, -1122, 767, 264, 374, 121, 1943, 1374, 1141, 1673, 25, 1470, 625, 1821, 1629, 51, 1991, 1131, 1661, 716, 1481, 120, 403, 1198, 1023, 1056, 199, 88, 696, 342, 313, 617, 706, 1749, 1546, 1995, 295, 1907, 343, 18, 1711, 45, 395, 1577, 1399, 668, 1484, 8, 282, 1111, 1852, 1182, 1402, 518, 1113, 1108, 1687, 1878, 808, 1599, 1864, 690, 1236, 1154, 595, 1275, 642, 492, 858, 1769, 1019, 891, 1421, 653, 986, 1201, 328, 741, 813, 769}

	s1LowWrong := find2lowest(slice1)
	s2Low := find2lowest(slice2)
	s3Low := find2lowest(slice3)

	fmt.Println("\nSecond lowest items")
	fmt.Println("slice 1 Wrong\t", slice1, "=", s1LowWrong)
	fmt.Println("remove duplicates ->", removeDuplicates(slice1))
	s1Low := find2lowest(removeDuplicates(slice1))
	fmt.Println("slice 1:\t", removeDuplicates(slice1), "=", s1Low)
	fmt.Println("slice 2:\t", slice2, "=", s2Low)
	fmt.Println("slice 3:\t", "[...] =", s3Low)

	fmt.Println("\nMerge sorted slices")
	fmt.Println(mergeSlices(slice1, slice2))

	fmt.Println("\nReverse sort")
	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println("slice 1 reversed:\t", slice1)
	sort.Ints(slice1) // back to normal sort
	fmt.Println("Back to normal sort:\t", slice1)
}

func find2lowest(s []int) int {
	sort.Ints(s)
	item := s[1]
	return item
}

func removeDuplicates(s []int) []int {
	sMap := intMap{}
	for _, number := range s {
		_, ok := sMap[number]
		if !ok {
			sMap[number] = struct{}{}
			//				struct{}  {}
			//				|  ^    | ^
			//				  type    empty element list
		}
	}
	filtered := intMapToSlice(sMap)
	return filtered
}

func intMapToSlice(m intMap) []int {
	intSlice := []int{}
	for key := range m {
		intSlice = append(intSlice, key)
	}
	return intSlice
}

func mergeSlices(s ...[]int) []int {
	merged := []int{}
	for _, slice := range s {
		merged = append(merged, slice...)
	}
	return merged
}
