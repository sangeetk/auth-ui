{{ with .Flash }}
    {{ if .success }}
    <p class="center green-text"> {{ .success }}</p>
    {{ end }}
    {{ if .notice }}
    <p class="center blue-text"> {{ .notice }}</p>
    {{ end }}
    {{ if .warning }}
    <p calss="center yellow-text"> {{ .warning }}</p>
    {{ end }}
    {{ if .error }}
    <p class="center red-text"> {{ .error }}</p>
    {{ end }}
{{ end }}
