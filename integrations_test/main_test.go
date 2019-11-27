package integrations

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/DATA-DOG/godog"
	"github.com/egor1344/banner/rotation_banner/proto/server"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var dbDsn = os.Getenv("DB_DSN")
var amqpDSN = os.Getenv("AMQP_DSN")
var queueName = os.Getenv("QUEUE_NAME")
var grpcPort = os.Getenv("API_GRPC_PORT")
var restPort = os.Getenv("API_REST_PORT")

type ServerTest struct {
	DB *sqlx.DB

	Client server.RotationBannerClient

	addBannerResponse       *server.AddBannerResponse
	delBannerResponse       *server.DelBannerResponse
	countTransitionResponse *server.CountTransitionResponse
	getBannerResponse       *server.GetBannerResponse

	responseStatusCode int
	responseBody       []byte
}

// connectDB - подключение к БД
func (test *ServerTest) connectDB(*gherkin.Feature) {
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	test.DB = db
}

// truncateDb - очистка таблиц для адекватной проверки фукционала
func (test *ServerTest) truncateDb(*gherkin.Feature) {
	tables := []string{"banners", "rotations", "slot", "soc_dem_group", "statistic"}
	for _, table := range tables {
		_, err := test.DB.Query("truncate table " + table + " restart identity cascade;")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Wait 10s for service availability...")
	time.Sleep(10 * time.Second)

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		GrpcContext(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"features/grpc"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}
	status = godog.RunWithOptions("integration", func(s *godog.Suite) {
		RestContext(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"features/rest"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
