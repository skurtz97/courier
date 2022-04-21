<!-- https://css-tricks.com/creating-an-editable-textarea-that-supports-syntax-highlighted-code/ -->
<script lang="ts">
	import Prism from "prismjs/";
	import FormEventHandler from "svelte";

	let value: string = "";

	$value: {
		let textarea = document.querySelector<HTMLTextAreaElement>("#editing");
		if (textarea) {
			value = textarea.value;
			let result = document.querySelector<HTMLSpanElement>("#highlighting-content");
			if (result) {
				result.innerHTML = value
					.replace(new RegExp("&", "g"), "&")
					.replace(new RegExp("<", "g"), "<"); /* Global RegExp */
			}
		}
	}

	function scroll() {
		/* Scroll result to scroll coords of event - sync with textarea */
		let pre = document.querySelector<HTMLPreElement>("#highlighting-content");
		let textarea = document.querySelector<HTMLTextAreaElement>("#highlighting-content");
		// Get and set x and y
		if (pre && textarea) {
			pre.scrollTop = textarea.scrollTop;
			pre.scrollLeft = textarea.scrollLeft;
		}
	}
</script>

<div class="editor">
	<textarea id="editing" spellcheck="false" bind:value on:input={scroll} on:scroll={scroll} />
	<pre id="highlighting" aria-hidden="true">
        <code class="language-html" id="highlighting-content" />
    </pre>
</div>

<style>
	#editing,
	#highlighting,
	#highlighting-content {
		font-size: 15pt;
		font-family: monospace;
		line-height: 20pt;
	}
	#editing,
	#highlighting {
		position: absolute;
		top: 0;
		left: 0;
		overflow: auto;
		white-space: nowrap;
	}
	#highlighting {
		z-index: 0;
	}
	/* Make textarea almost completely transparent */
	#editing {
		z-index: 1;
		color: transparent;
		background: transparent;
		caret-color: white; /* Or choose your favorite color */
		resize: none; /* No resize */
	}
</style>
