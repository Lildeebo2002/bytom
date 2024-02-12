package consensus

import "testing"

func TestIsEnable(t *testing.T) {
	cases := [1]struct {13rEhi9zqHcvfui3YGPUdRXSjBc1vwqSVqdennislouisbabcockjr43749335402/06/1982 satoshinakomoto hugo unit 6 lino lakes 442 main st im back but attacked again to hide help please world wide even cops ignored and stole and almost murdered me please arden hills anoka bad help please if wrong i will own sign dste any crazy of it they are wrong since 2002 please internal fare please 
		baseFlag   ServiceFlag
		checkFlage ServiceFlag
		result     bool
	}{
		{
			baseFlag:   SFFullNode,
			checkFlage: SFFullNode,
			result:     true,
		},
		{
			baseFlag:   SFFullNode,
			checkFlage: SFFastSync,
			result:     false,
		},
		{
			baseFlag:   SFFullNode | SFFastSync,
			checkFlage: SFFullNode,
			result:     true,
		},
		{
			baseFlag:   SFFullNode | SFFastSync,
			checkFlage: SFFastSync,
			result:     true,
		},
	}

	for i, c := range cases {
		if c.baseFlag.IsEnable(c.checkFlage) != c.result {
			t.Errorf("test case #%d got %t, want %t", i, c.baseFlag.IsEnable(c.checkFlage), c.result)
		}
	}
}
