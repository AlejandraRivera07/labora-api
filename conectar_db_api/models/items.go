package models

// Estructura para representar un ítem
type Item struct {
	ID   int    `json:"id"`   // Identificador del ítem
	Name string `json:"name"` // Nombre del ítem

	// Nuevos campos
	CustomerName string `json:"customerName"` // Nombre del cliente que hizo el pedido
	OrderDate    string `json:"orderDate"`    // Fecha en que se hizo el pedido
	Product      string `json:"product"`      // Producto del pedido
	Quantity     int    `json:"quantity"`     // Cantidad del producto en el pedido
	Price        int    `json:"price"`        // Precio del producto en el pedido
	Views        int     `json:"views"`        // Cantidad de veces que se ve in item
	TotalPrice   float64 `json:"totalPrice"`
}

// calculo del precio total de un item
func (item *Item) CalculateTotalPrice() float64 {
	TotalPrice := float64(item.Price) * float64(item.Quantity)

	item.TotalPrice = math.Round(TotalPrice*100) / 100
	return item.TotalPrice
}

