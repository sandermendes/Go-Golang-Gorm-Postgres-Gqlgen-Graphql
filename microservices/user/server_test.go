package user_test

import (
	"context"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sandermendes/Go-Golang-Gorm-Postgres-Gqlgen-Graphql/main/microservices/user"
	userv1 "github.com/sandermendes/Go-Golang-Gorm-Postgres-Gqlgen-Graphql/main/shared/protobufs/_generated/user/v1"
	serviceConnection "github.com/sandermendes/Go-Golang-Gorm-Postgres-Gqlgen-Graphql/main/shared/service_connection"
)

const USER_SERVICE_PORT = "6010"

var _ = Describe("User", Ordered, func() {
	var (
		userConn userv1.UserServiceClient
		err      error
		ctx      context.Context
	)

	BeforeAll(func() {
		go user.ListenGRPC(USER_SERVICE_PORT)

		defer func() {
			userConn, err = serviceConnection.GetUserConnection("localhost:" + USER_SERVICE_PORT)
			if err != nil {
				log.Fatalf("failed to connect to userService: %v", err)
			}

			// userService = user.NewService()
			ctx = context.TODO()
		}()
	})

	It("Create user with invalid email", func() {
		_, err := userConn.CreateUser(ctx, &userv1.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john-doe",
			Password:  "123456",
		})

		Expect(err).Should(HaveOccurred())
	})

	It("Create user with valid email", func() {
		userResponse, err := userConn.CreateUser(ctx, &userv1.CreateUserRequest{
			FirstName: "Teste",
			LastName:  "Test",
			Email:     "john-doe@dummy-corp.com",
			Password:  "123456",
		})
		createdUserId := userResponse.GetId()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(userResponse).ToNot(BeNil())
		Expect(createdUserId).ShouldNot(BeEmpty())
	})

	It("List User by Id or Email", func() {

		createUserResponse, _ := userConn.CreateUser(ctx, &userv1.CreateUserRequest{
			FirstName: "List",
			LastName:  "Test",
			Email:     "list-test@dummy-corp.com",
			Password:  "123456",
		})
		createdUserId := createUserResponse.GetId()
		createdUserEmail := createUserResponse.GetEmail()

		listByIdUserResponse, err := userConn.GetUser(ctx, &userv1.UpdateUserRequest{
			Id: createdUserId,
		})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(listByIdUserResponse).ToNot(BeNil())

		listByEmailUserResponse, err := userConn.GetUser(ctx, &userv1.UpdateUserRequest{
			Email: createdUserEmail,
		})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(listByEmailUserResponse).ToNot(BeNil())
	})

	It("Update User", func() {

		createUserResponse, _ := userConn.CreateUser(ctx, &userv1.CreateUserRequest{
			FirstName: "update",
			LastName:  "Test",
			Email:     "update-test@dummy-corp.com",
			Password:  "123456",
		})
		createdUserId := createUserResponse.GetId()

		updatedUserResponse, err := userConn.UpdateUser(ctx, &userv1.UpdateUserRequest{
			Id:        createdUserId,
			FirstName: "Teste Updated",
			LastName:  "Test Lastname",
			Email:     "update-test-1@dummy-corp.com",
		})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(updatedUserResponse).ToNot(BeNil())
	})

	It("Delete User", func() {
		createUserResponse, _ := userConn.CreateUser(ctx, &userv1.CreateUserRequest{
			FirstName: "update",
			LastName:  "Test",
			Email:     "delete-test@dummy-corp.com",
			Password:  "123456",
		})
		createdUserId := createUserResponse.GetId()

		userResponse, err := userConn.DeleteUser(ctx, &userv1.UpdateUserRequest{
			Id: createdUserId,
		})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(userResponse).ToNot(BeNil())
	})
})
