<script>
	import Navigation from './Navigation.svelte';
	import AboutMe from './AboutMe.svelte';
	import MyCredentials from './MyCredentials.svelte';

	export let dithering;

	window.setTimeout(() => {
        const canvas = document.getElementById('dithering');
        const context = canvas.getContext("2d");
        context.fillStyle = 'rgba(0, 0, 0, .6)';
        for(var x = 0; x < 20; x++) {
            for(var y = 0; y < 20; y++) {
                if((x + y) % 2 == 0) {
                    context.fillRect(x, y, 1, 1)
                }
            }
        }

        dithering = canvas.toDataURL('image/gif');
	});
</script>

<style>

	main {
		border: 8px solid #000;
		padding: 10px;
		border-radius: 10px;
		font-family: 'Coda', cursive;
		position: relative;
		height: calc(100vh - 160px);
		overflow: scroll;
	}

	.outer-border {
		border: 2px solid #000;
		border-radius: 10px;
		padding: 1px;
		width: 100%;
		max-width: 600px;
	}

	header {
		margin-top: -10px;
		margin-left: -10px;
		width: calc(100%);
		color: #FFF;
		padding-left: 10px;
		padding: 10px;
		background-color: #FFF;
		position: sticky;
		top: -10px;
		margin-bottom: 10px;
		align-items: center;
	}

	h1 {
		color: #000;
		font-family: 'Coda', cursive;
		font-size: 18px;
		display: inline-block;
		height: 32px;
	}

	header {
		display: flex;
		flex-direction: row;
		z-index: 101;
	}

	@media screen and (max-width: 480px) {
		h1 {
			font-size: 14px;
		}

		a {
			font-size: 12px;
		}
	}
</style>

<canvas id="dithering" width="20px" height="20px" style="display: none"/>
<div class="outer-border">
	<main>
		<header style="background-image: url({dithering})">
			Micro Portfolio
			<Navigation></Navigation>
		</header>

		<h1 id="about-me" >My name is Edmund.&nbsp;&nbsp;This is my micro-portfolio.</h1>
		<AboutMe></AboutMe>
		<MyCredentials></MyCredentials>
	</main>
</div>

