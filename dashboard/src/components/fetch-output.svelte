<script lang="ts">
  import { onMount } from "svelte";
  let posts = [];
  async function listPosts() {
    let res = await fetch("https://jsonplaceholder.typicode.com/posts");
    let json = await res.json();
    posts = json.map((post) => {
      return {
        id: post.id,
        title: post.title,
        body: post.body,
      };
    });
  }

  onMount(async () => {
    await listPosts();
    console.log(posts);
  });
</script>

<div class="output">
  <pre>
    {JSON.stringify(posts, null, 2)}
  </pre>
</div>

<style>
  div {
    overflow-y: scroll;
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    max-height: 100%;
  }
  pre {
    text-overflow: ellipsis;
  }
</style>
