package handler

import (
	"github.com/asaskevich/govalidator/v12"
	"github.com/gofiber/fiber/v3"
	"github.com/phamvlap/url-shortener/internal/config"
	apperrors "github.com/phamvlap/url-shortener/internal/errors"
	"github.com/phamvlap/url-shortener/internal/repository"
	"github.com/phamvlap/url-shortener/internal/utility"
)

const DefaultExpirationTime = 3600 // 1 hour in seconds

func RedirectURLHandler(c fiber.Ctx) error {
	// Get the short ID from the URL parameters
	shortID := c.Params("shortID")

	// Get the original URL from the repository using the short ID
	urlRepository := repository.NewURLRepository()
	originalURL, err := urlRepository.GetOriginalURL(shortID)

	if err != nil {
		return apperrors.NotFoundError("Short URL not found", err)
	}

	// Increment the counter for the short ID
	urlRepository.IncrCounter(shortID)

	return c.Redirect().Status(fiber.StatusMovedPermanently).To(originalURL)
}

type ShortenURLRequest struct {
	URL string `json:"url"`
}

type ShortenURLResponse struct {
	URL        string `json:"url"`
	ShortID    string `json:"short_id"`
	ShortURL   string `json:"short_url"`
	Expiration int64  `json:"expiration"`
}

func ShortenURLHandler(c fiber.Ctx) error {
	config := config.GetConfig()

	body := new(ShortenURLRequest)

	// Bind the request body to the ShortenURLRequest struct
	if err := c.Bind().Body(&body); err != nil {
		return apperrors.BadRequestError("Invalid request body", err)
	}

	// Validate the URL format
	if !govalidator.IsURL(body.URL) {
		return apperrors.BadRequestError("Invalid URL format", nil)
	}

	// Enforce HTTPS protocol
	body.URL = utility.EnforceHTTPSProtocol(body.URL)

	// Generate a unique short ID and construct the short URL
	shortID := utility.GenerateShortID()
	shortURL := config.App.Domain + "/" + shortID

	// Save the original URL and short ID mapping in the repository with an expiration time (e.g., 1 hour)
	urlRepository := repository.NewURLRepository()
	urlRepository.SaveURL(body.URL, shortID, DefaultExpirationTime)

	response := ShortenURLResponse{
		URL:        body.URL,
		ShortID:    shortID,
		ShortURL:   shortURL,
		Expiration: DefaultExpirationTime,
	}

	return c.JSON(response)
}
