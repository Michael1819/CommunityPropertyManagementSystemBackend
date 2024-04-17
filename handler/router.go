package handler

import (
	"github.com/form3tech-oss/jwt-go"
	"net/http"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var signingKey []byte

func InitRouter() http.Handler {
	signingKey = []byte("secret")

	middleware := jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	router := mux.NewRouter()

	router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
	router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")

	router.Handle("/time", middleware.Handler(http.HandlerFunc(timeHandler))).Methods("GET")

	//用户传标题内容
	router.Handle("/discussion", middleware.Handler(http.HandlerFunc(postDiscussionHandler))).Methods("POST")
	//论坛入口，显示列表，所有discussion
	router.Handle("/alldiscussions", middleware.Handler(http.HandlerFunc(getAllDiscussionsHandler))).Methods("GET")
	//看自己发了那些
	router.Handle("/mydiscussions", middleware.Handler(http.HandlerFunc(getMyDiscussionsHandler))).Methods("GET")
	//discussion本身所有的内容和所有的回复
	router.Handle("/discussion", middleware.Handler(http.HandlerFunc(getDiscussionDetailHandler))).Methods("PUT")
	//选定一条，删除
	router.Handle("/discussion", middleware.Handler(http.HandlerFunc(deleteDiscussionHandler))).Methods("DELETE")

	//选定一个，然后回复
	router.Handle("/reply", middleware.Handler(http.HandlerFunc(postReplyHandler))).Methods("POST")
	//看自己所有
	router.Handle("/reply", middleware.Handler(http.HandlerFunc(getMyRepliesHandler))).Methods("GET")
	router.Handle("/reply", middleware.Handler(http.HandlerFunc(deleteReplyHandler))).Methods("DELETE")

	//第三方使用，维修
	router.Handle("/maintenance", middleware.Handler(http.HandlerFunc(postMaintenanceHandler))).Methods("POST")
	//第三方使用，看所有的维修
	router.Handle("/allmaintenances", middleware.Handler(http.HandlerFunc(getAllMaintenancesHandler))).Methods("POST")
	//只看到自己的
	router.Handle("/mymaintenances", middleware.Handler(http.HandlerFunc(getMyMaintenancesHandler))).Methods("POST")
	//给第三方使用，
	router.Handle("/maintenance", middleware.Handler(http.HandlerFunc(putMaintenanceHandler))).Methods("PUT")

	//给第三方之用的bill账单，选定某个maintainnance order维修单，然后给对方发bill账单
	router.Handle("/bill", middleware.Handler(http.HandlerFunc(postBillHandler))).Methods("POST")
	//用户使用的账单
	router.Handle("/bill", middleware.Handler(http.HandlerFunc(getMyBillsHandler))).Methods("GET")

	//用户使用的支付
	router.Handle("/payment", middleware.Handler(http.HandlerFunc(postPaymentHandler))).Methods("POST")
	router.Handle("/payment", middleware.Handler(http.HandlerFunc(getMyPaymentsHandler))).Methods("GET")

	//把两列表合起来最后的综合，账单和支付的差值
	router.Handle("/balance", middleware.Handler(http.HandlerFunc(getMyBalanceHandlder))).Methods("GET")

	//应该是对应calendar板块
	//预约的公共设施的预约单。提交/查看/删除
	router.Handle("/reservation", middleware.Handler(http.HandlerFunc(postReservationHandler))).Methods("POST")
	router.Handle("/myreservations", middleware.Handler(http.HandlerFunc(getMyReservationsHandler))).Methods("GET")
	router.Handle("/reservation", middleware.Handler(http.HandlerFunc(deleteReservationHandler))).Methods("PUT")

	//公共设施。得到所有公共设施的数据
	router.Handle("/facility", middleware.Handler(http.HandlerFunc(getAllFacilitiesHandler))).Methods("GET")
	//得到已经被预约的所有公共设施的数据
	router.Handle("/facilityreservations", middleware.Handler(http.HandlerFunc(getFacilityReservationsHandler))).Methods("GET")

	//得到日历
	router.Handle("/calendar", middleware.Handler(http.HandlerFunc(getCalendarHandler))).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
