# Introdução A Linguagem GO

## Instalando o go.
Para instalar o compilador Go, basta fazer o download do instalador no site,
e necessário que já tenha o git.

https://golang.org/dl/

Escolha a plataforma.

## Windows
A instalção padrão. Confirme as telas e pronto.

## Linux / MacOS
Rode o commando 

``
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
``

e adicione o diretório /usr/local/go ao PATH:

No bash, uma forma rápida de fazer isso:

``
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
``

Crie um workspace, todo seu código, deve estar por convenção dentro desse diretório.

Em sistemas baseado em Unix(Linux, macOS) esse diretório é ~/go. Já em Windows, esse diretorio é %USERPROFILE%\go (C:\Users\TeuNome\go).

Crie um diretório src, dentro do workspace, é nele que escreveremos nosso código.

---
# Links Interessantes

https://www.youtube.com/watch?v=cN_DpYBzKso (Rob Pike - 'Concurrency Is Not Parallelism')

https://www.youtube.com/watch?v=CZ3wIuvmHeM (Mastering Chaos - A Netflix Guide to Microservices)