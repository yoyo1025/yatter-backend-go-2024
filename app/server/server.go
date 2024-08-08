package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"yatter-backend-go/app/config"
	"yatter-backend-go/app/dao"
	"yatter-backend-go/app/handler"
	"yatter-backend-go/app/usecase"
)

func Run() error {
	db, err := dao.NewDB(config.MySQLConfig())
	if err != nil {
		return err
	}
	defer db.Close()

	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)

	accountUsecase := usecase.NewAcocunt(db, dao.NewAccount(db))
	statusUsecase := usecase.NewStatus(db, dao.NewStatus(db))
	timelineUsecase := usecase.NewTimeline(db, dao.NewTimeline(db))

	r := handler.NewRouter(
		accountUsecase,
		statusUsecase,
		timelineUsecase,
		dao.NewAccount(db),
	)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel() // cancel関数を呼び出してリソースを解放する
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}
