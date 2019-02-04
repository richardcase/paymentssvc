package handlers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/handlers"
)

var _ = Describe("Create Payments Endpoints", func() {
	var (
		ts          *httptest.Server
		res         *http.Response
		config      config.Config
		restHandler http.Handler
		h           *handlers.Handlers
		err         error
		bodyBytes   []byte
	)

	BeforeEach(func() {
		bodyBytes, _ = ioutil.ReadFile("testdata/create_http_body.golden")

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

	Describe("when calling POST /payments", func() {

		Describe("with non existing payment id", func() {
			BeforeEach(func() {
				res, err = http.Post(ts.URL+"/payments", "application/json", strings.NewReader(string(bodyBytes)))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have HTTP status code 201", func() {
				Expect(res.StatusCode).Should(Equal(http.StatusCreated))
			})
			It("should have a Location header", func() {
				hdr := res.Header["Location"]
				Expect(hdr).To(Not(BeNil()))
				Expect(len(hdr)).To(Equal(1))

				loc := hdr[0]
				Expect(strings.HasPrefix(loc, "/")).To(BeTrue())
			})

			It("should have created the payment", func() {
				hdr := res.Header["Location"][0]
				parts := strings.Split(hdr, "/")
				paymentId, _ := uuid.Parse(parts[2])

				data, err := h.Repo.Find(context.Background(), paymentId)
				Expect(err).ShouldNot(HaveOccurred())

				domPayment := data.(*domain.Payment)

				Expect(domPayment.Version).Should(Equal(1))
				Expect(float32(domPayment.Attributes.Amount)).Should(Equal(float32(200.00)))
			})
		})
	})
})
