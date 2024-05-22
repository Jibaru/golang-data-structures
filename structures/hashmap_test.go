package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewHashMap(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	if hm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", hm.Size())
	}
}

func TestHashMap_PushAt(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	hm.PushAt(2, "two")
	hm.PushAt(3, "three")
	hm.PushAt(4, "four")
	hm.PushAt(5, "five")
	hm.PushAt(6, "six")
	hm.PushAt(7, "seven")
	hm.PushAt(8, "eight")
	hm.PushAt(9, "nine")
	hm.PushAt(10, "ten")
	if hm.Size() != 10 {
		t.Errorf("Expected size 10, got %d", hm.Size())
	}
}

func TestHashMap_PushAt_Overwrite(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	hm.PushAt(1, "uno")
	if hm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", hm.Size())
	}
}

func TestHashMap_PopAt(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	hm.PushAt(2, "two")
	hm.PushAt(3, "three")
	hm.PushAt(4, "four")
	hm.PushAt(5, "five")
	hm.PushAt(6, "six")
	hm.PushAt(7, "seven")
	hm.PushAt(8, "eight")
	hm.PushAt(9, "nine")
	hm.PushAt(10, "ten")
	hm.PopAt(5)
	if hm.Size() != 9 {
		t.Errorf("Expected size 9, got %d", hm.Size())
	}
}

func TestHashMap_PopAt_NotFound(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	_, err := hm.PopAt(2)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestHashMap_PopAt_Resize(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 0; i < 100; i++ {
		hm.PopAt(i)
	}
	if hm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", hm.Size())
	}
}

func TestHashMap_PopAt_Resize2(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 99; i >= 0; i-- {
		hm.PopAt(i)
	}
	if hm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", hm.Size())
	}
}

func TestHashMap_PopAt_Resize3(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 0; i < 50; i++ {
		hm.PopAt(i)
	}
	if hm.Size() != 50 {
		t.Errorf("Expected size 50, got %d", hm.Size())
	}
}

func TestHashMap_GetAt(t *testing.T) {
	hm := structures.NewHashMap[string, int]()
	hm.PushAt("one", 1)
	hm.PushAt("two", 2)
	hm.PushAt("three", 3)
	hm.PushAt("four", 4)
	hm.PushAt("five", 5)

	var val int
	var err error

	val, err = hm.GetAt("one")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	val, err = hm.GetAt("two")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val, err = hm.GetAt("three")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	val, err = hm.GetAt("four")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}

	val, err = hm.GetAt("five")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestHashMap_GetAt_NotFound(t *testing.T) {
	hm := structures.NewHashMap[string, int]()
	hm.PushAt("one", 1)
	_, err := hm.GetAt("two")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestHashMap_GetAt_Resize(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 0; i < 100; i++ {
		hm.GetAt(i)
	}
}

func TestHashMap_GetAt_Resize2(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 99; i >= 0; i-- {
		hm.GetAt(i)
	}
}

func TestHashMap_GetAt_Resize3(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	for i := 0; i < 100; i++ {
		hm.PushAt(i, "value")
	}
	for i := 0; i < 50; i++ {
		hm.GetAt(i)
	}
}

func TestHashMap_Has(t *testing.T) {
	hm := structures.NewHashMap[string, int]()
	hm.PushAt("one", 1)
	hm.PushAt("two", 2)
	hm.PushAt("three", 3)
	hm.PushAt("four", 4)
	hm.PushAt("five", 5)

	if !hm.Has("one") {
		t.Error("Expected true, got false")
	}
	if !hm.Has("two") {
		t.Error("Expected true, got false")
	}
	if !hm.Has("three") {
		t.Error("Expected true, got false")
	}
	if !hm.Has("four") {
		t.Error("Expected true, got false")
	}
	if !hm.Has("five") {
		t.Error("Expected true, got false")
	}
}

func TestHashMap_Has_NotFound(t *testing.T) {
	hm := structures.NewHashMap[string, int]()
	hm.PushAt("one", 1)
	if hm.Has("two") {
		t.Error("Expected false, got true")
	}
}

func TestHashMap_Size(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	hm.PushAt(2, "two")
	hm.PushAt(3, "three")
	hm.PushAt(4, "four")
	hm.PushAt(5, "five")
	hm.PushAt(6, "six")
	hm.PushAt(7, "seven")
	hm.PushAt(8, "eight")
	hm.PushAt(9, "nine")
	hm.PushAt(10, "ten")
	if hm.Size() != 10 {
		t.Errorf("Expected size 10, got %d", hm.Size())
	}
}

func TestHashMap_Find(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	hm.PushAt(2, "two")
	hm.PushAt(3, "three")
	hm.PushAt(4, "four")
	hm.PushAt(5, "five")
	hm.PushAt(6, "six")
	hm.PushAt(7, "seven")
	hm.PushAt(8, "eight")
	hm.PushAt(9, "nine")
	hm.PushAt(10, "ten")

	if !hm.Find("one") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("two") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("three") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("four") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("five") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("six") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("seven") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("eight") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("nine") {
		t.Error("Expected true, got false")
	}
	if !hm.Find("ten") {
		t.Error("Expected true, got false")
	}
}

func TestHashMap_Find_NotFound(t *testing.T) {
	hm := structures.NewHashMap[int, string]()
	hm.PushAt(1, "one")
	if hm.Find("two") {
		t.Error("Expected false, got true")
	}
}
