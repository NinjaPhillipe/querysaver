<script lang="ts">
  import { SlideToggle } from "@skeletonlabs/skeleton";
  import type { Tag } from "../lib/dataType";
  import { selectedTags } from "../lib/store";

  // type of query
  let orQuery = true;

  function removeTag(tagToRemove: Tag) {
    selectedTags.update((tags) =>
      tags.filter((tag) => tag.id !== tagToRemove.id)
    );
  }
</script>

<div class="flex">
  <div>
    {#each $selectedTags as tag (tag)}
      <div class="badge" style="background-color: {tag.color};">
        {tag.name}
        <button
          type="button"
          class="btn btn-sm variant-filled w-2 h-6 ml-2"
          on:click={() => removeTag(tag)}>x</button
        >
      </div>
    {/each}
  </div>
  <div class="flex flex-col">
    <SlideToggle name="queryType" checked={orQuery}>OR/AND</SlideToggle>
  </div>
</div>
