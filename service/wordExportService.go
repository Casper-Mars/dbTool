package service

import (
	"github.com/Casper-Mars/dbTool/pojo"
	"github.com/Casper-Mars/officeTool/color"
	"github.com/Casper-Mars/officeTool/document"
	"github.com/Casper-Mars/officeTool/measurement"
	"github.com/Casper-Mars/officeTool/schema/soo/wml"
	"log"
	"strconv"
)

func Export(tables []pojo.TableInfo, dbName string, storeLocation string) {

	doc := document.New()
	defer doc.Close()

	for _, k := range tables {
		createOneTable(k, doc)
	}

	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	if storeLocation == "" {
		storeLocation = "./"
	}
	doc.SaveToFile(storeLocation + dbName + ".docx")
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
