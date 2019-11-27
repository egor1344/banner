package integrations

import (
	"context"
	"errors"
	"log"

	"github.com/egor1344/banner/rotation_banner/proto/banner"
	"github.com/egor1344/banner/rotation_banner/proto/slot"
	"google.golang.org/grpc"

	"github.com/DATA-DOG/godog"
	"github.com/egor1344/banner/rotation_banner/proto/server"
	_ "github.com/jackc/pgx/stdlib"
)

func (test *ServerTest) iAddBannerGprcrequestTo(arg1 string) error {
	// Добавление события
	conn, err := grpc.Dial(arg1+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to grpc server ", arg1, err)
	}
	defer conn.Close()
	client := server.NewRotationBannerClient(conn)
	test.Client = client
	req := &server.AddBannerRequest{Banner: &banner.Banner{Id: 1, Slot: &slot.Slot{Id: 1}, Description: "Test banner"}}
	resp, err := test.Client.AddBanner(context.Background(), req)
	if err != nil {
		return err
	}
	test.addBannerResponse = resp
	return nil
}

func (test *ServerTest) theResponseAddBannerMustContainStatus() error {
	resp := test.addBannerResponse
	ev := resp.GetStatus()
	if !ev {
		return errors.New("wrong status")
	}
	err := resp.GetError()
	if err != "" {
		return errors.New(err)
	}
	return nil
}

func (test *ServerTest) iDelBannerGprcrequestTo(arg1 string) error {
	// Удаление баннера из ротации
	conn, err := grpc.Dial(arg1+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to grpc server ", arg1, err)
	}
	defer conn.Close()
	client := server.NewRotationBannerClient(conn)
	test.Client = client
	req := &server.DelBannerRequest{Id: 1}
	resp, err := test.Client.DelBanner(context.Background(), req)
	if err != nil {
		return err
	}
	test.delBannerResponse = resp
	return nil
}

func (test *ServerTest) theResponseDelBannerMustContainStatus() error {
	resp := test.delBannerResponse
	ev := resp.GetStatus()
	if !ev {
		return errors.New("wrong status")
	}
	err := resp.GetError()
	if err != "" {
		return errors.New(err)
	}
	return nil
}

func (test *ServerTest) iCountTransitionBannerGprcrequestTo(arg1 string) error {
	// Клик по баннеру
	conn, err := grpc.Dial(arg1+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to grpc server ", arg1, err)
	}
	defer conn.Close()
	client := server.NewRotationBannerClient(conn)
	test.Client = client
	req := &server.CountTransitionRequest{IdBanner: 1, IdSocDemGroup: 1, IdSlot: 1}
	resp, err := test.Client.CountTransition(context.Background(), req)
	if err != nil {
		return err
	}
	test.countTransitionResponse = resp
	return nil
}

func (test *ServerTest) theResponseCountTransitionMustContainStatus() error {
	resp := test.countTransitionResponse
	ev := resp.GetStatus()
	if !ev {
		return errors.New("wrong status")
	}
	err := resp.GetError()
	if err != "" {
		return errors.New(err)
	}
	return nil
}

func (test *ServerTest) iGetBannerBannerGprcrequestTo(arg1 string) error {
	// Удаление баннера из ротации
	conn, err := grpc.Dial(arg1+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to grpc server ", arg1, err)
	}
	defer conn.Close()
	client := server.NewRotationBannerClient(conn)
	test.Client = client
	req := &server.GetBannerRequest{IdSocDemGroup: 1, IdSlot: 1}
	resp, err := test.Client.GetBanner(context.Background(), req)
	if err != nil {
		return err
	}
	test.getBannerResponse = resp
	return nil
}

func (test *ServerTest) theResponseMustContainIdBanner() error {
	resp := test.getBannerResponse
	err := resp.GetError()
	if err != "" {
		return errors.New(err)
	}
	id := resp.GetIdBanner()
	if id != 1 {
		return errors.New("wrong id banner")
	}
	return nil
}

func GrpcContext(s *godog.Suite) {
	test := new(ServerTest)

	s.BeforeFeature(test.connectDB)
	s.BeforeFeature(test.truncateDb)
	s.AfterFeature(test.truncateDb)

	s.Step(`^I add banner gprc-request to "([^"]*)"$`, test.iAddBannerGprcrequestTo)
	s.Step(`^The response add banner must contain status$`, test.theResponseAddBannerMustContainStatus)
	s.Step(`^I del banner gprc-request to "([^"]*)"$`, test.iDelBannerGprcrequestTo)
	s.Step(`^The response del banner must contain status$`, test.theResponseDelBannerMustContainStatus)
	s.Step(`^I count transition banner gprc-request to "([^"]*)"$`, test.iCountTransitionBannerGprcrequestTo)
	s.Step(`^The response count transition must contain status$`, test.theResponseCountTransitionMustContainStatus)
	s.Step(`^I Get banner banner gprc-request to "([^"]*)"$`, test.iGetBannerBannerGprcrequestTo)
	s.Step(`^The response must contain id banner$`, test.theResponseMustContainIdBanner)
}
