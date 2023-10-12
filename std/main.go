package main

import (
	"fmt"
	"net/url"
)

func main() {
	str1 := "select Name,Age from login;"
	str2 := "select sex,Name,Age,Value,message from login;"

	set1 := convertToSet(str1)
	set2 := convertToSet(str2)

	intersection := getIntersection(set1, set2)
	union := getUnion(set1, set2)

	similarity := float64(len(intersection)) / float64(len(union))

	fmt.Printf("Similarity: %.2f%%\n", similarity*100)

	message := "insert+table+jk_mid_tbregister_channel%0A%28dtstatdate%2Cvgameappid%2Cplatid%2Cvopenid%2Clevel%2Cdregdate%2Cregchannel%2Cimei%2Cdtfirsttime%2Cdtfirsteffecttime%29%0ASELECT+%2720230905%27+as+dtstatdate%0A%2Cnvl%28a.vGameAppID%2Cb.vGameAppID%29+as+vGameAppID%0A%2Cnvl%28b.PlatID%2Ca.PlatID%29+as+PlatID%0A%2Cnvl%28a.vOpenID%2Cb.vOpenID%29+as+vOpenID%0A%2Cnvl%28a.level%2Cb.level%29+as+level%0A%2Cif%28b.vGameAppID+is+not+null%2Cb.dregdate%2C%2720230905%27%29+as+dregdate%0A%2Cif%28b.regchannel+is+null+or+b.regchannel+%3D+0%2Cnvl%28a.loginchannel%2C0%29%2Cb.regchannel%29+as+regchannel%0A%2Cif%28b.imei+is+not+null%2Cb.imei%2Ca.imei%29+as+imei%0A%2Cif%28b.vGameAppID+is+not+null%2Cb.dtfirsttime%2Ca.dtstattime%29+as+dtfirsttime%0A%2Cif%28b.regchannel%3D0%2Ca.dtstattime%2Cb.dtfirsteffecttime%29+as+dtfirsteffecttime%0AFROM+%0A%28%09%0A%09select+vgameappid%2Cplatid%2Cvopenid%2Clevel%2Cloginchannel%2Cimei%2Cdtstattime%0A%09from%0A%09%28%09%0A%09select%0A%09%09vgameappid%2Cplatid%2Cvopenid%2Clevel%2Cloginchannel%2Cimei%2Cdtstattime%0A%09%09%2Crow_number%28%29+over%28partition+by+vgameappid%2Cvopenid+order+by+dtstattime%29+rn%0A%09%09from+jk_mid_tblogin_channel%0A%09%09where+dtstatdate%3D%2720230905%27+and+platid+%3C%3E+255%0A%09%29+t%0A%09where+t.rn+%3D+1%0A%29+a+%0AFULL+OUTER+JOIN+%0A%28select+vgameappid%2Cplatid%2Cvopenid%2Clevel%2Cimei%2Cdregdate%2Cregchannel%2Cdtfirsttime%2Cdtfirsteffecttime+%0A--+from+jk_mid_tbregister_channel+where+dtStatDate%3D%2720230904%27%0Afrom+jk_mid_tbregister_channel+where+dtStatDate%3D%2720230904%27+and+platid+in%280%2C1%29%0A%29+b%0AON+a.vGameAppID%3Db.vGameAppID+AND+a.vOpenID%3Db.vOpenID"
	message, _ = url.QueryUnescape(message)
	fmt.Println(message)
}

// Convert string to character set
func convertToSet(str string) map[rune]bool {
	set := make(map[rune]bool)
	for _, char := range str {
		set[char] = true
	}
	return set
}

// Get the intersection of two character sets
func getIntersection(set1, set2 map[rune]bool) map[rune]bool {
	intersection := make(map[rune]bool)
	for char := range set1 {
		if set2[char] {
			intersection[char] = true
		}
	}
	return intersection
}

// Get the union of two character sets
func getUnion(set1, set2 map[rune]bool) map[rune]bool {
	union := make(map[rune]bool)
	for char := range set1 {
		union[char] = true
	}
	for char := range set2 {
		union[char] = true
	}
	return union
}
