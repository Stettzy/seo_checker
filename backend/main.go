package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SEOCheckRequest struct {
	URL string `json:"url"`
}

type MetaTag struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type HeadingStructure struct {
	Level string `json:"level"`
	Text  string `json:"text"`
	Count int    `json:"count"`
}

type ImageAnalysis struct {
	TotalImages    int      `json:"totalImages"`
	WithAlt        int      `json:"withAlt"`
	WithoutAlt     int      `json:"withoutAlt"`
	LargeImages    int      `json:"largeImages"`
	OptimizedCount int      `json:"optimizedCount"`
	ImageURLs      []string `json:"imageUrls"`
}

type LinkAnalysis struct {
	InternalLinks int      `json:"internalLinks"`
	ExternalLinks int      `json:"externalLinks"`
	BrokenLinks   int      `json:"brokenLinks"`
	NoFollowLinks int      `json:"noFollowLinks"`
	InternalURLs  []string `json:"internalUrls"`
	ExternalURLs  []string `json:"externalUrls"`
}

type PerformanceMetrics struct {
	PageSize      int64   `json:"pageSize"`
	LoadTime      float64 `json:"loadTime"`
	ResourceCount int     `json:"resourceCount"`
	JSCount       int     `json:"jsCount"`
	CSSCount      int     `json:"cssCount"`
	ImageCount    int     `json:"imageCount"`
}

type ContentAnalysis struct {
	WordCount        int                `json:"wordCount"`
	CharacterCount   int                `json:"characterCount"`
	ParagraphCount   int                `json:"paragraphCount"`
	ReadabilityScore float64            `json:"readabilityScore"`
	KeywordDensity   map[string]float64 `json:"keywordDensity"`
	TopKeywords      []string           `json:"topKeywords"`
}

type TechnicalSEO struct {
	HasSitemap       bool     `json:"hasSitemap"`
	HasRobotsTxt     bool     `json:"hasRobotsTxt"`
	IsHTTPS          bool     `json:"isHttps"`
	HasSSL           bool     `json:"hasSSL"`
	ServerResponse   int      `json:"serverResponse"`
	Canonical        string   `json:"canonical"`
	MobileResponsive bool     `json:"mobileResponsive"`
	AmpEnabled       bool     `json:"ampEnabled"`
	SchemaMarkup     []string `json:"schemaMarkup"`
	MetaKeywords     string   `json:"metaKeywords"`
	HasMetaKeywords  bool     `json:"hasMetaKeywords"`
}

type SocialMetaTags struct {
	OpenGraph map[string]string `json:"openGraph"`
	Twitter   map[string]string `json:"twitter"`
	LinkedIn  map[string]string `json:"linkedIn"`
}

type SEOCheckResponse struct {
	URL              string             `json:"url"`
	Score            int                `json:"score"`
	Title            string             `json:"title"`
	TitleLength      int                `json:"titleLength"`
	Description      string             `json:"description"`
	DescLength       int                `json:"descLength"`
	MetaTags         []MetaTag          `json:"metaTags"`
	HeadingStructure []HeadingStructure `json:"headingStructure"`
	Images           ImageAnalysis      `json:"images"`
	Links            LinkAnalysis       `json:"links"`
	Performance      PerformanceMetrics `json:"performance"`
	Content          ContentAnalysis    `json:"content"`
	Technical        TechnicalSEO       `json:"technical"`
	CriticalIssues   []Issue            `json:"criticalIssues"`
	Warnings         []Issue            `json:"warnings"`
	Recommendations  []Issue            `json:"recommendations"`
	LastUpdated      time.Time          `json:"lastUpdated"`
	LoadTime         float64            `json:"loadTime"`
	SocialMeta       SocialMetaTags     `json:"socialMeta"`
	Competitors      []string           `json:"competitors"`
}

type Issue struct {
	Category    string `json:"category"`
	Severity    string `json:"severity"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
	HowToFix    string `json:"howToFix"`
	Examples    string `json:"examples,omitempty"`
}

func main() {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/analyze", analyzeSEO)
	r.Run(":3000")
}

func analyzeSEO(c *gin.Context) {
	var req SEOCheckRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	startTime := time.Now()
	resp, err := client.Get(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch URL"})
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Calculate load time
	loadTime := time.Since(startTime).Seconds()

	// Create a new reader from the bytes for goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(bodyBytes)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse HTML"})
		return
	}

	seoResponse := SEOCheckResponse{
		URL:         req.URL,
		Score:       100,
		LastUpdated: time.Now(),
		LoadTime:    loadTime,
	}

	// Analyze title
	seoResponse.Title = doc.Find("title").Text()
	seoResponse.TitleLength = len(seoResponse.Title)
	if seoResponse.Title == "" {
		addIssue(&seoResponse, Issue{
			Category:    "Meta Tags",
			Severity:    "Critical",
			Title:       "Missing Title Tag",
			Description: "The page is missing a title tag, which is crucial for SEO.",
			Impact:      "Severely impacts search engine rankings and click-through rates.",
			HowToFix:    "Add a descriptive title tag between 50-60 characters that includes your main keyword.",
			Examples:    "<title>Your Primary Keyword - Your Brand Name</title>",
		})
		seoResponse.Score -= 10
	} else if len(seoResponse.Title) < 30 || len(seoResponse.Title) > 60 {
		addIssue(&seoResponse, Issue{
			Category:    "Meta Tags",
			Severity:    "Warning",
			Title:       "Title Length Not Optimal",
			Description: "Title tag length is not within the recommended range.",
			Impact:      "May be truncated in search results or not fully optimized for search engines.",
			HowToFix:    "Adjust title length to be between 30-60 characters while maintaining keyword relevance.",
		})
		seoResponse.Score -= 5
	}

	// Analyze meta description
	seoResponse.Description = doc.Find("meta[name='description']").AttrOr("content", "")
	seoResponse.DescLength = len(seoResponse.Description)
	if seoResponse.Description == "" {
		addIssue(&seoResponse, Issue{
			Category:    "Meta Tags",
			Severity:    "Critical",
			Title:       "Missing Meta Description",
			Description: "The page lacks a meta description tag.",
			Impact:      "Reduces click-through rates from search results and gives less control over search snippets.",
			HowToFix:    "Add a compelling meta description between 120-160 characters that includes your main keyword and a call to action.",
			Examples:    "<meta name=\"description\" content=\"Your compelling description here with keywords and call to action\">",
		})
		seoResponse.Score -= 10
	} else if len(seoResponse.Description) < 120 || len(seoResponse.Description) > 160 {
		addIssue(&seoResponse, Issue{
			Category:    "Meta Tags",
			Severity:    "Warning",
			Title:       "Meta Description Length Not Optimal",
			Description: "Meta description length is outside the recommended range.",
			Impact:      "May be truncated in search results or not fully optimized.",
			HowToFix:    "Adjust description length to be between 120-160 characters while maintaining relevance and including a call to action.",
		})
		seoResponse.Score -= 5
	}

	// Analyze meta tags
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		property, _ := s.Attr("property")
		content, _ := s.Attr("content")
		if name != "" || property != "" {
			seoResponse.MetaTags = append(seoResponse.MetaTags, MetaTag{
				Name:    name + property,
				Content: content,
			})
		}
	})

	// Analyze heading structure
	headingCounts := make(map[string]int)
	headingMap := make(map[string]HeadingStructure)
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		level := s.Get(0).Data
		headingCounts[level]++

		// Create or update heading structure
		if _, exists := headingMap[level]; !exists {
			headingMap[level] = HeadingStructure{
				Level: level,
				Text:  s.Text(),
				Count: 1,
			}
		} else {
			existing := headingMap[level]
			existing.Count++
			headingMap[level] = existing
		}
	})

	// Convert map to sorted slice
	var headings []HeadingStructure
	for _, level := range []string{"h1", "h2", "h3", "h4", "h5", "h6"} {
		if structure, exists := headingMap[level]; exists {
			headings = append(headings, structure)
		}
	}
	seoResponse.HeadingStructure = headings

	if headingCounts["h1"] == 0 {
		addIssue(&seoResponse, Issue{
			Category:    "Content Structure",
			Severity:    "Critical",
			Title:       "Missing H1 Heading",
			Description: "The page doesn't have an H1 heading.",
			Impact:      "Reduces search engines' understanding of page content hierarchy and main topic.",
			HowToFix:    "Add a single H1 heading that clearly describes the page's main topic and includes your primary keyword.",
		})
		seoResponse.Score -= 10
	} else if headingCounts["h1"] > 1 {
		addIssue(&seoResponse, Issue{
			Category:    "Content Structure",
			Severity:    "Warning",
			Title:       "Multiple H1 Headings",
			Description: "The page has multiple H1 headings.",
			Impact:      "May confuse search engines about the main topic of the page.",
			HowToFix:    "Keep only one H1 heading that best describes the page's main topic.",
		})
		seoResponse.Score -= 5
	}

	// Analyze images
	seoResponse.Images = analyzeImages(doc)
	if seoResponse.Images.WithoutAlt > 0 {
		addIssue(&seoResponse, Issue{
			Category:    "Accessibility & SEO",
			Severity:    "Warning",
			Title:       "Images Missing Alt Text",
			Description: fmt.Sprintf("%d images are missing alt text", seoResponse.Images.WithoutAlt),
			Impact:      "Reduces accessibility and image search potential.",
			HowToFix:    "Add descriptive alt text to all images that describes their content or function.",
		})
		seoResponse.Score -= 5
	}

	// Analyze links
	seoResponse.Links = analyzeLinks(doc, req.URL)
	if seoResponse.Links.BrokenLinks > 0 {
		addIssue(&seoResponse, Issue{
			Category:    "Technical SEO",
			Severity:    "Critical",
			Title:       "Broken Links Detected",
			Description: fmt.Sprintf("Found %d broken links", seoResponse.Links.BrokenLinks),
			Impact:      "Negatively impacts user experience and search engine crawling.",
			HowToFix:    "Fix or remove all broken links. Regularly monitor for broken links.",
		})
		seoResponse.Score -= 10
	}

	// Technical SEO checks
	seoResponse.Technical = analyzeTechnicalSEO(req.URL, resp, doc, &seoResponse)
	if !seoResponse.Technical.IsHTTPS {
		addIssue(&seoResponse, Issue{
			Category:    "Security & SEO",
			Severity:    "Critical",
			Title:       "Not Using HTTPS",
			Description: "The website is not served over HTTPS.",
			Impact:      "Reduces trust, security, and search engine rankings.",
			HowToFix:    "Install an SSL certificate and redirect all HTTP traffic to HTTPS.",
		})
		seoResponse.Score -= 15
	}

	// Performance metrics
	seoResponse.Performance = analyzePerformance(bodyBytes, doc)
	if seoResponse.Performance.PageSize > 5000000 { // 5MB
		addIssue(&seoResponse, Issue{
			Category:    "Performance",
			Severity:    "Warning",
			Title:       "Large Page Size",
			Description: "The page size exceeds 5MB.",
			Impact:      "Slows down page load time and affects user experience.",
			HowToFix:    "Optimize images, minify CSS/JS, and remove unnecessary resources.",
		})
		seoResponse.Score -= 5
	}

	// Content analysis
	seoResponse.Content = analyzeContent(doc)
	if seoResponse.Content.WordCount < 300 {
		addIssue(&seoResponse, Issue{
			Category:    "Content",
			Severity:    "Warning",
			Title:       "Thin Content",
			Description: "The page has less than 300 words.",
			Impact:      "May be considered thin content by search engines.",
			HowToFix:    "Add more valuable, relevant content to reach at least 300 words.",
		})
		seoResponse.Score -= 5
	}

	// Analyze social media meta tags
	seoResponse.SocialMeta = analyzeSocialMetaTags(doc)
	if len(seoResponse.SocialMeta.OpenGraph) == 0 && len(seoResponse.SocialMeta.Twitter) == 0 {
		addIssue(&seoResponse, Issue{
			Category:    "Social Media",
			Severity:    "Warning",
			Title:       "Missing Social Media Meta Tags",
			Description: "The page lacks Open Graph and Twitter Card meta tags",
			Impact:      "Reduces social media sharing effectiveness and appearance",
			HowToFix:    "Add Open Graph and Twitter Card meta tags for better social media integration",
			Examples:    "<meta property=\"og:title\" content=\"Your Title\">\n<meta name=\"twitter:card\" content=\"summary_large_image\">",
		})
		seoResponse.Score -= 5
	}

	// Check for structured data
	if len(seoResponse.Technical.SchemaMarkup) == 0 {
		addIssue(&seoResponse, Issue{
			Category:    "Technical SEO",
			Severity:    "Warning",
			Title:       "Missing Structured Data",
			Description: "No structured data (Schema.org) found on the page",
			Impact:      "Reduces rich snippet opportunities in search results",
			HowToFix:    "Implement relevant Schema.org markup for your content type",
		})
		seoResponse.Score -= 5
	}

	// Check content keyword optimization
	if len(seoResponse.Content.TopKeywords) > 0 {
		mainKeyword := seoResponse.Content.TopKeywords[0]
		if !strings.Contains(strings.ToLower(seoResponse.Title), strings.ToLower(mainKeyword)) {
			addIssue(&seoResponse, Issue{
				Category:    "Content Optimization",
				Severity:    "Warning",
				Title:       "Main Keyword Missing in Title",
				Description: fmt.Sprintf("The main keyword '%s' is not present in the title tag", mainKeyword),
				Impact:      "Reduces keyword relevance and potential rankings",
				HowToFix:    "Include the main keyword in the title tag while maintaining readability",
			})
			seoResponse.Score -= 5
		}
	}

	// Check for mobile optimization
	if !seoResponse.Technical.MobileResponsive {
		addIssue(&seoResponse, Issue{
			Category:    "Mobile Optimization",
			Severity:    "Critical",
			Title:       "Not Mobile-Friendly",
			Description: "The page may not be optimized for mobile devices",
			Impact:      "Severely impacts mobile search rankings and user experience",
			HowToFix:    "Implement responsive design and ensure proper mobile viewport configuration",
			Examples:    "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">",
		})
		seoResponse.Score -= 10
	}

	c.JSON(http.StatusOK, seoResponse)
}

func addIssue(response *SEOCheckResponse, issue Issue) {
	switch issue.Severity {
	case "Critical":
		response.CriticalIssues = append(response.CriticalIssues, issue)
	case "Warning":
		response.Warnings = append(response.Warnings, issue)
	default:
		response.Recommendations = append(response.Recommendations, issue)
	}
}

func analyzeImages(doc *goquery.Document) ImageAnalysis {
	analysis := ImageAnalysis{}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		analysis.TotalImages++
		if alt, exists := s.Attr("alt"); exists && alt != "" {
			analysis.WithAlt++
		} else {
			analysis.WithoutAlt++
		}
		if src, exists := s.Attr("src"); exists {
			analysis.ImageURLs = append(analysis.ImageURLs, src)
		}
	})
	return analysis
}

func analyzeLinks(doc *goquery.Document, baseURL string) LinkAnalysis {
	analysis := LinkAnalysis{}
	parsedBase, _ := url.Parse(baseURL)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}

		if parsedHref.Host == "" || parsedHref.Host == parsedBase.Host {
			analysis.InternalLinks++
			analysis.InternalURLs = append(analysis.InternalURLs, href)
		} else {
			analysis.ExternalLinks++
			analysis.ExternalURLs = append(analysis.ExternalURLs, href)
		}

		if rel, exists := s.Attr("rel"); exists && strings.Contains(rel, "nofollow") {
			analysis.NoFollowLinks++
		}
	})
	return analysis
}

func analyzeTechnicalSEO(baseURL string, resp *http.Response, doc *goquery.Document, seoResponse *SEOCheckResponse) TechnicalSEO {
	parsedURL, _ := url.Parse(baseURL)

	technical := TechnicalSEO{
		IsHTTPS:        parsedURL.Scheme == "https",
		ServerResponse: resp.StatusCode,
	}

	// Check for meta keywords
	if keywords, exists := doc.Find("meta[name='keywords']").Attr("content"); exists {
		technical.MetaKeywords = keywords
		technical.HasMetaKeywords = true

		// Add recommendation if keywords are found
		if len(keywords) > 0 {
			addIssue(seoResponse, Issue{
				Category:    "Meta Tags",
				Severity:    "Information",
				Title:       "Meta Keywords Tag Present",
				Description: "Meta keywords tag is present but has minimal SEO impact nowadays",
				Impact:      "Very low impact on search rankings as most modern search engines ignore this tag",
				HowToFix:    "Focus on quality content and proper keyword usage within content instead of meta keywords",
				Examples:    fmt.Sprintf("Current keywords: %s", keywords),
			})
		}
	}

	// Check for canonical
	technical.Canonical = doc.Find("link[rel='canonical']").AttrOr("href", "")

	// Check for sitemap
	sitemapURL := fmt.Sprintf("%s://%s/sitemap.xml", parsedURL.Scheme, parsedURL.Host)
	sitemapIndexURL := fmt.Sprintf("%s://%s/sitemap_index.xml", parsedURL.Scheme, parsedURL.Host)

	// Try common sitemap paths
	sitemapPaths := []string{
		sitemapURL,
		sitemapIndexURL,
		fmt.Sprintf("%s://%s/sitemap-index.xml", parsedURL.Scheme, parsedURL.Host),
		fmt.Sprintf("%s://%s/sitemap/sitemap.xml", parsedURL.Scheme, parsedURL.Host),
		fmt.Sprintf("%s://%s/sitemaps/sitemap.xml", parsedURL.Scheme, parsedURL.Host),
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, path := range sitemapPaths {
		resp, err := client.Get(path)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			// Check if it's a sitemap index
			if strings.Contains(string(bodyBytes), "<sitemapindex") {
				technical.HasSitemap = true
				break
			}

			// Check if it's a regular sitemap
			if strings.Contains(string(bodyBytes), "<urlset") {
				technical.HasSitemap = true
				break
			}
		}
	}

	// Check for robots.txt
	robotsURL := fmt.Sprintf("%s://%s/robots.txt", parsedURL.Scheme, parsedURL.Host)
	robotsResp, err := client.Get(robotsURL)
	if err == nil {
		defer robotsResp.Body.Close()
		technical.HasRobotsTxt = robotsResp.StatusCode == http.StatusOK
	}

	// Check for schema markup
	doc.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
		technical.SchemaMarkup = append(technical.SchemaMarkup, s.Text())
	})

	// Check for mobile viewport
	if _, exists := doc.Find("meta[name='viewport']").Attr("content"); exists {
		technical.MobileResponsive = true
	}

	// Check SSL
	technical.HasSSL = parsedURL.Scheme == "https" && resp.TLS != nil

	return technical
}

func analyzePerformance(bodyBytes []byte, doc *goquery.Document) PerformanceMetrics {
	metrics := PerformanceMetrics{
		PageSize: int64(len(bodyBytes)),
	}

	// Count resources
	metrics.JSCount = doc.Find("script").Length()
	metrics.CSSCount = doc.Find("link[rel='stylesheet']").Length()
	metrics.ImageCount = doc.Find("img").Length()
	metrics.ResourceCount = metrics.JSCount + metrics.CSSCount + metrics.ImageCount

	return metrics
}

// calculateReadabilityScore calculates the Flesch-Kincaid readability score
func calculateReadabilityScore(text string) float64 {
	// Split text into sentences (improved to handle common abbreviations)
	sentences := strings.FieldsFunc(text, func(r rune) bool {
		return r == '.' || r == '?' || r == '!'
	})

	// Filter out empty sentences and common abbreviations
	var validSentences []string
	for _, s := range sentences {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) > 0 && !strings.Contains(trimmed, "Mr") && !strings.Contains(trimmed, "Mrs") &&
			!strings.Contains(trimmed, "Dr") && !strings.Contains(trimmed, "Ph.D") {
			validSentences = append(validSentences, trimmed)
		}
	}

	sentenceCount := float64(len(validSentences))

	// Count words and syllables
	words := strings.Fields(text)
	wordCount := float64(len(words))

	// Count syllables (improved approach)
	syllableCount := float64(0)
	for _, word := range words {
		syllableCount += float64(countSyllables(strings.ToLower(word)))
	}

	// Avoid division by zero
	if wordCount == 0 || sentenceCount == 0 {
		return 0
	}

	// Modified Flesch Reading Ease formula with adjusted weights
	// Original: 206.835 - 1.015(words/sentences) - 84.6(syllables/words)
	// Adjusted: 206.835 - 0.846(words/sentences) - 74.3(syllables/words)
	score := 206.835 - 0.846*(wordCount/sentenceCount) - 74.3*(syllableCount/wordCount)

	// Clamp score between 0 and 100
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	// Round to 2 decimal places
	return math.Round(score*100) / 100
}

// countSyllables counts the number of syllables in a word (simplified approach)
func countSyllables(word string) int {
	count := 0
	vowels := "aeiouy"
	prevIsVowel := false

	// Handle special cases
	if len(word) <= 3 {
		return 1
	}

	// Remove 'e' from the end if it exists
	if strings.HasSuffix(word, "e") {
		word = word[:len(word)-1]
	}

	for _, char := range word {
		isVowel := strings.ContainsRune(vowels, char)
		if isVowel && !prevIsVowel {
			count++
		}
		prevIsVowel = isVowel
	}

	if count == 0 {
		count = 1
	}

	return count
}

func analyzeContent(doc *goquery.Document) ContentAnalysis {
	content := ContentAnalysis{
		KeywordDensity: make(map[string]float64),
	}

	text := doc.Find("body").Text()
	words := strings.Fields(text)
	content.WordCount = len(words)
	content.CharacterCount = len(text)
	content.ParagraphCount = doc.Find("p").Length()

	// Calculate readability score
	content.ReadabilityScore = calculateReadabilityScore(text)

	// Simple keyword analysis
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,!?()[]{}\"'"))
		if len(word) > 3 { // Skip short words
			wordCount[word]++
		}
	}

	// Calculate keyword density
	for word, count := range wordCount {
		density := float64(count) / float64(content.WordCount) * 100
		content.KeywordDensity[word] = density
	}

	// Get top keywords
	type wordFreq struct {
		word  string
		count int
	}
	var frequencies []wordFreq
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFreq{word, count})
	}
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].count > frequencies[j].count
	})

	// Get top 10 keywords
	for i := 0; i < 10 && i < len(frequencies); i++ {
		content.TopKeywords = append(content.TopKeywords, frequencies[i].word)
	}

	return content
}

func analyzeSocialMetaTags(doc *goquery.Document) SocialMetaTags {
	social := SocialMetaTags{
		OpenGraph: make(map[string]string),
		Twitter:   make(map[string]string),
		LinkedIn:  make(map[string]string),
	}

	// Open Graph tags
	doc.Find("meta[property^='og:']").Each(func(i int, s *goquery.Selection) {
		if prop, exists := s.Attr("property"); exists {
			if content, exists := s.Attr("content"); exists {
				social.OpenGraph[prop] = content
			}
		}
	})

	// Twitter Card tags
	doc.Find("meta[name^='twitter:']").Each(func(i int, s *goquery.Selection) {
		if name, exists := s.Attr("name"); exists {
			if content, exists := s.Attr("content"); exists {
				social.Twitter[name] = content
			}
		}
	})

	// LinkedIn tags
	doc.Find("meta[property^='linkedin:']").Each(func(i int, s *goquery.Selection) {
		if prop, exists := s.Attr("property"); exists {
			if content, exists := s.Attr("content"); exists {
				social.LinkedIn[prop] = content
			}
		}
	})

	return social
}
