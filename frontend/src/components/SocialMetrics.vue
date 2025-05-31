<template>
  <div class="bg-white rounded-lg shadow-sm p-6">
    <h3 class="text-lg font-medium text-gray-900 mb-4">Social Media Integration</h3>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- Open Graph -->
      <div>
        <h4 class="text-sm font-medium text-gray-700 mb-2">
          <div class="flex items-center">
            <svg class="h-5 w-5 text-blue-500 mr-2" viewBox="0 0 24 24" fill="currentColor">
              <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
            </svg>
            Open Graph
          </div>
        </h4>
        <div class="space-y-2">
          <div v-for="(value, key) in socialMeta.openGraph" :key="key" class="bg-gray-50 px-3 py-2 rounded-lg">
            <dt class="text-xs font-medium text-gray-500">{{ formatMetaKey(key) }}</dt>
            <dd class="mt-1 text-sm text-gray-900 truncate">{{ value }}</dd>
          </div>
          <div v-if="Object.keys(socialMeta.openGraph).length === 0" class="text-sm text-gray-500 italic">
            No Open Graph tags found
          </div>
        </div>
      </div>

      <!-- Twitter -->
      <div>
        <h4 class="text-sm font-medium text-gray-700 mb-2">
          <div class="flex items-center">
            <svg class="h-5 w-5 text-blue-400 mr-2" viewBox="0 0 24 24" fill="currentColor">
              <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
            </svg>
            Twitter Cards
          </div>
        </h4>
        <div class="space-y-2">
          <div v-for="(value, key) in socialMeta.twitter" :key="key" class="bg-gray-50 px-3 py-2 rounded-lg">
            <dt class="text-xs font-medium text-gray-500">{{ formatMetaKey(key) }}</dt>
            <dd class="mt-1 text-sm text-gray-900 truncate">{{ value }}</dd>
          </div>
          <div v-if="Object.keys(socialMeta.twitter).length === 0" class="text-sm text-gray-500 italic">
            No Twitter Card tags found
          </div>
        </div>
      </div>

      <!-- LinkedIn -->
      <div>
        <h4 class="text-sm font-medium text-gray-700 mb-2">
          <div class="flex items-center">
            <svg class="h-5 w-5 text-blue-600 mr-2" viewBox="0 0 24 24" fill="currentColor">
              <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
            </svg>
            LinkedIn
          </div>
        </h4>
        <div class="space-y-2">
          <div v-for="(value, key) in socialMeta.linkedIn" :key="key" class="bg-gray-50 px-3 py-2 rounded-lg">
            <dt class="text-xs font-medium text-gray-500">{{ formatMetaKey(key) }}</dt>
            <dd class="mt-1 text-sm text-gray-900 truncate">{{ value }}</dd>
          </div>
          <div v-if="Object.keys(socialMeta.linkedIn).length === 0" class="text-sm text-gray-500 italic">
            No LinkedIn tags found
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  socialMeta: {
    openGraph: Record<string, string>;
    twitter: Record<string, string>;
    linkedIn: Record<string, string>;
  };
}

const props = defineProps<Props>();

const formatMetaKey = (key: string) => {
  return key
    .split(/[:.]/g)
    .pop()
    ?.split(/(?=[A-Z])/)
    .join(' ')
    .toLowerCase()
    .replace(/^\w/, c => c.toUpperCase()) || key;
};
</script> 