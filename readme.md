# GoResize

Goresize is an small and simple script to resize all images inside a folder, including subfolder. 

## Usage

You can set the desired width and height as well as the quality using the following arguments:

| Flag | Type   | Default | Explanation                                                           |
|------|--------|---------|-----------------------------------------------------------------------|
| -d   | string | .       | Directory where to resize images                                      |
| -w   | int    | 1224    | Desired image width, in pixels. Setting it to 0 will keep the ratio.  |
| -h   | int    | 0       | Desired image height, in pixels. Setting it to 0 will keep the ratio. |
| -q   | int    | 80      | Desired image quality, from 0 to 100 (lower to better)                |

```
// This will resize all images inside mydir to 400x400px and save them with a quality of 70
$ goresize -d=mydir -w=400 -h=400 q=90

// This will resize all images inside mydir to a width of 400px maintaining the original ratio and save them with a quality of 90
$ goresize -d=mydir -w=400 q=70
```

## Image quality

Image quality argumentw will be applied only to images of type JPG or PNG.

For JPG images, the indicated ratio will be applied. 

For JPG images, the ratio will be converted to the PNG compression level of the (image/png)[https://golang.org/pkg/image/png/#CompressionLevel] go package, as follows:

| 100     | NoCompression      |
|---------|--------------------|
| 80 - 99 | DefaultCompression |
| 60 - 79 | BestSpeed          |
| 0 - 60  | BestCompression    |


## Contributåing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.