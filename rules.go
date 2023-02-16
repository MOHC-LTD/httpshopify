package httpshopify

import "github.com/MOHC-LTD/shopify/v2"

// RuleDTOs represents Shopify rules in HTTP requests and responses.
type RuleDTOs []RuleDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos RuleDTOs) ToShopify() shopify.Rules {
	rules := make(shopify.Rules, 0, len(dtos))

	for _, dto := range dtos {
		rules = append(rules, dto.ToShopify())
	}

	return rules
}

// RuleDTO represents Shopify rules in HTTP requests and responses.
type RuleDTO struct {
	Column    string `json:"column"`
	Relation  string `json:"relation"`
	Condition string `json:"condition"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto RuleDTO) ToShopify() shopify.Rule {
	return shopify.Rule{
		Column:    dto.Column,
		Relation:  dto.Relation,
		Condition: dto.Condition,
	}
}
