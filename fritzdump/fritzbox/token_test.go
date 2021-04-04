package fritzbox

import (
	"testing"
)

func TestParseChallengeToken(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?><SessionInfo><SID>0000000000000000</SID><Challenge>9ffb32f5</Challenge><BlockTime>0</BlockTime><Rights></Rights></SessionInfo>`)
	expected := "9ffb32f5"

	result, err := parseChallengeToken(input)
	if err != nil {
		t.Fatalf("Expected parseChallengeToken(%s) not to fail, but got error: %s", input, err)
	}

	if expected != result {
		t.Errorf("Expected parseChallengeToken(%s) to return %s, but got: %s", input, expected, result)
	}
}
func TestParseChallengeFailing(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?><SessionInfo><SID>0000000000000000</SID><Challenge>9ffb32f</Challenge><BlockTime>0</BlockTime><Rights></Rights></SessionInfo>`)

	result, err := parseChallengeToken(input)
	if err == nil {
		t.Errorf("Expected parseChallengeToken(%s) to fail, but got no error and result: %s", input, result)
	}
}

func TestParseSID(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?><SessionInfo><SID>807f66b3bd80c8d1</SID><Challenge>6ccb29e6</Challenge><BlockTime>0</BlockTime><Rights><Name>Dial</Name><Access>2</Access><Name>App</Name><Access>2</Access><Name>HomeAuto</Name><Access>2</Access><Name>BoxAdmin</Name><Access>2</Access><Name>Phone</Name><Access>2</Access><Name>NAS</Name><Access>2</Access></Rights></SessionInfo>`)
	expected := "807f66b3bd80c8d1"

	result, err := parseSID(input)
	if err != nil {
		t.Fatalf("Expected parseSID(%s) not to fail, but got error: %s", input, err)
	}

	if expected != result {
		t.Errorf("Expected parseSID(%s) to return %s, but got: %s", input, expected, result)
	}
}
func TestParseSIDFailing(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?><SessionInfo><SID>807f66b3bd80c8d</SID><Challenge>6ccb29e6</Challenge><BlockTime>0</BlockTime><Rights><Name>Dial</Name><Access>2</Access><Name>App</Name><Access>2</Access><Name>HomeAuto</Name><Access>2</Access><Name>BoxAdmin</Name><Access>2</Access><Name>Phone</Name><Access>2</Access><Name>NAS</Name><Access>2</Access></Rights></SessionInfo>`)

	result, err := parseSID(input)
	if err == nil {
		t.Errorf("Expected parseSID(%s) to fail, but got no error and result: %s", input, result)
	}
}

func TestCreateAuthenticationToken(t *testing.T) {
	challengeToken := "9ffb32f5"
	password := "123fritz"
	expected := "3b61a21a4624159bce6aa8623e9f6f83"

	result := createAuthenticationToken(challengeToken, password)

	if expected != result {
		t.Errorf("Expected createAuthenticationToken(%s, %s) to return %s, but got: %s", challengeToken, password, expected, result)
	}
}
