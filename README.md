# envstamp

Stamp out templates from environment variables

This was built to make it easier to modify config files in a docker
image using environment variables


## Installation

```sh
git clone https://github.com/jamrizzi/envstamp.git
make build
sudo make install
```


## Usage

Assume you have the following template at `/my/custom/template.txt`

```
Hello ${WOLRD}
```

Run the following . . .

```sh
export WOLRD=universe
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
