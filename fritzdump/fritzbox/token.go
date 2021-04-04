package fritzbox

import (
	"fmt"
	"regexp"

	"github.com/eg5846/getting-started-with-pcap/fritzdump/utils"
)

var (
	CHALLENGE_TOKEN_PATTERN = regexp.MustCompile(`<Challenge>([a-z0-9]{8}?)</Challenge>`)
	SID_PATTERN             = regexp.MustCompile(`<SID>([a-z0-9]{16}?)</SID>`)
)

func parse(bytes []byte, regex *regexp.Regexp) (string, error) {
	match := regex.FindSubmatch((bytes))
	if len(match) != 2 {
		err := fmt.Errorf("Unexpected match len: %d", len(match))
		return "", err
	}

	return string(match[1]), nil
}

// Parse challenge token from XML reponse (quick and dirty)
// <?xml version="1.0" encoding="utf-8"?>
// <SessionInfo>
//   <SID>0000000000000000</SID>
//   <Challenge>9ffb32f5</Challenge>
//   <BlockTime>0</BlockTime>
//   <Rights></Rights>
// </SessionInfo>
func parseChallengeToken(bytes []byte) (string, error) {
	return parse(bytes, CHALLENGE_TOKEN_PATTERN)
}

// Parse SID from XML response (quick and dirty)
// <?xml version="1.0" encoding="utf-8"?>
// <SessionInfo>
//   <SID>807f66b3bd80c8d1</SID>
//   <Challenge>6ccb29e6</Challenge>
//   <BlockTime>0</BlockTime>
//   <Rights>
//     <Name>Dial</Name>
//     <Access>2</Access>
//     <Name>App</Name>
//     <Access>2</Access>
//     <Name>HomeAuto</Name>
//     <Access>2</Access>
//     <Name>BoxAdmin</Name>
//     <Access>2</Access>
//     <Name>Phone</Name>
//     <Access>2</Access>
//     <Name>NAS</Name>
//     <Access>2</Access>
//   </Rights>
// </SessionInfo>
func parseSID(bytes []byte) (string, error) {
	return parse(bytes, SID_PATTERN)
}

func createAuthenticationToken(challengeToken string, password string) string {
	challengePassword := fmt.Sprintf("%s-%s", challengeToken, password)
	challengePasswordUtf16 := utils.ToUtf16(challengePassword)
	challengePasswordBytes := utils.ToBytes(challengePasswordUtf16)
	return utils.CalculateMd5Sum(challengePasswordBytes)
}
