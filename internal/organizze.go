package internal

import (
	"fmt"

	"github.com/viniciusgabrielfo/organizze-invoice-itau-converter/pkg/model"
	"github.com/xuri/excelize/v2"
)

const OrganizzeSheetName = "organizze-entries.xlsx"

func GenerateOrganizzeXLXSSheet(entries []model.Entry) error {
	f := excelize.NewFile()

	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}

		return nil
	}()

	_ = f.SetCellValue("Sheet1", "A1", "date")
	_ = f.SetCellValue("Sheet1", "B1", "description")
	_ = f.SetCellValue("Sheet1", "C1", "tag")
	_ = f.SetCellValue("Sheet1", "D1", "category")
	_ = f.SetCellValue("Sheet1", "E1", "amount")

	for i := 0; i < len(entries); i++ {
		row := i + 2

		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), entries[i].Date)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), entries[i].Description)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), entries[i].Tag)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), entries[i].Category)
		_ = f.SetCellFloat("Sheet1", fmt.Sprintf("E%d", row), entries[i].Value.AsMajorUnits(), 2, 32)
	}

	if err := f.SaveAs(OrganizzeSheetName); err != nil {
		return err
	}

	return nil
}
