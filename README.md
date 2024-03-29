# envstamp

Make docker containers easily configurable with environment variables

![](assets/envstamp.png)

This was built to make it easier to modify config files in a docker
image using environment variables

## Install

```sh
sudo curl -L -o /bin/envstamp https://github.com/jamrizzi/envstamp/releases/download/v0.2.0/envstamp
sudo chmod +x /bin/envstamp
```


## Build

```sh
git clone https://github.com/jamrizzi/envstamp.git
make build
```


## Usage

Assume you have the following template at `/my/custom/template.txt`

```
Hello ${WOLRD}
```

Run the following . . .

```sh
export WOLRD=galaxy
envstamp /my/custom/template.txt
```

All environment variables have been replaced in `/my/custom/template.txt`

```sh
cat /my/custom/template.txt

# Hello galaxy
```

To set a default value when an environment variable does not exist, use the following syntax
```
Hello ${WORLD:universe}
```
