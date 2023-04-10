package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isSpace(ch uint8) bool {
	return unicode.IsSpace(rune(ch))
}

func isOperator(ch uint8) bool {
	return ch == '='
}

func swallowBlank(sql string, start int) int {
	for ; start < len(sql) && isSpace(sql[start]); start++ {

	}
	return start
}

func readWord(sql string, start int) (string, int) {
	var word string
	start = swallowBlank(sql, start)
	var end = start
	if isOperator(sql[end]) {
		word = string(sql[end])
		end++
		return word, end
	}
	for ; end < len(sql) && !isSpace(sql[end]) && !isOperator(sql[end]); end++ {
	}
	//fmt.Printf("word = [%s]\n", sql[start:end])
	word = sql[start:end]
	return word, end
}

func parseToken(sql string) (tokens []string) {
	var start = 0
	var word string
	for word, start = readWord(sql, start); start < len(sql); word, start = readWord(sql, start) {
		tokens = append(tokens, word)
	}
	tokens = append(tokens, word)
	return
}

func parseTableNameAndEventName(tokens []string) (tableName, eventName string) {
	for index := 0; index < len(tokens); index++ {
		if strings.EqualFold(tokens[index], "from") {
			index++
			tableName = tokens[index]
			continue
		}

		if strings.EqualFold(tokens[index], "event") && strings.EqualFold(tokens[index+1], "=") {
			index += 2
			eventName = tokens[index]
			continue
		}
	}
	return
}

func printToken(tableName, eventName string) {
	fmt.Printf("tableName = [%s], eventName = [%s]\n", tableName, eventName)
}

func main() {
	printToken(parseTableNameAndEventName(parseToken("select count(*) from af.ads_sr_af_jpfxd_actuser_state_di partitions(p${YYYYMMDD}) p where event = 'ads_sr_mid_actv_di'")))
	printToken(parseTableNameAndEventName(parseToken("select count(*) from af.ads_sr_af_jpfxd_actuser_state_di partitions(p${YYYYMMDD}) p where event='ads_sr_mid_actv_di'")))
	printToken(parseTableNameAndEventName(parseToken("select count(*) from af.ads_sr_af_jpfxd_actuser_state_di partitions(p${YYYYMMDD}) p")))
	printToken(parseTableNameAndEventName(parseToken("select count(*) from af.af_event_1 partition(p${YYYYMMDD}) b where event='ads_sr_mid_actv_di'")))
}
