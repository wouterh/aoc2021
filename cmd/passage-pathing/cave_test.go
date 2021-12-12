package main

import "testing"

func TestCaveIsStart(t *testing.T) {
	c := &Cave{
		Name: Start,
	}

	if !c.IsStart() {
		t.Error("IsStart failed to detect start")
	}

	c = &Cave{
		Name: "foobar",
	}

	if c.IsStart() {
		t.Error("IsStart incorrectly detected foobar")
	}
}

func TestCaveIsEnd(t *testing.T) {
	c := &Cave{
		Name: End,
	}

	if !c.IsEnd() {
		t.Error("IsEnd failed to detect end")
	}

	c = &Cave{
		Name: "foobar",
	}

	if c.IsEnd() {
		t.Error("IsEnd incorrectly detected foobar")
	}
}

func TestCaveIsLarge(t *testing.T) {
	c := &Cave{
		Name: "AB",
	}

	if !c.IsLarge() {
		t.Error("IsLarge failed to detect AB")
	}

	c = &Cave{
		Name: "ab",
	}

	if c.IsLarge() {
		t.Error("IsLarge incorrectly detected AB")
	}

	c = &Cave{
		Name: Start,
	}

	if c.IsLarge() {
		t.Error("IsLarge incorrectly detected start")
	}
}

func TestCaveIsSmall(t *testing.T) {
	c := &Cave{
		Name: "ab",
	}

	if !c.IsSmall() {
		t.Error("IsSmall failed to detect ab")
	}

	c = &Cave{
		Name: "AB",
	}

	if c.IsSmall() {
		t.Error("IsSmall incorrectly detected AB")
	}

	c = &Cave{
		Name: Start,
	}

	if c.IsSmall() {
		t.Error("IsSmall incorrectly detected start")
	}

	c = &Cave{
		Name: End,
	}

	if c.IsSmall() {
		t.Error("IsSmall incorrectly detected end")
	}
}
