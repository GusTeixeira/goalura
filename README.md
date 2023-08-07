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
   
6. De forma mais simples, é possivel executar sem gerar o arquivo executável utilizando `go run`

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

## Váriaveis

Para declarar uma variável em Go, utilizamos a palavra var seguida do nome da variável mais o seu tipo. Como é um texto, o seu tipo será string, e vamos logo inicializá-la: `var nome string = "Douglas"`.

As váriaveis não tem necessidade de ser atribuídas logo na sua inicialização, elas podem ser apenas inicializadas e seu valor ser atribuido depois. Caso a váriavel não seja atribuída, ela será vazia

o Go não permite a não utilização de váriaveis, caso tenha declarado, ela tem que ser usada em alguma função.

## Armazenamento de Váriaveis

Para armazenar uma váriavel após a sua declaração, é necessário colocar um & antes para indicar o endereço da memória onde o valor indicado será alocado, exemplo:

`fmt.Scanf("%d", &comando)`

## Funções

Funções em go são simples e se parecem como qualquer outra linguagem, exemplo:

`func Teste()`

Funções podem não retornar nada e não receber nada como parametro, podem retornar porém não receber parametros ou podem retornar E receber parametros

## Funções com mais de um retorno

As funções podem ser configuradas para terem mais de um retorno,

Quando não queremos saber de um dos retornos, queremos ignorá-lo, nós utilizamos o operador de identificador em branco (_):

`  _, idade := devolveNomeEIdade()
    fmt.Println(idade)`

## Loopings

Em Go, você pode utilizar a estrutura de controle for para criar loops de repetição. Existem três tipos principais de loops que podem ser criados com a estrutura for:

1. Loop com apenas uma condição:
    
` for i := 0; i < 5; i++ {
    // Código a ser executado repetidamente
}`

Neste exemplo, o loop irá se repetir enquanto i for menor que 5. A cada iteração, o valor de i é incrementado em 1.

2. Loop com a condição de saída no meio do loop:

` for {
    // Código a ser executado repetidamente
    if condicaoDeSaida {
        break
    }
} `

Neste caso, o loop será executado indefinidamente até que a condição de saída seja atendida, momento em que o comando break será executado para sair do loop.

3. Loop utilizando um range (para percorrer elementos de um slice, array, map, string, etc.):

`slice := []int{1, 2, 3, 4, 5}
    for index, value := range slice {
        // index é o índice do elemento no slice
        // value é o valor do elemento naquele índice
    }`

Neste exemplo, o loop irá percorrer todos os elementos do slice e, a cada iteração, a variável index conterá o índice do elemento atual e a variável value conterá o valor do elemento naquele índice.

Os loops em Go funcionam de maneira similar a outras linguagens de programação, mas é importante ter cuidado para não criar loops infinitos acidentalmente, pois isso pode causar travamentos no programa. Certifique-se sempre de que haja uma condição de saída adequada em loops que não são do tipo "for { }".

## Slices e Arrays em Go

Em Go, tanto os arrays quanto os slices são utilizados para armazenar coleções de elementos do mesmo tipo. No entanto, eles têm algumas diferenças fundamentais em relação à sua natureza e funcionalidade.

### Arrays

Um array é uma coleção fixa de elementos com tamanho definido durante a sua declaração.
A sintaxe para declarar um array é a seguinte:

`var nomeDoArray [tamanho]tipo`

Os elementos de um array são acessados através de seus índices, que vão de 0 a tamanho-1.

### Slices

Um slice é uma estrutura dinâmica que representa uma seção de um array. Ele é mais flexível do que um array, pois seu tamanho pode mudar dinamicamente.
A sintaxe para declarar um slice é a seguinte:

`var nomeDoSlice []tipo`

Os elementos de um slice também são acessados através de seus índices, e a sintaxe é a mesma do array.

Slices são passados por referência em funções, o que significa que a função recebe um ponteiro para o slice original e qualquer modificação no slice dentro da função afetará o slice original.

### Principais Diferenças

O tamanho de um array é fixo e definido em tempo de compilação, enquanto o tamanho de um slice é dinâmico e pode ser alterado em tempo de execução.
Slices são mais versáteis e geralmente mais utilizados em Go, pois permitem criar estruturas de dados mais flexíveis sem se preocupar com o tamanho fixo.

Arrays são mais eficientes em termos de memória e processamento quando o tamanho é conhecido e imutável.
Exemplo de uso de slice:

`func main() {
    // Criando um slice com elementos
    numeros := []int{1, 2, 3, 4, 5}

    // Acessando elementos do slice
    fmt.Println(numeros[0]) // Saída: 1
    fmt.Println(numeros[2]) // Saída: 3

    // Adicionando elementos ao slice
    numeros = append(numeros, 6)
    fmt.Println(numeros) // Saída: [1 2 3 4 5 6]

    // Modificando elementos do slice
    numeros[0] = 10
    fmt.Println(numeros) // Saída: [10 2 3 4 5 6]

    // Criando um slice vazio e adicionando elementos posteriormente
    var outroSlice []int
    outroSlice = append(outroSlice, 100, 200, 300)
    fmt.Println(outroSlice) // Saída: [100 200 300]
}`

É importante entender as diferenças entre arrays e slices em Go e escolher a estrutura mais adequada para cada situação. Slices são mais flexíveis e são amplamente utilizados na linguagem devido à sua versatilidade e facilidade de uso.

## Contribuições e Issues:

Contribuições para a linguagem Go são bem-vindas! Caso encontre algum problema ou queira sugerir melhorias, abra uma issue no repositório oficial do projeto no GitHub (https://github.com/golang/go).

## Licença:

A linguagem Go é licenciada sob a Licença BSD. Para mais informações, consulte o arquivo de licença (LICENSE) do repositório oficial.
