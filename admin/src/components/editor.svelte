<!-- https://css-tricks.com/creating-an-editable-textarea-that-supports-syntax-highlighted-code/ -->
<script lang="ts" context="module">
	import Prism from "prismjs";
</script>

<script lang="ts">
	let editor: HTMLTextAreaElement;
	let content: HTMLSpanElement;
	let highlighter: HTMLPreElement;
	let value = "";

	function scroll() {
		if (highlighter && editor) {
			highlighter.scrollTop = editor.scrollTop;
			highlighter.scrollLeft = editor.scrollLeft;
		}
	}
	function update() {
		if (content) {
			if (value[value.length - 1] == "\n") {
				value += " ";
			}
			content.innerHTML = value
				.replace(new RegExp("&", "g"), "&")
				.replace(new RegExp("<", "g"), "<"); /* Global RegExp */
			Prism.highlightElement(content);
			scroll();
		}
	}

	// xpos and ypos are the width and height of the textarea elem. make them negative and add any margins / padding
	// to get the value for the pre elements relative position in order to place them on top of eachother.
</script>

<div class="container">
	<textarea
		id="editor"
		bind:value
		on:input={update}
		on:scroll={scroll}
		spellcheck="false"
		bind:this={editor}
	/>
	<pre id="highlighter" style:position="relative" bind:this={highlighter}>
			<code class="language-javascript" id="content" bind:this={content} />
		</pre>
</div>

<style>
	@import "prismjs/themes/prism.css";

	.container {
		flex: 1;
		display: grid;
	}
	/* sizing, border, etc. */
	#editor,
	#highlighter {
		grid-column: 1;
		grid-row: 1;
		margin: 0.5rem 0 0 0;
		padding: 0.5rem;
	}
	/* typography */
	#editor,
	#highlighter,
	#highlighter * {
		font-size: large;
		font-family: monospace;
		line-height: 20pt;
	}
	#highlighter {
		z-index: 0;
	}
	#editor {
		z-index: 1;
	}

	#editor {
		color: transparent;
		background: transparent;
		caret-color: white;
	}

	#editor,
	#highlighter {
		overflow: auto;
		white-space: nowrap;
	}

	#editor {
		resize: none;
	}

	p code {
		border-radius: 2px;
		background-color: #eee;
		color: #111;
	}

	/* position highlighter and editor on top of eachother */

	/* move textare in front of highlighter */
</style>
