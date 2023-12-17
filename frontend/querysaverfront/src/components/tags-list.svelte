<script lang="ts">
  import { onMount } from "svelte";

  // Define the type for a tag
  interface Tag {
    id: number;
    name: string;
    color: string;
  }

  let tags: Tag[] = [];

  onMount(async () => {
    const response = await fetch("http://localhost:8080/tags/query/all");
    tags = await response.json();
    console.log("tags");
    console.log(tags);
  });
</script>

<div class="tag-list">
  <h2>TAGS LIST</h2>
  {#each tags as tag (tag)}
    <span class="badge" style="background-color: {tag.color};">{tag.name}</span>
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
