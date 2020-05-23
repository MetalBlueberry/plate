# Plate, not just another go cli template tool.

There are several tools out there that provide similar functionality. A quick [google search](https://www.google.com/search?q=go+template+cli&oq=go+template+cli&aqs=chrome..69i57j69i64l3.2094j0j7&sourceid=chrome&ie=UTF-8) reveal that there are a lot of them out there. The difference of `Plate` vs other tools is the function `file` that allows you to write script-like templates to generate complex outputs.

## How it works

Plate reads the stdin as json format and output the executed template.

example.json

```json
{
    "field1": "value1",
    "field2": "value2"
}
```

example.tmpl

```txt
{{ range $k,$v := . -}}
Field {{ $k }} with value = {{ $v }}
{{ end -}}
```

running the following command from the folder containing both files.

```bash
plate < simple.json 
```

generates the following output

```txt
Field field1 with value = value1
Field field2 with value = value2
```

There is nothing really amazing in this behaviour. it is just a go template rendered. Let's add a little complexity.

## The "file" function

First thing is to define our default template. this will be the entry point.

example.tmpl

```txt
{{- define "default" -}}
// default template goes here
{{- end -}}
```

Now let's add template code to generate two files with different content based on this json input.

```json
{
    "file1": {
        "name": "file1.md",
        "content": "content of file 1"
    },
    "file2": {
        "name": "file2.md",
        "content": "content of file 2"
    }
}
```

```txt
{{- define "default" -}}
{{ range $k,$v := . -}}
Generating file {{ $k }} with value = {{ $v }}
{{ file $v.name "file" $v }}
{{ end -}}
{{- end -}}
```

The parameters for file function are the same as the template function from standard go templates plus a first extra parameter indicating the output file name. So in this example I'm creating a file with the name from the json data using the template "file" and passing the file data to the template. Next step is to define the template "file". You can do it in the same file or in another file within the same directory.

```txt
{{- define "default" -}}
{{ range $k,$v := . -}}
Generating file {{ $k }} with value = {{ $v }}
{{ file $v.name "file" $v }}
{{ end -}}
{{- end -}}

{{- define "file" -}}
# {{ .name }}
{{ .content }}
{{ end -}}
```

Running this code now generates the following output

```log
Generating file file1 with value = map[content:content of file 1 name:file1.md]
<nil>
Generating file file2 with value = map[content:content of file 2 name:file2.md]
<nil>
```

Also, two new files has been created within the directory.

file1.txt

```md
# file1.md
content of file 1
```

file2.txt

```md
# file2.md
content of file 2
```

The default template is acting like a log of the process. You can see that it prints the file that it's going to be generated, the content and `<nil>`. which is the return value from `file` function that as you expect, it's of the type `error`. We can easily improve the output of the log.

```txt
{{- define "default" -}}
{{ range $k,$v := . -}}
Generating file {{ $k }} with content = {{ quote $v.content }}
{{ $err := file $v.name "file" $v -}}
{{- if $err }}
Something went wrong with {{ $k }}
{{ $err }}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "file" -}}
# {{ .name }}
{{ .content }}
{{ end -}}
```

Now the output generated is easier to read and gives more info

```txt
Generating file "file1" with content = "content of file 1"
Generating file "file2" with content = "content of file 2"
```

The `quote` function is from [sprig](http://masterminds.github.io/sprig/) functions. All the spring functions has been registered and can be used in any template.

## The "stemplate" function

If you read the name of the function again, you may deduce what it does. `stemplate` > `string template` then it generates a string from a template. So yes, that's it.

In the standard go template system, the `template` function executes a template a prints the text to the current template. With `stemplate` you can execute a template and store the value in a variable to be used later.
