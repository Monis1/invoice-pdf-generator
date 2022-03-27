package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	//DE
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	//Header
	pdf.SetFont("Arial", "B", 25)
	pdf.Cellf(0, 0, "Receipt")
	pdf.Ln(24)
	pdf.SetFont("Arial", "B", 10)
	pdf.Cellf(0, 0, "ABC Inc.")
	pdf.SetFont("Arial", "", 10)
	pdf.Ln(8)
	pdf.Cellf(0, 0, "ABC Street, 3rd Floor")
	pdf.Ln(4)
	pdf.Cellf(0, 0, "123456 karachi")
	pdf.Ln(4)
	pdf.Cellf(0, 0, "Pakistan")

	//body
	pdf.SetFont("Arial", "", 10)
	pdf.Ln(16)
	pdf.Cellf(0, 0, "Date: 27 March 2022")
	pdf.Ln(4)
	pdf.Cellf(0, 0, "Order #: MS524637848488")
	pdf.Ln(16)
	pdf.Cellf(0, 0, "Sold To: Monis Ahmed")
	pdf.Ln(16)

	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		pdf.SetFont("Arial", "B", 10)

		pdf.CellFormat(20, 7, "Quantity", "1", 0, "", false, 0, "")
		pdf.CellFormat(80, 7, "Item", "1", 0, "", false, 0, "")
		pdf.CellFormat(35, 7, "Price", "1", 0, "", false, 0, "")
		pdf.CellFormat(35, 7, "Total", "1", 0, "", false, 0, "")

		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 10)
		pdf.SetX(left)
		pdf.CellFormat(20, 6, "1", "1", 0, "", false, 0, "")
		pdf.CellFormat(80, 6, "ABC Product", "1", 0, "", false, 0, "")
		pdf.CellFormat(35, 6, "$5000", "1", 0, "R", false, 0, "")
		pdf.CellFormat(35, 6, "$5000", "1", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}
	basicTable()

	pdf.Text(125, 120, "Subtotal")
	pdf.Text(182, 120, "$5000")
	pdf.Ln(4)
	pdf.Text(125, 124, "Tax")
	pdf.Text(182, 124, "$0")
	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.Text(125, 128, "Total")
	pdf.Text(182, 128, "$5000")
	pdf.SetFont("Arial", "", 10)

	//footer
	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(-50)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(0, 10, "Page	1 of 1",
			"", 0, "C", false, 0, "")
		pdf.Ln(16)
		pdf.SetFont("Arial", "B", 9)
		pdf.Cellf(0, 10, "Contact:")
		pdf.Text(160, 270, "Company Text Title")
		pdf.Ln(6)
		pdf.SetFont("Arial", "", 9)
		pdf.Text(160, 274, "Some Text")

		_, lineHt := pdf.GetFontSize()
		htmlStr := `<a href="mailto:abc.com">abc.com</a>`
		html := pdf.HTMLBasicNew()
		html.Write(lineHt, htmlStr)

	})

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		fmt.Println(err)
	}

}
