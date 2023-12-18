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
      <button
        type="button"
        class="btn btn-sm variant-filled w-2 h-6 ml-2"
        on:click={() => removeTag(tag)}>x</button
      >
    </div>
  {/each}
</div>
