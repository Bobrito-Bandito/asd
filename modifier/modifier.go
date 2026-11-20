package modifiers

import (
	"financialExchange/enum"
)

func ModifierForOrder(sellerQuantity float64) func(*enum.Order) *enum.Order {
	return func(order *enum.Order) *enum.Order {
		order.Filled += sellerQuantity
//функция обертка над анонимной функцией которая работает с ордерами
		if order.Filled >= order.Quantity {
			order.Status = enum.OrderStatusFilled
		} else {
			order.Status = enum.OrderStatusPartiallyFilled
		}

		return order
	}
}
