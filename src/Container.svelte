<script>
</script>

<style>
    div.group-box {
        padding: 10px;
        border: 2px solid #000;
        background-color: #FFF;
        overflow: hidden;
        margin-top: -6px;
        margin-left: -2px;
    }

    div.shadow-box {
        padding-top: 2px;
        margin-left: 2px;
        border-right: 3px solid #000;
        border-bottom: 3px solid #000;
    }
</style>

<div class="shadow-box">
    <div class="group-box">
        <slot></slot>
    </div>
</div>