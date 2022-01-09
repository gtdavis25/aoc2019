package main

import "testing"

func TestCut(t *testing.T) {
	for _, test := range []struct {
		n, size int
		want    []int
	}{
		{1, 7, []int{1, 2, 3, 4, 5, 6, 0}},
		{-1, 7, []int{6, 0, 1, 2, 3, 4, 5}},
	} {
		d := newDeck(test.size)
		d.cut(test.n)
		if got := getCards(d); !equals(got, test.want) {
			t.Errorf("cut(%d): got %v, want %v", test.n, got, test.want)
		}
	}
}

func TestDealWithIncrement(t *testing.T) {
	for _, test := range []struct {
		n, size int
		want    []int
	}{
		{2, 7, []int{0, 4, 1, 5, 2, 6, 3}},
	} {
		d := newDeck(test.size)
		d.dealWithIncrement(test.n)
		if got := getCards(d); !equals(got, test.want) {
			t.Errorf("dealWithIncrement(%d): got %v, want %v", test.n, got, test.want)
		}
	}
}

func TestDealIntoNewStack(t *testing.T) {
	for _, test := range []struct {
		size int
		want []int
	}{
		{7, []int{6, 5, 4, 3, 2, 1, 0}},
	} {
		d := newDeck(test.size)
		d.dealIntoNewStack()
		if got := getCards(d); !equals(got, test.want) {
			t.Errorf("dealIntoNewStack(): got %v, want %v", got, test.want)
		}
	}
}

func TestShuffle(t *testing.T) {
	for _, test := range []struct {
		size         int
		instructions []string
		want         []int
	}{
		{
			size:         5,
			instructions: []string{"cut 1", "deal with increment 2"},
			want:         []int{1, 4, 2, 0, 3},
		},
	} {
		d := newDeck(test.size)
		d.shuffle(test.instructions)
		if got := getCards(d); !equals(got, test.want) {
			t.Errorf("shuffle: got %v, want %v", got, test.want)
		}
	}
}

func getCards(d *deck) []int {
	cards := make([]int, d.size)
	for i := range cards {
		cards[i] = d.card(i)
	}

	return cards
}

func equals(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}
