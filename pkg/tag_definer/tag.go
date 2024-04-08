package tag_definer

import "strings"

type Tag string

var (
	Transport          = Tag("Transport")
	BarsAndRestaurants = Tag("Food & Drinks")
	Food               = Tag("Food & Drinks")
	Pharmacy           = Tag("Health & Personal Care")
	Market             = Tag("Mercado")
	Pet                = Tag("Pet")
	OnlineShopping     = Tag("Compras Online")
)

var categoryKeyWords = map[Tag][]string{
	Transport:          {"posto", "autopost", "conectcar", "estacion", "meuestar", "meu estacionamento", "uber", "99", "cabify"},
	BarsAndRestaurants: {"boteco", "bar", "restaurante", "churrascaria", "pizzaria", "padaria", "padoca", "padoka", "padok"},
	Food:               {"ifood", "rappi", "mcdonalds", "burger king", "subway", "kfc", "pizza hut", "dominos", "domino's", "ifd*"},
	Pharmacy:           {"panvel", "raia", "drogasil", "droga", "farmacia", "farmácia"},
	Market:             {"festval", "super beal", "mercado", "supermercado", "market4u", "mart minas"},
	Pet:                {"cobasi", "petz"},
	OnlineShopping:     {"shopee", "amazon", "aliexpress", "ebay", "wish", "magalu", "magazine luiza", "submarino", "americanas", "ponto frio", "casas bahia", "netshoes", "centauro", "zattini", "dafiti", "kanui", "renner", "cea", "marisa", "riachuelo", "havan", "havan.com", "havan.com.br", "havan.com"},
}

func GetTagFromDescription(description string) Tag {
	for tag, keys := range categoryKeyWords {
		for i := 0; i < len(keys); i++ {
			if strings.Contains(strings.ToLower(description), strings.ToLower(keys[i])) {
				return tag
			}
		}
	}

	return ""
}
