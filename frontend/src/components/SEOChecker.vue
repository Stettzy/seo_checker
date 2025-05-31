<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <!-- Hero Section -->
    <div
      class="w-full bg-gradient-to-br from-indigo-600 via-blue-600 to-blue-500 min-h-[400px] flex items-center justify-center relative overflow-hidden">
      <div
        class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZGVmcz48cGF0dGVybiBpZD0iZ3JpZCIgd2lkdGg9IjQwIiBoZWlnaHQ9IjQwIiBwYXR0ZXJuVW5pdHM9InVzZXJTcGFjZU9uVXNlIj48cGF0aCBkPSJNIDQwIDAgTCAwIDAgMCA0MCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSJ3aGl0ZSIgc3Ryb2tlLW9wYWNpdHk9IjAuMSIgc3Ryb2tlLXdpZHRoPSIxIi8+PC9wYXR0ZXJuPjwvZGVmcz48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2dyaWQpIi8+PC9zdmc+')] opacity-50">
      </div>
      <div class="max-w-5xl mx-auto w-full px-4 sm:px-6 lg:px-8 py-16 text-center relative z-10">
        <div class="flex justify-center mb-8">
          <SEOLogo class="text-white transform hover:scale-110 transition-transform duration-200" />
        </div>
        <h1 class="text-4xl md:text-5xl font-bold text-white mb-4 tracking-tight">
          Professional SEO Analysis
        </h1>
        <p class="text-lg text-white/90 mb-8 max-w-2xl mx-auto">
          Get a comprehensive analysis of your website's SEO performance and actionable recommendations to improve your
          rankings
        </p>

        <!-- URL Input Form -->
        <div class="max-w-3xl mx-auto">
          <form @submit.prevent="analyzeSEO" class="bg-white/10 backdrop-blur-lg rounded-xl p-4">
            <div class="flex flex-col md:flex-row gap-3">
              <div class="flex-1">
                <input type="url" name="url" id="url" v-model="url"
                  class="w-full px-4 py-2 bg-white/90 backdrop-blur text-gray-900 placeholder-gray-500 rounded-lg border-2 border-white/20 focus:border-white/40 focus:ring-2 focus:ring-white/20 focus:outline-none text-base transition-all duration-200"
                  placeholder="Enter your website URL (e.g., https://example.com)" required />
              </div>
              <button type="submit" :disabled="loading"
                class="px-6 py-2 bg-indigo-500 hover:bg-indigo-600 text-white font-medium rounded-lg shadow-lg hover:shadow-xl transition-all duration-200 disabled:opacity-50 text-base whitespace-nowrap flex items-center justify-center min-w-[140px]">
                <svg v-if="loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
                  xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
                  </path>
                </svg>
                {{ loading ? 'Analyzing...' : 'Analyze SEO' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Results Section -->
    <div class="w-full bg-gray-50" v-if="results">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <!-- Overall Score Card -->
        <div class="bg-white rounded-xl shadow p-6 mb-8 transform -translate-y-16">
          <div class="flex flex-col md:flex-row items-center justify-between gap-6">
            <div>
              <h2 class="text-2xl font-bold text-gray-900 mb-1">SEO Score</h2>
              <p class="text-base text-gray-600">Based on {{ (results.criticalIssues?.length || 0) +
                (results.warnings?.length || 0) }} checks</p>
            </div>
            <div class="flex items-center gap-3">
              <div class="text-6xl font-bold" :class="{
                'text-green-500': results.score >= 90,
                'text-yellow-500': results.score >= 70 && results.score < 90,
                'text-red-500': results.score < 70
              }">
                {{ results.score }}
              </div>
              <div class="text-2xl text-gray-400">/100</div>
            </div>
          </div>
        </div>

        <!-- Quick Stats -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          <div class="bg-white rounded-xl shadow-lg p-6 hover:shadow-xl transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-600">Page Load Time</p>
                <p class="text-2xl font-bold mt-2">{{ (results.loadTime || 0).toFixed(2) }}s</p>
              </div>
              <div :class="{
                'text-green-500': results.loadTime < 2,
                'text-yellow-500': results.loadTime >= 2 && results.loadTime < 3,
                'text-red-500': results.loadTime >= 3
              }">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-xl shadow p-6 hover:shadow-lg transition-shadow">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Critical Issues</p>
                <p class="text-2xl font-bold mt-1 text-red-500">{{ results.criticalIssues?.length || 0 }}</p>
              </div>
              <div class="text-red-500">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-xl shadow p-6 hover:shadow-lg transition-shadow">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Warnings</p>
                <p class="text-2xl font-bold mt-1 text-yellow-500">{{ results.warnings?.length || 0 }}</p>
              </div>
              <div class="text-yellow-500">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-xl shadow p-6 hover:shadow-lg transition-shadow">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Mobile Friendly</p>
                <p class="text-2xl font-bold mt-1"
                  :class="results.technical?.mobileResponsive ? 'text-green-500' : 'text-red-500'">
                  {{ results.technical?.mobileResponsive ? 'Yes' : 'No' }}
                </p>
              </div>
              <div :class="results.technical?.mobileResponsive ? 'text-green-500' : 'text-red-500'">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
                </svg>
              </div>
            </div>
          </div>
        </div>

        <!-- Issues Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
          <!-- Critical Issues -->
          <div v-if="results.criticalIssues?.length > 0" class="bg-white rounded-xl shadow p-6">
            <div class="flex items-center gap-2 mb-4">
              <div class="text-red-500">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <h3 class="text-lg font-bold text-gray-900">Critical Issues</h3>
            </div>
            <div class="space-y-3">
              <div v-for="(issue, index) in results.criticalIssues" :key="'critical-' + index"
                class="border-l-4 border-red-500 bg-red-50 p-4 rounded-r-lg">
                <h4 class="text-base font-medium text-red-800 mb-2">{{ issue.title }}</h4>
                <p class="text-sm text-red-700 mb-3">{{ issue.description }}</p>
                <div class="bg-white rounded-lg p-3 space-y-2 text-sm">
                  <p><span class="font-medium">Impact:</span> {{ issue.impact }}</p>
                  <p><span class="font-medium">How to Fix:</span> {{ issue.howToFix }}</p>
                  <p v-if="issue.examples" class="font-mono text-xs bg-red-50 p-2 rounded mt-2">{{ issue.examples }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Warnings -->
          <div v-if="results.warnings?.length > 0" class="bg-white rounded-xl shadow p-6">
            <div class="flex items-center gap-2 mb-4">
              <div class="text-yellow-500">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <h3 class="text-lg font-bold text-gray-900">Warnings</h3>
            </div>
            <div class="space-y-3">
              <div v-for="(warning, index) in results.warnings" :key="'warning-' + index"
                class="border-l-4 border-yellow-500 bg-yellow-50 p-4 rounded-r-lg">
                <h4 class="text-base font-medium text-yellow-800 mb-2">{{ warning.title }}</h4>
                <p class="text-sm text-yellow-700 mb-3">{{ warning.description }}</p>
                <div class="bg-white rounded-lg p-3 space-y-2 text-sm">
                  <p><span class="font-medium">Impact:</span> {{ warning.impact }}</p>
                  <p><span class="font-medium">How to Fix:</span> {{ warning.howToFix }}</p>
                  <p v-if="warning.examples" class="font-mono text-xs bg-yellow-50 p-2 rounded mt-2">{{ warning.examples
                    }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Meta Information -->
        <div class="bg-white rounded-xl shadow p-6 mb-8">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Meta Information</h3>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Title Tag</h4>
              <div class="bg-gray-50 rounded-lg p-4">
                <p class="text-sm text-gray-900 font-mono break-all">{{ results.title }}</p>
                <div class="mt-3 flex items-center gap-2">
                  <div class="h-2 flex-1 rounded-full bg-gray-200 overflow-hidden">
                    <div class="h-full rounded-full" :class="{
                      'bg-green-500': results.titleLength >= 30 && results.titleLength <= 60,
                      'bg-yellow-500': (results.titleLength > 60 && results.titleLength <= 70) || (results.titleLength >= 20 && results.titleLength < 30),
                      'bg-red-500': results.titleLength > 70 || results.titleLength < 20
                    }" :style="{ width: Math.min(results.titleLength / 0.7, 100) + '%' }"></div>
                  </div>
                  <span class="text-xs font-medium" :class="{
                    'text-green-600': results.titleLength >= 30 && results.titleLength <= 60,
                    'text-yellow-600': (results.titleLength > 60 && results.titleLength <= 70) || (results.titleLength >= 20 && results.titleLength < 30),
                    'text-red-600': results.titleLength > 70 || results.titleLength < 20
                  }">{{ results.titleLength }} chars</span>
                </div>
              </div>
            </div>

            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Meta Description</h4>
              <div class="bg-gray-50 rounded-lg p-4">
                <p class="text-sm text-gray-900 font-mono break-all">{{ results.description }}</p>
                <div class="mt-3 flex items-center gap-2">
                  <div class="h-2 flex-1 rounded-full bg-gray-200 overflow-hidden">
                    <div class="h-full rounded-full" :class="{
                      'bg-green-500': results.descLength >= 120 && results.descLength <= 160,
                      'bg-yellow-500': (results.descLength > 160 && results.descLength <= 180) || (results.descLength >= 80 && results.descLength < 120),
                      'bg-red-500': results.descLength > 180 || results.descLength < 80
                    }" :style="{ width: Math.min(results.descLength / 1.8, 100) + '%' }"></div>
                  </div>
                  <span class="text-xs font-medium" :class="{
                    'text-green-600': results.descLength >= 120 && results.descLength <= 160,
                    'text-yellow-600': (results.descLength > 160 && results.descLength <= 180) || (results.descLength >= 80 && results.descLength < 120),
                    'text-red-600': results.descLength > 180 || results.descLength < 80
                  }">{{ results.descLength }} chars</span>
                </div>
              </div>
            </div>

            <div class="mt-6">
              <h4 class="text-base font-medium text-gray-700 mb-3">Meta Keywords</h4>
              <div class="bg-gray-50 rounded-lg p-4">
                <p v-if="results.technical?.hasMetaKeywords" class="text-sm text-gray-900 font-mono break-all">
                  {{ results.technical.metaKeywords }}
                </p>
                <p v-else class="text-sm text-gray-500 italic">
                  No meta keywords defined (not critical for modern SEO)
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Website Structure -->
        <div class="bg-white rounded-xl shadow p-6 mb-8">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Website Structure</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Sitemap Analysis -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Sitemap Analysis</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Sitemap Present</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical.hasSitemap ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                      {{ results.technical.hasSitemap ? 'Yes' : 'Missing' }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Robots.txt</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical.hasRobotsTxt ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                      {{ results.technical.hasRobotsTxt ? 'Present' : 'Missing' }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Canonical URL</span>
                    <span class="text-xs font-medium text-gray-900 max-w-[200px] truncate"
                      :title="results.technical.canonical">
                      {{ results.technical.canonical || 'Not Set' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- URL Structure Score -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">URL Structure</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">URL Length</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium" :class="{
                      'bg-green-100 text-green-800': results.url.length <= 100,
                      'bg-yellow-100 text-yellow-800': results.url.length > 100 && results.url.length <= 150,
                      'bg-red-100 text-red-800': results.url.length > 150
                    }">
                      {{ results.url.length }} chars
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">URL Format</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="!results.url.includes('?') && !results.url.includes('#') ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                      {{ !results.url.includes('?') && !results.url.includes('#') ? 'Clean' : 'Has Parameters' }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">SSL/HTTPS</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical.hasSSL ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                      {{ results.technical.hasSSL ? 'Secure' : 'Not Secure' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Social Media Integration -->
        <SocialMetrics :social-meta="results.socialMeta" class="mb-8" />

        <!-- Performance Metrics -->
        <div class="bg-white rounded-xl shadow p-6 mb-8">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Performance Metrics</h3>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Resource Usage</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">JavaScript Files</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.performance.jsCount }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">CSS Files</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.performance.cssCount }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Images</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.performance.imageCount }}</span>
                  </div>
                </div>
              </div>
            </div>

            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Page Size</h4>
              <div class="bg-gray-50 rounded-lg p-4 text-center">
                <div class="text-2xl font-bold text-gray-900 mb-1">{{ formatBytes(results.performance.pageSize) }}</div>
                <p class="text-xs" :class="{
                  'text-green-600': results.performance.pageSize < 1500000,
                  'text-yellow-600': results.performance.pageSize >= 1500000 && results.performance.pageSize < 3000000,
                  'text-red-600': results.performance.pageSize >= 3000000
                }">{{ results.performance.pageSize < 1500000 ? 'Optimized' : results.performance.pageSize < 3000000
                  ? 'Could be improved' : 'Too large' }}</p>
              </div>
            </div>

            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Technical Checks</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">HTTPS</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical.isHttps ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                      {{ results.technical.isHttps ? 'Secure' : 'Not Secure' }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Mobile Friendly</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical.mobileResponsive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                      {{ results.technical.mobileResponsive ? 'Yes' : 'No' }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Schema Markup</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.technical?.schemaMarkup?.length > 0 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                      {{ results.technical?.schemaMarkup?.length > 0 ? 'Present' : 'Missing' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Content Analysis -->
        <div class="bg-white rounded-xl shadow p-6 mb-8 mt-12">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Content Analysis</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <!-- Content Stats -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Content Statistics</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Words</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.content.wordCount }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Paragraphs</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.content.paragraphCount }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Readability Score</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium" :class="{
                      'bg-green-100 text-green-800': results.content?.readabilityScore >= 80,
                      'bg-yellow-100 text-yellow-800': results.content?.readabilityScore >= 60 && results.content?.readabilityScore < 80,
                      'bg-red-100 text-red-800': !results.content?.readabilityScore || results.content.readabilityScore < 60
                    }">
                      {{ results.content?.readabilityScore ? `${results.content.readabilityScore}/100` : 'Not Available'
                      }}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Top Keywords -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Top Keywords</h4>
              <div class="bg-gray-50 rounded-lg p-4">
                <div class="space-y-2">
                  <div v-for="keyword in results.content?.topKeywords" :key="keyword"
                    class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">{{ keyword }}</span>
                    <span class="text-xs font-medium text-gray-900">
                      {{ (results.content?.keywordDensity[keyword] * 100).toFixed(1) }}%
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Heading Structure -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Heading Structure</h4>
              <div class="bg-gray-50 rounded-lg p-4">
                <div class="space-y-2">
                  <div v-for="heading in results.headingStructure" :key="heading.level"
                    class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">{{ heading.level }}</span>
                    <div class="flex items-center gap-2">
                      <span class="text-xs font-medium text-gray-900">{{ heading.count }}</span>
                      <span v-if="heading.level === 'h1' && heading.count > 1" class="text-xs text-red-600">(Too many
                        H1s)</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Link Analysis -->
        <div class="bg-white rounded-xl shadow p-6 mb-8">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Link Analysis</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Link Stats -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Link Statistics</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Internal Links</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.links.internalLinks }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">External Links</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.links.externalLinks }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Broken Links</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.links.brokenLinks === 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                      {{ results.links.brokenLinks }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">NoFollow Links</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.links.noFollowLinks }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Image Analysis -->
            <div>
              <h4 class="text-base font-medium text-gray-700 mb-3">Image Optimization</h4>
              <div class="space-y-3">
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Total Images</span>
                    <span class="text-base font-semibold text-gray-900">{{ results.images.totalImages }}</span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Missing Alt Text</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.images.withoutAlt === 0 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                      {{ results.images.withoutAlt }}
                    </span>
                  </div>
                </div>
                <div class="bg-gray-50 rounded-lg p-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">Large Images</span>
                    <span class="px-2 py-1 rounded-full text-xs font-medium"
                      :class="results.images.largeImages === 0 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                      {{ results.images.largeImages }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Recommendations -->
        <div v-if="results.recommendations?.length > 0" class="bg-white rounded-xl shadow p-6 mb-8">
          <div class="flex items-center gap-2 mb-4">
            <div class="text-blue-500">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="text-lg font-bold text-gray-900">Recommendations</h3>
          </div>
          <div class="space-y-3">
            <div v-for="(rec, index) in results.recommendations" :key="'rec-' + index"
              class="border-l-4 border-blue-500 bg-blue-50 p-4 rounded-r-lg">
              <h4 class="text-base font-medium text-blue-800 mb-2">{{ rec.title }}</h4>
              <p class="text-sm text-blue-700 mb-3">{{ rec.description }}</p>
              <div class="bg-white rounded-lg p-3 space-y-2 text-sm">
                <p><span class="font-medium">Impact:</span> {{ rec.impact }}</p>
                <p><span class="font-medium">How to Fix:</span> {{ rec.howToFix }}</p>
                <p v-if="rec.examples" class="font-mono text-xs bg-blue-50 p-2 rounded mt-2">{{ rec.examples }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Competitor Analysis -->
        <div v-if="results.competitors?.length > 0" class="bg-white rounded-xl shadow p-6">
          <h3 class="text-lg font-bold text-gray-900 mb-4">Competitor Analysis</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div v-for="competitor in results.competitors" :key="competitor" class="bg-gray-50 rounded-lg p-4">
              <a :href="competitor" target="_blank" rel="noopener noreferrer"
                class="text-sm text-blue-600 hover:text-blue-800 break-all">
                {{ competitor }}
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <footer class="mt-auto bg-gray-900 text-white py-8">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-col md:flex-row items-center justify-between gap-4">
          <div class="flex items-center gap-2">
            <SEOLogo class="text-white transform hover:scale-110 transition-transform duration-200" />
          </div>
          <div class="flex items-center gap-4">
            <a href="https://github.com/yourusername/seo-checker" target="_blank" rel="noopener noreferrer"
              class="flex items-center gap-2 text-gray-300 hover:text-white transition-colors">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd"
                  d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                  clip-rule="evenodd" />
              </svg>
              <span>View on GitHub</span>
            </a>
            <span class="text-gray-500">|</span>
            <span class="text-gray-400">Made with ❤️ by Stettzy</span>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import SocialMetrics from './SocialMetrics.vue'
import SEOLogo from './SEOLogo.vue'
import type { SEOCheckResponse } from '../types/seo'

const url = ref('')
const loading = ref(false)
const results = ref<SEOCheckResponse | null>(null)

const analyzeSEO = async () => {
  loading.value = true
  results.value = null // Reset results before new analysis
  try {
    const response = await axios.post<SEOCheckResponse>('http://localhost:3000/analyze', {
      url: url.value
    })
    results.value = response.data
    console.log('SEO Analysis Results:', response.data) // Add logging for debugging
  } catch (error) {
    console.error('Error analyzing SEO:', error)
    alert('Error analyzing the website. Please try again.')
  } finally {
    loading.value = false
  }
}

const formatBytes = (bytes: number) => {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`
}
</script>
