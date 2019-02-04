package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/handlers"
)

var _ = Describe("Delete Payments Endpoints", func() {
	Describe("when calling DELETE /payments/{id}", func() {
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
		)

		BeforeEach(func() {
			paymentId, _ = uuid.Parse("cae9aa62-0ea1-432b-baee-c0ff4b1d889e")

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

		Describe("with an incorrectly formatted payment id", func() {
			BeforeEach(func() {
				req, err = http.NewRequest("DELETE", "/payments/ABCDWRONGFORMAT", nil)
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

		Describe("with non existing payment id", func() {
			BeforeEach(func() {
				req, err = http.NewRequest("DELETE", "/payments/bae9aa62-0ea1-432b-baee-c0ff4b1d889e", nil)
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

		Describe("with 1 payment", func() {
			BeforeEach(func() {
				attributes = createDomainAttributes(100.00)
				err = h.CommandHandler.HandleCommand(context.Background(), &domain.Create{
					ID:         paymentId,
					Attributes: attributes,
				})
				Expect(err).NotTo(HaveOccurred())
				waitForEvent(h.EventBus, domain.Created)

				req, err = http.NewRequest("DELETE", "/payments/"+paymentId.String(), nil)
				req.Header.Set("Content-Type", "application/json")
				ts.Config.Handler.ServeHTTP(recorder, req)
				res = recorder.Result()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 204", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusNoContent))
			})
			//NOTE: the below sleep is a hack
			It("should have deleted the payment", func() {
				time.Sleep(1 * time.Second)
				data, errFind := h.Repo.Find(context.Background(), paymentId)
				Expect(data).To(BeNil())
				Expect(errFind).Should(HaveOccurred())
			})
		})
	})
})
