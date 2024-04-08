package category_definer

import "strings"

type Category string

var (
	Transport          = Category("Transport")
	BarsAndRestaurants = Category("Food & Drinks")
	Food               = Category("Food & Drinks")
	Pharmacy           = Category("Health & Personal Care")
	Market             = Category("Mercado")
	Pet                = Category("Pet")
	OnlineShopping     = Category("Compras Online")
)

var categoryKeyWords = map[Category][]string{
	Transport:          {"posto", "autopost", "conectcar", "estacion", "meuestar", "meu estacionamento", "uber", "99", "cabify"},
	BarsAndRestaurants: {"boteco", "bar", "restaurante", "churrascaria", "pizzaria", "padaria", "padoca", "padoka", "padok"},
	Food:               {"ifood", "rappi", "mcdonalds", "burger king", "subway", "kfc", "pizza hut", "dominos", "domino's", "ifd*"},
	Pharmacy:           {"panvel", "raia", "drogasil", "droga", "farmacia", "farmácia"},
	Market:             {"festval", "super beal", "mercado", "supermercado", "market4u", "mart minas"},
	Pet:                {"cobasi", "petz"},
	OnlineShopping:     {"shopee", "amazon", "aliexpress", "ebay", "wish", "magalu", "magazine luiza", "submarino", "americanas", "ponto frio", "casas bahia", "netshoes", "centauro", "zattini", "dafiti", "kanui", "renner", "cea", "marisa", "riachuelo", "havan", "havan.com", "havan.com.br", "havan.com"},
}

func GetCategoryFromDescription(description string) Category {
	for category, keys := range categoryKeyWords {
		for i := 0; i < len(keys); i++ {
			if strings.Contains(strings.ToLower(description), strings.ToLower(keys[i])) {
				return category
			}
		}
	}

	return ""
}
