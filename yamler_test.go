package main

import "testing"

var goodYaml = []byte(`
# this yaml is pretty cool
name: my cool yaml
data:
  goes: here
`)

var badYaml = []byte(`
# this isn't very cool yaml
name: my not so cool yaml
data:
  whoops: [who, needs, to, close, this
`)

func TestGoodParser(t *testing.T) {
	succ, err := parseYaml(goodYaml)
	if !succ {
		t.Errorf("Somehow it failed to parse good yaml: %s (%s)", succ, err)
	}
}

func TestBadParser(t *testing.T) {
	succ, err := parseYaml(badYaml)
	if succ {
		t.Errorf("Somehow it successfully parsed bad yaml: %s (%s)", succ, err)
	}
}
