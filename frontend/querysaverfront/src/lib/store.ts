import { writable } from 'svelte/store';
import type { Tag } from './dataType';
// Create a store for the selected tag
export const selectedTags = writable<Tag[]>([]);
