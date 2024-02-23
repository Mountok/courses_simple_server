package app

import (
	"context"
	"courses_serve/example/api"
	"courses_serve/example/api/middleware"
	"courses_serve/example/internals/app/db"
	"courses_serve/example/internals/app/handlers"
	"courses_serve/example/internals/app/processors"
	"courses_serve/example/internals/cfg"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppServer struct {
	config cfg.Cfg
	ctx    context.Context
	serv   *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer { //задаем поля нашего сервера, для его старта нам нужен контекст и конфигурация
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *AppServer) Serve() {
	log.Println("Server starting")
	log.Println(server.config.GetDBConnetcUrl())
	var err error
	//создаем пул соединений с БД и сохраним его для закрытия при остановке приложения
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBConnetcUrl())
	if err != nil {
		log.Fatalln(err)
	}

	subjectsStorage := db.NewSubjectsStorage(server.db)
	lessonsStorage := db.NewLessonsStorage(server.db)

	subjectsProcessor := processors.NewSubjectsProcessor(subjectsStorage)
	lessonsProcessor := processors.NewLessonsProcessor(lessonsStorage)

	subjectsHandler := handlers.NewSubjectsHandler(subjectsProcessor)
	lessonsHandler := handlers.NewLessonsHandler(lessonsProcessor)

	routes := api.CreateRoute(lessonsHandler, subjectsHandler) //хендлеры напрямую используются в путях
	routes.Use(middleware.RequestLog)                          //middleware используем здесь, хотя можно было бы и в CreateRoutes

	server.serv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.serv.ListenAndServe() //запускаем сервер
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return
}

func (server *AppServer) Shutdown() {
	log.Println("server stopped")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close() //закрываем соединение с БД
	defer func() {
		cancel()
	}()
	var err error
	if err = server.serv.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
