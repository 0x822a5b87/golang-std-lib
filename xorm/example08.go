package main

import "fmt"

func example08() {
	// Sum/SumInt：求某个字段的和，Sum返回float64，SumInt返回int64；
	// Sums/SumsInt：分别求某些字段的和，Sums返回[]float64，SumsInt返回[]int64。
	totalMoney, _ := engine.SumInt(&Sum{}, "money")
	logger.Info(fmt.Sprint("total money:", totalMoney))

	totalRate, _ := engine.Sum(&Sum{}, "rate")
	logger.Info(fmt.Sprint("total rate:", totalRate))

	totals, _ := engine.Sums(&Sum{}, "money", "rate")
	logger.Info(fmt.Sprint("total money:", totals[0], " & total rate: ", totals[1]))
}
