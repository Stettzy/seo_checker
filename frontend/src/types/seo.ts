export interface MetaTag {
    name: string;
    content: string;
}

export interface HeadingStructure {
    level: string;
    text: string;
    count: number;
}

export interface ImageAnalysis {
    totalImages: number;
    withAlt: number;
    withoutAlt: number;
    largeImages: number;
    optimizedCount: number;
    imageURLs: string[];
}

export interface LinkAnalysis {
    internalLinks: number;
    externalLinks: number;
    brokenLinks: number;
    noFollowLinks: number;
    internalURLs: string[];
    externalURLs: string[];
}

export interface PerformanceMetrics {
    pageSize: number;
    loadTime: number;
    resourceCount: number;
    jsCount: number;
    cssCount: number;
    imageCount: number;
}

export interface ContentAnalysis {
    wordCount: number;
    characterCount: number;
    paragraphCount: number;
    readabilityScore: number;
    keywordDensity: Record<string, number>;
    topKeywords: string[];
}

export interface TechnicalSEO {
    hasSitemap: boolean;
    hasRobotsTxt: boolean;
    isHttps: boolean;
    hasSSL: boolean;
    serverResponse: number;
    canonical: string;
    mobileResponsive: boolean;
    ampEnabled: boolean;
    schemaMarkup: string[];
    metaKeywords: string;
    hasMetaKeywords: boolean;
}

export interface Issue {
    category: string;
    severity: string;
    title: string;
    description: string;
    impact: string;
    howToFix: string;
    examples?: string;
}

export interface SocialMetaTags {
    openGraph: Record<string, string>;
    twitter: Record<string, string>;
    linkedIn: Record<string, string>;
}

export interface SEOCheckResponse {
    url: string;
    score: number;
    title: string;
    titleLength: number;
    description: string;
    descLength: number;
    metaTags: MetaTag[];
    headingStructure: HeadingStructure[];
    images: ImageAnalysis;
    links: LinkAnalysis;
    performance: PerformanceMetrics;
    content: ContentAnalysis;
    technical: TechnicalSEO;
    criticalIssues: Issue[];
    warnings: Issue[];
    recommendations: Issue[];
    lastUpdated: string;
    loadTime: number;
    socialMeta: SocialMetaTags;
    competitors: string[];
} 