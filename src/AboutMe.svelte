<script>
    import Container from './Container.svelte';
    let meSrc = 'assets/me.gif';
    let goofyMeSrc = 'assets/goofy_me.gif';

    let slideNumber = 0;

    function slideRight() {
        slideNumber = slideNumber > 0 ? slideNumber-1 : 0;
    }

    function slideLeft() {
        slideNumber = slideNumber < 1 ? slideNumber+1 : 1;
    }

</script>

<style>

    .slider{
        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
        overflow: hidden;
    }

    .slider .slide-item {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        flex: 0 0 100%;
        transition: transform 250ms;
    }

    img {
        border: 1px solid #000;
        margin-right: 10px;
        flex: 0 0 100px;
        height: auto;
        width: 100px;
        height: 120px;
        padding: 2px;
    }

    .personally img {
        width: 90px;
    }

    h2 {
        margin-top: 0;
        font-size: 16px;
        flex: 0 0 100%;
    }

    .details {
        padding: 0;
        margin: 0;
        font-size: 14px;
        display: flex;
        flex-direction: column;
        flex: 1 1 0;
    }

    p {
        margin-top: 0;
        margin-bottom: 10px;
    }

    p em {
        font-weight: bold;
        display: block;
        font-style: normal;
    }

    #buttons {
        width: 100%;
        display: flex;
        flex-direction: row;
        margin-top: 10px;
    }

    button {
        font-size: 14px;
        width: 100px;
    }

    button.next {
        align-self: flex-end;
        margin-left: auto;
    }
</style>

<Container>
    <h2>About Me</h2>
    <div class="slider">
        <div class="professionally slide-item" style="transform: translateX({-100 * slideNumber}%)">
            <img src="{meSrc}"
                 alt="A low resolution picture of my face."
                 aria-label="A low resolution picture of my face."/>
            <article class="details">
                <p>
                    <em>Professionally</em>
                    I'm a 30-something User Experience Engineer with a long history of delivering high-quality software, both as a team-member and also a team-lead.
                </p>
                <p>
                    I have extensive experience in both front-end and back-end software, as well as design experience and a passion for excellence.
                </p>
                <p>
                    I've been building software for over a decade, and I've worked for companies both large and small.
                </p>
            </article>
        </div>
        <div class="personally slide-item" style="transform: translateX({-100 * slideNumber}%)">
            <img src="{goofyMeSrc}"
                 alt="A low resolution picture of my face, slightly goofier than the other one."
                 aria-label="A low resolution picture of my face, slightly goofier than the other one."/>
            <article class="details">
                <p>
                    <em>Personally</em>
                    I'm a huge nerd.  I play a fair amount of Dungeons & Dragons with my friends, and I love just about anything cyberpunk.
                </p>
                <p>
                    I brew my own beer and wine, and I drink an inordinate amount of coffee.  Most of that coffee is converted into code.
                </p>
                <p>
                    I currently have two guinea pigs that I am mildly obsessed with.  They're adorable.
                </p>
            </article>
        </div>
    </div>
    <div id="buttons">
        <button class="previous" on:click={slideRight}>⇦</button>
        <button class="next" on:click={slideLeft}>⇨</button>
    </div>
</Container>