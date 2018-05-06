package simplevm

import "testing"

func TestFuntionMetaData(t *testing.T) {
	m := FunctionMetaData{name: "foo", nargs: 3, nlocals: 3, address: 0}
	m_name := m.GetName()
	m_nargs := m.GetNArgs()
	m_nlocals := m.GetNLocals()
	m_address := m.GetAddress()
	if m_name != "foo" {
		t.Errorf("m.GetName should return %v not %v", "foo", m_name)
	}

	if m_nargs != 3 {
		t.Errorf("m.GetNArgs should return %v not %v", 3, m_nargs)
	}

	if m_nlocals != 3 {
		t.Errorf("m.GetNLocals should return %v not %v", 3, m_nlocals)
	}

	if m_address != 0 {
		t.Errorf("m.GetAddress should return %v not %v", 3, m_address)
	}
}
