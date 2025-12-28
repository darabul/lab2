package payments

import "errors"

// Order представляет заказ в системе оплаты.
type Order struct {
	Amount int    // Сумма заказа в минимальных единицах валюты
	Status string // Текущий статус заказа
}

// NewOrder создает и возвращает новый заказ с указанной суммой.
// Начальный статус заказа устанавливается в "created".
func NewOrder(amount int) *Order {
	return &Order{Amount: amount, Status: "created"}
}

// IsValid проверяет корректность данных заказа.
// Возвращает true, если сумма заказа неотрицательна.
func (o *Order) IsValid() bool {
	return o.Amount >= 0
}

// PaymentService предоставляет методы для работы с оплатами заказов.
type PaymentService struct{}

// NewPaymentService создает и возвращает новый экземпляр PaymentService.
func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

// ProcessPayment обрабатывает оплату заказа.
// Проверяет статус и валидность заказа, обновляет статус в соответствии с результатом.
// Возвращает ошибку, если:
// - заказ уже оплачен (status == "paid")
// - заказ невалиден
// В случае успешной обработки возвращает nil.
func (ps *PaymentService) ProcessPayment(order *Order) error {
	if order.Status == "paid" {
		return errors.New("order is already paid")
	}

	if !order.IsValid() {
		order.Status = "payment_failed"
		return errors.New("error")
	}

	if order.Amount == 0 {
		order.Status = "confirmed"
		return nil
	}

	order.Status = "paid"
	return nil
}
