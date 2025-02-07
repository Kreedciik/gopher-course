package model

type Dashboard struct {
	Quantity      int      `json:"quantity"`
	AverageWeight float64  `json:"averageWeight"`
	HungryAnimals []Animal `json:"hungryAnimals"`
	SickAnimals   []Animal `json:"sickAnimals"`
	Food          []Stock  `json:"food"`
	Medicine      []Stock  `json:"medicine"`
}
