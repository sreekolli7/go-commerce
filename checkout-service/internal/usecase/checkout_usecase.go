package usecase

import "github.com/sreekolli7/checkout-service/internal/domain"

// I'm defining what my repository needs to do
// This is called an interface - it's like a contract
type OrderRepository interface {
	Create(order *domain.Order) error
	GetOrdersByUserID(userID int64) ([]domain.Order, error)
}

type CheckoutUsecase struct{ Repo OrderRepository }

func NewCheckoutUsecase(r OrderRepository) *CheckoutUsecase {
	return &CheckoutUsecase{Repo: r}
}

func (uc *CheckoutUsecase) Checkout(req *domain.CheckoutRequest) (*domain.Order, error) {
	// I'm calculating the total price of all items
	// This is my business logic for processing orders
	total := 0.0
	for _, it := range req.Items {
		total += float64(it.Quantity) * it.Price
	}

	// I'm creating a new order with all the details
	// This is the main business logic for checkout
	order := &domain.Order{
		UserID: req.UserID,
		Items:  req.Items,
		Total:  total,
		Status: "pending",
	}

	// I'm saving the order to the database
	// This is where the order gets stored permanently
	if err := uc.Repo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (uc *CheckoutUsecase) GetOrdersByUserID(userID int64) ([]domain.Order, error) {
	// I'm getting all orders for a specific user
	// This is my business logic for retrieving order history
	return uc.Repo.GetOrdersByUserID(userID)
}
