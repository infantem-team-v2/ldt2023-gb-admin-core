package dependency

import (
	"gb-admin-core/config"
	accountRepo "gb-admin-core/internal/account/repository"
	accountUC "gb-admin-core/internal/account/usecase"
	authRepo "gb-admin-core/internal/auth/repository"
	authUC "gb-admin-core/internal/auth/usecase"
	mdwHttp "gb-admin-core/internal/pkg/middleware/delivery/http"
	uiUC "gb-admin-core/internal/ui/usecase"
	"gb-admin-core/pkg/damqp/kafka"
	"gb-admin-core/pkg/damqp/rabbit"
	"gb-admin-core/pkg/terrors"
	"gb-admin-core/pkg/thttp"
	"gb-admin-core/pkg/thttp/server"
	"gb-admin-core/pkg/tlogger"
	"gb-admin-core/pkg/tsecure"
	tstorageCache "gb-admin-core/pkg/tstorage/cache"
	tstorageRelational "gb-admin-core/pkg/tstorage/relational"
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

	"uiUC": uiUC.BuildUiUseCase,

	"accountRepo": accountRepo.BuildPostgresRepository,
	"accountUC":   accountUC.BuildAccountUseCase,

	"app": server.BuildFiberApp,
}

const TagDI = "di"
