package model

import (
	"github.com/Rhymond/go-money"
	"github.com/viniciusgabrielfo/organizze-invoice-itau-converter/pkg/category_definer"
	"github.com/viniciusgabrielfo/organizze-invoice-itau-converter/pkg/tag_definer"
)

type Entry struct {
	Date        string
	Description string
	Category    category_definer.Category
	Tag         tag_definer.Tag
	Value       *money.Money
}
