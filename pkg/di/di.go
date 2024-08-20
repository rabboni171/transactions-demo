package di

import (
	"github.com/rabboni171/transactions-demo/configs"
	"github.com/rabboni171/transactions-demo/internal/handler"
	"github.com/rabboni171/transactions-demo/internal/service"
	"github.com/rabboni171/transactions-demo/pkg/db"
	"github.com/rabboni171/transactions-demo/pkg/logger"
	"github.com/rabboni171/transactions-demo/pkg/router"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"net/http"
)

func InitializeApp() *fx.App {
	return fx.New(
		fx.Provide(logger.NewLogger),

		// Provide the configuration
		fx.Provide(configs.InitConfig),

		// Provide the database with a constructor that takes config and logger
		fx.Provide(db.NewDatabase),
		fx.Provide(handler.NewHandler),
		fx.Provide(service.NewService),
		fx.Provide(db.NewDatabase),
		fx.Invoke(func(h *handler.Handler) {
			r := router.InitRoutes(h)
			http.Handle("/", r)
			go func() {
				err := http.ListenAndServe(":8080", nil)
				if err != nil {
				}
			}()
		}),
		fx.WithLogger(func(logger *logrus.Logger) fxevent.Logger {
			return &fxevent.ConsoleLogger{W: logger.Writer()}
		}),
	)

}
