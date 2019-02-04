package handlers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/handlers"
)

var _ = Describe("Update Payments Endpoints", func() {
	var (
		ts          *httptest.Server
		res         *http.Response
		req         *http.Request
		recorder    *httptest.ResponseRecorder
		config      config.Config
		restHandler http.Handler
		h           *handlers.Handlers
		err         error
		attributes  domain.PaymentAttributes
		paymentId   uuid.UUID
		bodyBytes   []byte
	)

	BeforeEach(func() {
		paymentId, _ = uuid.Parse("cae9aa62-0ea1-432b-baee-c0ff4b1d889e")
		bodyBytes, _ = ioutil.ReadFile("testdata/update_body_req.golden")

		config = createTestConfig()
		restHandler, h, err = createTestHandler(config)
		if err != nil {
			Fail(err.Error())
		}
		ts = httptest.NewServer(restHandler)
		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {
		ts.Close()
	})

	Describe("when calling PUT /payments/{id}", func() {
		Describe("with an incorrectly formatted payment id", func() {
			BeforeEach(func() {
				req, err = http.NewRequest("PUT", "/payments/ABCDWRONGFORMAT", strings.NewReader(string(bodyBytes)))
				req.Header.Set("Content-Type", "application/json")
				ts.Config.Handler.ServeHTTP(recorder, req)
				res = recorder.Result()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 422", func() {
				Expect(res.StatusCode).Should(Equal(422))
			})
		})

		Context("with non existing payment id", func() {
			BeforeEach(func() {
				req, err = http.NewRequest("PUT", "/payments/bae9aa62-0ea1-432b-baee-c0ff4b1d889e", strings.NewReader(string(bodyBytes)))
				req.Header.Set("Content-Type", "application/json")
				ts.Config.Handler.ServeHTTP(recorder, req)
				res = recorder.Result()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 404", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
			})
		})

		Context("with 1 payment", func() {
			BeforeEach(func() {
				attributes = createDomainAttributes(100.00)
				err = h.CommandHandler.HandleCommand(context.Background(), &domain.Create{
					ID:         paymentId,
					Attributes: attributes,
				})
				Expect(err).NotTo(HaveOccurred())
				waitForEvent(h.EventBus, domain.Created)

				req, err = http.NewRequest("PUT", "/payments/"+paymentId.String(), strings.NewReader(string(bodyBytes)))
				req.Header.Set("Content-Type", "application/json")
				ts.Config.Handler.ServeHTTP(recorder, req)
				res = recorder.Result()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 200", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusOK))
			})
			It("should have updated the payment", func() {
				time.Sleep(1 * time.Second)
				data, err := h.Repo.Find(context.Background(), paymentId)
				Expect(err).ShouldNot(HaveOccurred())

				domPayment := data.(*domain.Payment)

				Expect(domPayment.Version).Should(Equal(2))
				Expect(float32(domPayment.Attributes.Amount)).Should(Equal(float32(300.00)))
			})
		})
	})
})
