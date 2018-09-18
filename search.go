package yaxml

import (
	"encoding/xml"
    "regexp"
    "strings"
)

// InnerXML required for fields where <hlword> could be found
type InnerXML struct {
	Content string `xml:",innerxml" json:"content"`
}

// Extract highlighted text
func (x *InnerXML) HL() []string {
	s := x.Content
	s = strings.Replace(s, "</hlword> <hlword>", " ", -1)

	f := regexp.MustCompile(`<hlword>(.*?)</hlword>`)
	r := []string{}
	for _, pair := range f.FindAllStringSubmatch(s, -1) {
		r = append(r, pair[1])
	}

	return r
}

// Extract clean text
func (x *InnerXML) Text() string {
	s := x.Content
	s = strings.Replace(s, "<hlword>", "", -1)
	s = strings.Replace(s, "</hlword>", "", -1)
	return s
}

type YandexSearch struct {
	Request  YandexSearchRequest  `xml:"request" json:"request"`
	Response YandexSearchResponse `xml:"response" json:"response"`
}

func (ys *YandexSearch) Error() *Error {
	return &ys.Response.Error
}

func (ys *YandexSearch) Results() *[]YandexSearchResponseResultsGroupingGroup {
	return &ys.Response.Results.Grouping[0].Group
}

type YandexSearchRequest struct {
	XMLName 	xml.Name 						`xml:"request" json:"-"`
	Query       string              			`xml:"query" json:"query"`
	Page        int                 			`xml:"page" json:"page"`
	SortBy      YandexSearchRequestSortBy 		`xml:"sortby" json:"sortBy"`
	MaxPassages int 							`xml:"maxpassages" json:"maxPassages"`
	Groupings   YandexSearchRequestGroupings 	`xml:"groupings" json:"groupings"`
}

func (req *YandexSearchRequest) SetSortBy(sort, order string) {
	req.SortBy = YandexSearchRequestSortBy{
		SortBy: sort,
		Order: 	order,
	}
}

func (req *YandexSearchRequest) SetGroupBy(mode string, groupsOnPage, docsInGroup int) {
	var attr string
	if mode == "deep" {
		attr = "d"
	}
	req.Groupings.GroupBy = []YandexSearchRequestGroupingsGroupBy{
		YandexSearchRequestGroupingsGroupBy {
			Attr: attr,
			Mode: mode,
			GroupsOnPage: groupsOnPage,
			DocsInGroup: docsInGroup,
		},
	}
}

type YandexSearchRequestSortBy struct {
	SortBy   string 	`xml:",chardata" json:"sortBy"`
	Order    string 	`xml:"order,attr" json:""`
	Priority string 	`xml:"priority,attr" json:""`
}

type YandexSearchRequestGroupings struct {
	GroupBy 	[]YandexSearchRequestGroupingsGroupBy 	`xml:"groupby" json:"groupBy"`
}

type YandexSearchRequestGroupingsGroupBy struct {
	Attr         string `xml:"attr,attr" json:"attr"`
	Mode         string `xml:"mode,attr" json:"mode"`
	GroupsOnPage int    `xml:"groups-on-page,attr" json:"groupsOnPage"`
	DocsInGroup  int    `xml:"docs-in-group,attr" json:"docsInGroup"`
	CurCateg     int    `xml:"curcateg,attr" json:"-"` // ALARM if > -1
}

func NewYandexSearchRequest() *YandexSearchRequest {
	req := new(YandexSearchRequest)
	req.Page = 0
	req.SortBy = YandexSearchRequestSortBy{
		SortBy: "rlv",
		Order: "descending",
		Priority: "no",
	}
	req.MaxPassages = 4
	req.Groupings.GroupBy = []YandexSearchRequestGroupingsGroupBy{
		YandexSearchRequestGroupingsGroupBy {
			Attr: "d",
			Mode: "deep",
			GroupsOnPage: 100,
			DocsInGroup: 1,
		},
	}
	return req
}

type YandexSearchResponse struct {
	Date       	string                     				`xml:"date,attr" json:"date"`
	ReqID      	string                     				`xml:"reqid" json:"reqID"`
	Found      	[]YandexSearchResponseFound         	`xml:"found" json:"found"`
	FoundHuman 	string                     				`xml:"found-human" json:"foundHuman"`
	Misspell   	YandexSearchResponseMisspellOrReask 	`xml:"misspell" json:"misspell"`
	Reask      	YandexSearchResponseMisspellOrReask 	`xml:"reask" json:"reask"`
	Results    	YandexSearchResponseResults         	`xml:"results" json:"results"`
	Error 		Error 									`xml:"error" json:"error"`
}

type YandexSearchResponseFound struct {
	Found    int    `xml:",chardata" json:"found"`
	Priority string `xml:"priority,attr" json:"priority"`
}

type YandexSearchResponseMisspellOrReask struct {
	Rule       string   `xml:"rule" json:"rule"` // Misspell, KeyboardLayout, Volapyuk (ru translit)
	SourceText InnerXML `xml:"source-text" json:"sourceText"`
	Text       string   `xml:"text" json:"text"`
}

type YandexSearchResponseResults struct {
	Grouping []YandexSearchResponseResultsGrouping `xml:"grouping" json:"grouping"`
}

type YandexSearchResponseResultsGrouping struct {
	Attr           string                                			`xml:"attr,attr" json:"attr"`
	Mode           string                                			`xml:"mode,attr" json:"mode"`
	GroupsOnPage   int                                   			`xml:"groups-on-page,attr" json:"groupsOnPage"`
	DocsInGroup    int                                   			`xml:"docs-in-group,attr" json:"docsInGroup"`
	CurCateg       int                                   			`xml:"curcateg,attr" json:"-"`
	Found          []YandexSearchResponseResultsGroupingFound     	`xml:"found" json:"found"`
	FoundDocs      []YandexSearchResponseResultsGroupingFoundDocs 	`xml:"found-docs" json:"foundDocs"`
	FoundDocsHuman string                                			`xml:"found-docs-human" json:"foundDocsHuman"`
	Page           YandexSearchResponseResultsGroupingPage        	`xml:"page" json:"page"`
	Group          []YandexSearchResponseResultsGroupingGroup     	`xml:"group" json:"group"`
}

type YandexSearchResponseResultsGroupingFound struct {
	YandexSearchResponseFound
}

type YandexSearchResponseResultsGroupingFoundDocs struct {
	YandexSearchResponseFound
}

type YandexSearchResponseResultsGroupingPage struct {
	Page  int `xml:",chardata" json:"page"`
	First int `xml:"first,attr" json:"first"`
	Last  int `xml:"last,attr" json:"last"`
}

type YandexSearchResponseResultsGroupingGroup struct {
	Categ     []YandexSearchResponseResultsGroupingGroupCateg 	`xml:"categ" json:"-"` // ALARM !<categ attr=">d<"
	Doccount  int                                    			`xml:"doccount" json:"doccount"`
	Relevance Relevance                           				`xml:"relevance" json:"relevance"`
	Doc       []YandexSearchResponseResultsGroupingGroupDoc 	`xml:"doc" json:"doc"`
}

type YandexSearchResponseResultsGroupingGroupCateg struct {
	Categ string `:",chardata" json:"-"` // ALARM!!
	Attr  string `xml:"attr,attr" json:"attr"`
	Name  string `xml:"name,attr" json:"name"` // hostname
}

type YandexSearchResponseResultsGroupingGroupDoc struct {
	Id         string                                       			`xml:"id,attr" json:"id"`
	Relevance  Relevance                                 				`xml:"relevance" json:"relevance"`
	URL        string                                       			`xml:"url" json:"URL"`
	Domain     string                                       			`xml:"domain" json:"domain"`
	Title      InnerXML                                     			`xml:"title" json:"title"`
	Headline   InnerXML                                     			`xml:"headline" json:"headline"` // i.e. meta-description
	ModTime    string                                       			`xml:"modtime" json:"modTime"`
	Size       int                                          			`xml:"size" json:"size"`
	Charset    string                                       			`xml:"charset" json:"charset"`
	Passages   YandexSearchResponseResultsGroupingGroupDocPassages   	`xml:"passages" json:"passages"`
	Properties YandexSearchResponseResultsGroupingGroupDocProperties 	`xml:"properties" json:"properties"`
}

type YandexSearchResponseResultsGroupingGroupDocPassages struct {
	Passage []InnerXML `xml:"passage" json:"passage"`
}

type YandexSearchResponseResultsGroupingGroupDocProperties struct {
	// 0 — standart passage (on-page)
	// 1 — anchor based (on-link)
	PassagesType int    `xml:"_PassagesType" json:"passagesType"` // ALARM
	Lang         string `xml:"lang" json:"lang"`
}

type Relevance struct { // We miss you...
	Relevance string `xml:",chardata" json:"relevance"`
	Priority  string `xml:"priority,attr" json:"priority"`
}

func NewYandexSearch() *YandexSearch {
	ys := new(YandexSearch)
	return ys
}

