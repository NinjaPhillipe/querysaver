<script lang="ts">
  import { onMount } from "svelte";
  import { writable } from "svelte/store";

  // Define the type for a tag
  interface Tag {
    id: number;
    name: string;
    color: string;
  }

  let tags: Tag[] = [];

  // Create a store for the selected tag
  export const selectedTags = writable<Tag[]>([]);

  onMount(async () => {
    const response = await fetch("http://localhost:8080/tags/query/all");
    tags = await response.json();
    console.log("tags");
    console.log(tags);
  });

  function updateSelectedTagsList(tag: Tag) {
    console.log("selectedTag");
    console.log(tag);

    selectedTags.update((value) => {
      // Check if the tag is already in the store
      if (!value.some((t) => t.id === tag.id)) {
        // If it's not in the store, add it
        return [...value, tag];
      }
      // If it's already in the store, return the current value
      return value;
    });
  }

  // onclick;
</script>

<div class="tag-list border-2 flex flex-wrap">
  {#each tags as tag (tag)}
    <span
      class="badge"
      on:click={() => updateSelectedTagsList(tag)}
      style="background-color: {tag.color};">{tag.name}</span
    >
  {/each}
</div>
<div>
  <p>Selected tags: {$selectedTags.map((tag) => tag.name).join(", ")}</p>
</div>

<style>
  /* .tag-list {
    width: 100%;
    height: 200px; // adjust as needed
    overflow-y: auto;
    border: 1px solid #ccc;
    padding: 10px;
  }
  .tag {
    display: inline-block;
    margin: 5px;
    padding: 5px 10px;
    background-color: #f0f0f0;
    border-radius: 5px;
  } */
</style>
