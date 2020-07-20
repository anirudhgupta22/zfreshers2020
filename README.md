# zfreshers2020

This is a sample project with a small webapp in golang


#### Build the image
```
docker build --no-cache --tag hello-app .
```

#### Search for the image
```
docker images
```

#### Run the container
```
docker run -p "10000:8080" -it hello-app
```


#### Test on your browser
```
0.0.0.0:10000/hello
```
