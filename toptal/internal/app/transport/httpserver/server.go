package httpserver

type HttpServer struct {
	userService     UserService
	tokenService    TokenService
	bookService     BookService
	categoryService CategoryService
	cartService     CartService
}

func NewHttpServer(userService UserService, tokenService TokenService, bookService BookService,
	categoryService CategoryService, cartService CartService) HttpServer {
	return HttpServer{
		userService:     userService,
		tokenService:    tokenService,
		bookService:     bookService,
		categoryService: categoryService,
		cartService:     cartService,
	}
}
