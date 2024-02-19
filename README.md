# Picsum-SDK

Picsum GO (Golang) SDK for generating random images

## How to install

```shell
go get github.com/Hand-of-Doom/Picsum-SDK@latest
```

## Examples

### Generating a random image

```go
provider := picsum.NewImageProvider(
    picsum.NewImageOptions(500, 500), // width & height
)

image, err := provider.Load()
fmt.Println(image.ID) // id of that image

os.WriteFile("./output.jpg", image.Content, 0400)
```

### Getting an image by seed

```go
options := picsum.NewImageOptions(500, 500).
    Identifier("apple", picsum.Seed)

provider := picsum.NewImageProvider(options)
image, err := provider.Load()

os.WriteFile("./output.jpg", image.Content, 0400)
```

### Getting an image by id

```go
options := picsum.NewImageOptions(500, 500).
    Identifier("1", picsum.ID)

provider := picsum.NewImageProvider(options)
image, err := provider.Load()

os.WriteFile("./output.jpg", image.Content, 0400)
```

### Getting an image with blur & grayscale filter

```go
options := picsum.NewImageOptions(500, 500)
// You can set blur depth to zero to disable blur effect
options.Filter(
    /*is grayscale enabled*/ true, 
    /*blur depth*/ 5,
)

provider := picsum.NewImageProvider(options)
image, err := provider.Load()

os.WriteFile("./output.jpg", image.Content, 0400)
```

### Getting images in different formats

```go
optionsWebP := picsum.NewImageOptions(500, 500).Format(picsum.WebP)

optionsJPEG := picsum.NewImageOptions(500, 500).Format(picsum.JPG)

imageWebP, err := picsum.NewImageProvider(optionsWebP).Load()
imageJPEG, err := picsum.NewImageProvider(optionsJPEG).Load()

os.WriteFile("./output.webp", imageWebP.Content, 0400)
os.WriteFile("./output.jpg", imageJPEG.Content, 0400)
```

### Getting image details
```go
image, err := picsum.
    NewImageProvider(picsum.NewImageOptions(500, 500)).
    Load()

thisImageID := image.ID

details, err := picsum.GetImageDetails(thisImageID, picsum.ID)

// By seed
details, err := picsum.GetImageDetails("apple", picsum.Seed)
```

### Getting a list of images

```go
page := 1
limit := 100
list, err := picsum.GetImagesList(page, limit)

if list.LastPage {
    fmt.Println("Oh, this is a last page!")
}

fmt.Println(list.Value)
```
