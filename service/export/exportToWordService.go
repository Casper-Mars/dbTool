package export

import (
	"github.com/Casper-Mars/dbTool/pojo"
	"github.com/Casper-Mars/officeTool/color"
	"github.com/Casper-Mars/officeTool/document"
	"github.com/Casper-Mars/officeTool/measurement"
	"github.com/Casper-Mars/officeTool/schema/soo/wml"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type WordExportService struct {
}

func (service WordExportService) Export(dbName string, storeLocation string, tableInfos []pojo.TableInfo) {
	if len(dbName) == 0 {
		log.Printf("数据库名称为空")
		dbName = strconv.Itoa(rand.Int())
	}
	if len(storeLocation) == 0 {
		storeLocation = "." + string(os.PathSeparator)
	} else if !strings.HasSuffix(storeLocation, string(os.PathSeparator)) {
		storeLocation = storeLocation + string(os.PathSeparator)
	}
	export(tableInfos, dbName, storeLocation)
}

func export(tables []pojo.TableInfo, dbName string, storeLocation string) {

	doc := document.New()
	defer doc.Close()

	for _, k := range tables {
		createOneTable(k, doc)
	}

	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	err := doc.SaveToFile(storeLocation + dbName + ".docx")
	if err != nil {
		log.Println("导出" + dbName + "失败")
	}
}

func createOneTable(table pojo.TableInfo, doc *document.Document) {
	doc.AddParagraph().AddRun().AddText(table.TableName + "(" + table.Comment + ")")

	// First Table
	wordTable := doc.AddTable()
	// width of the page
	wordTable.Properties().SetWidthPercent(100)
	borders := wordTable.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)
	createTableHeader(&wordTable)
	for _, k := range table.Cols {
		row := wordTable.AddRow()
		getNameColCell(&row).AddParagraph().AddRun().AddText(k.ColName)
		getStandardCell(&row).AddParagraph().AddRun().AddText(k.ColType)
		if k.Len == 0 {
			getStandardCell(&row).AddParagraph().AddRun().AddText("-")
		} else {
			getStandardCell(&row).AddParagraph().AddRun().AddText(strconv.Itoa(k.Len))
		}
		if k.IsEmpty {
			getStandardCell(&row).AddParagraph().AddRun().AddText("是")
		} else {
			getStandardCell(&row).AddParagraph().AddRun().AddText("否")
		}
		if k.IsPrimary {
			getStandardCell(&row).AddParagraph().AddRun().AddText("是")
		} else {
			getStandardCell(&row).AddParagraph().AddRun().AddText("否")
		}
		getStandardCell(&row).AddParagraph().AddRun().AddText(k.Comment)
	}

}
func createTableHeader(wordTable *document.Table) {
	row := wordTable.AddRow()
	getNameColCell(&row).AddParagraph().AddRun().AddText("名称")
	getStandardCell(&row).AddParagraph().AddRun().AddText("类型")
	getStandardCell(&row).AddParagraph().AddRun().AddText("长度")
	getStandardCell(&row).AddParagraph().AddRun().AddText("可空")
	getStandardCell(&row).AddParagraph().AddRun().AddText("主键")
	getStandardCell(&row).AddParagraph().AddRun().AddText("描述")
}
func getStandardCell(row *document.Row) document.Cell {
	cell := row.AddCell()
	//cell.Properties().SetWidthPercent(0.16)
	return cell
}
func getNameColCell(row *document.Row) document.Cell {
	cell := row.AddCell()
	cell.Properties().SetWidth(1)
	return cell
}
