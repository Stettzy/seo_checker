<template>
    <div class="flex items-center gap-2">
        <svg class="w-12 h-12" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
            <!-- Magnifying Glass -->
            <circle cx="40" cy="40" r="20" fill="none" stroke="currentColor" stroke-width="8" class="animate-pulse" />
            <line x1="55" y1="55" x2="80" y2="80" stroke="currentColor" stroke-width="8" stroke-linecap="round" />

            <!-- Fun Eyes -->
            <circle cx="35" cy="35" r="5" fill="currentColor" class="animate-bounce" />
            <circle cx="45" cy="35" r="5" fill="currentColor" class="animate-bounce" />

            <!-- Smile -->
            <path d="M30 45 Q40 55 50 45" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" />

            <!-- SEO Text -->
            <text x="20" y="95" font-size="20" fill="currentColor" font-weight="bold">SEO</text>
        </svg>
        <span class="text-xl font-bold">SEOBuddy</span>
    </div>
</template>

<script setup lang="ts">
// The logo is fully contained in the template
</script>