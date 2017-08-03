package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	rawData, err := base.ParseCSVToInstances("train.csv", true)
	if err != nil {
		panic(err)
	}

	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.10)

	cls := knn.NewKnnClassifier("cosine", "linear", 8)
	cls.Weighted = true

	cls.Fit(trainData)

	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	fmt.Println(predictions)

	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
