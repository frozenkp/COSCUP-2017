package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	//"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	trainData, err := base.ParseCSVToInstances("train.csv", true)
	if err != nil {
		panic(err)
	}

	testData, err := base.ParseCSVToInstances("test.csv", true)
	if err != nil {
		panic(err)
	}

        cls := knn.NewKnnClassifier("cosine", "linear", 8)
	cls.Weighted = true

	cls.Fit(trainData)

	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	fmt.Println(predictions)

	/*confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))*/
}
