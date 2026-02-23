package main

// Root structures
type ScorePartWise struct {
	Work           *Work           `xml:"work"`
	MovementNumber *MovementNumber `xml:"movement-number"`
	MovementTitle  *MovementTitle  `xml:"movement-title"`
	Identification *Identification `xml:"identification"`
	Defaults       *Defaults       `xml:"defaults"`
	Credit         []Credit        `xml:"credit"`
	PartList       *PartList       `xml:"part-list"`
	Part           []Part          `xml:"part"`
	Version        string          `xml:"version,attr"`
}

type ScoreTimeWise struct {
	Work           *Work           `xml:"work"`
	MovementNumber *MovementNumber `xml:"movement-number"`
	MovementTitle  *MovementTitle  `xml:"movement-title"`
	Identification *Identification `xml:"identification"`
	Defaults       *Defaults       `xml:"defaults"`
	Credit         []Credit        `xml:"credit"`
	PartList       *PartList       `xml:"part-list"`
	Measure        []Measure       `xml:"measure"`
	Version        string          `xml:"version,attr"`
}

// Work-related structures
type Work struct {
	WorkNumber *WorkNumber `xml:"work-number"`
	WorkTitle  *WorkTitle  `xml:"work-title"`
	Opus       *Opus       `xml:"opus"`
}

type WorkNumber struct {
	Value string `xml:",chardata"`
}

type WorkTitle struct {
	Value string `xml:",chardata"`
}

type Opus struct {
	Value string `xml:",chardata"`
}

// Movement-related structures
type MovementNumber struct {
	Value string `xml:",chardata"`
}

type MovementTitle struct {
	Value string `xml:",chardata"`
}

// Identification structures
type Identification struct {
	Creator       []Creator      `xml:"creator"`
	Rights        []Rights       `xml:"rights"`
	Encoding      *Encoding      `xml:"encoding"`
	Source        *Source        `xml:"source"`
	Relation      []Relation     `xml:"relation"`
	Miscellaneous *Miscellaneous `xml:"miscellaneous"`
}

type Creator struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Rights struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Encoding struct {
	Software     []Software     `xml:"software"`
	EncodingDate []EncodingDate `xml:"encoding-date"`
	Encoder      []Encoder      `xml:"encoder"`
	Description  []Description  `xml:"description"`
	Supports     []Supports     `xml:"supports"`
}

type Software struct {
	Value string `xml:",chardata"`
}

type EncodingDate struct {
	Value string `xml:",chardata"`
}

type Encoder struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Description struct {
	Value string `xml:",chardata"`
}

type Supports struct {
	Type      string `xml:"type,attr"`
	Element   string `xml:"element,attr"`
	Attribute string `xml:"attribute,attr,omitempty"`
	Value     string `xml:"value,attr"`
}

type Source struct {
	Value string `xml:",chardata"`
}

type Relation struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Miscellaneous struct {
	MiscellaneousField []MiscellaneousField `xml:"miscellaneous-field"`
}

type MiscellaneousField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

// Defaults structures
type Defaults struct {
	Scaling       *Scaling        `xml:"scaling"`
	ConcertScore  *ConcertScore   `xml:"concert-score"`
	PageLayout    *PageLayout     `xml:"page-layout"`
	SystemLayout  *SystemLayout   `xml:"system-layout"`
	StaffLayout   []StaffLayout   `xml:"staff-layout"`
	Appearance    *Appearance     `xml:"appearance"`
	MusicFont     *MusicFont      `xml:"music-font"`
	WordFont      *WordFont       `xml:"word-font"`
	LyricFont     []LyricFont     `xml:"lyric-font"`
	LyricLanguage []LyricLanguage `xml:"lyric-language"`
}

type Scaling struct {
	Millimeters float32 `xml:"millimeters"`
	Tenths      float32 `xml:"tenths"`
}

type ConcertScore struct {
	// Empty element, presence indicates concert score
}

type PageLayout struct {
	PageHeight  *float32      `xml:"page-height"`
	PageWidth   *float32      `xml:"page-width"`
	PageMargins []PageMargins `xml:"page-margins"`
}

type PageMargins struct {
	Type         string  `xml:"type,attr,omitempty"`
	LeftMargin   float32 `xml:"left-margin"`
	RightMargin  float32 `xml:"right-margin"`
	TopMargin    float32 `xml:"top-margin"`
	BottomMargin float32 `xml:"bottom-margin"`
}

type SystemLayout struct {
	SystemMargins     *SystemMargins  `xml:"system-margins"`
	SystemDistance    *float32        `xml:"system-distance"`
	TopSystemDistance *float32        `xml:"top-system-distance"`
	SystemDividers    *SystemDividers `xml:"system-dividers"`
}

type SystemMargins struct {
	LeftMargin  float32 `xml:"left-margin"`
	RightMargin float32 `xml:"right-margin"`
}

type SystemDividers struct {
	LeftDivider  *SystemDivider `xml:"left-divider"`
	RightDivider *SystemDivider `xml:"right-divider"`
}

type SystemDivider struct {
	Type        string `xml:"type"`
	PrintObject string `xml:"print-object,attr,omitempty"`
}

type StaffLayout struct {
	Number        int      `xml:"number,attr,omitempty"`
	StaffDistance *float32 `xml:"staff-distance"`
}

type Appearance struct {
	LineWidth       []LineWidth       `xml:"line-width"`
	NoteSize        []NoteSize        `xml:"note-size"`
	Distance        []Distance        `xml:"distance"`
	Glyph           []Glyph           `xml:"glyph"`
	OtherAppearance []OtherAppearance `xml:"other-appearance"`
}

type LineWidth struct {
	Type  string  `xml:"type,attr"`
	Value float32 `xml:",chardata"`
}

type NoteSize struct {
	Type  string  `xml:"type,attr"`
	Value float32 `xml:",chardata"`
}

type Distance struct {
	Type  string  `xml:"type,attr"`
	Value float32 `xml:",chardata"`
}

type Glyph struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type OtherAppearance struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type MusicFont struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
}

type WordFont struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
}

type LyricFont struct {
	Number string `xml:"number,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
	Font   Font   `xml:"font"`
}

type Font struct {
	Family string `xml:"font-family,attr,omitempty"`
	Style  string `xml:"font-style,attr,omitempty"`
	Size   string `xml:"font-size,attr,omitempty"`
	Weight string `xml:"font-weight,attr,omitempty"`
}

type LyricLanguage struct {
	Number string `xml:"number,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
	Lang   string `xml:"xml:lang,attr"`
}

// Credit structures
type Credit struct {
	Page         int            `xml:"page,attr,omitempty"`
	CreditType   []CreditType   `xml:"credit-type"`
	Link         []Link         `xml:"link"`
	Bookmark     []Bookmark     `xml:"bookmark"`
	CreditImage  *CreditImage   `xml:"credit-image"`
	CreditWords  []CreditWords  `xml:"credit-words"`
	CreditSymbol []CreditSymbol `xml:"credit-symbol"`
}

type CreditType struct {
	Value string `xml:",chardata"`
}

type Link struct {
	Href     string `xml:"xlink:href,attr"`
	Type     string `xml:"xlink:type,attr,omitempty"`
	Role     string `xml:"xlink:role,attr,omitempty"`
	Title    string `xml:"xlink:title,attr,omitempty"`
	Show     string `xml:"xlink:show,attr,omitempty"`
	Actuate  string `xml:"xlink:actuate,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`
	Element  string `xml:"element,attr,omitempty"`
	Position int    `xml:"position,attr,omitempty"`
}

type Bookmark struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr,omitempty"`
	Element  string `xml:"element,attr,omitempty"`
	Position int    `xml:"position,attr,omitempty"`
}

type CreditImage struct {
	Source string  `xml:"source,attr"`
	Type   string  `xml:"type,attr"`
	Height float32 `xml:"height,attr,omitempty"`
	Width  float32 `xml:"width,attr,omitempty"`
}

type CreditWords struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Justify    string  `xml:"justify,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Enclosure  string  `xml:"enclosure,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type CreditSymbol struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

// Part list structures
type PartList struct {
	PartGroup []PartGroup `xml:"part-group"`
	ScorePart []ScorePart `xml:"score-part"`
}

type PartGroup struct {
	Type                     string                    `xml:"type,attr"`
	Number                   string                    `xml:"number,attr,omitempty"`
	GroupName                *GroupName                `xml:"group-name"`
	GroupNameDisplay         *GroupNameDisplay         `xml:"group-name-display"`
	GroupAbbreviation        *GroupAbbreviation        `xml:"group-abbreviation"`
	GroupAbbreviationDisplay *GroupAbbreviationDisplay `xml:"group-abbreviation-display"`
	GroupSymbol              *GroupSymbol              `xml:"group-symbol"`
	GroupBarline             *GroupBarline             `xml:"group-barline"`
	GroupTime                *GroupTime                `xml:"group-time"`
	Footnote                 *Footnote                 `xml:"footnote"`
	Level                    *Level                    `xml:"level"`
}

type GroupName struct {
	Value string `xml:",chardata"`
}

type GroupNameDisplay struct {
	PrintObject    string           `xml:"print-object,attr,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text"`
	AccidentalText []AccidentalText `xml:"accidental-text"`
}

type GroupAbbreviation struct {
	Value string `xml:",chardata"`
}

type GroupAbbreviationDisplay struct {
	PrintObject    string           `xml:"print-object,attr,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text"`
	AccidentalText []AccidentalText `xml:"accidental-text"`
}

type GroupSymbol struct {
	DefaultX  float32 `xml:"default-x,attr,omitempty"`
	DefaultY  float32 `xml:"default-y,attr,omitempty"`
	RelativeX float32 `xml:"relative-x,attr,omitempty"`
	RelativeY float32 `xml:"relative-y,attr,omitempty"`
	Color     string  `xml:"color,attr,omitempty"`
	Value     string  `xml:",chardata"`
}

type GroupBarline struct {
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

type GroupTime struct {
	// Empty element
}

type Footnote struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Justify    string  `xml:"justify,attr,omitempty"`
	Lang       string  `xml:"xml:lang,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Level struct {
	Reference string `xml:"reference,attr,omitempty"`
	Value     string `xml:",chardata"`
}

type ScorePart struct {
	ID                      string                   `xml:"id,attr"`
	Identification          *Identification          `xml:"identification"`
	PartName                PartName                 `xml:"part-name"`
	PartNameDisplay         *PartNameDisplay         `xml:"part-name-display"`
	PartAbbreviation        *PartAbbreviation        `xml:"part-abbreviation"`
	PartAbbreviationDisplay *PartAbbreviationDisplay `xml:"part-abbreviation-display"`
	Group                   []Group                  `xml:"group"`
	ScoreInstrument         []ScoreInstrument        `xml:"score-instrument"`
	Player                  []Player                 `xml:"player"`
	MidiDevice              []MidiDevice             `xml:"midi-device"`
	MidiInstrument          []MidiInstrument         `xml:"midi-instrument"`
	PartLink                []PartLink               `xml:"part-link"`
}

type PartName struct {
	PrintObject string `xml:"print-object,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type PartNameDisplay struct {
	PrintObject    string           `xml:"print-object,attr,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text"`
	AccidentalText []AccidentalText `xml:"accidental-text"`
}

type PartAbbreviation struct {
	PrintObject string `xml:"print-object,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type PartAbbreviationDisplay struct {
	PrintObject    string           `xml:"print-object,attr,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text"`
	AccidentalText []AccidentalText `xml:"accidental-text"`
}

type Group struct {
	Value string `xml:",chardata"`
}

type ScoreInstrument struct {
	ID                string             `xml:"id,attr"`
	InstrumentName         string             `xml:"instrument-name"`
	InstrumentAbbreviation string             `xml:"instrument-abbreviation,omitempty"`
	InstrumentSound        string             `xml:"instrument-sound,omitempty"`
	Solo              *Solo              `xml:"solo,omitempty"`
	Ensemble          *int               `xml:"ensemble,omitempty"`
	VirtualInstrument *VirtualInstrument `xml:"virtual-instrument,omitempty"`
}

type Solo struct {
	// Empty element
}

type VirtualInstrument struct {
	VirtualLibrary string `xml:"virtual-library,omitempty"`
	VirtualName    string `xml:"virtual-name,omitempty"`
}

type MidiDevice struct {
	ID    string `xml:"id,attr,omitempty"`
	Port  int    `xml:"port,attr,omitempty"`
	Value string `xml:",chardata"`
}

type MidiInstrument struct {
	ID            string   `xml:"id,attr"`
	MidiChannel   *int     `xml:"midi-channel,omitempty"`
	MidiName      *string  `xml:"midi-name,omitempty"`
	MidiBank      *int     `xml:"midi-bank,omitempty"`
	MidiProgram   *int     `xml:"midi-program,omitempty"`
	MidiUnpitched *int     `xml:"midi-unpitched,omitempty"`
	Volume        *float32 `xml:"volume,omitempty"`
	Pan           *float32 `xml:"pan,omitempty"`
	Elevation     *float32 `xml:"elevation,omitempty"`
}

// Part structures
type Part struct {
	ID      string    `xml:"id,attr"`
	Measure []Measure `xml:"measure"`
}

type Measure struct {
	Number         string      `xml:"number,attr"`
	Width          float32     `xml:"width,attr,omitempty"`
	Implicit       string      `xml:"implicit,attr,omitempty"`
	NonControlling string      `xml:"non-controlling,attr,omitempty"`
	Print          *Print      `xml:"print"`
	Attributes     *Attributes `xml:"attributes"`
	Backup         []Backup    `xml:"backup"`
	Forward        []Forward   `xml:"forward"`
	Direction      []Direction `xml:"direction"`
	Note           []Note      `xml:"note"`
	Figured        []Figured   `xml:"figured-bass"`
	Harmony        []Harmony   `xml:"harmony"`
	Sound          []Sound     `xml:"sound"`
	Barline        []Barline   `xml:"barline"`
	Grouping       []Grouping  `xml:"grouping"`
	Link           []Link      `xml:"link"`
	Bookmark       []Bookmark  `xml:"bookmark"`
}

// Note structures (simplified - this is one of the most complex parts)
type Note struct {
	Grace            *Grace            `xml:"grace"`
	Cue              *Cue              `xml:"cue"`
	Chord            *Chord            `xml:"chord"`
	Pitch            *Pitch            `xml:"pitch"`
	Unpitched        *Unpitched        `xml:"unpitched"`
	Rest             *Rest             `xml:"rest"`
	Duration         *int              `xml:"duration"`
	Tie              []Tie             `xml:"tie"`
	Instrument       *Instrument       `xml:"instrument"`
	Voice            *int              `xml:"voice"`
	Type             *NoteType         `xml:"type"`
	Dot              []Dot             `xml:"dot"`
	Accidental       *Accidental       `xml:"accidental"`
	TimeModification *TimeModification `xml:"time-modification"`
	Stem             *Stem             `xml:"stem"`
	Notehead         *Notehead         `xml:"notehead"`
	NoteheadText     *NoteheadText     `xml:"notehead-text"`
	Staff            *int              `xml:"staff"`
	Beam             []Beam            `xml:"beam"`
	Notations        []Notations       `xml:"notations"`
	Lyric            []Lyric           `xml:"lyric"`
	Play             *Play             `xml:"play"`
	Listen           *Listen           `xml:"listen"`
}

type Grace struct {
	StealTimePrevious  float32 `xml:"steal-time-previous,attr,omitempty"`
	StealTimeFollowing float32 `xml:"steal-time-following,attr,omitempty"`
	MakeTime           float32 `xml:"make-time,attr,omitempty"`
	Slash              string  `xml:"slash,attr,omitempty"`
}

type Chord struct {
	// Empty element
}

type Pitch struct {
	Step   string   `xml:"step"`
	Alter  *float32 `xml:"alter,omitempty"`
	Octave int      `xml:"octave"`
}

type Unpitched struct {
	DisplayStep   string `xml:"display-step"`
	DisplayOctave int    `xml:"display-octave"`
}

type Rest struct {
	DisplayStep   *string `xml:"display-step,omitempty"`
	DisplayOctave *int    `xml:"display-octave,omitempty"`
	Measure       string  `xml:"measure,attr,omitempty"`
}

type Tie struct {
	Type string `xml:"type,attr"`
}

type Instrument struct {
	ID string `xml:"id,attr"`
}

type NoteType struct {
	Size  string `xml:"size,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Dot struct {
	// Empty element or with print-style and placement
}

type Accidental struct {
	Cautionary  string `xml:"cautionary,attr,omitempty"`
	Editorial   string `xml:"editorial,attr,omitempty"`
	Parentheses string `xml:"parentheses,attr,omitempty"`
	Bracket     string `xml:"bracket,attr,omitempty"`
	Size        string `xml:"size,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TimeModification struct {
	ActualNotes int     `xml:"actual-notes"`
	NormalNotes int     `xml:"normal-notes"`
	NormalType  *string `xml:"normal-type,omitempty"`
	NormalDot   []Dot   `xml:"normal-dot,omitempty"`
}

type Stem struct {
	DefaultX  float32 `xml:"default-x,attr,omitempty"`
	DefaultY  float32 `xml:"default-y,attr,omitempty"`
	RelativeX float32 `xml:"relative-x,attr,omitempty"`
	RelativeY float32 `xml:"relative-y,attr,omitempty"`
	Color     string  `xml:"color,attr,omitempty"`
	Value     string  `xml:",chardata"`
}

type Notehead struct {
	Filled      string `xml:"filled,attr,omitempty"`
	Parentheses string `xml:"parentheses,attr,omitempty"`
	Color       string `xml:"color,attr,omitempty"`
	FontFamily  string `xml:"font-family,attr,omitempty"`
	FontStyle   string `xml:"font-style,attr,omitempty"`
	FontSize    string `xml:"font-size,attr,omitempty"`
	FontWeight  string `xml:"font-weight,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Beam struct {
	Number   int    `xml:"number,attr,omitempty"`
	Repeater string `xml:"repeater,attr,omitempty"`
	Fan      string `xml:"fan,attr,omitempty"`
	Color    string `xml:"color,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Notations struct {
	Tied           []Tied           `xml:"tied"`
	Slur           []Slur           `xml:"slur"`
	Tuplet         []Tuplet         `xml:"tuplet"`
	Glissando      []Glissando      `xml:"glissando"`
	Slide          []Slide          `xml:"slide"`
	Ornaments      *Ornaments       `xml:"ornaments"`
	Technical      *Technical       `xml:"technical"`
	Articulations  *Articulations   `xml:"articulations"`
	Dynamics       *Dynamics        `xml:"dynamics"`
	Fermata        []Fermata        `xml:"fermata"`
	Arpeggiate     *Arpeggiate      `xml:"arpeggiate"`
	NonArpeggiate  *NonArpeggiate   `xml:"non-arpeggiate"`
	AccidentalMark []AccidentalMark `xml:"accidental-mark"`
	OtherNotation  []OtherNotation  `xml:"other-notation"`
}

// ... Additional notation types would go here

type Lyric struct {
	Number       string        `xml:"number,attr,omitempty"`
	Name         string        `xml:"name,attr,omitempty"`
	Justify      string        `xml:"justify,attr,omitempty"`
	DefaultX     float32       `xml:"default-x,attr,omitempty"`
	DefaultY     float32       `xml:"default-y,attr,omitempty"`
	RelativeX    float32       `xml:"relative-x,attr,omitempty"`
	RelativeY    float32       `xml:"relative-y,attr,omitempty"`
	Placement    string        `xml:"placement,attr,omitempty"`
	Color        string        `xml:"color,attr,omitempty"`
	PrintObject  string        `xml:"print-object,attr,omitempty"`
	Syllabic     *Syllabic     `xml:"syllabic"`
	Text         Text          `xml:"text"`
	Elision      []Elision     `xml:"elision"`
	Extend       *Extend       `xml:"extend"`
	Laughing     *Laughing     `xml:"laughing"`
	Humming      *Humming      `xml:"humming"`
	EndLine      *EndLine      `xml:"end-line"`
	EndParagraph *EndParagraph `xml:"end-paragraph"`
	Footnote     *Footnote     `xml:"footnote"`
	Level        *Level        `xml:"level"`
}

type Syllabic struct {
	Value string `xml:",chardata"`
}

type Text struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type Elision struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type Extend struct {
	Type  string `xml:"type,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

type Laughing struct {
	// Empty element
}

type Humming struct {
	// Empty element
}

type EndLine struct {
	// Empty element
}

type EndParagraph struct {
	// Empty element
}

type Play struct {
	ID               string            `xml:"id,attr,omitempty"`
	Ipa              string            `xml:"ipa,omitempty"`
	Mute             string            `xml:"mute,omitempty"`
	SemiPitched      string            `xml:"semi-pitched,omitempty"`
	OtherPlay        *OtherPlay        `xml:"other-play"`
	InstrumentChange *InstrumentChange `xml:"instrument-change"`
}

// Print structures
type Print struct {
	StaffSpacing            *float32                 `xml:"staff-spacing,attr,omitempty"`
	NewSystem               string                   `xml:"new-system,attr,omitempty"`
	NewPage                 string                   `xml:"new-page,attr,omitempty"`
	BlankPage               int                      `xml:"blank-page,attr,omitempty"`
	PageNumber              string                   `xml:"page-number,attr,omitempty"`
	PageLayout              *PageLayout              `xml:"page-layout"`
	SystemLayout            *SystemLayout            `xml:"system-layout"`
	StaffLayout             []StaffLayout            `xml:"staff-layout"`
	MeasureLayout           *MeasureLayout           `xml:"measure-layout"`
	MeasureNumbering        *MeasureNumbering        `xml:"measure-numbering"`
	PartNameDisplay         *PartNameDisplay         `xml:"part-name-display"`
	PartAbbreviationDisplay *PartAbbreviationDisplay `xml:"part-abbreviation-display"`
}

type MeasureLayout struct {
	MeasureDistance *float32 `xml:"measure-distance"`
}

type MeasureNumbering struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

// Attributes structures
type Attributes struct {
	Divisions    *int           `xml:"divisions"`
	Key          []Key          `xml:"key"`
	Time         []Time         `xml:"time"`
	Staves       *int           `xml:"staves"`
	PartSymbol   *PartSymbol    `xml:"part-symbol"`
	Instruments  *int           `xml:"instruments"`
	Clef         []Clef         `xml:"clef"`
	StaffDetails []StaffDetails `xml:"staff-details"`
	Transpose    []Transpose    `xml:"transpose"`
	ForPart      *ForPart       `xml:"for-part"`
	Directive    []Directive    `xml:"directive"`
	MeasureStyle []MeasureStyle `xml:"measure-style"`
}

type Key struct {
	Number        int             `xml:"number,attr,omitempty"`
	PrintObject   string          `xml:"print-object,attr,omitempty"`
	Fifths        *int            `xml:"fifths"`
	Mode          *string         `xml:"mode"`
	Cancel        *Cancel         `xml:"cancel"`
	KeyStep       []KeyStep       `xml:"key-step"`
	KeyAlter      []KeyAlter      `xml:"key-alter"`
	KeyAccidental []KeyAccidental `xml:"key-accidental"`
	KeyOctave     []KeyOctave     `xml:"key-octave"`
}

type Cancel struct {
	Location string `xml:"location,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

type KeyStep struct {
	Value string `xml:",chardata"`
}

type KeyAlter struct {
	Value float32 `xml:",chardata"`
}

type KeyAccidental struct {
	Value string `xml:",chardata"`
}

type Time struct {
	Number          int              `xml:"number,attr,omitempty"`
	Symbol          string           `xml:"symbol,attr,omitempty"`
	PrintObject     string           `xml:"print-object,attr,omitempty"`
	Beats           []string         `xml:"beats"`
	BeatType        []string         `xml:"beat-type"`
	Interchangeable *Interchangeable `xml:"interchangeable"`
	SenzaMisura     *SenzaMisura     `xml:"senza-misura"`
}

type Interchangeable struct {
	Symbol       string   `xml:"symbol,attr,omitempty"`
	Separator    string   `xml:"separator,attr,omitempty"`
	TimeRelation string   `xml:"time-relation"`
	Beats        []string `xml:"beats"`
	BeatType     []string `xml:"beat-type"`
}

type SenzaMisura struct {
	Value string `xml:",chardata"`
}

type PartSymbol struct {
	TopStaff    int     `xml:"top-staff,attr,omitempty"`
	BottomStaff int     `xml:"bottom-staff,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Clef struct {
	Number           int    `xml:"number,attr,omitempty"`
	Additional       string `xml:"additional,attr,omitempty"`
	Size             string `xml:"size,attr,omitempty"`
	AfterBarline     string `xml:"after-barline,attr,omitempty"`
	PrintObject      string `xml:"print-object,attr,omitempty"`
	Sign             string `xml:"sign"`
	Line             *int   `xml:"line"`
	ClefOctaveChange *int   `xml:"clef-octave-change"`
}

type StaffDetails struct {
	Number       int           `xml:"number,attr,omitempty"`
	ShowFrets    string        `xml:"show-frets,attr,omitempty"`
	PrintObject  string        `xml:"print-object,attr,omitempty"`
	PrintSpacing string        `xml:"print-spacing,attr,omitempty"`
	StaffType    *StaffType    `xml:"staff-type"`
	StaffLines   *int          `xml:"staff-lines"`
	LineDetail   []LineDetail  `xml:"line-detail"`
	StaffTuning  []StaffTuning `xml:"staff-tuning"`
	Capo         *int          `xml:"capo"`
	StaffSize    *float32      `xml:"staff-size"`
}

type StaffType struct {
	Value string `xml:",chardata"`
}

type StaffTuning struct {
	Line         int      `xml:"line,attr"`
	TuningStep   string   `xml:"tuning-step"`
	TuningAlter  *float32 `xml:"tuning-alter"`
	TuningOctave int      `xml:"tuning-octave"`
}

type Transpose struct {
	Number       int     `xml:"number,attr,omitempty"`
	Diatonic     *int    `xml:"diatonic"`
	Chromatic    int     `xml:"chromatic"`
	OctaveChange *int    `xml:"octave-change"`
	Double       *Double `xml:"double"`
}

type Double struct {
	Above string `xml:"above,attr,omitempty"`
}

type ForPart struct {
	ID    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

type Directive struct {
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Halign      string  `xml:"halign,attr,omitempty"`
	Valign      string  `xml:"valign,attr,omitempty"`
	Lang        string  `xml:"xml:lang,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type MeasureStyle struct {
	Number        int            `xml:"number,attr,omitempty"`
	FontFamily    string         `xml:"font-family,attr,omitempty"`
	FontStyle     string         `xml:"font-style,attr,omitempty"`
	FontSize      string         `xml:"font-size,attr,omitempty"`
	FontWeight    string         `xml:"font-weight,attr,omitempty"`
	Color         string         `xml:"color,attr,omitempty"`
	ID            string         `xml:"id,attr,omitempty"`
	MultipleRest  *MultipleRest  `xml:"multiple-rest"`
	MeasureRepeat *MeasureRepeat `xml:"measure-repeat"`
	BeatRepeat    *BeatRepeat    `xml:"beat-repeat"`
	Slash         *Slash         `xml:"slash"`
}

type MultipleRest struct {
	UseSymbols string `xml:"use-symbols,attr,omitempty"`
	Value      int    `xml:",chardata"`
}

type MeasureRepeat struct {
	Type    string `xml:"type,attr"`
	Slashes *int   `xml:"slashes,attr,omitempty"`
}

type BeatRepeat struct {
	Type      string    `xml:"type,attr"`
	Slashes   int       `xml:"slashes"`
	UseDots   string    `xml:"use-dots,attr,omitempty"`
	SlashType *NoteType `xml:"slash-type"`
	SlashDot  []Dot     `xml:"slash-dot"`
}

type Slash struct {
	Type      string    `xml:"type,attr"`
	UseDots   string    `xml:"use-dots,attr,omitempty"`
	UseStems  string    `xml:"use-stems,attr,omitempty"`
	SlashType *NoteType `xml:"slash-type"`
	SlashDot  []Dot     `xml:"slash-dot"`
}

// Backup and Forward
type Backup struct {
	Duration int       `xml:"duration"`
	Footnote *Footnote `xml:"footnote"`
	Level    *Level    `xml:"level"`
}

type Forward struct {
	Duration int       `xml:"duration"`
	Footnote *Footnote `xml:"footnote"`
	Level    *Level    `xml:"level"`
	Voice    *int      `xml:"voice"`
	Staff    *int      `xml:"staff"`
}

// Direction structures
type Direction struct {
	Placement     string          `xml:"placement,attr,omitempty"`
	Directive     string          `xml:"directive,attr,omitempty"`
	System        string          `xml:"system,attr,omitempty"`
	ID            string          `xml:"id,attr,omitempty"`
	DirectionType []DirectionType `xml:"direction-type"`
	Offset        *Offset         `xml:"offset"`
	Footnote      *Footnote       `xml:"footnote"`
	Level         *Level          `xml:"level"`
	Voice         *int            `xml:"voice"`
	Staff         *int            `xml:"staff"`
	Sound         *Sound          `xml:"sound"`
	Listening     *Listening      `xml:"listening"`
}

type DirectionType struct {
	Rehearsal             []Rehearsal            `xml:"rehearsal"`
	Segno                 []Segno                `xml:"segno"`
	Coda                  []Coda                 `xml:"coda"`
	Words                 []Words                `xml:"words"`
	Symbol                []Symbol               `xml:"symbol"`
	Wedge                 *Wedge                 `xml:"wedge"`
	Dynamics              *Dynamics              `xml:"dynamics"`
	Dashes                *Dashes                `xml:"dashes"`
	Bracket               *Bracket               `xml:"bracket"`
	Pedal                 *Pedal                 `xml:"pedal"`
	Metronome             *Metronome             `xml:"metronome"`
	OctaveShift           *OctaveShift           `xml:"octave-shift"`
	HarpPedals            *HarpPedals            `xml:"harp-pedals"`
	Damp                  *Damp                  `xml:"damp"`
	DampAll               *DampAll               `xml:"damp-all"`
	EyeGlasses            *EyeGlasses            `xml:"eyeglasses"`
	StringMute            *StringMute            `xml:"string-mute"`
	Scordatura            *Scordatura            `xml:"scordatura"`
	Image                 *Image                 `xml:"image"`
	PrincipalVoice        *PrincipalVoice        `xml:"principal-voice"`
	Percussion            []Percussion           `xml:"percussion"`
	AccordionRegistration *AccordionRegistration `xml:"accordion-registration"`
	StaffDivide           *StaffDivide           `xml:"staff-divide"`
	OtherDirection        *OtherDirection        `xml:"other-direction"`
}

type Rehearsal struct {
	DefaultX      float32 `xml:"default-x,attr,omitempty"`
	DefaultY      float32 `xml:"default-y,attr,omitempty"`
	RelativeX     float32 `xml:"relative-x,attr,omitempty"`
	RelativeY     float32 `xml:"relative-y,attr,omitempty"`
	FontFamily    string  `xml:"font-family,attr,omitempty"`
	FontStyle     string  `xml:"font-style,attr,omitempty"`
	FontSize      string  `xml:"font-size,attr,omitempty"`
	FontWeight    string  `xml:"font-weight,attr,omitempty"`
	Color         string  `xml:"color,attr,omitempty"`
	Halign        string  `xml:"halign,attr,omitempty"`
	Valign        string  `xml:"valign,attr,omitempty"`
	Lang          string  `xml:"xml:lang,attr,omitempty"`
	Justify       string  `xml:"justify,attr,omitempty"`
	LetterSpacing string  `xml:"letter-spacing,attr,omitempty"`
	LineHeight    string  `xml:"line-height,attr,omitempty"`
	TextDirection string  `xml:"text-direction,attr,omitempty"`
	TextRotation  string  `xml:"text-rotation,attr,omitempty"`
	Enclosure     string  `xml:"enclosure,attr,omitempty"`
	ID            string  `xml:"id,attr,omitempty"`
	Value         string  `xml:",chardata"`
}

type Segno struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
}

type Coda struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
}

type Words struct {
	DefaultX      float32 `xml:"default-x,attr,omitempty"`
	DefaultY      float32 `xml:"default-y,attr,omitempty"`
	RelativeX     float32 `xml:"relative-x,attr,omitempty"`
	RelativeY     float32 `xml:"relative-y,attr,omitempty"`
	FontFamily    string  `xml:"font-family,attr,omitempty"`
	FontStyle     string  `xml:"font-style,attr,omitempty"`
	FontSize      string  `xml:"font-size,attr,omitempty"`
	FontWeight    string  `xml:"font-weight,attr,omitempty"`
	Color         string  `xml:"color,attr,omitempty"`
	Halign        string  `xml:"halign,attr,omitempty"`
	Valign        string  `xml:"valign,attr,omitempty"`
	Lang          string  `xml:"xml:lang,attr,omitempty"`
	Justify       string  `xml:"justify,attr,omitempty"`
	LetterSpacing string  `xml:"letter-spacing,attr,omitempty"`
	LineHeight    string  `xml:"line-height,attr,omitempty"`
	TextDirection string  `xml:"text-direction,attr,omitempty"`
	TextRotation  string  `xml:"text-rotation,attr,omitempty"`
	Enclosure     string  `xml:"enclosure,attr,omitempty"`
	ID            string  `xml:"id,attr,omitempty"`
	Value         string  `xml:",chardata"`
}

type Symbol struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Wedge struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	Niente      string  `xml:"niente,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Dashes struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Bracket struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineEnd     string  `xml:"line-end,attr,omitempty"`
	EndLength   float32 `xml:"end-length,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Pedal struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	Line        string  `xml:"line,attr,omitempty"`
	Sign        string  `xml:"sign,attr,omitempty"`
	Abbreviated string  `xml:"abbreviated,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Halign      string  `xml:"halign,attr,omitempty"`
	Valign      string  `xml:"valign,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Metronome struct {
	DefaultX          float32            `xml:"default-x,attr,omitempty"`
	DefaultY          float32            `xml:"default-y,attr,omitempty"`
	RelativeX         float32            `xml:"relative-x,attr,omitempty"`
	RelativeY         float32            `xml:"relative-y,attr,omitempty"`
	FontFamily        string             `xml:"font-family,attr,omitempty"`
	FontStyle         string             `xml:"font-style,attr,omitempty"`
	FontSize          string             `xml:"font-size,attr,omitempty"`
	FontWeight        string             `xml:"font-weight,attr,omitempty"`
	Color             string             `xml:"color,attr,omitempty"`
	Halign            string             `xml:"halign,attr,omitempty"`
	Valign            string             `xml:"valign,attr,omitempty"`
	Justify           string             `xml:"justify,attr,omitempty"`
	Parentheses       string             `xml:"parentheses,attr,omitempty"`
	ID                string             `xml:"id,attr,omitempty"`
	BeatUnit          []BeatUnit         `xml:"beat-unit"`
	BeatUnitDot       []Dot              `xml:"beat-unit-dot"`
	BeatUnitTied      []BeatUnitTied     `xml:"beat-unit-tied"`
	PerMinute         *PerMinute         `xml:"per-minute"`
	MetronomeArrows   *MetronomeArrows   `xml:"metronome-arrows"`
	MetronomeNote     []MetronomeNote    `xml:"metronome-note"`
	MetronomeRelation *MetronomeRelation `xml:"metronome-relation"`
}

type BeatUnit struct {
	Value string `xml:",chardata"`
}

type BeatUnitTied struct {
	BeatUnit    []BeatUnit `xml:"beat-unit"`
	BeatUnitDot []Dot      `xml:"beat-unit-dot"`
}

type PerMinute struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type MetronomeArrows struct {
	// Empty element
}

type MetronomeNote struct {
	MetronomeTuplet *MetronomeTuplet `xml:"metronome-tuplet"`
	MetronomeType   string           `xml:"metronome-type"`
	MetronomeDot    []Dot            `xml:"metronome-dot"`
	MetronomeBeam   []MetronomeBeam  `xml:"metronome-beam"`
	MetronomeTied   *MetronomeTied   `xml:"metronome-tied"`
}

type MetronomeTuplet struct {
	Type        string    `xml:"type,attr"`
	Bracket     string    `xml:"bracket,attr,omitempty"`
	ShowNumber  string    `xml:"show-number,attr,omitempty"`
	ActualNotes int       `xml:"actual-notes"`
	NormalNotes int       `xml:"normal-notes"`
	NormalType  *NoteType `xml:"normal-type"`
	NormalDot   []Dot     `xml:"normal-dot"`
}

type MetronomeBeam struct {
	Number int    `xml:"number,attr,omitempty"`
	Value  string `xml:",chardata"`
}

type MetronomeTied struct {
	Type string `xml:"type,attr"`
}

type MetronomeRelation struct {
	Value string `xml:",chardata"`
}

type OctaveShift struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	Size        int     `xml:"size,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type HarpPedals struct {
	DefaultX    float32       `xml:"default-x,attr,omitempty"`
	DefaultY    float32       `xml:"default-y,attr,omitempty"`
	RelativeX   float32       `xml:"relative-x,attr,omitempty"`
	RelativeY   float32       `xml:"relative-y,attr,omitempty"`
	FontFamily  string        `xml:"font-family,attr,omitempty"`
	FontStyle   string        `xml:"font-style,attr,omitempty"`
	FontSize    string        `xml:"font-size,attr,omitempty"`
	FontWeight  string        `xml:"font-weight,attr,omitempty"`
	Color       string        `xml:"color,attr,omitempty"`
	Halign      string        `xml:"halign,attr,omitempty"`
	Valign      string        `xml:"valign,attr,omitempty"`
	ID          string        `xml:"id,attr,omitempty"`
	PedalTuning []PedalTuning `xml:"pedal-tuning"`
}

type PedalTuning struct {
	PedalStep  string  `xml:"pedal-step"`
	PedalAlter float32 `xml:"pedal-alter"`
}

type Damp struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
}

type DampAll struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
}

type EyeGlasses struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
}

type StringMute struct {
	Type       string  `xml:"type,attr"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
}

type Scordatura struct {
	ID     string   `xml:"id,attr,omitempty"`
	Accord []Accord `xml:"accord"`
}

type Accord struct {
	TuningStep   string   `xml:"tuning-step"`
	TuningAlter  *float32 `xml:"tuning-alter"`
	TuningOctave int      `xml:"tuning-octave"`
	String       *int     `xml:"string,attr,omitempty"`
}

type Image struct {
	Source    string  `xml:"source,attr"`
	Type      string  `xml:"type,attr"`
	Height    float32 `xml:"height,attr,omitempty"`
	Width     float32 `xml:"width,attr,omitempty"`
	DefaultX  float32 `xml:"default-x,attr,omitempty"`
	DefaultY  float32 `xml:"default-y,attr,omitempty"`
	RelativeX float32 `xml:"relative-x,attr,omitempty"`
	RelativeY float32 `xml:"relative-y,attr,omitempty"`
	Halign    string  `xml:"halign,attr,omitempty"`
	Valign    string  `xml:"valign,attr,omitempty"`
	ID        string  `xml:"id,attr,omitempty"`
}

type PrincipalVoice struct {
	Type       string  `xml:"type,attr"`
	Symbol     string  `xml:"symbol,attr"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Percussion struct {
	DefaultX        float32          `xml:"default-x,attr,omitempty"`
	DefaultY        float32          `xml:"default-y,attr,omitempty"`
	RelativeX       float32          `xml:"relative-x,attr,omitempty"`
	RelativeY       float32          `xml:"relative-y,attr,omitempty"`
	FontFamily      string           `xml:"font-family,attr,omitempty"`
	FontStyle       string           `xml:"font-style,attr,omitempty"`
	FontSize        string           `xml:"font-size,attr,omitempty"`
	FontWeight      string           `xml:"font-weight,attr,omitempty"`
	Color           string           `xml:"color,attr,omitempty"`
	Halign          string           `xml:"halign,attr,omitempty"`
	Valign          string           `xml:"valign,attr,omitempty"`
	Enclosure       string           `xml:"enclosure,attr,omitempty"`
	ID              string           `xml:"id,attr,omitempty"`
	Glass           *Glass           `xml:"glass"`
	Metal           *Metal           `xml:"metal"`
	Wood            *Wood            `xml:"wood"`
	Pitched         *Pitched         `xml:"pitched"`
	Membrane        *Membrane        `xml:"membrane"`
	Effect          *Effect          `xml:"effect"`
	Timpani         *Timpani         `xml:"timpani"`
	Beater          *Beater          `xml:"beater"`
	Stick           *Stick           `xml:"stick"`
	StickLocation   *StickLocation   `xml:"stick-location"`
	OtherPercussion *OtherPercussion `xml:"other-percussion"`
}

type Glass struct {
	Value string `xml:",chardata"`
}

type Metal struct {
	Value string `xml:",chardata"`
}

type Wood struct {
	Value string `xml:",chardata"`
}

type Pitched struct {
	Value string `xml:",chardata"`
}

type Membrane struct {
	Value string `xml:",chardata"`
}

type Effect struct {
	Value string `xml:",chardata"`
}

type Timpani struct {
	// Empty element
}

type Beater struct {
	Tip   string `xml:"tip,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Stick struct {
	Tip           string `xml:"tip,attr,omitempty"`
	DashedCircle  string `xml:"dashed-circle,attr,omitempty"`
	Parentheses   string `xml:"parentheses,attr,omitempty"`
	StickType     string `xml:"stick-type"`
	StickMaterial string `xml:"stick-material"`
}

type StickLocation struct {
	Value string `xml:",chardata"`
}

type OtherPercussion struct {
	Value string `xml:",chardata"`
}

type AccordionRegistration struct {
	DefaultX        float32          `xml:"default-x,attr,omitempty"`
	DefaultY        float32          `xml:"default-y,attr,omitempty"`
	RelativeX       float32          `xml:"relative-x,attr,omitempty"`
	RelativeY       float32          `xml:"relative-y,attr,omitempty"`
	FontFamily      string           `xml:"font-family,attr,omitempty"`
	FontStyle       string           `xml:"font-style,attr,omitempty"`
	FontSize        string           `xml:"font-size,attr,omitempty"`
	FontWeight      string           `xml:"font-weight,attr,omitempty"`
	Color           string           `xml:"color,attr,omitempty"`
	Halign          string           `xml:"halign,attr,omitempty"`
	Valign          string           `xml:"valign,attr,omitempty"`
	ID              string           `xml:"id,attr,omitempty"`
	AccordionHigh   *AccordionHigh   `xml:"accordion-high"`
	AccordionMiddle *AccordionMiddle `xml:"accordion-middle"`
	AccordionLow    *AccordionLow    `xml:"accordion-low"`
}

type AccordionHigh struct {
	// Empty element
}

type AccordionMiddle struct {
	Value int `xml:",chardata"`
}

type AccordionLow struct {
	// Empty element
}

type StaffDivide struct {
	Type       string  `xml:"type,attr"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
}

type OtherDirection struct {
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Halign      string  `xml:"halign,attr,omitempty"`
	Valign      string  `xml:"valign,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Offset struct {
	Sound string `xml:"sound,attr,omitempty"`
	Value int    `xml:",chardata"`
}

type Listening struct {
	Sync           []Sync           `xml:"sync"`
	Wait           []Wait           `xml:"wait"`
	OtherListening []OtherListening `xml:"other-listening"`
}

type Sync struct {
	Type       string `xml:"type,attr"`
	Latency    int    `xml:"latency,attr,omitempty"`
	PlayerName string `xml:"player-name,attr,omitempty"`
	TimeOnly   string `xml:"time-only,attr,omitempty"`
}

type Wait struct {
	PlayerName string `xml:"player-name,attr,omitempty"`
	TimeOnly   string `xml:"time-only,attr,omitempty"`
}

type OtherListening struct {
	Type       string `xml:"type,attr"`
	PlayerName string `xml:"player-name,attr,omitempty"`
	TimeOnly   string `xml:"time-only,attr,omitempty"`
	Value      string `xml:",chardata"`
}

// Figured bass
type Figured struct {
	DefaultX     float32   `xml:"default-x,attr,omitempty"`
	DefaultY     float32   `xml:"default-y,attr,omitempty"`
	RelativeX    float32   `xml:"relative-x,attr,omitempty"`
	RelativeY    float32   `xml:"relative-y,attr,omitempty"`
	FontFamily   string    `xml:"font-family,attr,omitempty"`
	FontStyle    string    `xml:"font-style,attr,omitempty"`
	FontSize     string    `xml:"font-size,attr,omitempty"`
	FontWeight   string    `xml:"font-weight,attr,omitempty"`
	Color        string    `xml:"color,attr,omitempty"`
	PrintObject  string    `xml:"print-object,attr,omitempty"`
	PrintDot     string    `xml:"print-dot,attr,omitempty"`
	PrintSpacing string    `xml:"print-spacing,attr,omitempty"`
	PrintLyric   string    `xml:"print-lyric,attr,omitempty"`
	Parentheses  string    `xml:"parentheses,attr,omitempty"`
	ID           string    `xml:"id,attr,omitempty"`
	Figure       []Figure  `xml:"figure"`
	Duration     *int      `xml:"duration"`
	Footnote     *Footnote `xml:"footnote"`
	Level        *Level    `xml:"level"`
}

type Figure struct {
	Prefix       *StyleText `xml:"prefix"`
	FigureNumber *StyleText `xml:"figure-number"`
	Suffix       *StyleText `xml:"suffix"`
	Extend       *Extend    `xml:"extend"`
}

type StyleText struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

// Harmony
type Harmony struct {
	Type        string     `xml:"type,attr,omitempty"`
	PrintObject string     `xml:"print-object,attr,omitempty"`
	PrintFrame  string     `xml:"print-frame,attr,omitempty"`
	Parentheses string     `xml:"parentheses,attr,omitempty"`
	DefaultX    float32    `xml:"default-x,attr,omitempty"`
	DefaultY    float32    `xml:"default-y,attr,omitempty"`
	RelativeX   float32    `xml:"relative-x,attr,omitempty"`
	RelativeY   float32    `xml:"relative-y,attr,omitempty"`
	FontFamily  string     `xml:"font-family,attr,omitempty"`
	FontStyle   string     `xml:"font-style,attr,omitempty"`
	FontSize    string     `xml:"font-size,attr,omitempty"`
	FontWeight  string     `xml:"font-weight,attr,omitempty"`
	Color       string     `xml:"color,attr,omitempty"`
	Halign      string     `xml:"halign,attr,omitempty"`
	Valign      string     `xml:"valign,attr,omitempty"`
	Placement   string     `xml:"placement,attr,omitempty"`
	ID          string     `xml:"id,attr,omitempty"`
	Root        []Root     `xml:"root"`
	Numeral     *Numeral   `xml:"numeral"`
	Function    *StyleText `xml:"function"`
	Kind        Kind       `xml:"kind"`
	Inversion   *Inversion `xml:"inversion"`
	Bass        *Bass      `xml:"bass"`
	Degree      []Degree   `xml:"degree"`
	Frame       *Frame     `xml:"frame"`
	Offset      *Offset    `xml:"offset"`
	Footnote    *Footnote  `xml:"footnote"`
	Level       *Level     `xml:"level"`
	Staff       *int       `xml:"staff"`
}

type Root struct {
	RootStep  StyleText  `xml:"root-step"`
	RootAlter *RootAlter `xml:"root-alter"`
}

type RootAlter struct {
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Value       float32 `xml:",chardata"`
}

type Kind struct {
	UseSymbols   string  `xml:"use-symbols,attr,omitempty"`
	Text         string  `xml:"text,attr,omitempty"`
	StackDegrees string  `xml:"stack-degrees,attr,omitempty"`
	Parentheses  string  `xml:"parentheses-degrees,attr,omitempty"`
	Bracket      string  `xml:"bracket-degrees,attr,omitempty"`
	PrintObject  string  `xml:"print-object,attr,omitempty"`
	DefaultX     float32 `xml:"default-x,attr,omitempty"`
	DefaultY     float32 `xml:"default-y,attr,omitempty"`
	RelativeX    float32 `xml:"relative-x,attr,omitempty"`
	RelativeY    float32 `xml:"relative-y,attr,omitempty"`
	FontFamily   string  `xml:"font-family,attr,omitempty"`
	FontStyle    string  `xml:"font-style,attr,omitempty"`
	FontSize     string  `xml:"font-size,attr,omitempty"`
	FontWeight   string  `xml:"font-weight,attr,omitempty"`
	Color        string  `xml:"color,attr,omitempty"`
	Halign       string  `xml:"halign,attr,omitempty"`
	Valign       string  `xml:"valign,attr,omitempty"`
	Value        string  `xml:",chardata"`
}

type Inversion struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Value      int     `xml:",chardata"`
}

type Bass struct {
	BassSeparator *BassSeparator `xml:"bass-separator"`
	BassStep      StyleText     `xml:"bass-step"`
	BassAlter     *BassAlter    `xml:"bass-alter"`
}

type BassAlter struct {
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Value       float32 `xml:",chardata"`
}

type Degree struct {
	PrintObject string      `xml:"print-object,attr,omitempty"`
	DegreeValue DegreeValue `xml:"degree-value"`
	DegreeAlter DegreeAlter `xml:"degree-alter"`
	DegreeType  DegreeType  `xml:"degree-type"`
}

type DegreeValue struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Value      int     `xml:",chardata"`
}

type DegreeAlter struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	PlusSign   string  `xml:"plus-sign,attr,omitempty"`
	Value      float32 `xml:",chardata"`
}

type DegreeType struct {
	Text       string  `xml:"text,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Frame struct {
	DefaultX     float32     `xml:"default-x,attr,omitempty"`
	DefaultY     float32     `xml:"default-y,attr,omitempty"`
	RelativeX    float32     `xml:"relative-x,attr,omitempty"`
	RelativeY    float32     `xml:"relative-y,attr,omitempty"`
	Color        string      `xml:"color,attr,omitempty"`
	Halign       string      `xml:"halign,attr,omitempty"`
	Valign       string      `xml:"valign,attr,omitempty"`
	Height       float32     `xml:"height,attr,omitempty"`
	Width        float32     `xml:"width,attr,omitempty"`
	Unplayed     string      `xml:"unplayed,attr,omitempty"`
	ID           string      `xml:"id,attr,omitempty"`
	FrameStrings int         `xml:"frame-strings"`
	FrameFrets   int         `xml:"frame-frets"`
	FirstFret    *FirstFret  `xml:"first-fret"`
	FrameNote    []FrameNote `xml:"frame-note"`
}

type FirstFret struct {
	Text     string `xml:"text,attr,omitempty"`
	Location string `xml:"location,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

type FrameNote struct {
	String    int        `xml:"string"`
	Fret      int        `xml:"fret"`
	Fingering *Fingering `xml:"fingering"`
	Barre     *Barre     `xml:"barre"`
}

type Fingering struct {
	Substitution string  `xml:"substitution,attr,omitempty"`
	Alternate    string  `xml:"alternate,attr,omitempty"`
	DefaultX     float32 `xml:"default-x,attr,omitempty"`
	DefaultY     float32 `xml:"default-y,attr,omitempty"`
	RelativeX    float32 `xml:"relative-x,attr,omitempty"`
	RelativeY    float32 `xml:"relative-y,attr,omitempty"`
	FontFamily   string  `xml:"font-family,attr,omitempty"`
	FontStyle    string  `xml:"font-style,attr,omitempty"`
	FontSize     string  `xml:"font-size,attr,omitempty"`
	FontWeight   string  `xml:"font-weight,attr,omitempty"`
	Color        string  `xml:"color,attr,omitempty"`
	Placement    string  `xml:"placement,attr,omitempty"`
	Value        string  `xml:",chardata"`
}

type Barre struct {
	Type  string `xml:"type,attr"`
	Color string `xml:"color,attr,omitempty"`
}

// Sound
type Sound struct {
	Tempo          *float32        `xml:"tempo,attr,omitempty"`
	Dynamics       *float32        `xml:"dynamics,attr,omitempty"`
	Dacapo         string          `xml:"dacapo,attr,omitempty"`
	Segno          string          `xml:"segno,attr,omitempty"`
	Dalsegno       string          `xml:"dalsegno,attr,omitempty"`
	Coda           string          `xml:"coda,attr,omitempty"`
	Tocoda         string          `xml:"tocoda,attr,omitempty"`
	Divisions      *float32        `xml:"divisions,attr,omitempty"`
	ForwardRepeat  string          `xml:"forward-repeat,attr,omitempty"`
	Fine           string          `xml:"fine,attr,omitempty"`
	TimeOnly       string          `xml:"time-only,attr,omitempty"`
	Pizzicato      string          `xml:"pizzicato,attr,omitempty"`
	Pan            *float32        `xml:"pan,attr,omitempty"`
	Elevation      *float32        `xml:"elevation,attr,omitempty"`
	DamperPedal    string          `xml:"damper-pedal,attr,omitempty"`
	SoftPedal      string          `xml:"soft-pedal,attr,omitempty"`
	SostenutoPedal string          `xml:"sostenuto-pedal,attr,omitempty"`
	ID             string          `xml:"id,attr,omitempty"`
	Swing          *Swing          `xml:"swing"`
	MidiDevice     *MidiDevice     `xml:"midi-device"`
	MidiInstrument *MidiInstrument `xml:"midi-instrument"`
	Play           *Play           `xml:"play"`
}

type Swing struct {
	First      *int      `xml:"first"`
	Second     *int      `xml:"second"`
	SwingType  *string   `xml:"swing-type"`
	SwingStyle *string   `xml:"swing-style"`
	Straight   *Straight `xml:"straight"`
}

type Straight struct {
	// Empty element
}

// Barline
type Barline struct {
	Location string    `xml:"location,attr,omitempty"`
	Segno    *Segno    `xml:"segno"`
	Coda     *Coda     `xml:"coda"`
	Fermata  []Fermata `xml:"fermata"`
	BarStyle *BarStyle `xml:"bar-style"`
	Footnote *Footnote `xml:"footnote"`
	Level    *Level    `xml:"level"`
	WavyLine *WavyLine `xml:"wavy-line"`
	Ending   *Ending   `xml:"ending"`
	Repeat   *Repeat   `xml:"repeat"`
}

type BarStyle struct {
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

type WavyLine struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type Ending struct {
	Number      string  `xml:"number,attr"`
	Type        string  `xml:"type,attr"`
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	EndLength   float32 `xml:"end-length,attr,omitempty"`
	TextX       float32 `xml:"text-x,attr,omitempty"`
	TextY       float32 `xml:"text-y,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Repeat struct {
	Direction string `xml:"direction,attr"`
	Times     int    `xml:"times,attr,omitempty"`
	AfterJump string `xml:"after-jump,attr,omitempty"`
}

// Grouping
type Grouping struct {
	Type     string    `xml:"type,attr"`
	Number   string    `xml:"number,attr,omitempty"`
	MemberOf string    `xml:"member-of,attr,omitempty"`
	ID       string    `xml:"id,attr,omitempty"`
	Feature  []Feature `xml:"feature"`
}

type Feature struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Notation types
type Tied struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	Orientation string  `xml:"orientation,attr,omitempty"`
	BezierX     float32 `xml:"bezier-x,attr,omitempty"`
	BezierY     float32 `xml:"bezier-y,attr,omitempty"`
	BezierX2    float32 `xml:"bezier-x2,attr,omitempty"`
	BezierY2    float32 `xml:"bezier-y2,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Slur struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	Orientation string  `xml:"orientation,attr,omitempty"`
	BezierX     float32 `xml:"bezier-x,attr,omitempty"`
	BezierY     float32 `xml:"bezier-y,attr,omitempty"`
	BezierX2    float32 `xml:"bezier-x2,attr,omitempty"`
	BezierY2    float32 `xml:"bezier-y2,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
}

type Tuplet struct {
	Type         string         `xml:"type,attr"`
	Number       int            `xml:"number,attr,omitempty"`
	Bracket      string         `xml:"bracket,attr,omitempty"`
	ShowNumber   string         `xml:"show-number,attr,omitempty"`
	ShowType     string         `xml:"show-type,attr,omitempty"`
	LineShape    string         `xml:"line-shape,attr,omitempty"`
	DefaultX     float32        `xml:"default-x,attr,omitempty"`
	DefaultY     float32        `xml:"default-y,attr,omitempty"`
	RelativeX    float32        `xml:"relative-x,attr,omitempty"`
	RelativeY    float32        `xml:"relative-y,attr,omitempty"`
	Placement    string         `xml:"placement,attr,omitempty"`
	ID           string         `xml:"id,attr,omitempty"`
	TupletActual *TupletPortion `xml:"tuplet-actual"`
	TupletNormal *TupletPortion `xml:"tuplet-normal"`
}

type TupletPortion struct {
	TupletNumber *TupletNumber `xml:"tuplet-number"`
	TupletType   *TupletType   `xml:"tuplet-type"`
	TupletDot    []TupletDot   `xml:"tuplet-dot"`
}

type TupletNumber struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
	Value      int    `xml:",chardata"`
}

type TupletType struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type TupletDot struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
}

type Glissando struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Slide struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	BendSound   string  `xml:"bend-sound,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	FirstBeat   float32 `xml:"first-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Ornaments struct {
	TrillMark            []TrillMark            `xml:"trill-mark"`
	Turn                 []Turn                 `xml:"turn"`
	DelayedTurn          []DelayedTurn          `xml:"delayed-turn"`
	InvertedTurn         []InvertedTurn         `xml:"inverted-turn"`
	DelayedInvertedTurn  []DelayedInvertedTurn  `xml:"delayed-inverted-turn"`
	VerticalTurn         []VerticalTurn         `xml:"vertical-turn"`
	InvertedVerticalTurn []InvertedVerticalTurn `xml:"inverted-vertical-turn"`
	Shake                []Shake                `xml:"shake"`
	WavyLine             []WavyLine             `xml:"wavy-line"`
	Mordent              []Mordent              `xml:"mordent"`
	InvertedMordent      []InvertedMordent      `xml:"inverted-mordent"`
	Schleifer            []Schleifer            `xml:"schleifer"`
	Tremolo              []Tremolo              `xml:"tremolo"`
	Haydn                []Haydn                `xml:"haydn"`
	OtherOrnament        []OtherOrnament        `xml:"other-ornament"`
	AccidentalMark       []AccidentalMark       `xml:"accidental-mark"`
}

type TrillMark struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type Turn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
	Slash       string  `xml:"slash,attr,omitempty"`
}

type DelayedTurn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
	Slash       string  `xml:"slash,attr,omitempty"`
}

type InvertedTurn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
	Slash       string  `xml:"slash,attr,omitempty"`
}

type DelayedInvertedTurn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
	Slash       string  `xml:"slash,attr,omitempty"`
}

type VerticalTurn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type InvertedVerticalTurn struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type Shake struct {
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type Mordent struct {
	Long        string  `xml:"long,attr,omitempty"`
	Approach    string  `xml:"approach,attr,omitempty"`
	Departure   string  `xml:"departure,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type InvertedMordent struct {
	Long        string  `xml:"long,attr,omitempty"`
	Approach    string  `xml:"approach,attr,omitempty"`
	Departure   string  `xml:"departure,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	StartNote   string  `xml:"start-note,attr,omitempty"`
	TrillStep   string  `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn string  `xml:"two-note-turn,attr,omitempty"`
	Accelerate  string  `xml:"accelerate,attr,omitempty"`
	Beats       int     `xml:"beats,attr,omitempty"`
	SecondBeat  float32 `xml:"second-beat,attr,omitempty"`
	LastBeat    float32 `xml:"last-beat,attr,omitempty"`
}

type Schleifer struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Tremolo struct {
	Type       string  `xml:"type,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      int     `xml:",chardata"`
}

type Haydn struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type OtherOrnament struct {
	Type       string  `xml:"type,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type AccidentalMark struct {
	Parentheses string  `xml:"parentheses,attr,omitempty"`
	Bracket     string  `xml:"bracket,attr,omitempty"`
	Size        string  `xml:"size,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

type Technical struct {
	UpBow          []UpBow          `xml:"up-bow"`
	DownBow        []DownBow        `xml:"down-bow"`
	Harmonic       []Harmonic       `xml:"harmonic"`
	OpenString     []OpenString     `xml:"open-string"`
	ThumbPosition  []ThumbPosition  `xml:"thumb-position"`
	Fingering      []Fingering      `xml:"fingering"`
	Pluck          []Pluck          `xml:"pluck"`
	DoubleTongue   []DoubleTongue   `xml:"double-tongue"`
	TripleTongue   []TripleTongue   `xml:"triple-tongue"`
	Stopped        []Stopped        `xml:"stopped"`
	SnapPizzicato  []SnapPizzicato  `xml:"snap-pizzicato"`
	Fret           []Fret           `xml:"fret"`
	String         []StringNumber   `xml:"string"`
	HammerOn       []HammerOn       `xml:"hammer-on"`
	PullOff        []PullOff        `xml:"pull-off"`
	Bend           []Bend           `xml:"bend"`
	Tap            []Tap            `xml:"tap"`
	Heel           []Heel           `xml:"heel"`
	Toe            []Toe            `xml:"toe"`
	Fingernails    []Fingernails    `xml:"fingernails"`
	Hole           []Hole           `xml:"hole"`
	Arrow          []Arrow          `xml:"arrow"`
	Handbell       []Handbell       `xml:"handbell"`
	BrassBend      []BrassBend      `xml:"brass-bend"`
	Flip           []Flip           `xml:"flip"`
	Smear          []Smear          `xml:"smear"`
	Open           []Open           `xml:"open"`
	HalfMuted      []HalfMuted      `xml:"half-muted"`
	HarmonMute     []HarmonMute     `xml:"harmon-mute"`
	Golpe          []Golpe          `xml:"golpe"`
	OtherTechnical []OtherTechnical `xml:"other-technical"`
}

type UpBow struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type DownBow struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Harmonic struct {
	PrintObject   string         `xml:"print-object,attr,omitempty"`
	DefaultX      float32        `xml:"default-x,attr,omitempty"`
	DefaultY      float32        `xml:"default-y,attr,omitempty"`
	RelativeX     float32        `xml:"relative-x,attr,omitempty"`
	RelativeY     float32        `xml:"relative-y,attr,omitempty"`
	FontFamily    string         `xml:"font-family,attr,omitempty"`
	FontStyle     string         `xml:"font-style,attr,omitempty"`
	FontSize      string         `xml:"font-size,attr,omitempty"`
	FontWeight    string         `xml:"font-weight,attr,omitempty"`
	Color         string         `xml:"color,attr,omitempty"`
	Placement     string         `xml:"placement,attr,omitempty"`
	Natural       *Natural       `xml:"natural"`
	Artificial    *Artificial    `xml:"artificial"`
	BasePitch     *BasePitch     `xml:"base-pitch"`
	TouchingPitch *TouchingPitch `xml:"touching-pitch"`
	SoundingPitch *SoundingPitch `xml:"sounding-pitch"`
}

type Natural struct {
	// Empty element
}

type Artificial struct {
	// Empty element
}

type BasePitch struct {
	// Empty element
}

type TouchingPitch struct {
	// Empty element
}

type SoundingPitch struct {
	// Empty element
}

type OpenString struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type ThumbPosition struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Pluck struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type DoubleTongue struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type TripleTongue struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Stopped struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type SnapPizzicato struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Fret struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle  string `xml:"font-style,attr,omitempty"`
	FontSize   string `xml:"font-size,attr,omitempty"`
	FontWeight string `xml:"font-weight,attr,omitempty"`
	Color      string `xml:"color,attr,omitempty"`
	Value      int    `xml:",chardata"`
}

type StringNumber struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      int     `xml:",chardata"`
}

type HammerOn struct {
	Type       string  `xml:"type,attr"`
	Number     int     `xml:"number,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type PullOff struct {
	Type       string  `xml:"type,attr"`
	Number     int     `xml:"number,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Bend struct {
	Shape      string   `xml:"shape,attr,omitempty"`
	DefaultX   float32  `xml:"default-x,attr,omitempty"`
	DefaultY   float32  `xml:"default-y,attr,omitempty"`
	RelativeX  float32  `xml:"relative-x,attr,omitempty"`
	RelativeY  float32  `xml:"relative-y,attr,omitempty"`
	FontFamily string   `xml:"font-family,attr,omitempty"`
	FontStyle  string   `xml:"font-style,attr,omitempty"`
	FontSize   string   `xml:"font-size,attr,omitempty"`
	FontWeight string   `xml:"font-weight,attr,omitempty"`
	Color      string   `xml:"color,attr,omitempty"`
	Accelerate string   `xml:"accelerate,attr,omitempty"`
	Beats      int      `xml:"beats,attr,omitempty"`
	FirstBeat  float32  `xml:"first-beat,attr,omitempty"`
	LastBeat   float32  `xml:"last-beat,attr,omitempty"`
	BendAlter  float32  `xml:"bend-alter"`
	PreBend    *PreBend `xml:"pre-bend"`
	Release    *Release `xml:"release"`
	WithBar    *WithBar `xml:"with-bar"`
}

type PreBend struct {
	// Empty element
}

type Release struct {
	Offset float32 `xml:"offset,attr,omitempty"`
}

type WithBar struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Tap struct {
	Hand       string  `xml:"hand,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Heel struct {
	Substitution string  `xml:"substitution,attr,omitempty"`
	DefaultX     float32 `xml:"default-x,attr,omitempty"`
	DefaultY     float32 `xml:"default-y,attr,omitempty"`
	RelativeX    float32 `xml:"relative-x,attr,omitempty"`
	RelativeY    float32 `xml:"relative-y,attr,omitempty"`
	FontFamily   string  `xml:"font-family,attr,omitempty"`
	FontStyle    string  `xml:"font-style,attr,omitempty"`
	FontSize     string  `xml:"font-size,attr,omitempty"`
	FontWeight   string  `xml:"font-weight,attr,omitempty"`
	Color        string  `xml:"color,attr,omitempty"`
	Placement    string  `xml:"placement,attr,omitempty"`
}

type Toe struct {
	Substitution string  `xml:"substitution,attr,omitempty"`
	DefaultX     float32 `xml:"default-x,attr,omitempty"`
	DefaultY     float32 `xml:"default-y,attr,omitempty"`
	RelativeX    float32 `xml:"relative-x,attr,omitempty"`
	RelativeY    float32 `xml:"relative-y,attr,omitempty"`
	FontFamily   string  `xml:"font-family,attr,omitempty"`
	FontStyle    string  `xml:"font-style,attr,omitempty"`
	FontSize     string  `xml:"font-size,attr,omitempty"`
	FontWeight   string  `xml:"font-weight,attr,omitempty"`
	Color        string  `xml:"color,attr,omitempty"`
	Placement    string  `xml:"placement,attr,omitempty"`
}

type Fingernails struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Hole struct {
	HoleClosed *HoleClosed `xml:"hole-closed"`
	HoleShape  string      `xml:"hole-shape"`
	HoleType   string      `xml:"hole-type"`
	DefaultX   float32     `xml:"default-x,attr,omitempty"`
	DefaultY   float32     `xml:"default-y,attr,omitempty"`
	RelativeX  float32     `xml:"relative-x,attr,omitempty"`
	RelativeY  float32     `xml:"relative-y,attr,omitempty"`
	FontFamily string      `xml:"font-family,attr,omitempty"`
	FontStyle  string      `xml:"font-style,attr,omitempty"`
	FontSize   string      `xml:"font-size,attr,omitempty"`
	FontWeight string      `xml:"font-weight,attr,omitempty"`
	Color      string      `xml:"color,attr,omitempty"`
	Placement  string      `xml:"placement,attr,omitempty"`
}

type HoleClosed struct {
	Location string `xml:"location,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Arrow struct {
	ArrowDirection *ArrowDirection `xml:"arrow-direction"`
	ArrowStyle     *ArrowStyle     `xml:"arrow-style"`
	Arrowhead      *Arrowhead      `xml:"arrowhead"`
	CircularArrow  *CircularArrow  `xml:"circular-arrow"`
	DefaultX       float32         `xml:"default-x,attr,omitempty"`
	DefaultY       float32         `xml:"default-y,attr,omitempty"`
	RelativeX      float32         `xml:"relative-x,attr,omitempty"`
	RelativeY      float32         `xml:"relative-y,attr,omitempty"`
	FontFamily     string          `xml:"font-family,attr,omitempty"`
	FontStyle      string          `xml:"font-style,attr,omitempty"`
	FontSize       string          `xml:"font-size,attr,omitempty"`
	FontWeight     string          `xml:"font-weight,attr,omitempty"`
	Color          string          `xml:"color,attr,omitempty"`
	Placement      string          `xml:"placement,attr,omitempty"`
}

type ArrowDirection struct {
	Value string `xml:",chardata"`
}

type ArrowStyle struct {
	Value string `xml:",chardata"`
}

type Arrowhead struct {
	// Empty element
}

type CircularArrow struct {
	Value string `xml:",chardata"`
}

type Handbell struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type BrassBend struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Flip struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Smear struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Open struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type HalfMuted struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type HarmonMute struct {
	DefaultX     float32       `xml:"default-x,attr,omitempty"`
	DefaultY     float32       `xml:"default-y,attr,omitempty"`
	RelativeX    float32       `xml:"relative-x,attr,omitempty"`
	RelativeY    float32       `xml:"relative-y,attr,omitempty"`
	FontFamily   string        `xml:"font-family,attr,omitempty"`
	FontStyle    string        `xml:"font-style,attr,omitempty"`
	FontSize     string        `xml:"font-size,attr,omitempty"`
	FontWeight   string        `xml:"font-weight,attr,omitempty"`
	Color        string        `xml:"color,attr,omitempty"`
	Placement    string        `xml:"placement,attr,omitempty"`
	HarmonClosed *HarmonClosed `xml:"harmon-closed"`
	Value        string        `xml:",chardata"`
}

type HarmonClosed struct {
	Location string `xml:"location,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Golpe struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type OtherTechnical struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Articulations struct {
	Accent            []Accent            `xml:"accent"`
	StrongAccent      []StrongAccent      `xml:"strong-accent"`
	Staccato          []Staccato          `xml:"staccato"`
	Tenuto            []Tenuto            `xml:"tenuto"`
	DetachedLegato    []DetachedLegato    `xml:"detached-legato"`
	Staccatissimo     []Staccatissimo     `xml:"staccatissimo"`
	Spiccato          []Spiccato          `xml:"spiccato"`
	Scoop             []Scoop             `xml:"scoop"`
	Plop              []Plop              `xml:"plop"`
	Doit              []Doit              `xml:"doit"`
	Falloff           []Falloff           `xml:"falloff"`
	BreathMark        []BreathMark        `xml:"breath-mark"`
	Caesura           []Caesura           `xml:"caesura"`
	Stress            []Stress            `xml:"stress"`
	Unstress          []Unstress          `xml:"unstress"`
	SoftAccent        []SoftAccent        `xml:"soft-accent"`
	OtherArticulation []OtherArticulation `xml:"other-articulation"`
}

type Accent struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type StrongAccent struct {
	Type       string  `xml:"type,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Staccato struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Tenuto struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type DetachedLegato struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Staccatissimo struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Spiccato struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Scoop struct {
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
}

type Plop struct {
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
}

type Doit struct {
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
}

type Falloff struct {
	LineType    string  `xml:"line-type,attr,omitempty"`
	DashLength  float32 `xml:"dash-length,attr,omitempty"`
	SpaceLength float32 `xml:"space-length,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
}

type BreathMark struct {
	Line       string  `xml:"line,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Caesura struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Stress struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type Unstress struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type SoftAccent struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
}

type OtherArticulation struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Placement  string  `xml:"placement,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Dynamics struct {
	DefaultX      float32         `xml:"default-x,attr,omitempty"`
	DefaultY      float32         `xml:"default-y,attr,omitempty"`
	RelativeX     float32         `xml:"relative-x,attr,omitempty"`
	RelativeY     float32         `xml:"relative-y,attr,omitempty"`
	FontFamily    string          `xml:"font-family,attr,omitempty"`
	FontStyle     string          `xml:"font-style,attr,omitempty"`
	FontSize      string          `xml:"font-size,attr,omitempty"`
	FontWeight    string          `xml:"font-weight,attr,omitempty"`
	Color         string          `xml:"color,attr,omitempty"`
	Halign        string          `xml:"halign,attr,omitempty"`
	Valign        string          `xml:"valign,attr,omitempty"`
	Placement     string          `xml:"placement,attr,omitempty"`
	Underline     string          `xml:"underline,attr,omitempty"`
	Overline      string          `xml:"overline,attr,omitempty"`
	LineThrough   string          `xml:"line-through,attr,omitempty"`
	Enclosure     string          `xml:"enclosure,attr,omitempty"`
	ID            string          `xml:"id,attr,omitempty"`
	P             *Empty          `xml:"p"`
	PP            *Empty          `xml:"pp"`
	PPP           *Empty          `xml:"ppp"`
	PPPP          *Empty          `xml:"pppp"`
	PPPPP         *Empty          `xml:"ppppp"`
	PPPPPP        *Empty          `xml:"pppppp"`
	F             *Empty          `xml:"f"`
	FF            *Empty          `xml:"ff"`
	FFF           *Empty          `xml:"fff"`
	FFFF          *Empty          `xml:"ffff"`
	FFFFF         *Empty          `xml:"fffff"`
	FFFFFF        *Empty          `xml:"ffffff"`
	MP            *Empty          `xml:"mp"`
	MF            *Empty          `xml:"mf"`
	SF            *Empty          `xml:"sf"`
	SFP           *Empty          `xml:"sfp"`
	SFPP          *Empty          `xml:"sfpp"`
	FP            *Empty          `xml:"fp"`
	RF            *Empty          `xml:"rf"`
	RFZ           *Empty          `xml:"rfz"`
	SFZ           *Empty          `xml:"sfz"`
	SFFZ          *Empty          `xml:"sffz"`
	FZ            *Empty          `xml:"fz"`
	N             *Empty          `xml:"n"`
	PF            *Empty          `xml:"pf"`
	SFZP          *Empty          `xml:"sfzp"`
	OtherDynamics []OtherDynamics `xml:"other-dynamics"`
}

type Empty struct {
	// Empty element
}

type OtherDynamics struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Fermata struct {
	Type       string  `xml:"type,attr,omitempty"`
	Shape      string  `xml:"shape,attr,omitempty"`
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	ID         string  `xml:"id,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type Arpeggiate struct {
	Number    int     `xml:"number,attr,omitempty"`
	Direction string  `xml:"direction,attr,omitempty"`
	Unbroken  string  `xml:"unbroken,attr,omitempty"`
	DefaultX  float32 `xml:"default-x,attr,omitempty"`
	DefaultY  float32 `xml:"default-y,attr,omitempty"`
	RelativeX float32 `xml:"relative-x,attr,omitempty"`
	RelativeY float32 `xml:"relative-y,attr,omitempty"`
	Color     string  `xml:"color,attr,omitempty"`
	ID        string  `xml:"id,attr,omitempty"`
}

type NonArpeggiate struct {
	Type      string  `xml:"type,attr"`
	Number    int     `xml:"number,attr,omitempty"`
	DefaultX  float32 `xml:"default-x,attr,omitempty"`
	DefaultY  float32 `xml:"default-y,attr,omitempty"`
	RelativeX float32 `xml:"relative-x,attr,omitempty"`
	RelativeY float32 `xml:"relative-y,attr,omitempty"`
	Color     string  `xml:"color,attr,omitempty"`
	ID        string  `xml:"id,attr,omitempty"`
}

type OtherNotation struct {
	Type        string  `xml:"type,attr"`
	Number      int     `xml:"number,attr,omitempty"`
	PrintObject string  `xml:"print-object,attr,omitempty"`
	DefaultX    float32 `xml:"default-x,attr,omitempty"`
	DefaultY    float32 `xml:"default-y,attr,omitempty"`
	RelativeX   float32 `xml:"relative-x,attr,omitempty"`
	RelativeY   float32 `xml:"relative-y,attr,omitempty"`
	FontFamily  string  `xml:"font-family,attr,omitempty"`
	FontStyle   string  `xml:"font-style,attr,omitempty"`
	FontSize    string  `xml:"font-size,attr,omitempty"`
	FontWeight  string  `xml:"font-weight,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Placement   string  `xml:"placement,attr,omitempty"`
	ID          string  `xml:"id,attr,omitempty"`
	Value       string  `xml:",chardata"`
}

// Display text elements
type DisplayText struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Lang       string  `xml:"xml:lang,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

type AccidentalText struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Halign     string  `xml:"halign,attr,omitempty"`
	Valign     string  `xml:"valign,attr,omitempty"`
	Lang       string  `xml:"xml:lang,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

// Cue note
type Cue struct {
	// Empty element, presence indicates cue note
}

// Notehead text
type NoteheadText struct {
	DisplayText    []DisplayText    `xml:"display-text"`
	AccidentalText []AccidentalText `xml:"accidental-text"`
}

// Listen (note-level, MusicXML 4.0)
type Listen struct {
	Assess      []Assess      `xml:"assess"`
	Wait        []Wait        `xml:"wait"`
	OtherListen []OtherListen `xml:"other-listen"`
}

type Assess struct {
	Type     string `xml:"type,attr"`
	Player   string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
}

type OtherListen struct {
	Type     string `xml:"type,attr"`
	Player   string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// Other play
type OtherPlay struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Instrument change (MusicXML 4.0)
type InstrumentChange struct {
	ID                     string             `xml:"id,attr"`
	InstrumentName         string             `xml:"instrument-name"`
	InstrumentAbbreviation string             `xml:"instrument-abbreviation,omitempty"`
	InstrumentSound        string             `xml:"instrument-sound,omitempty"`
	Solo                   *Solo              `xml:"solo,omitempty"`
	Ensemble               *int               `xml:"ensemble,omitempty"`
	VirtualInstrument      *VirtualInstrument `xml:"virtual-instrument,omitempty"`
}

// Key octave
type KeyOctave struct {
	Number int    `xml:"number,attr"`
	Cancel string `xml:"cancel,attr,omitempty"`
	Value  int    `xml:",chardata"`
}

// Line detail (MusicXML 4.0)
type LineDetail struct {
	Line        int     `xml:"line,attr"`
	LineType    string  `xml:"line-type,attr,omitempty"`
	PrintObject string  `xml:"print-object,attr,omitempty"`
	Color       string  `xml:"color,attr,omitempty"`
	Width       float32 `xml:"width,attr,omitempty"`
}

// Numeral (MusicXML 4.0)
type Numeral struct {
	NumeralRoot  NumeralRoot   `xml:"numeral-root"`
	NumeralAlter *NumeralAlter `xml:"numeral-alter"`
	NumeralKey   *NumeralKey   `xml:"numeral-key"`
}

type NumeralRoot struct {
	Text  string `xml:"text,attr,omitempty"`
	Value int    `xml:",chardata"`
}

type NumeralAlter struct {
	PrintObject string  `xml:"print-object,attr,omitempty"`
	Location    string  `xml:"location,attr,omitempty"`
	Value       float32 `xml:",chardata"`
}

type NumeralKey struct {
	PrintObject   string `xml:"print-object,attr,omitempty"`
	NumeralFifths int    `xml:"numeral-fifths"`
	NumeralMode   string `xml:"numeral-mode"`
}

// Bass separator
type BassSeparator struct {
	DefaultX   float32 `xml:"default-x,attr,omitempty"`
	DefaultY   float32 `xml:"default-y,attr,omitempty"`
	RelativeX  float32 `xml:"relative-x,attr,omitempty"`
	RelativeY  float32 `xml:"relative-y,attr,omitempty"`
	FontFamily string  `xml:"font-family,attr,omitempty"`
	FontStyle  string  `xml:"font-style,attr,omitempty"`
	FontSize   string  `xml:"font-size,attr,omitempty"`
	FontWeight string  `xml:"font-weight,attr,omitempty"`
	Color      string  `xml:"color,attr,omitempty"`
	Value      string  `xml:",chardata"`
}

// Player (MusicXML 4.0)
type Player struct {
	ID         string `xml:"id,attr"`
	PlayerName string `xml:"player-name"`
}

// Part link (MusicXML 4.0)
type PartLink struct {
	Href           string           `xml:"xlink:href,attr"`
	Type           string           `xml:"xlink:type,attr,omitempty"`
	Role           string           `xml:"xlink:role,attr,omitempty"`
	Title          string           `xml:"xlink:title,attr,omitempty"`
	Show           string           `xml:"xlink:show,attr,omitempty"`
	Actuate        string           `xml:"xlink:actuate,attr,omitempty"`
	GroupLink      []GroupLink      `xml:"group-link"`
	InstrumentLink []InstrumentLink `xml:"instrument-link"`
	PartClef       *PartClef        `xml:"part-clef"`
	PartTranspose  *PartTranspose   `xml:"part-transpose"`
}

type GroupLink struct {
	Value string `xml:",chardata"`
}

type InstrumentLink struct {
	ID string `xml:"id,attr"`
}

type PartClef struct {
	Number           int    `xml:"number,attr,omitempty"`
	Sign             string `xml:"sign"`
	Line             *int   `xml:"line"`
	ClefOctaveChange *int   `xml:"clef-octave-change"`
}

type PartTranspose struct {
	Number       int     `xml:"number,attr,omitempty"`
	Diatonic     *int    `xml:"diatonic"`
	Chromatic    int     `xml:"chromatic"`
	OctaveChange *int    `xml:"octave-change"`
	Double       *Double `xml:"double"`
}
