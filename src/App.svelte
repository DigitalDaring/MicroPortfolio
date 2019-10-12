<script>
	import Navigation from './Navigation.svelte';

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
		height: calc(100vh - 80px);
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
	}

	h1 {
		color: #000;
		font-family: 'Coda', cursive;
		font-size: 18px;
	}

	greeting {
		display: flex;
		flex-direction: row;
	}
</style>

<canvas id="dithering" width="20px" height="20px" style="display: none"/>
<div class="outer-border">
	<main>
		<header style="background-image: url({dithering})">
			Micro Portfolio
		</header>
		<greeting>
			<h1>My name is Edmund</h1>
			<Navigation></Navigation>
		</greeting>
		
		<p>The year is 2019, and websites are out of control.</p>

		<p id="about-me">
			Hallward glanced round him with a puzzled expression. The room looked as if it had not been lived in for years. A faded Flemish tapestry, a curtained picture, an old Italian cassone, and an almost empty book-case--that was all that it seemed to contain, besides a chair and a table. As Dorian Gray was lighting a half-burned candle that was standing on the mantelshelf, he saw that the whole place was covered with dust and that the carpet was in holes. A mouse ran scuffling behind the wainscoting. There was a damp odour of mildew.

			"So you think that it is only God who sees the soul, Basil? Draw that curtain back, and you will see mine."

			The voice that spoke was cold and cruel. "You are mad, Dorian, or playing a part," muttered Hallward, frowning.

			"You won't? Then I must do it myself," said the young man, and he tore the curtain from its rod and flung it on the ground.

			An exclamation of horror broke from the painter's lips as he saw in the dim light the hideous face on the canvas grinning at him. There was something in its expression that filled him with disgust and loathing. Good heavens! it was Dorian Gray's own face that he was looking at! The horror, whatever it was, had not yet entirely spoiled that marvellous beauty. There was still some gold in the thinning hair and some scarlet on the sensual mouth. The sodden eyes had kept something of the loveliness of their blue, the noble curves had not yet completely passed away from chiselled nostrils and from plastic throat. Yes, it was Dorian himself. But who had done it? He seemed to recognize his own brushwork, and the frame was his own design. The idea was monstrous, yet he felt afraid. He seized the lighted candle, and held it to the picture. In the left-hand corner was his own name, traced in long letters of bright vermilion.

			It was some foul parody, some infamous ignoble satire. He had never done that. Still, it was his own picture. He knew it, and he felt as if his blood had changed in a moment from fire to sluggish ice. His own picture! What did it mean? Why had it altered? He turned and looked at Dorian Gray with the eyes of a sick man. His mouth twitched, and his parched tongue seemed unable to articulate. He passed his hand across his forehead. It was dank with clammy sweat.

			The young man was leaning against the mantelshelf, watching him with that strange expression that one sees on the faces of those who are absorbed in a play when some great artist is acting. There was neither real sorrow in it nor real joy. There was simply the passion of the spectator, with perhaps a flicker of triumph in his eyes. He had taken the flower out of his coat, and was smelling it, or pretending to do so.
		</p>
		<p id="credentials">
			"What does this mean?" cried Hallward, at last. His own voice sounded shrill and curious in his ears.

			"Years ago, when I was a boy," said Dorian Gray, crushing the flower in his hand, "you met me, flattered me, and taught me to be vain of my good looks. One day you introduced me to a friend of yours, who explained to me the wonder of youth, and you finished a portrait of me that revealed to me the wonder of beauty. In a mad moment that, even now, I don't know whether I regret or not, I made a wish, perhaps you would call it a prayer...."

			"I remember it! Oh, how well I remember it! No! the thing is impossible. The room is damp. Mildew has got into the canvas. The paints I used had some wretched mineral poison in them. I tell you the thing is impossible."

			"Ah, what is impossible?" murmured the young man, going over to the window and leaning his forehead against the cold, mist-stained glass.

			"You told me you had destroyed it."

			"I was wrong. It has destroyed me."

			"I don't believe it is my picture."
		</p>
	</main>
</div>

