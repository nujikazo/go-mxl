package examples

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Sample MusicXML parser demonstrating usage of the structures

func main() {
	// Check if filename is provided
	if len(os.Args) < 2 {
		// If no file provided, use embedded sample
		fmt.Println("No file provided, using embedded sample MusicXML")
		parseSampleXML()
		return
	}

	// Parse the provided file
	filename := os.Args[1]
	parseFile(filename)
}

// parseFile reads and parses a MusicXML file
func parseFile(filename string) {
	fmt.Printf("Parsing MusicXML file: %s\n", filename)
	fmt.Println(strings.Repeat("-", 50))

	// Read the file
	xmlData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse the XML
	var score ScorePartWise
	err = xml.Unmarshal(xmlData, &score)
	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}

	// Display the parsed information
	displayScoreInfo(&score)
}

// parseSampleXML demonstrates parsing with an embedded sample
func parseSampleXML() {
	sampleXML := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE score-partwise PUBLIC "-//Recordare//DTD MusicXML 3.1 Partwise//EN" "http://www.musicxml.org/dtds/partwise.dtd">
<score-partwise version="3.1">
  <work>
    <work-title>Sample Score</work-title>
  </work>
  <movement-title>Movement 1</movement-title>
  <identification>
    <creator type="composer">John Doe</creator>
    <creator type="lyricist">Jane Smith</creator>
    <rights>Copyright © 2024</rights>
    <encoding>
      <software>MusicXML Sample Parser</software>
      <encoding-date>2024-01-01</encoding-date>
    </encoding>
  </identification>
  <defaults>
    <scaling>
      <millimeters>7.05556</millimeters>
      <tenths>40</tenths>
    </scaling>
    <page-layout>
      <page-height>1683.36</page-height>
      <page-width>1190.88</page-width>
      <page-margins type="even">
        <left-margin>56.6929</left-margin>
        <right-margin>56.6929</right-margin>
        <top-margin>56.6929</top-margin>
        <bottom-margin>113.386</bottom-margin>
      </page-margins>
    </page-layout>
  </defaults>
  <credit page="1">
    <credit-words default-x="595.44" default-y="1627.31" justify="center" valign="top" font-size="24">Sample Score</credit-words>
  </credit>
  <part-list>
    <score-part id="P1">
      <part-name>Piano</part-name>
      <score-instrument id="P1-I1">
        <instrument-name>Piano</instrument-name>
      </score-instrument>
      <midi-device id="P1-I1" port="1"></midi-device>
      <midi-instrument id="P1-I1">
        <midi-channel>1</midi-channel>
        <midi-program>1</midi-program>
        <volume>78.7402</volume>
        <pan>0</pan>
      </midi-instrument>
    </score-part>
  </part-list>
  <part id="P1">
    <measure number="1" width="400">
      <print>
        <system-layout>
          <system-margins>
            <left-margin>50</left-margin>
            <right-margin>0</right-margin>
          </system-margins>
          <top-system-distance>170</top-system-distance>
        </system-layout>
      </print>
      <attributes>
        <divisions>4</divisions>
        <key>
          <fifths>0</fifths>
        </key>
        <time>
          <beats>4</beats>
          <beat-type>4</beat-type>
        </time>
        <clef>
          <sign>G</sign>
          <line>2</line>
        </clef>
      </attributes>
      <note>
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>quarter</type>
      </note>
      <note>
        <pitch>
          <step>D</step>
          <octave>4</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>quarter</type>
      </note>
      <note>
        <pitch>
          <step>E</step>
          <octave>4</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>quarter</type>
      </note>
      <note>
        <pitch>
          <step>F</step>
          <octave>4</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>quarter</type>
      </note>
    </measure>
    <measure number="2" width="300">
      <note>
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>8</duration>
        <voice>1</voice>
        <type>half</type>
        <notations>
          <dynamics default-x="-36" default-y="-40">
            <f/>
          </dynamics>
        </notations>
      </note>
      <note>
        <pitch>
          <step>A</step>
          <octave>4</octave>
        </pitch>
        <duration>8</duration>
        <voice>1</voice>
        <type>half</type>
      </note>
    </measure>
  </part>
</score-partwise>`

	// Parse the sample XML
	var score ScorePartWise
	err := xml.Unmarshal([]byte(sampleXML), &score)
	if err != nil {
		log.Fatalf("Error parsing sample XML: %v", err)
	}

	// Display the parsed information
	displayScoreInfo(&score)
}

// displayScoreInfo shows the parsed score information
func displayScoreInfo(score *ScorePartWise) {
	// Basic Information
	fmt.Println("SCORE INFORMATION")
	fmt.Println(strings.Repeat("-", 50))

	if score.Work != nil && score.Work.WorkTitle != nil {
		fmt.Printf("Work Title: %s\n", score.Work.WorkTitle.Value)
	}

	if score.MovementTitle != nil {
		fmt.Printf("Movement Title: %s\n", score.MovementTitle.Value)
	}

	fmt.Printf("MusicXML Version: %s\n", score.Version)

	// Identification
	if score.Identification != nil {
		fmt.Println("\nIDENTIFICATION")
		fmt.Println(strings.Repeat("-", 50))

		for _, creator := range score.Identification.Creator {
			fmt.Printf("%s: %s\n", strings.Title(creator.Type), creator.Value)
		}

		if score.Identification.Rights != nil && len(score.Identification.Rights) > 0 {
			fmt.Printf("Rights: %s\n", score.Identification.Rights[0].Value)
		}

		if score.Identification.Encoding != nil {
			if len(score.Identification.Encoding.Software) > 0 {
				fmt.Printf("Software: %s\n", score.Identification.Encoding.Software[0].Value)
			}
			if len(score.Identification.Encoding.EncodingDate) > 0 {
				fmt.Printf("Encoding Date: %s\n", score.Identification.Encoding.EncodingDate[0].Value)
			}
		}
	}

	// Defaults
	if score.Defaults != nil {
		fmt.Println("\nPAGE LAYOUT")
		fmt.Println(strings.Repeat("-", 50))

		if score.Defaults.Scaling != nil {
			fmt.Printf("Scaling: %.2f mm = %.0f tenths\n",
				score.Defaults.Scaling.Millimeters,
				score.Defaults.Scaling.Tenths)
		}

		if score.Defaults.PageLayout != nil {
			if score.Defaults.PageLayout.PageHeight != nil {
				fmt.Printf("Page Height: %.2f\n", *score.Defaults.PageLayout.PageHeight)
			}
			if score.Defaults.PageLayout.PageWidth != nil {
				fmt.Printf("Page Width: %.2f\n", *score.Defaults.PageLayout.PageWidth)
			}
		}
	}

	// Credits
	if len(score.Credit) > 0 {
		fmt.Println("\nCREDITS")
		fmt.Println(strings.Repeat("-", 50))
		for i, credit := range score.Credit {
			for _, words := range credit.CreditWords {
				fmt.Printf("Credit %d: %s\n", i+1, words.Value)
			}
		}
	}

	// Parts
	if score.PartList != nil {
		fmt.Println("\nPARTS")
		fmt.Println(strings.Repeat("-", 50))

		for _, scorePart := range score.PartList.ScorePart {
			fmt.Printf("\nPart ID: %s\n", scorePart.ID)
			fmt.Printf("Part Name: %s\n", scorePart.PartName.Value)

			// Instruments
			for _, inst := range scorePart.ScoreInstrument {
				fmt.Printf("  Instrument: %s (ID: %s)\n", inst.InstrumentName, inst.ID)
			}

			// MIDI Information
			for _, midi := range scorePart.MidiInstrument {
				fmt.Printf("  MIDI Channel: %d, Program: %d\n",
					*midi.MidiChannel, *midi.MidiProgram)
				if midi.Volume != nil {
					fmt.Printf("  Volume: %.2f\n", *midi.Volume)
				}
			}
		}
	}

	// Analyze Parts
	fmt.Println("\nMEASURE ANALYSIS")
	fmt.Println(strings.Repeat("-", 50))

	for _, part := range score.Part {
		fmt.Printf("\nPart: %s\n", part.ID)
		fmt.Printf("Total Measures: %d\n", len(part.Measure))

		// Analyze first few measures
		for i, measure := range part.Measure {
			if i >= 3 { // Only show first 3 measures
				break
			}

			fmt.Printf("\n  Measure %s:\n", measure.Number)

			// Attributes
			if measure.Attributes != nil {
				if measure.Attributes.Divisions != nil {
					fmt.Printf("    Divisions: %d\n", *measure.Attributes.Divisions)
				}

				// Time Signature
				for _, time := range measure.Attributes.Time {
					if len(time.Beats) > 0 && len(time.BeatType) > 0 {
						fmt.Printf("    Time Signature: %s/%s\n", time.Beats[0], time.BeatType[0])
					}
				}

				// Key Signature
				for _, key := range measure.Attributes.Key {
					if key.Fifths != nil {
						fmt.Printf("    Key Signature: %d fifths\n", *key.Fifths)
					}
				}

				// Clef
				for _, clef := range measure.Attributes.Clef {
					fmt.Printf("    Clef: %s line %d\n", clef.Sign, *clef.Line)
				}
			}

			// Notes
			noteCount := len(measure.Note)
			fmt.Printf("    Notes: %d\n", noteCount)

			// Show first few notes
			for j, note := range measure.Note {
				if j >= 5 { // Only show first 5 notes
					break
				}

				noteStr := "    Note: "
				if note.Pitch != nil {
					noteStr += fmt.Sprintf("%s%d", note.Pitch.Step, note.Pitch.Octave)
				} else if note.Rest != nil {
					noteStr += "Rest"
				}

				if note.Duration != nil {
					noteStr += fmt.Sprintf(" (dur: %d)", *note.Duration)
				}

				if note.Type != nil {
					noteStr += fmt.Sprintf(" [%s]", note.Type.Value)
				}

				// Check for dynamics
				for _, notation := range note.Notations {
					if notation.Dynamics != nil {
						noteStr += " (dynamics)"
					}
				}

				fmt.Println(noteStr)
			}

			if noteCount > 5 {
				fmt.Printf("    ... and %d more notes\n", noteCount-5)
			}
		}

		if len(part.Measure) > 3 {
			fmt.Printf("\n  ... and %d more measures\n", len(part.Measure)-3)
		}
	}

	// Summary Statistics
	fmt.Println("\nSUMMARY STATISTICS")
	fmt.Println(strings.Repeat("-", 50))

	totalNotes := 0
	totalMeasures := 0
	hasLyrics := false
	hasDynamics := false
	hasArticulations := false

	for _, part := range score.Part {
		totalMeasures += len(part.Measure)
		for _, measure := range part.Measure {
			totalNotes += len(measure.Note)

			for _, note := range measure.Note {
				if len(note.Lyric) > 0 {
					hasLyrics = true
				}

				for _, notation := range note.Notations {
					if notation.Dynamics != nil {
						hasDynamics = true
					}
					if notation.Articulations != nil {
						hasArticulations = true
					}
				}
			}
		}
	}

	fmt.Printf("Total Parts: %d\n", len(score.Part))
	fmt.Printf("Total Measures: %d\n", totalMeasures)
	fmt.Printf("Total Notes: %d\n", totalNotes)
	fmt.Printf("Has Lyrics: %v\n", hasLyrics)
	fmt.Printf("Has Dynamics: %v\n", hasDynamics)
	fmt.Printf("Has Articulations: %v\n", hasArticulations)
}

// Helper function to extract pitch name
func getPitchName(note *Note) string {
	if note.Pitch != nil {
		alter := ""
		if note.Pitch.Alter != nil {
			if *note.Pitch.Alter == 1 {
				alter = "#"
			} else if *note.Pitch.Alter == -1 {
				alter = "b"
			}
		}
		return fmt.Sprintf("%s%s%d", note.Pitch.Step, alter, note.Pitch.Octave)
	} else if note.Rest != nil {
		return "Rest"
	}
	return "Unknown"
}

// Helper function to get note duration in beats
func getNoteDurationInBeats(note *Note, divisions int) float64 {
	if note.Duration != nil && divisions > 0 {
		return float64(*note.Duration) / float64(divisions)
	}
	return 0
}

// Example of how to find specific elements
func findDynamicMarkings(score *ScorePartWise) []string {
	var dynamics []string

	for _, part := range score.Part {
		for _, measure := range part.Measure {
			// Check notes for dynamics
			for _, note := range measure.Note {
				for _, notation := range note.Notations {
					if notation.Dynamics != nil {
						if notation.Dynamics.P != nil {
							dynamics = append(dynamics, "p")
						}
						if notation.Dynamics.F != nil {
							dynamics = append(dynamics, "f")
						}
						if notation.Dynamics.MF != nil {
							dynamics = append(dynamics, "mf")
						}
						// ... check other dynamics
					}
				}
			}

			// Check directions for dynamics
			for _, direction := range measure.Direction {
				for _, dirType := range direction.DirectionType {
					if dirType.Dynamics != nil {
						// Process dynamics in directions
					}
				}
			}
		}
	}

	return dynamics
}

// Example of how to extract all lyrics
func extractLyrics(score *ScorePartWise) map[string][]string {
	lyrics := make(map[string][]string)

	for _, part := range score.Part {
		partLyrics := []string{}

		for _, measure := range part.Measure {
			for _, note := range measure.Note {
				for _, lyric := range note.Lyric {
					if lyric.Text.Value != "" {
						partLyrics = append(partLyrics, lyric.Text.Value)
					}
				}
			}
		}

		if len(partLyrics) > 0 {
			lyrics[part.ID] = partLyrics
		}
	}

	return lyrics
}

// Example of MIDI export preparation
func prepareMIDIData(score *ScorePartWise) {
	fmt.Println("\nMIDI PREPARATION EXAMPLE")
	fmt.Println(strings.Repeat("-", 50))

	for _, part := range score.Part {
		// Get MIDI settings from part-list
		var midiChannel, midiProgram int
		var volume float32

		if score.PartList != nil {
			for _, scorePart := range score.PartList.ScorePart {
				if scorePart.ID == part.ID && len(scorePart.MidiInstrument) > 0 {
					midi := scorePart.MidiInstrument[0]
					if midi.MidiChannel != nil {
						midiChannel = *midi.MidiChannel
					}
					if midi.MidiProgram != nil {
						midiProgram = *midi.MidiProgram
					}
					if midi.Volume != nil {
						volume = *midi.Volume
					}
				}
			}
		}

		fmt.Printf("Part %s: Channel=%d, Program=%d, Volume=%.2f\n",
			part.ID, midiChannel, midiProgram, volume)

		// Process measures for MIDI events
		currentTime := 0.0
		divisions := 4 // default

		for _, measure := range part.Measure {
			if measure.Attributes != nil && measure.Attributes.Divisions != nil {
				divisions = *measure.Attributes.Divisions
			}

			// Process notes
			for _, note := range measure.Note {
				if note.Pitch != nil && note.Duration != nil {
					duration := getNoteDurationInBeats(note, divisions)
					fmt.Printf("  Time=%.2f: %s (%.2f beats)\n",
						currentTime, getPitchName(note), duration)

					if note.Chord == nil { // Not a chord, advance time
						currentTime += duration
					}
				}
			}
		}

		break // Just show first part as example
	}
}
