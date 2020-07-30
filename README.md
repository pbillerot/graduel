# GRADUEL

Framework de développement d'application WEB en GO

**G Rad uel**
- **G** comme Go le langage que je vais pratiquer pour la 1ère fois
- **rad** comme "Rapid Application Development"
  pour développer un CRUD "Create Read Update Delete"
  Utilisation d'un dictionnaire JSON pour déclarer 
  - les rubriques, les vues et formulaires du CRUD
- tout ça progressivement -> grad**uel**ement

Application qui sera mise ensuite dans un container DOCKER - projet https://github.com/pbillerot/graduel-docker

## Installation de GO sur ma Debian Buster
```console
cd ~/Téléchargements
wget https://golang.org/dl/go1.14.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
```
##### dans .profil
```console
export PATH=$PATH:/usr/local/go/bin
export GOPATH=~/go
export SESSION_KEY=<key 32>
export SECRET_KEY=<key 32>
```
