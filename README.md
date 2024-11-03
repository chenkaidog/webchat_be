## run in docker
```shell
sudo docker build -t webchat:latest .
sudo docker run -itd --network=host --restart=unless-stopped -v webchat_log:/app/log webchat_be:latest
```