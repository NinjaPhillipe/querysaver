<script lang="ts">
  import type { Tag } from "../lib/dataType";
  import { selectedTags } from "../lib/store";
  import { searchResultsTags } from "../lib/store";

  function updateSelectedTagsList(tag: Tag) {
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
  {#each $searchResultsTags as tag (tag)}
    <button
      class="badge"
      on:click={() => updateSelectedTagsList(tag)}
      style="background-color: {tag.color};"
      >{tag.name}
    </button>
  {/each}
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
