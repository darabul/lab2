package payments

import "testing"

func TestNewOrder(t *testing.T) {
	o := NewOrder(100)
	if o.Amount != 100 {
		t.Errorf("expected 100, got %d", o.Amount)
	}
}

func TestNewOrder2(t *testing.T) {
	o := NewOrder(200)
	if o.Amount != 200 {
		t.Errorf("expected 100, got %d", o.Amount)
	}
}

func TestOrderInvalidAmount(t *testing.T) {
	o := NewOrder(-10)
	if o.IsValid() {
		t.Errorf("order with negative amount should be invalid")
	}
}

func TestOrderZeroAmount(t *testing.T) {
	o := NewOrder(0)
	if !o.IsValid() {
		t.Errorf("order with zero amount should be valid")
	}
}

func TestOrderPositiveAmount(t *testing.T) {
	o := NewOrder(100)
	if !o.IsValid() {
		t.Errorf("order with positive amount should be valid")
	}
}

func TestProcessPaymentInalidOrder(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(-10)

	err := ps.ProcessPayment(o)
	if err == nil {
		t.Errorf("expected error for invalid order")
	}
}

func TestProcessPaymentValidOrder(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(0)

	err := ps.ProcessPayment(o)
	if err != nil {
		t.Errorf("expected no error for valid order, got: %v", err)
	}
}

func TestOrderStatus(t *testing.T) {
	o := NewOrder(100)
	if o.Status != "created" {
		t.Errorf("expected status 'created', got %s", o.Status)
	}
}

func TestProcessPaymentChangesStatus(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(100)

	err := ps.ProcessPayment(o)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if o.Status != "paid" {
		t.Errorf("expected status 'paid', got %s", o.Status)
	}
}

func TestProcessPaymentWithErrorChangesStatus(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(-10)

	err := ps.ProcessPayment(o)
	if err == nil {
		t.Errorf("exprected error")
	}

	if o.Status != "payment_failed" {
		t.Errorf("expected status 'payment_failed', got %s", o.Status)
	}
}

func TestProcessPaymentAlreadyPaidOrder(t *testing.T) {
	ps := NewPaymentService()
	order := NewOrder(200)

	err := ps.ProcessPayment(order)
	if err != nil {
		t.Fatalf("unexpected error processing first payment: %v", err)
	}

	err = ps.ProcessPayment(order)
	if err == nil {
		t.Errorf("expected error when processing payment for already paid order")
	}

	if order.Status != "paid" {
		t.Errorf("expected status to remain 'paid', got %s", order.Status)
	}
}

func TestProcessPaymentZeroAmountOrder(t *testing.T) {
	ps := NewPaymentService()
	order := NewOrder(0)

	err := ps.ProcessPayment(order)
	if err != nil {
		t.Fatalf("unexpected error processing first payment: %v", err)
	}

	if order.Status != "confirmed" {
		t.Errorf("expected status to remain 'paid', got %s", order.Status)
	}
}
