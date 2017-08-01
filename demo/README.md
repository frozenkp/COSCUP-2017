# Handwriting recognition with KNN

## Preprocessing
### Image Format
- 20 pixel * 20 pixel
- only white and black

![sample image](http://i.imgur.com/qgQVatU.png)

### Split and Zoom
Split the image and zoom to the max size in 20*20 image.
![zoom](http://i.imgur.com/v1039zu.jpg)

### Transfer image to data
- white -> 0
- black -> 1
#### train.csv
```sh
./transfer ./trainData train.csv
```
#### test.csv
```sh
./transfer ./testData test.csv
```

## Model Training
### knn2.go
knn2.go is used for training model, it will split train.csv to `trainData` and `testData` for training.

| distance method | k | train:test | precision |
|:-:|:-:|:-:|:-:|
| cosine | 8 | 9:1 | 0.9677 |
| euclidean | 8 | 9:1 | 0.9035 |
| manhattan | 8 | 9:1 | 0.9032 |

## Prediction
### knn.go
knn.go is used for prediction, it will use train.csv as `trainData` and use test.csv as `testData`.

The output is the predictions of `testData` row by row.

```sh
Optimisations are switched off
Instances with 10 row(s) 1 attribute(s)
Attributes: 
*	FloatAttribute(target)

Data:
	0 
	1 
	2 
	3 
	4 
	5 
	6 
	7 
	6 
	9 
All rows displayed
```
