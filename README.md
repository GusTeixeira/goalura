# README - Go (Linguagem de Programação)

Este é o repositório destinado aos meus estudos da linguagem de programação Go, também conhecida como Golang. Criada pelos engenheiros do Google em 2009, o Go surgiu para resolver os problemas de compilação complicados e demorados enfrentados pela empresa, devido ao uso extensivo de C++ e C em seus sistemas.

## Características Principais:

1. **Rápida Compilação:** Uma das principais vantagens do Go é o rápido processo de compilação de programas, que resolveu o problema enfrentado pelo Google.

2. **Modularização:** O Go incentiva a modularização de funcionalidades em pacotes, que podem ser facilmente importados em aplicações conforme necessário.

3. **Sintaxe Simples:** Com apenas cerca de 25 palavras-chave, a sintaxe do Go é curta e fácil de entender, tornando o código simples e legível.

4. **Convenções Rápidas de Desenvolvimento:** O Go foi projetado para agilizar o desenvolvimento, com convenções que permitem que os programadores se concentrem no código em vez de discussões sobre como deve ser feito.

## Instalação e Uso:

1. Acesse o site oficial do Go (https://golang.org/) e siga as instruções para baixar e instalar a versão mais recente da linguagem para o seu sistema operacional.

2. Verifique se a instalação foi bem-sucedida executando o comando `go version` no terminal.

3. Crie e escreva seus programas Go usando seu editor de código preferido.

4. Compile o programa usando o comando `go build` seguido pelo nome do arquivo Go. O executável será gerado na mesma pasta.

5. Execute o programa compilado usando `./nome-do-executavel`.

## Go Workspace

No momento da instalação é definido um local onde ficará todo o nosso código Go, esse local é o Go Workspace. Por padrão essa pasta deve ser chamada go e deve ficar na raiz do nosso usuário. Isso vale para todos os sistemas operacionais, ou seja, na pasta do nosso usuário, nós criamos a pasta go.

Além disso, o Go Workspace deve possuir três pastas. A primeira delas é a pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando no nosso código, veremos isso mais à frente. Além da pasta pkg, também deve ter a pasta src, onde escreveremos o código fonte de cada aplicação, e a pasta bin, onde ficará os compilados do nosso código Go.

A estrutura final do Go Workspace ficará assim:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src

Mas se quisermos que o Go Workspace seja em outro lugar, pode? Pode ser, mas por padrão o instalador do Go já espera que essa pasta esteja na raiz do usuário, então, para nos facilitar, vamos seguir os padrões da linguagem.

## Contribuições e Issues:

Contribuições para a linguagem Go são bem-vindas! Caso encontre algum problema ou queira sugerir melhorias, abra uma issue no repositório oficial do projeto no GitHub (https://github.com/golang/go).

## Licença:

A linguagem Go é licenciada sob a Licença BSD. Para mais informações, consulte o arquivo de licença (LICENSE) do repositório oficial.
