package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
)

type couterMO_data struct {
	number    string // Номер дела
	url       string // Ссылка на дело
	datesP    string // Дата поступления
	kategory  string // Категория
	applicant string // Истец
	defendant string // Ответчик
	judge     string // Судья
	datesR    string // Дата решения
	decision  string // Решение
	datesZ    string //Дата вступления в законную силу
	akts      string // Судебные акты
}

var index_global int = 2

// const postfix string = "/modules.php?name=sud_delo&srv_num=1&name_op=r&page=MYPAGE&vnkod=50RS0002&srv_num=1&name_op=r&delo_id=1540005&case_type=0&new=0&G1_PARTS__NAMESS=&g1_case__CASE_NUMBERSS=&g1_case__JUDICIAL_UIDSS=&delo_table=g1_case&g1_case__ENTRY_DATE1D=DATEOT&g1_case__ENTRY_DATE2D=DATEDO&lawbookarticles%5B0%5D=%D1%EF%EE%F0%FB%2C+%E2%EE%E7%ED%E8%EA%E0%FE%F9%E8%E5+%E8%E7+%F2%F0%F3%E4%EE%E2%FB%F5+%EE%F2%ED%EE%F8%E5%ED%E8%E9&G1_CASE__JUDGE=&g1_case__RESULT_DATE1D=&g1_case__RESULT_DATE2D=&G1_CASE__RESULT=&G1_CASE__BUILDING_ID=&G1_CASE__COURT_STRUCT=&G1_EVENT__EVENT_NAME=&G1_EVENT__EVENT_DATEDD=&G1_PARTS__PARTS_TYPE=&G1_PARTS__INN_STRSS=&G1_PARTS__KPP_STRSS=&G1_PARTS__OGRN_STRSS=&G1_PARTS__OGRNIP_STRSS=&G1_RKN_ACCESS_RESTRICTION__RKN_REASON=&g1_rkn_access_restriction__RKN_RESTRICT_URLSS=&g1_requirement__ACCESSION_DATE1D=&g1_requirement__ACCESSION_DATE2D=&G1_REQUIREMENT__CATEGORY=&g1_requirement__ESSENCESS=&g1_requirement__JOIN_END_DATE1D=&g1_requirement__JOIN_END_DATE2D=&G1_REQUIREMENT__PUBLICATION_ID=&G1_DOCUMENT__PUBL_DATE1D=&G1_DOCUMENT__PUBL_DATE2D=&G1_CASE__VALIDITY_DATE1D=&G1_CASE__VALIDITY_DATE2D=&G1_ORDER_INFO__ORDER_DATE1D=&G1_ORDER_INFO__ORDER_DATE2D=&G1_ORDER_INFO__ORDER_NUMSS=&G1_ORDER_INFO__STATE_ID=&G1_ORDER_INFO__RECIP_ID=&Submit=%CD%E0%E9%F2%E8"
const postfix string = "/modules.php?name=sud_delo&srv_num=1&name_op=r&page=MYPAGE&vnkod=50RS0001&srv_num=1&name_op=r&delo_id=1540005&case_type=0&new=0&G1_PARTS__NAMESS=&g1_case__CASE_NUMBERSS=&g1_case__JUDICIAL_UIDSS=&delo_table=g1_case&g1_case__ENTRY_DATE1D=DATEOT&g1_case__ENTRY_DATE2D=DATEDO&lawbookarticles%5B0%5D=%D1%EF%EE%F0%FB%2C+%E2%EE%E7%ED%E8%EA%E0%FE%F9%E8%E5+%E8%E7+%F2%F0%F3%E4%EE%E2%FB%F5+%EE%F2%ED%EE%F8%E5%ED%E8%E9&G1_CASE__JUDGE=&g1_case__RESULT_DATE1D=&g1_case__RESULT_DATE2D=&G1_CASE__RESULT=&G1_CASE__BUILDING_ID=&G1_CASE__COURT_STRUCT=&G1_EVENT__EVENT_NAME=&G1_EVENT__EVENT_DATEDD=&G1_PARTS__PARTS_TYPE=&G1_PARTS__INN_STRSS=&G1_PARTS__KPP_STRSS=&G1_PARTS__OGRN_STRSS=&G1_PARTS__OGRNIP_STRSS=&G1_RKN_ACCESS_RESTRICTION__RKN_REASON=&g1_rkn_access_restriction__RKN_RESTRICT_URLSS=&g1_requirement__ACCESSION_DATE1D=&g1_requirement__ACCESSION_DATE2D=&G1_REQUIREMENT__CATEGORY=&g1_requirement__ESSENCESS=&g1_requirement__JOIN_END_DATE1D=&g1_requirement__JOIN_END_DATE2D=&G1_REQUIREMENT__PUBLICATION_ID=&G1_DOCUMENT__PUBL_DATE1D=&G1_DOCUMENT__PUBL_DATE2D=&G1_CASE__VALIDITY_DATE1D=&G1_CASE__VALIDITY_DATE2D=&G1_ORDER_INFO__ORDER_DATE1D=&G1_ORDER_INFO__ORDER_DATE2D=&G1_ORDER_INFO__ORDER_NUMSS=&G1_ORDER_INFO__STATE_ID=&G1_ORDER_INFO__RECIP_ID=&Submit=%CD%E0%E9%F2%E8"

//const asdasds string = "/modules.php?name=sud_delo&srv_num=1&name_op=r&page=2&     vnkod=50RS0001&srv_num=1&name_op=r&delo_id=1540005&case_type=0&new=0&G1_PARTS__NAMESS=&g1_case__CASE_NUMBERSS=&g1_case__JUDICIAL_UIDSS=&delo_table=g1_case&g1_case__ENTRY_DATE1D=19.09.2021&g1_case__ENTRY_DATE2D=25.09.2022&lawbookarticles%5B0%5D=%D1%EF%EE%F0%FB%2C+%E2%EE%E7%ED%E8%EA%E0%FE%F9%E8%E5+%E8%E7+%F2%F0%F3%E4%EE%E2%FB%F5+%EE%F2%ED%EE%F8%E5%ED%E8%E9&G1_CASE__JUDGE=&g1_case__RESULT_DATE1D=&g1_case__RESULT_DATE2D=&G1_CASE__RESULT=&G1_CASE__BUILDING_ID=&G1_CASE__COURT_STRUCT=&G1_EVENT__EVENT_NAME=&G1_EVENT__EVENT_DATEDD=&G1_PARTS__PARTS_TYPE=&G1_PARTS__INN_STRSS=&G1_PARTS__KPP_STRSS=&G1_PARTS__OGRN_STRSS=&G1_PARTS__OGRNIP_STRSS=&G1_RKN_ACCESS_RESTRICTION__RKN_REASON=&g1_rkn_access_restriction__RKN_RESTRICT_URLSS=&g1_requirement__ACCESSION_DATE1D=&g1_requirement__ACCESSION_DATE2D=&G1_REQUIREMENT__CATEGORY=&g1_requirement__ESSENCESS=&g1_requirement__JOIN_END_DATE1D=&g1_requirement__JOIN_END_DATE2D=&G1_REQUIREMENT__PUBLICATION_ID=&G1_DOCUMENT__PUBL_DATE1D=&G1_DOCUMENT__PUBL_DATE2D=&G1_CASE__VALIDITY_DATE1D=&G1_CASE__VALIDITY_DATE2D=&G1_ORDER_INFO__ORDER_DATE1D=&G1_ORDER_INFO__ORDER_DATE2D=&G1_ORDER_INFO__ORDER_NUMSS=&G1_ORDER_INFO__STATE_ID=&G1_ORDER_INFO__RECIP_ID=&Submit=%CD%E0%E9%F2%E8"

func main() {
	dt := time.Now()
	datedo := dt.Format("02.01.2006")
	dateot := dt.AddDate(0, -1, 0).Format("02.01.2006")
	//dateot = "15.11.2022"
	outputfilename := "Суды МО от " + dateot + " до " + datedo + ".xlsx"
	//fmt.Println(dateot, datedo)
	//var couters couterMO_data
	var index_link int
	var tecal_couter couterMO_data
	var tecal_link string
	var tecal_page int // Текущий лист
	var all_page int   // Всео листов
	var err_page error
	f, _ := excelize.OpenFile("links.xlsx")
	rows, _ := f.GetRows("main")
	f_excel := excelize.NewFile()
	f_excel.NewSheet("main")
	f_excel.DeleteSheet("Sheet1")
	c := colly.NewCollector()

	make_Title(f_excel)

	c.OnHTML("td[valign=bottom]", func(element *colly.HTMLElement) {
		if element.DOM.Find("a:last-child ").Text() == ">>" {
			attr_link, _ := element.DOM.Find("a:last-child ").Attr("href")
			attr_link = rows[index_link][0] + attr_link[1:]
			u, _ := url.Parse(attr_link)
			m, _ := url.ParseQuery(u.RawQuery)              //parse query param into map
			all_page, err_page = strconv.Atoi(m["page"][0]) //prints specific key value
			if err_page != nil {
				all_page = 1
			}
		}
	})

	c.OnHTML("#tablcont > tbody > tr", func(element *colly.HTMLElement) {
		tecal_couter.number = TrimAll(element.DOM.Find("td:nth-child(1)").Text())
		tecal_couter.datesP = TrimAll(element.DOM.Find("td:nth-child(2)").Text())
		tecal_couter.kategory, tecal_couter.applicant, tecal_couter.defendant = kategory_applicant_defendant(TrimAll(element.DOM.Find("td:nth-child(3)").Text()))
		tecal_couter.judge = TrimAll(element.DOM.Find("td:nth-child(4)").Text())
		tecal_couter.datesR = TrimAll(element.DOM.Find("td:nth-child(5)").Text())
		tecal_couter.decision = TrimAll(element.DOM.Find("td:nth-child(6)").Text())
		tecal_couter.datesZ = TrimAll(element.DOM.Find("td:nth-child(7)").Text())
		tecal_couter.akts = TrimAll(element.DOM.Find("td:nth-child(8)").Text())

		tecal_couter.url, _ = (element.DOM.Find("td:nth-child(1) > a[href]").Attr("href"))
		tecal_couter.url = tecal_link + tecal_couter.url
		if tecal_couter.number != "" {
			saveTypeOnXLSX(f_excel, tecal_couter)
			//couters = append(couters, tecal_couter)
		}
	})
	//link, _ := f.GetCellValue("main", "A"+strconv.Itoa(53)) // !

	for index_link := range rows {
		//link, _ := f.GetCellValue("main", "A"+strconv.Itoa(index_link+1))
		all_page = 1
		for tecal_page = 1; tecal_page <= all_page; tecal_page++ {
			link := rows[index_link][0] + postfix
			tecal_link = linkURLcouter(link) + postfix
			link = strings.Replace(link, "DATEOT", dateot, 1)
			link = strings.Replace(link, "DATEDO", datedo, 1)
			link = strings.Replace(link, "MYPAGE", strconv.Itoa(tecal_page), 1)

			c.Visit(link)
		}
	}
	//saveTypeOnXLSX(f_excel, couters)
	if err := f_excel.SaveAs(outputfilename); err != nil {
		fmt.Println(err)
	}
}

// Получить ячейку
func cells(y, x int) string {
	strCell, _ := excelize.CoordinatesToCellName(x, y)
	return strCell
}

// Создать шапку колонок
func make_Title(f *excelize.File) {
	f.SetCellValue("main", cells(1, 1), "Номер дела")
	f.SetCellValue("main", cells(1, 2), "Ссылка на дело")
	f.SetCellValue("main", cells(1, 3), "Дата поступления")
	f.SetCellValue("main", cells(1, 4), "Категория")
	f.SetCellValue("main", cells(1, 5), "Истец")
	f.SetCellValue("main", cells(1, 6), "Судья Арбитражного суда")
	f.SetCellValue("main", cells(1, 7), "Суд (наименование суда)")
	f.SetCellValue("main", cells(1, 8), "Название компании")
	f.SetCellValue("main", cells(1, 9), "ИНН")
	f.SetCellValue("main", cells(1, 10), "link")
	f.SetCellValue("main", cells(1, 11), "Судебные акты")
	f.SetCellValue("main", cells(1, 12), "Дата решения")
	f.SetCellValue("main", cells(1, 13), "Дата вступления в законную силу")
	f.SetCellValue("main", cells(1, 14), "Решение")
	f.SetCellValue("main", cells(1, 15), "Источник")
	f.SetCellValue("main", cells(1, 16), "Доступен")
}
func saveTypeOnXLSX(f *excelize.File, cou couterMO_data) {
	//for index := range cou {
	f.SetCellValue("main", cells(index_global, 1), cou.number)
	f.SetCellValue("main", cells(index_global, 2), cou.url)
	f.SetCellValue("main", cells(index_global, 3), cou.datesP)
	f.SetCellValue("main", cells(index_global, 4), cou.kategory)
	f.SetCellValue("main", cells(index_global, 5), cou.applicant)
	f.SetCellValue("main", cells(index_global, 6), cou.judge)
	f.SetCellValue("main", cells(index_global, 7), "")
	f.SetCellValue("main", cells(index_global, 8), cou.defendant)
	f.SetCellValue("main", cells(index_global, 9), "")
	f.SetCellValue("main", cells(index_global, 10), "ИНН Московская Область "+strings.ReplaceAll(cou.defendant, "\"", ""))

	f.SetCellValue("main", cells(index_global, 11), cou.akts)     // Судебные акты
	f.SetCellValue("main", cells(index_global, 12), cou.datesR)   // Дата решения
	f.SetCellValue("main", cells(index_global, 13), cou.datesZ)   // Дата вступления в законную силу
	f.SetCellValue("main", cells(index_global, 14), cou.decision) // Решение

	f.SetCellValue("main", cells(index_global, 15), "Мосгорсуд")
	f.SetCellValue("main", cells(index_global, 16), "Да")
	index_global++
	//}
}
func TrimAll(str string) string {
	return strings.TrimSpace(str)
}
func linkURLcouter(str string) string {
	return str[:strings.Index(str, "modules")-1]
}
func kategory_applicant_defendant(str string) (string, string, string) {
	str = strings.Replace(str, "КАТЕГОРИЯ:", "", 1)
	str = strings.Replace(str, "ИСТЕЦ(ЗАЯВИТЕЛЬ):", "<>", 1)
	str = strings.Replace(str, "ОТВЕТЧИК:", "<>", 1)
	sas := strings.Split(str, "<>")
	if len(sas) == 3 {
		return TrimAll(sas[0]), TrimAll(sas[1]), TrimAll(sas[2])
	} else {
		return "", "", ""
	}
}
