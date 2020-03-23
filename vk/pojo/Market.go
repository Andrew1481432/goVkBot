package pojo

type Price struct {
	Amount int `json:"amount" map:"amount"`
}

type Market struct {
	ID          int    `json:"id" map:"id"`
	OwnerID     int    `json:"owner_id" map:"owner_id"`
	Title       string `json:"title" map:"title"`
	Description string `json:"description" map:"description"`
	//	Price *Price        `json:"price" map:"price"`
}
