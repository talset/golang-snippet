### tab ###

int intarray
{{ .intarray }}
  Index number 3
  {{index .intarray 2}}

  range
  {{range .intarray}} {{.}} {{end}}

int intslice
{{ .intslice }}


Slice Of Slice string
  {{range .sliceofslice}}
    {{range .}} {{.}} {{end}}
  {{end}}

Map key/int
  {{ .map }}
  {{range .map}} {{.}} {{end}}

Struct Map
  {{ .structmap }}
  {{range .structmap}} {{.Lat}} {{.Long}} {{end}}

  {{ range $key, $value := .structmap }}
     {{ $key }}: {{ $value }}
  {{ end }}
