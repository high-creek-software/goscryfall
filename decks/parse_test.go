package decks_test

import (
	"bytes"
	"github.com/high-creek-software/goscryfall/decks"
	"log"
	"testing"
)

const (
	dinoDeck = `Commander
1 Gishath, Sun's Avatar

Deck
1 Apex Altisaur
1 Atzocan Seer
1 Bellowing Aegisaur
1 Burning Sun's Avatar
1 Carnage Tyrant
1 Etali, Primal Storm
1 Forerunner of the Empire
1 Ghalta, Primal Hunger
1 Giant Cindermaw
1 Goring Ceratops
1 Kinjalli's Caller
1 Kinjalli's Sunwing
1 Knight of the Stampede
1 Marauding Raptor
1 Otepec Huntmaster
1 Polyraptor
1 Quartzwood Crasher
1 Raging Regisaur
1 Raging Swordtooth
1 Ranging Raptors
1 Regisaur Alpha
1 Ripjaw Raptor
1 Runic Armasaur
1 Silverclad Ferocidons
1 Temple Altisaur
1 The Tarrasque
1 Thrashing Brontodon
1 Thundering Spineback
1 Topiary Stomper
1 Trapjaw Tyrant
1 Verdant Sun's Avatar
1 Wakening Sun's Avatar
1 Wayward Swordtooth
1 Zacama, Primal Calamity
1 Zetalpa, Primal Dawn
1 Beast Within
1 Boros Charm
1 Congregation at Dawn
1 Heroic Intervention
1 Return of the Wildspeaker
1 Swords to Plowshares
1 Tail Swipe
1 Commune with Dinosaurs
1 Cultivate
1 Farseek
1 Kodama's Reach
1 Majestic Genesis
1 Rampant Growth
1 Skyshroud Claim
1 Three Visits
1 Thunderherd Migration
1 Arcane Signet
1 Gruul Signet
1 Herald's Horn
1 Monster Manual
1 Sol Ring
1 Colossal Majesty
1 Elemental Bond
1 Garruk's Uprising
1 Mirari's Wake
1 Rhythm of the Wild
1 Unnatural Growth
1 Huatli, Warrior Poet
1 Canopy Vista
1 Cinder Glade
1 Clifftop Retreat
1 Command Tower
1 Evolving Wilds
1 Exotic Orchard
1 Jetmir's Garden
1 Jungle Shrine
1 Overgrown Farmland
1 Path of Ancestry
1 Rockfall Vale
1 Rogue's Passage
1 Rootbound Crag
1 Sacred Foundry
1 Secluded Courtyard
1 Stomping Ground
1 Sunpetal Grove
1 Temple Garden
1 Unclaimed Territory
3 Plains
4 Mountain
9 Forest
1 Lurking Predators

Sideboard
1 Shadow in the Warp`
)

func TestUnmarshal(t *testing.T) {
	buf := bytes.NewBufferString(dinoDeck)

	d, err := decks.Unmarshal("Dinos", buf)
	if err != nil {
		t.Error("error unmarshaling decks", err)
	}

	name := "Gishath, Sun's Avatar"
	if d.Commander.Name != name {
		t.Errorf("error parsing commander got: %s expected: %s", d.Companion.Name, name)
	}

	log.Println(*d)
}
