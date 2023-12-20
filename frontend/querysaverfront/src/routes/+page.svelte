<script lang="ts">
  import ActionBtn from "../components/action-btn.svelte";
  import SearchBar from "../components/search-bar.svelte";
  import SearchState from "../components/search-state.svelte";
  import TagsList from "../components/tags-list.svelte";
</script>

<div class="flex min-h-screen">
  <div class="w-1/4 bg-surface-800">
    <div><ActionBtn /></div>
    <div class="border-2 p-2">
      <h3>Search Tags:</h3>
      <div class="mt-2"><SearchBar /></div>
      <div class="mt-2"><TagsList /></div>
    </div>

    <div>CORESPONDING FILE</div>
  </div>
  <div class="w-3/4">
    <div class="h-1/6"><SearchState /></div>
    <div class="h-5/6">PREVIEW FILE</div>
  </div>
</div>
