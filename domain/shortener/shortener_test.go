package shortener_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
	"url-shortener/domain/shortener"
	mock_repository "url-shortener/mocks/repository"
)

var errSomething = errors.New("something error")

var _ = Describe("Post Service", func() {
	var (
		mockCtrl      *gomock.Controller
		shortenerUC   shortener.Service
		repo          *mock_repository.MockRepository
		mockLink      shortener.Link
		mockKeys      []string
		mockLinks     []shortener.Link
		mockLinksDesc []shortener.Link
		now           time.Time
		threeDays     time.Time
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockRepository(mockCtrl)
		shortenerUC = shortener.NewShortenerService(repo)
		now = time.Now()
		threeDays = now.AddDate(0, 0, 3)

		mockLink = shortener.Link{
			OriginalURL: "https://facebook.com/",
			ShortURL:    "bnyQI",
			Clicks:      0,
			ExpireAt:    &threeDays,
		}

		mockLinks = []shortener.Link{
			{
				OriginalURL: "https://facebook.com/",
				ShortURL:    "bnyQI",
				Clicks:      0,
				ExpireAt:    &threeDays,
			}, {
				OriginalURL: "https://facebook.com/",
				ShortURL:    "bnyQE",
				Clicks:      2,
				ExpireAt:    &threeDays,
			},
		}
		mockLinksDesc = []shortener.Link{
			{
				OriginalURL: "https://facebook.com/",
				ShortURL:    "bnyQE",
				Clicks:      2,
				ExpireAt:    &threeDays,
			},
			{
				OriginalURL: "https://facebook.com/",
				ShortURL:    "bnyQI",
				Clicks:      0,
				ExpireAt:    &threeDays,
			},
		}

		mockKeys = []string{"bnyQI", "bnyQE"}
	})

	Describe("CreateShortenedURL", func() {
		mockRequest := shortener.Link{
			OriginalURL: "https://facebook.com/",
			ShortURL:    "bnyQI",
			Clicks:      0,
			ExpireAt:    &threeDays,
		}

		It("return success", func() {
			repo.EXPECT().Store(gomock.Any()).Return(nil)
			repo.EXPECT().StoreKey(gomock.Any()).Return(nil)
			err := shortenerUC.CreateShortenedURL(&mockRequest)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			repo.EXPECT().Store(gomock.Any()).Return(errSomething)
			repo.EXPECT().StoreKey(gomock.Any()).Return(nil)
			err := shortenerUC.CreateShortenedURL(&mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("GetByShortenedURL", func() {
		mockRequest := "bnyQI"

		It("return success", func() {
			repo.EXPECT().GetByShortenedURL(gomock.Any()).Return(&mockLink, nil)
			repo.EXPECT().Store(gomock.Any()).Return(nil)
			link, err := shortenerUC.GetByShortenedURL(mockRequest)
			Expect(err).Should(Succeed())
			Expect(link).Should(Equal(&mockLink))
		})

		It("return error", func() {
			repo.EXPECT().GetByShortenedURL(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().Store(gomock.Any()).Return(nil)
			_, err := shortenerUC.GetByShortenedURL(mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("GetAll", func() {
		sortBy := "asc"
		sortByDesc := "desc"

		It("return success asc", func() {
			repo.EXPECT().GetAllKeys().Return(mockKeys, nil)
			for key := range mockKeys {
				repo.EXPECT().GetByShortenedURL(gomock.Any()).Return(&mockLinks[key], nil)
			}
			repo.EXPECT().SetKeys(gomock.Any()).Return(nil)
			links, err := shortenerUC.GetAll(sortBy)
			Expect(err).Should(Succeed())
			Expect(links).Should(Equal(mockLinks))
		})

		It("return success desc", func() {
			repo.EXPECT().GetAllKeys().Return(mockKeys, nil)
			for key := range mockKeys {
				repo.EXPECT().GetByShortenedURL(gomock.Any()).Return(&mockLinks[key], nil)
			}
			repo.EXPECT().SetKeys(gomock.Any()).Return(nil)
			links, err := shortenerUC.GetAll(sortByDesc)
			Expect(err).Should(Succeed())
			Expect(links).Should(Equal(mockLinksDesc))
		})

		It("return error", func() {
			repo.EXPECT().GetAllKeys().Return(mockKeys, errSomething)
			for key := range mockKeys {
				repo.EXPECT().GetByShortenedURL(gomock.Any()).Return(&mockLinks[key], nil)
			}
			repo.EXPECT().SetKeys(gomock.Any()).Return(nil)
			_, err := shortenerUC.GetAll(sortBy)
			Expect(err).Should(HaveOccurred())
		})
	})
})
