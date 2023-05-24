package dependency

import (
	"gb-auth-gate/config"
	accountRepo "gb-auth-gate/internal/account/repository"
	accountUC "gb-auth-gate/internal/account/usecase"
	authRepo "gb-auth-gate/internal/auth/repository"
	authUC "gb-auth-gate/internal/auth/usecase"
	calcUC "gb-auth-gate/internal/calculations/usecase"
	mdwHttp "gb-auth-gate/internal/pkg/middleware/delivery/http"
	uiUC "gb-auth-gate/internal/ui/usecase"
	"gb-auth-gate/pkg/damqp/kafka"
	"gb-auth-gate/pkg/damqp/rabbit"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp"
	"gb-auth-gate/pkg/thttp/server"
	"gb-auth-gate/pkg/tlogger"
	"gb-auth-gate/pkg/tsecure"
	tstorageCache "gb-auth-gate/pkg/tstorage/cache"
	tstorageRelational "gb-auth-gate/pkg/tstorage/relational"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config": config.BuildConfig,

	"fernet": tsecure.BuildFernetEncryptor,

	"postgres": tstorageRelational.BuildPostgres,
	"redis":    tstorageCache.BuildRedis,

	"httpClient": thttp.BuildHttpClient,

	"logger": tlogger.BuildLogger,

	"rabbit": rabbit.BuildRabbitMQ,
	"kafka":  kafka.BuildKafka,

	"authUC":   authUC.BuildAuthUsecase,
	"authRepo": authRepo.BuildPostgresRepository,

	"middleware":        mdwHttp.BuildMiddlewareManager,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,

	"calcUC": calcUC.BuildCalculationsUseCase,

	"uiUC": uiUC.BuildUiUseCase,

	"accountRepo": accountRepo.BuildPostgresRepository,
	"accountUC":   accountUC.BuildAccountUseCase,

	"app": server.BuildFiberApp,
}

const TagDI = "di"
