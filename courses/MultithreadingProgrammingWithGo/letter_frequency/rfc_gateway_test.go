package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// https://bismobaruno.medium.com/unit-test-http-request-in-golang-a96d146406e6

var rfcStubMaps = map[int]string{
	50: rfcStaticPageStubID50,
}

const rfcStaticPageStubID50 = "\n\nNetwork Working Group                                        J. Reynolds\nRequest for Comments: 1000                                     J. Postel\n                                                                     ISI\n                                                             August 1987\n\nObsoletes: RFCs 084, 100, 160, 170, 200, 598, 699, 800, 899, 999\n\n\n                THE REQUEST FOR COMMENTS REFERENCE GUIDE\n\n\nSTATUS OF THIS MEMO\n\n   This RFC is a reference guide for the Internet community which\n   summarizes of all the Request for Comments issued between April 1969\n   and March 1987.  This guide also categorizes the RFCs by topic.\n\nINTRODUCTION\n\n   This RFC Reference Guide is intended to provide a historical account\n   by categorizing and summarizing of the Request for Comments numbers 1\n   through 999 issued between the years 1969-1987.  These documents have\n   been crossed referenced to indicate which RFCs are current, obsolete,\n   or revised.  Distribution of this memo is unlimited.\n\nTHE ORIGINS OF RFCS - by Stephen D. Crocker\n\n   The DDN community now includes hundreds of nodes and thousands of\n   users, but once it was all a gleam in Larry Roberts' eye.  While much\n   of the development proceeded according to a grand plan, the design of\n   the protocols and the creation of the RFCs was largely accidental.\n\n   The procurement of the ARPANET was initiated in the summer of 1968 --\n   Remember Vietnam, flower children, etc?  There had been prior\n   experiments at various ARPA sites to link together computer systems,\n   but this was the first version to explore packet-switching on a grand\n   scale.  (\"ARPA\" didn't become \"DARPA\" until 1972.)  Unlike most of\n   the ARPA/IPTO procurements of the day, this was a competitive\n   procurement. The contract called for four IMPs to be delivered to\n   UCLA, SRI, UCSB and The University of Utah.  These sites were running\n   a Sigma 7 with the SEX operating system, an SDS 940 with the Genie\n   operating system, an IBM 360/75 with OS/MVT (or perhaps OS/MFT), and\n   a DEC PDP-10 with the Tenex operating system.  Options existed for\n   additional nodes if the first experiments were successful.  BBN won\n   the procurement in December 1968, but that gets ahead of this story.\n\n   Part of the reason for selecting these four sites was these were\n   existing ARPA computer science research contractors.  The precise\n   usage of the ARPANET was not spelled out in advance, and the research\n   community could be counted on to take some initiative.  To stimulate\n   this process, a meeting was called during the summer with\n   representatives from the selected sites, chaired by Elmer Shapiro\n\n\nReynolds & Postel                                               [Page 1]\n\f\n\n\nRFC 1000 - Request for Comments Reference Guide              August 1987\n\n\n   from SRI.  If memory serves me correctly, Jeff Rulifson came from\n   SRI, Ron Stoughton from UCSB, Steve Carr from Utah and I came from\n   UCLA. (Apologies to anyone I've left out; records are inaccessible or\n   lost at this point.)  At this point we knew only that the network was\n   coming, but the precise details weren't known.\n\n   That first meeting was seminal.  We had lots of questions -- how IMPs\n   and hosts would be connected, what hosts would say to each other, and\n   what applications would be supported.  No one had any answers, but\n   the prospects seemed exciting.  We found ourselves imagining all\n   kinds of possibilities -- interactive graphics, cooperating\n   processes, automatic data base query, electronic mail -- but no one\n   knew where to begin.  We weren't sure whether there was really room\n   to think hard about these problems; surely someone from the east\n   would be along by and by to bring the word.  But we did come to one\n   conclusion: We ought to meet again.  Over the next several months, we\n   managed to parlay that idea into a series of exchange meetings at\n   each of our sites, thereby setting the most important precedent in\n   protocol design.\n\n   The first few meetings were quite tenuous.  We had no official\n   charter...+311219 more"

func setupMockServer(code int, response json.RawMessage, mstimeOut ...int) *httptest.Server {
	timeout := 0
	if len(mstimeOut) > 0 {
		timeout = mstimeOut[0]
	}
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		})

	return httptest.NewServer(handler)
}

func Test_RFCGateway(t *testing.T) {
	type args struct {
		rfcID int
	}
	tests := []struct {
		name       string
		args       args
		want       string
		mockServer *httptest.Server
		wantErr    bool
	}{
		{
			name: "given a rfcID 50 should return all static page with success",
			args: args{
				rfcID: 50,
			},
			mockServer: setupMockServer(200, json.RawMessage(rfcStaticPageStubID50)),
			want:       rfcStaticPageStubID50,
			wantErr:    false,
		},
		{
			name: "given a rfcID 50 and httprequest reaches the time out out should return err",
			args: args{
				rfcID: 50,
			},
			mockServer: setupMockServer(500, json.RawMessage(rfcStaticPageStubID50), 1600),
			want:       "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.mockServer.Close()
			rfc := NewRFCGateway()
			rfc.URL = tt.mockServer.URL
			got, err := rfc.getRFC(tt.args.rfcID)

			if (err != nil) != tt.wantErr {
				t.Errorf("getRFCLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRFCLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
