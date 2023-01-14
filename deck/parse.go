package deck

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

const (
	sectionKeyCommander = "Commander"
	sectionKeyDeck      = "Deck"
	sectionKeySideboard = "Sideboard"
)

type sectionState int

const (
	sectionStateDeck sectionState = iota
	sectionStateCommander
	sectionStateSideboard
)

func Unmarshal(deckName string, r io.Reader) (*Deck, error) {
	parseState := sectionStateDeck
	deck := &Deck{Name: deckName}

	lines := bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)
	for lines.Scan() {
		line := lines.Text()

		if line == sectionKeyCommander {
			parseState = sectionStateCommander
			continue
		} else if line == sectionKeyDeck {
			parseState = sectionStateDeck
			continue
		} else if line == sectionKeySideboard {
			parseState = sectionStateSideboard
			continue
		} else if line == "" {
			continue
		}

		nameCount, err := parseLine(line)
		if err != nil {
			log.Println("error parsing name count:", line, err)
			continue
		}

		switch parseState {
		case sectionStateCommander:
			deck.Commander = nameCount
		case sectionStateDeck:
			deck.Deck = append(deck.Deck, nameCount)
		case sectionStateSideboard:
			deck.Sideboard = append(deck.Sideboard, nameCount)
		}

	}

	return deck, nil
}

func parseLine(line string) (NameCountPair, error) {
	line = strings.Trim(line, " ")
	firstSpaceIdx := strings.Index(line, " ")
	countPart := line[:firstSpaceIdx]
	namePart := line[firstSpaceIdx+1:]

	count, err := strconv.Atoi(countPart)
	if err != nil {
		return NameCountPair{}, err
	}

	return NameCountPair{Count: count, Name: namePart}, nil
}
