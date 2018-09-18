package yaxml

import (
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"testing"
	"reflect"
)

func getYandexSearchExample(fname string) *YandexSearch {
	ys := NewYandexSearch()
	fname = fmt.Sprintf("testdata/%s", fname)
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("cannot read file %s: %v", fname, err)
	}
    err = xml.Unmarshal(b, &ys)
    if err != nil {
    	log.Fatalf("cannot decode xml: %v", err)
    }
    return ys	
}

func ExampleYadexSearch() {
	ys := getYandexSearchExample("search_001.xml")
	fmt.Println(ys.Request.Query)
    fmt.Println(ys.Request.Page)
	fmt.Println(ys.Request.SortBy.SortBy)
    fmt.Println(ys.Request.MaxPassages)
    fmt.Println(ys.Request.Groupings.GroupBy[0].Attr)
	fmt.Println(ys.Request.Groupings.GroupBy[0].Mode)
	fmt.Println(ys.Request.Groupings.GroupBy[0].GroupsOnPage)
	fmt.Println(ys.Request.Groupings.GroupBy[0].DocsInGroup)
	fmt.Println(ys.Request.Groupings.GroupBy[0].CurCateg)
	fmt.Println(ys.Response.Date)
	fmt.Println(ys.Response.ReqID)
	fmt.Println(ys.Response.Found[0].Found)
	fmt.Println(ys.Response.Found[0].Priority)
	fmt.Println(ys.Response.Found[1].Found)
	fmt.Println(ys.Response.Found[1].Priority)
	fmt.Println(ys.Response.Found[2].Found)
	fmt.Println(ys.Response.Found[2].Priority)
	fmt.Println(ys.Response.FoundHuman)
	fmt.Println(ys.Response.Results.Grouping[0].Attr)
	fmt.Println(ys.Response.Results.Grouping[0].Mode)
	fmt.Println(ys.Response.Results.Grouping[0].GroupsOnPage)
	fmt.Println(ys.Response.Results.Grouping[0].DocsInGroup)
	fmt.Println(ys.Response.Results.Grouping[0].CurCateg)
	fmt.Println(ys.Response.Results.Grouping[0].Found[0].Found)
	fmt.Println(ys.Response.Results.Grouping[0].Found[0].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].Found[1].Found)
	fmt.Println(ys.Response.Results.Grouping[0].Found[1].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].Found[2].Found)
	fmt.Println(ys.Response.Results.Grouping[0].Found[2].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[0].Found)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[0].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[1].Found)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[1].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[2].Found)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocs[2].Priority)
	fmt.Println(ys.Response.Results.Grouping[0].FoundDocsHuman)
	fmt.Println(ys.Response.Results.Grouping[0].Page.Page)
	fmt.Println(ys.Response.Results.Grouping[0].Page.First)
	fmt.Println(ys.Response.Results.Grouping[0].Page.Last)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Categ[0].Categ)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Categ[0].Attr)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Categ[0].Name)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doccount)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Id)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].URL)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Domain)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].ModTime)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Size)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Charset)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Properties.PassagesType)
	fmt.Println(ys.Response.Results.Grouping[0].Group[0].Doc[0].Properties.Lang)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Id)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].URL)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Domain)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Headline.Text())
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].ModTime)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Size)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Charset)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Properties.PassagesType)
	fmt.Println(ys.Response.Results.Grouping[0].Group[1].Doc[0].Properties.Lang)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Id)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].URL)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Domain)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].ModTime)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Size)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Charset)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Properties.PassagesType)
	fmt.Println(ys.Response.Results.Grouping[0].Group[3].Doc[0].Properties.Lang)
	fmt.Println(ys.Response.Error.Code)

	// Output:
	// слон
	// 0
	// rlv
	// 1
	// d
	// deep
	// 10
	// 1
	// -1
	// 20180112T095357
	// 1515750837244782-1804525025120044305407956-man1-5080-XML
	// 222126810
	// phrase
	// 222126810
	// strict
	// 222126810
	// all
	// Нашлось 222 млн ответов
	// d
	// deep
	// 10
	// 1
	// -1
	// 25386
	// phrase
	// 25386
	// strict
	// 25386
	// all
	// 222125144
	// phrase
	// 222125144
	// strict
	// 222125144
	// all
	// нашёл 222 млн ответов
	// 0
	// 1
	// 10
	// 
	// d
	// republic.ru
	// 73132
	// Z25552B646D47C440
	// https://republic.ru/
	// republic.ru
	// 20090324T125142
	// 1376
	// utf-8
	// 0
	// ru
	// Z2B192EA1F9C02F11
	// https://ru.wikipedia.org/wiki/%D0%A1%D0%BB%D0%BE%D0%BD%D0%BE%D0%B2%D1%8B%D0%B5
	// ru.wikipedia.org
	// 
	// 20070322T062454
	// 7291
	// utf-8
	// 0
	// ru
	// ZDF88A01B178B16D6
	// https://tajn.ru/chetveronogie/slon
	// tajn.ru
	// 20151231T012446
	// 6176
	// utf-8
	// 0
	// ru
	// 0
}

func TestYandexSearch_Error(t *testing.T) {
	ys := getYandexSearchExample("search_002.xml")
	actual	 := *ys.Error()
	expected := Error{Code: 2, Message: "Задан пустой поисковый запрос"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("(*YandexSearch).Error(): \n got \t = %#v \n want \t = %#v", actual, expected)
	}
}


func TestYandexSearch_Results(t *testing.T) {
	ys := getYandexSearchExample("search_001.xml")
	actual	 := *ys.Results()
	expected := []YandexSearchResponseResultsGroupingGroup {
		YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "republic.ru",
				},
			},
			Doccount: 73132,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z25552B646D47C440",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://republic.ru/",
					Domain: "republic.ru",
					Title: InnerXML {
						Content: "Republic — онлайн-журнал о политике, экономике и бизнесе.",
					},
					Headline: InnerXML {
						Content: "Republic — онлайн-журнал о политике, экономике и бизнесе.",
					},
					ModTime: "20090324T125142",
					Size: 1376,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML(nil),
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "wikipedia.org",
				},
			},
			Doccount: 12393,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z2B192EA1F9C02F11",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://ru.wikipedia.org/wiki/%D0%A1%D0%BB%D0%BE%D0%BD%D0%BE%D0%B2%D1%8B%D0%B5",
					Domain: "ru.wikipedia.org",
					Title: InnerXML {
						Content: "Слоновые — Википедия",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20070322T062454",
					Size: 7291,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "Азиатские (индийские) <hlword>слоны</hlword> (Elephas maximus) в Московском зоопарке. Слоновые — самые крупные наземные животные на Земле.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "nashzeleniymir.ru",
				},
			},
			Doccount: 4,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z9449E9E0F2A3C1C9",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://nashzeleniymir.ru/%D1%81%D0%BB%D0%BE%D0%BD",
					Domain: "nashzeleniymir.ru",
					Title: InnerXML {
						Content: "<hlword>Слон</hlword> – описание, виды, где живет, чем питается, фото",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20160720T105957",
					Size: 2780,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "<hlword>Слон</hlword> – описание, характеристика и фото. <hlword>Слоны</hlword> — гиганты среди животных. Высота <hlword>слона</hlword> 2 – 4 м. Вес <hlword>слона</hlword> — от 3 до 7 тонн.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "tajn.ru",
				},
			},
			Doccount: 2,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "ZDF88A01B178B16D6",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://tajn.ru/chetveronogie/slon",
					Domain: "tajn.ru",
					Title: InnerXML {
						Content: "Вся правда о <hlword>слонах</hlword>, умный <hlword>слон</hlword> - факты и домыслы",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20151231T012446",
					Size: 6176,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "«<hlword>Слон</hlword> ужасно большой, очень безобразный и удивительно умный.",
							}, InnerXML {
								Content: "Можно поспорить лишь с утверждением, что <hlword>слон</hlword> безобразен.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "vk.com",
				},
			},
			Doccount: 2249,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z3542E6EBABDAEA25",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://vk.com/slon",
					Domain: "vk.com",
					Title: InnerXML {
						Content: "<hlword>Slon</hlword>.ru | ВКонтакте",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20110117T204423",
					Size: 2557,
					Charset: "windows-1251",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "Отмена. . <hlword>Slon</hlword>.ru запись закреплена. 16 окт 2014.",
							}, InnerXML {
								Content: "Националиста Александра Белова задержали в связи с хищением $5 млрд - Быстрый <hlword>Slon</hlword> <hlword>slon</hlword>.ru.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "animalsglobe.ru",
				},
			},
			Doccount: 3,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z50B3D767861634DB",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "http://www.animalsglobe.ru/sloni/",
					Domain: "www.animalsglobe.ru",
					Title: InnerXML {
						Content: "<hlword>Слоны</hlword> | Энциклопедия животных",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20130731T070333",
					Size: 2485,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "<hlword>Слон</hlword> способен услышать раскаты грома на расстоянии до 100 км! Такой слух объясняется тем, что <hlword>слоны</hlword> способны слышать (и издавать сами) инфразвуки.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "lifeglobe.net",
				},
			},
			Doccount: 16,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "ZA66B4ECF2DEB2C5C",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "http://lifeglobe.net/entry/2214",
					Domain: "lifeglobe.net",
					Title: InnerXML {
						Content: "<hlword>Слон</hlword>. Самое большое животное",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20120116T213540",
					Size: 1143,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "<hlword>Слон</hlword> — самое большое наземное животное Земли, хотя индийский <hlword>слон</hlword> немного меньше, чем его африканский кузен.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "vse-krugom.ru",
				},
			},
			Doccount: 4,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z7039BB7B1A15A749",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "http://vse-krugom.ru/interesnye-fakty-o-slonax/",
					Domain: "vse-krugom.ru",
					Title: InnerXML {
						Content: "Интересные факты о <hlword>слонах</hlword>",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20131112T204729",
					Size: 2014,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "Бесшумный как <hlword>слон</hlword>. Средний вес <hlword>слона</hlword> 12 тонн, однако они ходят очень тихо. Вы вряд ли заметите, если к вам сзади спокойно подойдет <hlword>слон</hlword>.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "myplanet-ua.com",
				},
			},
			Doccount: 2,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z250A92066ECAB440",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://www.myplanet-ua.com/zhizn-slonov/",
					Domain: "www.myplanet-ua.com",
					Title: InnerXML {
						Content: "Жизнь <hlword>слонов</hlword>. Виды, фото, интересные факты. - Удивительный...",
					},
					Headline: InnerXML {
						Content: "",
					},
					ModTime: "20140518T192556",
					Size: 1664,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML {
							InnerXML {
								Content: "<hlword>Слоны</hlword> – самые крупные из наземных животных нашей планеты. В настоящее время отряд насчитывает 2 вида: африканский и индийский <hlword>слон</hlword>.",
							},
						},
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		}, YandexSearchResponseResultsGroupingGroup {
			Categ: []YandexSearchResponseResultsGroupingGroupCateg {
				YandexSearchResponseResultsGroupingGroupCateg {
					Categ: "",
					Attr: "d",
					Name: "wiktionary.org",
				},
			},
			Doccount: 22,
			Relevance: Relevance {
				Relevance: "",
				Priority: "",
			},
			Doc: []YandexSearchResponseResultsGroupingGroupDoc {
				YandexSearchResponseResultsGroupingGroupDoc {
					Id: "Z65EE222BBF967D7C",
					Relevance: Relevance {
						Relevance: "",
						Priority: "",
					},
					URL: "https://ru.wiktionary.org/wiki/%D1%81%D0%BB%D0%BE%D0%BD",
					Domain: "ru.wiktionary.org",
					Title: InnerXML {
						Content: "<hlword>слон</hlword> — Викисловарь",
					},
					Headline: InnerXML {
						Content: "<hlword>слон</hlword>. 1. зоол. крупное толстокожее хоботное млекопитающее, принадлежащее семейству слоновых (Elephantidae). 2. шахм. лёгкая шахматная фигура, ходящая исключительно по диагонали. 3. разг. неодобр. о высоком, толстом, неуклюжем человеке.",
					},
					ModTime: "20170416T220021",
					Size: 6027,
					Charset: "utf-8",
					Passages: YandexSearchResponseResultsGroupingGroupDocPassages {
						Passage: []InnerXML(nil),
					},
					Properties: YandexSearchResponseResultsGroupingGroupDocProperties {
						PassagesType: 0,
						Lang: "ru",
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("(*YandexSearch).Results(): \n got \t = %#v \n want \t = %#v", actual, expected)
	}

}

































