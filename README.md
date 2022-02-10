# Gamebook graphing thing

To run locally:
```
cd ~/go/src/github.com/abworrall/grapher
go run ./app/graphapp
```
Now visit [http://localhost:8080/]


To deploy into Google Cloud AppEngine:
```
cd ~/go/src/github.com/abworrall/grapher
gcloud --project=worrall-io app deploy ./app/grapherapp/dispatch.yaml
gcloud --project=worrall-io app deploy ./app/grapherapp/
```

