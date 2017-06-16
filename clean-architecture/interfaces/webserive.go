package interfaces

import (
	"fmt"
	"github.com/linh-snoopy/go-examples/clean-architecture/usecases"
	"io"
	"net/http"
	"strconv"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	Add(userId, orderId, itemId int) error
}

type WebServiceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler WebServiceHandler) ShowOrder(res http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.FormValue("userId"))
	orderId, _ := strconv.Atoi(r.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}
