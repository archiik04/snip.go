package handlers

import (
	"context"
	"net/http"
	"time"

	"snip-go/store"
	"snip-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type URLHandler struct {
	store   *store.RedisStore
	baseURL string
}

func NewURLHandler(store *store.RedisStore, baseURL string) *URLHandler {
	return &URLHandler{store: store, baseURL: baseURL}
}

type ShortenRequest struct {
	URL        string `json:"url" binding:"required,url"`
	CustomCode string `json:"custom_code,omitempty"`
	TTLDays    int    `json:"ttl_days,omitempty"`
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := req.CustomCode
	if code == "" {
		var err error
		code, err = utils.GenerateShortCode(7)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate code"})
			return
		}
	}

	ttl := time.Duration(0)
	if req.TTLDays > 0 {
		ttl = time.Duration(req.TTLDays) * 24 * time.Hour
	}

	ctx := context.Background()
	if err := h.store.Set(ctx, code, req.URL, ttl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url":    h.baseURL + "/" + code,
		"code":         code,
		"original_url": req.URL,
		"expires_in":   req.TTLDays,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")
	ctx := context.Background()

	originalURL, err := h.store.Get(ctx, code)
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found or expired"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "lookup failed"})
		return
	}

	go h.store.IncrClicks(ctx, code)
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func (h *URLHandler) Stats(c *gin.Context) {
	code := c.Param("code")
	ctx := context.Background()

	originalURL, err := h.store.Get(ctx, code)
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	clicks, _ := h.store.GetClicks(ctx, code)
	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"original_url": originalURL,
		"clicks":       clicks,
		"short_url":    h.baseURL + "/" + code,
	})
}
