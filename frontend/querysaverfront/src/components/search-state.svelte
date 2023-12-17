<script lang="ts">
  import type { Tag } from "../lib/dataType";
  import { selectedTags } from "../lib/store";

  function removeTag(tagToRemove: Tag) {
    selectedTags.update((tags) =>
      tags.filter((tag) => tag.id !== tagToRemove.id)
    );
  }
</script>

<div>
  {#each $selectedTags as tag (tag)}
    <div class="badge" style="background-color: {tag.color};">
      {tag.name}
      <button class="border-2 px-1" on:click={() => removeTag(tag)}> x </button>
    </div>
  {/each}
</div>
