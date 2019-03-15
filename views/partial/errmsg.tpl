{{ with .Error }}
<div class="row margin">
    <div class="input-field col s12">
    <p class="center red-text"> {{ . }}</p>
    </div>
</div>
{{ end }}
