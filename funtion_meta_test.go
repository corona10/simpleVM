package simplevm

import "testing"

func TestFuntionMetaData(t *testing.T) {
	m := FunctionMetaData{name: "foo", nargs: 3, nlocals: 3, address: 0}
	name := m.GetName()
	nargs := m.GetNArgs()
	nlocals := m.GetNLocals()
	address := m.GetAddress()
	if name != "foo" {
		t.Errorf("m.GetName should return %v not %v", "foo", name)
	}

	if nargs != 3 {
		t.Errorf("m.GetNArgs should return %v not %v", 3, nargs)
	}

	if nlocals != 3 {
		t.Errorf("m.GetNLocals should return %v not %v", 3, nlocals)
	}

	if address != 0 {
		t.Errorf("m.GetAddress should return %v not %v", 3, address)
	}
}
