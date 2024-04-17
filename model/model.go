package model

// 看postman
// sign in的地方，返回回来的username role token。用token看对应是哪个用户
type SigninResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

// 在sign up前端不需要id，只需要username用户名和password密码选个role权限
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// 只subject主题和content内容需要传
type Discussion struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	LastUpdateTime string `json:"last_update_time"`
}

// postreply。只传discussion_id内容的id和 content内容本身
type Reply struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	DiscussionId int    `json:"discussion_id"`
	Content      string `json:"content"`
	ReplyTime    string `json:"reply_time"`
}

// discussiondetail。返回discussion（本身内容），和所有reply（回复）
type DiscussionDto struct {
	Discussion Discussion `json:"discussion"`
	Replies    []Reply    `json:"replies"`
}

// 维修。只需要传subject主题和content内容
type Maintenance struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	Reply          string `json:"reply"`
	Completed      bool   `json:"completed"`
	LastUpdateTime string `json:"last_update_time"`
}

// 账单。postbill。传username用户名 和maintenance_id，确定给谁。item是物品留言，amount多少钱
type Bill struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	MaintenanceId int    `json:"maintenance_id"`
	Item          string `json:"item"`
	Amount        int    `json:"amount"`
	BillTime      string `json:"bill_time"`
}

// （用户用的）支付：id，用户名，名称，总金额，时间。前端只传item和amount
type Payment struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Item            string `json:"item"`
	ApartmentNumber string `json:"apartment_number"`
	Amount          int    `json:"amount"`
	PaymentTime     string `json:"payment_time"`
}

// 传列表，所有的bills和所有的payment支付，balance=bills-payments
type BalanceDto struct {
	Balance  int       `json:"balance"`
	Bills    []Bill    `json:"bills"`
	Payments []Payment `json:"payments"`
}

type Facility struct {
	Id           int    `json:"id"`
	FacilityName string `json:"facility_name"`
	Description  string `json:"description"`
}

type Reservation struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	FacilityName    string `json:"facility_name"`
	Remark          string `json:"remark"`
	ReservationDate string `json:"reservation_date"`
	StartHour       int    `json:"start_hour"`
	EndHour         int    `json:"end_hour"`
}

type ReservationRequest struct {
	FacilityName string `json:"facility_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
