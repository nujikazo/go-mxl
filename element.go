package main

type ScorePartwise struct {
	*Work           `xml:"work"`
	*MovementNumber `xml:"movement-number"`
	*MovementTitle  `xml:"movement-title"`
	*Identification `xml:"identification"`
	*Defaults       `xml:"defaults"`
	*Credit         `xml:"credit"`
	*PartList       `xml:"part-list"`
	Version         string `xml:"version,attr"`
}

type Work struct{}
type MovementNumber struct{}
type MovementTitle struct{}
type Identification struct{}
type Defaults struct{}
type Credit struct{}
type PartList struct{}
