# works-hai-backend
## how to use
```sh
make
./main
```
### generate presigned url in terminal
google CLI を[インストール](https://cloud.google.com/sdk/docs/install?hl=ja)
```sh
pip install pyopenssl
export CLOUDSDK_PYTHON_SITEPACKAGES=1
gcloud init
gcloud storage sign-url -m PUT gs://BUCKET_NAME/IMAGE_NAME.png --private-key-file=./config/path/to/envkey.json --duration=10m
```