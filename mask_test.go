package idmask

import (
	"strconv"
	"testing"
)

func TestIDMask(t *testing.T) {
	originalIDs := []uint{1, 3, 5, 8}
	bits := Mask(originalIDs)
	btoi, err := strconv.ParseInt(reverse("010101001"), 2, 64)
	if err != nil {
		t.Fatalf("unable to converte binary to int - `%v`", err)
	}
	if got, want := bits, int(btoi); got != want {
		t.Errorf("got `%d` (%b), want `%d` (%b)", got, got, want, want)
	}

	restoredIDs := Unmask(bits)
	if got, want := restoredIDs, originalIDs; !listEqual(got, want) {
		t.Errorf("got %v, want %v", restoredIDs, originalIDs)
	}
}

func TestIDMaskLarger(t *testing.T) {
	originalIDs := []uint{1, 3, 5, 8, 140, 150, 6, 2, 62, 63, 63, 127, 128, 129}

	bits := Mask(originalIDs)
	restoredIDs := Unmask(bits)

	if got, want := restoredIDs, originalIDs; !listEqual(got, want) {
		t.Errorf("got %v, want %v", restoredIDs, originalIDs)
	}
}

func listEqual(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	hashA := make(map[uint]int)

	for _, element := range a {
		if value, ok := hashA[element]; ok {
			hashA[element] = value + 1
		} else {
			hashA[element] = 1
		}
	}

	for _, element := range b {
		if value, ok := hashA[element]; ok {
			hashA[element] = value - 1
		} else {
			return false
		}
	}

	for _, value := range hashA {
		if value != 0 {
			return false
		}
	}

	return true
}

func reverse(s string) string {
	reversed := make([]rune, len(s))
	for i, char := range s {
		reversed[len(s)-(i+1)] = char
	}
	return string(reversed)
}
