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

type Work struct {
	*WorkNumber `xml:"work-number"`
	*WorkTitle  `xml:"work-title"`
	*Opus       `xml:"opus"`
}
type WorkNumber struct{}
type WorkTitle struct{}
type Opus struct{}

type MovementNumber struct{}

type MovementTitle struct{}

type Identification struct {
	*Creator       `xml:"creator"`
	*Rights        `xml:"rights"`
	*Encoding      `xml:"encoding"`
	*Source        `xml:"source"`
	*Relation      `xml:"relation"`
	*Miscellaneous `xml:"miscellaneous"`
}
type Creator struct{}
type Rights struct{}
type Encoding struct{}
type Source struct{}
type Relation struct{}
type Miscellaneous struct{}

type Defaults struct {
	*Scaling       `xml:"scaling"`
	*ConcertScore  `xml:"concert0score"`
	*PageLayout    `xml:"page-layout"`
	*SystemLayout  `xml:"system-layout"`
	*StaffLayout   `xml:"staff-layout"`
	*Appearance    `xml:"appearance"`
	*MusicFont     `xml:"music-font"`
	*WordFont      `xml:"word-font"`
	*LyricFont     `xml:"lyric-font"`
	*LyricLanguage `xml:"lyric-language"`
}
type Scaling struct{}
type ConcertScore struct{}
type PageLayout struct{}
type SystemLayout struct{}
type StaffLayout struct{}
type Appearance struct{}
type MusicFont struct{}
type WordFont struct{}
type LyricFont struct{}
type LyricLanguage struct{}

type Credit struct {
	*Credit       `xml:"credit"`
	*CreditType   `xml:"credit-type"`
	*Link         `xml:"link"`
	*Bookmark     `xml:"bookmark"`
	*CreditImage  `xml:"credit-image"`
	*CreditWords  `xml:"credit-words"`
	*CreditSymbol `xml:"credit-symbol"`
}
type CreditType struct{}
type Link struct{}
type Bookmark struct{}
type CreditImage struct{}
type CreditWords struct{}
type CreditSymbol struct{}

type PartList struct {
	*PartGroup `xml:"part-group"`
	*ScorePart `xml:"score-part"`
}
type PartGroup struct{}
type ScorePart struct{}
