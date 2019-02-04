package handlers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/handlers"
)

var _ = Describe("Get Payments Endpoints", func() {
	var (
		ts          *httptest.Server
		res         *http.Response
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
	})

	AfterEach(func() {
		ts.Close()
	})

	Describe("when calling GET /payments/{id}", func() {
		Describe("with an incorrectly formatted payment id", func() {
			BeforeEach(func() {
				res, err = http.Get(ts.URL + "/payments/ABCDWRONGFORMAT")
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
				res, err = http.Get(ts.URL + "/payments/" + paymentId.String())
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 404", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
			})
		})

		Describe("with 1 existing payment", func() {
			BeforeEach(func() {
				attributes = createDomainAttributes(100.00)
				paymentId, _ = uuid.Parse("cae9aa62-0ea1-432b-baee-c0ff4b1d889e")

				err = h.CommandHandler.HandleCommand(context.Background(), &domain.Create{
					ID:         paymentId,
					Attributes: attributes,
				})
				Expect(err).NotTo(HaveOccurred())

				res, err = http.Get(ts.URL + "/payments/" + paymentId.String())
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 200", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusOK))
			})
			It("should have a content type of JSON", func() {
				Expect(res.Header["Content-Type"][0]).Should(Equal("application/json"))
			})
			It("should have JSON body with 1 payment", func() {
				g, err := ioutil.ReadFile("testdata/get_single_resp.golden")
				if err != nil {
					GinkgoT().Fatalf("failed reading .golden: %s", err)
				}
				bodyBytes, _ := ioutil.ReadAll(res.Body)
				Expect(string(bodyBytes)).Should(MatchJSON(string(g)))
			})

			Describe("after updating the payment", func() {
				BeforeEach(func() {
					attributes = createDomainAttributes(300.00)

					err = h.CommandHandler.HandleCommand(context.Background(), &domain.Update{
						ID:         paymentId,
						Attributes: attributes,
					})
					Expect(err).NotTo(HaveOccurred())

					res, err = http.Get(ts.URL + "/payments/" + paymentId.String())
				})

				It("should not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})
				It("should have HTTP status code 200", func() {
					Expect(res.StatusCode).Should(Equal(http.StatusOK))
				})
				It("should have a content type of JSON", func() {
					Expect(res.Header["Content-Type"][0]).Should(Equal("application/json"))
				})
				It("should have JSON body with 1 payment that is version 2", func() {
					g, err := ioutil.ReadFile("testdata/get_single_after_update_resp.golden")
					if err != nil {
						GinkgoT().Fatalf("failed reading .golden: %s", err)
					}
					bodyBytes, _ := ioutil.ReadAll(res.Body)
					Expect(string(bodyBytes)).Should(MatchJSON(string(g)))
				})

				Describe("with deleting the payment", func() {
					BeforeEach(func() {
						err = h.CommandHandler.HandleCommand(context.Background(), &domain.Delete{
							ID:     paymentId,
							Reason: "testing",
						})
						Expect(err).NotTo(HaveOccurred())

						res, err = http.Get(ts.URL + "/payments/" + paymentId.String())
					})

					It("should not error", func() {
						Expect(err).NotTo(HaveOccurred())
					})
					It("should have HTTP status code 404", func() {
						Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
					})
				})
			})
		})
	})
})
