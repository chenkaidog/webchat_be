## run in docker
```shell
sudo docker build -t webchat_be:latest .
sudo docker run -itd --network=host --restart=unless-stopped -v webchat_log:/app/log webchat_be:latest
```