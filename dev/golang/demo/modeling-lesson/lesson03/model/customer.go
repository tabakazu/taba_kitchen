package model

// 会員顧客
type Customer struct {
	ID        int        // ID
	Name      string     // 名称
	Email     string     // メールアドレス
	Addresses []*Address // 配送先リスト
}

func NewCustomer(id int, name, email string) *Customer {
	return &Customer{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
